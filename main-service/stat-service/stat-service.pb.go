// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.4
// source: stat-service.proto

package stat_service

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

type GetViewsAndLikesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId string `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
}

func (x *GetViewsAndLikesRequest) Reset() {
	*x = GetViewsAndLikesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stat_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetViewsAndLikesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetViewsAndLikesRequest) ProtoMessage() {}

func (x *GetViewsAndLikesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stat_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetViewsAndLikesRequest.ProtoReflect.Descriptor instead.
func (*GetViewsAndLikesRequest) Descriptor() ([]byte, []int) {
	return file_stat_service_proto_rawDescGZIP(), []int{0}
}

func (x *GetViewsAndLikesRequest) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

type GetViewsAndLikesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Views  int32  `protobuf:"varint,1,opt,name=views,proto3" json:"views,omitempty"`
	Likes  int32  `protobuf:"varint,2,opt,name=likes,proto3" json:"likes,omitempty"`
	TaskId string `protobuf:"bytes,3,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
}

func (x *GetViewsAndLikesResponse) Reset() {
	*x = GetViewsAndLikesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stat_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetViewsAndLikesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetViewsAndLikesResponse) ProtoMessage() {}

func (x *GetViewsAndLikesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stat_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetViewsAndLikesResponse.ProtoReflect.Descriptor instead.
func (*GetViewsAndLikesResponse) Descriptor() ([]byte, []int) {
	return file_stat_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetViewsAndLikesResponse) GetViews() int32 {
	if x != nil {
		return x.Views
	}
	return 0
}

func (x *GetViewsAndLikesResponse) GetLikes() int32 {
	if x != nil {
		return x.Likes
	}
	return 0
}

func (x *GetViewsAndLikesResponse) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

type GetTop5LikedOrViewedTasksRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Likes bool `protobuf:"varint,1,opt,name=likes,proto3" json:"likes,omitempty"`
}

func (x *GetTop5LikedOrViewedTasksRequest) Reset() {
	*x = GetTop5LikedOrViewedTasksRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stat_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTop5LikedOrViewedTasksRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTop5LikedOrViewedTasksRequest) ProtoMessage() {}

func (x *GetTop5LikedOrViewedTasksRequest) ProtoReflect() protoreflect.Message {
	mi := &file_stat_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTop5LikedOrViewedTasksRequest.ProtoReflect.Descriptor instead.
func (*GetTop5LikedOrViewedTasksRequest) Descriptor() ([]byte, []int) {
	return file_stat_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetTop5LikedOrViewedTasksRequest) GetLikes() bool {
	if x != nil {
		return x.Likes
	}
	return false
}

type GetTop5LikedOrViewedTasksResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tasks []*GetTop5LikedOrViewedTasksResponse_TaskIdAndCount `protobuf:"bytes,1,rep,name=tasks,proto3" json:"tasks,omitempty"`
}

func (x *GetTop5LikedOrViewedTasksResponse) Reset() {
	*x = GetTop5LikedOrViewedTasksResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stat_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTop5LikedOrViewedTasksResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTop5LikedOrViewedTasksResponse) ProtoMessage() {}

func (x *GetTop5LikedOrViewedTasksResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stat_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTop5LikedOrViewedTasksResponse.ProtoReflect.Descriptor instead.
func (*GetTop5LikedOrViewedTasksResponse) Descriptor() ([]byte, []int) {
	return file_stat_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetTop5LikedOrViewedTasksResponse) GetTasks() []*GetTop5LikedOrViewedTasksResponse_TaskIdAndCount {
	if x != nil {
		return x.Tasks
	}
	return nil
}

type GetTop3LikedUsersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserIds []*GetTop3LikedUsersResponse_UserIdAndLikeCount `protobuf:"bytes,1,rep,name=user_ids,json=userIds,proto3" json:"user_ids,omitempty"`
}

func (x *GetTop3LikedUsersResponse) Reset() {
	*x = GetTop3LikedUsersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stat_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTop3LikedUsersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTop3LikedUsersResponse) ProtoMessage() {}

func (x *GetTop3LikedUsersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_stat_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTop3LikedUsersResponse.ProtoReflect.Descriptor instead.
func (*GetTop3LikedUsersResponse) Descriptor() ([]byte, []int) {
	return file_stat_service_proto_rawDescGZIP(), []int{4}
}

func (x *GetTop3LikedUsersResponse) GetUserIds() []*GetTop3LikedUsersResponse_UserIdAndLikeCount {
	if x != nil {
		return x.UserIds
	}
	return nil
}

type GetTop5LikedOrViewedTasksResponse_TaskIdAndCount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId string `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	Count  int32  `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *GetTop5LikedOrViewedTasksResponse_TaskIdAndCount) Reset() {
	*x = GetTop5LikedOrViewedTasksResponse_TaskIdAndCount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stat_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTop5LikedOrViewedTasksResponse_TaskIdAndCount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTop5LikedOrViewedTasksResponse_TaskIdAndCount) ProtoMessage() {}

