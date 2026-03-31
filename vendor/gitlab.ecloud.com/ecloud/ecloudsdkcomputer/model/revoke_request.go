// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type RevokeRequest struct {


	RevokeBody *RevokeBody `json:"revokeBody,omitempty"`
}

func (s RevokeRequest) String() string {
	return utils.Beautify(s)
}

func (s RevokeRequest) GoString() string {
	return s.String()
}

func (s RevokeRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *RevokeRequest) SetRevokeBody(v *RevokeBody) *RevokeRequest {
	s.RevokeBody = v
	return s
}


type RevokeRequestBuilder struct {
	s *RevokeRequest
}

func NewRevokeRequestBuilder() *RevokeRequestBuilder {
	s := &RevokeRequest{}
	b := &RevokeRequestBuilder{s: s}
	return b
}

func (b *RevokeRequestBuilder) RevokeBody(v *RevokeBody) *RevokeRequestBuilder {
	b.s.RevokeBody = v
	return b
}

func (b *RevokeRequestBuilder) Build() *RevokeRequest {
	return b.s
}


