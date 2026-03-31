// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type StartAutoRenewNatResponse struct {

    // 响应数据
	Data *StartAutoRenewNatResponseData `json:"data,omitempty"`
    // 请求ID
	RequestId *string `json:"requestId,omitempty"`
    // 响应信息
	RespMsg *string `json:"respMsg,omitempty"`
    // 响应编码
	RespCode *string `json:"respCode,omitempty"`
}

func (s StartAutoRenewNatResponse) String() string {
	return utils.Beautify(s)
}

func (s StartAutoRenewNatResponse) GoString() string {
	return s.String()
}

func (s StartAutoRenewNatResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *StartAutoRenewNatResponse) SetData(v *StartAutoRenewNatResponseData) *StartAutoRenewNatResponse {
	s.Data = v
	return s
}

func (s *StartAutoRenewNatResponse) SetRequestId(v string) *StartAutoRenewNatResponse {
	s.RequestId = &v
	return s
}

func (s *StartAutoRenewNatResponse) SetRespMsg(v string) *StartAutoRenewNatResponse {
	s.RespMsg = &v
	return s
}

func (s *StartAutoRenewNatResponse) SetRespCode(v string) *StartAutoRenewNatResponse {
	s.RespCode = &v
	return s
}


type StartAutoRenewNatResponseBuilder struct {
	s *StartAutoRenewNatResponse
}

func NewStartAutoRenewNatResponseBuilder() *StartAutoRenewNatResponseBuilder {
	s := &StartAutoRenewNatResponse{}
	b := &StartAutoRenewNatResponseBuilder{s: s}
	return b
}

func (b *StartAutoRenewNatResponseBuilder) Data(v *StartAutoRenewNatResponseData) *StartAutoRenewNatResponseBuilder {
	b.s.Data = v
	return b
}

func (b *StartAutoRenewNatResponseBuilder) RequestId(v string) *StartAutoRenewNatResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *StartAutoRenewNatResponseBuilder) RespMsg(v string) *StartAutoRenewNatResponseBuilder {
	b.s.RespMsg = &v
	return b
}

func (b *StartAutoRenewNatResponseBuilder) RespCode(v string) *StartAutoRenewNatResponseBuilder {
	b.s.RespCode = &v
	return b
}

func (b *StartAutoRenewNatResponseBuilder) Build() *StartAutoRenewNatResponse {
	return b.s
}


