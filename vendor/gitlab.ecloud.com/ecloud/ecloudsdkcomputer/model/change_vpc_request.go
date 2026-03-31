// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ChangeVpcRequest struct {


	ChangeVpcBody *ChangeVpcBody `json:"changeVpcBody,omitempty"`
}

func (s ChangeVpcRequest) String() string {
	return utils.Beautify(s)
}

func (s ChangeVpcRequest) GoString() string {
	return s.String()
}

func (s ChangeVpcRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ChangeVpcRequest) SetChangeVpcBody(v *ChangeVpcBody) *ChangeVpcRequest {
	s.ChangeVpcBody = v
	return s
}


type ChangeVpcRequestBuilder struct {
	s *ChangeVpcRequest
}

func NewChangeVpcRequestBuilder() *ChangeVpcRequestBuilder {
	s := &ChangeVpcRequest{}
	b := &ChangeVpcRequestBuilder{s: s}
	return b
}

func (b *ChangeVpcRequestBuilder) ChangeVpcBody(v *ChangeVpcBody) *ChangeVpcRequestBuilder {
	b.s.ChangeVpcBody = v
	return b
}

func (b *ChangeVpcRequestBuilder) Build() *ChangeVpcRequest {
	return b.s
}


