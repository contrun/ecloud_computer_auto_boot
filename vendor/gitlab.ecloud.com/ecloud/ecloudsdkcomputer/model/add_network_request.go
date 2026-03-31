// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type AddNetworkRequest struct {


	AddNetworkBody *AddNetworkBody `json:"addNetworkBody,omitempty"`
}

func (s AddNetworkRequest) String() string {
	return utils.Beautify(s)
}

func (s AddNetworkRequest) GoString() string {
	return s.String()
}

func (s AddNetworkRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *AddNetworkRequest) SetAddNetworkBody(v *AddNetworkBody) *AddNetworkRequest {
	s.AddNetworkBody = v
	return s
}


type AddNetworkRequestBuilder struct {
	s *AddNetworkRequest
}

func NewAddNetworkRequestBuilder() *AddNetworkRequestBuilder {
	s := &AddNetworkRequest{}
	b := &AddNetworkRequestBuilder{s: s}
	return b
}

func (b *AddNetworkRequestBuilder) AddNetworkBody(v *AddNetworkBody) *AddNetworkRequestBuilder {
	b.s.AddNetworkBody = v
	return b
}

func (b *AddNetworkRequestBuilder) Build() *AddNetworkRequest {
	return b.s
}


