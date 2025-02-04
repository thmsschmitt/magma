//
//Copyright 2020 The Magma Authors.
//
//This source code is licensed under the BSD-style license found in the
//LICENSE file in the root directory of this source tree.
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.10.0
// source: orc8r/protos/identity.proto

package protos

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

type Identity_Wildcard_Type int32

const (
	Identity_Wildcard_Gateway  Identity_Wildcard_Type = 0
	Identity_Wildcard_Operator Identity_Wildcard_Type = 1
	Identity_Wildcard_Network  Identity_Wildcard_Type = 2
)

// Enum value maps for Identity_Wildcard_Type.
var (
	Identity_Wildcard_Type_name = map[int32]string{
		0: "Gateway",
		1: "Operator",
		2: "Network",
	}
	Identity_Wildcard_Type_value = map[string]int32{
		"Gateway":  0,
		"Operator": 1,
		"Network":  2,
	}
)

func (x Identity_Wildcard_Type) Enum() *Identity_Wildcard_Type {
	p := new(Identity_Wildcard_Type)
	*p = x
	return p
}

func (x Identity_Wildcard_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Identity_Wildcard_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_orc8r_protos_identity_proto_enumTypes[0].Descriptor()
}

func (Identity_Wildcard_Type) Type() protoreflect.EnumType {
	return &file_orc8r_protos_identity_proto_enumTypes[0]
}

func (x Identity_Wildcard_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Identity_Wildcard_Type.Descriptor instead.
func (Identity_Wildcard_Type) EnumDescriptor() ([]byte, []int) {
	return file_orc8r_protos_identity_proto_rawDescGZIP(), []int{0, 0, 0}
}

type Identity struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// NOTE: Every Value type below should have corresponding entry in
	// identityTypeNameTable in magma/protos/identity_helper.go
	//
	// Types that are assignable to Value:
	//	*Identity_Gateway_
	//	*Identity_Operator
	//	*Identity_Network
	//	*Identity_Wildcard_
	Value isIdentity_Value `protobuf_oneof:"Value"`
}

func (x *Identity) Reset() {
	*x = Identity{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orc8r_protos_identity_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Identity) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Identity) ProtoMessage() {}

func (x *Identity) ProtoReflect() protoreflect.Message {
	mi := &file_orc8r_protos_identity_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Identity.ProtoReflect.Descriptor instead.
func (*Identity) Descriptor() ([]byte, []int) {
	return file_orc8r_protos_identity_proto_rawDescGZIP(), []int{0}
}

func (m *Identity) GetValue() isIdentity_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (x *Identity) GetGateway() *Identity_Gateway {
	if x, ok := x.GetValue().(*Identity_Gateway_); ok {
		return x.Gateway
	}
	return nil
}

func (x *Identity) GetOperator() string {
	if x, ok := x.GetValue().(*Identity_Operator); ok {
		return x.Operator
	}
	return ""
}

func (x *Identity) GetNetwork() string {
	if x, ok := x.GetValue().(*Identity_Network); ok {
		return x.Network
	}
	return ""
}

func (x *Identity) GetWildcard() *Identity_Wildcard {
	if x, ok := x.GetValue().(*Identity_Wildcard_); ok {
		return x.Wildcard
	}
	return nil
}

type isIdentity_Value interface {
	isIdentity_Value()
}

type Identity_Gateway_ struct {
	Gateway *Identity_Gateway `protobuf:"bytes,1,opt,name=gateway,proto3,oneof"`
}

type Identity_Operator struct {
	Operator string `protobuf:"bytes,2,opt,name=operator,proto3,oneof"` // unique operator ID (user name), encoded in CN of operator certificate
}

type Identity_Network struct {
	Network string `protobuf:"bytes,3,opt,name=network,proto3,oneof"` // unique network ID, used by REST access control
}

type Identity_Wildcard_ struct {
	Wildcard *Identity_Wildcard `protobuf:"bytes,11,opt,name=wildcard,proto3,oneof"`
}

func (*Identity_Gateway_) isIdentity_Value() {}

func (*Identity_Operator) isIdentity_Value() {}

func (*Identity_Network) isIdentity_Value() {}

func (*Identity_Wildcard_) isIdentity_Value() {}

// --------------------------------------------------------------------------
// AccessGatewayID uniquely identifies an access gateway across the system
// Based on the usage context it could represent either hardware or logical
// AG ID.
// --------------------------------------------------------------------------
type AccessGatewayID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *AccessGatewayID) Reset() {
	*x = AccessGatewayID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orc8r_protos_identity_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AccessGatewayID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AccessGatewayID) ProtoMessage() {}

