// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v5.27.3
// source: core_api.proto

package core_api

import (
	_ "github.com/xh-polaris/schedule-core-api/biz/application/dto/http"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_core_api_proto protoreflect.FileDescriptor

var file_core_api_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x11, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f,
	0x61, 0x70, 0x69, 0x1a, 0x0f, 0x68, 0x74, 0x74, 0x70, 0x2f, 0x68, 0x74, 0x74, 0x70, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x63,
	0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x32, 0xaa, 0x06, 0x0a, 0x08, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70,
	0x69, 0x12, 0x58, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x12, 0x1c, 0x2e, 0x73, 0x63,
	0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e,
	0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x73, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x69,
	0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x73, 0x70, 0x22, 0x11, 0xd2, 0xc1, 0x18, 0x0d, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2f, 0x73, 0x69, 0x67, 0x6e, 0x5f, 0x75, 0x70, 0x12, 0x58, 0x0a, 0x06, 0x53,
	0x69, 0x67, 0x6e, 0x49, 0x6e, 0x12, 0x1c, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e,
	0x52, 0x65, 0x71, 0x1a, 0x1d, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x63,
	0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65,
	0x73, 0x70, 0x22, 0x11, 0xd2, 0xc1, 0x18, 0x0d, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x73, 0x69,
	0x67, 0x6e, 0x5f, 0x69, 0x6e, 0x12, 0x73, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x12, 0x24, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x25, 0x2e,
	0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70,
	0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x22, 0x14, 0xd2, 0xc1, 0x18, 0x10, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x12, 0x95, 0x01, 0x0a, 0x18, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x46, 0x72, 0x6f,
	0x6d, 0x4f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x12, 0x2e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x4f, 0x72,
	0x69, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x2f, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x46, 0x72, 0x6f, 0x6d, 0x4f, 0x72,
	0x69, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x22, 0x18, 0xd2, 0xc1, 0x18, 0x14, 0x2f, 0x73,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x6f,
	0x72, 0x69, 0x12, 0x7b, 0x0a, 0x0f, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x73, 0x12, 0x25, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65,
	0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x26, 0x2e, 0x73,
	0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x22, 0x19, 0xd2, 0xc1, 0x18, 0x15, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x2f, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x12,
	0x73, 0x0a, 0x0e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c,
	0x65, 0x12, 0x24, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x72,
	0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x71, 0x1a, 0x25, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x52, 0x65, 0x73, 0x70, 0x22, 0x14,
	0xd2, 0xc1, 0x18, 0x10, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x12, 0x6b, 0x0a, 0x0c, 0x47, 0x65, 0x74, 0x53, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x73, 0x12, 0x22, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2e,
	0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74, 0x53, 0x63, 0x68, 0x65,
	0x64, 0x75, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x71, 0x1a, 0x23, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x2e, 0x47, 0x65, 0x74,
	0x53, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x22, 0x12, 0xf2,
	0xc1, 0x18, 0x0e, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2f, 0x6c, 0x69, 0x73,
	0x74, 0x42, 0x87, 0x01, 0x0a, 0x26, 0x63, 0x6f, 0x6d, 0x2e, 0x78, 0x68, 0x70, 0x6f, 0x6c, 0x61,
	0x72, 0x69, 0x73, 0x2e, 0x69, 0x64, 0x6c, 0x67, 0x65, 0x6e, 0x2e, 0x73, 0x63, 0x68, 0x65, 0x64,
	0x75, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x42, 0x0c, 0x43, 0x6f,
	0x72, 0x65, 0x41, 0x70, 0x69, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x4d, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x78, 0x68, 0x2d, 0x70, 0x6f, 0x6c, 0x61,
	0x72, 0x69, 0x73, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75, 0x6c, 0x65, 0x2d, 0x63, 0x6f, 0x72,
	0x65, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x62, 0x69, 0x7a, 0x2f, 0x61, 0x70, 0x70, 0x6c, 0x69, 0x63,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x64, 0x74, 0x6f, 0x2f, 0x73, 0x63, 0x68, 0x65, 0x64, 0x75,
	0x6c, 0x65, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x5f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var file_core_api_proto_goTypes = []interface{}{
	(*SignUpReq)(nil),                    // 0: schedule.core_api.SignUpReq
	(*SignInReq)(nil),                    // 1: schedule.core_api.SignInReq
	(*CreateScheduleReq)(nil),            // 2: schedule.core_api.CreateScheduleReq
	(*CreateScheduleFromOriginReq)(nil),  // 3: schedule.core_api.CreateScheduleFromOriginReq
	(*CreateSchedulesReq)(nil),           // 4: schedule.core_api.CreateSchedulesReq
	(*UpdateScheduleReq)(nil),            // 5: schedule.core_api.UpdateScheduleReq
	(*GetSchedulesReq)(nil),              // 6: schedule.core_api.GetSchedulesReq
	(*SignUpResp)(nil),                   // 7: schedule.core_api.SignUpResp
	(*SignInResp)(nil),                   // 8: schedule.core_api.SignInResp
	(*CreateScheduleResp)(nil),           // 9: schedule.core_api.CreateScheduleResp
	(*CreateScheduleFromOriginResp)(nil), // 10: schedule.core_api.CreateScheduleFromOriginResp
	(*CreateSchedulesResp)(nil),          // 11: schedule.core_api.CreateSchedulesResp
	(*UpdateScheduleResp)(nil),           // 12: schedule.core_api.UpdateScheduleResp
	(*GetSchedulesResp)(nil),             // 13: schedule.core_api.GetSchedulesResp
}
var file_core_api_proto_depIdxs = []int32{
	0,  // 0: schedule.core_api.core_api.SignUp:input_type -> schedule.core_api.SignUpReq
	1,  // 1: schedule.core_api.core_api.SignIn:input_type -> schedule.core_api.SignInReq
	2,  // 2: schedule.core_api.core_api.CreateSchedule:input_type -> schedule.core_api.CreateScheduleReq
	3,  // 3: schedule.core_api.core_api.CreateScheduleFromOrigin:input_type -> schedule.core_api.CreateScheduleFromOriginReq
	4,  // 4: schedule.core_api.core_api.CreateSchedules:input_type -> schedule.core_api.CreateSchedulesReq
	5,  // 5: schedule.core_api.core_api.UpdateSchedule:input_type -> schedule.core_api.UpdateScheduleReq
	6,  // 6: schedule.core_api.core_api.GetSchedules:input_type -> schedule.core_api.GetSchedulesReq
	7,  // 7: schedule.core_api.core_api.SignUp:output_type -> schedule.core_api.SignUpResp
	8,  // 8: schedule.core_api.core_api.SignIn:output_type -> schedule.core_api.SignInResp
	9,  // 9: schedule.core_api.core_api.CreateSchedule:output_type -> schedule.core_api.CreateScheduleResp
	10, // 10: schedule.core_api.core_api.CreateScheduleFromOrigin:output_type -> schedule.core_api.CreateScheduleFromOriginResp
	11, // 11: schedule.core_api.core_api.CreateSchedules:output_type -> schedule.core_api.CreateSchedulesResp
	12, // 12: schedule.core_api.core_api.UpdateSchedule:output_type -> schedule.core_api.UpdateScheduleResp
	13, // 13: schedule.core_api.core_api.GetSchedules:output_type -> schedule.core_api.GetSchedulesResp
	7,  // [7:14] is the sub-list for method output_type
	0,  // [0:7] is the sub-list for method input_type
	0,  // [0:0] is the sub-list for extension type_name
	0,  // [0:0] is the sub-list for extension extendee
	0,  // [0:0] is the sub-list for field type_name
}

func file_core_api_proto_init() {
	if File_core_api_proto != nil {
		return
	}
	file_schedule_core_api_common_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_core_api_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_core_api_proto_goTypes,
		DependencyIndexes: file_core_api_proto_depIdxs,
	}.Build()
	File_core_api_proto = out.File
	file_core_api_proto_rawDesc = nil
	file_core_api_proto_goTypes = nil
	file_core_api_proto_depIdxs = nil
}
