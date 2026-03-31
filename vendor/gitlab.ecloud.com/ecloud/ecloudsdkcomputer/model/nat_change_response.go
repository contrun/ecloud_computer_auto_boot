// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type NatChangeResponse struct {

    // 响应数据
	Data *NatChangeResponseData `json:"data,omitempty"`
    // 请求ID
	RequestId *string `json:"requestId,omitempty"`
    // 响应信息
	RespMsg *string `json:"respMsg,omitempty"`
    // 响应编码
	RespCode *string `json:"respCode,omitempty"`
}

func (s NatChangeResponse) String() string {
	return utils.Beautify(s)
}

func (s NatChangeResponse) GoString() string {
	return s.String()
}

func (s NatChangeResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *NatChangeResponse) SetData(v *NatChangeResponseData) *NatChangeResponse {
	s.Data = v
	return s
}

func (s *NatChangeResponse) SetRequestId(v string) *NatChangeResponse {
	s.RequestId = &v
	return s
}

func (s *NatChangeResponse) SetRespMsg(v string) *NatChangeResponse {
	s.RespMsg = &v
	return s
}

func (s *NatChangeResponse) SetRespCode(v string) *NatChangeResponse {
	s.RespCode = &v
	return s
}


type NatChangeResponseBuilder struct {
	s *NatChangeResponse
}

func NewNatChangeResponseBuilder() *NatChangeResponseBuilder {
	s := &NatChangeResponse{}
	b := &NatChangeResponseBuilder{s: s}
	return b
}

func (b *NatChangeResponseBuilder) Data(v *NatChangeResponseData) *NatChangeResponseBuilder {
	b.s.Data = v
	return b
}

func (b *NatChangeResponseBuilder) RequestId(v string) *NatChangeResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *NatChangeResponseBuilder) RespMsg(v string) *NatChangeResponseBuilder {
	b.s.RespMsg = &v
	return b
}

func (b *NatChangeResponseBuilder) RespCode(v string) *NatChangeResponseBuilder {
	b.s.RespCode = &v
	return b
}

func (b *NatChangeResponseBuilder) Build() *NatChangeResponse {
	return b.s
}


