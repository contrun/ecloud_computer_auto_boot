// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type LoginUsageStatisticsPageResponseBody struct {

    // 总数据数
	Total *int32 `json:"total,omitempty"`
    // 使用统计数据列表
	Records *[]LoginUsageStatisticsPageResponseRecords `json:"records,omitempty"`
}

func (s LoginUsageStatisticsPageResponseBody) String() string {
	return utils.Beautify(s)
}

func (s LoginUsageStatisticsPageResponseBody) GoString() string {
	return s.String()
}

func (s LoginUsageStatisticsPageResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *LoginUsageStatisticsPageResponseBody) SetTotal(v int32) *LoginUsageStatisticsPageResponseBody {
	s.Total = &v
	return s
}

func (s *LoginUsageStatisticsPageResponseBody) SetRecords(v []LoginUsageStatisticsPageResponseRecords) *LoginUsageStatisticsPageResponseBody {
	s.Records = &v
	return s
}


type LoginUsageStatisticsPageResponseBodyBuilder struct {
	s *LoginUsageStatisticsPageResponseBody
}

func NewLoginUsageStatisticsPageResponseBodyBuilder() *LoginUsageStatisticsPageResponseBodyBuilder {
	s := &LoginUsageStatisticsPageResponseBody{}
	b := &LoginUsageStatisticsPageResponseBodyBuilder{s: s}
	return b
}

func (b *LoginUsageStatisticsPageResponseBodyBuilder) Total(v int32) *LoginUsageStatisticsPageResponseBodyBuilder {
	b.s.Total = &v
	return b
}

func (b *LoginUsageStatisticsPageResponseBodyBuilder) Records(v []LoginUsageStatisticsPageResponseRecords) *LoginUsageStatisticsPageResponseBodyBuilder {
	b.s.Records = &v
	return b
}

func (b *LoginUsageStatisticsPageResponseBodyBuilder) Build() *LoginUsageStatisticsPageResponseBody {
	return b.s
}


