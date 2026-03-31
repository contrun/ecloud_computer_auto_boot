// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type RenewComputerResponse struct {

    // 响应数据
	Data *RenewComputerResponseData `json:"data,omitempty"`
    // 请求ID
	RequestId *string `json:"requestId,omitempty"`
    // 响应信息
	RespMsg *string `json:"respMsg,omitempty"`
    // 响应码
	RespCode *string `json:"respCode,omitempty"`
}

func (s RenewComputerResponse) String() string {
	return utils.Beautify(s)
}

func (s RenewComputerResponse) GoString() string {
	return s.String()
}

func (s RenewComputerResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *RenewComputerResponse) SetData(v *RenewComputerResponseData) *RenewComputerResponse {
	s.Data = v
	return s
}

func (s *RenewComputerResponse) SetRequestId(v string) *RenewComputerResponse {
	s.RequestId = &v
	return s
}

func (s *RenewComputerResponse) SetRespMsg(v string) *RenewComputerResponse {
	s.RespMsg = &v
	return s
}

func (s *RenewComputerResponse) SetRespCode(v string) *RenewComputerResponse {
	s.RespCode = &v
	return s
}


type RenewComputerResponseBuilder struct {
	s *RenewComputerResponse
}

func NewRenewComputerResponseBuilder() *RenewComputerResponseBuilder {
	s := &RenewComputerResponse{}
	b := &RenewComputerResponseBuilder{s: s}
	return b
}

func (b *RenewComputerResponseBuilder) Data(v *RenewComputerResponseData) *RenewComputerResponseBuilder {
	b.s.Data = v
	return b
}

func (b *RenewComputerResponseBuilder) RequestId(v string) *RenewComputerResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *RenewComputerResponseBuilder) RespMsg(v string) *RenewComputerResponseBuilder {
	b.s.RespMsg = &v
	return b
}

func (b *RenewComputerResponseBuilder) RespCode(v string) *RenewComputerResponseBuilder {
	b.s.RespCode = &v
	return b
}

func (b *RenewComputerResponseBuilder) Build() *RenewComputerResponse {
	return b.s
}


