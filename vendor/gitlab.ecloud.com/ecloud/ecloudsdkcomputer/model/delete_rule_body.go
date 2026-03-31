// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeleteRuleBody struct {
    position.Body
    // 规则UID
	RuleUid *string `json:"ruleUid,omitempty"`
    // 安全组UID
	SecurityGroupUid *string `json:"securityGroupUid,omitempty"`
}

func (s DeleteRuleBody) String() string {
	return utils.Beautify(s)
}

func (s DeleteRuleBody) GoString() string {
	return s.String()
}

func (s DeleteRuleBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteRuleBody) SetRuleUid(v string) *DeleteRuleBody {
	s.RuleUid = &v
	return s
}

func (s *DeleteRuleBody) SetSecurityGroupUid(v string) *DeleteRuleBody {
	s.SecurityGroupUid = &v
	return s
}


type DeleteRuleBodyBuilder struct {
	s *DeleteRuleBody
}

func NewDeleteRuleBodyBuilder() *DeleteRuleBodyBuilder {
	s := &DeleteRuleBody{}
	b := &DeleteRuleBodyBuilder{s: s}
	return b
}

func (b *DeleteRuleBodyBuilder) RuleUid(v string) *DeleteRuleBodyBuilder {
	b.s.RuleUid = &v
	return b
}

func (b *DeleteRuleBodyBuilder) SecurityGroupUid(v string) *DeleteRuleBodyBuilder {
	b.s.SecurityGroupUid = &v
	return b
}

func (b *DeleteRuleBodyBuilder) Build() *DeleteRuleBody {
	return b.s
}


