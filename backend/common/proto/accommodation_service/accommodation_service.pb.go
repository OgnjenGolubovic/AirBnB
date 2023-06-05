// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.30.0
// 	protoc        v4.22.3
// source: accommodation_service/accommodation_service.proto

package accommodation_service

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type Accommodation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name              string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Location          string `protobuf:"bytes,3,opt,name=location,proto3" json:"location,omitempty"`
	Benefits          string `protobuf:"bytes,4,opt,name=benefits,proto3" json:"benefits,omitempty"`
	Photos            string `protobuf:"bytes,5,opt,name=photos,proto3" json:"photos,omitempty"`
	MinGuest          int64  `protobuf:"varint,6,opt,name=minGuest,proto3" json:"minGuest,omitempty"`
	MaxGuest          int64  `protobuf:"varint,7,opt,name=maxGuest,proto3" json:"maxGuest,omitempty"`
	HostId            string `protobuf:"bytes,8,opt,name=hostId,proto3" json:"hostId,omitempty"`
	AutomaticApproval bool   `protobuf:"varint,9,opt,name=automaticApproval,proto3" json:"automaticApproval,omitempty"`
}

func (x *Accommodation) Reset() {
	*x = Accommodation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accommodation_service_accommodation_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Accommodation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Accommodation) ProtoMessage() {}

func (x *Accommodation) ProtoReflect() protoreflect.Message {
	mi := &file_accommodation_service_accommodation_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Accommodation.ProtoReflect.Descriptor instead.
func (*Accommodation) Descriptor() ([]byte, []int) {
	return file_accommodation_service_accommodation_service_proto_rawDescGZIP(), []int{0}
}

func (x *Accommodation) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Accommodation) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Accommodation) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *Accommodation) GetBenefits() string {
	if x != nil {
		return x.Benefits
	}
	return ""
}

func (x *Accommodation) GetPhotos() string {
	if x != nil {
		return x.Photos
	}
	return ""
}

func (x *Accommodation) GetMinGuest() int64 {
	if x != nil {
		return x.MinGuest
	}
	return 0
}

func (x *Accommodation) GetMaxGuest() int64 {
	if x != nil {
		return x.MaxGuest
	}
	return 0
}

func (x *Accommodation) GetHostId() string {
	if x != nil {
		return x.HostId
	}
	return ""
}

func (x *Accommodation) GetAutomaticApproval() bool {
	if x != nil {
		return x.AutomaticApproval
	}
	return false
}

type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetRequest) Reset() {
	*x = GetRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accommodation_service_accommodation_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequest) ProtoMessage() {}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	mi := &file_accommodation_service_accommodation_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequest.ProtoReflect.Descriptor instead.
func (*GetRequest) Descriptor() ([]byte, []int) {
	return file_accommodation_service_accommodation_service_proto_rawDescGZIP(), []int{1}
}

func (x *GetRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Accommodation *Accommodation `protobuf:"bytes,1,opt,name=accommodation,proto3" json:"accommodation,omitempty"`
}

func (x *GetResponse) Reset() {
	*x = GetResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accommodation_service_accommodation_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetResponse) ProtoMessage() {}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	mi := &file_accommodation_service_accommodation_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetResponse.ProtoReflect.Descriptor instead.
func (*GetResponse) Descriptor() ([]byte, []int) {
	return file_accommodation_service_accommodation_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetResponse) GetAccommodation() *Accommodation {
	if x != nil {
		return x.Accommodation
	}
	return nil
}

type DateRange struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartDate string `protobuf:"bytes,1,opt,name=startDate,proto3" json:"startDate,omitempty"`
	EndDate   string `protobuf:"bytes,2,opt,name=endDate,proto3" json:"endDate,omitempty"`
}

func (x *DateRange) Reset() {
	*x = DateRange{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accommodation_service_accommodation_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DateRange) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DateRange) ProtoMessage() {}

