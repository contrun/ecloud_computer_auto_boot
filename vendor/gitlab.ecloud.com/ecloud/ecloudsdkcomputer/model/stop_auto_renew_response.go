// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type StopAutoRenewResponse struct {

    // 响应数据
	Data *StopAutoRenewResponseData `json:"data,omitempty"`
    // 请求ID
	RequestId *string `json:"requestId,omitempty"`
    // 响应信息
	RespMsg *string `json:"respMsg,omitempty"`
    // 响应码
	RespCode *string `json:"respCode,omitempty"`
}

func (s StopAutoRenewResponse) String() string {
	return utils.Beautify(s)
}

func (s StopAutoRenewResponse) GoString() string {
	return s.String()
}

func (s StopAutoRenewResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *StopAutoRenewResponse) SetData(v *StopAutoRenewResponseData) *StopAutoRenewResponse {
	s.Data = v
	return s
}

func (s *StopAutoRenewResponse) SetRequestId(v string) *StopAutoRenewResponse {
	s.RequestId = &v
	return s
}

func (s *StopAutoRenewResponse) SetRespMsg(v string) *StopAutoRenewResponse {
	s.RespMsg = &v
	return s
}

func (s *StopAutoRenewResponse) SetRespCode(v string) *StopAutoRenewResponse {
	s.RespCode = &v
	return s
}


type StopAutoRenewResponseBuilder struct {
	s *StopAutoRenewResponse
}

func NewStopAutoRenewResponseBuilder() *StopAutoRenewResponseBuilder {
	s := &StopAutoRenewResponse{}
	b := &StopAutoRenewResponseBuilder{s: s}
	return b
}

func (b *StopAutoRenewResponseBuilder) Data(v *StopAutoRenewResponseData) *StopAutoRenewResponseBuilder {
	b.s.Data = v
	return b
}

func (b *StopAutoRenewResponseBuilder) RequestId(v string) *StopAutoRenewResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *StopAutoRenewResponseBuilder) RespMsg(v string) *StopAutoRenewResponseBuilder {
	b.s.RespMsg = &v
	return b
}

func (b *StopAutoRenewResponseBuilder) RespCode(v string) *StopAutoRenewResponseBuilder {
	b.s.RespCode = &v
	return b
}

func (b *StopAutoRenewResponseBuilder) Build() *StopAutoRenewResponse {
	return b.s
}


