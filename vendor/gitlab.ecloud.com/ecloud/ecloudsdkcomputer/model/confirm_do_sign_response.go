// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ConfirmDoSignResponse struct {

    // 响应数据
	Data *string `json:"data,omitempty"`
    // 请求ID
	RequestId *string `json:"requestId,omitempty"`
    // 响应信息
	RespMsg *string `json:"respMsg,omitempty"`
    // 响应码
	RespCode *string `json:"respCode,omitempty"`
}

func (s ConfirmDoSignResponse) String() string {
	return utils.Beautify(s)
}

func (s ConfirmDoSignResponse) GoString() string {
	return s.String()
}

func (s ConfirmDoSignResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ConfirmDoSignResponse) SetData(v string) *ConfirmDoSignResponse {
	s.Data = &v
	return s
}

func (s *ConfirmDoSignResponse) SetRequestId(v string) *ConfirmDoSignResponse {
	s.RequestId = &v
	return s
}

func (s *ConfirmDoSignResponse) SetRespMsg(v string) *ConfirmDoSignResponse {
	s.RespMsg = &v
	return s
}

func (s *ConfirmDoSignResponse) SetRespCode(v string) *ConfirmDoSignResponse {
	s.RespCode = &v
	return s
}


type ConfirmDoSignResponseBuilder struct {
	s *ConfirmDoSignResponse
}

func NewConfirmDoSignResponseBuilder() *ConfirmDoSignResponseBuilder {
	s := &ConfirmDoSignResponse{}
	b := &ConfirmDoSignResponseBuilder{s: s}
	return b
}

func (b *ConfirmDoSignResponseBuilder) Data(v string) *ConfirmDoSignResponseBuilder {
	b.s.Data = &v
	return b
}

func (b *ConfirmDoSignResponseBuilder) RequestId(v string) *ConfirmDoSignResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *ConfirmDoSignResponseBuilder) RespMsg(v string) *ConfirmDoSignResponseBuilder {
	b.s.RespMsg = &v
	return b
}

func (b *ConfirmDoSignResponseBuilder) RespCode(v string) *ConfirmDoSignResponseBuilder {
	b.s.RespCode = &v
	return b
}

func (b *ConfirmDoSignResponseBuilder) Build() *ConfirmDoSignResponse {
	return b.s
}


