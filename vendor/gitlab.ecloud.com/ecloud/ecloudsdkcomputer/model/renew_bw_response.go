// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type RenewBwResponse struct {

    // 响应数据
	Data *RenewBwResponseData `json:"data,omitempty"`
    // 请求ID
	RequestId *string `json:"requestId,omitempty"`
    // 响应信息
	RespMsg *string `json:"respMsg,omitempty"`
    // 响应编码
	RespCode *string `json:"respCode,omitempty"`
}

func (s RenewBwResponse) String() string {
	return utils.Beautify(s)
}

func (s RenewBwResponse) GoString() string {
	return s.String()
}

func (s RenewBwResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *RenewBwResponse) SetData(v *RenewBwResponseData) *RenewBwResponse {
	s.Data = v
	return s
}

func (s *RenewBwResponse) SetRequestId(v string) *RenewBwResponse {
	s.RequestId = &v
	return s
}

func (s *RenewBwResponse) SetRespMsg(v string) *RenewBwResponse {
	s.RespMsg = &v
	return s
}

func (s *RenewBwResponse) SetRespCode(v string) *RenewBwResponse {
	s.RespCode = &v
	return s
}


type RenewBwResponseBuilder struct {
	s *RenewBwResponse
}

func NewRenewBwResponseBuilder() *RenewBwResponseBuilder {
	s := &RenewBwResponse{}
	b := &RenewBwResponseBuilder{s: s}
	return b
}

func (b *RenewBwResponseBuilder) Data(v *RenewBwResponseData) *RenewBwResponseBuilder {
	b.s.Data = v
	return b
}

func (b *RenewBwResponseBuilder) RequestId(v string) *RenewBwResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *RenewBwResponseBuilder) RespMsg(v string) *RenewBwResponseBuilder {
	b.s.RespMsg = &v
	return b
}

func (b *RenewBwResponseBuilder) RespCode(v string) *RenewBwResponseBuilder {
	b.s.RespCode = &v
	return b
}

func (b *RenewBwResponseBuilder) Build() *RenewBwResponse {
	return b.s
}