func (x *DateRange) ProtoReflect() protoreflect.Message {
	mi := &file_accommodation_service_accommodation_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DateRange.ProtoReflect.Descriptor instead.
func (*DateRange) Descriptor() ([]byte, []int) {
	return file_accommodation_service_accommodation_service_proto_rawDescGZIP(), []int{3}
}

func (x *DateRange) GetStartDate() string {
	if x != nil {
		return x.StartDate
	}
	return ""
}

func (x *DateRange) GetEndDate() string {
	if x != nil {
		return x.EndDate
	}
	return ""
}

type DateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Dates []*DateRange `protobuf:"bytes,1,rep,name=dates,proto3" json:"dates,omitempty"`
}

func (x *DateResponse) Reset() {
	*x = DateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accommodation_service_accommodation_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DateResponse) ProtoMessage() {}

func (x *DateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_accommodation_service_accommodation_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DateResponse.ProtoReflect.Descriptor instead.
func (*DateResponse) Descriptor() ([]byte, []int) {
	return file_accommodation_service_accommodation_service_proto_rawDescGZIP(), []int{4}
}

func (x *DateResponse) GetDates() []*DateRange {
	if x != nil {
		return x.Dates
	}
	return nil
}

type GetAllRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllRequest) Reset() {
	*x = GetAllRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accommodation_service_accommodation_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllRequest) ProtoMessage() {}

func (x *GetAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_accommodation_service_accommodation_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllRequest.ProtoReflect.Descriptor instead.
func (*GetAllRequest) Descriptor() ([]byte, []int) {
	return file_accommodation_service_accommodation_service_proto_rawDescGZIP(), []int{5}
}

type GetAllResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Accommodations []*Accommodation `protobuf:"bytes,1,rep,name=accommodations,proto3" json:"accommodations,omitempty"`
}

func (x *GetAllResponse) Reset() {
	*x = GetAllResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accommodation_service_accommodation_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllResponse) ProtoMessage() {}

func (x *GetAllResponse) ProtoReflect() protoreflect.Message {
	mi := &file_accommodation_service_accommodation_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllResponse.ProtoReflect.Descriptor instead.
func (*GetAllResponse) Descriptor() ([]byte, []int) {
	return file_accommodation_service_accommodation_service_proto_rawDescGZIP(), []int{6}
}

func (x *GetAllResponse) GetAccommodations() []*Accommodation {
	if x != nil {
		return x.Accommodations
	}
	return nil
}

type CreateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Accommodation *Accommodation `protobuf:"bytes,1,opt,name=accommodation,proto3" json:"accommodation,omitempty"`
}

func (x *CreateRequest) Reset() {
	*x = CreateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accommodation_service_accommodation_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateRequest) ProtoMessage() {}

func (x *CreateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_accommodation_service_accommodation_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateRequest.ProtoReflect.Descriptor instead.
func (*CreateRequest) Descriptor() ([]byte, []int) {
	return file_accommodation_service_accommodation_service_proto_rawDescGZIP(), []int{7}
}

func (x *CreateRequest) GetAccommodation() *Accommodation {
	if x != nil {
		return x.Accommodation
	}
	return nil
}

type CreateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Accommodation *Accommodation `protobuf:"bytes,1,opt,name=accommodation,proto3" json:"accommodation,omitempty"`
}

func (x *CreateResponse) Reset() {
	*x = CreateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_accommodation_service_accommodation_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateResponse) ProtoMessage() {}

