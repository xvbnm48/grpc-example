// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.2
// source: student/student.proto

package student

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// DataStudentClient is the client API for DataStudent service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type DataStudentClient interface {
	FindStudentByEmail(ctx context.Context, in *Student, opts ...grpc.CallOption) (*Student, error)
}

type dataStudentClient struct {
	cc grpc.ClientConnInterface
}

func NewDataStudentClient(cc grpc.ClientConnInterface) DataStudentClient {
	return &dataStudentClient{cc}
}

func (c *dataStudentClient) FindStudentByEmail(ctx context.Context, in *Student, opts ...grpc.CallOption) (*Student, error) {
	out := new(Student)
	err := c.cc.Invoke(ctx, "/student.DataStudent/FindStudentByEmail", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DataStudentServer is the server API for DataStudent service.
// All implementations must embed UnimplementedDataStudentServer
// for forward compatibility
type DataStudentServer interface {
	FindStudentByEmail(context.Context, *Student) (*Student, error)
	mustEmbedUnimplementedDataStudentServer()
}

// UnimplementedDataStudentServer must be embedded to have forward compatible implementations.
type UnimplementedDataStudentServer struct {
}

func (UnimplementedDataStudentServer) FindStudentByEmail(context.Context, *Student) (*Student, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FindStudentByEmail not implemented")
}
func (UnimplementedDataStudentServer) mustEmbedUnimplementedDataStudentServer() {}

// UnsafeDataStudentServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DataStudentServer will
// result in compilation errors.
type UnsafeDataStudentServer interface {
	mustEmbedUnimplementedDataStudentServer()
}

func RegisterDataStudentServer(s grpc.ServiceRegistrar, srv DataStudentServer) {
	s.RegisterService(&DataStudent_ServiceDesc, srv)
}

func _DataStudent_FindStudentByEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Student)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DataStudentServer).FindStudentByEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/student.DataStudent/FindStudentByEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DataStudentServer).FindStudentByEmail(ctx, req.(*Student))
	}
	return interceptor(ctx, in, info, handler)
}

// DataStudent_ServiceDesc is the grpc.ServiceDesc for DataStudent service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var DataStudent_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "student.DataStudent",
	HandlerType: (*DataStudentServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "FindStudentByEmail",
			Handler:    _DataStudent_FindStudentByEmail_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "student/student.proto",
}
