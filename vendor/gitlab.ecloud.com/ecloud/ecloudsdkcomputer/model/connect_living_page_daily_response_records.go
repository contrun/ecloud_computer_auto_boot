// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ConnectLivingPageDailyResponseRecords struct {

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

func (s ConnectLivingPageDailyResponseRecords) String() string {
	return utils.Beautify(s)
}

func (s ConnectLivingPageDailyResponseRecords) GoString() string {
	return s.String()
}

func (s ConnectLivingPageDailyResponseRecords) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ConnectLivingPageDailyResponseRecords) SetDailyTime(v string) *ConnectLivingPageDailyResponseRecords {
	s.DailyTime = &v
	return s
}

func (s *ConnectLivingPageDailyResponseRecords) SetLiveMachineNumber(v int32) *ConnectLivingPageDailyResponseRecords {
	s.LiveMachineNumber = &v
	return s
}

func (s *ConnectLivingPageDailyResponseRecords) SetAvgDailyTotalOnline(v string) *ConnectLivingPageDailyResponseRecords {
	s.AvgDailyTotalOnline = &v
	return s
}

func (s *ConnectLivingPageDailyResponseRecords) SetMonthlyTime(v string) *ConnectLivingPageDailyResponseRecords {
	s.MonthlyTime = &v
	return s
}

func (s *ConnectLivingPageDailyResponseRecords) SetTotalOnline(v string) *ConnectLivingPageDailyResponseRecords {
	s.TotalOnline = &v
	return s
}

func (s *ConnectLivingPageDailyResponseRecords) SetAveDailyOnlineTime(v string) *ConnectLivingPageDailyResponseRecords {
	s.AveDailyOnlineTime = &v
	return s
}

func (s *ConnectLivingPageDailyResponseRecords) SetAveOnlineTime(v string) *ConnectLivingPageDailyResponseRecords {
	s.AveOnlineTime = &v
	return s
}

func (s *ConnectLivingPageDailyResponseRecords) SetTotalOnlineTime(v string) *ConnectLivingPageDailyResponseRecords {
	s.TotalOnlineTime = &v
	return s
}


type ConnectLivingPageDailyResponseRecordsBuilder struct {
	s *ConnectLivingPageDailyResponseRecords
}

func NewConnectLivingPageDailyResponseRecordsBuilder() *ConnectLivingPageDailyResponseRecordsBuilder {
	s := &ConnectLivingPageDailyResponseRecords{}
	b := &ConnectLivingPageDailyResponseRecordsBuilder{s: s}
	return b
}

func (b *ConnectLivingPageDailyResponseRecordsBuilder) DailyTime(v string) *ConnectLivingPageDailyResponseRecordsBuilder {
	b.s.DailyTime = &v
	return b
}

func (b *ConnectLivingPageDailyResponseRecordsBuilder) LiveMachineNumber(v int32) *ConnectLivingPageDailyResponseRecordsBuilder {
	b.s.LiveMachineNumber = &v
	return b
}

func (b *ConnectLivingPageDailyResponseRecordsBuilder) AvgDailyTotalOnline(v string) *ConnectLivingPageDailyResponseRecordsBuilder {
	b.s.AvgDailyTotalOnline = &v
	return b
}

func (b *ConnectLivingPageDailyResponseRecordsBuilder) MonthlyTime(v string) *ConnectLivingPageDailyResponseRecordsBuilder {
	b.s.MonthlyTime = &v
	return b
}

func (b *ConnectLivingPageDailyResponseRecordsBuilder) TotalOnline(v string) *ConnectLivingPageDailyResponseRecordsBuilder {
	b.s.TotalOnline = &v
	return b
}

func (b *ConnectLivingPageDailyResponseRecordsBuilder) AveDailyOnlineTime(v string) *ConnectLivingPageDailyResponseRecordsBuilder {
	b.s.AveDailyOnlineTime = &v
	return b
}

func (b *ConnectLivingPageDailyResponseRecordsBuilder) AveOnlineTime(v string) *ConnectLivingPageDailyResponseRecordsBuilder {
	b.s.AveOnlineTime = &v
	return b
}

func (b *ConnectLivingPageDailyResponseRecordsBuilder) TotalOnlineTime(v string) *ConnectLivingPageDailyResponseRecordsBuilder {
	b.s.TotalOnlineTime = &v
	return b
}

func (b *ConnectLivingPageDailyResponseRecordsBuilder) Build() *ConnectLivingPageDailyResponseRecords {
	return b.s
}


