// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ProductsGovResponse struct {

    // 响应数据
	Data *[]ProductsGovResponseData `json:"data,omitempty"`
    // 请求ID
	RequestId *string `json:"requestId,omitempty"`
    // 响应信息
	RespMsg *string `json:"respMsg,omitempty"`
    // 响应码
	RespCode *string `json:"respCode,omitempty"`
}

func (s ProductsGovResponse) String() string {
	return utils.Beautify(s)
}

func (s ProductsGovResponse) GoString() string {
	return s.String()
}

func (s ProductsGovResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ProductsGovResponse) SetData(v []ProductsGovResponseData) *ProductsGovResponse {
	s.Data = &v
	return s
}

func (s *ProductsGovResponse) SetRequestId(v string) *ProductsGovResponse {
	s.RequestId = &v
	return s
}

func (s *ProductsGovResponse) SetRespMsg(v string) *ProductsGovResponse {
	s.RespMsg = &v
	return s
}

func (s *ProductsGovResponse) SetRespCode(v string) *ProductsGovResponse {
	s.RespCode = &v
	return s
}


type ProductsGovResponseBuilder struct {
	s *ProductsGovResponse
}

func NewProductsGovResponseBuilder() *ProductsGovResponseBuilder {
	s := &ProductsGovResponse{}
	b := &ProductsGovResponseBuilder{s: s}
	return b
}

func (b *ProductsGovResponseBuilder) Data(v []ProductsGovResponseData) *ProductsGovResponseBuilder {
	b.s.Data = &v
	return b
}

func (b *ProductsGovResponseBuilder) RequestId(v string) *ProductsGovResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *ProductsGovResponseBuilder) RespMsg(v string) *ProductsGovResponseBuilder {
	b.s.RespMsg = &v
	return b
}

func (b *ProductsGovResponseBuilder) RespCode(v string) *ProductsGovResponseBuilder {
	b.s.RespCode = &v
	return b
}

func (b *ProductsGovResponseBuilder) Build() *ProductsGovResponse {
	return b.s
}


