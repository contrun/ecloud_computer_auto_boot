// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetSecurityGroupRuleInfoBody struct {
    position.Body
    // 当前页起始下标
	Start *int32 `json:"start,omitempty"`
    // 每页显示条数
	PageSize *int32 `json:"pageSize,omitempty"`
    // 当前页
	CurrentPage *int32 `json:"currentPage,omitempty"`
    // 安全组UID
	SecurityGroupUid *string `json:"securityGroupUid,omitempty"`
}

func (s GetSecurityGroupRuleInfoBody) String() string {
	return utils.Beautify(s)
}

func (s GetSecurityGroupRuleInfoBody) GoString() string {
	return s.String()
}

func (s GetSecurityGroupRuleInfoBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetSecurityGroupRuleInfoBody) SetStart(v int32) *GetSecurityGroupRuleInfoBody {
	s.Start = &v
	return s
}

func (s *GetSecurityGroupRuleInfoBody) SetPageSize(v int32) *GetSecurityGroupRuleInfoBody {
	s.PageSize = &v
	return s
}

func (s *GetSecurityGroupRuleInfoBody) SetCurrentPage(v int32) *GetSecurityGroupRuleInfoBody {
	s.CurrentPage = &v
	return s
}

func (s *GetSecurityGroupRuleInfoBody) SetSecurityGroupUid(v string) *GetSecurityGroupRuleInfoBody {
	s.SecurityGroupUid = &v
	return s
}


type GetSecurityGroupRuleInfoBodyBuilder struct {
	s *GetSecurityGroupRuleInfoBody
}

func NewGetSecurityGroupRuleInfoBodyBuilder() *GetSecurityGroupRuleInfoBodyBuilder {
	s := &GetSecurityGroupRuleInfoBody{}
	b := &GetSecurityGroupRuleInfoBodyBuilder{s: s}
	return b
}

func (b *GetSecurityGroupRuleInfoBodyBuilder) Start(v int32) *GetSecurityGroupRuleInfoBodyBuilder {
	b.s.Start = &v
	return b
}

func (b *GetSecurityGroupRuleInfoBodyBuilder) PageSize(v int32) *GetSecurityGroupRuleInfoBodyBuilder {
	b.s.PageSize = &v
	return b
}

func (b *GetSecurityGroupRuleInfoBodyBuilder) CurrentPage(v int32) *GetSecurityGroupRuleInfoBodyBuilder {
	b.s.CurrentPage = &v
	return b
}

func (b *GetSecurityGroupRuleInfoBodyBuilder) SecurityGroupUid(v string) *GetSecurityGroupRuleInfoBodyBuilder {
	b.s.SecurityGroupUid = &v
	return b
}

func (b *GetSecurityGroupRuleInfoBodyBuilder) Build() *GetSecurityGroupRuleInfoBody {
	return b.s
}


