package server

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"cloud.google.com/go/bigtable"
	pb "github.com/datacommonsorg/reconciliation/internal/proto"
	"github.com/datacommonsorg/reconciliation/internal/util"
	"github.com/golang/geo/s2"
	"google.golang.org/protobuf/encoding/protojson"
)

const (
	gridSize         float64 = 0.2
	geoJsonPredicate string  = "geoJsonCoordinates"
)

// TODO(Spaceenter): Also add place types to the results.

// ResolveCoordinates implements API for ReconServer.ResolveCoordinates.
func (s *Server) ResolveCoordinates(ctx context.Context, in *pb.ResolveCoordinatesRequest) (
	*pb.ResolveCoordinatesResponse, error) {
	// Map: lat^lng => normalized lat^lng.
	normCoordinateMap := map[string]string{}
	coordinateLookupKeys := map[string]struct{}{}

	// Read request.
	for _, coordinate := range in.GetCoordinates() {
		nKey := normalizedCoordinateKey(coordinate)
		normCoordinateMap[coordinateKey(coordinate)] = nKey
		coordinateLookupKeys[nKey] = struct{}{}
	}

	// Read coordinate recon cache.
	reconRowList := bigtable.RowList{}
	for key := range coordinateLookupKeys {
		reconRowList = append(reconRowList,
			fmt.Sprintf("%s%s", util.BtCoordinateReconPrefix, key))
	}
	reconDataMap, err := bigTableReadRowsParallel(ctx, s.btTable, reconRowList,
		func(rowKey string) (string, error) {
			return strings.TrimPrefix(rowKey, util.BtCoordinateReconPrefix), nil
		},
		func(dcid string, jsonRaw []byte) (interface{}, error) {
			var recon pb.CoordinateRecon
			if err := protojson.Unmarshal(jsonRaw, &recon); err != nil {
				return nil, err
			}
			return &recon, nil
		})
	if err != nil {
		return nil, err
	}

	// Collect places that don't fully cover the tiles that the coordinates are in.
	questionablePlaces := map[string]struct{}{}
	for _, recon := range reconDataMap {
		for _, place := range recon.(*pb.CoordinateRecon).GetPlaces() {
			if !place.GetFull() {
				questionablePlaces[place.GetDcid()] = struct{}{}
			}
		}
	}

	// Read place GeoJson cache.
	geoJsonRowList := bigtable.RowList{}
	for place := range questionablePlaces {
		geoJsonRowList = append(geoJsonRowList,
			fmt.Sprintf("%s%s^%s", util.BtOutPropValPrefix, place, geoJsonPredicate))
	}
	geoJsonDataMap, err := bigTableReadRowsParallel(ctx, s.btTable, geoJsonRowList,
		func(rowKey string) (string, error) {
			l := strings.TrimPrefix(rowKey, util.BtOutPropValPrefix)
			return strings.TrimSuffix(l, fmt.Sprintf("^%s", geoJsonPredicate)), nil
		},
		func(dcid string, jsonRaw []byte) (interface{}, error) {
			var info pb.EntityInfoCollection
			if err := protojson.Unmarshal(jsonRaw, &info); err != nil {
				return nil, err
			}
			return &info, nil
		})
	if err != nil {
		return nil, err
	}
	geoJsonMap := map[string]string{}
	for place, info := range geoJsonDataMap {
		// A place should only have a single geoJsonCooridnates out arc.
		typedInfo := info.(*pb.EntityInfoCollection)
		if typedInfo.GetTotalCount() != 1 {
			continue
		}
		geoJsonMap[place] = typedInfo.GetEntities()[0].GetValue()
	}

	// Assemble response.
	res := &pb.ResolveCoordinatesResponse{}
	for _, co := range in.GetCoordinates() {
		nKey := normCoordinateMap[coordinateKey(co)]

		recon, ok := reconDataMap[nKey]
		if !ok {
			continue
		}

		placeCoordinates := &pb.ResolveCoordinatesResponse_PlaceCoordinate{}
		for _, place := range recon.(*pb.CoordinateRecon).GetPlaces() {
			if place.GetFull() {
				placeCoordinates.PlaceDcids = append(placeCoordinates.PlaceDcids,
					place.GetDcid())
			} else { // Not fully cover the tile.
				geoJson, ok := geoJsonMap[place.GetDcid()]
				if !ok {
					continue
				}
				contained, err := isContainedIn(geoJson,
					co.GetLatitude(), co.GetLongitude())
				if err != nil {
					return res, err
				}
				if contained {
					placeCoordinates.PlaceDcids = append(placeCoordinates.PlaceDcids,
						place.GetDcid())
				}
			}
		}

		res.PlaceCoordinates = append(res.PlaceCoordinates, placeCoordinates)
	}

	return res, nil
}

