// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetSubnetRequest struct {


	GetSubnetBody *GetSubnetBody `json:"getSubnetBody,omitempty"`
}

func (s GetSubnetRequest) String() string {
	return utils.Beautify(s)
}

func (s GetSubnetRequest) GoString() string {
	return s.String()
}

func (s GetSubnetRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetSubnetRequest) SetGetSubnetBody(v *GetSubnetBody) *GetSubnetRequest {
	s.GetSubnetBody = v
	return s
}


type GetSubnetRequestBuilder struct {
	s *GetSubnetRequest
}

func NewGetSubnetRequestBuilder() *GetSubnetRequestBuilder {
	s := &GetSubnetRequest{}
	b := &GetSubnetRequestBuilder{s: s}
	return b
}

func (b *GetSubnetRequestBuilder) GetSubnetBody(v *GetSubnetBody) *GetSubnetRequestBuilder {
	b.s.GetSubnetBody = v
	return b
}

func (b *GetSubnetRequestBuilder) Build() *GetSubnetRequest {
	return b.s
}


