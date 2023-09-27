// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/plugins/kubeflow/mpi.proto

package plugins

import (
	fmt "fmt"
	core "github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/core"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Proto for plugin that enables distributed training using https://github.com/kubeflow/mpi-operator
type DistributedMPITrainingTask struct {
	// Worker replicas spec
	WorkerReplicas *DistributedMPITrainingReplicaSpec `protobuf:"bytes,1,opt,name=worker_replicas,json=workerReplicas,proto3" json:"worker_replicas,omitempty"`
	// Master replicas spec
	LauncherReplicas *DistributedMPITrainingReplicaSpec `protobuf:"bytes,2,opt,name=launcher_replicas,json=launcherReplicas,proto3" json:"launcher_replicas,omitempty"`
	// RunPolicy encapsulates various runtime policies of the distributed training
	// job, for example how to clean up resources and how long the job can stay
	// active.
	RunPolicy *RunPolicy `protobuf:"bytes,3,opt,name=run_policy,json=runPolicy,proto3" json:"run_policy,omitempty"`
	// Number of slots per worker
	Slots                int32    `protobuf:"varint,4,opt,name=slots,proto3" json:"slots,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DistributedMPITrainingTask) Reset()         { *m = DistributedMPITrainingTask{} }
func (m *DistributedMPITrainingTask) String() string { return proto.CompactTextString(m) }
func (*DistributedMPITrainingTask) ProtoMessage()    {}
func (*DistributedMPITrainingTask) Descriptor() ([]byte, []int) {
	return fileDescriptor_298b02c608b0cddf, []int{0}
}

func (m *DistributedMPITrainingTask) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DistributedMPITrainingTask.Unmarshal(m, b)
}
func (m *DistributedMPITrainingTask) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DistributedMPITrainingTask.Marshal(b, m, deterministic)
}
func (m *DistributedMPITrainingTask) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DistributedMPITrainingTask.Merge(m, src)
}
func (m *DistributedMPITrainingTask) XXX_Size() int {
	return xxx_messageInfo_DistributedMPITrainingTask.Size(m)
}
func (m *DistributedMPITrainingTask) XXX_DiscardUnknown() {
	xxx_messageInfo_DistributedMPITrainingTask.DiscardUnknown(m)
}

var xxx_messageInfo_DistributedMPITrainingTask proto.InternalMessageInfo

func (m *DistributedMPITrainingTask) GetWorkerReplicas() *DistributedMPITrainingReplicaSpec {
	if m != nil {
		return m.WorkerReplicas
	}
	return nil
}

func (m *DistributedMPITrainingTask) GetLauncherReplicas() *DistributedMPITrainingReplicaSpec {
	if m != nil {
		return m.LauncherReplicas
	}
	return nil
}

func (m *DistributedMPITrainingTask) GetRunPolicy() *RunPolicy {
	if m != nil {
		return m.RunPolicy
	}
	return nil
}

func (m *DistributedMPITrainingTask) GetSlots() int32 {
	if m != nil {
		return m.Slots
	}
	return 0
}

// Replica specification for distributed MPI training
type DistributedMPITrainingReplicaSpec struct {
	// Number of replicas
	Replicas int32 `protobuf:"varint,1,opt,name=replicas,proto3" json:"replicas,omitempty"`
	// Image used for the replica group
	Image string `protobuf:"bytes,2,opt,name=image,proto3" json:"image,omitempty"`
	// Resources required for the replica group
	Resources *core.Resources `protobuf:"bytes,3,opt,name=resources,proto3" json:"resources,omitempty"`
	// Restart policy determines whether pods will be restarted when they exit
	RestartPolicy RestartPolicy `protobuf:"varint,4,opt,name=restart_policy,json=restartPolicy,proto3,enum=flyteidl.plugins.kubeflow.RestartPolicy" json:"restart_policy,omitempty"`
	// MPI sometimes requires different command set for different replica groups
	Command              []string `protobuf:"bytes,5,rep,name=command,proto3" json:"command,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DistributedMPITrainingReplicaSpec) Reset()         { *m = DistributedMPITrainingReplicaSpec{} }
func (m *DistributedMPITrainingReplicaSpec) String() string { return proto.CompactTextString(m) }
func (*DistributedMPITrainingReplicaSpec) ProtoMessage()    {}
func (*DistributedMPITrainingReplicaSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_298b02c608b0cddf, []int{1}
}

func (m *DistributedMPITrainingReplicaSpec) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DistributedMPITrainingReplicaSpec.Unmarshal(m, b)
}
func (m *DistributedMPITrainingReplicaSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DistributedMPITrainingReplicaSpec.Marshal(b, m, deterministic)
}
func (m *DistributedMPITrainingReplicaSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DistributedMPITrainingReplicaSpec.Merge(m, src)
}
func (m *DistributedMPITrainingReplicaSpec) XXX_Size() int {
	return xxx_messageInfo_DistributedMPITrainingReplicaSpec.Size(m)
}
func (m *DistributedMPITrainingReplicaSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_DistributedMPITrainingReplicaSpec.DiscardUnknown(m)
}

var xxx_messageInfo_DistributedMPITrainingReplicaSpec proto.InternalMessageInfo

func (m *DistributedMPITrainingReplicaSpec) GetReplicas() int32 {
	if m != nil {
		return m.Replicas
	}
	return 0
}

func (m *DistributedMPITrainingReplicaSpec) GetImage() string {
	if m != nil {
		return m.Image
	}
	return ""
}

func (m *DistributedMPITrainingReplicaSpec) GetResources() *core.Resources {
	if m != nil {
		return m.Resources
	}
	return nil
}

func (m *DistributedMPITrainingReplicaSpec) GetRestartPolicy() RestartPolicy {
	if m != nil {
		return m.RestartPolicy
	}
	return RestartPolicy_RESTART_POLICY_NEVER
}

func (m *DistributedMPITrainingReplicaSpec) GetCommand() []string {
	if m != nil {
		return m.Command
	}
	return nil
}

func init() {
	proto.RegisterType((*DistributedMPITrainingTask)(nil), "flyteidl.plugins.kubeflow.DistributedMPITrainingTask")
	proto.RegisterType((*DistributedMPITrainingReplicaSpec)(nil), "flyteidl.plugins.kubeflow.DistributedMPITrainingReplicaSpec")
}

func init() {
	proto.RegisterFile("flyteidl/plugins/kubeflow/mpi.proto", fileDescriptor_298b02c608b0cddf)
}

var fileDescriptor_298b02c608b0cddf = []byte{
	// 373 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x52, 0x4d, 0x6b, 0xdb, 0x30,
	0x00, 0xc5, 0xc9, 0xbc, 0xcd, 0x1a, 0xcb, 0x36, 0xb1, 0x83, 0xe7, 0x53, 0x96, 0x8d, 0xe1, 0xcb,
	0x2c, 0xc8, 0x60, 0xa5, 0xd0, 0x53, 0xdb, 0x4b, 0x0f, 0xa5, 0x41, 0xcd, 0xa9, 0x97, 0x20, 0x2b,
	0x8a, 0x23, 0x2c, 0x4b, 0x46, 0x1f, 0x84, 0xfc, 0xc4, 0xfe, 0xa3, 0x1e, 0x8b, 0x3f, 0x93, 0x96,
	0x26, 0xbd, 0xf4, 0xe6, 0x67, 0x3d, 0xbd, 0x0f, 0xfb, 0x81, 0x5f, 0x2b, 0xb1, 0xb5, 0x8c, 0x2f,
	0x05, 0x2a, 0x85, 0xcb, 0xb8, 0x34, 0x28, 0x77, 0x29, 0x5b, 0x09, 0xb5, 0x41, 0x45, 0xc9, 0x93,
	0x52, 0x2b, 0xab, 0xe0, 0x8f, 0x8e, 0x94, 0xb4, 0xa4, 0xa4, 0x23, 0x45, 0xfd, 0x11, 0xa2, 0x4a,
	0x33, 0x64, 0x89, 0xc9, 0x4d, 0x73, 0x2b, 0xfa, 0x73, 0x58, 0x9a, 0xaa, 0xa2, 0x50, 0xb2, 0xe1,
	0x4d, 0xee, 0x07, 0x20, 0xba, 0xe4, 0xc6, 0x6a, 0x9e, 0x3a, 0xcb, 0x96, 0xd7, 0xb3, 0xab, 0xb9,
	0x26, 0x5c, 0x72, 0x99, 0xcd, 0x89, 0xc9, 0x21, 0x03, 0x5f, 0x36, 0x4a, 0xe7, 0x4c, 0x2f, 0x34,
	0x2b, 0x05, 0xa7, 0xc4, 0x84, 0xde, 0xd8, 0x8b, 0x3f, 0x4d, 0xcf, 0x92, 0x83, 0xb1, 0x92, 0x97,
	0xf5, 0x70, 0x23, 0x70, 0x5b, 0x32, 0x8a, 0x47, 0x8d, 0x68, 0xfb, 0xca, 0x40, 0x0e, 0xbe, 0x09,
	0xe2, 0x24, 0x5d, 0xef, 0x1b, 0x0d, 0xde, 0xc0, 0xe8, 0x6b, 0x27, 0xdb, 0x5b, 0x5d, 0x00, 0xa0,
	0x9d, 0x5c, 0x94, 0x4a, 0x70, 0xba, 0x0d, 0x87, 0xb5, 0xc7, 0xef, 0x23, 0x1e, 0xd8, 0xc9, 0x59,
	0xcd, 0xc5, 0x81, 0xee, 0x1e, 0xe1, 0x77, 0xe0, 0x1b, 0xa1, 0xac, 0x09, 0xdf, 0x8d, 0xbd, 0xd8,
	0xc7, 0x0d, 0x98, 0x3c, 0x78, 0xe0, 0xe7, 0xab, 0x91, 0x60, 0x04, 0x3e, 0x3e, 0xf9, 0x96, 0x3e,
	0xee, 0x71, 0xa5, 0xcb, 0x0b, 0x92, 0xb1, 0xba, 0x7b, 0x80, 0x1b, 0x00, 0xff, 0x83, 0x40, 0x33,
	0xa3, 0x9c, 0xa6, 0xcc, 0xb4, 0x89, 0xc3, 0x5d, 0xe2, 0xea, 0xd7, 0x27, 0xb8, 0x3b, 0xc7, 0x3b,
	0x2a, 0xbc, 0x01, 0x23, 0xcd, 0x8c, 0x25, 0xda, 0x76, 0x75, 0xab, 0xb8, 0xa3, 0x69, 0x7c, 0xac,
	0x6e, 0x73, 0xa1, 0xad, 0xfc, 0x59, 0xef, 0x43, 0x18, 0x82, 0x0f, 0xd5, 0x78, 0x88, 0x5c, 0x86,
	0xfe, 0x78, 0x18, 0x07, 0xb8, 0x83, 0xe7, 0xa7, 0x77, 0x27, 0x19, 0xb7, 0x6b, 0x97, 0x26, 0x54,
	0x15, 0xa8, 0x96, 0x57, 0x3a, 0x43, 0xfd, 0x08, 0x33, 0x26, 0x51, 0x99, 0xfe, 0xcd, 0x14, 0x7a,
	0xbe, 0xcb, 0xf4, 0x7d, 0x3d, 0xc4, 0x7f, 0x8f, 0x01, 0x00, 0x00, 0xff, 0xff, 0x31, 0xf8, 0x1a,
	0x19, 0x0d, 0x03, 0x00, 0x00,
}