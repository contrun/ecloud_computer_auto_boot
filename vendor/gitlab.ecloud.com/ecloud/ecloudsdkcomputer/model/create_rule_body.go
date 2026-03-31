// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreateRuleBody struct {
    position.Body

	SecurityGroupRuleList *[]CreateRuleRequestSecurityGroupRuleList `json:"securityGroupRuleList,omitempty"`
    // 安全组UID
	SecurityGroupUid *string `json:"securityGroupUid,omitempty"`
}

func (s CreateRuleBody) String() string {
	return utils.Beautify(s)
}

func (s CreateRuleBody) GoString() string {
	return s.String()
}

func (s CreateRuleBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateRuleBody) SetSecurityGroupRuleList(v []CreateRuleRequestSecurityGroupRuleList) *CreateRuleBody {
	s.SecurityGroupRuleList = &v
	return s
}

func (s *CreateRuleBody) SetSecurityGroupUid(v string) *CreateRuleBody {
	s.SecurityGroupUid = &v
	return s
}


type CreateRuleBodyBuilder struct {
	s *CreateRuleBody
}

func NewCreateRuleBodyBuilder() *CreateRuleBodyBuilder {
	s := &CreateRuleBody{}
	b := &CreateRuleBodyBuilder{s: s}
	return b
}

func (b *CreateRuleBodyBuilder) SecurityGroupRuleList(v []CreateRuleRequestSecurityGroupRuleList) *CreateRuleBodyBuilder {
	b.s.SecurityGroupRuleList = &v
	return b
}

func (b *CreateRuleBodyBuilder) SecurityGroupUid(v string) *CreateRuleBodyBuilder {
	b.s.SecurityGroupUid = &v
	return b
}

func (b *CreateRuleBodyBuilder) Build() *CreateRuleBody {
	return b.s
}


