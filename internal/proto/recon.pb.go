// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0-devel
// 	protoc        v3.14.0
// source: recon.proto

package proto

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type IdWithProperty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Prop *string `protobuf:"bytes,1,opt,name=prop" json:"prop,omitempty"`
	Val  *string `protobuf:"bytes,2,opt,name=val" json:"val,omitempty"`
}

func (x *IdWithProperty) Reset() {
	*x = IdWithProperty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recon_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IdWithProperty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IdWithProperty) ProtoMessage() {}

func (x *IdWithProperty) ProtoReflect() protoreflect.Message {
	mi := &file_recon_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IdWithProperty.ProtoReflect.Descriptor instead.
func (*IdWithProperty) Descriptor() ([]byte, []int) {
	return file_recon_proto_rawDescGZIP(), []int{0}
}

func (x *IdWithProperty) GetProp() string {
	if x != nil && x.Prop != nil {
		return *x.Prop
	}
	return ""
}

func (x *IdWithProperty) GetVal() string {
	if x != nil && x.Val != nil {
		return *x.Val
	}
	return ""
}

type EntityIds struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids []*IdWithProperty `protobuf:"bytes,1,rep,name=ids" json:"ids,omitempty"`
}

func (x *EntityIds) Reset() {
	*x = EntityIds{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recon_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntityIds) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntityIds) ProtoMessage() {}

func (x *EntityIds) ProtoReflect() protoreflect.Message {
	mi := &file_recon_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntityIds.ProtoReflect.Descriptor instead.
func (*EntityIds) Descriptor() ([]byte, []int) {
	return file_recon_proto_rawDescGZIP(), []int{1}
}

func (x *EntityIds) GetIds() []*IdWithProperty {
	if x != nil {
		return x.Ids
	}
	return nil
}

// An entity is represented by a subgraph, which contains itself and its neighbors.
type EntitySubGraph struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// REQUIRED: source_id must be a key within `sub_graph.nodes`, or one of the `ids`.
	SourceId *string `protobuf:"bytes,1,opt,name=source_id,json=sourceId" json:"source_id,omitempty"`
	// Types that are assignable to GraphRepresentation:
	//	*EntitySubGraph_SubGraph
	//	*EntitySubGraph_EntityIds
	GraphRepresentation isEntitySubGraph_GraphRepresentation `protobuf_oneof:"graph_representation"`
}

func (x *EntitySubGraph) Reset() {
	*x = EntitySubGraph{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recon_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntitySubGraph) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntitySubGraph) ProtoMessage() {}

func (x *EntitySubGraph) ProtoReflect() protoreflect.Message {
	mi := &file_recon_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntitySubGraph.ProtoReflect.Descriptor instead.
func (*EntitySubGraph) Descriptor() ([]byte, []int) {
	return file_recon_proto_rawDescGZIP(), []int{2}
}

func (x *EntitySubGraph) GetSourceId() string {
	if x != nil && x.SourceId != nil {
		return *x.SourceId
	}
	return ""
}

func (m *EntitySubGraph) GetGraphRepresentation() isEntitySubGraph_GraphRepresentation {
	if m != nil {
		return m.GraphRepresentation
	}
	return nil
}

func (x *EntitySubGraph) GetSubGraph() *McfGraph {
	if x, ok := x.GetGraphRepresentation().(*EntitySubGraph_SubGraph); ok {
		return x.SubGraph
	}
	return nil
}

func (x *EntitySubGraph) GetEntityIds() *EntityIds {
	if x, ok := x.GetGraphRepresentation().(*EntitySubGraph_EntityIds); ok {
		return x.EntityIds
	}
	return nil
}

type isEntitySubGraph_GraphRepresentation interface {
	isEntitySubGraph_GraphRepresentation()
}

type EntitySubGraph_SubGraph struct {
	SubGraph *McfGraph `protobuf:"bytes,2,opt,name=sub_graph,json=subGraph,oneof"`
}

type EntitySubGraph_EntityIds struct {
	EntityIds *EntityIds `protobuf:"bytes,3,opt,name=entity_ids,json=entityIds,oneof"`
}

func (*EntitySubGraph_SubGraph) isEntitySubGraph_GraphRepresentation() {}

func (*EntitySubGraph_EntityIds) isEntitySubGraph_GraphRepresentation() {}

type EntityPair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntityOne *EntitySubGraph `protobuf:"bytes,1,opt,name=entity_one,json=entityOne" json:"entity_one,omitempty"`
	EntityTwo *EntitySubGraph `protobuf:"bytes,2,opt,name=entity_two,json=entityTwo" json:"entity_two,omitempty"`
}

