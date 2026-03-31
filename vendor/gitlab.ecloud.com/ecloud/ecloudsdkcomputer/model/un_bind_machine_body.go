// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UnBindMachineBody struct {
    position.Body
    // 策略id
	PolicyId *string `json:"policyId,omitempty"`
    // 桌面id列表
	MachineIdList []string `json:"machineIdList,omitempty"`
}

func (s UnBindMachineBody) String() string {
	return utils.Beautify(s)
}

func (s UnBindMachineBody) GoString() string {
	return s.String()
}

func (s UnBindMachineBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UnBindMachineBody) SetPolicyId(v string) *UnBindMachineBody {
	s.PolicyId = &v
	return s
}

func (s *UnBindMachineBody) SetMachineIdList(v []string) *UnBindMachineBody {
	s.MachineIdList = v
	return s
}


type UnBindMachineBodyBuilder struct {
	s *UnBindMachineBody
}

func NewUnBindMachineBodyBuilder() *UnBindMachineBodyBuilder {
	s := &UnBindMachineBody{}
	b := &UnBindMachineBodyBuilder{s: s}
	return b
}

func (b *UnBindMachineBodyBuilder) PolicyId(v string) *UnBindMachineBodyBuilder {
	b.s.PolicyId = &v
	return b
}

func (b *UnBindMachineBodyBuilder) MachineIdList(v []string) *UnBindMachineBodyBuilder {
	b.s.MachineIdList = v
	return b
}

func (b *UnBindMachineBodyBuilder) Build() *UnBindMachineBody {
	return b.s
}


