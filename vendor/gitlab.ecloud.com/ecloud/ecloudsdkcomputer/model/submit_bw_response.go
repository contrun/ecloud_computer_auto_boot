// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type SubmitBwResponse struct {

    // 响应数据
	Data *SubmitBwResponseData `json:"data,omitempty"`
    // 请求ID
	RequestId *string `json:"requestId,omitempty"`
    // 响应信息
	RespMsg *string `json:"respMsg,omitempty"`
    // 响应编码
	RespCode *string `json:"respCode,omitempty"`
}

func (s SubmitBwResponse) String() string {
	return utils.Beautify(s)
}

func (s SubmitBwResponse) GoString() string {
	return s.String()
}

func (s SubmitBwResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *SubmitBwResponse) SetData(v *SubmitBwResponseData) *SubmitBwResponse {
	s.Data = v
	return s
}

func (s *SubmitBwResponse) SetRequestId(v string) *SubmitBwResponse {
	s.RequestId = &v
	return s
}

func (s *SubmitBwResponse) SetRespMsg(v string) *SubmitBwResponse {
	s.RespMsg = &v
	return s
}

func (s *SubmitBwResponse) SetRespCode(v string) *SubmitBwResponse {
	s.RespCode = &v
	return s
}


type SubmitBwResponseBuilder struct {
	s *SubmitBwResponse
}

func NewSubmitBwResponseBuilder() *SubmitBwResponseBuilder {
	s := &SubmitBwResponse{}
	b := &SubmitBwResponseBuilder{s: s}
	return b
}

func (b *SubmitBwResponseBuilder) Data(v *SubmitBwResponseData) *SubmitBwResponseBuilder {
	b.s.Data = v
	return b
}

func (b *SubmitBwResponseBuilder) RequestId(v string) *SubmitBwResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *SubmitBwResponseBuilder) RespMsg(v string) *SubmitBwResponseBuilder {
	b.s.RespMsg = &v
	return b
}

func (b *SubmitBwResponseBuilder) RespCode(v string) *SubmitBwResponseBuilder {
	b.s.RespCode = &v
	return b
}

func (b *SubmitBwResponseBuilder) Build() *SubmitBwResponse {
	return b.s
}


