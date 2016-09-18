// Code generated by protoc-gen-go.
// source: downlinkQueue.proto
// DO NOT EDIT!

package api

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/google/api"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type EnqueueDownlinkQueueItemRequest struct {
	// hex encoded DevEUI
	DevEUI string `protobuf:"bytes,1,opt,name=devEUI" json:"devEUI,omitempty"`
	// random reference (used on ack notification)
	Reference string `protobuf:"bytes,2,opt,name=reference" json:"reference,omitempty"`
	// requires an ack from the node
	Confirmed bool `protobuf:"varint,3,opt,name=confirmed" json:"confirmed,omitempty"`
	// FPort to be used
	FPort uint32 `protobuf:"varint,4,opt,name=fPort" json:"fPort,omitempty"`
	// base64 encoded data
	Data []byte `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *EnqueueDownlinkQueueItemRequest) Reset()                    { *m = EnqueueDownlinkQueueItemRequest{} }
func (m *EnqueueDownlinkQueueItemRequest) String() string            { return proto.CompactTextString(m) }
func (*EnqueueDownlinkQueueItemRequest) ProtoMessage()               {}
func (*EnqueueDownlinkQueueItemRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{0} }

type EnqueueDownlinkQueueItemResponse struct {
}

func (m *EnqueueDownlinkQueueItemResponse) Reset()         { *m = EnqueueDownlinkQueueItemResponse{} }
func (m *EnqueueDownlinkQueueItemResponse) String() string { return proto.CompactTextString(m) }
func (*EnqueueDownlinkQueueItemResponse) ProtoMessage()    {}
func (*EnqueueDownlinkQueueItemResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor3, []int{1}
}

type DeleteDownlinkQeueueItemRequest struct {
	// ID of the queue item
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *DeleteDownlinkQeueueItemRequest) Reset()                    { *m = DeleteDownlinkQeueueItemRequest{} }
func (m *DeleteDownlinkQeueueItemRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteDownlinkQeueueItemRequest) ProtoMessage()               {}
func (*DeleteDownlinkQeueueItemRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{2} }

type DeleteDownlinkQueueItemResponse struct {
}

func (m *DeleteDownlinkQueueItemResponse) Reset()                    { *m = DeleteDownlinkQueueItemResponse{} }
func (m *DeleteDownlinkQueueItemResponse) String() string            { return proto.CompactTextString(m) }
func (*DeleteDownlinkQueueItemResponse) ProtoMessage()               {}
func (*DeleteDownlinkQueueItemResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{3} }

type DownlinkQueueItem struct {
	// id of the queue item
	Id int64 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	// hex encoded DevEUI
	DevEUI string `protobuf:"bytes,2,opt,name=devEUI" json:"devEUI,omitempty"`
	// random reference (used on ack notification)
	Reference string `protobuf:"bytes,3,opt,name=reference" json:"reference,omitempty"`
	// requires an ack from the node
	Confirmed bool `protobuf:"varint,4,opt,name=confirmed" json:"confirmed,omitempty"`
	// the transmission is pending (waiting for an ack)
	Pending bool `protobuf:"varint,5,opt,name=pending" json:"pending,omitempty"`
	// FPort to be used
	FPort uint32 `protobuf:"varint,6,opt,name=fPort" json:"fPort,omitempty"`
	// base64 encoded data
	Data []byte `protobuf:"bytes,7,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *DownlinkQueueItem) Reset()                    { *m = DownlinkQueueItem{} }
func (m *DownlinkQueueItem) String() string            { return proto.CompactTextString(m) }
func (*DownlinkQueueItem) ProtoMessage()               {}
func (*DownlinkQueueItem) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{4} }

type ListDownlinkQueueItemsRequest struct {
	// hex encoded DevEUI
	DevEUI string `protobuf:"bytes,1,opt,name=devEUI" json:"devEUI,omitempty"`
}

