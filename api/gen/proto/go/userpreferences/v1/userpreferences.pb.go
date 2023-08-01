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
// 	protoc-gen-go v1.30.0
// 	protoc        (unknown)
// source: teleport/userpreferences/v1/userpreferences.proto

package userpreferencesv1

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

// UserPreferences is a collection of different user changeable preferences for the frontend.
type UserPreferences struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// assist is the preferences for the Teleport Assist.
	Assist *AssistUserPreferences `protobuf:"bytes,1,opt,name=assist,proto3" json:"assist,omitempty"`
	// theme is the theme of the frontend.
	Theme Theme `protobuf:"varint,2,opt,name=theme,proto3,enum=teleport.userpreferences.v1.Theme" json:"theme,omitempty"`
	// onboard is the preferences from the onboarding questionnaire.
	Onboard *OnboardUserPreferences `protobuf:"bytes,3,opt,name=onboard,proto3" json:"onboard,omitempty"`
}

func (x *UserPreferences) Reset() {
	*x = UserPreferences{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_userpreferences_v1_userpreferences_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserPreferences) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserPreferences) ProtoMessage() {}

func (x *UserPreferences) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_userpreferences_v1_userpreferences_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserPreferences.ProtoReflect.Descriptor instead.
func (*UserPreferences) Descriptor() ([]byte, []int) {
	return file_teleport_userpreferences_v1_userpreferences_proto_rawDescGZIP(), []int{0}
}

func (x *UserPreferences) GetAssist() *AssistUserPreferences {
	if x != nil {
		return x.Assist
	}
	return nil
}

func (x *UserPreferences) GetTheme() Theme {
	if x != nil {
		return x.Theme
	}
	return Theme_THEME_UNSPECIFIED
}

func (x *UserPreferences) GetOnboard() *OnboardUserPreferences {
	if x != nil {
		return x.Onboard
	}
	return nil
}

// GetUserPreferencesRequest is a request to get the user preferences.
type GetUserPreferencesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetUserPreferencesRequest) Reset() {
	*x = GetUserPreferencesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_userpreferences_v1_userpreferences_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserPreferencesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserPreferencesRequest) ProtoMessage() {}

func (x *GetUserPreferencesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_userpreferences_v1_userpreferences_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserPreferencesRequest.ProtoReflect.Descriptor instead.
func (*GetUserPreferencesRequest) Descriptor() ([]byte, []int) {
	return file_teleport_userpreferences_v1_userpreferences_proto_rawDescGZIP(), []int{1}
}

// GetUserPreferencesResponse is a response to get the user preferences.
type GetUserPreferencesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// preferences is the user preferences.
	Preferences *UserPreferences `protobuf:"bytes,1,opt,name=preferences,proto3" json:"preferences,omitempty"`
}

func (x *GetUserPreferencesResponse) Reset() {
	*x = GetUserPreferencesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_userpreferences_v1_userpreferences_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserPreferencesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserPreferencesResponse) ProtoMessage() {}

func (x *GetUserPreferencesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_userpreferences_v1_userpreferences_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserPreferencesResponse.ProtoReflect.Descriptor instead.
func (*GetUserPreferencesResponse) Descriptor() ([]byte, []int) {
	return file_teleport_userpreferences_v1_userpreferences_proto_rawDescGZIP(), []int{2}
}

func (x *GetUserPreferencesResponse) GetPreferences() *UserPreferences {
	if x != nil {
		return x.Preferences
	}
	return nil
}

// UpsertUserPreferencesRequest is a request to create or update the user preferences.
type UpsertUserPreferencesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// preferences is the new user preferences to set.
	Preferences *UserPreferences `protobuf:"bytes,1,opt,name=preferences,proto3" json:"preferences,omitempty"`
}

func (x *UpsertUserPreferencesRequest) Reset() {
	*x = UpsertUserPreferencesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_teleport_userpreferences_v1_userpreferences_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpsertUserPreferencesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpsertUserPreferencesRequest) ProtoMessage() {}

