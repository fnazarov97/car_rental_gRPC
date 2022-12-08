// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.9
// source: car.proto

package car_service

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

type CreateCarRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Model   string `protobuf:"bytes,1,opt,name=model,proto3" json:"model,omitempty"`
	Color   string `protobuf:"bytes,2,opt,name=color,proto3" json:"color,omitempty"`
	Year    string `protobuf:"bytes,3,opt,name=year,proto3" json:"year,omitempty"`
	Mileage string `protobuf:"bytes,4,opt,name=mileage,proto3" json:"mileage,omitempty"`
	BrandId string `protobuf:"bytes,5,opt,name=brand_id,json=brandId,proto3" json:"brand_id,omitempty"`
}

func (x *CreateCarRequest) Reset() {
	*x = CreateCarRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCarRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCarRequest) ProtoMessage() {}

func (x *CreateCarRequest) ProtoReflect() protoreflect.Message {
	mi := &file_car_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCarRequest.ProtoReflect.Descriptor instead.
func (*CreateCarRequest) Descriptor() ([]byte, []int) {
	return file_car_proto_rawDescGZIP(), []int{0}
}

func (x *CreateCarRequest) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *CreateCarRequest) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

func (x *CreateCarRequest) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

func (x *CreateCarRequest) GetMileage() string {
	if x != nil {
		return x.Mileage
	}
	return ""
}

func (x *CreateCarRequest) GetBrandId() string {
	if x != nil {
		return x.BrandId
	}
	return ""
}

type Car struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CarId     string `protobuf:"bytes,1,opt,name=car_id,json=carId,proto3" json:"car_id,omitempty"`
	Model     string `protobuf:"bytes,2,opt,name=model,proto3" json:"model,omitempty"`
	Color     string `protobuf:"bytes,3,opt,name=color,proto3" json:"color,omitempty"`
	Year      string `protobuf:"bytes,4,opt,name=year,proto3" json:"year,omitempty"`
	Mileage   string `protobuf:"bytes,5,opt,name=mileage,proto3" json:"mileage,omitempty"`
	BrandId   string `protobuf:"bytes,6,opt,name=brand_id,json=brandId,proto3" json:"brand_id,omitempty"`
	CreatedAt string `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *Car) Reset() {
	*x = Car{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Car) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Car) ProtoMessage() {}

func (x *Car) ProtoReflect() protoreflect.Message {
	mi := &file_car_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Car.ProtoReflect.Descriptor instead.
func (*Car) Descriptor() ([]byte, []int) {
	return file_car_proto_rawDescGZIP(), []int{1}
}

func (x *Car) GetCarId() string {
	if x != nil {
		return x.CarId
	}
	return ""
}

func (x *Car) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *Car) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

func (x *Car) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

func (x *Car) GetMileage() string {
	if x != nil {
		return x.Mileage
	}
	return ""
}

func (x *Car) GetBrandId() string {
	if x != nil {
		return x.BrandId
	}
	return ""
}

func (x *Car) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *Car) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type UpdateCarRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Model   string `protobuf:"bytes,1,opt,name=model,proto3" json:"model,omitempty"`
	Color   string `protobuf:"bytes,2,opt,name=color,proto3" json:"color,omitempty"`
	Year    string `protobuf:"bytes,3,opt,name=year,proto3" json:"year,omitempty"`
	Mileage string `protobuf:"bytes,4,opt,name=mileage,proto3" json:"mileage,omitempty"`
	BrandId string `protobuf:"bytes,5,opt,name=brand_id,json=brandId,proto3" json:"brand_id,omitempty"`
}

func (x *UpdateCarRequest) Reset() {
	*x = UpdateCarRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateCarRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateCarRequest) ProtoMessage() {}

func (x *UpdateCarRequest) ProtoReflect() protoreflect.Message {
	mi := &file_car_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateCarRequest.ProtoReflect.Descriptor instead.
func (*UpdateCarRequest) Descriptor() ([]byte, []int) {
	return file_car_proto_rawDescGZIP(), []int{2}
}

func (x *UpdateCarRequest) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *UpdateCarRequest) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

func (x *UpdateCarRequest) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

func (x *UpdateCarRequest) GetMileage() string {
	if x != nil {
		return x.Mileage
	}
	return ""
}

func (x *UpdateCarRequest) GetBrandId() string {
	if x != nil {
		return x.BrandId
	}
	return ""
}

type DeleteCarRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteCarRequest) Reset() {
	*x = DeleteCarRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCarRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCarRequest) ProtoMessage() {}

func (x *DeleteCarRequest) ProtoReflect() protoreflect.Message {
	mi := &file_car_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCarRequest.ProtoReflect.Descriptor instead.
func (*DeleteCarRequest) Descriptor() ([]byte, []int) {
	return file_car_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteCarRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetCarByIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetCarByIDRequest) Reset() {
	*x = GetCarByIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCarByIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCarByIDRequest) ProtoMessage() {}

func (x *GetCarByIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_car_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCarByIDRequest.ProtoReflect.Descriptor instead.
func (*GetCarByIDRequest) Descriptor() ([]byte, []int) {
	return file_car_proto_rawDescGZIP(), []int{4}
}

func (x *GetCarByIDRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetCarByIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	CarId     string                    `protobuf:"bytes,1,opt,name=car_id,json=carId,proto3" json:"car_id,omitempty"`
	Model     string                    `protobuf:"bytes,2,opt,name=model,proto3" json:"model,omitempty"`
	Color     string                    `protobuf:"bytes,3,opt,name=color,proto3" json:"color,omitempty"`
	Year      string                    `protobuf:"bytes,4,opt,name=year,proto3" json:"year,omitempty"`
	Mileage   string                    `protobuf:"bytes,5,opt,name=mileage,proto3" json:"mileage,omitempty"`
	Brand     *GetCarByIDResponse_Brand `protobuf:"bytes,6,opt,name=brand,proto3" json:"brand,omitempty"`
	CreatedAt string                    `protobuf:"bytes,7,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt string                    `protobuf:"bytes,8,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *GetCarByIDResponse) Reset() {
	*x = GetCarByIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCarByIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCarByIDResponse) ProtoMessage() {}

