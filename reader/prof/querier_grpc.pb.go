// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v3.21.12
// source: querier.proto

package prof

import (
	context "context"
	v11 "github.com/metrico/qryn/reader/prof/google/v1"
	v1 "github.com/metrico/qryn/reader/prof/types/v1"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	QuerierService_ProfileTypes_FullMethodName           = "/querier.v1.QuerierService/ProfileTypes"
	QuerierService_LabelValues_FullMethodName            = "/querier.v1.QuerierService/LabelValues"
	QuerierService_LabelNames_FullMethodName             = "/querier.v1.QuerierService/LabelNames"
	QuerierService_Series_FullMethodName                 = "/querier.v1.QuerierService/Series"
	QuerierService_SelectMergeStacktraces_FullMethodName = "/querier.v1.QuerierService/SelectMergeStacktraces"
	QuerierService_SelectMergeSpanProfile_FullMethodName = "/querier.v1.QuerierService/SelectMergeSpanProfile"
	QuerierService_SelectMergeProfile_FullMethodName     = "/querier.v1.QuerierService/SelectMergeProfile"
	QuerierService_SelectSeries_FullMethodName           = "/querier.v1.QuerierService/SelectSeries"
	QuerierService_Diff_FullMethodName                   = "/querier.v1.QuerierService/Diff"
	QuerierService_GetProfileStats_FullMethodName        = "/querier.v1.QuerierService/GetProfileStats"
	QuerierService_AnalyzeQuery_FullMethodName           = "/querier.v1.QuerierService/AnalyzeQuery"
)

// QuerierServiceClient is the client API for QuerierService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QuerierServiceClient interface {
	// ProfileType returns a list of the existing profile types.
	ProfileTypes(ctx context.Context, in *ProfileTypesRequest, opts ...grpc.CallOption) (*ProfileTypesResponse, error)
	// LabelValues returns the existing label values for the provided label names.
	LabelValues(ctx context.Context, in *v1.LabelValuesRequest, opts ...grpc.CallOption) (*v1.LabelValuesResponse, error)
	// LabelNames returns a list of the existing label names.
	LabelNames(ctx context.Context, in *v1.LabelNamesRequest, opts ...grpc.CallOption) (*v1.LabelNamesResponse, error)
	// Series returns profiles series matching the request. A series is a unique label set.
	Series(ctx context.Context, in *SeriesRequest, opts ...grpc.CallOption) (*SeriesResponse, error)
	// SelectMergeStacktraces returns matching profiles aggregated in a flamegraph format. It will combine samples from within the same callstack, with each element being grouped by its function name.
	SelectMergeStacktraces(ctx context.Context, in *SelectMergeStacktracesRequest, opts ...grpc.CallOption) (*SelectMergeStacktracesResponse, error)
	// SelectMergeSpanProfile returns matching profiles aggregated in a flamegraph format. It will combine samples from within the same callstack, with each element being grouped by its function name.
	SelectMergeSpanProfile(ctx context.Context, in *SelectMergeSpanProfileRequest, opts ...grpc.CallOption) (*SelectMergeSpanProfileResponse, error)
	// SelectMergeProfile returns matching profiles aggregated in pprof format. It will contain all information stored (so including filenames and line number, if ingested).
	SelectMergeProfile(ctx context.Context, in *SelectMergeProfileRequest, opts ...grpc.CallOption) (*v11.Profile, error)
	// SelectSeries returns a time series for the total sum of the requested profiles.
	SelectSeries(ctx context.Context, in *SelectSeriesRequest, opts ...grpc.CallOption) (*SelectSeriesResponse, error)
	// Diff returns a diff of two profiles
	Diff(ctx context.Context, in *DiffRequest, opts ...grpc.CallOption) (*DiffResponse, error)
	// GetProfileStats returns profile stats for the current tenant.
	GetProfileStats(ctx context.Context, in *v1.GetProfileStatsRequest, opts ...grpc.CallOption) (*v1.GetProfileStatsResponse, error)
	AnalyzeQuery(ctx context.Context, in *AnalyzeQueryRequest, opts ...grpc.CallOption) (*AnalyzeQueryResponse, error)
}

type querierServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewQuerierServiceClient(cc grpc.ClientConnInterface) QuerierServiceClient {
	return &querierServiceClient{cc}
}

func (c *querierServiceClient) ProfileTypes(ctx context.Context, in *ProfileTypesRequest, opts ...grpc.CallOption) (*ProfileTypesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(ProfileTypesResponse)
	err := c.cc.Invoke(ctx, QuerierService_ProfileTypes_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *querierServiceClient) LabelValues(ctx context.Context, in *v1.LabelValuesRequest, opts ...grpc.CallOption) (*v1.LabelValuesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(v1.LabelValuesResponse)
	err := c.cc.Invoke(ctx, QuerierService_LabelValues_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *querierServiceClient) LabelNames(ctx context.Context, in *v1.LabelNamesRequest, opts ...grpc.CallOption) (*v1.LabelNamesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(v1.LabelNamesResponse)
	err := c.cc.Invoke(ctx, QuerierService_LabelNames_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *querierServiceClient) Series(ctx context.Context, in *SeriesRequest, opts ...grpc.CallOption) (*SeriesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SeriesResponse)
	err := c.cc.Invoke(ctx, QuerierService_Series_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *querierServiceClient) SelectMergeStacktraces(ctx context.Context, in *SelectMergeStacktracesRequest, opts ...grpc.CallOption) (*SelectMergeStacktracesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SelectMergeStacktracesResponse)
	err := c.cc.Invoke(ctx, QuerierService_SelectMergeStacktraces_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *querierServiceClient) SelectMergeSpanProfile(ctx context.Context, in *SelectMergeSpanProfileRequest, opts ...grpc.CallOption) (*SelectMergeSpanProfileResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SelectMergeSpanProfileResponse)
	err := c.cc.Invoke(ctx, QuerierService_SelectMergeSpanProfile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *querierServiceClient) SelectMergeProfile(ctx context.Context, in *SelectMergeProfileRequest, opts ...grpc.CallOption) (*v11.Profile, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(v11.Profile)
	err := c.cc.Invoke(ctx, QuerierService_SelectMergeProfile_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *querierServiceClient) SelectSeries(ctx context.Context, in *SelectSeriesRequest, opts ...grpc.CallOption) (*SelectSeriesResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(SelectSeriesResponse)
	err := c.cc.Invoke(ctx, QuerierService_SelectSeries_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *querierServiceClient) Diff(ctx context.Context, in *DiffRequest, opts ...grpc.CallOption) (*DiffResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DiffResponse)
	err := c.cc.Invoke(ctx, QuerierService_Diff_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *querierServiceClient) GetProfileStats(ctx context.Context, in *v1.GetProfileStatsRequest, opts ...grpc.CallOption) (*v1.GetProfileStatsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(v1.GetProfileStatsResponse)
	err := c.cc.Invoke(ctx, QuerierService_GetProfileStats_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *querierServiceClient) AnalyzeQuery(ctx context.Context, in *AnalyzeQueryRequest, opts ...grpc.CallOption) (*AnalyzeQueryResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(AnalyzeQueryResponse)
	err := c.cc.Invoke(ctx, QuerierService_AnalyzeQuery_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QuerierServiceServer is the server API for QuerierService service.
// All implementations must embed UnimplementedQuerierServiceServer
// for forward compatibility.
type QuerierServiceServer interface {
	// ProfileType returns a list of the existing profile types.
	ProfileTypes(context.Context, *ProfileTypesRequest) (*ProfileTypesResponse, error)
	// LabelValues returns the existing label values for the provided label names.
	LabelValues(context.Context, *v1.LabelValuesRequest) (*v1.LabelValuesResponse, error)
	// LabelNames returns a list of the existing label names.
	LabelNames(context.Context, *v1.LabelNamesRequest) (*v1.LabelNamesResponse, error)
	// Series returns profiles series matching the request. A series is a unique label set.
	Series(context.Context, *SeriesRequest) (*SeriesResponse, error)
	// SelectMergeStacktraces returns matching profiles aggregated in a flamegraph format. It will combine samples from within the same callstack, with each element being grouped by its function name.
	SelectMergeStacktraces(context.Context, *SelectMergeStacktracesRequest) (*SelectMergeStacktracesResponse, error)
	// SelectMergeSpanProfile returns matching profiles aggregated in a flamegraph format. It will combine samples from within the same callstack, with each element being grouped by its function name.
	SelectMergeSpanProfile(context.Context, *SelectMergeSpanProfileRequest) (*SelectMergeSpanProfileResponse, error)
	// SelectMergeProfile returns matching profiles aggregated in pprof format. It will contain all information stored (so including filenames and line number, if ingested).
	SelectMergeProfile(context.Context, *SelectMergeProfileRequest) (*v11.Profile, error)
	// SelectSeries returns a time series for the total sum of the requested profiles.
	SelectSeries(context.Context, *SelectSeriesRequest) (*SelectSeriesResponse, error)
	// Diff returns a diff of two profiles
	Diff(context.Context, *DiffRequest) (*DiffResponse, error)
	// GetProfileStats returns profile stats for the current tenant.
	GetProfileStats(context.Context, *v1.GetProfileStatsRequest) (*v1.GetProfileStatsResponse, error)
	AnalyzeQuery(context.Context, *AnalyzeQueryRequest) (*AnalyzeQueryResponse, error)
	mustEmbedUnimplementedQuerierServiceServer()
}

// UnimplementedQuerierServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedQuerierServiceServer struct{}

func (UnimplementedQuerierServiceServer) ProfileTypes(context.Context, *ProfileTypesRequest) (*ProfileTypesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ProfileTypes not implemented")
}
func (UnimplementedQuerierServiceServer) LabelValues(context.Context, *v1.LabelValuesRequest) (*v1.LabelValuesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LabelValues not implemented")
}
func (UnimplementedQuerierServiceServer) LabelNames(context.Context, *v1.LabelNamesRequest) (*v1.LabelNamesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method LabelNames not implemented")
}
func (UnimplementedQuerierServiceServer) Series(context.Context, *SeriesRequest) (*SeriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Series not implemented")
}
func (UnimplementedQuerierServiceServer) SelectMergeStacktraces(context.Context, *SelectMergeStacktracesRequest) (*SelectMergeStacktracesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SelectMergeStacktraces not implemented")
}
func (UnimplementedQuerierServiceServer) SelectMergeSpanProfile(context.Context, *SelectMergeSpanProfileRequest) (*SelectMergeSpanProfileResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SelectMergeSpanProfile not implemented")
}
func (UnimplementedQuerierServiceServer) SelectMergeProfile(context.Context, *SelectMergeProfileRequest) (*v11.Profile, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SelectMergeProfile not implemented")
}
func (UnimplementedQuerierServiceServer) SelectSeries(context.Context, *SelectSeriesRequest) (*SelectSeriesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SelectSeries not implemented")
}
func (UnimplementedQuerierServiceServer) Diff(context.Context, *DiffRequest) (*DiffResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Diff not implemented")
}
func (UnimplementedQuerierServiceServer) GetProfileStats(context.Context, *v1.GetProfileStatsRequest) (*v1.GetProfileStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProfileStats not implemented")
}
func (UnimplementedQuerierServiceServer) AnalyzeQuery(context.Context, *AnalyzeQueryRequest) (*AnalyzeQueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AnalyzeQuery not implemented")
}
func (UnimplementedQuerierServiceServer) mustEmbedUnimplementedQuerierServiceServer() {}
func (UnimplementedQuerierServiceServer) testEmbeddedByValue()                        {}

// UnsafeQuerierServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QuerierServiceServer will
// result in compilation errors.
type UnsafeQuerierServiceServer interface {
	mustEmbedUnimplementedQuerierServiceServer()
}

func RegisterQuerierServiceServer(s grpc.ServiceRegistrar, srv QuerierServiceServer) {
	// If the following call pancis, it indicates UnimplementedQuerierServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&QuerierService_ServiceDesc, srv)
}

func _QuerierService_ProfileTypes_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProfileTypesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuerierServiceServer).ProfileTypes(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuerierService_ProfileTypes_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuerierServiceServer).ProfileTypes(ctx, req.(*ProfileTypesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuerierService_LabelValues_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.LabelValuesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuerierServiceServer).LabelValues(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuerierService_LabelValues_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuerierServiceServer).LabelValues(ctx, req.(*v1.LabelValuesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuerierService_LabelNames_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.LabelNamesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuerierServiceServer).LabelNames(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuerierService_LabelNames_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuerierServiceServer).LabelNames(ctx, req.(*v1.LabelNamesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuerierService_Series_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SeriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuerierServiceServer).Series(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuerierService_Series_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuerierServiceServer).Series(ctx, req.(*SeriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuerierService_SelectMergeStacktraces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SelectMergeStacktracesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuerierServiceServer).SelectMergeStacktraces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuerierService_SelectMergeStacktraces_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuerierServiceServer).SelectMergeStacktraces(ctx, req.(*SelectMergeStacktracesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuerierService_SelectMergeSpanProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SelectMergeSpanProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuerierServiceServer).SelectMergeSpanProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuerierService_SelectMergeSpanProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuerierServiceServer).SelectMergeSpanProfile(ctx, req.(*SelectMergeSpanProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuerierService_SelectMergeProfile_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SelectMergeProfileRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuerierServiceServer).SelectMergeProfile(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuerierService_SelectMergeProfile_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuerierServiceServer).SelectMergeProfile(ctx, req.(*SelectMergeProfileRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuerierService_SelectSeries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SelectSeriesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuerierServiceServer).SelectSeries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuerierService_SelectSeries_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuerierServiceServer).SelectSeries(ctx, req.(*SelectSeriesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuerierService_Diff_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DiffRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuerierServiceServer).Diff(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuerierService_Diff_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuerierServiceServer).Diff(ctx, req.(*DiffRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuerierService_GetProfileStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(v1.GetProfileStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuerierServiceServer).GetProfileStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuerierService_GetProfileStats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuerierServiceServer).GetProfileStats(ctx, req.(*v1.GetProfileStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QuerierService_AnalyzeQuery_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AnalyzeQueryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QuerierServiceServer).AnalyzeQuery(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: QuerierService_AnalyzeQuery_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QuerierServiceServer).AnalyzeQuery(ctx, req.(*AnalyzeQueryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// QuerierService_ServiceDesc is the grpc.ServiceDesc for QuerierService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var QuerierService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "querier.v1.QuerierService",
	HandlerType: (*QuerierServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ProfileTypes",
			Handler:    _QuerierService_ProfileTypes_Handler,
		},
		{
			MethodName: "LabelValues",
			Handler:    _QuerierService_LabelValues_Handler,
		},
		{
			MethodName: "LabelNames",
			Handler:    _QuerierService_LabelNames_Handler,
		},
		{
			MethodName: "Series",
			Handler:    _QuerierService_Series_Handler,
		},
		{
			MethodName: "SelectMergeStacktraces",
			Handler:    _QuerierService_SelectMergeStacktraces_Handler,
		},
		{
			MethodName: "SelectMergeSpanProfile",
			Handler:    _QuerierService_SelectMergeSpanProfile_Handler,
		},
		{
			MethodName: "SelectMergeProfile",
			Handler:    _QuerierService_SelectMergeProfile_Handler,
		},
		{
			MethodName: "SelectSeries",
			Handler:    _QuerierService_SelectSeries_Handler,
		},
		{
			MethodName: "Diff",
			Handler:    _QuerierService_Diff_Handler,
		},
		{
			MethodName: "GetProfileStats",
			Handler:    _QuerierService_GetProfileStats_Handler,
		},
		{
			MethodName: "AnalyzeQuery",
			Handler:    _QuerierService_AnalyzeQuery_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "querier.proto",
}