func (x *EntityPair) Reset() {
	*x = EntityPair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recon_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EntityPair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EntityPair) ProtoMessage() {}

func (x *EntityPair) ProtoReflect() protoreflect.Message {
	mi := &file_recon_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EntityPair.ProtoReflect.Descriptor instead.
func (*EntityPair) Descriptor() ([]byte, []int) {
	return file_recon_proto_rawDescGZIP(), []int{3}
}

func (x *EntityPair) GetEntityOne() *EntitySubGraph {
	if x != nil {
		return x.EntityOne
	}
	return nil
}

func (x *EntityPair) GetEntityTwo() *EntitySubGraph {
	if x != nil {
		return x.EntityTwo
	}
	return nil
}

type CompareEntitiesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EntityPairs []*EntityPair `protobuf:"bytes,1,rep,name=entity_pairs,json=entityPairs" json:"entity_pairs,omitempty"`
}

func (x *CompareEntitiesRequest) Reset() {
	*x = CompareEntitiesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recon_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompareEntitiesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompareEntitiesRequest) ProtoMessage() {}

func (x *CompareEntitiesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_recon_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompareEntitiesRequest.ProtoReflect.Descriptor instead.
func (*CompareEntitiesRequest) Descriptor() ([]byte, []int) {
	return file_recon_proto_rawDescGZIP(), []int{4}
}

func (x *CompareEntitiesRequest) GetEntityPairs() []*EntityPair {
	if x != nil {
		return x.EntityPairs
	}
	return nil
}

type CompareEntitiesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Comparisons []*CompareEntitiesResponse_Comparison `protobuf:"bytes,1,rep,name=comparisons" json:"comparisons,omitempty"`
}

func (x *CompareEntitiesResponse) Reset() {
	*x = CompareEntitiesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recon_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompareEntitiesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompareEntitiesResponse) ProtoMessage() {}

func (x *CompareEntitiesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_recon_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompareEntitiesResponse.ProtoReflect.Descriptor instead.
func (*CompareEntitiesResponse) Descriptor() ([]byte, []int) {
	return file_recon_proto_rawDescGZIP(), []int{5}
}

func (x *CompareEntitiesResponse) GetComparisons() []*CompareEntitiesResponse_Comparison {
	if x != nil {
		return x.Comparisons
	}
	return nil
}

type ResolveEntitiesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Entities []*EntitySubGraph `protobuf:"bytes,1,rep,name=entities" json:"entities,omitempty"`
	// The properties of IDs to find. If empty, all known IDs are returned.
	WantedIdProperties []string `protobuf:"bytes,2,rep,name=wanted_id_properties,json=wantedIdProperties" json:"wanted_id_properties,omitempty"`
}

func (x *ResolveEntitiesRequest) Reset() {
	*x = ResolveEntitiesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recon_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResolveEntitiesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResolveEntitiesRequest) ProtoMessage() {}

func (x *ResolveEntitiesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_recon_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResolveEntitiesRequest.ProtoReflect.Descriptor instead.
func (*ResolveEntitiesRequest) Descriptor() ([]byte, []int) {
	return file_recon_proto_rawDescGZIP(), []int{6}
}

func (x *ResolveEntitiesRequest) GetEntities() []*EntitySubGraph {
	if x != nil {
		return x.Entities
	}
	return nil
}

func (x *ResolveEntitiesRequest) GetWantedIdProperties() []string {
	if x != nil {
		return x.WantedIdProperties
	}
	return nil
}

type ResolveEntitiesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ResolvedEntities []*ResolveEntitiesResponse_ResolvedEntity `protobuf:"bytes,1,rep,name=resolved_entities,json=resolvedEntities" json:"resolved_entities,omitempty"`
}

func (x *ResolveEntitiesResponse) Reset() {
	*x = ResolveEntitiesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recon_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResolveEntitiesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResolveEntitiesResponse) ProtoMessage() {}

