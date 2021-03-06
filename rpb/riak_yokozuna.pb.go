// Code generated by protoc-gen-go.
// source: riak_yokozuna.proto
// DO NOT EDIT!

package rpb

import proto "code.google.com/p/goprotobuf/proto"
import json "encoding/json"
import math "math"

// Reference proto, json, and math imports to suppress error if they are not otherwise used.
var _ = proto.Marshal
var _ = &json.SyntaxError{}
var _ = math.Inf

type RpbYokozunaIndex struct {
	Name             []byte  `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Schema           []byte  `protobuf:"bytes,2,opt,name=schema" json:"schema,omitempty"`
	NVal             *uint32 `protobuf:"varint,3,opt,name=n_val" json:"n_val,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RpbYokozunaIndex) Reset()         { *m = RpbYokozunaIndex{} }
func (m *RpbYokozunaIndex) String() string { return proto.CompactTextString(m) }
func (*RpbYokozunaIndex) ProtoMessage()    {}

func (m *RpbYokozunaIndex) GetName() []byte {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *RpbYokozunaIndex) GetSchema() []byte {
	if m != nil {
		return m.Schema
	}
	return nil
}

func (m *RpbYokozunaIndex) GetNVal() uint32 {
	if m != nil && m.NVal != nil {
		return *m.NVal
	}
	return 0
}

type RpbYokozunaIndexGetReq struct {
	Name             []byte `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RpbYokozunaIndexGetReq) Reset()         { *m = RpbYokozunaIndexGetReq{} }
func (m *RpbYokozunaIndexGetReq) String() string { return proto.CompactTextString(m) }
func (*RpbYokozunaIndexGetReq) ProtoMessage()    {}

func (m *RpbYokozunaIndexGetReq) GetName() []byte {
	if m != nil {
		return m.Name
	}
	return nil
}

type RpbYokozunaIndexGetResp struct {
	Index            []*RpbYokozunaIndex `protobuf:"bytes,1,rep,name=index" json:"index,omitempty"`
	XXX_unrecognized []byte              `json:"-"`
}

func (m *RpbYokozunaIndexGetResp) Reset()         { *m = RpbYokozunaIndexGetResp{} }
func (m *RpbYokozunaIndexGetResp) String() string { return proto.CompactTextString(m) }
func (*RpbYokozunaIndexGetResp) ProtoMessage()    {}

func (m *RpbYokozunaIndexGetResp) GetIndex() []*RpbYokozunaIndex {
	if m != nil {
		return m.Index
	}
	return nil
}

type RpbYokozunaIndexPutReq struct {
	Index            *RpbYokozunaIndex `protobuf:"bytes,1,req,name=index" json:"index,omitempty"`
	XXX_unrecognized []byte            `json:"-"`
}

func (m *RpbYokozunaIndexPutReq) Reset()         { *m = RpbYokozunaIndexPutReq{} }
func (m *RpbYokozunaIndexPutReq) String() string { return proto.CompactTextString(m) }
func (*RpbYokozunaIndexPutReq) ProtoMessage()    {}

func (m *RpbYokozunaIndexPutReq) GetIndex() *RpbYokozunaIndex {
	if m != nil {
		return m.Index
	}
	return nil
}

type RpbYokozunaIndexDeleteReq struct {
	Name             []byte `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RpbYokozunaIndexDeleteReq) Reset()         { *m = RpbYokozunaIndexDeleteReq{} }
func (m *RpbYokozunaIndexDeleteReq) String() string { return proto.CompactTextString(m) }
func (*RpbYokozunaIndexDeleteReq) ProtoMessage()    {}

func (m *RpbYokozunaIndexDeleteReq) GetName() []byte {
	if m != nil {
		return m.Name
	}
	return nil
}

type RpbYokozunaSchema struct {
	Name             []byte `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Content          []byte `protobuf:"bytes,2,opt,name=content" json:"content,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RpbYokozunaSchema) Reset()         { *m = RpbYokozunaSchema{} }
func (m *RpbYokozunaSchema) String() string { return proto.CompactTextString(m) }
func (*RpbYokozunaSchema) ProtoMessage()    {}

func (m *RpbYokozunaSchema) GetName() []byte {
	if m != nil {
		return m.Name
	}
	return nil
}

func (m *RpbYokozunaSchema) GetContent() []byte {
	if m != nil {
		return m.Content
	}
	return nil
}

type RpbYokozunaSchemaPutReq struct {
	Schema           *RpbYokozunaSchema `protobuf:"bytes,1,req,name=schema" json:"schema,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *RpbYokozunaSchemaPutReq) Reset()         { *m = RpbYokozunaSchemaPutReq{} }
func (m *RpbYokozunaSchemaPutReq) String() string { return proto.CompactTextString(m) }
func (*RpbYokozunaSchemaPutReq) ProtoMessage()    {}

func (m *RpbYokozunaSchemaPutReq) GetSchema() *RpbYokozunaSchema {
	if m != nil {
		return m.Schema
	}
	return nil
}

type RpbYokozunaSchemaGetReq struct {
	Name             []byte `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *RpbYokozunaSchemaGetReq) Reset()         { *m = RpbYokozunaSchemaGetReq{} }
func (m *RpbYokozunaSchemaGetReq) String() string { return proto.CompactTextString(m) }
func (*RpbYokozunaSchemaGetReq) ProtoMessage()    {}

func (m *RpbYokozunaSchemaGetReq) GetName() []byte {
	if m != nil {
		return m.Name
	}
	return nil
}

type RpbYokozunaSchemaGetResp struct {
	Schema           *RpbYokozunaSchema `protobuf:"bytes,1,req,name=schema" json:"schema,omitempty"`
	XXX_unrecognized []byte             `json:"-"`
}

func (m *RpbYokozunaSchemaGetResp) Reset()         { *m = RpbYokozunaSchemaGetResp{} }
func (m *RpbYokozunaSchemaGetResp) String() string { return proto.CompactTextString(m) }
func (*RpbYokozunaSchemaGetResp) ProtoMessage()    {}

func (m *RpbYokozunaSchemaGetResp) GetSchema() *RpbYokozunaSchema {
	if m != nil {
		return m.Schema
	}
	return nil
}

func init() {
}
