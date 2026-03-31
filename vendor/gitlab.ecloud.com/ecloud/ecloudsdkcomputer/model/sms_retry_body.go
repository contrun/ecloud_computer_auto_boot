// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type SmsRetryBody struct {
    position.Body
    // 实例id。示例值：CCA-XXX
	InstanceId *string `json:"instanceId,omitempty"`
}

func (s SmsRetryBody) String() string {
	return utils.Beautify(s)
}

func (s SmsRetryBody) GoString() string {
	return s.String()
}

func (s SmsRetryBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *SmsRetryBody) SetInstanceId(v string) *SmsRetryBody {
	s.InstanceId = &v
	return s
}


type SmsRetryBodyBuilder struct {
	s *SmsRetryBody
}

func NewSmsRetryBodyBuilder() *SmsRetryBodyBuilder {
	s := &SmsRetryBody{}
	b := &SmsRetryBodyBuilder{s: s}
	return b
}

func (b *SmsRetryBodyBuilder) InstanceId(v string) *SmsRetryBodyBuilder {
	b.s.InstanceId = &v
	return b
}

func (b *SmsRetryBodyBuilder) Build() *SmsRetryBody {
	return b.s
}


