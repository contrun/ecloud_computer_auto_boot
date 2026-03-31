// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ConnectLivingPageMonthlyResponseRecords struct {

    // 日期
	DailyTime *string `json:"dailyTime,omitempty"`
    // 活跃桌面数
	LiveMachineNumber *int32 `json:"liveMachineNumber,omitempty"`
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

func (s ConnectLivingPageMonthlyResponseRecords) String() string {
	return utils.Beautify(s)
}

func (s ConnectLivingPageMonthlyResponseRecords) GoString() string {
	return s.String()
}

func (s ConnectLivingPageMonthlyResponseRecords) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ConnectLivingPageMonthlyResponseRecords) SetDailyTime(v string) *ConnectLivingPageMonthlyResponseRecords {
	s.DailyTime = &v
	return s
}

func (s *ConnectLivingPageMonthlyResponseRecords) SetLiveMachineNumber(v int32) *ConnectLivingPageMonthlyResponseRecords {
	s.LiveMachineNumber = &v
	return s
}

func (s *ConnectLivingPageMonthlyResponseRecords) SetAvgDailyTotalOnline(v string) *ConnectLivingPageMonthlyResponseRecords {
	s.AvgDailyTotalOnline = &v
	return s
}

func (s *ConnectLivingPageMonthlyResponseRecords) SetMonthlyTime(v string) *ConnectLivingPageMonthlyResponseRecords {
	s.MonthlyTime = &v
	return s
}

func (s *ConnectLivingPageMonthlyResponseRecords) SetTotalOnline(v string) *ConnectLivingPageMonthlyResponseRecords {
	s.TotalOnline = &v
	return s
}

func (s *ConnectLivingPageMonthlyResponseRecords) SetAveDailyOnlineTime(v string) *ConnectLivingPageMonthlyResponseRecords {
	s.AveDailyOnlineTime = &v
	return s
}

func (s *ConnectLivingPageMonthlyResponseRecords) SetAveOnlineTime(v string) *ConnectLivingPageMonthlyResponseRecords {
	s.AveOnlineTime = &v
	return s
}

func (s *ConnectLivingPageMonthlyResponseRecords) SetTotalOnlineTime(v string) *ConnectLivingPageMonthlyResponseRecords {
	s.TotalOnlineTime = &v
	return s
}


type ConnectLivingPageMonthlyResponseRecordsBuilder struct {
	s *ConnectLivingPageMonthlyResponseRecords
}

func NewConnectLivingPageMonthlyResponseRecordsBuilder() *ConnectLivingPageMonthlyResponseRecordsBuilder {
	s := &ConnectLivingPageMonthlyResponseRecords{}
	b := &ConnectLivingPageMonthlyResponseRecordsBuilder{s: s}
	return b
}

func (b *ConnectLivingPageMonthlyResponseRecordsBuilder) DailyTime(v string) *ConnectLivingPageMonthlyResponseRecordsBuilder {
	b.s.DailyTime = &v
	return b
}

func (b *ConnectLivingPageMonthlyResponseRecordsBuilder) LiveMachineNumber(v int32) *ConnectLivingPageMonthlyResponseRecordsBuilder {
	b.s.LiveMachineNumber = &v
	return b
}

func (b *ConnectLivingPageMonthlyResponseRecordsBuilder) AvgDailyTotalOnline(v string) *ConnectLivingPageMonthlyResponseRecordsBuilder {
	b.s.AvgDailyTotalOnline = &v
	return b
}

func (b *ConnectLivingPageMonthlyResponseRecordsBuilder) MonthlyTime(v string) *ConnectLivingPageMonthlyResponseRecordsBuilder {
	b.s.MonthlyTime = &v
	return b
}

func (b *ConnectLivingPageMonthlyResponseRecordsBuilder) TotalOnline(v string) *ConnectLivingPageMonthlyResponseRecordsBuilder {
	b.s.TotalOnline = &v
	return b
}

func (b *ConnectLivingPageMonthlyResponseRecordsBuilder) AveDailyOnlineTime(v string) *ConnectLivingPageMonthlyResponseRecordsBuilder {
	b.s.AveDailyOnlineTime = &v
	return b
}

func (b *ConnectLivingPageMonthlyResponseRecordsBuilder) AveOnlineTime(v string) *ConnectLivingPageMonthlyResponseRecordsBuilder {
	b.s.AveOnlineTime = &v
	return b
}

func (b *ConnectLivingPageMonthlyResponseRecordsBuilder) TotalOnlineTime(v string) *ConnectLivingPageMonthlyResponseRecordsBuilder {
	b.s.TotalOnlineTime = &v
	return b
}

func (b *ConnectLivingPageMonthlyResponseRecordsBuilder) Build() *ConnectLivingPageMonthlyResponseRecords {
	return b.s
}


