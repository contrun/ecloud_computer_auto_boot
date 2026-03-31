// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UnifyInstanceQueryResponse struct {

    // 响应数据
	Data *UnifyInstanceQueryResponseData `json:"data,omitempty"`
    // 请求ID
	RequestId *string `json:"requestId,omitempty"`
    // 响应信息
	RespMsg *string `json:"respMsg,omitempty"`
    // 响应码
	RespCode *string `json:"respCode,omitempty"`
}

func (s UnifyInstanceQueryResponse) String() string {
	return utils.Beautify(s)
}

func (s UnifyInstanceQueryResponse) GoString() string {
	return s.String()
}

func (s UnifyInstanceQueryResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UnifyInstanceQueryResponse) SetData(v *UnifyInstanceQueryResponseData) *UnifyInstanceQueryResponse {
	s.Data = v
	return s
}

func (s *UnifyInstanceQueryResponse) SetRequestId(v string) *UnifyInstanceQueryResponse {
	s.RequestId = &v
	return s
}

func (s *UnifyInstanceQueryResponse) SetRespMsg(v string) *UnifyInstanceQueryResponse {
	s.RespMsg = &v
	return s
}

func (s *UnifyInstanceQueryResponse) SetRespCode(v string) *UnifyInstanceQueryResponse {
	s.RespCode = &v
	return s
}


type UnifyInstanceQueryResponseBuilder struct {
	s *UnifyInstanceQueryResponse
}

func NewUnifyInstanceQueryResponseBuilder() *UnifyInstanceQueryResponseBuilder {
	s := &UnifyInstanceQueryResponse{}
	b := &UnifyInstanceQueryResponseBuilder{s: s}
	return b
}

func (b *UnifyInstanceQueryResponseBuilder) Data(v *UnifyInstanceQueryResponseData) *UnifyInstanceQueryResponseBuilder {
	b.s.Data = v
	return b
}

func (b *UnifyInstanceQueryResponseBuilder) RequestId(v string) *UnifyInstanceQueryResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *UnifyInstanceQueryResponseBuilder) RespMsg(v string) *UnifyInstanceQueryResponseBuilder {
	b.s.RespMsg = &v
	return b
}

func (b *UnifyInstanceQueryResponseBuilder) RespCode(v string) *UnifyInstanceQueryResponseBuilder {
	b.s.RespCode = &v
	return b
}

func (b *UnifyInstanceQueryResponseBuilder) Build() *UnifyInstanceQueryResponse {
	return b.s
}