func coordinateKey(c *pb.ResolveCoordinatesRequest_Coordinate) string {
	return fmt.Sprintf("%f^%f", c.GetLatitude(), c.GetLongitude())
}

func normalizedCoordinateKey(c *pb.ResolveCoordinatesRequest_Coordinate) string {
	// Normalize to South-West of the grid points.
	lat := float64(int((c.GetLatitude()+90.0)/gridSize))*gridSize - 90
	lng := float64(int((c.GetLongitude()+180.0)/gridSize))*gridSize - 180
	return fmt.Sprintf("%.1f^%.1f", lat, lng)
}

type Polygon struct {
	Loops [][][]float64
}

type MultiPolygon struct {
	Polygons [][][][]float64
}

type GeoJson struct {
	Type         string          `json:"type"`
	Coordinates  json.RawMessage `json:"coordinates"`
	Polygon      Polygon         `json:"-"`
	MultiPolygon MultiPolygon    `json:"-"`
}

func buildS2Loops(loops [][][]float64) ([]*s2.Loop, error) {
	res := []*s2.Loop{}
	for i, loop := range loops {
		if l := len(loop); l < 4 {
			return nil, fmt.Errorf("geoJson requires >= 4 points for a loop, got %d", l)
		}

		s2Points := []s2.Point{}
		// NOTE: We have to skip the last point when constructing the s2Loop.
		// In GeoJson, the last point is the same as the first point for a loop.
		// If not skipping, it sometimes leads to wrong result for containment calculation.
		for _, point := range loop[:len(loop)-1] {
			if len(point) != 2 {
				return nil, fmt.Errorf("wrong point format: %+v", point)
			}
			// NOTE: GeoJson has longitude comes before latitude.
			s2Points = append(s2Points,
				s2.PointFromLatLng(s2.LatLngFromDegrees(point[1], point[0])))
		}
		s2Loop := s2.LoopFromPoints(s2Points)
		if i == 0 {
			// The first ring of a polygon is a shell, it should be normalized to counter-clockwise.
			//
			// This step ensures that the planar polygon loop follows the "right-hand rule"
			// and reverses the orientation when that is not the case. This is specified by
			// RFC 7946 GeoJSON spec (https://tools.ietf.org/html/rfc7946), but is commonly
			// disregarded. Since orientation is easy to deduce on the plane, we assume the
			// obvious orientation is intended. We reverse orientation to ensure that all
			// loops follow the right-hand rule. This corresponds to S2's "interior-on-the-
			// left rule", and allows us to create these polygon as oriented S2 polygons.
			//
			// Also see https://en.wikipedia.org/wiki/Curve_orientation.
			s2Loop.Normalize()
		}
		res = append(res, s2Loop)
	}
	return res, nil
}

func parseGeoJson(geoJson string) (*s2.Polygon, error) {
	g := &GeoJson{}
	if err := json.Unmarshal([]byte(geoJson), g); err != nil {
		return nil, err
	}

	switch g.Type {
	case "Polygon":
		if err := json.Unmarshal(g.Coordinates, &g.Polygon.Loops); err != nil {
			return nil, err
		}
		s2Loops, err := buildS2Loops(g.Polygon.Loops)
		if err != nil {
			return nil, err
		}
		return s2.PolygonFromOrientedLoops(s2Loops), nil
	case "MultiPolygon":
		if err := json.Unmarshal(g.Coordinates, &g.MultiPolygon.Polygons); err != nil {
			return nil, err
		}
		s2Loops := []*s2.Loop{}
		for _, polygon := range g.MultiPolygon.Polygons {
			lps, err := buildS2Loops(polygon)
			if err != nil {
				return nil, err
			}
			s2Loops = append(s2Loops, lps...)
		}
		return s2.PolygonFromOrientedLoops(s2Loops), nil
	default:
		return nil, fmt.Errorf("unrecognized GeoJson object: %+v", g.Type)
	}
}

func isContainedIn(geoJson string, lat float64, lng float64) (bool, error) {
	s2Polygon, err := parseGeoJson(geoJson)
	if err != nil {
		return false, err
	}
	s2Point := s2.PointFromLatLng(s2.LatLngFromDegrees(lat, lng))
	return s2Polygon.ContainsPoint(s2Point), nil
}
