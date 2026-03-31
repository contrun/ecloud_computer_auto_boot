// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetVpcRequest struct {


	GetVpcBody *GetVpcBody `json:"getVpcBody,omitempty"`
}

func (s GetVpcRequest) String() string {
	return utils.Beautify(s)
}

func (s GetVpcRequest) GoString() string {
	return s.String()
}

func (s GetVpcRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetVpcRequest) SetGetVpcBody(v *GetVpcBody) *GetVpcRequest {
	s.GetVpcBody = v
	return s
}


type GetVpcRequestBuilder struct {
	s *GetVpcRequest
}

func NewGetVpcRequestBuilder() *GetVpcRequestBuilder {
	s := &GetVpcRequest{}
	b := &GetVpcRequestBuilder{s: s}
	return b
}

func (b *GetVpcRequestBuilder) GetVpcBody(v *GetVpcBody) *GetVpcRequestBuilder {
	b.s.GetVpcBody = v
	return b
}

func (b *GetVpcRequestBuilder) Build() *GetVpcRequest {
	return b.s
}


