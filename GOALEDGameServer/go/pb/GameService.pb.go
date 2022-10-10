// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: GameService.proto

package pb

import (
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

type PlayerData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Key   []string `protobuf:"bytes,2,rep,name=key,proto3" json:"key,omitempty"`
	Value []string `protobuf:"bytes,3,rep,name=value,proto3" json:"value,omitempty"`
}

func (x *PlayerData) Reset() {
	*x = PlayerData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GameService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PlayerData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PlayerData) ProtoMessage() {}

func (x *PlayerData) ProtoReflect() protoreflect.Message {
	mi := &file_GameService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PlayerData.ProtoReflect.Descriptor instead.
func (*PlayerData) Descriptor() ([]byte, []int) {
	return file_GameService_proto_rawDescGZIP(), []int{0}
}

func (x *PlayerData) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PlayerData) GetKey() []string {
	if x != nil {
		return x.Key
	}
	return nil
}

func (x *PlayerData) GetValue() []string {
	if x != nil {
		return x.Value
	}
	return nil
}

type Room struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Owner     string   `protobuf:"bytes,3,opt,name=owner,proto3" json:"owner,omitempty"`
	PlayerIds []string `protobuf:"bytes,4,rep,name=player_ids,json=playerIds,proto3" json:"player_ids,omitempty"`
}

func (x *Room) Reset() {
	*x = Room{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GameService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Room) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Room) ProtoMessage() {}

func (x *Room) ProtoReflect() protoreflect.Message {
	mi := &file_GameService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Room.ProtoReflect.Descriptor instead.
func (*Room) Descriptor() ([]byte, []int) {
	return file_GameService_proto_rawDescGZIP(), []int{1}
}

func (x *Room) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Room) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Room) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *Room) GetPlayerIds() []string {
	if x != nil {
		return x.PlayerIds
	}
	return nil
}

type Vec3 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	X float32 `protobuf:"fixed32,1,opt,name=x,proto3" json:"x,omitempty"`
	Y float32 `protobuf:"fixed32,2,opt,name=y,proto3" json:"y,omitempty"`
	Z float32 `protobuf:"fixed32,3,opt,name=z,proto3" json:"z,omitempty"`
}

func (x *Vec3) Reset() {
	*x = Vec3{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GameService_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Vec3) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Vec3) ProtoMessage() {}

func (x *Vec3) ProtoReflect() protoreflect.Message {
	mi := &file_GameService_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Vec3.ProtoReflect.Descriptor instead.
func (*Vec3) Descriptor() ([]byte, []int) {
	return file_GameService_proto_rawDescGZIP(), []int{2}
}

func (x *Vec3) GetX() float32 {
	if x != nil {
		return x.X
	}
	return 0
}

func (x *Vec3) GetY() float32 {
	if x != nil {
		return x.Y
	}
	return 0
}

func (x *Vec3) GetZ() float32 {
	if x != nil {
		return x.Z
	}
	return 0
}

type Object struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id       string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Owner    string            `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Prefub   string            `protobuf:"bytes,3,opt,name=prefub,proto3" json:"prefub,omitempty"`
	Position *Vec3             `protobuf:"bytes,4,opt,name=position,proto3" json:"position,omitempty"`
	Rotation *Vec3             `protobuf:"bytes,5,opt,name=rotation,proto3" json:"rotation,omitempty"`
	Scale    *Vec3             `protobuf:"bytes,6,opt,name=scale,proto3" json:"scale,omitempty"`
	Rpc      string            `protobuf:"bytes,7,opt,name=rpc,proto3" json:"rpc,omitempty"`
	Args     map[string]string `protobuf:"bytes,8,rep,name=args,proto3" json:"args,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Object) Reset() {
	*x = Object{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GameService_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Object) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Object) ProtoMessage() {}

func (x *Object) ProtoReflect() protoreflect.Message {
	mi := &file_GameService_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Object.ProtoReflect.Descriptor instead.
func (*Object) Descriptor() ([]byte, []int) {
	return file_GameService_proto_rawDescGZIP(), []int{3}
}

func (x *Object) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Object) GetOwner() string {
	if x != nil {
		return x.Owner
	}
	return ""
}

func (x *Object) GetPrefub() string {
	if x != nil {
		return x.Prefub
	}
	return ""
}

func (x *Object) GetPosition() *Vec3 {
	if x != nil {
		return x.Position
	}
	return nil
}

func (x *Object) GetRotation() *Vec3 {
	if x != nil {
		return x.Rotation
	}
	return nil
}

