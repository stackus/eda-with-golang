// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.1
// source: event_message.proto

package jetstream

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EventMessage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               string                          `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	EventName        string                          `protobuf:"bytes,2,opt,name=event_name,json=eventName,proto3" json:"event_name,omitempty"`
	AggregateName    string                          `protobuf:"bytes,3,opt,name=aggregate_name,json=aggregateName,proto3" json:"aggregate_name,omitempty"`
	AggregateId      string                          `protobuf:"bytes,4,opt,name=aggregate_id,json=aggregateId,proto3" json:"aggregate_id,omitempty"`
	AggregateVersion int32                           `protobuf:"varint,5,opt,name=aggregate_version,json=aggregateVersion,proto3" json:"aggregate_version,omitempty"`
	OccurredAt       *timestamppb.Timestamp          `protobuf:"bytes,6,opt,name=occurred_at,json=occurredAt,proto3" json:"occurred_at,omitempty"`
	Headers          map[string]*EventMessage_Values `protobuf:"bytes,7,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Data             []byte                          `protobuf:"bytes,8,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *EventMessage) Reset() {
	*x = EventMessage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_message_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventMessage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventMessage) ProtoMessage() {}

func (x *EventMessage) ProtoReflect() protoreflect.Message {
	mi := &file_event_message_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventMessage.ProtoReflect.Descriptor instead.
func (*EventMessage) Descriptor() ([]byte, []int) {
	return file_event_message_proto_rawDescGZIP(), []int{0}
}

func (x *EventMessage) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *EventMessage) GetEventName() string {
	if x != nil {
		return x.EventName
	}
	return ""
}

func (x *EventMessage) GetAggregateName() string {
	if x != nil {
		return x.AggregateName
	}
	return ""
}

func (x *EventMessage) GetAggregateId() string {
	if x != nil {
		return x.AggregateId
	}
	return ""
}

func (x *EventMessage) GetAggregateVersion() int32 {
	if x != nil {
		return x.AggregateVersion
	}
	return 0
}

func (x *EventMessage) GetOccurredAt() *timestamppb.Timestamp {
	if x != nil {
		return x.OccurredAt
	}
	return nil
}

func (x *EventMessage) GetHeaders() map[string]*EventMessage_Values {
	if x != nil {
		return x.Headers
	}
	return nil
}

func (x *EventMessage) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

type EventMessage_Values struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Values []string `protobuf:"bytes,1,rep,name=values,proto3" json:"values,omitempty"`
}

func (x *EventMessage_Values) Reset() {
	*x = EventMessage_Values{}
	if protoimpl.UnsafeEnabled {
		mi := &file_event_message_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EventMessage_Values) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EventMessage_Values) ProtoMessage() {}

func (x *EventMessage_Values) ProtoReflect() protoreflect.Message {
	mi := &file_event_message_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EventMessage_Values.ProtoReflect.Descriptor instead.
func (*EventMessage_Values) Descriptor() ([]byte, []int) {
	return file_event_message_proto_rawDescGZIP(), []int{0, 0}
}

func (x *EventMessage_Values) GetValues() []string {
	if x != nil {
		return x.Values
	}
	return nil
}

var File_event_message_proto protoreflect.FileDescriptor

var file_event_message_proto_rawDesc = []byte{
	0x0a, 0x13, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x6a, 0x65, 0x74, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0xc3, 0x03, 0x0a, 0x0c, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x5f, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x67, 0x67, 0x72, 0x65,
	0x67, 0x61, 0x74, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x61, 0x67, 0x67, 0x72,
	0x65, 0x67, 0x61, 0x74, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x11, 0x61,
	0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x65, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x10, 0x61, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74,
	0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x3b, 0x0a, 0x0b, 0x6f, 0x63, 0x63, 0x75,
	0x72, 0x72, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e,
	0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e,
	0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x6f, 0x63, 0x63, 0x75, 0x72,
	0x72, 0x65, 0x64, 0x41, 0x74, 0x12, 0x3e, 0x0a, 0x07, 0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73,
	0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x6a, 0x65, 0x74, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e,
	0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07, 0x68, 0x65,
	0x61, 0x64, 0x65, 0x72, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x08, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x1a, 0x20, 0x0a, 0x06, 0x56, 0x61, 0x6c,
	0x75, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x09, 0x52, 0x06, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x1a, 0x5a, 0x0a, 0x0c, 0x48,
	0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x34, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x6a,
	0x65, 0x74, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x73, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x42, 0x8c, 0x01, 0x0a, 0x0d, 0x63, 0x6f, 0x6d, 0x2e,
	0x6a, 0x65, 0x74, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x42, 0x11, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x24,
	0x65, 0x64, 0x61, 0x2d, 0x69, 0x6e, 0x2d, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x63, 0x68,
	0x37, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x6a, 0x65, 0x74, 0x73, 0x74,
	0x72, 0x65, 0x61, 0x6d, 0xa2, 0x02, 0x03, 0x4a, 0x58, 0x58, 0xaa, 0x02, 0x09, 0x4a, 0x65, 0x74,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0xca, 0x02, 0x09, 0x4a, 0x65, 0x74, 0x73, 0x74, 0x72, 0x65,
	0x61, 0x6d, 0xe2, 0x02, 0x15, 0x4a, 0x65, 0x74, 0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x5c, 0x47,
	0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x09, 0x4a, 0x65, 0x74,
	0x73, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_event_message_proto_rawDescOnce sync.Once
	file_event_message_proto_rawDescData = file_event_message_proto_rawDesc
)

func file_event_message_proto_rawDescGZIP() []byte {
	file_event_message_proto_rawDescOnce.Do(func() {
		file_event_message_proto_rawDescData = protoimpl.X.CompressGZIP(file_event_message_proto_rawDescData)
	})
	return file_event_message_proto_rawDescData
}

var file_event_message_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_event_message_proto_goTypes = []interface{}{
	(*EventMessage)(nil),          // 0: jetstream.EventMessage
	(*EventMessage_Values)(nil),   // 1: jetstream.EventMessage.Values
	nil,                           // 2: jetstream.EventMessage.HeadersEntry
	(*timestamppb.Timestamp)(nil), // 3: google.protobuf.Timestamp
}
var file_event_message_proto_depIdxs = []int32{
	3, // 0: jetstream.EventMessage.occurred_at:type_name -> google.protobuf.Timestamp
	2, // 1: jetstream.EventMessage.headers:type_name -> jetstream.EventMessage.HeadersEntry
	1, // 2: jetstream.EventMessage.HeadersEntry.value:type_name -> jetstream.EventMessage.Values
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_event_message_proto_init() }
func file_event_message_proto_init() {
	if File_event_message_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_event_message_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventMessage); i {
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
		file_event_message_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EventMessage_Values); i {
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
			RawDescriptor: file_event_message_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_event_message_proto_goTypes,
		DependencyIndexes: file_event_message_proto_depIdxs,
		MessageInfos:      file_event_message_proto_msgTypes,
	}.Build()
	File_event_message_proto = out.File
	file_event_message_proto_rawDesc = nil
	file_event_message_proto_goTypes = nil
	file_event_message_proto_depIdxs = nil
}