func (x *AccessGatewayID) ProtoReflect() protoreflect.Message {
	mi := &file_orc8r_protos_identity_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AccessGatewayID.ProtoReflect.Descriptor instead.
func (*AccessGatewayID) Descriptor() ([]byte, []int) {
	return file_orc8r_protos_identity_proto_rawDescGZIP(), []int{1}
}

func (x *AccessGatewayID) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Identity_Wildcard struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type Identity_Wildcard_Type `protobuf:"varint,1,opt,name=type,proto3,enum=magma.orc8r.Identity_Wildcard_Type" json:"type,omitempty"`
}

func (x *Identity_Wildcard) Reset() {
	*x = Identity_Wildcard{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orc8r_protos_identity_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Identity_Wildcard) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Identity_Wildcard) ProtoMessage() {}

func (x *Identity_Wildcard) ProtoReflect() protoreflect.Message {
	mi := &file_orc8r_protos_identity_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Identity_Wildcard.ProtoReflect.Descriptor instead.
func (*Identity_Wildcard) Descriptor() ([]byte, []int) {
	return file_orc8r_protos_identity_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Identity_Wildcard) GetType() Identity_Wildcard_Type {
	if x != nil {
		return x.Type
	}
	return Identity_Wildcard_Gateway
}

type Identity_Gateway struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	HardwareId string `protobuf:"bytes,1,opt,name=hardware_id,json=hardwareId,proto3" json:"hardware_id,omitempty"` // unique hardware ID of a gateway encoded in CN of GW certificate
	NetworkId  string `protobuf:"bytes,2,opt,name=network_id,json=networkId,proto3" json:"network_id,omitempty"`
	LogicalId  string `protobuf:"bytes,3,opt,name=logical_id,json=logicalId,proto3" json:"logical_id,omitempty"`
}

func (x *Identity_Gateway) Reset() {
	*x = Identity_Gateway{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orc8r_protos_identity_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Identity_Gateway) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Identity_Gateway) ProtoMessage() {}

func (x *Identity_Gateway) ProtoReflect() protoreflect.Message {
	mi := &file_orc8r_protos_identity_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Identity_Gateway.ProtoReflect.Descriptor instead.
func (*Identity_Gateway) Descriptor() ([]byte, []int) {
	return file_orc8r_protos_identity_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Identity_Gateway) GetHardwareId() string {
	if x != nil {
		return x.HardwareId
	}
	return ""
}

func (x *Identity_Gateway) GetNetworkId() string {
	if x != nil {
		return x.NetworkId
	}
	return ""
}

func (x *Identity_Gateway) GetLogicalId() string {
	if x != nil {
		return x.LogicalId
	}
	return ""
}

// Identities list wrapper
type Identity_List struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	List []*Identity `protobuf:"bytes,1,rep,name=list,proto3" json:"list,omitempty"`
}

func (x *Identity_List) Reset() {
	*x = Identity_List{}
	if protoimpl.UnsafeEnabled {
		mi := &file_orc8r_protos_identity_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Identity_List) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Identity_List) ProtoMessage() {}

