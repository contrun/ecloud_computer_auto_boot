// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateSNatRuleRequest struct {


	UpdateSNatRuleBody *UpdateSNatRuleBody `json:"updateSNatRuleBody,omitempty"`
}

func (s UpdateSNatRuleRequest) String() string {
	return utils.Beautify(s)
}

func (s UpdateSNatRuleRequest) GoString() string {
	return s.String()
}

func (s UpdateSNatRuleRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateSNatRuleRequest) SetUpdateSNatRuleBody(v *UpdateSNatRuleBody) *UpdateSNatRuleRequest {
	s.UpdateSNatRuleBody = v
	return s
}


type UpdateSNatRuleRequestBuilder struct {
	s *UpdateSNatRuleRequest
}

func NewUpdateSNatRuleRequestBuilder() *UpdateSNatRuleRequestBuilder {
	s := &UpdateSNatRuleRequest{}
	b := &UpdateSNatRuleRequestBuilder{s: s}
	return b
}

func (b *UpdateSNatRuleRequestBuilder) UpdateSNatRuleBody(v *UpdateSNatRuleBody) *UpdateSNatRuleRequestBuilder {
	b.s.UpdateSNatRuleBody = v
	return b
}

func (b *UpdateSNatRuleRequestBuilder) Build() *UpdateSNatRuleRequest {
	return b.s
}


