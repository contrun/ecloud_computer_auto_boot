// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BindInstanceListBody struct {
    position.Body
    // 定时任务uuid
	PolicyUid *string `json:"policyUid,omitempty"`
    // 云电脑id
	MachineId *string `json:"machineId,omitempty"`
    // 订购资源实例id
	InstanceId *string `json:"instanceId,omitempty"`
    // 分页大小
	PageSize *int32 `json:"pageSize,omitempty"`
    // 当前页
	CurrentPage *int32 `json:"currentPage,omitempty"`
    // 登录账号
	UserName *string `json:"userName,omitempty"`
    // 云电脑名称
	MachineName *string `json:"machineName,omitempty"`
}

func (s BindInstanceListBody) String() string {
	return utils.Beautify(s)
}

func (s BindInstanceListBody) GoString() string {
	return s.String()
}

func (s BindInstanceListBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BindInstanceListBody) SetPolicyUid(v string) *BindInstanceListBody {
	s.PolicyUid = &v
	return s
}

func (s *BindInstanceListBody) SetMachineId(v string) *BindInstanceListBody {
	s.MachineId = &v
	return s
}

func (s *BindInstanceListBody) SetInstanceId(v string) *BindInstanceListBody {
	s.InstanceId = &v
	return s
}

func (s *BindInstanceListBody) SetPageSize(v int32) *BindInstanceListBody {
	s.PageSize = &v
	return s
}

func (s *BindInstanceListBody) SetCurrentPage(v int32) *BindInstanceListBody {
	s.CurrentPage = &v
	return s
}

func (s *BindInstanceListBody) SetUserName(v string) *BindInstanceListBody {
	s.UserName = &v
	return s
}

func (s *BindInstanceListBody) SetMachineName(v string) *BindInstanceListBody {
	s.MachineName = &v
	return s
}


type BindInstanceListBodyBuilder struct {
	s *BindInstanceListBody
}

func NewBindInstanceListBodyBuilder() *BindInstanceListBodyBuilder {
	s := &BindInstanceListBody{}
	b := &BindInstanceListBodyBuilder{s: s}
	return b
}

func (b *BindInstanceListBodyBuilder) PolicyUid(v string) *BindInstanceListBodyBuilder {
	b.s.PolicyUid = &v
	return b
}

func (b *BindInstanceListBodyBuilder) MachineId(v string) *BindInstanceListBodyBuilder {
	b.s.MachineId = &v
	return b
}

func (b *BindInstanceListBodyBuilder) InstanceId(v string) *BindInstanceListBodyBuilder {
	b.s.InstanceId = &v
	return b
}

func (b *BindInstanceListBodyBuilder) PageSize(v int32) *BindInstanceListBodyBuilder {
	b.s.PageSize = &v
	return b
}

func (b *BindInstanceListBodyBuilder) CurrentPage(v int32) *BindInstanceListBodyBuilder {
	b.s.CurrentPage = &v
	return b
}

func (b *BindInstanceListBodyBuilder) UserName(v string) *BindInstanceListBodyBuilder {
	b.s.UserName = &v
	return b
}

func (b *BindInstanceListBodyBuilder) MachineName(v string) *BindInstanceListBodyBuilder {
	b.s.MachineName = &v
	return b
}

func (b *BindInstanceListBodyBuilder) Build() *BindInstanceListBody {
	return b.s
}