func (x *ResolveEntitiesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_recon_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResolveEntitiesResponse.ProtoReflect.Descriptor instead.
func (*ResolveEntitiesResponse) Descriptor() ([]byte, []int) {
	return file_recon_proto_rawDescGZIP(), []int{7}
}

func (x *ResolveEntitiesResponse) GetResolvedEntities() []*ResolveEntitiesResponse_ResolvedEntity {
	if x != nil {
		return x.ResolvedEntities
	}
	return nil
}

type CompareEntitiesResponse_Comparison struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SourceIdOne *string  `protobuf:"bytes,1,opt,name=source_id_one,json=sourceIdOne" json:"source_id_one,omitempty"`
	SourceIdTwo *string  `protobuf:"bytes,2,opt,name=source_id_two,json=sourceIdTwo" json:"source_id_two,omitempty"`
	Probability *float64 `protobuf:"fixed64,3,opt,name=probability" json:"probability,omitempty"`
}

func (x *CompareEntitiesResponse_Comparison) Reset() {
	*x = CompareEntitiesResponse_Comparison{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recon_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompareEntitiesResponse_Comparison) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompareEntitiesResponse_Comparison) ProtoMessage() {}

func (x *CompareEntitiesResponse_Comparison) ProtoReflect() protoreflect.Message {
	mi := &file_recon_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompareEntitiesResponse_Comparison.ProtoReflect.Descriptor instead.
func (*CompareEntitiesResponse_Comparison) Descriptor() ([]byte, []int) {
	return file_recon_proto_rawDescGZIP(), []int{5, 0}
}

func (x *CompareEntitiesResponse_Comparison) GetSourceIdOne() string {
	if x != nil && x.SourceIdOne != nil {
		return *x.SourceIdOne
	}
	return ""
}

func (x *CompareEntitiesResponse_Comparison) GetSourceIdTwo() string {
	if x != nil && x.SourceIdTwo != nil {
		return *x.SourceIdTwo
	}
	return ""
}

func (x *CompareEntitiesResponse_Comparison) GetProbability() float64 {
	if x != nil && x.Probability != nil {
		return *x.Probability
	}
	return 0
}

type ResolveEntitiesResponse_ResolvedId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ids         []*IdWithProperty `protobuf:"bytes,1,rep,name=ids" json:"ids,omitempty"`
	Probability *float64          `protobuf:"fixed64,2,opt,name=probability" json:"probability,omitempty"`
}

func (x *ResolveEntitiesResponse_ResolvedId) Reset() {
	*x = ResolveEntitiesResponse_ResolvedId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recon_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResolveEntitiesResponse_ResolvedId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResolveEntitiesResponse_ResolvedId) ProtoMessage() {}

func (x *ResolveEntitiesResponse_ResolvedId) ProtoReflect() protoreflect.Message {
	mi := &file_recon_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResolveEntitiesResponse_ResolvedId.ProtoReflect.Descriptor instead.
func (*ResolveEntitiesResponse_ResolvedId) Descriptor() ([]byte, []int) {
	return file_recon_proto_rawDescGZIP(), []int{7, 0}
}

func (x *ResolveEntitiesResponse_ResolvedId) GetIds() []*IdWithProperty {
	if x != nil {
		return x.Ids
	}
	return nil
}

func (x *ResolveEntitiesResponse_ResolvedId) GetProbability() float64 {
	if x != nil && x.Probability != nil {
		return *x.Probability
	}
	return 0
}

type ResolveEntitiesResponse_ResolvedEntity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SourceId    *string                               `protobuf:"bytes,1,opt,name=source_id,json=sourceId" json:"source_id,omitempty"`
	ResolvedIds []*ResolveEntitiesResponse_ResolvedId `protobuf:"bytes,2,rep,name=resolved_ids,json=resolvedIds" json:"resolved_ids,omitempty"`
}

func (x *ResolveEntitiesResponse_ResolvedEntity) Reset() {
	*x = ResolveEntitiesResponse_ResolvedEntity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_recon_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ResolveEntitiesResponse_ResolvedEntity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ResolveEntitiesResponse_ResolvedEntity) ProtoMessage() {}

