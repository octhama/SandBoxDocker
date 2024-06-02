// Code generated by protoc-gen-go. DO NOT EDIT.
// source: vtctldata.proto

package vtctldata

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	logutil "vitess.io/vitess/go/vt/proto/logutil"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// ExecuteVtctlCommandRequest is the payload for ExecuteVtctlCommand.
// timeouts are in nanoseconds.
type ExecuteVtctlCommandRequest struct {
	Args                 []string `protobuf:"bytes,1,rep,name=args,proto3" json:"args,omitempty"`
	ActionTimeout        int64    `protobuf:"varint,2,opt,name=action_timeout,json=actionTimeout,proto3" json:"action_timeout,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExecuteVtctlCommandRequest) Reset()         { *m = ExecuteVtctlCommandRequest{} }
func (m *ExecuteVtctlCommandRequest) String() string { return proto.CompactTextString(m) }
func (*ExecuteVtctlCommandRequest) ProtoMessage()    {}
func (*ExecuteVtctlCommandRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f41247b323a1ab2e, []int{0}
}

func (m *ExecuteVtctlCommandRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecuteVtctlCommandRequest.Unmarshal(m, b)
}
func (m *ExecuteVtctlCommandRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecuteVtctlCommandRequest.Marshal(b, m, deterministic)
}
func (m *ExecuteVtctlCommandRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecuteVtctlCommandRequest.Merge(m, src)
}
func (m *ExecuteVtctlCommandRequest) XXX_Size() int {
	return xxx_messageInfo_ExecuteVtctlCommandRequest.Size(m)
}
func (m *ExecuteVtctlCommandRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecuteVtctlCommandRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ExecuteVtctlCommandRequest proto.InternalMessageInfo

func (m *ExecuteVtctlCommandRequest) GetArgs() []string {
	if m != nil {
		return m.Args
	}
	return nil
}

func (m *ExecuteVtctlCommandRequest) GetActionTimeout() int64 {
	if m != nil {
		return m.ActionTimeout
	}
	return 0
}

// ExecuteVtctlCommandResponse is streamed back by ExecuteVtctlCommand.
type ExecuteVtctlCommandResponse struct {
	Event                *logutil.Event `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *ExecuteVtctlCommandResponse) Reset()         { *m = ExecuteVtctlCommandResponse{} }
func (m *ExecuteVtctlCommandResponse) String() string { return proto.CompactTextString(m) }
func (*ExecuteVtctlCommandResponse) ProtoMessage()    {}
func (*ExecuteVtctlCommandResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f41247b323a1ab2e, []int{1}
}

func (m *ExecuteVtctlCommandResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExecuteVtctlCommandResponse.Unmarshal(m, b)
}
func (m *ExecuteVtctlCommandResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExecuteVtctlCommandResponse.Marshal(b, m, deterministic)
}
func (m *ExecuteVtctlCommandResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExecuteVtctlCommandResponse.Merge(m, src)
}
func (m *ExecuteVtctlCommandResponse) XXX_Size() int {
	return xxx_messageInfo_ExecuteVtctlCommandResponse.Size(m)
}
func (m *ExecuteVtctlCommandResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ExecuteVtctlCommandResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ExecuteVtctlCommandResponse proto.InternalMessageInfo

func (m *ExecuteVtctlCommandResponse) GetEvent() *logutil.Event {
	if m != nil {
		return m.Event
	}
	return nil
}

// TableMaterializeSttings contains the settings for one table.
type TableMaterializeSettings struct {
	TargetTable string `protobuf:"bytes,1,opt,name=target_table,json=targetTable,proto3" json:"target_table,omitempty"`
	// source_expression is a select statement.
	SourceExpression string `protobuf:"bytes,2,opt,name=source_expression,json=sourceExpression,proto3" json:"source_expression,omitempty"`
	// create_ddl contains the DDL to create the target table.
	// If empty, the target table must already exist.
	// if "copy", the target table DDL is the same as the source table.
	CreateDdl            string   `protobuf:"bytes,3,opt,name=create_ddl,json=createDdl,proto3" json:"create_ddl,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TableMaterializeSettings) Reset()         { *m = TableMaterializeSettings{} }
func (m *TableMaterializeSettings) String() string { return proto.CompactTextString(m) }
func (*TableMaterializeSettings) ProtoMessage()    {}
func (*TableMaterializeSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_f41247b323a1ab2e, []int{2}
}

func (m *TableMaterializeSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TableMaterializeSettings.Unmarshal(m, b)
}
func (m *TableMaterializeSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TableMaterializeSettings.Marshal(b, m, deterministic)
}
func (m *TableMaterializeSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TableMaterializeSettings.Merge(m, src)
}
func (m *TableMaterializeSettings) XXX_Size() int {
	return xxx_messageInfo_TableMaterializeSettings.Size(m)
}
func (m *TableMaterializeSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_TableMaterializeSettings.DiscardUnknown(m)
}

var xxx_messageInfo_TableMaterializeSettings proto.InternalMessageInfo

func (m *TableMaterializeSettings) GetTargetTable() string {
	if m != nil {
		return m.TargetTable
	}
	return ""
}

func (m *TableMaterializeSettings) GetSourceExpression() string {
	if m != nil {
		return m.SourceExpression
	}
	return ""
}

func (m *TableMaterializeSettings) GetCreateDdl() string {
	if m != nil {
		return m.CreateDdl
	}
	return ""
}

// MaterializeSettings contains the settings for the Materialize command.
type MaterializeSettings struct {
	// workflow is the name of the workflow.
	Workflow       string `protobuf:"bytes,1,opt,name=workflow,proto3" json:"workflow,omitempty"`
	SourceKeyspace string `protobuf:"bytes,2,opt,name=source_keyspace,json=sourceKeyspace,proto3" json:"source_keyspace,omitempty"`
	TargetKeyspace string `protobuf:"bytes,3,opt,name=target_keyspace,json=targetKeyspace,proto3" json:"target_keyspace,omitempty"`
	// stop_after_copy specifies if vreplication should be stopped after copying.
	StopAfterCopy bool                        `protobuf:"varint,4,opt,name=stop_after_copy,json=stopAfterCopy,proto3" json:"stop_after_copy,omitempty"`
	TableSettings []*TableMaterializeSettings `protobuf:"bytes,5,rep,name=table_settings,json=tableSettings,proto3" json:"table_settings,omitempty"`
	// optional parameters.
	Cell                 string   `protobuf:"bytes,6,opt,name=cell,proto3" json:"cell,omitempty"`
	TabletTypes          string   `protobuf:"bytes,7,opt,name=tablet_types,json=tabletTypes,proto3" json:"tablet_types,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MaterializeSettings) Reset()         { *m = MaterializeSettings{} }
func (m *MaterializeSettings) String() string { return proto.CompactTextString(m) }
func (*MaterializeSettings) ProtoMessage()    {}
func (*MaterializeSettings) Descriptor() ([]byte, []int) {
	return fileDescriptor_f41247b323a1ab2e, []int{3}
}

func (m *MaterializeSettings) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MaterializeSettings.Unmarshal(m, b)
}
func (m *MaterializeSettings) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MaterializeSettings.Marshal(b, m, deterministic)
}
func (m *MaterializeSettings) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MaterializeSettings.Merge(m, src)
}
func (m *MaterializeSettings) XXX_Size() int {
	return xxx_messageInfo_MaterializeSettings.Size(m)
}
func (m *MaterializeSettings) XXX_DiscardUnknown() {
	xxx_messageInfo_MaterializeSettings.DiscardUnknown(m)
}

var xxx_messageInfo_MaterializeSettings proto.InternalMessageInfo

func (m *MaterializeSettings) GetWorkflow() string {
	if m != nil {
		return m.Workflow
	}
	return ""
}

func (m *MaterializeSettings) GetSourceKeyspace() string {
	if m != nil {
		return m.SourceKeyspace
	}
	return ""
}

func (m *MaterializeSettings) GetTargetKeyspace() string {
	if m != nil {
		return m.TargetKeyspace
	}
	return ""
}

func (m *MaterializeSettings) GetStopAfterCopy() bool {
	if m != nil {
		return m.StopAfterCopy
	}
	return false
}

func (m *MaterializeSettings) GetTableSettings() []*TableMaterializeSettings {
	if m != nil {
		return m.TableSettings
	}
	return nil
}

func (m *MaterializeSettings) GetCell() string {
	if m != nil {
		return m.Cell
	}
	return ""
}

func (m *MaterializeSettings) GetTabletTypes() string {
	if m != nil {
		return m.TabletTypes
	}
	return ""
}

func init() {
	proto.RegisterType((*ExecuteVtctlCommandRequest)(nil), "vtctldata.ExecuteVtctlCommandRequest")
	proto.RegisterType((*ExecuteVtctlCommandResponse)(nil), "vtctldata.ExecuteVtctlCommandResponse")
	proto.RegisterType((*TableMaterializeSettings)(nil), "vtctldata.TableMaterializeSettings")
	proto.RegisterType((*MaterializeSettings)(nil), "vtctldata.MaterializeSettings")
}

func init() { proto.RegisterFile("vtctldata.proto", fileDescriptor_f41247b323a1ab2e) }

