// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BindMachineBody struct {
    position.Body
    // 策略id
	PolicyId *string `json:"policyId,omitempty"`
    // 桌面id列表
	MachineIdList []string `json:"machineIdList,omitempty"`
}

func (s BindMachineBody) String() string {
	return utils.Beautify(s)
}

func (s BindMachineBody) GoString() string {
	return s.String()
}

func (s BindMachineBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BindMachineBody) SetPolicyId(v string) *BindMachineBody {
	s.PolicyId = &v
	return s
}

func (s *BindMachineBody) SetMachineIdList(v []string) *BindMachineBody {
	s.MachineIdList = v
	return s
}


type BindMachineBodyBuilder struct {
	s *BindMachineBody
}

func NewBindMachineBodyBuilder() *BindMachineBodyBuilder {
	s := &BindMachineBody{}
	b := &BindMachineBodyBuilder{s: s}
	return b
}

func (b *BindMachineBodyBuilder) PolicyId(v string) *BindMachineBodyBuilder {
	b.s.PolicyId = &v
	return b
}

func (b *BindMachineBodyBuilder) MachineIdList(v []string) *BindMachineBodyBuilder {
	b.s.MachineIdList = v
	return b
}

func (b *BindMachineBodyBuilder) Build() *BindMachineBody {
	return b.s
}