func (x *ResolveEntitiesResponse_ResolvedEntity) ProtoReflect() protoreflect.Message {
	mi := &file_recon_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ResolveEntitiesResponse_ResolvedEntity.ProtoReflect.Descriptor instead.
func (*ResolveEntitiesResponse_ResolvedEntity) Descriptor() ([]byte, []int) {
	return file_recon_proto_rawDescGZIP(), []int{7, 1}
}

func (x *ResolveEntitiesResponse_ResolvedEntity) GetSourceId() string {
	if x != nil && x.SourceId != nil {
		return *x.SourceId
	}
	return ""
}

func (x *ResolveEntitiesResponse_ResolvedEntity) GetResolvedIds() []*ResolveEntitiesResponse_ResolvedId {
	if x != nil {
		return x.ResolvedIds
	}
	return nil
}

var File_recon_proto protoreflect.FileDescriptor

var file_recon_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x72, 0x65, 0x63, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x64,
	0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x09, 0x4d, 0x63, 0x66, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x36, 0x0a, 0x0e, 0x49, 0x64, 0x57, 0x69, 0x74, 0x68, 0x50, 0x72, 0x6f,
	0x70, 0x65, 0x72, 0x74, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x72, 0x6f, 0x70, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x70, 0x72, 0x6f, 0x70, 0x12, 0x10, 0x0a, 0x03, 0x76, 0x61, 0x6c,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x76, 0x61, 0x6c, 0x22, 0x3a, 0x0a, 0x09, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x73, 0x12, 0x2d, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x2e, 0x49, 0x64, 0x57, 0x69, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72,
	0x74, 0x79, 0x52, 0x03, 0x69, 0x64, 0x73, 0x22, 0xb4, 0x01, 0x0a, 0x0e, 0x45, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x53, 0x75, 0x62, 0x47, 0x72, 0x61, 0x70, 0x68, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73,
	0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x34, 0x0a, 0x09, 0x73, 0x75, 0x62, 0x5f, 0x67,
	0x72, 0x61, 0x70, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x64, 0x61, 0x74,
	0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x4d, 0x63, 0x66, 0x47, 0x72, 0x61, 0x70,
	0x68, 0x48, 0x00, 0x52, 0x08, 0x73, 0x75, 0x62, 0x47, 0x72, 0x61, 0x70, 0x68, 0x12, 0x37, 0x0a,
	0x0a, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x16, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e,
	0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x49, 0x64, 0x73, 0x48, 0x00, 0x52, 0x09, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x49, 0x64, 0x73, 0x42, 0x16, 0x0a, 0x14, 0x67, 0x72, 0x61, 0x70, 0x68, 0x5f,
	0x72, 0x65, 0x70, 0x72, 0x65, 0x73, 0x65, 0x6e, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x84,
	0x01, 0x0a, 0x0a, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x61, 0x69, 0x72, 0x12, 0x3a, 0x0a,
	0x0a, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e,
	0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x53, 0x75, 0x62, 0x47, 0x72, 0x61, 0x70, 0x68, 0x52, 0x09,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x4f, 0x6e, 0x65, 0x12, 0x3a, 0x0a, 0x0a, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x5f, 0x74, 0x77, 0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e,
	0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x45, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x53, 0x75, 0x62, 0x47, 0x72, 0x61, 0x70, 0x68, 0x52, 0x09, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x54, 0x77, 0x6f, 0x22, 0x54, 0x0a, 0x16, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x65,
	0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x3a, 0x0a, 0x0c, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x5f, 0x70, 0x61, 0x69, 0x72, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x2e, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x61, 0x69, 0x72, 0x52, 0x0b,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x61, 0x69, 0x72, 0x73, 0x22, 0xe4, 0x01, 0x0a, 0x17,
	0x43, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x51, 0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x70, 0x61,
	0x72, 0x69, 0x73, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61,
	0x72, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x69, 0x73, 0x6f, 0x6e, 0x52, 0x0b, 0x63,
	0x6f, 0x6d, 0x70, 0x61, 0x72, 0x69, 0x73, 0x6f, 0x6e, 0x73, 0x1a, 0x76, 0x0a, 0x0a, 0x43, 0x6f,
	0x6d, 0x70, 0x61, 0x72, 0x69, 0x73, 0x6f, 0x6e, 0x12, 0x22, 0x0a, 0x0d, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x69, 0x64, 0x5f, 0x6f, 0x6e, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x4f, 0x6e, 0x65, 0x12, 0x22, 0x0a, 0x0d,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x69, 0x64, 0x5f, 0x74, 0x77, 0x6f, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x49, 0x64, 0x54, 0x77, 0x6f,
	0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x62, 0x61, 0x62, 0x69, 0x6c, 0x69, 0x74, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x62, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x22, 0x83, 0x01, 0x0a, 0x16, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a,
	0x08, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1b, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x53, 0x75, 0x62, 0x47, 0x72, 0x61, 0x70, 0x68, 0x52, 0x08, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x30, 0x0a, 0x14, 0x77, 0x61, 0x6e, 0x74, 0x65, 0x64,
	0x5f, 0x69, 0x64, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x12, 0x77, 0x61, 0x6e, 0x74, 0x65, 0x64, 0x49, 0x64, 0x50, 0x72,
	0x6f, 0x70, 0x65, 0x72, 0x74, 0x69, 0x65, 0x73, 0x22, 0xde, 0x02, 0x0a, 0x17, 0x52, 0x65, 0x73,
	0x6f, 0x6c, 0x76, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x60, 0x0a, 0x11, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64,
	0x5f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x33, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x52, 0x65,
	0x73, 0x6f, 0x6c, 0x76, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x52, 0x10, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x1a, 0x5d, 0x0a, 0x0a, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76,
	0x65, 0x64, 0x49, 0x64, 0x12, 0x2d, 0x0a, 0x03, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x1b, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e,
	0x49, 0x64, 0x57, 0x69, 0x74, 0x68, 0x50, 0x72, 0x6f, 0x70, 0x65, 0x72, 0x74, 0x79, 0x52, 0x03,
	0x69, 0x64, 0x73, 0x12, 0x20, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x62, 0x61, 0x62, 0x69, 0x6c, 0x69,
	0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x01, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x62, 0x61, 0x62,
	0x69, 0x6c, 0x69, 0x74, 0x79, 0x1a, 0x81, 0x01, 0x0a, 0x0e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76,
	0x65, 0x64, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x49, 0x64, 0x12, 0x52, 0x0a, 0x0c, 0x72, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65,
	0x64, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2f, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76,
	0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x49, 0x64, 0x52, 0x0b, 0x72, 0x65,
	0x73, 0x6f, 0x6c, 0x76, 0x65, 0x64, 0x49, 0x64, 0x73, 0x32, 0xfb, 0x01, 0x0a, 0x05, 0x52, 0x65,
	0x63, 0x6f, 0x6e, 0x12, 0x78, 0x0a, 0x0f, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x65, 0x45, 0x6e,
	0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x23, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x65, 0x45, 0x6e, 0x74, 0x69,
	0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x72,
	0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x22, 0x0f, 0x2f, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x2f, 0x63, 0x6f, 0x6d, 0x70, 0x61, 0x72, 0x65, 0x3a, 0x01, 0x2a, 0x12, 0x78, 0x0a,
	0x0f, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x12, 0x23, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x73, 0x2e, 0x52,
	0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x63, 0x6f, 0x6d, 0x6d,
	0x6f, 0x6e, 0x73, 0x2e, 0x52, 0x65, 0x73, 0x6f, 0x6c, 0x76, 0x65, 0x45, 0x6e, 0x74, 0x69, 0x74,
	0x69, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x14, 0x22, 0x0f, 0x2f, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2f, 0x72, 0x65, 0x73,
	0x6f, 0x6c, 0x76, 0x65, 0x3a, 0x01, 0x2a, 0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f,
}

