// Copyright 2023 Gravitational, Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: teleport/accesslist/v1/accesslist_service.proto

package accesslistv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// GetAccessListsRequest is the request for getting all access lists.
type GetAccessListsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAccessListsRequest) Reset() {
	*x = GetAccessListsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_accesslist_v1_accesslist_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccessListsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccessListsRequest) ProtoMessage() {}

func (x *GetAccessListsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_accesslist_v1_accesslist_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccessListsRequest.ProtoReflect.Descriptor instead.
func (*GetAccessListsRequest) Descriptor() ([]byte, []int) {
	return file_teleport_accesslist_v1_accesslist_service_proto_rawDescGZIP(), []int{0}
}

// GetAccessListsResponse is the response for getting all access lists.
type GetAccessListsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// access_lists is the list of access lists.
	AccessLists []*AccessList `protobuf:"bytes,1,rep,name=access_lists,json=accessLists,proto3" json:"access_lists,omitempty"`
}

func (x *GetAccessListsResponse) Reset() {
	*x = GetAccessListsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_accesslist_v1_accesslist_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccessListsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccessListsResponse) ProtoMessage() {}

func (x *GetAccessListsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_accesslist_v1_accesslist_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccessListsResponse.ProtoReflect.Descriptor instead.
func (*GetAccessListsResponse) Descriptor() ([]byte, []int) {
	return file_teleport_accesslist_v1_accesslist_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetAccessListsResponse) GetAccessLists() []*AccessList {
	if x != nil {
		return x.AccessLists
	}
	return nil
}

// ListAccessListsRequest is the request for getting paginated access lists.
type ListAccessListsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// page_size is the size of the page to request.
	PageSize int32 `protobuf:"varint,1,opt,name=page_size,json=pageSize,proto3" json:"page_size,omitempty"`
	// next_token is the page token.
	NextToken string `protobuf:"bytes,2,opt,name=next_token,json=nextToken,proto3" json:"next_token,omitempty"`
}

func (x *ListAccessListsRequest) Reset() {
	*x = ListAccessListsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_accesslist_v1_accesslist_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAccessListsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAccessListsRequest) ProtoMessage() {}

func (x *ListAccessListsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_accesslist_v1_accesslist_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAccessListsRequest.ProtoReflect.Descriptor instead.
func (*ListAccessListsRequest) Descriptor() ([]byte, []int) {
	return file_teleport_accesslist_v1_accesslist_service_proto_rawDescGZIP(), []int{2}
}

func (x *ListAccessListsRequest) GetPageSize() int32 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *ListAccessListsRequest) GetNextToken() string {
	if x != nil {
		return x.NextToken
	}
	return ""
}

// ListAccessListsResponse is the response for getting paginated access lists.
type ListAccessListsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// access_lists is the list of access lists.
	AccessLists []*AccessList `protobuf:"bytes,1,rep,name=access_lists,json=accessLists,proto3" json:"access_lists,omitempty"`
	// next_token is the next page token.
	NextToken string `protobuf:"bytes,2,opt,name=next_token,json=nextToken,proto3" json:"next_token,omitempty"`
}

func (x *ListAccessListsResponse) Reset() {
	*x = ListAccessListsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_accesslist_v1_accesslist_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListAccessListsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListAccessListsResponse) ProtoMessage() {}

func (x *ListAccessListsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_accesslist_v1_accesslist_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListAccessListsResponse.ProtoReflect.Descriptor instead.
func (*ListAccessListsResponse) Descriptor() ([]byte, []int) {
	return file_teleport_accesslist_v1_accesslist_service_proto_rawDescGZIP(), []int{3}
}

func (x *ListAccessListsResponse) GetAccessLists() []*AccessList {
	if x != nil {
		return x.AccessLists
	}
	return nil
}

func (x *ListAccessListsResponse) GetNextToken() string {
	if x != nil {
		return x.NextToken
	}
	return ""
}

// GetAccessListRequest is the request for retrieving an access list.
type GetAccessListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name is the name of the access list to retrieve.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetAccessListRequest) Reset() {
	*x = GetAccessListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_accesslist_v1_accesslist_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAccessListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccessListRequest) ProtoMessage() {}

