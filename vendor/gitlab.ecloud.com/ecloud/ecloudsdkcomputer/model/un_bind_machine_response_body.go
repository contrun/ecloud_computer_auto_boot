// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UnBindMachineResponseBody struct {

    // 描述
	AllocateMsg *string `json:"allocateMsg,omitempty"`
    // 失败的桌面id
	MachineId *string `json:"machineId,omitempty"`
    // 分配结果code，失败fail
	AllocateCode *string `json:"allocateCode,omitempty"`
}

func (s UnBindMachineResponseBody) String() string {
	return utils.Beautify(s)
}

func (s UnBindMachineResponseBody) GoString() string {
	return s.String()
}

func (s UnBindMachineResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UnBindMachineResponseBody) SetAllocateMsg(v string) *UnBindMachineResponseBody {
	s.AllocateMsg = &v
	return s
}

func (s *UnBindMachineResponseBody) SetMachineId(v string) *UnBindMachineResponseBody {
	s.MachineId = &v
	return s
}

func (s *UnBindMachineResponseBody) SetAllocateCode(v string) *UnBindMachineResponseBody {
	s.AllocateCode = &v
	return s
}


type UnBindMachineResponseBodyBuilder struct {
	s *UnBindMachineResponseBody
}

func NewUnBindMachineResponseBodyBuilder() *UnBindMachineResponseBodyBuilder {
	s := &UnBindMachineResponseBody{}
	b := &UnBindMachineResponseBodyBuilder{s: s}
	return b
}

func (b *UnBindMachineResponseBodyBuilder) AllocateMsg(v string) *UnBindMachineResponseBodyBuilder {
	b.s.AllocateMsg = &v
	return b
}

func (b *UnBindMachineResponseBodyBuilder) MachineId(v string) *UnBindMachineResponseBodyBuilder {
	b.s.MachineId = &v
	return b
}

func (b *UnBindMachineResponseBodyBuilder) AllocateCode(v string) *UnBindMachineResponseBodyBuilder {
	b.s.AllocateCode = &v
	return b
}

func (b *UnBindMachineResponseBodyBuilder) Build() *UnBindMachineResponseBody {
	return b.s
}


