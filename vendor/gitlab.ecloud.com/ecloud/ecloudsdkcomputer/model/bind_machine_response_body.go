// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BindMachineResponseBody struct {

    // 描述
	AllocateMsg *string `json:"allocateMsg,omitempty"`
    // 失败的桌面id
	MachineId *string `json:"machineId,omitempty"`
    // 分配结果code，失败fail
	AllocateCode *string `json:"allocateCode,omitempty"`
}

func (s BindMachineResponseBody) String() string {
	return utils.Beautify(s)
}

func (s BindMachineResponseBody) GoString() string {
	return s.String()
}

func (s BindMachineResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BindMachineResponseBody) SetAllocateMsg(v string) *BindMachineResponseBody {
	s.AllocateMsg = &v
	return s
}

func (s *BindMachineResponseBody) SetMachineId(v string) *BindMachineResponseBody {
	s.MachineId = &v
	return s
}

func (s *BindMachineResponseBody) SetAllocateCode(v string) *BindMachineResponseBody {
	s.AllocateCode = &v
	return s
}


type BindMachineResponseBodyBuilder struct {
	s *BindMachineResponseBody
}

func NewBindMachineResponseBodyBuilder() *BindMachineResponseBodyBuilder {
	s := &BindMachineResponseBody{}
	b := &BindMachineResponseBodyBuilder{s: s}
	return b
}

func (b *BindMachineResponseBodyBuilder) AllocateMsg(v string) *BindMachineResponseBodyBuilder {
	b.s.AllocateMsg = &v
	return b
}

func (b *BindMachineResponseBodyBuilder) MachineId(v string) *BindMachineResponseBodyBuilder {
	b.s.MachineId = &v
	return b
}

func (b *BindMachineResponseBodyBuilder) AllocateCode(v string) *BindMachineResponseBodyBuilder {
	b.s.AllocateCode = &v
	return b
}

func (b *BindMachineResponseBodyBuilder) Build() *BindMachineResponseBody {
	return b.s
}


