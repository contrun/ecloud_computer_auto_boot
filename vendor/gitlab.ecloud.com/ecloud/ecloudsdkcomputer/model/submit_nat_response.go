// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type SubmitNatResponse struct {

    // 响应数据
	Data *SubmitNatResponseData `json:"data,omitempty"`
    // 请求ID
	RequestId *string `json:"requestId,omitempty"`
    // 响应信息
	RespMsg *string `json:"respMsg,omitempty"`
    // 响应编码
	RespCode *string `json:"respCode,omitempty"`
}

func (s SubmitNatResponse) String() string {
	return utils.Beautify(s)
}

func (s SubmitNatResponse) GoString() string {
	return s.String()
}

func (s SubmitNatResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *SubmitNatResponse) SetData(v *SubmitNatResponseData) *SubmitNatResponse {
	s.Data = v
	return s
}

func (s *SubmitNatResponse) SetRequestId(v string) *SubmitNatResponse {
	s.RequestId = &v
	return s
}

func (s *SubmitNatResponse) SetRespMsg(v string) *SubmitNatResponse {
	s.RespMsg = &v
	return s
}

func (s *SubmitNatResponse) SetRespCode(v string) *SubmitNatResponse {
	s.RespCode = &v
	return s
}


type SubmitNatResponseBuilder struct {
	s *SubmitNatResponse
}

func NewSubmitNatResponseBuilder() *SubmitNatResponseBuilder {
	s := &SubmitNatResponse{}
	b := &SubmitNatResponseBuilder{s: s}
	return b
}

func (b *SubmitNatResponseBuilder) Data(v *SubmitNatResponseData) *SubmitNatResponseBuilder {
	b.s.Data = v
	return b
}

func (b *SubmitNatResponseBuilder) RequestId(v string) *SubmitNatResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *SubmitNatResponseBuilder) RespMsg(v string) *SubmitNatResponseBuilder {
	b.s.RespMsg = &v
	return b
}

func (b *SubmitNatResponseBuilder) RespCode(v string) *SubmitNatResponseBuilder {
	b.s.RespCode = &v
	return b
}

func (b *SubmitNatResponseBuilder) Build() *SubmitNatResponse {
	return b.s
}


