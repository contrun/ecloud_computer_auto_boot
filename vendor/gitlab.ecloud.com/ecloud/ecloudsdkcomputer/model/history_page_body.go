// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type HistoryPageBody struct {
    position.Body
    // 分页大小
	PageSize *string `json:"pageSize,omitempty"`
    // 开始时间
	StartTime *string `json:"startTime,omitempty"`
    // 结束时间
	EndTime *string `json:"endTime,omitempty"`
    // 当前页
	Page *string `json:"page,omitempty"`
    // 登录账号
	UserName *string `json:"userName,omitempty"`
    // 设备名称
	DeviceName *string `json:"deviceName,omitempty"`
}

func (s HistoryPageBody) String() string {
	return utils.Beautify(s)
}

func (s HistoryPageBody) GoString() string {
	return s.String()
}

func (s HistoryPageBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *HistoryPageBody) SetPageSize(v string) *HistoryPageBody {
	s.PageSize = &v
	return s
}

func (s *HistoryPageBody) SetStartTime(v string) *HistoryPageBody {
	s.StartTime = &v
	return s
}

func (s *HistoryPageBody) SetEndTime(v string) *HistoryPageBody {
	s.EndTime = &v
	return s
}

func (s *HistoryPageBody) SetPage(v string) *HistoryPageBody {
	s.Page = &v
	return s
}

func (s *HistoryPageBody) SetUserName(v string) *HistoryPageBody {
	s.UserName = &v
	return s
}

func (s *HistoryPageBody) SetDeviceName(v string) *HistoryPageBody {
	s.DeviceName = &v
	return s
}


type HistoryPageBodyBuilder struct {
	s *HistoryPageBody
}

func NewHistoryPageBodyBuilder() *HistoryPageBodyBuilder {
	s := &HistoryPageBody{}
	b := &HistoryPageBodyBuilder{s: s}
	return b
}

func (b *HistoryPageBodyBuilder) PageSize(v string) *HistoryPageBodyBuilder {
	b.s.PageSize = &v
	return b
}

func (b *HistoryPageBodyBuilder) StartTime(v string) *HistoryPageBodyBuilder {
	b.s.StartTime = &v
	return b
}

func (b *HistoryPageBodyBuilder) EndTime(v string) *HistoryPageBodyBuilder {
	b.s.EndTime = &v
	return b
}

func (b *HistoryPageBodyBuilder) Page(v string) *HistoryPageBodyBuilder {
	b.s.Page = &v
	return b
}

func (b *HistoryPageBodyBuilder) UserName(v string) *HistoryPageBodyBuilder {
	b.s.UserName = &v
	return b
}

func (b *HistoryPageBodyBuilder) DeviceName(v string) *HistoryPageBodyBuilder {
	b.s.DeviceName = &v
	return b
}

func (b *HistoryPageBodyBuilder) Build() *HistoryPageBody {
	return b.s
}


