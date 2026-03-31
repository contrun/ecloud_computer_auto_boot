// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ConnectLivingPageDailyResponseBody struct {

    // 总数据数
	Total *int32 `json:"total,omitempty"`
    // 活跃度数据列表
	Records *[]ConnectLivingPageDailyResponseRecords `json:"records,omitempty"`
}

func (s ConnectLivingPageDailyResponseBody) String() string {
	return utils.Beautify(s)
}

func (s ConnectLivingPageDailyResponseBody) GoString() string {
	return s.String()
}

func (s ConnectLivingPageDailyResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ConnectLivingPageDailyResponseBody) SetTotal(v int32) *ConnectLivingPageDailyResponseBody {
	s.Total = &v
	return s
}

func (s *ConnectLivingPageDailyResponseBody) SetRecords(v []ConnectLivingPageDailyResponseRecords) *ConnectLivingPageDailyResponseBody {
	s.Records = &v
	return s
}


type ConnectLivingPageDailyResponseBodyBuilder struct {
	s *ConnectLivingPageDailyResponseBody
}

func NewConnectLivingPageDailyResponseBodyBuilder() *ConnectLivingPageDailyResponseBodyBuilder {
	s := &ConnectLivingPageDailyResponseBody{}
	b := &ConnectLivingPageDailyResponseBodyBuilder{s: s}
	return b
}

func (b *ConnectLivingPageDailyResponseBodyBuilder) Total(v int32) *ConnectLivingPageDailyResponseBodyBuilder {
	b.s.Total = &v
	return b
}

func (b *ConnectLivingPageDailyResponseBodyBuilder) Records(v []ConnectLivingPageDailyResponseRecords) *ConnectLivingPageDailyResponseBodyBuilder {
	b.s.Records = &v
	return b
}

func (b *ConnectLivingPageDailyResponseBodyBuilder) Build() *ConnectLivingPageDailyResponseBody {
	return b.s
}