var fileDescriptor_f41247b323a1ab2e = []byte{
	// 422 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x92, 0xd1, 0x6e, 0xd3, 0x30,
	0x14, 0x86, 0x95, 0xb5, 0x1b, 0xeb, 0x29, 0x4d, 0xc1, 0xdc, 0x58, 0x45, 0x48, 0xa1, 0xc0, 0x88,
	0x84, 0xd4, 0x48, 0xe3, 0x09, 0xa0, 0xf4, 0x06, 0xc4, 0x4d, 0xa8, 0x40, 0xe2, 0x26, 0x72, 0x93,
	0xb3, 0xc8, 0x9a, 0x1b, 0x07, 0xfb, 0xa4, 0x5b, 0x79, 0x03, 0x5e, 0x86, 0x67, 0x44, 0xb6, 0xb3,
	0x70, 0xb3, 0xdd, 0x1d, 0x7f, 0xe7, 0xb7, 0xfd, 0x9f, 0x5f, 0x07, 0xe6, 0x07, 0x2a, 0x49, 0x55,
	0x82, 0xc4, 0xaa, 0x35, 0x9a, 0x34, 0x9b, 0x0c, 0x60, 0x31, 0x53, 0xba, 0xee, 0x48, 0xaa, 0xd0,
	0x59, 0xfe, 0x80, 0xc5, 0xe6, 0x16, 0xcb, 0x8e, 0xf0, 0xbb, 0x93, 0xac, 0xf5, 0x7e, 0x2f, 0x9a,
	0x2a, 0xc7, 0x5f, 0x1d, 0x5a, 0x62, 0x0c, 0xc6, 0xc2, 0xd4, 0x96, 0x47, 0xc9, 0x28, 0x9d, 0xe4,
	0xbe, 0x66, 0x6f, 0x20, 0x16, 0x25, 0x49, 0xdd, 0x14, 0x24, 0xf7, 0xa8, 0x3b, 0xe2, 0x27, 0x49,
	0x94, 0x8e, 0xf2, 0x59, 0xa0, 0xdb, 0x00, 0x97, 0x6b, 0x78, 0x7e, 0xef, 0xc3, 0xb6, 0xd5, 0x8d,
	0x45, 0xf6, 0x1a, 0x4e, 0xf1, 0x80, 0x0d, 0xf1, 0x28, 0x89, 0xd2, 0xe9, 0x65, 0xbc, 0xba, 0xb3,
	0xb5, 0x71, 0x34, 0x0f, 0xcd, 0xe5, 0x9f, 0x08, 0xf8, 0x56, 0xec, 0x14, 0x7e, 0x15, 0x84, 0x46,
	0x0a, 0x25, 0x7f, 0xe3, 0x37, 0x24, 0x92, 0x4d, 0x6d, 0xd9, 0x4b, 0x78, 0x4c, 0xc2, 0xd4, 0x48,
	0x05, 0x39, 0x89, 0x7f, 0x69, 0x92, 0x4f, 0x03, 0xf3, 0xb7, 0xd8, 0x3b, 0x78, 0x6a, 0x75, 0x67,
	0x4a, 0x2c, 0xf0, 0xb6, 0x35, 0x68, 0xad, 0xd4, 0x8d, 0xb7, 0x3b, 0xc9, 0x9f, 0x84, 0xc6, 0x66,
	0xe0, 0xec, 0x05, 0x40, 0x69, 0x50, 0x10, 0x16, 0x55, 0xa5, 0xf8, 0xc8, 0xab, 0x26, 0x81, 0x7c,
	0xaa, 0xd4, 0xf2, 0xef, 0x09, 0x3c, 0xbb, 0xcf, 0xc6, 0x02, 0xce, 0x6f, 0xb4, 0xb9, 0xbe, 0x52,
	0xfa, 0xa6, 0xb7, 0x30, 0x9c, 0xd9, 0x5b, 0x98, 0xf7, 0xff, 0x5f, 0xe3, 0xd1, 0xb6, 0xa2, 0xc4,
	0xfe, 0xf7, 0x38, 0xe0, 0x2f, 0x3d, 0x75, 0xc2, 0x7e, 0x96, 0x41, 0x18, 0x0c, 0xc4, 0x01, 0x0f,
	0xc2, 0x0b, 0x98, 0x5b, 0xd2, 0x6d, 0x21, 0xae, 0x08, 0x4d, 0x51, 0xea, 0xf6, 0xc8, 0xc7, 0x49,
	0x94, 0x9e, 0xe7, 0x33, 0x87, 0x3f, 0x38, 0xba, 0xd6, 0xed, 0x91, 0x7d, 0x86, 0xd8, 0xa7, 0x52,
	0xd8, 0xde, 0x27, 0x3f, 0x4d, 0x46, 0xe9, 0xf4, 0xf2, 0xd5, 0xea, 0xff, 0x6e, 0x3c, 0x94, 0x6c,
	0x3e, 0xf3, 0x57, 0x87, 0x09, 0x19, 0x8c, 0x4b, 0x54, 0x8a, 0x9f, 0x79, 0x47, 0xbe, 0x0e, 0xe1,
	0xef, 0x94, 0x0b, 0xff, 0xd8, 0xa2, 0xe5, 0x8f, 0xee, 0xc2, 0x77, 0x6c, 0xeb, 0xd0, 0xc7, 0xf4,
	0xe7, 0xc5, 0x41, 0x12, 0x5a, 0xbb, 0x92, 0x3a, 0x0b, 0x55, 0x56, 0xeb, 0xec, 0x40, 0x99, 0x5f,
	0xbd, 0x6c, 0x30, 0xb2, 0x3b, 0xf3, 0xe0, 0xfd, 0xbf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x46, 0x37,
	0xd0, 0x53, 0xb8, 0x02, 0x00, 0x00,
}