func (x *GetTop5LikedOrViewedTasksResponse_TaskIdAndCount) ProtoReflect() protoreflect.Message {
	mi := &file_stat_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTop5LikedOrViewedTasksResponse_TaskIdAndCount.ProtoReflect.Descriptor instead.
func (*GetTop5LikedOrViewedTasksResponse_TaskIdAndCount) Descriptor() ([]byte, []int) {
	return file_stat_service_proto_rawDescGZIP(), []int{3, 0}
}

func (x *GetTop5LikedOrViewedTasksResponse_TaskIdAndCount) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *GetTop5LikedOrViewedTasksResponse_TaskIdAndCount) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

type GetTop3LikedUsersResponse_UserIdAndLikeCount struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	LikeCount int32  `protobuf:"varint,2,opt,name=like_count,json=likeCount,proto3" json:"like_count,omitempty"`
}

func (x *GetTop3LikedUsersResponse_UserIdAndLikeCount) Reset() {
	*x = GetTop3LikedUsersResponse_UserIdAndLikeCount{}
	if protoimpl.UnsafeEnabled {
		mi := &file_stat_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetTop3LikedUsersResponse_UserIdAndLikeCount) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTop3LikedUsersResponse_UserIdAndLikeCount) ProtoMessage() {}

func (x *GetTop3LikedUsersResponse_UserIdAndLikeCount) ProtoReflect() protoreflect.Message {
	mi := &file_stat_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTop3LikedUsersResponse_UserIdAndLikeCount.ProtoReflect.Descriptor instead.
func (*GetTop3LikedUsersResponse_UserIdAndLikeCount) Descriptor() ([]byte, []int) {
	return file_stat_service_proto_rawDescGZIP(), []int{4, 0}
}

func (x *GetTop3LikedUsersResponse_UserIdAndLikeCount) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *GetTop3LikedUsersResponse_UserIdAndLikeCount) GetLikeCount() int32 {
	if x != nil {
		return x.LikeCount
	}
	return 0
}

var File_stat_service_proto protoreflect.FileDescriptor

var file_stat_service_proto_rawDesc = []byte{
	0x0a, 0x12, 0x73, 0x74, 0x61, 0x74, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x32, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x56, 0x69, 0x65, 0x77, 0x73, 0x41, 0x6e, 0x64,
	0x4c, 0x69, 0x6b, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x74,
	0x61, 0x73, 0x6b, 0x49, 0x64, 0x22, 0x5f, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x56, 0x69, 0x65, 0x77,
	0x73, 0x41, 0x6e, 0x64, 0x4c, 0x69, 0x6b, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x69, 0x65, 0x77, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05,
	0x52, 0x05, 0x76, 0x69, 0x65, 0x77, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69, 0x6b, 0x65, 0x73,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6b, 0x65, 0x73, 0x12, 0x17, 0x0a,
	0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x22, 0x38, 0x0a, 0x20, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70,
	0x35, 0x4c, 0x69, 0x6b, 0x65, 0x64, 0x4f, 0x72, 0x56, 0x69, 0x65, 0x77, 0x65, 0x64, 0x54, 0x61,
	0x73, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69,
	0x6b, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x05, 0x6c, 0x69, 0x6b, 0x65, 0x73,
	0x22, 0xad, 0x01, 0x0a, 0x21, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x35, 0x4c, 0x69, 0x6b, 0x65,
	0x64, 0x4f, 0x72, 0x56, 0x69, 0x65, 0x77, 0x65, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x47, 0x0a, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x35, 0x4c,
	0x69, 0x6b, 0x65, 0x64, 0x4f, 0x72, 0x56, 0x69, 0x65, 0x77, 0x65, 0x64, 0x54, 0x61, 0x73, 0x6b,
	0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x64,
	0x41, 0x6e, 0x64, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x1a,
	0x3f, 0x0a, 0x0e, 0x54, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x41, 0x6e, 0x64, 0x43, 0x6f, 0x75, 0x6e,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x22, 0xb3, 0x01, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x33, 0x4c, 0x69, 0x6b, 0x65,
	0x64, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x48,
	0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x2d, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x33, 0x4c, 0x69, 0x6b, 0x65, 0x64, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x55, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x41, 0x6e, 0x64, 0x4c, 0x69, 0x6b, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x73, 0x1a, 0x4c, 0x0a, 0x12, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x41, 0x6e, 0x64, 0x4c, 0x69, 0x6b, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x6c, 0x69, 0x6b, 0x65, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x6c, 0x69, 0x6b,
	0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x32, 0x89, 0x02, 0x0a, 0x0b, 0x53, 0x74, 0x61, 0x74, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x56, 0x69, 0x65,
	0x77, 0x73, 0x41, 0x6e, 0x64, 0x4c, 0x69, 0x6b, 0x65, 0x73, 0x12, 0x18, 0x2e, 0x47, 0x65, 0x74,
	0x56, 0x69, 0x65, 0x77, 0x73, 0x41, 0x6e, 0x64, 0x4c, 0x69, 0x6b, 0x65, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x69, 0x65, 0x77, 0x73, 0x41,
	0x6e, 0x64, 0x4c, 0x69, 0x6b, 0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x00, 0x12, 0x64, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x35, 0x4c, 0x69, 0x6b, 0x65,
	0x64, 0x4f, 0x72, 0x56, 0x69, 0x65, 0x77, 0x65, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x12, 0x21,
	0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x35, 0x4c, 0x69, 0x6b, 0x65, 0x64, 0x4f, 0x72, 0x56,
	0x69, 0x65, 0x77, 0x65, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x22, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x35, 0x4c, 0x69, 0x6b, 0x65, 0x64,
	0x4f, 0x72, 0x56, 0x69, 0x65, 0x77, 0x65, 0x64, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x54, 0x6f,
	0x70, 0x33, 0x4c, 0x69, 0x6b, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x16, 0x2e, 0x67,
	0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x1a, 0x1a, 0x2e, 0x47, 0x65, 0x74, 0x54, 0x6f, 0x70, 0x33, 0x4c, 0x69,
	0x6b, 0x65, 0x64, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x00, 0x42, 0x2b, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2d, 0x74, 0x72, 0x61, 0x63, 0x6b,
	0x65, 0x72, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_stat_service_proto_rawDescOnce sync.Once
	file_stat_service_proto_rawDescData = file_stat_service_proto_rawDesc
)

