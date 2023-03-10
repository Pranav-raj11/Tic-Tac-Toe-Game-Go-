// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        (unknown)
// source: proto/games/tictactoe/board.proto

package tictactoe

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

type BoardRow struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cells  []int32 `protobuf:"varint,1,rep,packed,name=cells,proto3" json:"cells,omitempty"`
	Length int32   `protobuf:"varint,2,opt,name=length,proto3" json:"length,omitempty"`
}

func (x *BoardRow) Reset() {
	*x = BoardRow{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_games_tictactoe_board_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BoardRow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BoardRow) ProtoMessage() {}

func (x *BoardRow) ProtoReflect() protoreflect.Message {
	mi := &file_proto_games_tictactoe_board_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BoardRow.ProtoReflect.Descriptor instead.
func (*BoardRow) Descriptor() ([]byte, []int) {
	return file_proto_games_tictactoe_board_proto_rawDescGZIP(), []int{0}
}

func (x *BoardRow) GetCells() []int32 {
	if x != nil {
		return x.Cells
	}
	return nil
}

func (x *BoardRow) GetLength() int32 {
	if x != nil {
		return x.Length
	}
	return 0
}

type Board struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cells     []*BoardRow `protobuf:"bytes,1,rep,name=cells,proto3" json:"cells,omitempty"`
	Dimension int32       `protobuf:"varint,2,opt,name=dimension,proto3" json:"dimension,omitempty"`
}

func (x *Board) Reset() {
	*x = Board{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_games_tictactoe_board_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Board) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Board) ProtoMessage() {}

func (x *Board) ProtoReflect() protoreflect.Message {
	mi := &file_proto_games_tictactoe_board_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Board.ProtoReflect.Descriptor instead.
func (*Board) Descriptor() ([]byte, []int) {
	return file_proto_games_tictactoe_board_proto_rawDescGZIP(), []int{1}
}

func (x *Board) GetCells() []*BoardRow {
	if x != nil {
		return x.Cells
	}
	return nil
}

func (x *Board) GetDimension() int32 {
	if x != nil {
		return x.Dimension
	}
	return 0
}

var File_proto_games_tictactoe_board_proto protoreflect.FileDescriptor

var file_proto_games_tictactoe_board_proto_rawDesc = []byte{
	0x0a, 0x21, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x2f, 0x74, 0x69,
	0x63, 0x74, 0x61, 0x63, 0x74, 0x6f, 0x65, 0x2f, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x09, 0x74, 0x69, 0x63, 0x74, 0x61, 0x63, 0x74, 0x6f, 0x65, 0x22, 0x38,
	0x0a, 0x08, 0x42, 0x6f, 0x61, 0x72, 0x64, 0x52, 0x6f, 0x77, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x65,
	0x6c, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x05, 0x52, 0x05, 0x63, 0x65, 0x6c, 0x6c, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x06, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x22, 0x50, 0x0a, 0x05, 0x42, 0x6f, 0x61, 0x72,
	0x64, 0x12, 0x29, 0x0a, 0x05, 0x63, 0x65, 0x6c, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x13, 0x2e, 0x74, 0x69, 0x63, 0x74, 0x61, 0x63, 0x74, 0x6f, 0x65, 0x2e, 0x42, 0x6f, 0x61,
	0x72, 0x64, 0x52, 0x6f, 0x77, 0x52, 0x05, 0x63, 0x65, 0x6c, 0x6c, 0x73, 0x12, 0x1c, 0x0a, 0x09,
	0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x09, 0x64, 0x69, 0x6d, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x6b, 0x65, 0x72, 0x74, 0x6f, 0x78, 0x36,
	0x36, 0x32, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x2d, 0x73, 0x69, 0x74, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x2f, 0x74, 0x69, 0x63, 0x74, 0x61, 0x63, 0x74,
	0x6f, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_games_tictactoe_board_proto_rawDescOnce sync.Once
	file_proto_games_tictactoe_board_proto_rawDescData = file_proto_games_tictactoe_board_proto_rawDesc
)

func file_proto_games_tictactoe_board_proto_rawDescGZIP() []byte {
	file_proto_games_tictactoe_board_proto_rawDescOnce.Do(func() {
		file_proto_games_tictactoe_board_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_games_tictactoe_board_proto_rawDescData)
	})
	return file_proto_games_tictactoe_board_proto_rawDescData
}

var file_proto_games_tictactoe_board_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_games_tictactoe_board_proto_goTypes = []interface{}{
	(*BoardRow)(nil), // 0: tictactoe.BoardRow
	(*Board)(nil),    // 1: tictactoe.Board
}
var file_proto_games_tictactoe_board_proto_depIdxs = []int32{
	0, // 0: tictactoe.Board.cells:type_name -> tictactoe.BoardRow
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_games_tictactoe_board_proto_init() }
func file_proto_games_tictactoe_board_proto_init() {
	if File_proto_games_tictactoe_board_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_games_tictactoe_board_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BoardRow); i {
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
		file_proto_games_tictactoe_board_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Board); i {
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
			RawDescriptor: file_proto_games_tictactoe_board_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_games_tictactoe_board_proto_goTypes,
		DependencyIndexes: file_proto_games_tictactoe_board_proto_depIdxs,
		MessageInfos:      file_proto_games_tictactoe_board_proto_msgTypes,
	}.Build()
	File_proto_games_tictactoe_board_proto = out.File
	file_proto_games_tictactoe_board_proto_rawDesc = nil
	file_proto_games_tictactoe_board_proto_goTypes = nil
	file_proto_games_tictactoe_board_proto_depIdxs = nil
}
