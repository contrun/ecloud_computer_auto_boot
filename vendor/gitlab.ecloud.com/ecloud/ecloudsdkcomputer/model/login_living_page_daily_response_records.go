// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type LoginLivingPageDailyResponseRecords struct {

    // 活跃用户数
	LiveUserNum *int32 `json:"liveUserNum,omitempty"`
    // 日期
	DailyTime *string `json:"dailyTime,omitempty"`
    // 平均每天总在线时长
	AvgDailyTotalOnline *string `json:"avgDailyTotalOnline,omitempty"`
    // 日期
	MonthlyTime *string `json:"monthlyTime,omitempty"`
    // 总在线时长
	TotalOnline *string `json:"totalOnline,omitempty"`
    // 平均每天在线时长
	AveDailyOnlineTime *string `json:"aveDailyOnlineTime,omitempty"`
    // 平均在线时长
	AveOnlineTime *string `json:"aveOnlineTime,omitempty"`
    // 总在线时长
	TotalOnlineTime *string `json:"totalOnlineTime,omitempty"`
}

func (s LoginLivingPageDailyResponseRecords) String() string {
	return utils.Beautify(s)
}

func (s LoginLivingPageDailyResponseRecords) GoString() string {
	return s.String()
}

func (s LoginLivingPageDailyResponseRecords) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *LoginLivingPageDailyResponseRecords) SetLiveUserNum(v int32) *LoginLivingPageDailyResponseRecords {
	s.LiveUserNum = &v
	return s
}

func (s *LoginLivingPageDailyResponseRecords) SetDailyTime(v string) *LoginLivingPageDailyResponseRecords {
	s.DailyTime = &v
	return s
}

func (s *LoginLivingPageDailyResponseRecords) SetAvgDailyTotalOnline(v string) *LoginLivingPageDailyResponseRecords {
	s.AvgDailyTotalOnline = &v
	return s
}

func (s *LoginLivingPageDailyResponseRecords) SetMonthlyTime(v string) *LoginLivingPageDailyResponseRecords {
	s.MonthlyTime = &v
	return s
}

func (s *LoginLivingPageDailyResponseRecords) SetTotalOnline(v string) *LoginLivingPageDailyResponseRecords {
	s.TotalOnline = &v
	return s
}

func (s *LoginLivingPageDailyResponseRecords) SetAveDailyOnlineTime(v string) *LoginLivingPageDailyResponseRecords {
	s.AveDailyOnlineTime = &v
	return s
}

func (s *LoginLivingPageDailyResponseRecords) SetAveOnlineTime(v string) *LoginLivingPageDailyResponseRecords {
	s.AveOnlineTime = &v
	return s
}

func (s *LoginLivingPageDailyResponseRecords) SetTotalOnlineTime(v string) *LoginLivingPageDailyResponseRecords {
	s.TotalOnlineTime = &v
	return s
}


type LoginLivingPageDailyResponseRecordsBuilder struct {
	s *LoginLivingPageDailyResponseRecords
}

func NewLoginLivingPageDailyResponseRecordsBuilder() *LoginLivingPageDailyResponseRecordsBuilder {
	s := &LoginLivingPageDailyResponseRecords{}
	b := &LoginLivingPageDailyResponseRecordsBuilder{s: s}
	return b
}

func (b *LoginLivingPageDailyResponseRecordsBuilder) LiveUserNum(v int32) *LoginLivingPageDailyResponseRecordsBuilder {
	b.s.LiveUserNum = &v
	return b
}

func (b *LoginLivingPageDailyResponseRecordsBuilder) DailyTime(v string) *LoginLivingPageDailyResponseRecordsBuilder {
	b.s.DailyTime = &v
	return b
}

func (b *LoginLivingPageDailyResponseRecordsBuilder) AvgDailyTotalOnline(v string) *LoginLivingPageDailyResponseRecordsBuilder {
	b.s.AvgDailyTotalOnline = &v
	return b
}

func (b *LoginLivingPageDailyResponseRecordsBuilder) MonthlyTime(v string) *LoginLivingPageDailyResponseRecordsBuilder {
	b.s.MonthlyTime = &v
	return b
}

func (b *LoginLivingPageDailyResponseRecordsBuilder) TotalOnline(v string) *LoginLivingPageDailyResponseRecordsBuilder {
	b.s.TotalOnline = &v
	return b
}

func (b *LoginLivingPageDailyResponseRecordsBuilder) AveDailyOnlineTime(v string) *LoginLivingPageDailyResponseRecordsBuilder {
	b.s.AveDailyOnlineTime = &v
	return b
}

func (b *LoginLivingPageDailyResponseRecordsBuilder) AveOnlineTime(v string) *LoginLivingPageDailyResponseRecordsBuilder {
	b.s.AveOnlineTime = &v
	return b
}

func (b *LoginLivingPageDailyResponseRecordsBuilder) TotalOnlineTime(v string) *LoginLivingPageDailyResponseRecordsBuilder {
	b.s.TotalOnlineTime = &v
	return b
}

func (b *LoginLivingPageDailyResponseRecordsBuilder) Build() *LoginLivingPageDailyResponseRecords {
	return b.s
}


