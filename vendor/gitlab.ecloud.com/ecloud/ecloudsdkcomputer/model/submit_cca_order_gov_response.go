// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type SubmitCcaOrderGovResponse struct {

    // 响应数据
	Data *SubmitCcaOrderGovResponseData `json:"data,omitempty"`
    // 请求ID
	RequestId *string `json:"requestId,omitempty"`
    // 响应信息
	RespMsg *string `json:"respMsg,omitempty"`
    // 响应码
	RespCode *string `json:"respCode,omitempty"`
}

func (s SubmitCcaOrderGovResponse) String() string {
	return utils.Beautify(s)
}

func (s SubmitCcaOrderGovResponse) GoString() string {
	return s.String()
}

func (s SubmitCcaOrderGovResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *SubmitCcaOrderGovResponse) SetData(v *SubmitCcaOrderGovResponseData) *SubmitCcaOrderGovResponse {
	s.Data = v
	return s
}

func (s *SubmitCcaOrderGovResponse) SetRequestId(v string) *SubmitCcaOrderGovResponse {
	s.RequestId = &v
	return s
}

func (s *SubmitCcaOrderGovResponse) SetRespMsg(v string) *SubmitCcaOrderGovResponse {
	s.RespMsg = &v
	return s
}

func (s *SubmitCcaOrderGovResponse) SetRespCode(v string) *SubmitCcaOrderGovResponse {
	s.RespCode = &v
	return s
}


type SubmitCcaOrderGovResponseBuilder struct {
	s *SubmitCcaOrderGovResponse
}

func NewSubmitCcaOrderGovResponseBuilder() *SubmitCcaOrderGovResponseBuilder {
	s := &SubmitCcaOrderGovResponse{}
	b := &SubmitCcaOrderGovResponseBuilder{s: s}
	return b
}

func (b *SubmitCcaOrderGovResponseBuilder) Data(v *SubmitCcaOrderGovResponseData) *SubmitCcaOrderGovResponseBuilder {
	b.s.Data = v
	return b
}

func (b *SubmitCcaOrderGovResponseBuilder) RequestId(v string) *SubmitCcaOrderGovResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *SubmitCcaOrderGovResponseBuilder) RespMsg(v string) *SubmitCcaOrderGovResponseBuilder {
	b.s.RespMsg = &v
	return b
}

func (b *SubmitCcaOrderGovResponseBuilder) RespCode(v string) *SubmitCcaOrderGovResponseBuilder {
	b.s.RespCode = &v
	return b
}

func (b *SubmitCcaOrderGovResponseBuilder) Build() *SubmitCcaOrderGovResponse {
	return b.s
}


