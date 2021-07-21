// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nft/nft.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// BaseNFT defines a non-fungible token
type BaseNFT struct {
	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	URI   string `protobuf:"bytes,3,opt,name=uri,proto3" json:"uri,omitempty"`
	Data  string `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	Owner string `protobuf:"bytes,5,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (m *BaseNFT) Reset()         { *m = BaseNFT{} }
func (m *BaseNFT) String() string { return proto.CompactTextString(m) }
func (*BaseNFT) ProtoMessage()    {}
func (*BaseNFT) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe8ab7e15b7f0646, []int{0}
}
func (m *BaseNFT) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BaseNFT) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BaseNFT.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BaseNFT) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BaseNFT.Merge(m, src)
}
func (m *BaseNFT) XXX_Size() int {
	return m.Size()
}
func (m *BaseNFT) XXX_DiscardUnknown() {
	xxx_messageInfo_BaseNFT.DiscardUnknown(m)
}

var xxx_messageInfo_BaseNFT proto.InternalMessageInfo

// Denom defines a type of NFT
type Denom struct {
	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Schema  string `protobuf:"bytes,3,opt,name=schema,proto3" json:"schema,omitempty"`
	Creator string `protobuf:"bytes,4,opt,name=creator,proto3" json:"creator,omitempty"`
}

func (m *Denom) Reset()         { *m = Denom{} }
func (m *Denom) String() string { return proto.CompactTextString(m) }
func (*Denom) ProtoMessage()    {}
func (*Denom) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe8ab7e15b7f0646, []int{1}
}
func (m *Denom) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Denom) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Denom.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Denom) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Denom.Merge(m, src)
}
func (m *Denom) XXX_Size() int {
	return m.Size()
}
func (m *Denom) XXX_DiscardUnknown() {
	xxx_messageInfo_Denom.DiscardUnknown(m)
}

var xxx_messageInfo_Denom proto.InternalMessageInfo

// IDCollection defines a type of collection with specified ID
type IDCollection struct {
	DenomId  string   `protobuf:"bytes,1,opt,name=denom_id,json=denomId,proto3" json:"denom_id,omitempty" yaml:"denom_id"`
	TokenIds []string `protobuf:"bytes,2,rep,name=token_ids,json=tokenIds,proto3" json:"token_ids,omitempty" yaml:"token_ids"`
}

func (m *IDCollection) Reset()         { *m = IDCollection{} }
func (m *IDCollection) String() string { return proto.CompactTextString(m) }
func (*IDCollection) ProtoMessage()    {}
func (*IDCollection) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe8ab7e15b7f0646, []int{2}
}
func (m *IDCollection) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IDCollection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IDCollection.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IDCollection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IDCollection.Merge(m, src)
}
func (m *IDCollection) XXX_Size() int {
	return m.Size()
}
func (m *IDCollection) XXX_DiscardUnknown() {
	xxx_messageInfo_IDCollection.DiscardUnknown(m)
}

var xxx_messageInfo_IDCollection proto.InternalMessageInfo

