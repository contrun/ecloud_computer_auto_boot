// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateNetworkRequest struct {


	UpdateNetworkBody *UpdateNetworkBody `json:"updateNetworkBody,omitempty"`
}

func (s UpdateNetworkRequest) String() string {
	return utils.Beautify(s)
}

func (s UpdateNetworkRequest) GoString() string {
	return s.String()
}

func (s UpdateNetworkRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateNetworkRequest) SetUpdateNetworkBody(v *UpdateNetworkBody) *UpdateNetworkRequest {
	s.UpdateNetworkBody = v
	return s
}


type UpdateNetworkRequestBuilder struct {
	s *UpdateNetworkRequest
}

func NewUpdateNetworkRequestBuilder() *UpdateNetworkRequestBuilder {
	s := &UpdateNetworkRequest{}
	b := &UpdateNetworkRequestBuilder{s: s}
	return b
}

func (b *UpdateNetworkRequestBuilder) UpdateNetworkBody(v *UpdateNetworkBody) *UpdateNetworkRequestBuilder {
	b.s.UpdateNetworkBody = v
	return b
}

func (b *UpdateNetworkRequestBuilder) Build() *UpdateNetworkRequest {
	return b.s
}


