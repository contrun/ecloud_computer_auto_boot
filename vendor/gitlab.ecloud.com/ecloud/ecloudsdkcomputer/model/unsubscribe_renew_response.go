// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UnsubscribeRenewResponse struct {

    // 响应数据
	Data *UnsubscribeRenewResponseData `json:"data,omitempty"`
    // 请求ID
	RequestId *string `json:"requestId,omitempty"`
    // 响应信息
	RespMsg *string `json:"respMsg,omitempty"`
    // 响应码
	RespCode *string `json:"respCode,omitempty"`
}

func (s UnsubscribeRenewResponse) String() string {
	return utils.Beautify(s)
}

func (s UnsubscribeRenewResponse) GoString() string {
	return s.String()
}

func (s UnsubscribeRenewResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UnsubscribeRenewResponse) SetData(v *UnsubscribeRenewResponseData) *UnsubscribeRenewResponse {
	s.Data = v
	return s
}

func (s *UnsubscribeRenewResponse) SetRequestId(v string) *UnsubscribeRenewResponse {
	s.RequestId = &v
	return s
}

func (s *UnsubscribeRenewResponse) SetRespMsg(v string) *UnsubscribeRenewResponse {
	s.RespMsg = &v
	return s
}

func (s *UnsubscribeRenewResponse) SetRespCode(v string) *UnsubscribeRenewResponse {
	s.RespCode = &v
	return s
}


type UnsubscribeRenewResponseBuilder struct {
	s *UnsubscribeRenewResponse
}

func NewUnsubscribeRenewResponseBuilder() *UnsubscribeRenewResponseBuilder {
	s := &UnsubscribeRenewResponse{}
	b := &UnsubscribeRenewResponseBuilder{s: s}
	return b
}

func (b *UnsubscribeRenewResponseBuilder) Data(v *UnsubscribeRenewResponseData) *UnsubscribeRenewResponseBuilder {
	b.s.Data = v
	return b
}

func (b *UnsubscribeRenewResponseBuilder) RequestId(v string) *UnsubscribeRenewResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *UnsubscribeRenewResponseBuilder) RespMsg(v string) *UnsubscribeRenewResponseBuilder {
	b.s.RespMsg = &v
	return b
}

func (b *UnsubscribeRenewResponseBuilder) RespCode(v string) *UnsubscribeRenewResponseBuilder {
	b.s.RespCode = &v
	return b
}

func (b *UnsubscribeRenewResponseBuilder) Build() *UnsubscribeRenewResponse {
	return b.s
}


