// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ProductsNatResponse struct {

    // 响应数据
	Data *[]ProductsNatResponseData `json:"data,omitempty"`
    // 请求ID
	RequestId *string `json:"requestId,omitempty"`
    // 响应信息
	RespMsg *string `json:"respMsg,omitempty"`
    // 响应编码
	RespCode *string `json:"respCode,omitempty"`
}

func (s ProductsNatResponse) String() string {
	return utils.Beautify(s)
}

func (s ProductsNatResponse) GoString() string {
	return s.String()
}

func (s ProductsNatResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ProductsNatResponse) SetData(v []ProductsNatResponseData) *ProductsNatResponse {
	s.Data = &v
	return s
}

func (s *ProductsNatResponse) SetRequestId(v string) *ProductsNatResponse {
	s.RequestId = &v
	return s
}

func (s *ProductsNatResponse) SetRespMsg(v string) *ProductsNatResponse {
	s.RespMsg = &v
	return s
}

func (s *ProductsNatResponse) SetRespCode(v string) *ProductsNatResponse {
	s.RespCode = &v
	return s
}


type ProductsNatResponseBuilder struct {
	s *ProductsNatResponse
}

func NewProductsNatResponseBuilder() *ProductsNatResponseBuilder {
	s := &ProductsNatResponse{}
	b := &ProductsNatResponseBuilder{s: s}
	return b
}

func (b *ProductsNatResponseBuilder) Data(v []ProductsNatResponseData) *ProductsNatResponseBuilder {
	b.s.Data = &v
	return b
}

func (b *ProductsNatResponseBuilder) RequestId(v string) *ProductsNatResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *ProductsNatResponseBuilder) RespMsg(v string) *ProductsNatResponseBuilder {
	b.s.RespMsg = &v
	return b
}

func (b *ProductsNatResponseBuilder) RespCode(v string) *ProductsNatResponseBuilder {
	b.s.RespCode = &v
	return b
}

func (b *ProductsNatResponseBuilder) Build() *ProductsNatResponse {
	return b.s
}


