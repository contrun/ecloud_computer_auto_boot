// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ConnectUsageStatisticsPageResponseBody struct {

    // 总数据数
	Total *int32 `json:"total,omitempty"`
    // 使用统计数据列表
	Records *[]ConnectUsageStatisticsPageResponseRecords `json:"records,omitempty"`
}

func (s ConnectUsageStatisticsPageResponseBody) String() string {
	return utils.Beautify(s)
}

func (s ConnectUsageStatisticsPageResponseBody) GoString() string {
	return s.String()
}

func (s ConnectUsageStatisticsPageResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ConnectUsageStatisticsPageResponseBody) SetTotal(v int32) *ConnectUsageStatisticsPageResponseBody {
	s.Total = &v
	return s
}

func (s *ConnectUsageStatisticsPageResponseBody) SetRecords(v []ConnectUsageStatisticsPageResponseRecords) *ConnectUsageStatisticsPageResponseBody {
	s.Records = &v
	return s
}


type ConnectUsageStatisticsPageResponseBodyBuilder struct {
	s *ConnectUsageStatisticsPageResponseBody
}

func NewConnectUsageStatisticsPageResponseBodyBuilder() *ConnectUsageStatisticsPageResponseBodyBuilder {
	s := &ConnectUsageStatisticsPageResponseBody{}
	b := &ConnectUsageStatisticsPageResponseBodyBuilder{s: s}
	return b
}

func (b *ConnectUsageStatisticsPageResponseBodyBuilder) Total(v int32) *ConnectUsageStatisticsPageResponseBodyBuilder {
	b.s.Total = &v
	return b
}

func (b *ConnectUsageStatisticsPageResponseBodyBuilder) Records(v []ConnectUsageStatisticsPageResponseRecords) *ConnectUsageStatisticsPageResponseBodyBuilder {
	b.s.Records = &v
	return b
}

func (b *ConnectUsageStatisticsPageResponseBodyBuilder) Build() *ConnectUsageStatisticsPageResponseBody {
	return b.s
}


