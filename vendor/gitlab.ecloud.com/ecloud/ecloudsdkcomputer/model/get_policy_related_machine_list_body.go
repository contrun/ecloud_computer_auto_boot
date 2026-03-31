// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetPolicyRelatedMachineListBody struct {
    position.Body
    // 云电脑id
	MachineId *string `json:"machineId,omitempty"`
    // 策略id
	PolicyId *string `json:"policyId,omitempty"`
    // 分页大小
	PageSize *int32 `json:"pageSize,omitempty"`
    // 当前页
	CurrentPage *int32 `json:"currentPage,omitempty"`
    // 登录账号
	UserName *string `json:"userName,omitempty"`
    // 云电脑名称
	MachineName *string `json:"machineName,omitempty"`
    // 策略分配状态 0:绑定成功,1:解绑成功,2:处理中,3:绑定失败,4:解绑失败
	Status *string `json:"status,omitempty"`
}

func (s GetPolicyRelatedMachineListBody) String() string {
	return utils.Beautify(s)
}

func (s GetPolicyRelatedMachineListBody) GoString() string {
	return s.String()
}

func (s GetPolicyRelatedMachineListBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetPolicyRelatedMachineListBody) SetMachineId(v string) *GetPolicyRelatedMachineListBody {
	s.MachineId = &v
	return s
}

func (s *GetPolicyRelatedMachineListBody) SetPolicyId(v string) *GetPolicyRelatedMachineListBody {
	s.PolicyId = &v
	return s
}

func (s *GetPolicyRelatedMachineListBody) SetPageSize(v int32) *GetPolicyRelatedMachineListBody {
	s.PageSize = &v
	return s
}

func (s *GetPolicyRelatedMachineListBody) SetCurrentPage(v int32) *GetPolicyRelatedMachineListBody {
	s.CurrentPage = &v
	return s
}

func (s *GetPolicyRelatedMachineListBody) SetUserName(v string) *GetPolicyRelatedMachineListBody {
	s.UserName = &v
	return s
}

func (s *GetPolicyRelatedMachineListBody) SetMachineName(v string) *GetPolicyRelatedMachineListBody {
	s.MachineName = &v
	return s
}

func (s *GetPolicyRelatedMachineListBody) SetStatus(v string) *GetPolicyRelatedMachineListBody {
	s.Status = &v
	return s
}


type GetPolicyRelatedMachineListBodyBuilder struct {
	s *GetPolicyRelatedMachineListBody
}

func NewGetPolicyRelatedMachineListBodyBuilder() *GetPolicyRelatedMachineListBodyBuilder {
	s := &GetPolicyRelatedMachineListBody{}
	b := &GetPolicyRelatedMachineListBodyBuilder{s: s}
	return b
}

func (b *GetPolicyRelatedMachineListBodyBuilder) MachineId(v string) *GetPolicyRelatedMachineListBodyBuilder {
	b.s.MachineId = &v
	return b
}

func (b *GetPolicyRelatedMachineListBodyBuilder) PolicyId(v string) *GetPolicyRelatedMachineListBodyBuilder {
	b.s.PolicyId = &v
	return b
}

func (b *GetPolicyRelatedMachineListBodyBuilder) PageSize(v int32) *GetPolicyRelatedMachineListBodyBuilder {
	b.s.PageSize = &v
	return b
}

func (b *GetPolicyRelatedMachineListBodyBuilder) CurrentPage(v int32) *GetPolicyRelatedMachineListBodyBuilder {
	b.s.CurrentPage = &v
	return b
}

func (b *GetPolicyRelatedMachineListBodyBuilder) UserName(v string) *GetPolicyRelatedMachineListBodyBuilder {
	b.s.UserName = &v
	return b
}

func (b *GetPolicyRelatedMachineListBodyBuilder) MachineName(v string) *GetPolicyRelatedMachineListBodyBuilder {
	b.s.MachineName = &v
	return b
}

func (b *GetPolicyRelatedMachineListBodyBuilder) Status(v string) *GetPolicyRelatedMachineListBodyBuilder {
	b.s.Status = &v
	return b
}

func (b *GetPolicyRelatedMachineListBodyBuilder) Build() *GetPolicyRelatedMachineListBody {
	return b.s
}


