// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BwChangeResponse struct {

    // 响应数据
	Data *BwChangeResponseData `json:"data,omitempty"`
    // 请求ID
	RequestId *string `json:"requestId,omitempty"`
    // 响应信息
	RespMsg *string `json:"respMsg,omitempty"`
    // 响应编码
	RespCode *string `json:"respCode,omitempty"`
}

func (s BwChangeResponse) String() string {
	return utils.Beautify(s)
}

func (s BwChangeResponse) GoString() string {
	return s.String()
}

func (s BwChangeResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BwChangeResponse) SetData(v *BwChangeResponseData) *BwChangeResponse {
	s.Data = v
	return s
}

func (s *BwChangeResponse) SetRequestId(v string) *BwChangeResponse {
	s.RequestId = &v
	return s
}

func (s *BwChangeResponse) SetRespMsg(v string) *BwChangeResponse {
	s.RespMsg = &v
	return s
}

func (s *BwChangeResponse) SetRespCode(v string) *BwChangeResponse {
	s.RespCode = &v
	return s
}


type BwChangeResponseBuilder struct {
	s *BwChangeResponse
}

func NewBwChangeResponseBuilder() *BwChangeResponseBuilder {
	s := &BwChangeResponse{}
	b := &BwChangeResponseBuilder{s: s}
	return b
}

func (b *BwChangeResponseBuilder) Data(v *BwChangeResponseData) *BwChangeResponseBuilder {
	b.s.Data = v
	return b
}

func (b *BwChangeResponseBuilder) RequestId(v string) *BwChangeResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *BwChangeResponseBuilder) RespMsg(v string) *BwChangeResponseBuilder {
	b.s.RespMsg = &v
	return b
}

func (b *BwChangeResponseBuilder) RespCode(v string) *BwChangeResponseBuilder {
	b.s.RespCode = &v
	return b
}

func (b *BwChangeResponseBuilder) Build() *BwChangeResponse {
	return b.s
}


