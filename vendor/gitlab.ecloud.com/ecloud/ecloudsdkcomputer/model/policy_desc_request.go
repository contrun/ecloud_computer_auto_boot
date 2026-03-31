// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type PolicyDescRequest struct {


	PolicyDescQuery *PolicyDescQuery `json:"policyDescQuery,omitempty"`
}

func (s PolicyDescRequest) String() string {
	return utils.Beautify(s)
}

func (s PolicyDescRequest) GoString() string {
	return s.String()
}

func (s PolicyDescRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *PolicyDescRequest) SetPolicyDescQuery(v *PolicyDescQuery) *PolicyDescRequest {
	s.PolicyDescQuery = v
	return s
}


type PolicyDescRequestBuilder struct {
	s *PolicyDescRequest
}

func NewPolicyDescRequestBuilder() *PolicyDescRequestBuilder {
	s := &PolicyDescRequest{}
	b := &PolicyDescRequestBuilder{s: s}
	return b
}

func (b *PolicyDescRequestBuilder) PolicyDescQuery(v *PolicyDescQuery) *PolicyDescRequestBuilder {
	b.s.PolicyDescQuery = v
	return b
}

func (b *PolicyDescRequestBuilder) Build() *PolicyDescRequest {
	return b.s
}


