// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetNetworkSelectRequest struct {


	GetNetworkSelectBody *GetNetworkSelectBody `json:"getNetworkSelectBody,omitempty"`
}

func (s GetNetworkSelectRequest) String() string {
	return utils.Beautify(s)
}

func (s GetNetworkSelectRequest) GoString() string {
	return s.String()
}

func (s GetNetworkSelectRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetNetworkSelectRequest) SetGetNetworkSelectBody(v *GetNetworkSelectBody) *GetNetworkSelectRequest {
	s.GetNetworkSelectBody = v
	return s
}


type GetNetworkSelectRequestBuilder struct {
	s *GetNetworkSelectRequest
}

func NewGetNetworkSelectRequestBuilder() *GetNetworkSelectRequestBuilder {
	s := &GetNetworkSelectRequest{}
	b := &GetNetworkSelectRequestBuilder{s: s}
	return b
}

func (b *GetNetworkSelectRequestBuilder) GetNetworkSelectBody(v *GetNetworkSelectBody) *GetNetworkSelectRequestBuilder {
	b.s.GetNetworkSelectBody = v
	return b
}

func (b *GetNetworkSelectRequestBuilder) Build() *GetNetworkSelectRequest {
	return b.s
}


