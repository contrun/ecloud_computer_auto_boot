// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreateRuleRequest struct {


	CreateRuleBody *CreateRuleBody `json:"createRuleBody,omitempty"`
}

func (s CreateRuleRequest) String() string {
	return utils.Beautify(s)
}

func (s CreateRuleRequest) GoString() string {
	return s.String()
}

func (s CreateRuleRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateRuleRequest) SetCreateRuleBody(v *CreateRuleBody) *CreateRuleRequest {
	s.CreateRuleBody = v
	return s
}


type CreateRuleRequestBuilder struct {
	s *CreateRuleRequest
}

func NewCreateRuleRequestBuilder() *CreateRuleRequestBuilder {
	s := &CreateRuleRequest{}
	b := &CreateRuleRequestBuilder{s: s}
	return b
}

func (b *CreateRuleRequestBuilder) CreateRuleBody(v *CreateRuleBody) *CreateRuleRequestBuilder {
	b.s.CreateRuleBody = v
	return b
}

func (b *CreateRuleRequestBuilder) Build() *CreateRuleRequest {
	return b.s
}


