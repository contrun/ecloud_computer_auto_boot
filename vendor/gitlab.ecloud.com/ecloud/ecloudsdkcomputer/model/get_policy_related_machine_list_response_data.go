// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetPolicyRelatedMachineListResponseData struct {

    // 云电脑id
	MachineId *string `json:"machineId,omitempty"`
    // 策略id
	PolicyId *string `json:"policyId,omitempty"`
    // 策略名称
	PolicyName *string `json:"policyName,omitempty"`
    // 云电脑绑定策略状态,0:绑定成功,1:解绑成功,2:处理中,3:绑定失败,4:解绑失败
	DeskBindExtPolicyStatus *string `json:"deskBindExtPolicyStatus,omitempty"`
    // 登录账号
	UserName *string `json:"userName,omitempty"`
    // 云电脑名称
	MachineName *string `json:"machineName,omitempty"`
}

func (s GetPolicyRelatedMachineListResponseData) String() string {
	return utils.Beautify(s)
}

func (s GetPolicyRelatedMachineListResponseData) GoString() string {
	return s.String()
}

func (s GetPolicyRelatedMachineListResponseData) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetPolicyRelatedMachineListResponseData) SetMachineId(v string) *GetPolicyRelatedMachineListResponseData {
	s.MachineId = &v
	return s
}

func (s *GetPolicyRelatedMachineListResponseData) SetPolicyId(v string) *GetPolicyRelatedMachineListResponseData {
	s.PolicyId = &v
	return s
}

func (s *GetPolicyRelatedMachineListResponseData) SetPolicyName(v string) *GetPolicyRelatedMachineListResponseData {
	s.PolicyName = &v
	return s
}

func (s *GetPolicyRelatedMachineListResponseData) SetDeskBindExtPolicyStatus(v string) *GetPolicyRelatedMachineListResponseData {
	s.DeskBindExtPolicyStatus = &v
	return s
}

func (s *GetPolicyRelatedMachineListResponseData) SetUserName(v string) *GetPolicyRelatedMachineListResponseData {
	s.UserName = &v
	return s
}

func (s *GetPolicyRelatedMachineListResponseData) SetMachineName(v string) *GetPolicyRelatedMachineListResponseData {
	s.MachineName = &v
	return s
}


type GetPolicyRelatedMachineListResponseDataBuilder struct {
	s *GetPolicyRelatedMachineListResponseData
}

func NewGetPolicyRelatedMachineListResponseDataBuilder() *GetPolicyRelatedMachineListResponseDataBuilder {
	s := &GetPolicyRelatedMachineListResponseData{}
	b := &GetPolicyRelatedMachineListResponseDataBuilder{s: s}
	return b
}

func (b *GetPolicyRelatedMachineListResponseDataBuilder) MachineId(v string) *GetPolicyRelatedMachineListResponseDataBuilder {
	b.s.MachineId = &v
	return b
}

func (b *GetPolicyRelatedMachineListResponseDataBuilder) PolicyId(v string) *GetPolicyRelatedMachineListResponseDataBuilder {
	b.s.PolicyId = &v
	return b
}

func (b *GetPolicyRelatedMachineListResponseDataBuilder) PolicyName(v string) *GetPolicyRelatedMachineListResponseDataBuilder {
	b.s.PolicyName = &v
	return b
}

func (b *GetPolicyRelatedMachineListResponseDataBuilder) DeskBindExtPolicyStatus(v string) *GetPolicyRelatedMachineListResponseDataBuilder {
	b.s.DeskBindExtPolicyStatus = &v
	return b
}

func (b *GetPolicyRelatedMachineListResponseDataBuilder) UserName(v string) *GetPolicyRelatedMachineListResponseDataBuilder {
	b.s.UserName = &v
	return b
}

func (b *GetPolicyRelatedMachineListResponseDataBuilder) MachineName(v string) *GetPolicyRelatedMachineListResponseDataBuilder {
	b.s.MachineName = &v
	return b
}

func (b *GetPolicyRelatedMachineListResponseDataBuilder) Build() *GetPolicyRelatedMachineListResponseData {
	return b.s
}


