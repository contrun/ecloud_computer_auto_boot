// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ActiveConnectPageBody struct {
    position.Body
    // 每页个数
	PageSize *int32 `json:"pageSize,omitempty"`
    // 开始时间
	StartTime *string `json:"startTime,omitempty"`
    // 结束时间
	EndTime *string `json:"endTime,omitempty"`
    // 当前页数
	Page *int32 `json:"page,omitempty"`
    // 登录账号
	UserName *string `json:"userName,omitempty"`
    // 设备名称
	DeviceName *string `json:"deviceName,omitempty"`
    // 电脑名称
	MachineName *string `json:"machineName,omitempty"`
}

func (s ActiveConnectPageBody) String() string {
	return utils.Beautify(s)
}

func (s ActiveConnectPageBody) GoString() string {
	return s.String()
}

func (s ActiveConnectPageBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ActiveConnectPageBody) SetPageSize(v int32) *ActiveConnectPageBody {
	s.PageSize = &v
	return s
}

func (s *ActiveConnectPageBody) SetStartTime(v string) *ActiveConnectPageBody {
	s.StartTime = &v
	return s
}

func (s *ActiveConnectPageBody) SetEndTime(v string) *ActiveConnectPageBody {
	s.EndTime = &v
	return s
}

func (s *ActiveConnectPageBody) SetPage(v int32) *ActiveConnectPageBody {
	s.Page = &v
	return s
}

func (s *ActiveConnectPageBody) SetUserName(v string) *ActiveConnectPageBody {
	s.UserName = &v
	return s
}

func (s *ActiveConnectPageBody) SetDeviceName(v string) *ActiveConnectPageBody {
	s.DeviceName = &v
	return s
}

func (s *ActiveConnectPageBody) SetMachineName(v string) *ActiveConnectPageBody {
	s.MachineName = &v
	return s
}


type ActiveConnectPageBodyBuilder struct {
	s *ActiveConnectPageBody
}

func NewActiveConnectPageBodyBuilder() *ActiveConnectPageBodyBuilder {
	s := &ActiveConnectPageBody{}
	b := &ActiveConnectPageBodyBuilder{s: s}
	return b
}

func (b *ActiveConnectPageBodyBuilder) PageSize(v int32) *ActiveConnectPageBodyBuilder {
	b.s.PageSize = &v
	return b
}

func (b *ActiveConnectPageBodyBuilder) StartTime(v string) *ActiveConnectPageBodyBuilder {
	b.s.StartTime = &v
	return b
}

func (b *ActiveConnectPageBodyBuilder) EndTime(v string) *ActiveConnectPageBodyBuilder {
	b.s.EndTime = &v
	return b
}

func (b *ActiveConnectPageBodyBuilder) Page(v int32) *ActiveConnectPageBodyBuilder {
	b.s.Page = &v
	return b
}

func (b *ActiveConnectPageBodyBuilder) UserName(v string) *ActiveConnectPageBodyBuilder {
	b.s.UserName = &v
	return b
}

func (b *ActiveConnectPageBodyBuilder) DeviceName(v string) *ActiveConnectPageBodyBuilder {
	b.s.DeviceName = &v
	return b
}

func (b *ActiveConnectPageBodyBuilder) MachineName(v string) *ActiveConnectPageBodyBuilder {
	b.s.MachineName = &v
	return b
}

func (b *ActiveConnectPageBodyBuilder) Build() *ActiveConnectPageBody {
	return b.s
}


