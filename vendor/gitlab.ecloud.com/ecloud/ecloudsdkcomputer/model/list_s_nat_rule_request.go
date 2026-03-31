// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListSNatRuleRequest struct {


	ListSNatRuleBody *ListSNatRuleBody `json:"listSNatRuleBody,omitempty"`
}

func (s ListSNatRuleRequest) String() string {
	return utils.Beautify(s)
}

func (s ListSNatRuleRequest) GoString() string {
	return s.String()
}

func (s ListSNatRuleRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListSNatRuleRequest) SetListSNatRuleBody(v *ListSNatRuleBody) *ListSNatRuleRequest {
	s.ListSNatRuleBody = v
	return s
}


type ListSNatRuleRequestBuilder struct {
	s *ListSNatRuleRequest
}

func NewListSNatRuleRequestBuilder() *ListSNatRuleRequestBuilder {
	s := &ListSNatRuleRequest{}
	b := &ListSNatRuleRequestBuilder{s: s}
	return b
}

func (b *ListSNatRuleRequestBuilder) ListSNatRuleBody(v *ListSNatRuleBody) *ListSNatRuleRequestBuilder {
	b.s.ListSNatRuleBody = v
	return b
}

func (b *ListSNatRuleRequestBuilder) Build() *ListSNatRuleRequest {
	return b.s
}


