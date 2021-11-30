package server

import (
	"io/ioutil"
	"path"
	"runtime"
	"testing"
)

func TestIsContainedIn(t *testing.T) {
	for _, c := range []struct {
		geoJsonFileName   string
		lat               float64
		lng               float64
		wantIsContainedIn bool
	}{
		{
			"mountain_view_geo_json.json",
			37.42,
			-122.08,
			true,
		},
		{
			"mexico_geo_json.json",
			32.41,
			-102.11,
			false,
		},
		{
			"mexico_geo_json.json",
			26.55,
			-102.85,
			true,
		},
	} {
		_, filename, _, _ := runtime.Caller(0)
		geoJsonFilePath := path.Join(
			path.Dir(filename), "test_data", c.geoJsonFileName)
		geoJsonBytes, err := ioutil.ReadFile(geoJsonFilePath)
		if err != nil {
			t.Errorf("ioutil.ReadFile(%s) = %s", c.geoJsonFileName, err)
			continue
		}
		contained, err := isContainedIn(string(geoJsonBytes), c.lat, c.lng)
		if err != nil {
			t.Errorf("isContainedIn(%s) = %s", c.geoJsonFileName, err)
			continue
		}
		if contained != c.wantIsContainedIn {
			t.Errorf("isContainedIn(%s) = %t, want %t",
				c.geoJsonFileName, contained, c.wantIsContainedIn)
		}
	}
}
