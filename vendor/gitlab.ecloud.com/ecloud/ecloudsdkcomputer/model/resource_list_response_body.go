// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ResourceListResponseBody struct {

    // 是否配置关机还原
	ConfigShutdownRevert *bool `json:"configShutdownRevert,omitempty"`
    // 资源实例id
	InstanceId *string `json:"instanceId,omitempty"`
    // 桌面id
	MachineId *string `json:"machineId,omitempty"`
    // 邮箱
	Mail *string `json:"mail,omitempty"`
    // 手机号
	Mobile *string `json:"mobile,omitempty"`
    // 登录用户
	UserName *string `json:"userName,omitempty"`
    // 登录用户
	UserId *string `json:"userId,omitempty"`
    // 绑定状态，0：未关联，1：关联当前策略，2：关联其他策略
	BindStatus *string `json:"bindStatus,omitempty"`
    // 桌面名称
	MachineName *string `json:"machineName,omitempty"`
}

func (s ResourceListResponseBody) String() string {
	return utils.Beautify(s)
}

func (s ResourceListResponseBody) GoString() string {
	return s.String()
}

func (s ResourceListResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ResourceListResponseBody) SetConfigShutdownRevert(v bool) *ResourceListResponseBody {
	s.ConfigShutdownRevert = &v
	return s
}

func (s *ResourceListResponseBody) SetInstanceId(v string) *ResourceListResponseBody {
	s.InstanceId = &v
	return s
}

func (s *ResourceListResponseBody) SetMachineId(v string) *ResourceListResponseBody {
	s.MachineId = &v
	return s
}

func (s *ResourceListResponseBody) SetMail(v string) *ResourceListResponseBody {
	s.Mail = &v
	return s
}

func (s *ResourceListResponseBody) SetMobile(v string) *ResourceListResponseBody {
	s.Mobile = &v
	return s
}

func (s *ResourceListResponseBody) SetUserName(v string) *ResourceListResponseBody {
	s.UserName = &v
	return s
}

func (s *ResourceListResponseBody) SetUserId(v string) *ResourceListResponseBody {
	s.UserId = &v
	return s
}

func (s *ResourceListResponseBody) SetBindStatus(v string) *ResourceListResponseBody {
	s.BindStatus = &v
	return s
}

func (s *ResourceListResponseBody) SetMachineName(v string) *ResourceListResponseBody {
	s.MachineName = &v
	return s
}


type ResourceListResponseBodyBuilder struct {
	s *ResourceListResponseBody
}

func NewResourceListResponseBodyBuilder() *ResourceListResponseBodyBuilder {
	s := &ResourceListResponseBody{}
	b := &ResourceListResponseBodyBuilder{s: s}
	return b
}

func (b *ResourceListResponseBodyBuilder) ConfigShutdownRevert(v bool) *ResourceListResponseBodyBuilder {
	b.s.ConfigShutdownRevert = &v
	return b
}

func (b *ResourceListResponseBodyBuilder) InstanceId(v string) *ResourceListResponseBodyBuilder {
	b.s.InstanceId = &v
	return b
}

func (b *ResourceListResponseBodyBuilder) MachineId(v string) *ResourceListResponseBodyBuilder {
	b.s.MachineId = &v
	return b
}

func (b *ResourceListResponseBodyBuilder) Mail(v string) *ResourceListResponseBodyBuilder {
	b.s.Mail = &v
	return b
}

func (b *ResourceListResponseBodyBuilder) Mobile(v string) *ResourceListResponseBodyBuilder {
	b.s.Mobile = &v
	return b
}

func (b *ResourceListResponseBodyBuilder) UserName(v string) *ResourceListResponseBodyBuilder {
	b.s.UserName = &v
	return b
}

func (b *ResourceListResponseBodyBuilder) UserId(v string) *ResourceListResponseBodyBuilder {
	b.s.UserId = &v
	return b
}

func (b *ResourceListResponseBodyBuilder) BindStatus(v string) *ResourceListResponseBodyBuilder {
	b.s.BindStatus = &v
	return b
}

func (b *ResourceListResponseBodyBuilder) MachineName(v string) *ResourceListResponseBodyBuilder {
	b.s.MachineName = &v
	return b
}

func (b *ResourceListResponseBodyBuilder) Build() *ResourceListResponseBody {
	return b.s
}


