// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type OrderChangeResponse struct {

    // 响应数据
	Data *OrderChangeResponseData `json:"data,omitempty"`
    // 请求ID
	RequestId *string `json:"requestId,omitempty"`
    // 响应信息
	RespMsg *string `json:"respMsg,omitempty"`
    // 响应码
	RespCode *string `json:"respCode,omitempty"`
}

func (s OrderChangeResponse) String() string {
	return utils.Beautify(s)
}

func (s OrderChangeResponse) GoString() string {
	return s.String()
}

func (s OrderChangeResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *OrderChangeResponse) SetData(v *OrderChangeResponseData) *OrderChangeResponse {
	s.Data = v
	return s
}

func (s *OrderChangeResponse) SetRequestId(v string) *OrderChangeResponse {
	s.RequestId = &v
	return s
}

func (s *OrderChangeResponse) SetRespMsg(v string) *OrderChangeResponse {
	s.RespMsg = &v
	return s
}

func (s *OrderChangeResponse) SetRespCode(v string) *OrderChangeResponse {
	s.RespCode = &v
	return s
}


type OrderChangeResponseBuilder struct {
	s *OrderChangeResponse
}

func NewOrderChangeResponseBuilder() *OrderChangeResponseBuilder {
	s := &OrderChangeResponse{}
	b := &OrderChangeResponseBuilder{s: s}
	return b
}

func (b *OrderChangeResponseBuilder) Data(v *OrderChangeResponseData) *OrderChangeResponseBuilder {
	b.s.Data = v
	return b
}

func (b *OrderChangeResponseBuilder) RequestId(v string) *OrderChangeResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *OrderChangeResponseBuilder) RespMsg(v string) *OrderChangeResponseBuilder {
	b.s.RespMsg = &v
	return b
}

func (b *OrderChangeResponseBuilder) RespCode(v string) *OrderChangeResponseBuilder {
	b.s.RespCode = &v
	return b
}

func (b *OrderChangeResponseBuilder) Build() *OrderChangeResponse {
	return b.s
}


