// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type HistoryConnectPageBody struct {
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

func (s HistoryConnectPageBody) String() string {
	return utils.Beautify(s)
}

func (s HistoryConnectPageBody) GoString() string {
	return s.String()
}

func (s HistoryConnectPageBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *HistoryConnectPageBody) SetPageSize(v int32) *HistoryConnectPageBody {
	s.PageSize = &v
	return s
}

func (s *HistoryConnectPageBody) SetStartTime(v string) *HistoryConnectPageBody {
	s.StartTime = &v
	return s
}

func (s *HistoryConnectPageBody) SetEndTime(v string) *HistoryConnectPageBody {
	s.EndTime = &v
	return s
}

func (s *HistoryConnectPageBody) SetPage(v int32) *HistoryConnectPageBody {
	s.Page = &v
	return s
}

func (s *HistoryConnectPageBody) SetUserName(v string) *HistoryConnectPageBody {
	s.UserName = &v
	return s
}

func (s *HistoryConnectPageBody) SetDeviceName(v string) *HistoryConnectPageBody {
	s.DeviceName = &v
	return s
}

func (s *HistoryConnectPageBody) SetMachineName(v string) *HistoryConnectPageBody {
	s.MachineName = &v
	return s
}


type HistoryConnectPageBodyBuilder struct {
	s *HistoryConnectPageBody
}

func NewHistoryConnectPageBodyBuilder() *HistoryConnectPageBodyBuilder {
	s := &HistoryConnectPageBody{}
	b := &HistoryConnectPageBodyBuilder{s: s}
	return b
}

func (b *HistoryConnectPageBodyBuilder) PageSize(v int32) *HistoryConnectPageBodyBuilder {
	b.s.PageSize = &v
	return b
}

func (b *HistoryConnectPageBodyBuilder) StartTime(v string) *HistoryConnectPageBodyBuilder {
	b.s.StartTime = &v
	return b
}

func (b *HistoryConnectPageBodyBuilder) EndTime(v string) *HistoryConnectPageBodyBuilder {
	b.s.EndTime = &v
	return b
}

func (b *HistoryConnectPageBodyBuilder) Page(v int32) *HistoryConnectPageBodyBuilder {
	b.s.Page = &v
	return b
}

func (b *HistoryConnectPageBodyBuilder) UserName(v string) *HistoryConnectPageBodyBuilder {
	b.s.UserName = &v
	return b
}

func (b *HistoryConnectPageBodyBuilder) DeviceName(v string) *HistoryConnectPageBodyBuilder {
	b.s.DeviceName = &v
	return b
}

func (b *HistoryConnectPageBodyBuilder) MachineName(v string) *HistoryConnectPageBodyBuilder {
	b.s.MachineName = &v
	return b
}

func (b *HistoryConnectPageBodyBuilder) Build() *HistoryConnectPageBody {
	return b.s
}


