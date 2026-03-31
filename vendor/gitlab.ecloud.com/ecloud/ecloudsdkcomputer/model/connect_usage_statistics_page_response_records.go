// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ConnectUsageStatisticsPageResponseRecords struct {

    // 平均在线时长
	AvgOnlineTime *string `json:"avgOnlineTime,omitempty"`
    // 在线总时长
	OnlineTotalTime *string `json:"onlineTotalTime,omitempty"`
    // 所属渠道
	OrderSourceName *string `json:"orderSourceName,omitempty"`
    // 所属企业
	CompanyName *string `json:"companyName,omitempty"`
    // 用户名/授权用户
	UserName *string `json:"userName,omitempty"`
    // 连接次数
	ConnectCount *string `json:"connectCount,omitempty"`
    // 用户id
	UserId *string `json:"userId,omitempty"`
    // 桌面名称
	MachineName *string `json:"machineName,omitempty"`
    // 最近一次连接时间
	LatestOnlineTime *int64 `json:"latestOnlineTime,omitempty"`
    // 桌面Id
	MachineId *string `json:"machineId,omitempty"`
    // 所属渠道code
	OrderSourceCode *string `json:"orderSourceCode,omitempty"`
    // 租户表主键Id
	TenantId *int64 `json:"tenantId,omitempty"`
    // 租户名
	MopUserName *string `json:"mopUserName,omitempty"`
    // 最长在线时长
	LongestOnlineTime *string `json:"longestOnlineTime,omitempty"`
}

func (s ConnectUsageStatisticsPageResponseRecords) String() string {
	return utils.Beautify(s)
}

func (s ConnectUsageStatisticsPageResponseRecords) GoString() string {
	return s.String()
}

func (s ConnectUsageStatisticsPageResponseRecords) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ConnectUsageStatisticsPageResponseRecords) SetAvgOnlineTime(v string) *ConnectUsageStatisticsPageResponseRecords {
	s.AvgOnlineTime = &v
	return s
}

func (s *ConnectUsageStatisticsPageResponseRecords) SetOnlineTotalTime(v string) *ConnectUsageStatisticsPageResponseRecords {
	s.OnlineTotalTime = &v
	return s
}

func (s *ConnectUsageStatisticsPageResponseRecords) SetOrderSourceName(v string) *ConnectUsageStatisticsPageResponseRecords {
	s.OrderSourceName = &v
	return s
}

func (s *ConnectUsageStatisticsPageResponseRecords) SetCompanyName(v string) *ConnectUsageStatisticsPageResponseRecords {
	s.CompanyName = &v
	return s
}

func (s *ConnectUsageStatisticsPageResponseRecords) SetUserName(v string) *ConnectUsageStatisticsPageResponseRecords {
	s.UserName = &v
	return s
}

func (s *ConnectUsageStatisticsPageResponseRecords) SetConnectCount(v string) *ConnectUsageStatisticsPageResponseRecords {
	s.ConnectCount = &v
	return s
}

func (s *ConnectUsageStatisticsPageResponseRecords) SetUserId(v string) *ConnectUsageStatisticsPageResponseRecords {
	s.UserId = &v
	return s
}

func (s *ConnectUsageStatisticsPageResponseRecords) SetMachineName(v string) *ConnectUsageStatisticsPageResponseRecords {
	s.MachineName = &v
	return s
}

func (s *ConnectUsageStatisticsPageResponseRecords) SetLatestOnlineTime(v int64) *ConnectUsageStatisticsPageResponseRecords {
	s.LatestOnlineTime = &v
	return s
}

func (s *ConnectUsageStatisticsPageResponseRecords) SetMachineId(v string) *ConnectUsageStatisticsPageResponseRecords {
	s.MachineId = &v
	return s
}

func (s *ConnectUsageStatisticsPageResponseRecords) SetOrderSourceCode(v string) *ConnectUsageStatisticsPageResponseRecords {
	s.OrderSourceCode = &v
	return s
}

func (s *ConnectUsageStatisticsPageResponseRecords) SetTenantId(v int64) *ConnectUsageStatisticsPageResponseRecords {
	s.TenantId = &v
	return s
}

