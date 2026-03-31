// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetSecurityGroupRuleInfoRequest struct {


	GetSecurityGroupRuleInfoBody *GetSecurityGroupRuleInfoBody `json:"getSecurityGroupRuleInfoBody,omitempty"`
}

func (s GetSecurityGroupRuleInfoRequest) String() string {
	return utils.Beautify(s)
}

func (s GetSecurityGroupRuleInfoRequest) GoString() string {
	return s.String()
}

func (s GetSecurityGroupRuleInfoRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetSecurityGroupRuleInfoRequest) SetGetSecurityGroupRuleInfoBody(v *GetSecurityGroupRuleInfoBody) *GetSecurityGroupRuleInfoRequest {
	s.GetSecurityGroupRuleInfoBody = v
	return s
}


type GetSecurityGroupRuleInfoRequestBuilder struct {
	s *GetSecurityGroupRuleInfoRequest
}

func NewGetSecurityGroupRuleInfoRequestBuilder() *GetSecurityGroupRuleInfoRequestBuilder {
	s := &GetSecurityGroupRuleInfoRequest{}
	b := &GetSecurityGroupRuleInfoRequestBuilder{s: s}
	return b
}

func (b *GetSecurityGroupRuleInfoRequestBuilder) GetSecurityGroupRuleInfoBody(v *GetSecurityGroupRuleInfoBody) *GetSecurityGroupRuleInfoRequestBuilder {
	b.s.GetSecurityGroupRuleInfoBody = v
	return b
}

func (b *GetSecurityGroupRuleInfoRequestBuilder) Build() *GetSecurityGroupRuleInfoRequest {
	return b.s
}


