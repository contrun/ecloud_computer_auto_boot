// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ReallocateRequestInstances struct {

    // 订购资源实例ID
	InstanceId *string `json:"instanceId,omitempty"`
    // 云电脑ID
	MachineId *string `json:"machineId,omitempty"`
}

func (s ReallocateRequestInstances) String() string {
	return utils.Beautify(s)
}

func (s ReallocateRequestInstances) GoString() string {
	return s.String()
}

func (s ReallocateRequestInstances) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ReallocateRequestInstances) SetInstanceId(v string) *ReallocateRequestInstances {
	s.InstanceId = &v
	return s
}

func (s *ReallocateRequestInstances) SetMachineId(v string) *ReallocateRequestInstances {
	s.MachineId = &v
	return s
}


type ReallocateRequestInstancesBuilder struct {
	s *ReallocateRequestInstances
}

func NewReallocateRequestInstancesBuilder() *ReallocateRequestInstancesBuilder {
	s := &ReallocateRequestInstances{}
	b := &ReallocateRequestInstancesBuilder{s: s}
	return b
}

func (b *ReallocateRequestInstancesBuilder) InstanceId(v string) *ReallocateRequestInstancesBuilder {
	b.s.InstanceId = &v
	return b
}

func (b *ReallocateRequestInstancesBuilder) MachineId(v string) *ReallocateRequestInstancesBuilder {
	b.s.MachineId = &v
	return b
}

func (b *ReallocateRequestInstancesBuilder) Build() *ReallocateRequestInstances {
	return b.s
}


