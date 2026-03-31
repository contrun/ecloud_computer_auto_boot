// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetUserPolicyDetailsResponseBody struct {

    // 用户锁定设置
	LockRuleList *[]GetUserPolicyDetailsResponseLockRuleList `json:"lockRuleList,omitempty"`
    // 策略名称
	PolicyName *string `json:"policyName,omitempty"`
    // 策略来源 1:云管 2:BOP 3:CHBN 4:企业门户,目前可以默认给固定值4
	SourceType *int32 `json:"sourceType,omitempty"`
    // 企业门户 组织id
	CcempOrganizeId *string `json:"ccempOrganizeId,omitempty"`
    // 策略描述
	PolicyDesc *string `json:"policyDesc,omitempty"`
    // 租户id
	TenantId *string `json:"tenantId,omitempty"`
    // 接入ip白名单设置
	IpRuleList *[]GetUserPolicyDetailsResponseIpRuleList `json:"ipRuleList,omitempty"`
    // 用户锁定类型| 0:未锁定 1:锁定 2:按固定时间锁定
	LockType *int32 `json:"lockType,omitempty"`
    // 策略id  为空时 表示新增  不为空 表示编辑
	UserPolicyId *string `json:"userPolicyId,omitempty"`
}

func (s GetUserPolicyDetailsResponseBody) String() string {
	return utils.Beautify(s)
}

func (s GetUserPolicyDetailsResponseBody) GoString() string {
	return s.String()
}

func (s GetUserPolicyDetailsResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetUserPolicyDetailsResponseBody) SetLockRuleList(v []GetUserPolicyDetailsResponseLockRuleList) *GetUserPolicyDetailsResponseBody {
	s.LockRuleList = &v
	return s
}

func (s *GetUserPolicyDetailsResponseBody) SetPolicyName(v string) *GetUserPolicyDetailsResponseBody {
	s.PolicyName = &v
	return s
}

func (s *GetUserPolicyDetailsResponseBody) SetSourceType(v int32) *GetUserPolicyDetailsResponseBody {
	s.SourceType = &v
	return s
}

func (s *GetUserPolicyDetailsResponseBody) SetCcempOrganizeId(v string) *GetUserPolicyDetailsResponseBody {
	s.CcempOrganizeId = &v
	return s
}

func (s *GetUserPolicyDetailsResponseBody) SetPolicyDesc(v string) *GetUserPolicyDetailsResponseBody {
	s.PolicyDesc = &v
	return s
}

func (s *GetUserPolicyDetailsResponseBody) SetTenantId(v string) *GetUserPolicyDetailsResponseBody {
	s.TenantId = &v
	return s
}

func (s *GetUserPolicyDetailsResponseBody) SetIpRuleList(v []GetUserPolicyDetailsResponseIpRuleList) *GetUserPolicyDetailsResponseBody {
	s.IpRuleList = &v
	return s
}

func (s *GetUserPolicyDetailsResponseBody) SetLockType(v int32) *GetUserPolicyDetailsResponseBody {
	s.LockType = &v
	return s
}

func (s *GetUserPolicyDetailsResponseBody) SetUserPolicyId(v string) *GetUserPolicyDetailsResponseBody {
	s.UserPolicyId = &v
	return s
}


type GetUserPolicyDetailsResponseBodyBuilder struct {
	s *GetUserPolicyDetailsResponseBody
}

func NewGetUserPolicyDetailsResponseBodyBuilder() *GetUserPolicyDetailsResponseBodyBuilder {
	s := &GetUserPolicyDetailsResponseBody{}
	b := &GetUserPolicyDetailsResponseBodyBuilder{s: s}
	return b
}

func (b *GetUserPolicyDetailsResponseBodyBuilder) LockRuleList(v []GetUserPolicyDetailsResponseLockRuleList) *GetUserPolicyDetailsResponseBodyBuilder {
	b.s.LockRuleList = &v
	return b
}

func (b *GetUserPolicyDetailsResponseBodyBuilder) PolicyName(v string) *GetUserPolicyDetailsResponseBodyBuilder {
	b.s.PolicyName = &v
	return b
}

func (b *GetUserPolicyDetailsResponseBodyBuilder) SourceType(v int32) *GetUserPolicyDetailsResponseBodyBuilder {
	b.s.SourceType = &v
	return b
}

func (b *GetUserPolicyDetailsResponseBodyBuilder) CcempOrganizeId(v string) *GetUserPolicyDetailsResponseBodyBuilder {
	b.s.CcempOrganizeId = &v
	return b
}

func (b *GetUserPolicyDetailsResponseBodyBuilder) PolicyDesc(v string) *GetUserPolicyDetailsResponseBodyBuilder {
	b.s.PolicyDesc = &v
	return b
}

func (b *GetUserPolicyDetailsResponseBodyBuilder) TenantId(v string) *GetUserPolicyDetailsResponseBodyBuilder {
	b.s.TenantId = &v
	return b
}

func (b *GetUserPolicyDetailsResponseBodyBuilder) IpRuleList(v []GetUserPolicyDetailsResponseIpRuleList) *GetUserPolicyDetailsResponseBodyBuilder {
	b.s.IpRuleList = &v
	return b
}

func (b *GetUserPolicyDetailsResponseBodyBuilder) LockType(v int32) *GetUserPolicyDetailsResponseBodyBuilder {
	b.s.LockType = &v
	return b
}

func (b *GetUserPolicyDetailsResponseBodyBuilder) UserPolicyId(v string) *GetUserPolicyDetailsResponseBodyBuilder {
	b.s.UserPolicyId = &v
	return b
}

func (b *GetUserPolicyDetailsResponseBodyBuilder) Build() *GetUserPolicyDetailsResponseBody {
	return b.s
}