func (x *CreateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_accommodation_service_accommodation_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateResponse.ProtoReflect.Descriptor instead.
func (*CreateResponse) Descriptor() ([]byte, []int) {
	return file_accommodation_service_accommodation_service_proto_rawDescGZIP(), []int{8}
}

func (x *CreateResponse) GetAccommodation() *Accommodation {
	if x != nil {
		return x.Accommodation
	}
	return nil
}

var File_accommodation_service_accommodation_service_proto protoreflect.FileDescriptor

var file_accommodation_service_accommodation_service_proto_rawDesc = []byte{
	0x0a, 0x31, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61,
	0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0x81, 0x02, 0x0a, 0x0d, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x62, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x62, 0x65, 0x6e, 0x65, 0x66, 0x69, 0x74, 0x73, 0x12, 0x16,
	0x0a, 0x06, 0x70, 0x68, 0x6f, 0x74, 0x6f, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x70, 0x68, 0x6f, 0x74, 0x6f, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x69, 0x6e, 0x47, 0x75, 0x65,
	0x73, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x69, 0x6e, 0x47, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x61, 0x78, 0x47, 0x75, 0x65, 0x73, 0x74, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x6d, 0x61, 0x78, 0x47, 0x75, 0x65, 0x73, 0x74, 0x12, 0x16,
	0x0a, 0x06, 0x68, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x68, 0x6f, 0x73, 0x74, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x11, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61,
	0x74, 0x69, 0x63, 0x41, 0x70, 0x70, 0x72, 0x6f, 0x76, 0x61, 0x6c, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x11, 0x61, 0x75, 0x74, 0x6f, 0x6d, 0x61, 0x74, 0x69, 0x63, 0x41, 0x70, 0x70, 0x72,
	0x6f, 0x76, 0x61, 0x6c, 0x22, 0x1c, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x22, 0x51, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x42, 0x0a, 0x0d, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f,
	0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x43, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x65, 0x52, 0x61, 0x6e,
	0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x65, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x22, 0x3e, 0x0a, 0x0c, 0x44, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x05, 0x64, 0x61,
	0x74, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x61, 0x63, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x44, 0x61, 0x74, 0x65, 0x52, 0x61,
	0x6e, 0x67, 0x65, 0x52, 0x05, 0x64, 0x61, 0x74, 0x65, 0x73, 0x22, 0x0f, 0x0a, 0x0d, 0x47, 0x65,
	0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x56, 0x0a, 0x0e, 0x47,
	0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x44, 0x0a,
	0x0e, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0e, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x22, 0x53, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x42, 0x0a, 0x0d, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x63,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x63, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x61, 0x63, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x54, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x0d, 0x61, 0x63,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1c, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x0d, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x32, 0x99,
	0x03, 0x0a, 0x14, 0x41, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x59, 0x0a, 0x03, 0x47, 0x65, 0x74, 0x12, 0x19,
	0x2e, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47,
	0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x61, 0x63, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1b, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x15, 0x12, 0x13, 0x2f,
	0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x7b, 0x69,
	0x64, 0x7d, 0x12, 0x65, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x46, 0x72, 0x65, 0x65,
	0x44, 0x61, 0x74, 0x65, 0x73, 0x12, 0x19, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1b, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x44, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x1a, 0x82,
	0xd3, 0xe4, 0x93, 0x02, 0x14, 0x12, 0x12, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x67, 0x65, 0x74, 0x12, 0x5d, 0x0a, 0x06, 0x47, 0x65, 0x74,
	0x41, 0x6c, 0x6c, 0x12, 0x1c, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1d, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12, 0x0e, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x6d,
	0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x60, 0x0a, 0x06, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x12, 0x1c, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1d, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x19, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x13, 0x3a, 0x01, 0x2a, 0x22, 0x0e, 0x2f, 0x61, 0x63, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x42, 0x48, 0x5a, 0x46, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x4f, 0x67, 0x6e, 0x6a, 0x65, 0x6e, 0x47,
	0x6f, 0x6c, 0x75, 0x62, 0x6f, 0x76, 0x69, 0x63, 0x2f, 0x41, 0x69, 0x72, 0x42, 0x6e, 0x42, 0x2f,
	0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x61,
	0x63, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x64, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_accommodation_service_accommodation_service_proto_rawDescOnce sync.Once
	file_accommodation_service_accommodation_service_proto_rawDescData = file_accommodation_service_accommodation_service_proto_rawDesc
)

func file_accommodation_service_accommodation_service_proto_rawDescGZIP() []byte {
	file_accommodation_service_accommodation_service_proto_rawDescOnce.Do(func() {
		file_accommodation_service_accommodation_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_accommodation_service_accommodation_service_proto_rawDescData)
	})
	return file_accommodation_service_accommodation_service_proto_rawDescData
}

