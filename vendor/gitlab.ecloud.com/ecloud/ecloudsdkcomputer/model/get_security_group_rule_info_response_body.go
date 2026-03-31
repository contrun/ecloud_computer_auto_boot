// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetSecurityGroupRuleInfoResponseBody struct {

    // 规则列表
	RuleLists *[]GetSecurityGroupRuleInfoResponseRuleLists `json:"ruleLists,omitempty"`
    // 规则列表总行数
	RuleTotalCount *int32 `json:"ruleTotalCount,omitempty"`
}

func (s GetSecurityGroupRuleInfoResponseBody) String() string {
	return utils.Beautify(s)
}

func (s GetSecurityGroupRuleInfoResponseBody) GoString() string {
	return s.String()
}

func (s GetSecurityGroupRuleInfoResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetSecurityGroupRuleInfoResponseBody) SetRuleLists(v []GetSecurityGroupRuleInfoResponseRuleLists) *GetSecurityGroupRuleInfoResponseBody {
	s.RuleLists = &v
	return s
}

func (s *GetSecurityGroupRuleInfoResponseBody) SetRuleTotalCount(v int32) *GetSecurityGroupRuleInfoResponseBody {
	s.RuleTotalCount = &v
	return s
}


type GetSecurityGroupRuleInfoResponseBodyBuilder struct {
	s *GetSecurityGroupRuleInfoResponseBody
}

func NewGetSecurityGroupRuleInfoResponseBodyBuilder() *GetSecurityGroupRuleInfoResponseBodyBuilder {
	s := &GetSecurityGroupRuleInfoResponseBody{}
	b := &GetSecurityGroupRuleInfoResponseBodyBuilder{s: s}
	return b
}

func (b *GetSecurityGroupRuleInfoResponseBodyBuilder) RuleLists(v []GetSecurityGroupRuleInfoResponseRuleLists) *GetSecurityGroupRuleInfoResponseBodyBuilder {
	b.s.RuleLists = &v
	return b
}

func (b *GetSecurityGroupRuleInfoResponseBodyBuilder) RuleTotalCount(v int32) *GetSecurityGroupRuleInfoResponseBodyBuilder {
	b.s.RuleTotalCount = &v
	return b
}

func (b *GetSecurityGroupRuleInfoResponseBodyBuilder) Build() *GetSecurityGroupRuleInfoResponseBody {
	return b.s
}