func (x *GetCarByIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_car_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCarByIDResponse.ProtoReflect.Descriptor instead.
func (*GetCarByIDResponse) Descriptor() ([]byte, []int) {
	return file_car_proto_rawDescGZIP(), []int{5}
}

func (x *GetCarByIDResponse) GetCarId() string {
	if x != nil {
		return x.CarId
	}
	return ""
}

func (x *GetCarByIDResponse) GetModel() string {
	if x != nil {
		return x.Model
	}
	return ""
}

func (x *GetCarByIDResponse) GetColor() string {
	if x != nil {
		return x.Color
	}
	return ""
}

func (x *GetCarByIDResponse) GetYear() string {
	if x != nil {
		return x.Year
	}
	return ""
}

func (x *GetCarByIDResponse) GetMileage() string {
	if x != nil {
		return x.Mileage
	}
	return ""
}

func (x *GetCarByIDResponse) GetBrand() *GetCarByIDResponse_Brand {
	if x != nil {
		return x.Brand
	}
	return nil
}

func (x *GetCarByIDResponse) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *GetCarByIDResponse) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

type GetCarListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Offset int32  `protobuf:"varint,1,opt,name=offset,proto3" json:"offset,omitempty"`
	Limit  int32  `protobuf:"varint,2,opt,name=limit,proto3" json:"limit,omitempty"`
	Search string `protobuf:"bytes,3,opt,name=search,proto3" json:"search,omitempty"`
}

func (x *GetCarListRequest) Reset() {
	*x = GetCarListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCarListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCarListRequest) ProtoMessage() {}

func (x *GetCarListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_car_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCarListRequest.ProtoReflect.Descriptor instead.
func (*GetCarListRequest) Descriptor() ([]byte, []int) {
	return file_car_proto_rawDescGZIP(), []int{6}
}

func (x *GetCarListRequest) GetOffset() int32 {
	if x != nil {
		return x.Offset
	}
	return 0
}

func (x *GetCarListRequest) GetLimit() int32 {
	if x != nil {
		return x.Limit
	}
	return 0
}

func (x *GetCarListRequest) GetSearch() string {
	if x != nil {
		return x.Search
	}
	return ""
}

type GetCarListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Cars []*Car `protobuf:"bytes,1,rep,name=Cars,proto3" json:"Cars,omitempty"`
}

func (x *GetCarListResponse) Reset() {
	*x = GetCarListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCarListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCarListResponse) ProtoMessage() {}