// Owner defines a type of owner
type Owner struct {
	Address       string         `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	IDCollections []IDCollection `protobuf:"bytes,2,rep,name=id_collections,json=idCollections,proto3" json:"id_collections" yaml:"idcs"`
}

func (m *Owner) Reset()         { *m = Owner{} }
func (m *Owner) String() string { return proto.CompactTextString(m) }
func (*Owner) ProtoMessage()    {}
func (*Owner) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe8ab7e15b7f0646, []int{3}
}
func (m *Owner) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Owner) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Owner.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Owner) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Owner.Merge(m, src)
}
func (m *Owner) XXX_Size() int {
	return m.Size()
}
func (m *Owner) XXX_DiscardUnknown() {
	xxx_messageInfo_Owner.DiscardUnknown(m)
}

var xxx_messageInfo_Owner proto.InternalMessageInfo

// Collection defines a type of collection
type Collection struct {
	Denom Denom     `protobuf:"bytes,1,opt,name=denom,proto3" json:"denom"`
	NFTs  []BaseNFT `protobuf:"bytes,2,rep,name=nfts,proto3" json:"nfts"`
}

func (m *Collection) Reset()         { *m = Collection{} }
func (m *Collection) String() string { return proto.CompactTextString(m) }
func (*Collection) ProtoMessage()    {}
func (*Collection) Descriptor() ([]byte, []int) {
	return fileDescriptor_fe8ab7e15b7f0646, []int{4}
}
func (m *Collection) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Collection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Collection.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Collection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Collection.Merge(m, src)
}
func (m *Collection) XXX_Size() int {
	return m.Size()
}
func (m *Collection) XXX_DiscardUnknown() {
	xxx_messageInfo_Collection.DiscardUnknown(m)
}

var xxx_messageInfo_Collection proto.InternalMessageInfo

func init() {
	proto.RegisterType((*BaseNFT)(nil), "zebi.cric.nft.BaseNFT")
	proto.RegisterType((*Denom)(nil), "zebi.cric.nft.Denom")
	proto.RegisterType((*IDCollection)(nil), "zebi.cric.nft.IDCollection")
	proto.RegisterType((*Owner)(nil), "zebi.cric.nft.Owner")
	proto.RegisterType((*Collection)(nil), "zebi.cric.nft.Collection")
}

func init() { proto.RegisterFile("nft/nft.proto", fileDescriptor_fe8ab7e15b7f0646) }

var fileDescriptor_fe8ab7e15b7f0646 = []byte{
	// 471 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xb1, 0x6e, 0xdb, 0x3c,
	0x10, 0xc7, 0x25, 0x5b, 0x8a, 0xe3, 0x73, 0x9c, 0xef, 0x03, 0x6b, 0x04, 0x6a, 0x0b, 0x48, 0x86,
	0xd0, 0x21, 0x93, 0xd4, 0xa6, 0x4b, 0x91, 0xa1, 0x83, 0x1a, 0x04, 0xf0, 0x92, 0x02, 0x42, 0xba,
	0x74, 0x31, 0x64, 0x91, 0xb6, 0x89, 0x5a, 0x64, 0x20, 0x32, 0x48, 0xd2, 0xa1, 0xcf, 0xd0, 0xbe,
	0x41, 0x1f, 0xc7, 0x63, 0xc6, 0x4e, 0x42, 0x2b, 0x2f, 0x9d, 0xfd, 0x04, 0x05, 0x49, 0x29, 0x70,
	0x33, 0x75, 0xbb, 0xfb, 0xdf, 0x9d, 0x7e, 0xff, 0x13, 0x0f, 0x86, 0x6c, 0x2e, 0x63, 0x36, 0x97,
	0xd1, 0x55, 0xc9, 0x25, 0x47, 0xc3, 0xcf, 0x64, 0x46, 0xa3, 0xbc, 0xa4, 0x79, 0xc4, 0xe6, 0xf2,
	0xd9, 0x68, 0xc1, 0x17, 0x5c, 0x57, 0x62, 0x15, 0x99, 0xa6, 0xf0, 0x16, 0x7a, 0x49, 0x26, 0xc8,
	0xc5, 0xf9, 0x25, 0x3a, 0x84, 0x0e, 0xc5, 0x9e, 0x3d, 0xb6, 0x8f, 0xfb, 0x69, 0x87, 0x62, 0x84,
	0xc0, 0x61, 0x59, 0x41, 0xbc, 0x8e, 0x56, 0x74, 0x8c, 0x9e, 0x42, 0xf7, 0xba, 0xa4, 0x5e, 0x57,
	0x49, 0x49, 0xaf, 0xae, 0x82, 0xee, 0x87, 0x74, 0x92, 0x2a, 0x4d, 0xb5, 0xe3, 0x4c, 0x66, 0x9e,
	0x63, 0xda, 0x55, 0x8c, 0x46, 0xe0, 0xf2, 0x1b, 0x46, 0x4a, 0xcf, 0xd5, 0xa2, 0x49, 0x4e, 0x9d,
	0xdf, 0xdf, 0x03, 0x3b, 0xcc, 0xc1, 0x3d, 0x23, 0x8c, 0x17, 0xff, 0xc4, 0x3d, 0x82, 0x3d, 0x91,
	0x2f, 0x49, 0x91, 0x19, 0x74, 0xda, 0x64, 0xc8, 0x83, 0x5e, 0x5e, 0x92, 0x4c, 0xf2, 0xb2, 0xe1,
	0xb6, 0x69, 0x03, 0xb9, 0x81, 0x83, 0xc9, 0xd9, 0x3b, 0xbe, 0x5a, 0x91, 0x5c, 0x52, 0xce, 0x50,
	0x04, 0xfb, 0x58, 0x41, 0xa7, 0x2d, 0x31, 0x79, 0xb2, 0xad, 0x82, 0xff, 0xee, 0xb2, 0x62, 0x75,
	0x1a, 0xb6, 0x95, 0x30, 0xed, 0xe9, 0x70, 0x82, 0xd1, 0x2b, 0xe8, 0x4b, 0xfe, 0x89, 0xb0, 0x29,
	0xc5, 0xc2, 0xeb, 0x8c, 0xbb, 0xc7, 0xfd, 0x64, 0xb4, 0xad, 0x82, 0xff, 0xcd, 0xc0, 0x43, 0x29,
	0x4c, 0xf7, 0x75, 0x3c, 0xc1, 0xa2, 0x01, 0x7f, 0xb3, 0xc1, 0x7d, 0xaf, 0xb6, 0x55, 0x16, 0x33,
	0x8c, 0x4b, 0x22, 0x44, 0xb3, 0x63, 0x9b, 0xa2, 0x25, 0x1c, 0x52, 0x3c, 0xcd, 0x1f, 0xdc, 0x19,
	0xc2, 0xe0, 0xe4, 0x79, 0xf4, 0xd7, 0xcb, 0x45, 0xbb, 0x1b, 0x24, 0x2f, 0xd6, 0x55, 0x60, 0xd5,
	0x55, 0x30, 0xdc, 0x55, 0xc5, 0xb6, 0x0a, 0x06, 0xc6, 0x13, 0xc5, 0xb9, 0x08, 0xd3, 0x21, 0xc5,
	0x3b, 0xd5, 0xc6, 0xd3, 0x17, 0x80, 0x9d, 0x5f, 0xf1, 0x12, 0x5c, 0xbd, 0xa5, 0x76, 0x35, 0x38,
	0x19, 0x3d, 0x82, 0xea, 0xb7, 0x49, 0x1c, 0x45, 0x4b, 0x4d, 0x23, 0x7a, 0x03, 0x0e, 0x9b, 0xcb,
	0xd6, 0xe5, 0xd1, 0xa3, 0x81, 0xe6, 0x8c, 0x92, 0x83, 0xc6, 0xa0, 0x73, 0x71, 0x7e, 0x29, 0x52,
	0x3d, 0x61, 0xf8, 0xc9, 0xdb, 0xf5, 0x2f, 0xdf, 0x5a, 0xd7, 0xbe, 0x7d, 0x5f, 0xfb, 0xf6, 0xcf,
	0xda, 0xb7, 0xbf, 0x6e, 0x7c, 0xeb, 0x7e, 0xe3, 0x5b, 0x3f, 0x36, 0xbe, 0xf5, 0x71, 0xbc, 0xa0,
	0x72, 0x79, 0x3d, 0x8b, 0x72, 0x5e, 0xc4, 0xea, 0xcb, 0xb1, 0xfa, 0x72, 0x7c, 0xab, 0x0e, 0x3a,
	0x96, 0x77, 0x57, 0x44, 0xcc, 0xf6, 0xf4, 0xc9, 0xbe, 0xfe, 0x13, 0x00, 0x00, 0xff, 0xff, 0x2a,
	0x1f, 0x33, 0x51, 0xe8, 0x02, 0x00, 0x00,
}

func (this *BaseNFT) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*BaseNFT)
	if !ok {
		that2, ok := that.(BaseNFT)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Id != that1.Id {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.URI != that1.URI {
		return false
	}
	if this.Data != that1.Data {
		return false
	}
	if this.Owner != that1.Owner {
		return false
	}
	return true
}
func (this *Denom) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Denom)
	if !ok {
		that2, ok := that.(Denom)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Id != that1.Id {
		return false
	}
	if this.Name != that1.Name {
		return false
	}
	if this.Schema != that1.Schema {
		return false
	}
	if this.Creator != that1.Creator {
		return false
	}
	return true
}
func (this *IDCollection) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*IDCollection)
	if !ok {
		that2, ok := that.(IDCollection)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.DenomId != that1.DenomId {
		return false
	}
	if len(this.TokenIds) != len(that1.TokenIds) {
		return false
	}
	for i := range this.TokenIds {
		if this.TokenIds[i] != that1.TokenIds[i] {
			return false
		}
	}
	return true
}
func (this *Owner) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Owner)
	if !ok {
		that2, ok := that.(Owner)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Address != that1.Address {
		return false
	}
	if len(this.IDCollections) != len(that1.IDCollections) {
		return false
	}
	for i := range this.IDCollections {
		if !this.IDCollections[i].Equal(&that1.IDCollections[i]) {
			return false
		}
	}
	return true
}
func (this *Collection) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Collection)
	if !ok {
		that2, ok := that.(Collection)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if !this.Denom.Equal(&that1.Denom) {
		return false
	}
	if len(this.NFTs) != len(that1.NFTs) {
		return false
	}
	for i := range this.NFTs {
		if !this.NFTs[i].Equal(&that1.NFTs[i]) {
			return false
		}
	}
	return true
}
func (m *BaseNFT) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BaseNFT) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BaseNFT) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintNft(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Data) > 0 {
		i -= len(m.Data)
		copy(dAtA[i:], m.Data)
		i = encodeVarintNft(dAtA, i, uint64(len(m.Data)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.URI) > 0 {
		i -= len(m.URI)
		copy(dAtA[i:], m.URI)
		i = encodeVarintNft(dAtA, i, uint64(len(m.URI)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintNft(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintNft(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Denom) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Denom) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Denom) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintNft(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Schema) > 0 {
		i -= len(m.Schema)
		copy(dAtA[i:], m.Schema)
		i = encodeVarintNft(dAtA, i, uint64(len(m.Schema)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintNft(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintNft(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *IDCollection) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IDCollection) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IDCollection) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.TokenIds) > 0 {
		for iNdEx := len(m.TokenIds) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.TokenIds[iNdEx])
			copy(dAtA[i:], m.TokenIds[iNdEx])
			i = encodeVarintNft(dAtA, i, uint64(len(m.TokenIds[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.DenomId) > 0 {
		i -= len(m.DenomId)
		copy(dAtA[i:], m.DenomId)
		i = encodeVarintNft(dAtA, i, uint64(len(m.DenomId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Owner) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Owner) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Owner) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.IDCollections) > 0 {
		for iNdEx := len(m.IDCollections) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.IDCollections[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintNft(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintNft(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Collection) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Collection) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Collection) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.NFTs) > 0 {
		for iNdEx := len(m.NFTs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.NFTs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintNft(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	{
		size, err := m.Denom.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintNft(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintNft(dAtA []byte, offset int, v uint64) int {
	offset -= sovNft(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BaseNFT) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovNft(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovNft(uint64(l))
	}
	l = len(m.URI)
	if l > 0 {
		n += 1 + l + sovNft(uint64(l))
	}
	l = len(m.Data)
	if l > 0 {
		n += 1 + l + sovNft(uint64(l))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovNft(uint64(l))
	}
	return n
}

func (m *Denom) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovNft(uint64(l))
	}
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovNft(uint64(l))
	}
	l = len(m.Schema)
	if l > 0 {
		n += 1 + l + sovNft(uint64(l))
	}
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovNft(uint64(l))
	}
	return n
}

func (m *IDCollection) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.DenomId)
	if l > 0 {
		n += 1 + l + sovNft(uint64(l))
	}
	if len(m.TokenIds) > 0 {
		for _, s := range m.TokenIds {
			l = len(s)
			n += 1 + l + sovNft(uint64(l))
		}
	}
	return n
}

func (m *Owner) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovNft(uint64(l))
	}
	if len(m.IDCollections) > 0 {
		for _, e := range m.IDCollections {
			l = e.Size()
			n += 1 + l + sovNft(uint64(l))
		}
	}
	return n
}

func (m *Collection) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Denom.Size()
	n += 1 + l + sovNft(uint64(l))
	if len(m.NFTs) > 0 {
		for _, e := range m.NFTs {
			l = e.Size()
			n += 1 + l + sovNft(uint64(l))
		}
	}
	return n
}

func sovNft(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozNft(x uint64) (n int) {
	return sovNft(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BaseNFT) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNft
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: BaseNFT: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BaseNFT: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNft
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNft
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field URI", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNft
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.URI = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNft
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Data = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNft
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNft(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNft
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Denom) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNft
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Denom: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Denom: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNft
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNft
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Schema", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNft
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Schema = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNft
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNft(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNft
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *IDCollection) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNft
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: IDCollection: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IDCollection: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DenomId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNft
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.DenomId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenIds", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNft
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenIds = append(m.TokenIds, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNft(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNft
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Owner) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNft
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Owner: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Owner: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthNft
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IDCollections", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthNft
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IDCollections = append(m.IDCollections, IDCollection{})
			if err := m.IDCollections[len(m.IDCollections)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNft(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNft
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Collection) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowNft
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Collection: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Collection: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Denom", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthNft
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Denom.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NFTs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowNft
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthNft
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthNft
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NFTs = append(m.NFTs, BaseNFT{})
			if err := m.NFTs[len(m.NFTs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipNft(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthNft
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipNft(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowNft
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowNft
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowNft
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthNft
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupNft
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthNft
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthNft        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowNft          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupNft = fmt.Errorf("proto: unexpected end of group")
)