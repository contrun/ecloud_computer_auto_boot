// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetVpcSelectRequest struct {


	GetVpcSelectBody *GetVpcSelectBody `json:"getVpcSelectBody,omitempty"`
}

func (s GetVpcSelectRequest) String() string {
	return utils.Beautify(s)
}

func (s GetVpcSelectRequest) GoString() string {
	return s.String()
}

func (s GetVpcSelectRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetVpcSelectRequest) SetGetVpcSelectBody(v *GetVpcSelectBody) *GetVpcSelectRequest {
	s.GetVpcSelectBody = v
	return s
}


type GetVpcSelectRequestBuilder struct {
	s *GetVpcSelectRequest
}

func NewGetVpcSelectRequestBuilder() *GetVpcSelectRequestBuilder {
	s := &GetVpcSelectRequest{}
	b := &GetVpcSelectRequestBuilder{s: s}
	return b
}

func (b *GetVpcSelectRequestBuilder) GetVpcSelectBody(v *GetVpcSelectBody) *GetVpcSelectRequestBuilder {
	b.s.GetVpcSelectBody = v
	return b
}

func (b *GetVpcSelectRequestBuilder) Build() *GetVpcSelectRequest {
	return b.s
}


