// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type StartAutoRenewResponse struct {

    // 响应数据
	Data *StartAutoRenewResponseData `json:"data,omitempty"`
    // 请求ID
	RequestId *string `json:"requestId,omitempty"`
    // 响应信息
	RespMsg *string `json:"respMsg,omitempty"`
    // 响应码
	RespCode *string `json:"respCode,omitempty"`
}

func (s StartAutoRenewResponse) String() string {
	return utils.Beautify(s)
}

func (s StartAutoRenewResponse) GoString() string {
	return s.String()
}

func (s StartAutoRenewResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *StartAutoRenewResponse) SetData(v *StartAutoRenewResponseData) *StartAutoRenewResponse {
	s.Data = v
	return s
}

func (s *StartAutoRenewResponse) SetRequestId(v string) *StartAutoRenewResponse {
	s.RequestId = &v
	return s
}

func (s *StartAutoRenewResponse) SetRespMsg(v string) *StartAutoRenewResponse {
	s.RespMsg = &v
	return s
}

func (s *StartAutoRenewResponse) SetRespCode(v string) *StartAutoRenewResponse {
	s.RespCode = &v
	return s
}


type StartAutoRenewResponseBuilder struct {
	s *StartAutoRenewResponse
}

func NewStartAutoRenewResponseBuilder() *StartAutoRenewResponseBuilder {
	s := &StartAutoRenewResponse{}
	b := &StartAutoRenewResponseBuilder{s: s}
	return b
}

func (b *StartAutoRenewResponseBuilder) Data(v *StartAutoRenewResponseData) *StartAutoRenewResponseBuilder {
	b.s.Data = v
	return b
}

func (b *StartAutoRenewResponseBuilder) RequestId(v string) *StartAutoRenewResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *StartAutoRenewResponseBuilder) RespMsg(v string) *StartAutoRenewResponseBuilder {
	b.s.RespMsg = &v
	return b
}

func (b *StartAutoRenewResponseBuilder) RespCode(v string) *StartAutoRenewResponseBuilder {
	b.s.RespCode = &v
	return b
}

func (b *StartAutoRenewResponseBuilder) Build() *StartAutoRenewResponse {
	return b.s
}


