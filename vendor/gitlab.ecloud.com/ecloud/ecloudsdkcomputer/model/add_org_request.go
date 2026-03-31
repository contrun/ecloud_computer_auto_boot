// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type AddOrgRequest struct {


	AddOrgBody *AddOrgBody `json:"addOrgBody,omitempty"`
}

func (s AddOrgRequest) String() string {
	return utils.Beautify(s)
}

func (s AddOrgRequest) GoString() string {
	return s.String()
}

func (s AddOrgRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *AddOrgRequest) SetAddOrgBody(v *AddOrgBody) *AddOrgRequest {
	s.AddOrgBody = v
	return s
}


type AddOrgRequestBuilder struct {
	s *AddOrgRequest
}

func NewAddOrgRequestBuilder() *AddOrgRequestBuilder {
	s := &AddOrgRequest{}
	b := &AddOrgRequestBuilder{s: s}
	return b
}

func (b *AddOrgRequestBuilder) AddOrgBody(v *AddOrgBody) *AddOrgRequestBuilder {
	b.s.AddOrgBody = v
	return b
}

func (b *AddOrgRequestBuilder) Build() *AddOrgRequest {
	return b.s
}


