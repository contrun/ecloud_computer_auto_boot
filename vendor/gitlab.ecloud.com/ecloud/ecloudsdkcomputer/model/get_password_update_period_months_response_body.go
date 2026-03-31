// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetPasswordUpdatePeriodMonthsResponseBody struct {

    // 密码需要几个月更新一次，若为-1，则密码不设置过期时间
	PasswordUpdatePeriodMonths *int32 `json:"passwordUpdatePeriodMonths,omitempty"`
}

func (s GetPasswordUpdatePeriodMonthsResponseBody) String() string {
	return utils.Beautify(s)
}

func (s GetPasswordUpdatePeriodMonthsResponseBody) GoString() string {
	return s.String()
}

func (s GetPasswordUpdatePeriodMonthsResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetPasswordUpdatePeriodMonthsResponseBody) SetPasswordUpdatePeriodMonths(v int32) *GetPasswordUpdatePeriodMonthsResponseBody {
	s.PasswordUpdatePeriodMonths = &v
	return s
}


type GetPasswordUpdatePeriodMonthsResponseBodyBuilder struct {
	s *GetPasswordUpdatePeriodMonthsResponseBody
}

func NewGetPasswordUpdatePeriodMonthsResponseBodyBuilder() *GetPasswordUpdatePeriodMonthsResponseBodyBuilder {
	s := &GetPasswordUpdatePeriodMonthsResponseBody{}
	b := &GetPasswordUpdatePeriodMonthsResponseBodyBuilder{s: s}
	return b
}

func (b *GetPasswordUpdatePeriodMonthsResponseBodyBuilder) PasswordUpdatePeriodMonths(v int32) *GetPasswordUpdatePeriodMonthsResponseBodyBuilder {
	b.s.PasswordUpdatePeriodMonths = &v
	return b
}

func (b *GetPasswordUpdatePeriodMonthsResponseBodyBuilder) Build() *GetPasswordUpdatePeriodMonthsResponseBody {
	return b.s
}