func (x *UpsertUserPreferencesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_teleport_userpreferences_v1_userpreferences_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpsertUserPreferencesRequest.ProtoReflect.Descriptor instead.
func (*UpsertUserPreferencesRequest) Descriptor() ([]byte, []int) {
	return file_teleport_userpreferences_v1_userpreferences_proto_rawDescGZIP(), []int{3}
}

func (x *UpsertUserPreferencesRequest) GetPreferences() *UserPreferences {
	if x != nil {
		return x.Preferences
	}
	return nil
}

var File_teleport_userpreferences_v1_userpreferences_proto protoreflect.FileDescriptor

var file_teleport_userpreferences_v1_userpreferences_proto_rawDesc = []byte{
	0x0a, 0x31, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x70,
	0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x28, 0x74,
	0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x70, 0x72, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x61, 0x73, 0x73, 0x69, 0x73,
	0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x27, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x74, 0x68, 0x65, 0x6d, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xe6, 0x01, 0x0a, 0x0f,
	0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x12,
	0x4a, 0x0a, 0x06, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x32, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x70,
	0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x73,
	0x73, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x73, 0x52, 0x06, 0x61, 0x73, 0x73, 0x69, 0x73, 0x74, 0x12, 0x38, 0x0a, 0x05, 0x74,
	0x68, 0x65, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x22, 0x2e, 0x74, 0x65, 0x6c,
	0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72,
	0x65, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x54, 0x68, 0x65, 0x6d, 0x65, 0x52, 0x05,
	0x74, 0x68, 0x65, 0x6d, 0x65, 0x12, 0x4d, 0x0a, 0x07, 0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72,
	0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x55, 0x73, 0x65, 0x72,
	0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x52, 0x07, 0x6f, 0x6e, 0x62,
	0x6f, 0x61, 0x72, 0x64, 0x22, 0x2b, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50,
	0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x4a, 0x04, 0x08, 0x01, 0x10, 0x02, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x6c, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x4e, 0x0a, 0x0b, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x73, 0x52, 0x0b, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x22,
	0x7e, 0x0a, 0x1c, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x65,
	0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x4e, 0x0a, 0x0b, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2c, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x73, 0x52, 0x0b, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x4a,
	0x04, 0x08, 0x02, 0x10, 0x03, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x32,
	0x8c, 0x02, 0x0a, 0x16, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e,
	0x63, 0x65, 0x73, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x85, 0x01, 0x0a, 0x12, 0x47,
	0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65,
	0x73, 0x12, 0x36, 0x2e, 0x74, 0x65, 0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x75, 0x73, 0x65,
	0x72, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63,
	0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x37, 0x2e, 0x74, 0x65, 0x6c, 0x65,
	0x70, 0x6f, 0x72, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x50,
	0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x6a, 0x0a, 0x15, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x39, 0x2e, 0x74, 0x65,
	0x6c, 0x65, 0x70, 0x6f, 0x72, 0x74, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x70, 0x72, 0x65, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x70, 0x73, 0x65, 0x72, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x50, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x59,
	0x5a, 0x57, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x72, 0x61,
	0x76, 0x69, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x2f, 0x74, 0x65, 0x6c, 0x65, 0x70,
	0x6f, 0x72, 0x74, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x70, 0x72, 0x65, 0x66, 0x65, 0x72, 0x65,
	0x6e, 0x63, 0x65, 0x73, 0x2f, 0x76, 0x31, 0x3b, 0x75, 0x73, 0x65, 0x72, 0x70, 0x72, 0x65, 0x66,
	0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x73, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_teleport_userpreferences_v1_userpreferences_proto_rawDescOnce sync.Once
	file_teleport_userpreferences_v1_userpreferences_proto_rawDescData = file_teleport_userpreferences_v1_userpreferences_proto_rawDesc
)

func file_teleport_userpreferences_v1_userpreferences_proto_rawDescGZIP() []byte {
	file_teleport_userpreferences_v1_userpreferences_proto_rawDescOnce.Do(func() {
		file_teleport_userpreferences_v1_userpreferences_proto_rawDescData = protoimpl.X.CompressGZIP(file_teleport_userpreferences_v1_userpreferences_proto_rawDescData)
	})
	return file_teleport_userpreferences_v1_userpreferences_proto_rawDescData
}

var file_teleport_userpreferences_v1_userpreferences_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_teleport_userpreferences_v1_userpreferences_proto_goTypes = []interface{}{
	(*UserPreferences)(nil),              // 0: teleport.userpreferences.v1.UserPreferences
	(*GetUserPreferencesRequest)(nil),    // 1: teleport.userpreferences.v1.GetUserPreferencesRequest
	(*GetUserPreferencesResponse)(nil),   // 2: teleport.userpreferences.v1.GetUserPreferencesResponse
	(*UpsertUserPreferencesRequest)(nil), // 3: teleport.userpreferences.v1.UpsertUserPreferencesRequest
	(*AssistUserPreferences)(nil),        // 4: teleport.userpreferences.v1.AssistUserPreferences
	(Theme)(0),                           // 5: teleport.userpreferences.v1.Theme
	(*OnboardUserPreferences)(nil),       // 6: teleport.userpreferences.v1.OnboardUserPreferences
	(*emptypb.Empty)(nil),                // 7: google.protobuf.Empty
}
var file_teleport_userpreferences_v1_userpreferences_proto_depIdxs = []int32{
	4, // 0: teleport.userpreferences.v1.UserPreferences.assist:type_name -> teleport.userpreferences.v1.AssistUserPreferences
	5, // 1: teleport.userpreferences.v1.UserPreferences.theme:type_name -> teleport.userpreferences.v1.Theme
	6, // 2: teleport.userpreferences.v1.UserPreferences.onboard:type_name -> teleport.userpreferences.v1.OnboardUserPreferences
	0, // 3: teleport.userpreferences.v1.GetUserPreferencesResponse.preferences:type_name -> teleport.userpreferences.v1.UserPreferences
	0, // 4: teleport.userpreferences.v1.UpsertUserPreferencesRequest.preferences:type_name -> teleport.userpreferences.v1.UserPreferences
	1, // 5: teleport.userpreferences.v1.UserPreferencesService.GetUserPreferences:input_type -> teleport.userpreferences.v1.GetUserPreferencesRequest
	3, // 6: teleport.userpreferences.v1.UserPreferencesService.UpsertUserPreferences:input_type -> teleport.userpreferences.v1.UpsertUserPreferencesRequest
	2, // 7: teleport.userpreferences.v1.UserPreferencesService.GetUserPreferences:output_type -> teleport.userpreferences.v1.GetUserPreferencesResponse
	7, // 8: teleport.userpreferences.v1.UserPreferencesService.UpsertUserPreferences:output_type -> google.protobuf.Empty
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_teleport_userpreferences_v1_userpreferences_proto_init() }
func file_teleport_userpreferences_v1_userpreferences_proto_init() {
	if File_teleport_userpreferences_v1_userpreferences_proto != nil {
		return
	}
	file_teleport_userpreferences_v1_assist_proto_init()
	file_teleport_userpreferences_v1_onboard_proto_init()
	file_teleport_userpreferences_v1_theme_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_teleport_userpreferences_v1_userpreferences_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserPreferences); i {
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
		file_teleport_userpreferences_v1_userpreferences_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserPreferencesRequest); i {
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
		file_teleport_userpreferences_v1_userpreferences_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserPreferencesResponse); i {
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
		file_teleport_userpreferences_v1_userpreferences_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpsertUserPreferencesRequest); i {
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
			RawDescriptor: file_teleport_userpreferences_v1_userpreferences_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_teleport_userpreferences_v1_userpreferences_proto_goTypes,
		DependencyIndexes: file_teleport_userpreferences_v1_userpreferences_proto_depIdxs,
		MessageInfos:      file_teleport_userpreferences_v1_userpreferences_proto_msgTypes,
	}.Build()
	File_teleport_userpreferences_v1_userpreferences_proto = out.File
	file_teleport_userpreferences_v1_userpreferences_proto_rawDesc = nil
	file_teleport_userpreferences_v1_userpreferences_proto_goTypes = nil
	file_teleport_userpreferences_v1_userpreferences_proto_depIdxs = nil
}