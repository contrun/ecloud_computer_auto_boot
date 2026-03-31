// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type AllocationBody struct {
    position.Body
    // 实例ID列表
	Instances *[]AllocationRequestInstances `json:"instances,omitempty"`
    // 云电脑用户id
	UserId *string `json:"userId,omitempty"`
}

func (s AllocationBody) String() string {
	return utils.Beautify(s)
}

func (s AllocationBody) GoString() string {
	return s.String()
}

func (s AllocationBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *AllocationBody) SetInstances(v []AllocationRequestInstances) *AllocationBody {
	s.Instances = &v
	return s
}

func (s *AllocationBody) SetUserId(v string) *AllocationBody {
	s.UserId = &v
	return s
}


type AllocationBodyBuilder struct {
	s *AllocationBody
}

func NewAllocationBodyBuilder() *AllocationBodyBuilder {
	s := &AllocationBody{}
	b := &AllocationBodyBuilder{s: s}
	return b
}

func (b *AllocationBodyBuilder) Instances(v []AllocationRequestInstances) *AllocationBodyBuilder {
	b.s.Instances = &v
	return b
}

func (b *AllocationBodyBuilder) UserId(v string) *AllocationBodyBuilder {
	b.s.UserId = &v
	return b
}

func (b *AllocationBodyBuilder) Build() *AllocationBody {
	return b.s
}


