// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateRuleRequest struct {


	UpdateRuleBody *UpdateRuleBody `json:"updateRuleBody,omitempty"`
}

func (s UpdateRuleRequest) String() string {
	return utils.Beautify(s)
}

func (s UpdateRuleRequest) GoString() string {
	return s.String()
}

func (s UpdateRuleRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateRuleRequest) SetUpdateRuleBody(v *UpdateRuleBody) *UpdateRuleRequest {
	s.UpdateRuleBody = v
	return s
}


type UpdateRuleRequestBuilder struct {
	s *UpdateRuleRequest
}

func NewUpdateRuleRequestBuilder() *UpdateRuleRequestBuilder {
	s := &UpdateRuleRequest{}
	b := &UpdateRuleRequestBuilder{s: s}
	return b
}

func (b *UpdateRuleRequestBuilder) UpdateRuleBody(v *UpdateRuleBody) *UpdateRuleRequestBuilder {
	b.s.UpdateRuleBody = v
	return b
}

func (b *UpdateRuleRequestBuilder) Build() *UpdateRuleRequest {
	return b.s
}


