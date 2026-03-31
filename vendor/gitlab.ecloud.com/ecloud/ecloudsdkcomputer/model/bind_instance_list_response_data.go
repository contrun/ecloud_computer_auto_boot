// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BindInstanceListResponseData struct {

    // 订购资源实例id
	InstanceId *string `json:"instanceId,omitempty"`
    // 云电脑id
	MachineId *string `json:"machineId,omitempty"`
    // 手机号
	Mobile *string `json:"mobile,omitempty"`
    // 登录账号
	MopUserName *string `json:"mopUserName,omitempty"`
    // 登录账号
	UserName *string `json:"userName,omitempty"`
    // 邮箱
	Email *string `json:"email,omitempty"`
    // 云电脑名称
	MachineName *string `json:"machineName,omitempty"`
}

func (s BindInstanceListResponseData) String() string {
	return utils.Beautify(s)
}

func (s BindInstanceListResponseData) GoString() string {
	return s.String()
}

func (s BindInstanceListResponseData) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BindInstanceListResponseData) SetInstanceId(v string) *BindInstanceListResponseData {
	s.InstanceId = &v
	return s
}

func (s *BindInstanceListResponseData) SetMachineId(v string) *BindInstanceListResponseData {
	s.MachineId = &v
	return s
}

func (s *BindInstanceListResponseData) SetMobile(v string) *BindInstanceListResponseData {
	s.Mobile = &v
	return s
}

func (s *BindInstanceListResponseData) SetMopUserName(v string) *BindInstanceListResponseData {
	s.MopUserName = &v
	return s
}

func (s *BindInstanceListResponseData) SetUserName(v string) *BindInstanceListResponseData {
	s.UserName = &v
	return s
}

func (s *BindInstanceListResponseData) SetEmail(v string) *BindInstanceListResponseData {
	s.Email = &v
	return s
}

func (s *BindInstanceListResponseData) SetMachineName(v string) *BindInstanceListResponseData {
	s.MachineName = &v
	return s
}


type BindInstanceListResponseDataBuilder struct {
	s *BindInstanceListResponseData
}

func NewBindInstanceListResponseDataBuilder() *BindInstanceListResponseDataBuilder {
	s := &BindInstanceListResponseData{}
	b := &BindInstanceListResponseDataBuilder{s: s}
	return b
}

func (b *BindInstanceListResponseDataBuilder) InstanceId(v string) *BindInstanceListResponseDataBuilder {
	b.s.InstanceId = &v
	return b
}

func (b *BindInstanceListResponseDataBuilder) MachineId(v string) *BindInstanceListResponseDataBuilder {
	b.s.MachineId = &v
	return b
}

func (b *BindInstanceListResponseDataBuilder) Mobile(v string) *BindInstanceListResponseDataBuilder {
	b.s.Mobile = &v
	return b
}

func (b *BindInstanceListResponseDataBuilder) MopUserName(v string) *BindInstanceListResponseDataBuilder {
	b.s.MopUserName = &v
	return b
}

func (b *BindInstanceListResponseDataBuilder) UserName(v string) *BindInstanceListResponseDataBuilder {
	b.s.UserName = &v
	return b
}

func (b *BindInstanceListResponseDataBuilder) Email(v string) *BindInstanceListResponseDataBuilder {
	b.s.Email = &v
	return b
}

func (b *BindInstanceListResponseDataBuilder) MachineName(v string) *BindInstanceListResponseDataBuilder {
	b.s.MachineName = &v
	return b
}

func (b *BindInstanceListResponseDataBuilder) Build() *BindInstanceListResponseData {
	return b.s
}