func file_stat_service_proto_rawDescGZIP() []byte {
	file_stat_service_proto_rawDescOnce.Do(func() {
		file_stat_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_stat_service_proto_rawDescData)
	})
	return file_stat_service_proto_rawDescData
}

var file_stat_service_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_stat_service_proto_goTypes = []interface{}{
	(*GetViewsAndLikesRequest)(nil),                          // 0: GetViewsAndLikesRequest
	(*GetViewsAndLikesResponse)(nil),                         // 1: GetViewsAndLikesResponse
	(*GetTop5LikedOrViewedTasksRequest)(nil),                 // 2: GetTop5LikedOrViewedTasksRequest
	(*GetTop5LikedOrViewedTasksResponse)(nil),                // 3: GetTop5LikedOrViewedTasksResponse
	(*GetTop3LikedUsersResponse)(nil),                        // 4: GetTop3LikedUsersResponse
	(*GetTop5LikedOrViewedTasksResponse_TaskIdAndCount)(nil), // 5: GetTop5LikedOrViewedTasksResponse.TaskIdAndCount
	(*GetTop3LikedUsersResponse_UserIdAndLikeCount)(nil),     // 6: GetTop3LikedUsersResponse.UserIdAndLikeCount
	(*emptypb.Empty)(nil),                                    // 7: google.protobuf.Empty
}
var file_stat_service_proto_depIdxs = []int32{
	5, // 0: GetTop5LikedOrViewedTasksResponse.tasks:type_name -> GetTop5LikedOrViewedTasksResponse.TaskIdAndCount
	6, // 1: GetTop3LikedUsersResponse.user_ids:type_name -> GetTop3LikedUsersResponse.UserIdAndLikeCount
	0, // 2: StatService.GetViewsAndLikes:input_type -> GetViewsAndLikesRequest
	2, // 3: StatService.GetTop5LikedOrViewedTasks:input_type -> GetTop5LikedOrViewedTasksRequest
	7, // 4: StatService.GetTop3LikedUsers:input_type -> google.protobuf.Empty
	1, // 5: StatService.GetViewsAndLikes:output_type -> GetViewsAndLikesResponse
	3, // 6: StatService.GetTop5LikedOrViewedTasks:output_type -> GetTop5LikedOrViewedTasksResponse
	4, // 7: StatService.GetTop3LikedUsers:output_type -> GetTop3LikedUsersResponse
	5, // [5:8] is the sub-list for method output_type
	2, // [2:5] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_stat_service_proto_init() }
func file_stat_service_proto_init() {
	if File_stat_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_stat_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetViewsAndLikesRequest); i {
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
		file_stat_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetViewsAndLikesResponse); i {
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
		file_stat_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTop5LikedOrViewedTasksRequest); i {
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
		file_stat_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTop5LikedOrViewedTasksResponse); i {
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
		file_stat_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTop3LikedUsersResponse); i {
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
		file_stat_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTop5LikedOrViewedTasksResponse_TaskIdAndCount); i {
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
		file_stat_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetTop3LikedUsersResponse_UserIdAndLikeCount); i {
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
			RawDescriptor: file_stat_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_stat_service_proto_goTypes,
		DependencyIndexes: file_stat_service_proto_depIdxs,
		MessageInfos:      file_stat_service_proto_msgTypes,
	}.Build()
	File_stat_service_proto = out.File
	file_stat_service_proto_rawDesc = nil
	file_stat_service_proto_goTypes = nil
	file_stat_service_proto_depIdxs = nil
}
