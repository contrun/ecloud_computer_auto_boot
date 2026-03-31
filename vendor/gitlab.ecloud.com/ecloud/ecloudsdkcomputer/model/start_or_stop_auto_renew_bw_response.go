// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type StartOrStopAutoRenewBwResponse struct {

    // 响应数据
	Data *StartOrStopAutoRenewBwResponseData `json:"data,omitempty"`
    // 请求ID
	RequestId *string `json:"requestId,omitempty"`
    // 响应信息
	RespMsg *string `json:"respMsg,omitempty"`
    // 响应编码
	RespCode *string `json:"respCode,omitempty"`
}

func (s StartOrStopAutoRenewBwResponse) String() string {
	return utils.Beautify(s)
}

func (s StartOrStopAutoRenewBwResponse) GoString() string {
	return s.String()
}

func (s StartOrStopAutoRenewBwResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *StartOrStopAutoRenewBwResponse) SetData(v *StartOrStopAutoRenewBwResponseData) *StartOrStopAutoRenewBwResponse {
	s.Data = v
	return s
}

func (s *StartOrStopAutoRenewBwResponse) SetRequestId(v string) *StartOrStopAutoRenewBwResponse {
	s.RequestId = &v
	return s
}

func (s *StartOrStopAutoRenewBwResponse) SetRespMsg(v string) *StartOrStopAutoRenewBwResponse {
	s.RespMsg = &v
	return s
}

func (s *StartOrStopAutoRenewBwResponse) SetRespCode(v string) *StartOrStopAutoRenewBwResponse {
	s.RespCode = &v
	return s
}


type StartOrStopAutoRenewBwResponseBuilder struct {
	s *StartOrStopAutoRenewBwResponse
}

func NewStartOrStopAutoRenewBwResponseBuilder() *StartOrStopAutoRenewBwResponseBuilder {
	s := &StartOrStopAutoRenewBwResponse{}
	b := &StartOrStopAutoRenewBwResponseBuilder{s: s}
	return b
}

func (b *StartOrStopAutoRenewBwResponseBuilder) Data(v *StartOrStopAutoRenewBwResponseData) *StartOrStopAutoRenewBwResponseBuilder {
	b.s.Data = v
	return b
}

func (b *StartOrStopAutoRenewBwResponseBuilder) RequestId(v string) *StartOrStopAutoRenewBwResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *StartOrStopAutoRenewBwResponseBuilder) RespMsg(v string) *StartOrStopAutoRenewBwResponseBuilder {
	b.s.RespMsg = &v
	return b
}

func (b *StartOrStopAutoRenewBwResponseBuilder) RespCode(v string) *StartOrStopAutoRenewBwResponseBuilder {
	b.s.RespCode = &v
	return b
}

func (b *StartOrStopAutoRenewBwResponseBuilder) Build() *StartOrStopAutoRenewBwResponse {
	return b.s
}


