// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type LoginLivingPageMonthlyResponseRecords struct {

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

func (s LoginLivingPageMonthlyResponseRecords) String() string {
	return utils.Beautify(s)
}

func (s LoginLivingPageMonthlyResponseRecords) GoString() string {
	return s.String()
}

func (s LoginLivingPageMonthlyResponseRecords) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *LoginLivingPageMonthlyResponseRecords) SetLiveUserNum(v int32) *LoginLivingPageMonthlyResponseRecords {
	s.LiveUserNum = &v
	return s
}

func (s *LoginLivingPageMonthlyResponseRecords) SetDailyTime(v string) *LoginLivingPageMonthlyResponseRecords {
	s.DailyTime = &v
	return s
}

func (s *LoginLivingPageMonthlyResponseRecords) SetAvgDailyTotalOnline(v string) *LoginLivingPageMonthlyResponseRecords {
	s.AvgDailyTotalOnline = &v
	return s
}

func (s *LoginLivingPageMonthlyResponseRecords) SetMonthlyTime(v string) *LoginLivingPageMonthlyResponseRecords {
	s.MonthlyTime = &v
	return s
}

func (s *LoginLivingPageMonthlyResponseRecords) SetTotalOnline(v string) *LoginLivingPageMonthlyResponseRecords {
	s.TotalOnline = &v
	return s
}

func (s *LoginLivingPageMonthlyResponseRecords) SetAveDailyOnlineTime(v string) *LoginLivingPageMonthlyResponseRecords {
	s.AveDailyOnlineTime = &v
	return s
}

func (s *LoginLivingPageMonthlyResponseRecords) SetAveOnlineTime(v string) *LoginLivingPageMonthlyResponseRecords {
	s.AveOnlineTime = &v
	return s
}

func (s *LoginLivingPageMonthlyResponseRecords) SetTotalOnlineTime(v string) *LoginLivingPageMonthlyResponseRecords {
	s.TotalOnlineTime = &v
	return s
}


type LoginLivingPageMonthlyResponseRecordsBuilder struct {
	s *LoginLivingPageMonthlyResponseRecords
}

func NewLoginLivingPageMonthlyResponseRecordsBuilder() *LoginLivingPageMonthlyResponseRecordsBuilder {
	s := &LoginLivingPageMonthlyResponseRecords{}
	b := &LoginLivingPageMonthlyResponseRecordsBuilder{s: s}
	return b
}

func (b *LoginLivingPageMonthlyResponseRecordsBuilder) LiveUserNum(v int32) *LoginLivingPageMonthlyResponseRecordsBuilder {
	b.s.LiveUserNum = &v
	return b
}

func (b *LoginLivingPageMonthlyResponseRecordsBuilder) DailyTime(v string) *LoginLivingPageMonthlyResponseRecordsBuilder {
	b.s.DailyTime = &v
	return b
}

func (b *LoginLivingPageMonthlyResponseRecordsBuilder) AvgDailyTotalOnline(v string) *LoginLivingPageMonthlyResponseRecordsBuilder {
	b.s.AvgDailyTotalOnline = &v
	return b
}

func (b *LoginLivingPageMonthlyResponseRecordsBuilder) MonthlyTime(v string) *LoginLivingPageMonthlyResponseRecordsBuilder {
	b.s.MonthlyTime = &v
	return b
}

func (b *LoginLivingPageMonthlyResponseRecordsBuilder) TotalOnline(v string) *LoginLivingPageMonthlyResponseRecordsBuilder {
	b.s.TotalOnline = &v
	return b
}

func (b *LoginLivingPageMonthlyResponseRecordsBuilder) AveDailyOnlineTime(v string) *LoginLivingPageMonthlyResponseRecordsBuilder {
	b.s.AveDailyOnlineTime = &v
	return b
}

func (b *LoginLivingPageMonthlyResponseRecordsBuilder) AveOnlineTime(v string) *LoginLivingPageMonthlyResponseRecordsBuilder {
	b.s.AveOnlineTime = &v
	return b
}

func (b *LoginLivingPageMonthlyResponseRecordsBuilder) TotalOnlineTime(v string) *LoginLivingPageMonthlyResponseRecordsBuilder {
	b.s.TotalOnlineTime = &v
	return b
}

func (b *LoginLivingPageMonthlyResponseRecordsBuilder) Build() *LoginLivingPageMonthlyResponseRecords {
	return b.s
}


