// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type SmsRetryRequest struct {


	SmsRetryBody *SmsRetryBody `json:"smsRetryBody,omitempty"`
}

func (s SmsRetryRequest) String() string {
	return utils.Beautify(s)
}

func (s SmsRetryRequest) GoString() string {
	return s.String()
}

func (s SmsRetryRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *SmsRetryRequest) SetSmsRetryBody(v *SmsRetryBody) *SmsRetryRequest {
	s.SmsRetryBody = v
	return s
}


type SmsRetryRequestBuilder struct {
	s *SmsRetryRequest
}

func NewSmsRetryRequestBuilder() *SmsRetryRequestBuilder {
	s := &SmsRetryRequest{}
	b := &SmsRetryRequestBuilder{s: s}
	return b
}

func (b *SmsRetryRequestBuilder) SmsRetryBody(v *SmsRetryBody) *SmsRetryRequestBuilder {
	b.s.SmsRetryBody = v
	return b
}

func (b *SmsRetryRequestBuilder) Build() *SmsRetryRequest {
	return b.s
}


