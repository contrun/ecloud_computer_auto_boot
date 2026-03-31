// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type PolicyListRequest struct {


	PolicyListBody *PolicyListBody `json:"policyListBody,omitempty"`
}

func (s PolicyListRequest) String() string {
	return utils.Beautify(s)
}

func (s PolicyListRequest) GoString() string {
	return s.String()
}

func (s PolicyListRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *PolicyListRequest) SetPolicyListBody(v *PolicyListBody) *PolicyListRequest {
	s.PolicyListBody = v
	return s
}


type PolicyListRequestBuilder struct {
	s *PolicyListRequest
}

func NewPolicyListRequestBuilder() *PolicyListRequestBuilder {
	s := &PolicyListRequest{}
	b := &PolicyListRequestBuilder{s: s}
	return b
}

func (b *PolicyListRequestBuilder) PolicyListBody(v *PolicyListBody) *PolicyListRequestBuilder {
	b.s.PolicyListBody = v
	return b
}

func (b *PolicyListRequestBuilder) Build() *PolicyListRequest {
	return b.s
}


