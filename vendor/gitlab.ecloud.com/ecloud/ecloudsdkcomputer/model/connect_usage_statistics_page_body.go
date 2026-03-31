// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ConnectUsageStatisticsPageBody struct {
    position.Body
    // 电脑id
	MachineId *string `json:"machineId,omitempty"`
    // 分页大小
	PageSize *int32 `json:"pageSize,omitempty"`
    // 开始时间
	StartTime *string `json:"startTime,omitempty"`
    // 结束时间
	EndTime *string `json:"endTime,omitempty"`
    // 当前页
	Page *int32 `json:"page,omitempty"`
    // 用户名
	UserName *string `json:"userName,omitempty"`
    // 电脑名称
	MachineName *string `json:"machineName,omitempty"`
}

func (s ConnectUsageStatisticsPageBody) String() string {
	return utils.Beautify(s)
}

func (s ConnectUsageStatisticsPageBody) GoString() string {
	return s.String()
}

func (s ConnectUsageStatisticsPageBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ConnectUsageStatisticsPageBody) SetMachineId(v string) *ConnectUsageStatisticsPageBody {
	s.MachineId = &v
	return s
}

func (s *ConnectUsageStatisticsPageBody) SetPageSize(v int32) *ConnectUsageStatisticsPageBody {
	s.PageSize = &v
	return s
}

func (s *ConnectUsageStatisticsPageBody) SetStartTime(v string) *ConnectUsageStatisticsPageBody {
	s.StartTime = &v
	return s
}

func (s *ConnectUsageStatisticsPageBody) SetEndTime(v string) *ConnectUsageStatisticsPageBody {
	s.EndTime = &v
	return s
}

func (s *ConnectUsageStatisticsPageBody) SetPage(v int32) *ConnectUsageStatisticsPageBody {
	s.Page = &v
	return s
}

func (s *ConnectUsageStatisticsPageBody) SetUserName(v string) *ConnectUsageStatisticsPageBody {
	s.UserName = &v
	return s
}

func (s *ConnectUsageStatisticsPageBody) SetMachineName(v string) *ConnectUsageStatisticsPageBody {
	s.MachineName = &v
	return s
}


type ConnectUsageStatisticsPageBodyBuilder struct {
	s *ConnectUsageStatisticsPageBody
}

func NewConnectUsageStatisticsPageBodyBuilder() *ConnectUsageStatisticsPageBodyBuilder {
	s := &ConnectUsageStatisticsPageBody{}
	b := &ConnectUsageStatisticsPageBodyBuilder{s: s}
	return b
}

func (b *ConnectUsageStatisticsPageBodyBuilder) MachineId(v string) *ConnectUsageStatisticsPageBodyBuilder {
	b.s.MachineId = &v
	return b
}

func (b *ConnectUsageStatisticsPageBodyBuilder) PageSize(v int32) *ConnectUsageStatisticsPageBodyBuilder {
	b.s.PageSize = &v
	return b
}

func (b *ConnectUsageStatisticsPageBodyBuilder) StartTime(v string) *ConnectUsageStatisticsPageBodyBuilder {
	b.s.StartTime = &v
	return b
}

func (b *ConnectUsageStatisticsPageBodyBuilder) EndTime(v string) *ConnectUsageStatisticsPageBodyBuilder {
	b.s.EndTime = &v
	return b
}

func (b *ConnectUsageStatisticsPageBodyBuilder) Page(v int32) *ConnectUsageStatisticsPageBodyBuilder {
	b.s.Page = &v
	return b
}

func (b *ConnectUsageStatisticsPageBodyBuilder) UserName(v string) *ConnectUsageStatisticsPageBodyBuilder {
	b.s.UserName = &v
	return b
}

func (b *ConnectUsageStatisticsPageBodyBuilder) MachineName(v string) *ConnectUsageStatisticsPageBodyBuilder {
	b.s.MachineName = &v
	return b
}

func (b *ConnectUsageStatisticsPageBodyBuilder) Build() *ConnectUsageStatisticsPageBody {
	return b.s
}