func (s *ConnectUsageStatisticsPageResponseRecords) SetMopUserName(v string) *ConnectUsageStatisticsPageResponseRecords {
	s.MopUserName = &v
	return s
}

func (s *ConnectUsageStatisticsPageResponseRecords) SetLongestOnlineTime(v string) *ConnectUsageStatisticsPageResponseRecords {
	s.LongestOnlineTime = &v
	return s
}


type ConnectUsageStatisticsPageResponseRecordsBuilder struct {
	s *ConnectUsageStatisticsPageResponseRecords
}

func NewConnectUsageStatisticsPageResponseRecordsBuilder() *ConnectUsageStatisticsPageResponseRecordsBuilder {
	s := &ConnectUsageStatisticsPageResponseRecords{}
	b := &ConnectUsageStatisticsPageResponseRecordsBuilder{s: s}
	return b
}

func (b *ConnectUsageStatisticsPageResponseRecordsBuilder) AvgOnlineTime(v string) *ConnectUsageStatisticsPageResponseRecordsBuilder {
	b.s.AvgOnlineTime = &v
	return b
}

func (b *ConnectUsageStatisticsPageResponseRecordsBuilder) OnlineTotalTime(v string) *ConnectUsageStatisticsPageResponseRecordsBuilder {
	b.s.OnlineTotalTime = &v
	return b
}

func (b *ConnectUsageStatisticsPageResponseRecordsBuilder) OrderSourceName(v string) *ConnectUsageStatisticsPageResponseRecordsBuilder {
	b.s.OrderSourceName = &v
	return b
}

func (b *ConnectUsageStatisticsPageResponseRecordsBuilder) CompanyName(v string) *ConnectUsageStatisticsPageResponseRecordsBuilder {
	b.s.CompanyName = &v
	return b
}

func (b *ConnectUsageStatisticsPageResponseRecordsBuilder) UserName(v string) *ConnectUsageStatisticsPageResponseRecordsBuilder {
	b.s.UserName = &v
	return b
}

func (b *ConnectUsageStatisticsPageResponseRecordsBuilder) ConnectCount(v string) *ConnectUsageStatisticsPageResponseRecordsBuilder {
	b.s.ConnectCount = &v
	return b
}

func (b *ConnectUsageStatisticsPageResponseRecordsBuilder) UserId(v string) *ConnectUsageStatisticsPageResponseRecordsBuilder {
	b.s.UserId = &v
	return b
}

func (b *ConnectUsageStatisticsPageResponseRecordsBuilder) MachineName(v string) *ConnectUsageStatisticsPageResponseRecordsBuilder {
	b.s.MachineName = &v
	return b
}

func (b *ConnectUsageStatisticsPageResponseRecordsBuilder) LatestOnlineTime(v int64) *ConnectUsageStatisticsPageResponseRecordsBuilder {
	b.s.LatestOnlineTime = &v
	return b
}

func (b *ConnectUsageStatisticsPageResponseRecordsBuilder) MachineId(v string) *ConnectUsageStatisticsPageResponseRecordsBuilder {
	b.s.MachineId = &v
	return b
}

func (b *ConnectUsageStatisticsPageResponseRecordsBuilder) OrderSourceCode(v string) *ConnectUsageStatisticsPageResponseRecordsBuilder {
	b.s.OrderSourceCode = &v
	return b
}

func (b *ConnectUsageStatisticsPageResponseRecordsBuilder) TenantId(v int64) *ConnectUsageStatisticsPageResponseRecordsBuilder {
	b.s.TenantId = &v
	return b
}

func (b *ConnectUsageStatisticsPageResponseRecordsBuilder) MopUserName(v string) *ConnectUsageStatisticsPageResponseRecordsBuilder {
	b.s.MopUserName = &v
	return b
}

func (b *ConnectUsageStatisticsPageResponseRecordsBuilder) LongestOnlineTime(v string) *ConnectUsageStatisticsPageResponseRecordsBuilder {
	b.s.LongestOnlineTime = &v
	return b
}

func (b *ConnectUsageStatisticsPageResponseRecordsBuilder) Build() *ConnectUsageStatisticsPageResponseRecords {
	return b.s
}


