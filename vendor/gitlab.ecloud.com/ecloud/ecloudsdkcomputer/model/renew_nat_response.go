// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type RenewNatResponse struct {

    // 响应数据
	Data *RenewNatResponseData `json:"data,omitempty"`
    // 请求ID
	RequestId *string `json:"requestId,omitempty"`
    // 响应信息
	RespMsg *string `json:"respMsg,omitempty"`
    // 响应编码
	RespCode *string `json:"respCode,omitempty"`
}

func (s RenewNatResponse) String() string {
	return utils.Beautify(s)
}

func (s RenewNatResponse) GoString() string {
	return s.String()
}

func (s RenewNatResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *RenewNatResponse) SetData(v *RenewNatResponseData) *RenewNatResponse {
	s.Data = v
	return s
}

func (s *RenewNatResponse) SetRequestId(v string) *RenewNatResponse {
	s.RequestId = &v
	return s
}

func (s *RenewNatResponse) SetRespMsg(v string) *RenewNatResponse {
	s.RespMsg = &v
	return s
}

func (s *RenewNatResponse) SetRespCode(v string) *RenewNatResponse {
	s.RespCode = &v
	return s
}


type RenewNatResponseBuilder struct {
	s *RenewNatResponse
}

func NewRenewNatResponseBuilder() *RenewNatResponseBuilder {
	s := &RenewNatResponse{}
	b := &RenewNatResponseBuilder{s: s}
	return b
}

func (b *RenewNatResponseBuilder) Data(v *RenewNatResponseData) *RenewNatResponseBuilder {
	b.s.Data = v
	return b
}

func (b *RenewNatResponseBuilder) RequestId(v string) *RenewNatResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *RenewNatResponseBuilder) RespMsg(v string) *RenewNatResponseBuilder {
	b.s.RespMsg = &v
	return b
}

func (b *RenewNatResponseBuilder) RespCode(v string) *RenewNatResponseBuilder {
	b.s.RespCode = &v
	return b
}

func (b *RenewNatResponseBuilder) Build() *RenewNatResponse {
	return b.s
}