func (x *Identity_List) ProtoReflect() protoreflect.Message {
	mi := &file_orc8r_protos_identity_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Identity_List.ProtoReflect.Descriptor instead.
func (*Identity_List) Descriptor() ([]byte, []int) {
	return file_orc8r_protos_identity_proto_rawDescGZIP(), []int{0, 2}
}

func (x *Identity_List) GetList() []*Identity {
	if x != nil {
		return x.List
	}
	return nil
}

var File_orc8r_protos_identity_proto protoreflect.FileDescriptor

var file_orc8r_protos_identity_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x69,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x6d,
	0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x22, 0xd8, 0x03, 0x0a, 0x08, 0x49,
	0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x39, 0x0a, 0x07, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1d, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61,
	0x2e, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e,
	0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x48, 0x00, 0x52, 0x07, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x12, 0x1c, 0x0a, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x6f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x12, 0x1a, 0x0a, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x48, 0x00, 0x52, 0x07, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x12, 0x3c, 0x0a, 0x08,
	0x77, 0x69, 0x6c, 0x64, 0x63, 0x61, 0x72, 0x64, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2e, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x57, 0x69, 0x6c, 0x64, 0x63, 0x61, 0x72, 0x64, 0x48, 0x00,
	0x52, 0x08, 0x77, 0x69, 0x6c, 0x64, 0x63, 0x61, 0x72, 0x64, 0x1a, 0x73, 0x0a, 0x08, 0x57, 0x69,
	0x6c, 0x64, 0x63, 0x61, 0x72, 0x64, 0x12, 0x37, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x23, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63,
	0x38, 0x72, 0x2e, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x57, 0x69, 0x6c, 0x64,
	0x63, 0x61, 0x72, 0x64, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74, 0x79, 0x70, 0x65, 0x22,
	0x2e, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0b, 0x0a, 0x07, 0x47, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x10, 0x00, 0x12, 0x0c, 0x0a, 0x08, 0x4f, 0x70, 0x65, 0x72, 0x61, 0x74, 0x6f, 0x72,
	0x10, 0x01, 0x12, 0x0b, 0x0a, 0x07, 0x4e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x10, 0x02, 0x1a,
	0x68, 0x0a, 0x07, 0x47, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x12, 0x1f, 0x0a, 0x0b, 0x68, 0x61,
	0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x68, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6e, 0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x6f,
	0x67, 0x69, 0x63, 0x61, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09,
	0x6c, 0x6f, 0x67, 0x69, 0x63, 0x61, 0x6c, 0x49, 0x64, 0x1a, 0x31, 0x0a, 0x04, 0x4c, 0x69, 0x73,
	0x74, 0x12, 0x29, 0x0a, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x6d, 0x61, 0x67, 0x6d, 0x61, 0x2e, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2e, 0x49, 0x64,
	0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x52, 0x04, 0x6c, 0x69, 0x73, 0x74, 0x42, 0x07, 0x0a, 0x05,
	0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x21, 0x0a, 0x0f, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x47,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x49, 0x44, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x42, 0x1b, 0x5a, 0x19, 0x6d, 0x61, 0x67, 0x6d,
	0x61, 0x2f, 0x6f, 0x72, 0x63, 0x38, 0x72, 0x2f, 0x6c, 0x69, 0x62, 0x2f, 0x67, 0x6f, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_orc8r_protos_identity_proto_rawDescOnce sync.Once
	file_orc8r_protos_identity_proto_rawDescData = file_orc8r_protos_identity_proto_rawDesc
)

func file_orc8r_protos_identity_proto_rawDescGZIP() []byte {
	file_orc8r_protos_identity_proto_rawDescOnce.Do(func() {
		file_orc8r_protos_identity_proto_rawDescData = protoimpl.X.CompressGZIP(file_orc8r_protos_identity_proto_rawDescData)
	})
	return file_orc8r_protos_identity_proto_rawDescData
}

var file_orc8r_protos_identity_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_orc8r_protos_identity_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_orc8r_protos_identity_proto_goTypes = []interface{}{
	(Identity_Wildcard_Type)(0), // 0: magma.orc8r.Identity.Wildcard.Type
	(*Identity)(nil),            // 1: magma.orc8r.Identity
	(*AccessGatewayID)(nil),     // 2: magma.orc8r.AccessGatewayID
	(*Identity_Wildcard)(nil),   // 3: magma.orc8r.Identity.Wildcard
	(*Identity_Gateway)(nil),    // 4: magma.orc8r.Identity.Gateway
	(*Identity_List)(nil),       // 5: magma.orc8r.Identity.List
}
var file_orc8r_protos_identity_proto_depIdxs = []int32{
	4, // 0: magma.orc8r.Identity.gateway:type_name -> magma.orc8r.Identity.Gateway
	3, // 1: magma.orc8r.Identity.wildcard:type_name -> magma.orc8r.Identity.Wildcard
	0, // 2: magma.orc8r.Identity.Wildcard.type:type_name -> magma.orc8r.Identity.Wildcard.Type
	1, // 3: magma.orc8r.Identity.List.list:type_name -> magma.orc8r.Identity
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_orc8r_protos_identity_proto_init() }
func file_orc8r_protos_identity_proto_init() {
	if File_orc8r_protos_identity_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_orc8r_protos_identity_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Identity); i {
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
		file_orc8r_protos_identity_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AccessGatewayID); i {
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
		file_orc8r_protos_identity_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Identity_Wildcard); i {
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
		file_orc8r_protos_identity_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Identity_Gateway); i {
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
		file_orc8r_protos_identity_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Identity_List); i {
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
	file_orc8r_protos_identity_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Identity_Gateway_)(nil),
		(*Identity_Operator)(nil),
		(*Identity_Network)(nil),
		(*Identity_Wildcard_)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_orc8r_protos_identity_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_orc8r_protos_identity_proto_goTypes,
		DependencyIndexes: file_orc8r_protos_identity_proto_depIdxs,
		EnumInfos:         file_orc8r_protos_identity_proto_enumTypes,
		MessageInfos:      file_orc8r_protos_identity_proto_msgTypes,
	}.Build()
	File_orc8r_protos_identity_proto = out.File
	file_orc8r_protos_identity_proto_rawDesc = nil
	file_orc8r_protos_identity_proto_goTypes = nil
	file_orc8r_protos_identity_proto_depIdxs = nil
}
