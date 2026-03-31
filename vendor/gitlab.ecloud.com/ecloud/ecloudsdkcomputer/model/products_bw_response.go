// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ProductsBwResponse struct {

    // 响应数据
	Data *[]ProductsBwResponseData `json:"data,omitempty"`
    // 请求ID
	RequestId *string `json:"requestId,omitempty"`
    // 响应信息
	RespMsg *string `json:"respMsg,omitempty"`
    // 响应编码
	RespCode *string `json:"respCode,omitempty"`
}

func (s ProductsBwResponse) String() string {
	return utils.Beautify(s)
}

func (s ProductsBwResponse) GoString() string {
	return s.String()
}

func (s ProductsBwResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ProductsBwResponse) SetData(v []ProductsBwResponseData) *ProductsBwResponse {
	s.Data = &v
	return s
}

func (s *ProductsBwResponse) SetRequestId(v string) *ProductsBwResponse {
	s.RequestId = &v
	return s
}

func (s *ProductsBwResponse) SetRespMsg(v string) *ProductsBwResponse {
	s.RespMsg = &v
	return s
}

func (s *ProductsBwResponse) SetRespCode(v string) *ProductsBwResponse {
	s.RespCode = &v
	return s
}


type ProductsBwResponseBuilder struct {
	s *ProductsBwResponse
}

func NewProductsBwResponseBuilder() *ProductsBwResponseBuilder {
	s := &ProductsBwResponse{}
	b := &ProductsBwResponseBuilder{s: s}
	return b
}

func (b *ProductsBwResponseBuilder) Data(v []ProductsBwResponseData) *ProductsBwResponseBuilder {
	b.s.Data = &v
	return b
}

func (b *ProductsBwResponseBuilder) RequestId(v string) *ProductsBwResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *ProductsBwResponseBuilder) RespMsg(v string) *ProductsBwResponseBuilder {
	b.s.RespMsg = &v
	return b
}

func (b *ProductsBwResponseBuilder) RespCode(v string) *ProductsBwResponseBuilder {
	b.s.RespCode = &v
	return b
}

func (b *ProductsBwResponseBuilder) Build() *ProductsBwResponse {
	return b.s
}


