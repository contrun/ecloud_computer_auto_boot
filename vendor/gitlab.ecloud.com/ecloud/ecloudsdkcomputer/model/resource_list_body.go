// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ResourceListBody struct {
    position.Body
    // 定时任务uid
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

func (s ResourceListBody) String() string {
	return utils.Beautify(s)
}

func (s ResourceListBody) GoString() string {
	return s.String()
}

func (s ResourceListBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ResourceListBody) SetPolicyUid(v string) *ResourceListBody {
	s.PolicyUid = &v
	return s
}

func (s *ResourceListBody) SetMachineId(v string) *ResourceListBody {
	s.MachineId = &v
	return s
}

func (s *ResourceListBody) SetInstanceId(v string) *ResourceListBody {
	s.InstanceId = &v
	return s
}

func (s *ResourceListBody) SetPageSize(v int32) *ResourceListBody {
	s.PageSize = &v
	return s
}

func (s *ResourceListBody) SetCurrentPage(v int32) *ResourceListBody {
	s.CurrentPage = &v
	return s
}

func (s *ResourceListBody) SetUserName(v string) *ResourceListBody {
	s.UserName = &v
	return s
}

func (s *ResourceListBody) SetMachineName(v string) *ResourceListBody {
	s.MachineName = &v
	return s
}


type ResourceListBodyBuilder struct {
	s *ResourceListBody
}

func NewResourceListBodyBuilder() *ResourceListBodyBuilder {
	s := &ResourceListBody{}
	b := &ResourceListBodyBuilder{s: s}
	return b
}

func (b *ResourceListBodyBuilder) PolicyUid(v string) *ResourceListBodyBuilder {
	b.s.PolicyUid = &v
	return b
}

func (b *ResourceListBodyBuilder) MachineId(v string) *ResourceListBodyBuilder {
	b.s.MachineId = &v
	return b
}

func (b *ResourceListBodyBuilder) InstanceId(v string) *ResourceListBodyBuilder {
	b.s.InstanceId = &v
	return b
}

func (b *ResourceListBodyBuilder) PageSize(v int32) *ResourceListBodyBuilder {
	b.s.PageSize = &v
	return b
}

func (b *ResourceListBodyBuilder) CurrentPage(v int32) *ResourceListBodyBuilder {
	b.s.CurrentPage = &v
	return b
}

func (b *ResourceListBodyBuilder) UserName(v string) *ResourceListBodyBuilder {
	b.s.UserName = &v
	return b
}

func (b *ResourceListBodyBuilder) MachineName(v string) *ResourceListBodyBuilder {
	b.s.MachineName = &v
	return b
}

func (b *ResourceListBodyBuilder) Build() *ResourceListBody {
	return b.s
}


