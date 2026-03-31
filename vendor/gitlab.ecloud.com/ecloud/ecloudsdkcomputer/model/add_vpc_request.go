// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type AddVpcRequest struct {


	AddVpcBody *AddVpcBody `json:"addVpcBody,omitempty"`
}

func (s AddVpcRequest) String() string {
	return utils.Beautify(s)
}

func (s AddVpcRequest) GoString() string {
	return s.String()
}

func (s AddVpcRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *AddVpcRequest) SetAddVpcBody(v *AddVpcBody) *AddVpcRequest {
	s.AddVpcBody = v
	return s
}


type AddVpcRequestBuilder struct {
	s *AddVpcRequest
}

func NewAddVpcRequestBuilder() *AddVpcRequestBuilder {
	s := &AddVpcRequest{}
	b := &AddVpcRequestBuilder{s: s}
	return b
}

func (b *AddVpcRequestBuilder) AddVpcBody(v *AddVpcBody) *AddVpcRequestBuilder {
	b.s.AddVpcBody = v
	return b
}

func (b *AddVpcRequestBuilder) Build() *AddVpcRequest {
	return b.s
}


