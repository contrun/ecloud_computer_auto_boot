// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateVpcRequest struct {


	UpdateVpcBody *UpdateVpcBody `json:"updateVpcBody,omitempty"`
}

func (s UpdateVpcRequest) String() string {
	return utils.Beautify(s)
}

func (s UpdateVpcRequest) GoString() string {
	return s.String()
}

func (s UpdateVpcRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateVpcRequest) SetUpdateVpcBody(v *UpdateVpcBody) *UpdateVpcRequest {
	s.UpdateVpcBody = v
	return s
}


type UpdateVpcRequestBuilder struct {
	s *UpdateVpcRequest
}

func NewUpdateVpcRequestBuilder() *UpdateVpcRequestBuilder {
	s := &UpdateVpcRequest{}
	b := &UpdateVpcRequestBuilder{s: s}
	return b
}

func (b *UpdateVpcRequestBuilder) UpdateVpcBody(v *UpdateVpcBody) *UpdateVpcRequestBuilder {
	b.s.UpdateVpcBody = v
	return b
}

func (b *UpdateVpcRequestBuilder) Build() *UpdateVpcRequest {
	return b.s
}


