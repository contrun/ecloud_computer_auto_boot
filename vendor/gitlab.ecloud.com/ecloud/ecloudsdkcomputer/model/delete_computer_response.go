// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeleteComputerResponse struct {

    // 响应数据
	Data *DeleteComputerResponseData `json:"data,omitempty"`
    // 请求ID
	RequestId *string `json:"requestId,omitempty"`
    // 响应信息
	RespMsg *string `json:"respMsg,omitempty"`
    // 响应码
	RespCode *string `json:"respCode,omitempty"`
}

func (s DeleteComputerResponse) String() string {
	return utils.Beautify(s)
}

func (s DeleteComputerResponse) GoString() string {
	return s.String()
}

func (s DeleteComputerResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteComputerResponse) SetData(v *DeleteComputerResponseData) *DeleteComputerResponse {
	s.Data = v
	return s
}

func (s *DeleteComputerResponse) SetRequestId(v string) *DeleteComputerResponse {
	s.RequestId = &v
	return s
}

func (s *DeleteComputerResponse) SetRespMsg(v string) *DeleteComputerResponse {
	s.RespMsg = &v
	return s
}

func (s *DeleteComputerResponse) SetRespCode(v string) *DeleteComputerResponse {
	s.RespCode = &v
	return s
}


type DeleteComputerResponseBuilder struct {
	s *DeleteComputerResponse
}

func NewDeleteComputerResponseBuilder() *DeleteComputerResponseBuilder {
	s := &DeleteComputerResponse{}
	b := &DeleteComputerResponseBuilder{s: s}
	return b
}

func (b *DeleteComputerResponseBuilder) Data(v *DeleteComputerResponseData) *DeleteComputerResponseBuilder {
	b.s.Data = v
	return b
}

func (b *DeleteComputerResponseBuilder) RequestId(v string) *DeleteComputerResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *DeleteComputerResponseBuilder) RespMsg(v string) *DeleteComputerResponseBuilder {
	b.s.RespMsg = &v
	return b
}

func (b *DeleteComputerResponseBuilder) RespCode(v string) *DeleteComputerResponseBuilder {
	b.s.RespCode = &v
	return b
}

func (b *DeleteComputerResponseBuilder) Build() *DeleteComputerResponse {
	return b.s
}


