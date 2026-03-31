// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ConnectLivingPageMonthlyResponseBody struct {

    // 总数据数
	Total *int32 `json:"total,omitempty"`
    // 活跃度数据列表
	Records *[]ConnectLivingPageMonthlyResponseRecords `json:"records,omitempty"`
}

func (s ConnectLivingPageMonthlyResponseBody) String() string {
	return utils.Beautify(s)
}

func (s ConnectLivingPageMonthlyResponseBody) GoString() string {
	return s.String()
}

func (s ConnectLivingPageMonthlyResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ConnectLivingPageMonthlyResponseBody) SetTotal(v int32) *ConnectLivingPageMonthlyResponseBody {
	s.Total = &v
	return s
}

func (s *ConnectLivingPageMonthlyResponseBody) SetRecords(v []ConnectLivingPageMonthlyResponseRecords) *ConnectLivingPageMonthlyResponseBody {
	s.Records = &v
	return s
}


type ConnectLivingPageMonthlyResponseBodyBuilder struct {
	s *ConnectLivingPageMonthlyResponseBody
}

func NewConnectLivingPageMonthlyResponseBodyBuilder() *ConnectLivingPageMonthlyResponseBodyBuilder {
	s := &ConnectLivingPageMonthlyResponseBody{}
	b := &ConnectLivingPageMonthlyResponseBodyBuilder{s: s}
	return b
}

func (b *ConnectLivingPageMonthlyResponseBodyBuilder) Total(v int32) *ConnectLivingPageMonthlyResponseBodyBuilder {
	b.s.Total = &v
	return b
}

func (b *ConnectLivingPageMonthlyResponseBodyBuilder) Records(v []ConnectLivingPageMonthlyResponseRecords) *ConnectLivingPageMonthlyResponseBodyBuilder {
	b.s.Records = &v
	return b
}

func (b *ConnectLivingPageMonthlyResponseBodyBuilder) Build() *ConnectLivingPageMonthlyResponseBody {
	return b.s
}