func (x *GetAccessListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_accesslist_v1_accesslist_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccessListRequest.ProtoReflect.Descriptor instead.
func (*GetAccessListRequest) Descriptor() ([]byte, []int) {
	return file_teleport_accesslist_v1_accesslist_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetAccessListRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// UpsertAccessListRequest is the request for upserting an access list.
type UpsertAccessListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// access_list is the access list to upsert.
	AccessList *AccessList `protobuf:"bytes,1,opt,name=access_list,json=accessList,proto3" json:"access_list,omitempty"`
}

func (x *UpsertAccessListRequest) Reset() {
	*x = UpsertAccessListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_accesslist_v1_accesslist_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpsertAccessListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertAccessListRequest) ProtoMessage() {}

func (x *UpsertAccessListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_accesslist_v1_accesslist_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertAccessListRequest.ProtoReflect.Descriptor instead.
func (*UpsertAccessListRequest) Descriptor() ([]byte, []int) {
	return file_teleport_accesslist_v1_accesslist_service_proto_rawDescGZIP(), []int{5}
}

func (x *UpsertAccessListRequest) GetAccessList() *AccessList {
	if x != nil {
		return x.AccessList
	}
	return nil
}

// DeleteAccessListRequest is the request for deleting an access list.
type DeleteAccessListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// name is the name of the access list to delete.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *DeleteAccessListRequest) Reset() {
	*x = DeleteAccessListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_accesslist_v1_accesslist_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAccessListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAccessListRequest) ProtoMessage() {}

func (x *DeleteAccessListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_accesslist_v1_accesslist_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAccessListRequest.ProtoReflect.Descriptor instead.
func (*DeleteAccessListRequest) Descriptor() ([]byte, []int) {
	return file_teleport_accesslist_v1_accesslist_service_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteAccessListRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

// DeleteAllAccessListRequest is the request for deleting all access lists.
type DeleteAllAccessListsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteAllAccessListsRequest) Reset() {
	*x = DeleteAllAccessListsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_accesslist_v1_accesslist_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteAllAccessListsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteAllAccessListsRequest) ProtoMessage() {}

func (x *DeleteAllAccessListsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_accesslist_v1_accesslist_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteAllAccessListsRequest.ProtoReflect.Descriptor instead.
func (*DeleteAllAccessListsRequest) Descriptor() ([]byte, []int) {
	return file_teleport_accesslist_v1_accesslist_service_proto_rawDescGZIP(), []int{7}
}

var File_teleport_accesslist_v1_accesslist_service_proto protoreflect.FileDescriptor

var file_teleport_accesslist_v1_accesslist_service_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6c,
	0x69, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x16, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x27, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x61,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x17, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x73, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x5f, 0x0a, 0x16, 0x47, 0x65, 0x74, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x45, 0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x0b, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x22, 0x54, 0x0a, 0x16, 0x4c, 0x69, 0x73,
	0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x61, 0x67, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a, 0x65,
	0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x65, 0x78, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22,
	0x7f, 0x0a, 0x17, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x73,
	0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x45, 0x0a, 0x0c, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x22, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x61, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x4c, 0x69, 0x73, 0x74, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x73, 0x74,
	0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x6e, 0x65, 0x78, 0x74, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x65, 0x78, 0x74, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x22, 0x2a, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x73,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x5e, 0x0a, 0x17,
	0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x43, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x74,
	0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6c, 0x69,
	0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x0a, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x2d, 0x0a, 0x17,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x73, 0x74,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x1d, 0x0a, 0x1b, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x6c, 0x6c, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x69,
	0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x32, 0x86, 0x05, 0x0a, 0x11, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x6f, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x73,
	0x74, 0x73, 0x12, 0x2d, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x61, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41,
	0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x2e, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x61, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x72, 0x0a, 0x0f, 0x4c, 0x69, 0x73, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c,
	0x69, 0x73, 0x74, 0x73, 0x12, 0x2e, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x2f, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69,
	0x73, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x61, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2c, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e,
	0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63,
	0x63, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x67, 0x0a, 0x10, 0x55, 0x70, 0x73, 0x65,
	0x72, 0x74, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2f, 0x2e, 0x74,
	0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6c, 0x69,
	0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x41, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e,
	0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6c,
	0x69, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x5b, 0x0a, 0x10, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x2f, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74,
	0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x63,
	0x0a, 0x14, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x6c, 0x6c, 0x41, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x4c, 0x69, 0x73, 0x74, 0x73, 0x12, 0x33, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x2e, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6c, 0x69, 0x73, 0x74, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x41, 0x6c, 0x6c, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x4c,
	0x69, 0x73, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x42, 0x58, 0x5a, 0x56, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x67, 0x72, 0x61, 0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x2f,
	0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f,
	0x72, 0x74, 0x2f, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6c, 0x69, 0x73, 0x74, 0x2f, 0x76, 0x31,
	0x3b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x6c, 0x69, 0x73, 0x74, 0x76, 0x31, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_teleport_accesslist_v1_accesslist_service_proto_rawDescOnce sync.Once
	file_teleport_accesslist_v1_accesslist_service_proto_rawDescData = file_teleport_accesslist_v1_accesslist_service_proto_rawDesc
)

func file_teleport_accesslist_v1_accesslist_service_proto_rawDescGZIP() []byte {
	file_teleport_accesslist_v1_accesslist_service_proto_rawDescOnce.Do(func() {
		file_teleport_accesslist_v1_accesslist_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_teleport_accesslist_v1_accesslist_service_proto_rawDescData)
	})
	return file_teleport_accesslist_v1_accesslist_service_proto_rawDescData
}

var file_teleport_accesslist_v1_accesslist_service_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_teleport_accesslist_v1_accesslist_service_proto_goTypes = []interface{}{
	(*GetAccessListsRequest)(nil),       // 0: teleport.accesslist.v1.GetAccessListsRequest
	(*GetAccessListsResponse)(nil),      // 1: teleport.accesslist.v1.GetAccessListsResponse
	(*ListAccessListsRequest)(nil),      // 2: teleport.accesslist.v1.ListAccessListsRequest
	(*ListAccessListsResponse)(nil),     // 3: teleport.accesslist.v1.ListAccessListsResponse
	(*GetAccessListRequest)(nil),        // 4: teleport.accesslist.v1.GetAccessListRequest
	(*UpsertAccessListRequest)(nil),     // 5: teleport.accesslist.v1.UpsertAccessListRequest
	(*DeleteAccessListRequest)(nil),     // 6: teleport.accesslist.v1.DeleteAccessListRequest
	(*DeleteAllAccessListsRequest)(nil), // 7: teleport.accesslist.v1.DeleteAllAccessListsRequest
	(*AccessList)(nil),                  // 8: teleport.accesslist.v1.AccessList
	(*emptypb.Empty)(nil),               // 9: google.protobuf.Empty
}
var file_teleport_accesslist_v1_accesslist_service_proto_depIdxs = []int32{
	8, // 0: teleport.accesslist.v1.GetAccessListsResponse.access_lists:type_name -> teleport.accesslist.v1.AccessList
	8, // 1: teleport.accesslist.v1.ListAccessListsResponse.access_lists:type_name -> teleport.accesslist.v1.AccessList
	8, // 2: teleport.accesslist.v1.UpsertAccessListRequest.access_list:type_name -> teleport.accesslist.v1.AccessList
	0, // 3: teleport.accesslist.v1.AccessListService.GetAccessLists:input_type -> teleport.accesslist.v1.GetAccessListsRequest
	2, // 4: teleport.accesslist.v1.AccessListService.ListAccessLists:input_type -> teleport.accesslist.v1.ListAccessListsRequest
	4, // 5: teleport.accesslist.v1.AccessListService.GetAccessList:input_type -> teleport.accesslist.v1.GetAccessListRequest
	5, // 6: teleport.accesslist.v1.AccessListService.UpsertAccessList:input_type -> teleport.accesslist.v1.UpsertAccessListRequest
	6, // 7: teleport.accesslist.v1.AccessListService.DeleteAccessList:input_type -> teleport.accesslist.v1.DeleteAccessListRequest
	7, // 8: teleport.accesslist.v1.AccessListService.DeleteAllAccessLists:input_type -> teleport.accesslist.v1.DeleteAllAccessListsRequest
	1, // 9: teleport.accesslist.v1.AccessListService.GetAccessLists:output_type -> teleport.accesslist.v1.GetAccessListsResponse
	3, // 10: teleport.accesslist.v1.AccessListService.ListAccessLists:output_type -> teleport.accesslist.v1.ListAccessListsResponse
	8, // 11: teleport.accesslist.v1.AccessListService.GetAccessList:output_type -> teleport.accesslist.v1.AccessList
	8, // 12: teleport.accesslist.v1.AccessListService.UpsertAccessList:output_type -> teleport.accesslist.v1.AccessList
	9, // 13: teleport.accesslist.v1.AccessListService.DeleteAccessList:output_type -> google.protobuf.Empty
	9, // 14: teleport.accesslist.v1.AccessListService.DeleteAllAccessLists:output_type -> google.protobuf.Empty
	9, // [9:15] is the sub-list for method output_type
	3, // [3:9] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_teleport_accesslist_v1_accesslist_service_proto_init() }
func file_teleport_accesslist_v1_accesslist_service_proto_init() {
	if File_teleport_accesslist_v1_accesslist_service_proto != nil {
		return
	}
	file_teleport_accesslist_v1_accesslist_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_teleport_accesslist_v1_accesslist_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccessListsRequest); i {
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
		file_teleport_accesslist_v1_accesslist_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccessListsResponse); i {
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
		file_teleport_accesslist_v1_accesslist_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAccessListsRequest); i {
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
		file_teleport_accesslist_v1_accesslist_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListAccessListsResponse); i {
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
		file_teleport_accesslist_v1_accesslist_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAccessListRequest); i {
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
		file_teleport_accesslist_v1_accesslist_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpsertAccessListRequest); i {
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
		file_teleport_accesslist_v1_accesslist_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAccessListRequest); i {
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
		file_teleport_accesslist_v1_accesslist_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteAllAccessListsRequest); i {
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
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_teleport_accesslist_v1_accesslist_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_teleport_accesslist_v1_accesslist_service_proto_goTypes,
		DependencyIndexes: file_teleport_accesslist_v1_accesslist_service_proto_depIdxs,
		MessageInfos:      file_teleport_accesslist_v1_accesslist_service_proto_msgTypes,
	}.Build()
	File_teleport_accesslist_v1_accesslist_service_proto = out.File
	file_teleport_accesslist_v1_accesslist_service_proto_rawDesc = nil
	file_teleport_accesslist_v1_accesslist_service_proto_goTypes = nil
	file_teleport_accesslist_v1_accesslist_service_proto_depIdxs = nil
}