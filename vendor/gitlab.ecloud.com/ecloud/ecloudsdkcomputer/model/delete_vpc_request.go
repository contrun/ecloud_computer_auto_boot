// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeleteVpcRequest struct {


	DeleteVpcBody *DeleteVpcBody `json:"deleteVpcBody,omitempty"`
}

func (s DeleteVpcRequest) String() string {
	return utils.Beautify(s)
}

func (s DeleteVpcRequest) GoString() string {
	return s.String()
}

func (s DeleteVpcRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteVpcRequest) SetDeleteVpcBody(v *DeleteVpcBody) *DeleteVpcRequest {
	s.DeleteVpcBody = v
	return s
}


type DeleteVpcRequestBuilder struct {
	s *DeleteVpcRequest
}

func NewDeleteVpcRequestBuilder() *DeleteVpcRequestBuilder {
	s := &DeleteVpcRequest{}
	b := &DeleteVpcRequestBuilder{s: s}
	return b
}

func (b *DeleteVpcRequestBuilder) DeleteVpcBody(v *DeleteVpcBody) *DeleteVpcRequestBuilder {
	b.s.DeleteVpcBody = v
	return b
}

func (b *DeleteVpcRequestBuilder) Build() *DeleteVpcRequest {
	return b.s
}


