// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UnsubscribeNatResponse struct {

    // 响应数据
	Data *UnsubscribeNatResponseData `json:"data,omitempty"`
    // 请求ID
	RequestId *string `json:"requestId,omitempty"`
    // 响应信息
	RespMsg *string `json:"respMsg,omitempty"`
    // 响应编码
	RespCode *string `json:"respCode,omitempty"`
}

func (s UnsubscribeNatResponse) String() string {
	return utils.Beautify(s)
}

func (s UnsubscribeNatResponse) GoString() string {
	return s.String()
}

func (s UnsubscribeNatResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UnsubscribeNatResponse) SetData(v *UnsubscribeNatResponseData) *UnsubscribeNatResponse {
	s.Data = v
	return s
}

func (s *UnsubscribeNatResponse) SetRequestId(v string) *UnsubscribeNatResponse {
	s.RequestId = &v
	return s
}

func (s *UnsubscribeNatResponse) SetRespMsg(v string) *UnsubscribeNatResponse {
	s.RespMsg = &v
	return s
}

func (s *UnsubscribeNatResponse) SetRespCode(v string) *UnsubscribeNatResponse {
	s.RespCode = &v
	return s
}


type UnsubscribeNatResponseBuilder struct {
	s *UnsubscribeNatResponse
}

func NewUnsubscribeNatResponseBuilder() *UnsubscribeNatResponseBuilder {
	s := &UnsubscribeNatResponse{}
	b := &UnsubscribeNatResponseBuilder{s: s}
	return b
}

func (b *UnsubscribeNatResponseBuilder) Data(v *UnsubscribeNatResponseData) *UnsubscribeNatResponseBuilder {
	b.s.Data = v
	return b
}

func (b *UnsubscribeNatResponseBuilder) RequestId(v string) *UnsubscribeNatResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *UnsubscribeNatResponseBuilder) RespMsg(v string) *UnsubscribeNatResponseBuilder {
	b.s.RespMsg = &v
	return b
}

func (b *UnsubscribeNatResponseBuilder) RespCode(v string) *UnsubscribeNatResponseBuilder {
	b.s.RespCode = &v
	return b
}

func (b *UnsubscribeNatResponseBuilder) Build() *UnsubscribeNatResponse {
	return b.s
}