func (x *GetCarListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_car_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCarListResponse.ProtoReflect.Descriptor instead.
func (*GetCarListResponse) Descriptor() ([]byte, []int) {
	return file_car_proto_rawDescGZIP(), []int{7}
}

func (x *GetCarListResponse) GetCars() []*Car {
	if x != nil {
		return x.Cars
	}
	return nil
}

type GetCarByIDResponse_Brand struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BrandId     string `protobuf:"bytes,1,opt,name=brand_id,json=brandId,proto3" json:"brand_id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Discription string `protobuf:"bytes,3,opt,name=discription,proto3" json:"discription,omitempty"`
	CreatedAt   string `protobuf:"bytes,4,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	UpdatedAt   string `protobuf:"bytes,5,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
}

func (x *GetCarByIDResponse_Brand) Reset() {
	*x = GetCarByIDResponse_Brand{}
	if protoimpl.UnsafeEnabled {
		mi := &file_car_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCarByIDResponse_Brand) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCarByIDResponse_Brand) ProtoMessage() {}

func (x *GetCarByIDResponse_Brand) ProtoReflect() protoreflect.Message {
	mi := &file_car_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCarByIDResponse_Brand.ProtoReflect.Descriptor instead.
func (*GetCarByIDResponse_Brand) Descriptor() ([]byte, []int) {
	return file_car_proto_rawDescGZIP(), []int{5, 0}
}

func (x *GetCarByIDResponse_Brand) GetBrandId() string {
	if x != nil {
		return x.BrandId
	}
	return ""
}

func (x *GetCarByIDResponse_Brand) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetCarByIDResponse_Brand) GetDiscription() string {
	if x != nil {
		return x.Discription
	}
	return ""
}

func (x *GetCarByIDResponse_Brand) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *GetCarByIDResponse_Brand) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

var File_car_proto protoreflect.FileDescriptor

var file_car_proto_rawDesc = []byte{
	0x0a, 0x09, 0x63, 0x61, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x67, 0x65, 0x6e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x87, 0x01, 0x0a, 0x10, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x43, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f,
	0x64, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x69,
	0x6c, 0x65, 0x61, 0x67, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x69, 0x6c,
	0x65, 0x61, 0x67, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x5f, 0x69, 0x64,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x49, 0x64, 0x22,
	0xcf, 0x01, 0x0a, 0x03, 0x43, 0x61, 0x72, 0x12, 0x15, 0x0a, 0x06, 0x63, 0x61, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x61, 0x72, 0x49, 0x64, 0x12, 0x14,
	0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65,
	0x61, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x69, 0x6c, 0x65, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x69, 0x6c, 0x65, 0x61, 0x67, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x72, 0x61, 0x6e,
	0x64, 0x5f, 0x69, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x62, 0x72, 0x61, 0x6e,
	0x64, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x22, 0x87, 0x01, 0x0a, 0x10, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05,
	0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x6c,
	0x6f, 0x72, 0x12, 0x12, 0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x79, 0x65, 0x61, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x69, 0x6c, 0x65, 0x61, 0x67,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x69, 0x6c, 0x65, 0x61, 0x67, 0x65,
	0x12, 0x19, 0x0a, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x49, 0x64, 0x22, 0x22, 0x0a, 0x10, 0x44,
	0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22,
	0x23, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x42, 0x79, 0x49, 0x44, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x96, 0x03, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x42,
	0x79, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x15, 0x0a, 0x06, 0x63,
	0x61, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x61, 0x72,
	0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x6c, 0x6f,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x63, 0x6f, 0x6c, 0x6f, 0x72, 0x12, 0x12,
	0x0a, 0x04, 0x79, 0x65, 0x61, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x79, 0x65,
	0x61, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x69, 0x6c, 0x65, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x69, 0x6c, 0x65, 0x61, 0x67, 0x65, 0x12, 0x38, 0x0a, 0x05,
	0x62, 0x72, 0x61, 0x6e, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x67, 0x65,
	0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x42, 0x79, 0x49,
	0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x52,
	0x05, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x74, 0x1a, 0x96, 0x01, 0x0a, 0x05, 0x42, 0x72, 0x61, 0x6e, 0x64, 0x12, 0x19,
	0x0a, 0x08, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x62, 0x72, 0x61, 0x6e, 0x64, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a,
	0x0b, 0x64, 0x69, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x64, 0x69, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d,
	0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x22, 0x59, 0x0a,
	0x11, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x6f, 0x66, 0x66, 0x73, 0x65, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x6c, 0x69,
	0x6d, 0x69, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6c, 0x69, 0x6d, 0x69, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x65, 0x61, 0x72, 0x63, 0x68, 0x22, 0x37, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x43,
	0x61, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21,
	0x0a, 0x04, 0x43, 0x61, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x61, 0x72, 0x52, 0x04, 0x43, 0x61, 0x72,
	0x73, 0x32, 0xd0, 0x02, 0x0a, 0x0a, 0x43, 0x61, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x38, 0x0a, 0x09, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x12, 0x1a, 0x2e,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43,
	0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x61, 0x72, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x09, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x12, 0x1a, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43,
	0x61, 0x72, 0x22, 0x00, 0x12, 0x38, 0x0a, 0x09, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x43, 0x61,
	0x72, 0x12, 0x1a, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x43, 0x61, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0d, 0x2e,
	0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x43, 0x61, 0x72, 0x22, 0x00, 0x12, 0x49,
	0x0a, 0x0a, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x42, 0x79, 0x49, 0x44, 0x12, 0x1b, 0x2e, 0x67,
	0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x42, 0x79,
	0x49, 0x44, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x42, 0x79, 0x49, 0x44, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x0a, 0x47, 0x65, 0x74,
	0x43, 0x61, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1b, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x1c, 0x2e, 0x67, 0x65, 0x6e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e,
	0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x0f, 0x5a, 0x0d, 0x2e, 0x2f, 0x63, 0x61, 0x72, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_car_proto_rawDescOnce sync.Once
	file_car_proto_rawDescData = file_car_proto_rawDesc
)

