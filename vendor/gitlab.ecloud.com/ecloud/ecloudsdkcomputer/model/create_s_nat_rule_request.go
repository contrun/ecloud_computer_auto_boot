// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreateSNatRuleRequest struct {


	CreateSNatRuleBody *CreateSNatRuleBody `json:"createSNatRuleBody,omitempty"`
}

func (s CreateSNatRuleRequest) String() string {
	return utils.Beautify(s)
}

func (s CreateSNatRuleRequest) GoString() string {
	return s.String()
}

func (s CreateSNatRuleRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateSNatRuleRequest) SetCreateSNatRuleBody(v *CreateSNatRuleBody) *CreateSNatRuleRequest {
	s.CreateSNatRuleBody = v
	return s
}


type CreateSNatRuleRequestBuilder struct {
	s *CreateSNatRuleRequest
}

func NewCreateSNatRuleRequestBuilder() *CreateSNatRuleRequestBuilder {
	s := &CreateSNatRuleRequest{}
	b := &CreateSNatRuleRequestBuilder{s: s}
	return b
}

func (b *CreateSNatRuleRequestBuilder) CreateSNatRuleBody(v *CreateSNatRuleBody) *CreateSNatRuleRequestBuilder {
	b.s.CreateSNatRuleBody = v
	return b
}

func (b *CreateSNatRuleRequestBuilder) Build() *CreateSNatRuleRequest {
	return b.s
}


