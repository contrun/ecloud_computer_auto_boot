// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetZoneVpcRequest struct {


	GetZoneVpcBody *GetZoneVpcBody `json:"getZoneVpcBody,omitempty"`
}

func (s GetZoneVpcRequest) String() string {
	return utils.Beautify(s)
}

func (s GetZoneVpcRequest) GoString() string {
	return s.String()
}

func (s GetZoneVpcRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetZoneVpcRequest) SetGetZoneVpcBody(v *GetZoneVpcBody) *GetZoneVpcRequest {
	s.GetZoneVpcBody = v
	return s
}


type GetZoneVpcRequestBuilder struct {
	s *GetZoneVpcRequest
}

func NewGetZoneVpcRequestBuilder() *GetZoneVpcRequestBuilder {
	s := &GetZoneVpcRequest{}
	b := &GetZoneVpcRequestBuilder{s: s}
	return b
}

func (b *GetZoneVpcRequestBuilder) GetZoneVpcBody(v *GetZoneVpcBody) *GetZoneVpcRequestBuilder {
	b.s.GetZoneVpcBody = v
	return b
}

func (b *GetZoneVpcRequestBuilder) Build() *GetZoneVpcRequest {
	return b.s
}