func (m *ListDownlinkQueueItemsRequest) Reset()                    { *m = ListDownlinkQueueItemsRequest{} }
func (m *ListDownlinkQueueItemsRequest) String() string            { return proto.CompactTextString(m) }
func (*ListDownlinkQueueItemsRequest) ProtoMessage()               {}
func (*ListDownlinkQueueItemsRequest) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{5} }

type ListDownlinkQueueItemsResponse struct {
	Items []*DownlinkQueueItem `protobuf:"bytes,1,rep,name=items" json:"items,omitempty"`
}

func (m *ListDownlinkQueueItemsResponse) Reset()                    { *m = ListDownlinkQueueItemsResponse{} }
func (m *ListDownlinkQueueItemsResponse) String() string            { return proto.CompactTextString(m) }
func (*ListDownlinkQueueItemsResponse) ProtoMessage()               {}
func (*ListDownlinkQueueItemsResponse) Descriptor() ([]byte, []int) { return fileDescriptor3, []int{6} }

func (m *ListDownlinkQueueItemsResponse) GetItems() []*DownlinkQueueItem {
	if m != nil {
		return m.Items
	}
	return nil
}

func init() {
	proto.RegisterType((*EnqueueDownlinkQueueItemRequest)(nil), "api.EnqueueDownlinkQueueItemRequest")
	proto.RegisterType((*EnqueueDownlinkQueueItemResponse)(nil), "api.EnqueueDownlinkQueueItemResponse")
	proto.RegisterType((*DeleteDownlinkQeueueItemRequest)(nil), "api.DeleteDownlinkQeueueItemRequest")
	proto.RegisterType((*DeleteDownlinkQueueItemResponse)(nil), "api.DeleteDownlinkQueueItemResponse")
	proto.RegisterType((*DownlinkQueueItem)(nil), "api.DownlinkQueueItem")
	proto.RegisterType((*ListDownlinkQueueItemsRequest)(nil), "api.ListDownlinkQueueItemsRequest")
	proto.RegisterType((*ListDownlinkQueueItemsResponse)(nil), "api.ListDownlinkQueueItemsResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for DownlinkQueue service

type DownlinkQueueClient interface {
	// Enqueue adds the given item to the queue.
	Enqueue(ctx context.Context, in *EnqueueDownlinkQueueItemRequest, opts ...grpc.CallOption) (*EnqueueDownlinkQueueItemResponse, error)
	// Delete deletes an item from the queue.
	Delete(ctx context.Context, in *DeleteDownlinkQeueueItemRequest, opts ...grpc.CallOption) (*DeleteDownlinkQueueItemResponse, error)
	// List lists the items in the queue for the given devEUI.
	List(ctx context.Context, in *ListDownlinkQueueItemsRequest, opts ...grpc.CallOption) (*ListDownlinkQueueItemsResponse, error)
}

type downlinkQueueClient struct {
	cc *grpc.ClientConn
}

func NewDownlinkQueueClient(cc *grpc.ClientConn) DownlinkQueueClient {
	return &downlinkQueueClient{cc}
}

func (c *downlinkQueueClient) Enqueue(ctx context.Context, in *EnqueueDownlinkQueueItemRequest, opts ...grpc.CallOption) (*EnqueueDownlinkQueueItemResponse, error) {
	out := new(EnqueueDownlinkQueueItemResponse)
	err := grpc.Invoke(ctx, "/api.DownlinkQueue/Enqueue", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *downlinkQueueClient) Delete(ctx context.Context, in *DeleteDownlinkQeueueItemRequest, opts ...grpc.CallOption) (*DeleteDownlinkQueueItemResponse, error) {
	out := new(DeleteDownlinkQueueItemResponse)
	err := grpc.Invoke(ctx, "/api.DownlinkQueue/Delete", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *downlinkQueueClient) List(ctx context.Context, in *ListDownlinkQueueItemsRequest, opts ...grpc.CallOption) (*ListDownlinkQueueItemsResponse, error) {
	out := new(ListDownlinkQueueItemsResponse)
	err := grpc.Invoke(ctx, "/api.DownlinkQueue/List", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DownlinkQueue service

type DownlinkQueueServer interface {
	// Enqueue adds the given item to the queue.
	Enqueue(context.Context, *EnqueueDownlinkQueueItemRequest) (*EnqueueDownlinkQueueItemResponse, error)
	// Delete deletes an item from the queue.
	Delete(context.Context, *DeleteDownlinkQeueueItemRequest) (*DeleteDownlinkQueueItemResponse, error)
	// List lists the items in the queue for the given devEUI.
	List(context.Context, *ListDownlinkQueueItemsRequest) (*ListDownlinkQueueItemsResponse, error)
}

func RegisterDownlinkQueueServer(s *grpc.Server, srv DownlinkQueueServer) {
	s.RegisterService(&_DownlinkQueue_serviceDesc, srv)
}

func _DownlinkQueue_Enqueue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnqueueDownlinkQueueItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DownlinkQueueServer).Enqueue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DownlinkQueue/Enqueue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DownlinkQueueServer).Enqueue(ctx, req.(*EnqueueDownlinkQueueItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DownlinkQueue_Delete_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteDownlinkQeueueItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DownlinkQueueServer).Delete(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DownlinkQueue/Delete",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DownlinkQueueServer).Delete(ctx, req.(*DeleteDownlinkQeueueItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _DownlinkQueue_List_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListDownlinkQueueItemsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DownlinkQueueServer).List(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.DownlinkQueue/List",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DownlinkQueueServer).List(ctx, req.(*ListDownlinkQueueItemsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DownlinkQueue_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.DownlinkQueue",
	HandlerType: (*DownlinkQueueServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Enqueue",
			Handler:    _DownlinkQueue_Enqueue_Handler,
		},
		{
			MethodName: "Delete",
			Handler:    _DownlinkQueue_Delete_Handler,
		},
		{
			MethodName: "List",
			Handler:    _DownlinkQueue_List_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: fileDescriptor3,
}

func init() { proto.RegisterFile("downlinkQueue.proto", fileDescriptor3) }

var fileDescriptor3 = []byte{
	// 430 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x53, 0x51, 0x8e, 0xd3, 0x30,
	0x10, 0x95, 0x93, 0x34, 0xdd, 0x1d, 0x58, 0x24, 0x0c, 0xda, 0x0d, 0x61, 0x4b, 0x83, 0x17, 0xa4,
	0x68, 0x85, 0x5a, 0x51, 0x3e, 0x90, 0xf8, 0x6e, 0x3f, 0x2a, 0x21, 0x04, 0x91, 0x38, 0x40, 0xa8,
	0xa7, 0x95, 0x45, 0x6b, 0xa7, 0x89, 0x03, 0x1f, 0xa8, 0x3f, 0x5c, 0x81, 0x03, 0x70, 0x10, 0x8e,
	0xc1, 0x15, 0xb8, 0x01, 0x17, 0x40, 0x71, 0x82, 0xd2, 0x36, 0x6d, 0xf3, 0x97, 0x99, 0x79, 0x33,
	0x2f, 0xf3, 0xde, 0x18, 0x1e, 0x70, 0xf5, 0x55, 0x2e, 0x85, 0xfc, 0xfc, 0x21, 0xc7, 0x1c, 0x07,
	0x49, 0xaa, 0xb4, 0xa2, 0x76, 0x9c, 0x08, 0xff, 0x7a, 0xa1, 0xd4, 0x62, 0x89, 0xc3, 0x38, 0x11,
	0xc3, 0x58, 0x4a, 0xa5, 0x63, 0x2d, 0x94, 0xcc, 0x4a, 0x08, 0xfb, 0x49, 0xa0, 0x3f, 0x91, 0xeb,
	0xa2, 0x69, 0xbc, 0x3d, 0x61, 0xaa, 0x71, 0x15, 0xe1, 0x3a, 0xc7, 0x4c, 0xd3, 0x4b, 0x70, 0x39,
	0x7e, 0x99, 0x7c, 0x9c, 0x7a, 0x24, 0x20, 0xe1, 0x79, 0x54, 0x45, 0xf4, 0x1a, 0xce, 0x53, 0x9c,
	0x63, 0x8a, 0x72, 0x86, 0x9e, 0x65, 0x4a, 0x75, 0xa2, 0xa8, 0xce, 0x94, 0x9c, 0x8b, 0x74, 0x85,
	0xdc, 0xb3, 0x03, 0x12, 0x9e, 0x45, 0x75, 0x82, 0x3e, 0x84, 0xce, 0xfc, 0xbd, 0x4a, 0xb5, 0xe7,
	0x04, 0x24, 0xbc, 0x88, 0xca, 0x80, 0x52, 0x70, 0x78, 0xac, 0x63, 0xaf, 0x13, 0x90, 0xf0, 0x6e,
	0x64, 0xbe, 0x19, 0x83, 0xe0, 0xf8, 0x0f, 0x66, 0x89, 0x92, 0x19, 0xb2, 0x97, 0xd0, 0x1f, 0xe3,
	0x12, 0x75, 0x0d, 0xc1, 0xfd, 0x25, 0xee, 0x81, 0x25, 0xb8, 0x59, 0xc0, 0x8e, 0x2c, 0xc1, 0xd9,
	0xd3, 0x46, 0x4b, 0x63, 0xea, 0x2f, 0x02, 0xf7, 0x1b, 0xd5, 0xfd, 0x41, 0x5b, 0xea, 0x58, 0xc7,
	0xd5, 0xb1, 0x4f, 0xaa, 0xe3, 0xec, 0xab, 0xe3, 0x41, 0x37, 0x41, 0xc9, 0x85, 0x5c, 0x18, 0x29,
	0xce, 0xa2, 0xff, 0x61, 0xad, 0x9b, 0x7b, 0x48, 0xb7, 0xee, 0x96, 0x6e, 0xaf, 0xa1, 0xf7, 0x56,
	0x64, 0xba, 0xb1, 0x40, 0xd6, 0x62, 0x2b, 0x7b, 0x07, 0x4f, 0x8e, 0x35, 0x96, 0xc2, 0xd0, 0x17,
	0xd0, 0x11, 0x45, 0xc2, 0x23, 0x81, 0x1d, 0xde, 0x19, 0x5d, 0x0e, 0xe2, 0x44, 0x0c, 0x9a, 0x3a,
	0x96, 0xa0, 0xd1, 0x5f, 0x0b, 0x2e, 0x76, 0x8a, 0x34, 0x87, 0x6e, 0x65, 0x29, 0x7d, 0x66, 0x7a,
	0x5b, 0x2e, 0xd0, 0x7f, 0xde, 0x82, 0xaa, 0x0c, 0xeb, 0x7d, 0xff, 0xfd, 0xe7, 0x87, 0x75, 0xc5,
	0xa8, 0x39, 0xf6, 0x9d, 0x17, 0xf1, 0x86, 0xdc, 0xd2, 0x1c, 0xdc, 0xd2, 0xf2, 0x8a, 0xb5, 0xe5,
	0x64, 0xfc, 0x83, 0xa8, 0x06, 0x69, 0xdf, 0x90, 0x3e, 0xba, 0xbd, 0x6a, 0x92, 0x0e, 0xbf, 0x09,
	0xbe, 0xa1, 0x1a, 0x9c, 0x42, 0x4f, 0xca, 0xcc, 0xb8, 0x93, 0x9e, 0xf8, 0x37, 0x27, 0x31, 0x15,
	0xe3, 0x8d, 0x61, 0xec, 0xd1, 0xc7, 0x87, 0x18, 0x4b, 0x13, 0x37, 0x9f, 0x5c, 0xf3, 0xbe, 0x5f,
	0xfd, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x18, 0xa7, 0x22, 0x52, 0x19, 0x04, 0x00, 0x00,
}