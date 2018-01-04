// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cli_to_hub.proto

/*
Package idl is a generated protocol buffer package.

It is generated from these files:
	cli_to_hub.proto
	hub_to_agent.proto

It has these top-level messages:
	PingRequest
	PingReply
	StatusUpgradeRequest
	StatusUpgradeReply
	UpgradeStepStatus
	CheckConfigRequest
	CheckConfigReply
	CountPerDb
	CheckObjectCountRequest
	CheckObjectCountReply
	CheckVersionRequest
	CheckVersionReply
	CheckDiskUsageRequest
	CheckDiskUsageReply
	PrepareShutdownClustersRequest
	PrepareShutdownClustersReply
	PrepareInitClusterRequest
	PrepareInitClusterReply
	UpgradeConvertMasterRequest
	UpgradeConvertMasterReply
	CheckUpgradeStatusRequest
	CheckUpgradeStatusReply
	FileSysUsage
	CheckDiskUsageRequestToAgent
	CheckDiskUsageReplyFromAgent
*/
package idl

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type UpgradeSteps int32

const (
	UpgradeSteps_UNKNOWN_STEP         UpgradeSteps = 0
	UpgradeSteps_CHECK_CONFIG         UpgradeSteps = 1
	UpgradeSteps_SEGINSTALL           UpgradeSteps = 2
	UpgradeSteps_PREPARE_INIT_CLUSTER UpgradeSteps = 3
	UpgradeSteps_MASTERUPGRADE        UpgradeSteps = 4
)

var UpgradeSteps_name = map[int32]string{
	0: "UNKNOWN_STEP",
	1: "CHECK_CONFIG",
	2: "SEGINSTALL",
	3: "PREPARE_INIT_CLUSTER",
	4: "MASTERUPGRADE",
}
var UpgradeSteps_value = map[string]int32{
	"UNKNOWN_STEP":         0,
	"CHECK_CONFIG":         1,
	"SEGINSTALL":           2,
	"PREPARE_INIT_CLUSTER": 3,
	"MASTERUPGRADE":        4,
}

func (x UpgradeSteps) String() string {
	return proto.EnumName(UpgradeSteps_name, int32(x))
}
func (UpgradeSteps) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type StepStatus int32

const (
	StepStatus_UNKNOWN_STATUS StepStatus = 0
	StepStatus_PENDING        StepStatus = 1
	StepStatus_RUNNING        StepStatus = 2
	StepStatus_COMPLETE       StepStatus = 3
	StepStatus_FAILED         StepStatus = 4
)

var StepStatus_name = map[int32]string{
	0: "UNKNOWN_STATUS",
	1: "PENDING",
	2: "RUNNING",
	3: "COMPLETE",
	4: "FAILED",
}
var StepStatus_value = map[string]int32{
	"UNKNOWN_STATUS": 0,
	"PENDING":        1,
	"RUNNING":        2,
	"COMPLETE":       3,
	"FAILED":         4,
}

