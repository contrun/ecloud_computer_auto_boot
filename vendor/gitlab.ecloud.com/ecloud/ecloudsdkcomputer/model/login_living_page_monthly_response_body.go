// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type LoginLivingPageMonthlyResponseBody struct {

    // 总数据数
	Total *int32 `json:"total,omitempty"`
    // 活跃度数据列表
	Records *[]LoginLivingPageMonthlyResponseRecords `json:"records,omitempty"`
}

func (s LoginLivingPageMonthlyResponseBody) String() string {
	return utils.Beautify(s)
}

func (s LoginLivingPageMonthlyResponseBody) GoString() string {
	return s.String()
}

func (s LoginLivingPageMonthlyResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *LoginLivingPageMonthlyResponseBody) SetTotal(v int32) *LoginLivingPageMonthlyResponseBody {
	s.Total = &v
	return s
}

func (s *LoginLivingPageMonthlyResponseBody) SetRecords(v []LoginLivingPageMonthlyResponseRecords) *LoginLivingPageMonthlyResponseBody {
	s.Records = &v
	return s
}


type LoginLivingPageMonthlyResponseBodyBuilder struct {
	s *LoginLivingPageMonthlyResponseBody
}

func NewLoginLivingPageMonthlyResponseBodyBuilder() *LoginLivingPageMonthlyResponseBodyBuilder {
	s := &LoginLivingPageMonthlyResponseBody{}
	b := &LoginLivingPageMonthlyResponseBodyBuilder{s: s}
	return b
}

func (b *LoginLivingPageMonthlyResponseBodyBuilder) Total(v int32) *LoginLivingPageMonthlyResponseBodyBuilder {
	b.s.Total = &v
	return b
}

func (b *LoginLivingPageMonthlyResponseBodyBuilder) Records(v []LoginLivingPageMonthlyResponseRecords) *LoginLivingPageMonthlyResponseBodyBuilder {
	b.s.Records = &v
	return b
}

func (b *LoginLivingPageMonthlyResponseBodyBuilder) Build() *LoginLivingPageMonthlyResponseBody {
	return b.s
}