var file_accommodation_service_accommodation_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_accommodation_service_accommodation_service_proto_goTypes = []interface{}{
	(*Accommodation)(nil),  // 0: accommodation.Accommodation
	(*GetRequest)(nil),     // 1: accommodation.GetRequest
	(*GetResponse)(nil),    // 2: accommodation.GetResponse
	(*DateRange)(nil),      // 3: accommodation.DateRange
	(*DateResponse)(nil),   // 4: accommodation.DateResponse
	(*GetAllRequest)(nil),  // 5: accommodation.GetAllRequest
	(*GetAllResponse)(nil), // 6: accommodation.GetAllResponse
	(*CreateRequest)(nil),  // 7: accommodation.CreateRequest
	(*CreateResponse)(nil), // 8: accommodation.CreateResponse
}
var file_accommodation_service_accommodation_service_proto_depIdxs = []int32{
	0, // 0: accommodation.GetResponse.accommodation:type_name -> accommodation.Accommodation
	3, // 1: accommodation.DateResponse.dates:type_name -> accommodation.DateRange
	0, // 2: accommodation.GetAllResponse.accommodations:type_name -> accommodation.Accommodation
	0, // 3: accommodation.CreateRequest.accommodation:type_name -> accommodation.Accommodation
	0, // 4: accommodation.CreateResponse.accommodation:type_name -> accommodation.Accommodation
	1, // 5: accommodation.AccommodationService.Get:input_type -> accommodation.GetRequest
	1, // 6: accommodation.AccommodationService.GetAllFreeDates:input_type -> accommodation.GetRequest
	5, // 7: accommodation.AccommodationService.GetAll:input_type -> accommodation.GetAllRequest
	7, // 8: accommodation.AccommodationService.Create:input_type -> accommodation.CreateRequest
	2, // 9: accommodation.AccommodationService.Get:output_type -> accommodation.GetResponse
	4, // 10: accommodation.AccommodationService.GetAllFreeDates:output_type -> accommodation.DateResponse
	6, // 11: accommodation.AccommodationService.GetAll:output_type -> accommodation.GetAllResponse
	8, // 12: accommodation.AccommodationService.Create:output_type -> accommodation.CreateResponse
	9, // [9:13] is the sub-list for method output_type
	5, // [5:9] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_accommodation_service_accommodation_service_proto_init() }
func file_accommodation_service_accommodation_service_proto_init() {
	if File_accommodation_service_accommodation_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_accommodation_service_accommodation_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Accommodation); i {
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
		file_accommodation_service_accommodation_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetRequest); i {
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
		file_accommodation_service_accommodation_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetResponse); i {
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
		file_accommodation_service_accommodation_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DateRange); i {
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
		file_accommodation_service_accommodation_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DateResponse); i {
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
		file_accommodation_service_accommodation_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllRequest); i {
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
		file_accommodation_service_accommodation_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllResponse); i {
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
		file_accommodation_service_accommodation_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateRequest); i {
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
		file_accommodation_service_accommodation_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateResponse); i {
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
			RawDescriptor: file_accommodation_service_accommodation_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_accommodation_service_accommodation_service_proto_goTypes,
		DependencyIndexes: file_accommodation_service_accommodation_service_proto_depIdxs,
		MessageInfos:      file_accommodation_service_accommodation_service_proto_msgTypes,
	}.Build()
	File_accommodation_service_accommodation_service_proto = out.File
	file_accommodation_service_accommodation_service_proto_rawDesc = nil
	file_accommodation_service_accommodation_service_proto_goTypes = nil
	file_accommodation_service_accommodation_service_proto_depIdxs = nil
}
