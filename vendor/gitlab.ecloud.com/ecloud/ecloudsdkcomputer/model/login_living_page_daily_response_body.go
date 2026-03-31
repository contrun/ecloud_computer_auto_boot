// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type LoginLivingPageDailyResponseBody struct {

    // 总数据数
	Total *int32 `json:"total,omitempty"`
    // 活跃度数据列表
	Records *[]LoginLivingPageDailyResponseRecords `json:"records,omitempty"`
}

func (s LoginLivingPageDailyResponseBody) String() string {
	return utils.Beautify(s)
}

func (s LoginLivingPageDailyResponseBody) GoString() string {
	return s.String()
}

func (s LoginLivingPageDailyResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *LoginLivingPageDailyResponseBody) SetTotal(v int32) *LoginLivingPageDailyResponseBody {
	s.Total = &v
	return s
}

func (s *LoginLivingPageDailyResponseBody) SetRecords(v []LoginLivingPageDailyResponseRecords) *LoginLivingPageDailyResponseBody {
	s.Records = &v
	return s
}


type LoginLivingPageDailyResponseBodyBuilder struct {
	s *LoginLivingPageDailyResponseBody
}

func NewLoginLivingPageDailyResponseBodyBuilder() *LoginLivingPageDailyResponseBodyBuilder {
	s := &LoginLivingPageDailyResponseBody{}
	b := &LoginLivingPageDailyResponseBodyBuilder{s: s}
	return b
}

func (b *LoginLivingPageDailyResponseBodyBuilder) Total(v int32) *LoginLivingPageDailyResponseBodyBuilder {
	b.s.Total = &v
	return b
}

func (b *LoginLivingPageDailyResponseBodyBuilder) Records(v []LoginLivingPageDailyResponseRecords) *LoginLivingPageDailyResponseBodyBuilder {
	b.s.Records = &v
	return b
}

func (b *LoginLivingPageDailyResponseBodyBuilder) Build() *LoginLivingPageDailyResponseBody {
	return b.s
}