func (x *Object) GetScale() *Vec3 {
	if x != nil {
		return x.Scale
	}
	return nil
}

func (x *Object) GetRpc() string {
	if x != nil {
		return x.Rpc
	}
	return ""
}

func (x *Object) GetArgs() map[string]string {
	if x != nil {
		return x.Args
	}
	return nil
}

type SyncPlayerDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId   string `protobuf:"bytes,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	PlayerId string `protobuf:"bytes,2,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
}

func (x *SyncPlayerDataRequest) Reset() {
	*x = SyncPlayerDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GameService_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncPlayerDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncPlayerDataRequest) ProtoMessage() {}

func (x *SyncPlayerDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_GameService_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncPlayerDataRequest.ProtoReflect.Descriptor instead.
func (*SyncPlayerDataRequest) Descriptor() ([]byte, []int) {
	return file_GameService_proto_rawDescGZIP(), []int{4}
}

func (x *SyncPlayerDataRequest) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *SyncPlayerDataRequest) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

type SendPlayerDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PlayerId string      `protobuf:"bytes,1,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	Data     *PlayerData `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *SendPlayerDataResponse) Reset() {
	*x = SendPlayerDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GameService_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendPlayerDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendPlayerDataResponse) ProtoMessage() {}

func (x *SendPlayerDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_GameService_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendPlayerDataResponse.ProtoReflect.Descriptor instead.
func (*SendPlayerDataResponse) Descriptor() ([]byte, []int) {
	return file_GameService_proto_rawDescGZIP(), []int{5}
}

func (x *SendPlayerDataResponse) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

func (x *SendPlayerDataResponse) GetData() *PlayerData {
	if x != nil {
		return x.Data
	}
	return nil
}

type SyncObjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId   string `protobuf:"bytes,1,opt,name=room_id,json=roomId,proto3" json:"room_id,omitempty"`
	PlayerId string `protobuf:"bytes,2,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
}

func (x *SyncObjectRequest) Reset() {
	*x = SyncObjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GameService_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncObjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncObjectRequest) ProtoMessage() {}

func (x *SyncObjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_GameService_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncObjectRequest.ProtoReflect.Descriptor instead.
func (*SyncObjectRequest) Descriptor() ([]byte, []int) {
	return file_GameService_proto_rawDescGZIP(), []int{6}
}

func (x *SyncObjectRequest) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *SyncObjectRequest) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

type SyncObjectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Object []*Object `protobuf:"bytes,1,rep,name=object,proto3" json:"object,omitempty"`
}

func (x *SyncObjectResponse) Reset() {
	*x = SyncObjectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GameService_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncObjectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncObjectResponse) ProtoMessage() {}

func (x *SyncObjectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_GameService_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncObjectResponse.ProtoReflect.Descriptor instead.
func (*SyncObjectResponse) Descriptor() ([]byte, []int) {
	return file_GameService_proto_rawDescGZIP(), []int{7}
}

func (x *SyncObjectResponse) GetObject() []*Object {
	if x != nil {
		return x.Object
	}
	return nil
}

type SendObjectRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId   string    `protobuf:"bytes,1,opt,name=Room_id,json=RoomId,proto3" json:"Room_id,omitempty"`
	PlayerId string    `protobuf:"bytes,2,opt,name=player_id,json=playerId,proto3" json:"player_id,omitempty"`
	Object   []*Object `protobuf:"bytes,3,rep,name=object,proto3" json:"object,omitempty"`
}

func (x *SendObjectRequest) Reset() {
	*x = SendObjectRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GameService_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendObjectRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendObjectRequest) ProtoMessage() {}

func (x *SendObjectRequest) ProtoReflect() protoreflect.Message {
	mi := &file_GameService_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendObjectRequest.ProtoReflect.Descriptor instead.
func (*SendObjectRequest) Descriptor() ([]byte, []int) {
	return file_GameService_proto_rawDescGZIP(), []int{8}
}

func (x *SendObjectRequest) GetRoomId() string {
	if x != nil {
		return x.RoomId
	}
	return ""
}

func (x *SendObjectRequest) GetPlayerId() string {
	if x != nil {
		return x.PlayerId
	}
	return ""
}

func (x *SendObjectRequest) GetObject() []*Object {
	if x != nil {
		return x.Object
	}
	return nil
}

type SendObjectResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SendObjectResponse) Reset() {
	*x = SendObjectResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_GameService_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendObjectResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendObjectResponse) ProtoMessage() {}

func (x *SendObjectResponse) ProtoReflect() protoreflect.Message {
	mi := &file_GameService_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendObjectResponse.ProtoReflect.Descriptor instead.
func (*SendObjectResponse) Descriptor() ([]byte, []int) {
	return file_GameService_proto_rawDescGZIP(), []int{9}
}

func (x *SendObjectResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_GameService_proto protoreflect.FileDescriptor

var file_GameService_proto_rawDesc = []byte{
	0x0a, 0x11, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x22, 0x44, 0x0a, 0x0a, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x02, 0x20, 0x03, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x03, 0x20, 0x03, 0x28, 0x09, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x5f, 0x0a, 0x04, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12,
	0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09, 0x52, 0x09, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x73, 0x22, 0x30, 0x0a, 0x04, 0x56, 0x65, 0x63, 0x33, 0x12,
	0x0c, 0x0a, 0x01, 0x78, 0x18, 0x01, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x78, 0x12, 0x0c, 0x0a,
	0x01, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x79, 0x12, 0x0c, 0x0a, 0x01, 0x7a,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x02, 0x52, 0x01, 0x7a, 0x22, 0xcb, 0x02, 0x0a, 0x06, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x70, 0x72,
	0x65, 0x66, 0x75, 0x62, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x72, 0x65, 0x66,
	0x75, 0x62, 0x12, 0x2d, 0x0a, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x56, 0x65, 0x63, 0x33, 0x52, 0x08, 0x70, 0x6f, 0x73, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x2d, 0x0a, 0x08, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x56, 0x65, 0x63, 0x33, 0x52, 0x08, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x27, 0x0a, 0x05, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x11, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x56, 0x65,
	0x63, 0x33, 0x52, 0x05, 0x73, 0x63, 0x61, 0x6c, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x72, 0x70, 0x63,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x72, 0x70, 0x63, 0x12, 0x31, 0x0a, 0x04, 0x61,
	0x72, 0x67, 0x73, 0x18, 0x08, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x47, 0x61, 0x6d, 0x65,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x2e, 0x41,
	0x72, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x61, 0x72, 0x67, 0x73, 0x1a, 0x37,
	0x0a, 0x09, 0x41, 0x72, 0x67, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x4d, 0x0a, 0x15, 0x53, 0x79, 0x6e, 0x63, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x22, 0x62, 0x0a, 0x16, 0x53, 0x65, 0x6e, 0x64, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2b, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x47, 0x61,
	0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x49, 0x0a, 0x11, 0x53, 0x79,
	0x6e, 0x63, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x72, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x72, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x49, 0x64, 0x22, 0x41, 0x0a, 0x12, 0x53, 0x79, 0x6e, 0x63, 0x4f, 0x62, 0x6a,
	0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x06, 0x6f,
	0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x47, 0x61,
	0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x22, 0x76, 0x0a, 0x11, 0x53, 0x65, 0x6e, 0x64,
	0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a,
	0x07, 0x52, 0x6f, 0x6f, 0x6d, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x70, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x03, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x06, 0x6f, 0x62, 0x6a, 0x65, 0x63, 0x74,
	0x22, 0x2e, 0x0a, 0x12, 0x53, 0x65, 0x6e, 0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x32, 0x90, 0x03, 0x0a, 0x0b, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x34, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x11,
	0x2e, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x6f, 0x6f,
	0x6d, 0x1a, 0x11, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x52, 0x6f, 0x6f, 0x6d, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x0e, 0x53, 0x79, 0x6e, 0x63, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x22, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x47,
	0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x44, 0x61, 0x74, 0x61, 0x22, 0x00, 0x30, 0x01, 0x12, 0x52, 0x0a, 0x0e, 0x53, 0x65, 0x6e,
	0x64, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x17, 0x2e, 0x47, 0x61,
	0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72,
	0x44, 0x61, 0x74, 0x61, 0x1a, 0x23, 0x2e, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x53, 0x65, 0x6e, 0x64, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x44, 0x61, 0x74,
	0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x28, 0x01, 0x12, 0x51, 0x0a,
	0x0a, 0x53, 0x79, 0x6e, 0x63, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1e, 0x2e, 0x47, 0x61,
	0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x47, 0x61,
	0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x4f, 0x62,
	0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x30, 0x01,
	0x12, 0x51, 0x0a, 0x0a, 0x53, 0x65, 0x6e, 0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x1e,
	0x2e, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x6e,
	0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f,
	0x2e, 0x47, 0x61, 0x6d, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x65, 0x6e,
	0x64, 0x4f, 0x62, 0x6a, 0x65, 0x63, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x28, 0x01, 0x42, 0x0a, 0x5a, 0x08, 0x2e, 0x2e, 0x2f, 0x67, 0x6f, 0x2f, 0x70, 0x62, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_GameService_proto_rawDescOnce sync.Once
	file_GameService_proto_rawDescData = file_GameService_proto_rawDesc
)

func file_GameService_proto_rawDescGZIP() []byte {
	file_GameService_proto_rawDescOnce.Do(func() {
		file_GameService_proto_rawDescData = protoimpl.X.CompressGZIP(file_GameService_proto_rawDescData)
	})
	return file_GameService_proto_rawDescData
}

var file_GameService_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_GameService_proto_goTypes = []interface{}{
	(*PlayerData)(nil),             // 0: GameService.PlayerData
	(*Room)(nil),                   // 1: GameService.Room
	(*Vec3)(nil),                   // 2: GameService.Vec3
	(*Object)(nil),                 // 3: GameService.Object
	(*SyncPlayerDataRequest)(nil),  // 4: GameService.SyncPlayerDataRequest
	(*SendPlayerDataResponse)(nil), // 5: GameService.SendPlayerDataResponse
	(*SyncObjectRequest)(nil),      // 6: GameService.SyncObjectRequest
	(*SyncObjectResponse)(nil),     // 7: GameService.SyncObjectResponse
	(*SendObjectRequest)(nil),      // 8: GameService.SendObjectRequest
	(*SendObjectResponse)(nil),     // 9: GameService.SendObjectResponse
	nil,                            // 10: GameService.Object.ArgsEntry
}
var file_GameService_proto_depIdxs = []int32{
	2,  // 0: GameService.Object.position:type_name -> GameService.Vec3
	2,  // 1: GameService.Object.rotation:type_name -> GameService.Vec3
	2,  // 2: GameService.Object.scale:type_name -> GameService.Vec3
	10, // 3: GameService.Object.args:type_name -> GameService.Object.ArgsEntry
	0,  // 4: GameService.SendPlayerDataResponse.data:type_name -> GameService.PlayerData
	3,  // 5: GameService.SyncObjectResponse.object:type_name -> GameService.Object
	3,  // 6: GameService.SendObjectRequest.object:type_name -> GameService.Object
	1,  // 7: GameService.GameService.CreateRoom:input_type -> GameService.Room
	4,  // 8: GameService.GameService.SyncPlayerData:input_type -> GameService.SyncPlayerDataRequest
	0,  // 9: GameService.GameService.SendPlayerData:input_type -> GameService.PlayerData
	6,  // 10: GameService.GameService.SyncObject:input_type -> GameService.SyncObjectRequest
	8,  // 11: GameService.GameService.SendObject:input_type -> GameService.SendObjectRequest
	1,  // 12: GameService.GameService.CreateRoom:output_type -> GameService.Room
	0,  // 13: GameService.GameService.SyncPlayerData:output_type -> GameService.PlayerData
	5,  // 14: GameService.GameService.SendPlayerData:output_type -> GameService.SendPlayerDataResponse
	7,  // 15: GameService.GameService.SyncObject:output_type -> GameService.SyncObjectResponse
	9,  // 16: GameService.GameService.SendObject:output_type -> GameService.SendObjectResponse
	12, // [12:17] is the sub-list for method output_type
	7,  // [7:12] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_GameService_proto_init() }
func file_GameService_proto_init() {
	if File_GameService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_GameService_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PlayerData); i {
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
		file_GameService_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Room); i {
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
		file_GameService_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Vec3); i {
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
		file_GameService_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Object); i {
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
		file_GameService_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncPlayerDataRequest); i {
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
		file_GameService_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendPlayerDataResponse); i {
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
		file_GameService_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncObjectRequest); i {
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
		file_GameService_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncObjectResponse); i {
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
		file_GameService_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendObjectRequest); i {
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
		file_GameService_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendObjectResponse); i {
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
			RawDescriptor: file_GameService_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_GameService_proto_goTypes,
		DependencyIndexes: file_GameService_proto_depIdxs,
		MessageInfos:      file_GameService_proto_msgTypes,
	}.Build()
	File_GameService_proto = out.File
	file_GameService_proto_rawDesc = nil
	file_GameService_proto_goTypes = nil
	file_GameService_proto_depIdxs = nil
}