func (x StepStatus) String() string {
	return proto.EnumName(StepStatus_name, int32(x))
}
func (StepStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type PingRequest struct {
}

func (m *PingRequest) Reset()                    { *m = PingRequest{} }
func (m *PingRequest) String() string            { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()               {}
func (*PingRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type PingReply struct {
}

func (m *PingReply) Reset()                    { *m = PingReply{} }
func (m *PingReply) String() string            { return proto.CompactTextString(m) }
func (*PingReply) ProtoMessage()               {}
func (*PingReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type StatusUpgradeRequest struct {
}

func (m *StatusUpgradeRequest) Reset()                    { *m = StatusUpgradeRequest{} }
func (m *StatusUpgradeRequest) String() string            { return proto.CompactTextString(m) }
func (*StatusUpgradeRequest) ProtoMessage()               {}
func (*StatusUpgradeRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type StatusUpgradeReply struct {
	ListOfUpgradeStepStatuses []*UpgradeStepStatus `protobuf:"bytes,1,rep,name=listOfUpgradeStepStatuses" json:"listOfUpgradeStepStatuses,omitempty"`
}

func (m *StatusUpgradeReply) Reset()                    { *m = StatusUpgradeReply{} }
func (m *StatusUpgradeReply) String() string            { return proto.CompactTextString(m) }
func (*StatusUpgradeReply) ProtoMessage()               {}
func (*StatusUpgradeReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *StatusUpgradeReply) GetListOfUpgradeStepStatuses() []*UpgradeStepStatus {
	if m != nil {
		return m.ListOfUpgradeStepStatuses
	}
	return nil
}

type UpgradeStepStatus struct {
	Step   UpgradeSteps `protobuf:"varint,1,opt,name=step,enum=idl.UpgradeSteps" json:"step,omitempty"`
	Status StepStatus   `protobuf:"varint,2,opt,name=status,enum=idl.StepStatus" json:"status,omitempty"`
}

func (m *UpgradeStepStatus) Reset()                    { *m = UpgradeStepStatus{} }
func (m *UpgradeStepStatus) String() string            { return proto.CompactTextString(m) }
func (*UpgradeStepStatus) ProtoMessage()               {}
func (*UpgradeStepStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *UpgradeStepStatus) GetStep() UpgradeSteps {
	if m != nil {
		return m.Step
	}
	return UpgradeSteps_UNKNOWN_STEP
}

func (m *UpgradeStepStatus) GetStatus() StepStatus {
	if m != nil {
		return m.Status
	}
	return StepStatus_UNKNOWN_STATUS
}

type CheckConfigRequest struct {
	DbPort int32 `protobuf:"varint,1,opt,name=dbPort" json:"dbPort,omitempty"`
}

func (m *CheckConfigRequest) Reset()                    { *m = CheckConfigRequest{} }
func (m *CheckConfigRequest) String() string            { return proto.CompactTextString(m) }
func (*CheckConfigRequest) ProtoMessage()               {}
func (*CheckConfigRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *CheckConfigRequest) GetDbPort() int32 {
	if m != nil {
		return m.DbPort
	}
	return 0
}

// Consider removing the status as errors are/should be put on the error field.
type CheckConfigReply struct {
	ConfigStatus string `protobuf:"bytes,1,opt,name=configStatus" json:"configStatus,omitempty"`
}

func (m *CheckConfigReply) Reset()                    { *m = CheckConfigReply{} }
func (m *CheckConfigReply) String() string            { return proto.CompactTextString(m) }
func (*CheckConfigReply) ProtoMessage()               {}
func (*CheckConfigReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *CheckConfigReply) GetConfigStatus() string {
	if m != nil {
		return m.ConfigStatus
	}
	return ""
}

type CountPerDb struct {
	DbName    string `protobuf:"bytes,1,opt,name=DbName" json:"DbName,omitempty"`
	AoCount   int32  `protobuf:"varint,2,opt,name=AoCount" json:"AoCount,omitempty"`
	HeapCount int32  `protobuf:"varint,3,opt,name=HeapCount" json:"HeapCount,omitempty"`
}

func (m *CountPerDb) Reset()                    { *m = CountPerDb{} }
func (m *CountPerDb) String() string            { return proto.CompactTextString(m) }
func (*CountPerDb) ProtoMessage()               {}
func (*CountPerDb) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *CountPerDb) GetDbName() string {
	if m != nil {
		return m.DbName
	}
	return ""
}

func (m *CountPerDb) GetAoCount() int32 {
	if m != nil {
		return m.AoCount
	}
	return 0
}

func (m *CountPerDb) GetHeapCount() int32 {
	if m != nil {
		return m.HeapCount
	}
	return 0
}

type CheckObjectCountRequest struct {
	DbPort int32 `protobuf:"varint,1,opt,name=DbPort" json:"DbPort,omitempty"`
}

func (m *CheckObjectCountRequest) Reset()                    { *m = CheckObjectCountRequest{} }
func (m *CheckObjectCountRequest) String() string            { return proto.CompactTextString(m) }
func (*CheckObjectCountRequest) ProtoMessage()               {}
func (*CheckObjectCountRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *CheckObjectCountRequest) GetDbPort() int32 {
	if m != nil {
		return m.DbPort
	}
	return 0
}

type CheckObjectCountReply struct {
	ListOfCounts []*CountPerDb `protobuf:"bytes,1,rep,name=list_of_counts,json=listOfCounts" json:"list_of_counts,omitempty"`
}

func (m *CheckObjectCountReply) Reset()                    { *m = CheckObjectCountReply{} }
func (m *CheckObjectCountReply) String() string            { return proto.CompactTextString(m) }
func (*CheckObjectCountReply) ProtoMessage()               {}
func (*CheckObjectCountReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *CheckObjectCountReply) GetListOfCounts() []*CountPerDb {
	if m != nil {
		return m.ListOfCounts
	}
	return nil
}

type CheckVersionRequest struct {
	DbPort int32  `protobuf:"varint,1,opt,name=DbPort" json:"DbPort,omitempty"`
	Host   string `protobuf:"bytes,2,opt,name=Host" json:"Host,omitempty"`
}

func (m *CheckVersionRequest) Reset()                    { *m = CheckVersionRequest{} }
func (m *CheckVersionRequest) String() string            { return proto.CompactTextString(m) }
func (*CheckVersionRequest) ProtoMessage()               {}
func (*CheckVersionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *CheckVersionRequest) GetDbPort() int32 {
	if m != nil {
		return m.DbPort
	}
	return 0
}

func (m *CheckVersionRequest) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

type CheckVersionReply struct {
	IsVersionCompatible bool `protobuf:"varint,1,opt,name=IsVersionCompatible" json:"IsVersionCompatible,omitempty"`
}

func (m *CheckVersionReply) Reset()                    { *m = CheckVersionReply{} }
func (m *CheckVersionReply) String() string            { return proto.CompactTextString(m) }
func (*CheckVersionReply) ProtoMessage()               {}
func (*CheckVersionReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *CheckVersionReply) GetIsVersionCompatible() bool {
	if m != nil {
		return m.IsVersionCompatible
	}
	return false
}

type CheckDiskUsageRequest struct {
}

func (m *CheckDiskUsageRequest) Reset()                    { *m = CheckDiskUsageRequest{} }
func (m *CheckDiskUsageRequest) String() string            { return proto.CompactTextString(m) }
func (*CheckDiskUsageRequest) ProtoMessage()               {}
func (*CheckDiskUsageRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

type CheckDiskUsageReply struct {
	SegmentFileSysUsage []string `protobuf:"bytes,1,rep,name=SegmentFileSysUsage" json:"SegmentFileSysUsage,omitempty"`
}

func (m *CheckDiskUsageReply) Reset()                    { *m = CheckDiskUsageReply{} }
func (m *CheckDiskUsageReply) String() string            { return proto.CompactTextString(m) }
func (*CheckDiskUsageReply) ProtoMessage()               {}
func (*CheckDiskUsageReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *CheckDiskUsageReply) GetSegmentFileSysUsage() []string {
	if m != nil {
		return m.SegmentFileSysUsage
	}
	return nil
}

type PrepareShutdownClustersRequest struct {
	OldBinDir string `protobuf:"bytes,1,opt,name=oldBinDir" json:"oldBinDir,omitempty"`
	NewBinDir string `protobuf:"bytes,2,opt,name=newBinDir" json:"newBinDir,omitempty"`
}

func (m *PrepareShutdownClustersRequest) Reset()                    { *m = PrepareShutdownClustersRequest{} }
func (m *PrepareShutdownClustersRequest) String() string            { return proto.CompactTextString(m) }
func (*PrepareShutdownClustersRequest) ProtoMessage()               {}
func (*PrepareShutdownClustersRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *PrepareShutdownClustersRequest) GetOldBinDir() string {
	if m != nil {
		return m.OldBinDir
	}
	return ""
}

func (m *PrepareShutdownClustersRequest) GetNewBinDir() string {
	if m != nil {
		return m.NewBinDir
	}
	return ""
}

type PrepareShutdownClustersReply struct {
}

func (m *PrepareShutdownClustersReply) Reset()                    { *m = PrepareShutdownClustersReply{} }
func (m *PrepareShutdownClustersReply) String() string            { return proto.CompactTextString(m) }
func (*PrepareShutdownClustersReply) ProtoMessage()               {}
func (*PrepareShutdownClustersReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

type PrepareInitClusterRequest struct {
	DbPort int32 `protobuf:"varint,1,opt,name=dbPort" json:"dbPort,omitempty"`
}

func (m *PrepareInitClusterRequest) Reset()                    { *m = PrepareInitClusterRequest{} }
func (m *PrepareInitClusterRequest) String() string            { return proto.CompactTextString(m) }
func (*PrepareInitClusterRequest) ProtoMessage()               {}
func (*PrepareInitClusterRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *PrepareInitClusterRequest) GetDbPort() int32 {
	if m != nil {
		return m.DbPort
	}
	return 0
}

type PrepareInitClusterReply struct {
}

func (m *PrepareInitClusterReply) Reset()                    { *m = PrepareInitClusterReply{} }
func (m *PrepareInitClusterReply) String() string            { return proto.CompactTextString(m) }
func (*PrepareInitClusterReply) ProtoMessage()               {}
func (*PrepareInitClusterReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

type UpgradeConvertMasterRequest struct {
	OldBinDir  string `protobuf:"bytes,1,opt,name=OldBinDir" json:"OldBinDir,omitempty"`
	OldDataDir string `protobuf:"bytes,2,opt,name=OldDataDir" json:"OldDataDir,omitempty"`
	NewBinDir  string `protobuf:"bytes,3,opt,name=NewBinDir" json:"NewBinDir,omitempty"`
	NewDataDir string `protobuf:"bytes,4,opt,name=NewDataDir" json:"NewDataDir,omitempty"`
}

func (m *UpgradeConvertMasterRequest) Reset()                    { *m = UpgradeConvertMasterRequest{} }
func (m *UpgradeConvertMasterRequest) String() string            { return proto.CompactTextString(m) }
func (*UpgradeConvertMasterRequest) ProtoMessage()               {}
func (*UpgradeConvertMasterRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{18} }

func (m *UpgradeConvertMasterRequest) GetOldBinDir() string {
	if m != nil {
		return m.OldBinDir
	}
	return ""
}

func (m *UpgradeConvertMasterRequest) GetOldDataDir() string {
	if m != nil {
		return m.OldDataDir
	}
	return ""
}

func (m *UpgradeConvertMasterRequest) GetNewBinDir() string {
	if m != nil {
		return m.NewBinDir
	}
	return ""
}

func (m *UpgradeConvertMasterRequest) GetNewDataDir() string {
	if m != nil {
		return m.NewDataDir
	}
	return ""
}

type UpgradeConvertMasterReply struct {
}

func (m *UpgradeConvertMasterReply) Reset()                    { *m = UpgradeConvertMasterReply{} }
func (m *UpgradeConvertMasterReply) String() string            { return proto.CompactTextString(m) }
func (*UpgradeConvertMasterReply) ProtoMessage()               {}
func (*UpgradeConvertMasterReply) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{19} }

func init() {
	proto.RegisterType((*PingRequest)(nil), "idl.PingRequest")
	proto.RegisterType((*PingReply)(nil), "idl.PingReply")
	proto.RegisterType((*StatusUpgradeRequest)(nil), "idl.StatusUpgradeRequest")
	proto.RegisterType((*StatusUpgradeReply)(nil), "idl.StatusUpgradeReply")
	proto.RegisterType((*UpgradeStepStatus)(nil), "idl.UpgradeStepStatus")
	proto.RegisterType((*CheckConfigRequest)(nil), "idl.CheckConfigRequest")
	proto.RegisterType((*CheckConfigReply)(nil), "idl.CheckConfigReply")
	proto.RegisterType((*CountPerDb)(nil), "idl.CountPerDb")
	proto.RegisterType((*CheckObjectCountRequest)(nil), "idl.CheckObjectCountRequest")
	proto.RegisterType((*CheckObjectCountReply)(nil), "idl.CheckObjectCountReply")
	proto.RegisterType((*CheckVersionRequest)(nil), "idl.CheckVersionRequest")
	proto.RegisterType((*CheckVersionReply)(nil), "idl.CheckVersionReply")
	proto.RegisterType((*CheckDiskUsageRequest)(nil), "idl.CheckDiskUsageRequest")
	proto.RegisterType((*CheckDiskUsageReply)(nil), "idl.CheckDiskUsageReply")
	proto.RegisterType((*PrepareShutdownClustersRequest)(nil), "idl.PrepareShutdownClustersRequest")
	proto.RegisterType((*PrepareShutdownClustersReply)(nil), "idl.PrepareShutdownClustersReply")
	proto.RegisterType((*PrepareInitClusterRequest)(nil), "idl.PrepareInitClusterRequest")
	proto.RegisterType((*PrepareInitClusterReply)(nil), "idl.PrepareInitClusterReply")
	proto.RegisterType((*UpgradeConvertMasterRequest)(nil), "idl.UpgradeConvertMasterRequest")
	proto.RegisterType((*UpgradeConvertMasterReply)(nil), "idl.UpgradeConvertMasterReply")
	proto.RegisterEnum("idl.UpgradeSteps", UpgradeSteps_name, UpgradeSteps_value)
	proto.RegisterEnum("idl.StepStatus", StepStatus_name, StepStatus_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for CliToHub service

type CliToHubClient interface {
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error)
	StatusUpgrade(ctx context.Context, in *StatusUpgradeRequest, opts ...grpc.CallOption) (*StatusUpgradeReply, error)
	CheckConfig(ctx context.Context, in *CheckConfigRequest, opts ...grpc.CallOption) (*CheckConfigReply, error)
	CheckObjectCount(ctx context.Context, in *CheckObjectCountRequest, opts ...grpc.CallOption) (*CheckObjectCountReply, error)
	CheckVersion(ctx context.Context, in *CheckVersionRequest, opts ...grpc.CallOption) (*CheckVersionReply, error)
	CheckDiskUsage(ctx context.Context, in *CheckDiskUsageRequest, opts ...grpc.CallOption) (*CheckDiskUsageReply, error)
	PrepareInitCluster(ctx context.Context, in *PrepareInitClusterRequest, opts ...grpc.CallOption) (*PrepareInitClusterReply, error)
	PrepareShutdownClusters(ctx context.Context, in *PrepareShutdownClustersRequest, opts ...grpc.CallOption) (*PrepareShutdownClustersReply, error)
	UpgradeConvertMaster(ctx context.Context, in *UpgradeConvertMasterRequest, opts ...grpc.CallOption) (*UpgradeConvertMasterReply, error)
}

type cliToHubClient struct {
	cc *grpc.ClientConn
}

func NewCliToHubClient(cc *grpc.ClientConn) CliToHubClient {
	return &cliToHubClient{cc}
}

func (c *cliToHubClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingReply, error) {
	out := new(PingReply)
	err := grpc.Invoke(ctx, "/idl.CliToHub/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cliToHubClient) StatusUpgrade(ctx context.Context, in *StatusUpgradeRequest, opts ...grpc.CallOption) (*StatusUpgradeReply, error) {
	out := new(StatusUpgradeReply)
	err := grpc.Invoke(ctx, "/idl.CliToHub/StatusUpgrade", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cliToHubClient) CheckConfig(ctx context.Context, in *CheckConfigRequest, opts ...grpc.CallOption) (*CheckConfigReply, error) {
	out := new(CheckConfigReply)
	err := grpc.Invoke(ctx, "/idl.CliToHub/CheckConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cliToHubClient) CheckObjectCount(ctx context.Context, in *CheckObjectCountRequest, opts ...grpc.CallOption) (*CheckObjectCountReply, error) {
	out := new(CheckObjectCountReply)
	err := grpc.Invoke(ctx, "/idl.CliToHub/CheckObjectCount", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cliToHubClient) CheckVersion(ctx context.Context, in *CheckVersionRequest, opts ...grpc.CallOption) (*CheckVersionReply, error) {
	out := new(CheckVersionReply)
	err := grpc.Invoke(ctx, "/idl.CliToHub/CheckVersion", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cliToHubClient) CheckDiskUsage(ctx context.Context, in *CheckDiskUsageRequest, opts ...grpc.CallOption) (*CheckDiskUsageReply, error) {
	out := new(CheckDiskUsageReply)
	err := grpc.Invoke(ctx, "/idl.CliToHub/CheckDiskUsage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cliToHubClient) PrepareInitCluster(ctx context.Context, in *PrepareInitClusterRequest, opts ...grpc.CallOption) (*PrepareInitClusterReply, error) {
	out := new(PrepareInitClusterReply)
	err := grpc.Invoke(ctx, "/idl.CliToHub/PrepareInitCluster", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cliToHubClient) PrepareShutdownClusters(ctx context.Context, in *PrepareShutdownClustersRequest, opts ...grpc.CallOption) (*PrepareShutdownClustersReply, error) {
	out := new(PrepareShutdownClustersReply)
	err := grpc.Invoke(ctx, "/idl.CliToHub/PrepareShutdownClusters", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cliToHubClient) UpgradeConvertMaster(ctx context.Context, in *UpgradeConvertMasterRequest, opts ...grpc.CallOption) (*UpgradeConvertMasterReply, error) {
	out := new(UpgradeConvertMasterReply)
	err := grpc.Invoke(ctx, "/idl.CliToHub/UpgradeConvertMaster", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CliToHub service

type CliToHubServer interface {
	Ping(context.Context, *PingRequest) (*PingReply, error)
	StatusUpgrade(context.Context, *StatusUpgradeRequest) (*StatusUpgradeReply, error)
	CheckConfig(context.Context, *CheckConfigRequest) (*CheckConfigReply, error)
	CheckObjectCount(context.Context, *CheckObjectCountRequest) (*CheckObjectCountReply, error)
	CheckVersion(context.Context, *CheckVersionRequest) (*CheckVersionReply, error)
	CheckDiskUsage(context.Context, *CheckDiskUsageRequest) (*CheckDiskUsageReply, error)
	PrepareInitCluster(context.Context, *PrepareInitClusterRequest) (*PrepareInitClusterReply, error)
	PrepareShutdownClusters(context.Context, *PrepareShutdownClustersRequest) (*PrepareShutdownClustersReply, error)
	UpgradeConvertMaster(context.Context, *UpgradeConvertMasterRequest) (*UpgradeConvertMasterReply, error)
}

func RegisterCliToHubServer(s *grpc.Server, srv CliToHubServer) {
	s.RegisterService(&_CliToHub_serviceDesc, srv)
}

func _CliToHub_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CliToHubServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/idl.CliToHub/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CliToHubServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CliToHub_StatusUpgrade_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StatusUpgradeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CliToHubServer).StatusUpgrade(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/idl.CliToHub/StatusUpgrade",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CliToHubServer).StatusUpgrade(ctx, req.(*StatusUpgradeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CliToHub_CheckConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CliToHubServer).CheckConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/idl.CliToHub/CheckConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CliToHubServer).CheckConfig(ctx, req.(*CheckConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CliToHub_CheckObjectCount_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckObjectCountRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CliToHubServer).CheckObjectCount(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/idl.CliToHub/CheckObjectCount",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CliToHubServer).CheckObjectCount(ctx, req.(*CheckObjectCountRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CliToHub_CheckVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CliToHubServer).CheckVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/idl.CliToHub/CheckVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CliToHubServer).CheckVersion(ctx, req.(*CheckVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CliToHub_CheckDiskUsage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckDiskUsageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CliToHubServer).CheckDiskUsage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/idl.CliToHub/CheckDiskUsage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CliToHubServer).CheckDiskUsage(ctx, req.(*CheckDiskUsageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CliToHub_PrepareInitCluster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrepareInitClusterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CliToHubServer).PrepareInitCluster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/idl.CliToHub/PrepareInitCluster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CliToHubServer).PrepareInitCluster(ctx, req.(*PrepareInitClusterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CliToHub_PrepareShutdownClusters_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrepareShutdownClustersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CliToHubServer).PrepareShutdownClusters(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/idl.CliToHub/PrepareShutdownClusters",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CliToHubServer).PrepareShutdownClusters(ctx, req.(*PrepareShutdownClustersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CliToHub_UpgradeConvertMaster_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpgradeConvertMasterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CliToHubServer).UpgradeConvertMaster(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/idl.CliToHub/UpgradeConvertMaster",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CliToHubServer).UpgradeConvertMaster(ctx, req.(*UpgradeConvertMasterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CliToHub_serviceDesc = grpc.ServiceDesc{
	ServiceName: "idl.CliToHub",
	HandlerType: (*CliToHubServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _CliToHub_Ping_Handler,
		},
		{
			MethodName: "StatusUpgrade",
			Handler:    _CliToHub_StatusUpgrade_Handler,
		},
		{
			MethodName: "CheckConfig",
			Handler:    _CliToHub_CheckConfig_Handler,
		},
		{
			MethodName: "CheckObjectCount",
			Handler:    _CliToHub_CheckObjectCount_Handler,
		},
		{
			MethodName: "CheckVersion",
			Handler:    _CliToHub_CheckVersion_Handler,
		},
		{
			MethodName: "CheckDiskUsage",
			Handler:    _CliToHub_CheckDiskUsage_Handler,
		},
		{
			MethodName: "PrepareInitCluster",
			Handler:    _CliToHub_PrepareInitCluster_Handler,
		},
		{
			MethodName: "PrepareShutdownClusters",
			Handler:    _CliToHub_PrepareShutdownClusters_Handler,
		},
		{
			MethodName: "UpgradeConvertMaster",
			Handler:    _CliToHub_UpgradeConvertMaster_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cli_to_hub.proto",
}

func init() { proto.RegisterFile("cli_to_hub.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 869 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x56, 0x7f, 0x6f, 0xe2, 0x46,
	0x10, 0x0d, 0x81, 0xfc, 0x60, 0x20, 0xd4, 0x99, 0xcb, 0x25, 0xe0, 0x43, 0x28, 0x75, 0x55, 0xf5,
	0x14, 0x55, 0x51, 0x9b, 0x53, 0xfb, 0x6f, 0xc5, 0xd9, 0x4e, 0x40, 0x47, 0x8c, 0x65, 0x9b, 0x56,
	0xaa, 0x4e, 0x42, 0x06, 0x36, 0x89, 0xef, 0x1c, 0xdb, 0xb5, 0x97, 0x46, 0x7c, 0x94, 0x7e, 0x8e,
	0x7e, 0xc1, 0xd3, 0xee, 0x1a, 0x6c, 0xc0, 0xdc, 0xfd, 0xc7, 0xbc, 0x37, 0xf3, 0x66, 0x67, 0xd7,
	0xf3, 0x12, 0x90, 0xa6, 0xbe, 0x37, 0xa6, 0xe1, 0xf8, 0x69, 0x3e, 0xb9, 0x8e, 0xe2, 0x90, 0x86,
	0x58, 0xf6, 0x66, 0xbe, 0x72, 0x02, 0x35, 0xd3, 0x0b, 0x1e, 0x2d, 0xf2, 0xcf, 0x9c, 0x24, 0x54,
	0xa9, 0x41, 0x55, 0x84, 0x91, 0xbf, 0x50, 0xce, 0xe1, 0xcc, 0xa6, 0x2e, 0x9d, 0x27, 0xa3, 0xe8,
	0x31, 0x76, 0x67, 0x64, 0x99, 0xf4, 0x09, 0x70, 0x03, 0x8f, 0xfc, 0x05, 0x3a, 0xd0, 0xf2, 0xbd,
	0x84, 0x0e, 0x1f, 0x52, 0xd4, 0xa6, 0x24, 0x12, 0x69, 0x24, 0x69, 0x96, 0x2e, 0xcb, 0x6f, 0x6b,
	0x37, 0xe7, 0xd7, 0xde, 0xcc, 0xbf, 0xde, 0xe2, 0xad, 0xdd, 0x85, 0xca, 0x14, 0x4e, 0xb7, 0x60,
	0xfc, 0x11, 0x2a, 0x09, 0x25, 0x51, 0xb3, 0x74, 0x59, 0x7a, 0xdb, 0xb8, 0x39, 0xdd, 0x54, 0x4d,
	0x2c, 0x4e, 0xe3, 0x4f, 0x70, 0x98, 0xf0, 0x82, 0xe6, 0x3e, 0x4f, 0xfc, 0x8e, 0x27, 0xe6, 0xfa,
	0xa6, 0xb4, 0xf2, 0x33, 0xa0, 0xfa, 0x44, 0xa6, 0x9f, 0xd5, 0x30, 0x78, 0xf0, 0x96, 0x77, 0x81,
	0xe7, 0x70, 0x38, 0x9b, 0x98, 0x61, 0x4c, 0x79, 0x9f, 0x03, 0x2b, 0x8d, 0x94, 0xdf, 0x41, 0x5a,
	0xcb, 0x66, 0xc3, 0x2b, 0x50, 0x9f, 0xf2, 0x50, 0x28, 0xf3, 0x8a, 0xaa, 0xb5, 0x86, 0x29, 0x1f,
	0x01, 0xd4, 0x70, 0x1e, 0x50, 0x93, 0xc4, 0xda, 0x84, 0xa9, 0x6b, 0x13, 0xc3, 0x7d, 0x26, 0x69,
	0x6e, 0x1a, 0x61, 0x13, 0x8e, 0xba, 0x21, 0xcf, 0xe3, 0xa7, 0x3e, 0xb0, 0x96, 0x21, 0xb6, 0xa1,
	0xda, 0x23, 0x6e, 0x24, 0xb8, 0x32, 0xe7, 0x32, 0x40, 0xf9, 0x15, 0x2e, 0xf8, 0xa9, 0x86, 0x93,
	0x4f, 0x64, 0x4a, 0x39, 0x96, 0x1b, 0x44, 0x5b, 0x1b, 0x44, 0x44, 0x8a, 0x01, 0xaf, 0xb7, 0x4b,
	0xd8, 0x34, 0xbf, 0x41, 0x83, 0xbd, 0xc8, 0x38, 0x7c, 0x18, 0x4f, 0x19, 0xba, 0x7c, 0x3f, 0x71,
	0x81, 0xd9, 0x10, 0x56, 0x5d, 0x3c, 0x1c, 0x47, 0x12, 0xa5, 0x0b, 0xaf, 0xb8, 0xde, 0x9f, 0x24,
	0x4e, 0xbc, 0x30, 0xf8, 0x46, 0x7b, 0x44, 0xa8, 0xf4, 0xc2, 0x44, 0x8c, 0x59, 0xb5, 0xf8, 0x6f,
	0x45, 0x87, 0xd3, 0x75, 0x09, 0x76, 0x9c, 0x5f, 0xe0, 0x55, 0x3f, 0x49, 0x11, 0x35, 0x7c, 0x8e,
	0x5c, 0xea, 0x4d, 0x7c, 0x71, 0x6f, 0xc7, 0x56, 0x11, 0xa5, 0x5c, 0xa4, 0x93, 0x69, 0x5e, 0xf2,
	0x79, 0x94, 0xb8, 0x8f, 0xab, 0x4f, 0xf7, 0x2e, 0x3d, 0x62, 0x8e, 0x48, 0x3b, 0xd8, 0xe4, 0xf1,
	0x99, 0x04, 0xf4, 0xd6, 0xf3, 0x89, 0xbd, 0x48, 0x38, 0xc7, 0xa7, 0xae, 0x5a, 0x45, 0x94, 0xf2,
	0x11, 0x3a, 0x66, 0x4c, 0x22, 0x37, 0x26, 0xf6, 0xd3, 0x9c, 0xce, 0xc2, 0x97, 0x40, 0xf5, 0xe7,
	0x09, 0x25, 0x71, 0xb2, 0x1c, 0xbb, 0x0d, 0xd5, 0xd0, 0x9f, 0xbd, 0xf7, 0x02, 0xcd, 0x8b, 0xd3,
	0x37, 0xce, 0x00, 0xc6, 0x06, 0xe4, 0x25, 0x65, 0xc5, 0x0d, 0x64, 0x80, 0xd2, 0x81, 0xf6, 0x4e,
	0x75, 0xb6, 0x99, 0xef, 0xa0, 0x95, 0xf2, 0xfd, 0xc0, 0xa3, 0x29, 0xf7, 0xad, 0xef, 0xb6, 0x05,
	0x17, 0x45, 0x45, 0x4c, 0xef, 0xbf, 0x12, 0xbc, 0x49, 0x17, 0x48, 0x0d, 0x83, 0x7f, 0x49, 0x4c,
	0xef, 0xdd, 0xbc, 0x64, 0x1b, 0xaa, 0xc3, 0xcd, 0x59, 0x56, 0x00, 0x76, 0x00, 0x86, 0xfe, 0x4c,
	0x73, 0xa9, 0x9b, 0x0d, 0x93, 0x43, 0x58, 0xb5, 0xb1, 0x9a, 0xb5, 0x2c, 0xaa, 0x57, 0x00, 0xab,
	0x36, 0xc8, 0xcb, 0xb2, 0xba, 0x22, 0xaa, 0x33, 0x44, 0x79, 0x03, 0xad, 0xe2, 0xa3, 0x45, 0xfe,
	0xe2, 0x2a, 0x84, 0x7a, 0x7e, 0xf1, 0x51, 0x82, 0xfa, 0xc8, 0xf8, 0x60, 0x0c, 0xff, 0x32, 0xc6,
	0xb6, 0xa3, 0x9b, 0xd2, 0x1e, 0x43, 0xd4, 0x9e, 0xae, 0x7e, 0x18, 0xab, 0x43, 0xe3, 0xb6, 0x7f,
	0x27, 0x95, 0xb0, 0x01, 0x60, 0xeb, 0x77, 0x7d, 0xc3, 0x76, 0xba, 0x83, 0x81, 0xb4, 0x8f, 0x4d,
	0x38, 0x33, 0x2d, 0xdd, 0xec, 0x5a, 0xfa, 0xb8, 0x6f, 0xf4, 0x9d, 0xb1, 0x3a, 0x18, 0xd9, 0x8e,
	0x6e, 0x49, 0x65, 0x3c, 0x85, 0x93, 0xfb, 0x2e, 0xfb, 0x3d, 0x32, 0xef, 0xac, 0xae, 0xa6, 0x4b,
	0x95, 0x2b, 0x07, 0x20, 0x67, 0x44, 0x08, 0x8d, 0xac, 0x5d, 0xd7, 0x19, 0xd9, 0xd2, 0x1e, 0xd6,
	0xe0, 0xc8, 0xd4, 0x0d, 0xad, 0x6f, 0xb0, 0x5e, 0x35, 0x38, 0xb2, 0x46, 0x86, 0xc1, 0x82, 0x7d,
	0xac, 0xc3, 0xb1, 0x3a, 0xbc, 0x37, 0x07, 0xba, 0xa3, 0x4b, 0x65, 0x04, 0x38, 0xbc, 0xed, 0xf6,
	0x07, 0xba, 0x26, 0x55, 0x6e, 0xfe, 0x3f, 0x80, 0x63, 0xd5, 0xf7, 0x9c, 0xb0, 0x37, 0x9f, 0xe0,
	0x15, 0x54, 0x98, 0x07, 0xa3, 0xc4, 0xb7, 0x2d, 0xe7, 0xce, 0x72, 0x23, 0x87, 0xb0, 0x67, 0xdb,
	0x43, 0x1d, 0x4e, 0xd6, 0xac, 0x18, 0x5b, 0xa9, 0xc7, 0x6d, 0xdb, 0xb6, 0x7c, 0x51, 0x44, 0x09,
	0x99, 0x3f, 0xa0, 0x96, 0xb3, 0x34, 0x14, 0x99, 0xdb, 0x96, 0x28, 0xbf, 0xde, 0x26, 0x84, 0x80,
	0x91, 0x7a, 0x62, 0xce, 0x4a, 0xb0, 0x9d, 0x25, 0x6f, 0x9b, 0x92, 0x2c, 0xef, 0x60, 0x85, 0xde,
	0x7b, 0xa8, 0xe7, 0x7d, 0x00, 0x9b, 0x59, 0xf6, 0xba, 0xbb, 0xc8, 0xe7, 0x05, 0x8c, 0xd0, 0xe8,
	0x41, 0x63, 0x7d, 0xd7, 0x31, 0xd7, 0x73, 0xd3, 0x19, 0xe4, 0x66, 0x21, 0x27, 0x94, 0x1c, 0xc0,
	0xed, 0xcd, 0xc1, 0x8e, 0x78, 0x8d, 0x5d, 0x7b, 0x28, 0xb7, 0x77, 0xf2, 0x42, 0x75, 0xba, 0xda,
	0xc7, 0xcd, 0x25, 0xc7, 0x1f, 0xf2, 0xa5, 0x3b, 0x0c, 0x46, 0xfe, 0xfe, 0xeb, 0x49, 0xa2, 0xc9,
	0xdf, 0x70, 0x56, 0xb4, 0x3d, 0x78, 0x99, 0xff, 0xa3, 0x59, 0xb4, 0xf3, 0x72, 0xe7, 0x2b, 0x19,
	0x5c, 0x7b, 0x72, 0xc8, 0xff, 0x8f, 0x78, 0xf7, 0x25, 0x00, 0x00, 0xff, 0xff, 0x3d, 0x73, 0x32,
	0x0b, 0x5b, 0x08, 0x00, 0x00,
}