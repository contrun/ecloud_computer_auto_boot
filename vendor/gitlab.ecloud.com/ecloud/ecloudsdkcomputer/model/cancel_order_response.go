// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CancelOrderResponse struct {

    // 响应数据
	Data *string `json:"data,omitempty"`
    // 请求ID
	RequestId *string `json:"requestId,omitempty"`
    // 响应信息
	RespMsg *string `json:"respMsg,omitempty"`
    // 响应码
	RespCode *string `json:"respCode,omitempty"`
}

func (s CancelOrderResponse) String() string {
	return utils.Beautify(s)
}

func (s CancelOrderResponse) GoString() string {
	return s.String()
}

func (s CancelOrderResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CancelOrderResponse) SetData(v string) *CancelOrderResponse {
	s.Data = &v
	return s
}

func (s *CancelOrderResponse) SetRequestId(v string) *CancelOrderResponse {
	s.RequestId = &v
	return s
}

func (s *CancelOrderResponse) SetRespMsg(v string) *CancelOrderResponse {
	s.RespMsg = &v
	return s
}

func (s *CancelOrderResponse) SetRespCode(v string) *CancelOrderResponse {
	s.RespCode = &v
	return s
}


type CancelOrderResponseBuilder struct {
	s *CancelOrderResponse
}

func NewCancelOrderResponseBuilder() *CancelOrderResponseBuilder {
	s := &CancelOrderResponse{}
	b := &CancelOrderResponseBuilder{s: s}
	return b
}

func (b *CancelOrderResponseBuilder) Data(v string) *CancelOrderResponseBuilder {
	b.s.Data = &v
	return b
}

func (b *CancelOrderResponseBuilder) RequestId(v string) *CancelOrderResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *CancelOrderResponseBuilder) RespMsg(v string) *CancelOrderResponseBuilder {
	b.s.RespMsg = &v
	return b
}

func (b *CancelOrderResponseBuilder) RespCode(v string) *CancelOrderResponseBuilder {
	b.s.RespCode = &v
	return b
}

func (b *CancelOrderResponseBuilder) Build() *CancelOrderResponse {
	return b.s
}


