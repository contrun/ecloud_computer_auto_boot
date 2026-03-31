// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type AllocationRequestInstances struct {

    // 实例ID
	InstanceId *string `json:"instanceId,omitempty"`
}

func (s AllocationRequestInstances) String() string {
	return utils.Beautify(s)
}

func (s AllocationRequestInstances) GoString() string {
	return s.String()
}

func (s AllocationRequestInstances) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *AllocationRequestInstances) SetInstanceId(v string) *AllocationRequestInstances {
	s.InstanceId = &v
	return s
}


type AllocationRequestInstancesBuilder struct {
	s *AllocationRequestInstances
}

func NewAllocationRequestInstancesBuilder() *AllocationRequestInstancesBuilder {
	s := &AllocationRequestInstances{}
	b := &AllocationRequestInstancesBuilder{s: s}
	return b
}

func (b *AllocationRequestInstancesBuilder) InstanceId(v string) *AllocationRequestInstancesBuilder {
	b.s.InstanceId = &v
	return b
}

func (b *AllocationRequestInstancesBuilder) Build() *AllocationRequestInstances {
	return b.s
}


