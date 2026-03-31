// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type SubmitCcaOrderResponse struct {

    // 响应数据
	Data *SubmitCcaOrderResponseData `json:"data,omitempty"`
    // 请求ID
	RequestId *string `json:"requestId,omitempty"`
    // 响应信息
	RespMsg *string `json:"respMsg,omitempty"`
    // 响应码
	RespCode *string `json:"respCode,omitempty"`
}

func (s SubmitCcaOrderResponse) String() string {
	return utils.Beautify(s)
}

func (s SubmitCcaOrderResponse) GoString() string {
	return s.String()
}

func (s SubmitCcaOrderResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *SubmitCcaOrderResponse) SetData(v *SubmitCcaOrderResponseData) *SubmitCcaOrderResponse {
	s.Data = v
	return s
}

func (s *SubmitCcaOrderResponse) SetRequestId(v string) *SubmitCcaOrderResponse {
	s.RequestId = &v
	return s
}

func (s *SubmitCcaOrderResponse) SetRespMsg(v string) *SubmitCcaOrderResponse {
	s.RespMsg = &v
	return s
}

func (s *SubmitCcaOrderResponse) SetRespCode(v string) *SubmitCcaOrderResponse {
	s.RespCode = &v
	return s
}


type SubmitCcaOrderResponseBuilder struct {
	s *SubmitCcaOrderResponse
}

func NewSubmitCcaOrderResponseBuilder() *SubmitCcaOrderResponseBuilder {
	s := &SubmitCcaOrderResponse{}
	b := &SubmitCcaOrderResponseBuilder{s: s}
	return b
}

func (b *SubmitCcaOrderResponseBuilder) Data(v *SubmitCcaOrderResponseData) *SubmitCcaOrderResponseBuilder {
	b.s.Data = v
	return b
}

func (b *SubmitCcaOrderResponseBuilder) RequestId(v string) *SubmitCcaOrderResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *SubmitCcaOrderResponseBuilder) RespMsg(v string) *SubmitCcaOrderResponseBuilder {
	b.s.RespMsg = &v
	return b
}

func (b *SubmitCcaOrderResponseBuilder) RespCode(v string) *SubmitCcaOrderResponseBuilder {
	b.s.RespCode = &v
	return b
}

func (b *SubmitCcaOrderResponseBuilder) Build() *SubmitCcaOrderResponse {
	return b.s
}


