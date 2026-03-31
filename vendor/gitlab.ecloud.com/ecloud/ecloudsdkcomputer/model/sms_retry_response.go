// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type SmsRetryResponse struct {

    // 响应数据
	Data *string `json:"data,omitempty"`
    // 请求ID
	RequestId *string `json:"requestId,omitempty"`
    // 响应信息
	RespMsg *string `json:"respMsg,omitempty"`
    // 响应码
	RespCode *string `json:"respCode,omitempty"`
}

func (s SmsRetryResponse) String() string {
	return utils.Beautify(s)
}

func (s SmsRetryResponse) GoString() string {
	return s.String()
}

func (s SmsRetryResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *SmsRetryResponse) SetData(v string) *SmsRetryResponse {
	s.Data = &v
	return s
}

func (s *SmsRetryResponse) SetRequestId(v string) *SmsRetryResponse {
	s.RequestId = &v
	return s
}

func (s *SmsRetryResponse) SetRespMsg(v string) *SmsRetryResponse {
	s.RespMsg = &v
	return s
}

func (s *SmsRetryResponse) SetRespCode(v string) *SmsRetryResponse {
	s.RespCode = &v
	return s
}


type SmsRetryResponseBuilder struct {
	s *SmsRetryResponse
}

func NewSmsRetryResponseBuilder() *SmsRetryResponseBuilder {
	s := &SmsRetryResponse{}
	b := &SmsRetryResponseBuilder{s: s}
	return b
}

func (b *SmsRetryResponseBuilder) Data(v string) *SmsRetryResponseBuilder {
	b.s.Data = &v
	return b
}

func (b *SmsRetryResponseBuilder) RequestId(v string) *SmsRetryResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *SmsRetryResponseBuilder) RespMsg(v string) *SmsRetryResponseBuilder {
	b.s.RespMsg = &v
	return b
}

func (b *SmsRetryResponseBuilder) RespCode(v string) *SmsRetryResponseBuilder {
	b.s.RespCode = &v
	return b
}

func (b *SmsRetryResponseBuilder) Build() *SmsRetryResponse {
	return b.s
}


