// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type SetEnforcePasswordPeriodMonthRequest struct {


	SetEnforcePasswordPeriodMonthBody *SetEnforcePasswordPeriodMonthBody `json:"setEnforcePasswordPeriodMonthBody,omitempty"`
}

func (s SetEnforcePasswordPeriodMonthRequest) String() string {
	return utils.Beautify(s)
}

func (s SetEnforcePasswordPeriodMonthRequest) GoString() string {
	return s.String()
}

func (s SetEnforcePasswordPeriodMonthRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *SetEnforcePasswordPeriodMonthRequest) SetSetEnforcePasswordPeriodMonthBody(v *SetEnforcePasswordPeriodMonthBody) *SetEnforcePasswordPeriodMonthRequest {
	s.SetEnforcePasswordPeriodMonthBody = v
	return s
}


type SetEnforcePasswordPeriodMonthRequestBuilder struct {
	s *SetEnforcePasswordPeriodMonthRequest
}

func NewSetEnforcePasswordPeriodMonthRequestBuilder() *SetEnforcePasswordPeriodMonthRequestBuilder {
	s := &SetEnforcePasswordPeriodMonthRequest{}
	b := &SetEnforcePasswordPeriodMonthRequestBuilder{s: s}
	return b
}

func (b *SetEnforcePasswordPeriodMonthRequestBuilder) SetEnforcePasswordPeriodMonthBody(v *SetEnforcePasswordPeriodMonthBody) *SetEnforcePasswordPeriodMonthRequestBuilder {
	b.s.SetEnforcePasswordPeriodMonthBody = v
	return b
}

func (b *SetEnforcePasswordPeriodMonthRequestBuilder) Build() *SetEnforcePasswordPeriodMonthRequest {
	return b.s
}


