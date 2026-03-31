// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetNetworkRequest struct {


	GetNetworkBody *GetNetworkBody `json:"getNetworkBody,omitempty"`
}

func (s GetNetworkRequest) String() string {
	return utils.Beautify(s)
}

func (s GetNetworkRequest) GoString() string {
	return s.String()
}

func (s GetNetworkRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetNetworkRequest) SetGetNetworkBody(v *GetNetworkBody) *GetNetworkRequest {
	s.GetNetworkBody = v
	return s
}


type GetNetworkRequestBuilder struct {
	s *GetNetworkRequest
}

func NewGetNetworkRequestBuilder() *GetNetworkRequestBuilder {
	s := &GetNetworkRequest{}
	b := &GetNetworkRequestBuilder{s: s}
	return b
}

func (b *GetNetworkRequestBuilder) GetNetworkBody(v *GetNetworkBody) *GetNetworkRequestBuilder {
	b.s.GetNetworkBody = v
	return b
}

func (b *GetNetworkRequestBuilder) Build() *GetNetworkRequest {
	return b.s
}