var (
	file_recon_proto_rawDescOnce sync.Once
	file_recon_proto_rawDescData = file_recon_proto_rawDesc
)

func file_recon_proto_rawDescGZIP() []byte {
	file_recon_proto_rawDescOnce.Do(func() {
		file_recon_proto_rawDescData = protoimpl.X.CompressGZIP(file_recon_proto_rawDescData)
	})
	return file_recon_proto_rawDescData
}

var file_recon_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_recon_proto_goTypes = []interface{}{
	(*IdWithProperty)(nil),                         // 0: datacommons.IdWithProperty
	(*EntityIds)(nil),                              // 1: datacommons.EntityIds
	(*EntitySubGraph)(nil),                         // 2: datacommons.EntitySubGraph
	(*EntityPair)(nil),                             // 3: datacommons.EntityPair
	(*CompareEntitiesRequest)(nil),                 // 4: datacommons.CompareEntitiesRequest
	(*CompareEntitiesResponse)(nil),                // 5: datacommons.CompareEntitiesResponse
	(*ResolveEntitiesRequest)(nil),                 // 6: datacommons.ResolveEntitiesRequest
	(*ResolveEntitiesResponse)(nil),                // 7: datacommons.ResolveEntitiesResponse
	(*CompareEntitiesResponse_Comparison)(nil),     // 8: datacommons.CompareEntitiesResponse.Comparison
	(*ResolveEntitiesResponse_ResolvedId)(nil),     // 9: datacommons.ResolveEntitiesResponse.ResolvedId
	(*ResolveEntitiesResponse_ResolvedEntity)(nil), // 10: datacommons.ResolveEntitiesResponse.ResolvedEntity
	(*McfGraph)(nil),                               // 11: datacommons.McfGraph
}
var file_recon_proto_depIdxs = []int32{
	0,  // 0: datacommons.EntityIds.ids:type_name -> datacommons.IdWithProperty
	11, // 1: datacommons.EntitySubGraph.sub_graph:type_name -> datacommons.McfGraph
	1,  // 2: datacommons.EntitySubGraph.entity_ids:type_name -> datacommons.EntityIds
	2,  // 3: datacommons.EntityPair.entity_one:type_name -> datacommons.EntitySubGraph
	2,  // 4: datacommons.EntityPair.entity_two:type_name -> datacommons.EntitySubGraph
	3,  // 5: datacommons.CompareEntitiesRequest.entity_pairs:type_name -> datacommons.EntityPair
	8,  // 6: datacommons.CompareEntitiesResponse.comparisons:type_name -> datacommons.CompareEntitiesResponse.Comparison
	2,  // 7: datacommons.ResolveEntitiesRequest.entities:type_name -> datacommons.EntitySubGraph
	10, // 8: datacommons.ResolveEntitiesResponse.resolved_entities:type_name -> datacommons.ResolveEntitiesResponse.ResolvedEntity
	0,  // 9: datacommons.ResolveEntitiesResponse.ResolvedId.ids:type_name -> datacommons.IdWithProperty
	9,  // 10: datacommons.ResolveEntitiesResponse.ResolvedEntity.resolved_ids:type_name -> datacommons.ResolveEntitiesResponse.ResolvedId
	4,  // 11: datacommons.Recon.CompareEntities:input_type -> datacommons.CompareEntitiesRequest
	6,  // 12: datacommons.Recon.ResolveEntities:input_type -> datacommons.ResolveEntitiesRequest
	5,  // 13: datacommons.Recon.CompareEntities:output_type -> datacommons.CompareEntitiesResponse
	7,  // 14: datacommons.Recon.ResolveEntities:output_type -> datacommons.ResolveEntitiesResponse
	13, // [13:15] is the sub-list for method output_type
	11, // [11:13] is the sub-list for method input_type
	11, // [11:11] is the sub-list for extension type_name
	11, // [11:11] is the sub-list for extension extendee
	0,  // [0:11] is the sub-list for field type_name
}

func init() { file_recon_proto_init() }
func file_recon_proto_init() {
	if File_recon_proto != nil {
		return
	}
	file_Mcf_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_recon_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IdWithProperty); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_recon_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EntityIds); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_recon_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EntitySubGraph); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_recon_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EntityPair); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_recon_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompareEntitiesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_recon_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompareEntitiesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_recon_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResolveEntitiesRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_recon_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResolveEntitiesResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_recon_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompareEntitiesResponse_Comparison); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_recon_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResolveEntitiesResponse_ResolvedId); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_recon_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ResolveEntitiesResponse_ResolvedEntity); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_recon_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*EntitySubGraph_SubGraph)(nil),
		(*EntitySubGraph_EntityIds)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_recon_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_recon_proto_goTypes,
		DependencyIndexes: file_recon_proto_depIdxs,
		MessageInfos:      file_recon_proto_msgTypes,
	}.Build()
	File_recon_proto = out.File
	file_recon_proto_rawDesc = nil
	file_recon_proto_goTypes = nil
	file_recon_proto_depIdxs = nil
}