func file_car_proto_rawDescGZIP() []byte {
	file_car_proto_rawDescOnce.Do(func() {
		file_car_proto_rawDescData = protoimpl.X.CompressGZIP(file_car_proto_rawDescData)
	})
	return file_car_proto_rawDescData
}

var file_car_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_car_proto_goTypes = []interface{}{
	(*CreateCarRequest)(nil),         // 0: genproto.CreateCarRequest
	(*Car)(nil),                      // 1: genproto.Car
	(*UpdateCarRequest)(nil),         // 2: genproto.UpdateCarRequest
	(*DeleteCarRequest)(nil),         // 3: genproto.DeleteCarRequest
	(*GetCarByIDRequest)(nil),        // 4: genproto.GetCarByIDRequest
	(*GetCarByIDResponse)(nil),       // 5: genproto.GetCarByIDResponse
	(*GetCarListRequest)(nil),        // 6: genproto.GetCarListRequest
	(*GetCarListResponse)(nil),       // 7: genproto.GetCarListResponse
	(*GetCarByIDResponse_Brand)(nil), // 8: genproto.GetCarByIDResponse.Brand
}
var file_car_proto_depIdxs = []int32{
	8, // 0: genproto.GetCarByIDResponse.brand:type_name -> genproto.GetCarByIDResponse.Brand
	1, // 1: genproto.GetCarListResponse.Cars:type_name -> genproto.Car
	0, // 2: genproto.CarService.CreateCar:input_type -> genproto.CreateCarRequest
	2, // 3: genproto.CarService.UpdateCar:input_type -> genproto.UpdateCarRequest
	3, // 4: genproto.CarService.DeleteCar:input_type -> genproto.DeleteCarRequest
	4, // 5: genproto.CarService.GetCarByID:input_type -> genproto.GetCarByIDRequest
	6, // 6: genproto.CarService.GetCarList:input_type -> genproto.GetCarListRequest
	1, // 7: genproto.CarService.CreateCar:output_type -> genproto.Car
	1, // 8: genproto.CarService.UpdateCar:output_type -> genproto.Car
	1, // 9: genproto.CarService.DeleteCar:output_type -> genproto.Car
	5, // 10: genproto.CarService.GetCarByID:output_type -> genproto.GetCarByIDResponse
	7, // 11: genproto.CarService.GetCarList:output_type -> genproto.GetCarListResponse
	7, // [7:12] is the sub-list for method output_type
	2, // [2:7] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_car_proto_init() }
func file_car_proto_init() {
	if File_car_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_car_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCarRequest); i {
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
		file_car_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Car); i {
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
		file_car_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateCarRequest); i {
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
		file_car_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCarRequest); i {
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
		file_car_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCarByIDRequest); i {
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
		file_car_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCarByIDResponse); i {
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
		file_car_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCarListRequest); i {
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
		file_car_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCarListResponse); i {
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
		file_car_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCarByIDResponse_Brand); i {
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
			RawDescriptor: file_car_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_car_proto_goTypes,
		DependencyIndexes: file_car_proto_depIdxs,
		MessageInfos:      file_car_proto_msgTypes,
	}.Build()
	File_car_proto = out.File
	file_car_proto_rawDesc = nil
	file_car_proto_goTypes = nil
	file_car_proto_depIdxs = nil
}