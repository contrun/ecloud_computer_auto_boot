// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetUserDetailByPolicyIdResponseRecords struct {

    // 用户类型 0 管理员激活, 1 用户激活
	ActivationType *string `json:"activationType,omitempty"`
    // 所属组织名称（所属最小组织）
	OrgName *string `json:"orgName,omitempty"`
    // 邮箱
	Mail *string `json:"mail,omitempty"`
    // 已分配电脑数量
	BindCount *int32 `json:"bindCount,omitempty"`
    // 手机号
	Mobile *string `json:"mobile,omitempty"`
    // 用户关联状态, 0: 未关联, 1：和当前用户策略关联, 2:和其他用户策略关联
	AssociationState *int32 `json:"associationState,omitempty"`
    // 备注
	Description *string `json:"description,omitempty"`
    // 所属租户名称
	MopUserName *string `json:"mopUserName,omitempty"`
    // 用户id
	Id *string `json:"id,omitempty"`
    // 用户授权状态 0 未授权, 1 已授权
	AuthorizationState *string `json:"authorizationState,omitempty"`
    // 登录账号
	UserName *string `json:"userName,omitempty"`
    // 用户状态 0 已激活, 1 未激活
	Status *string `json:"status,omitempty"`
}

func (s GetUserDetailByPolicyIdResponseRecords) String() string {
	return utils.Beautify(s)
}

func (s GetUserDetailByPolicyIdResponseRecords) GoString() string {
	return s.String()
}

func (s GetUserDetailByPolicyIdResponseRecords) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetUserDetailByPolicyIdResponseRecords) SetActivationType(v string) *GetUserDetailByPolicyIdResponseRecords {
	s.ActivationType = &v
	return s
}

func (s *GetUserDetailByPolicyIdResponseRecords) SetOrgName(v string) *GetUserDetailByPolicyIdResponseRecords {
	s.OrgName = &v
	return s
}

func (s *GetUserDetailByPolicyIdResponseRecords) SetMail(v string) *GetUserDetailByPolicyIdResponseRecords {
	s.Mail = &v
	return s
}

func (s *GetUserDetailByPolicyIdResponseRecords) SetBindCount(v int32) *GetUserDetailByPolicyIdResponseRecords {
	s.BindCount = &v
	return s
}

func (s *GetUserDetailByPolicyIdResponseRecords) SetMobile(v string) *GetUserDetailByPolicyIdResponseRecords {
	s.Mobile = &v
	return s
}

func (s *GetUserDetailByPolicyIdResponseRecords) SetAssociationState(v int32) *GetUserDetailByPolicyIdResponseRecords {
	s.AssociationState = &v
	return s
}

func (s *GetUserDetailByPolicyIdResponseRecords) SetDescription(v string) *GetUserDetailByPolicyIdResponseRecords {
	s.Description = &v
	return s
}

func (s *GetUserDetailByPolicyIdResponseRecords) SetMopUserName(v string) *GetUserDetailByPolicyIdResponseRecords {
	s.MopUserName = &v
	return s
}

func (s *GetUserDetailByPolicyIdResponseRecords) SetId(v string) *GetUserDetailByPolicyIdResponseRecords {
	s.Id = &v
	return s
}

func (s *GetUserDetailByPolicyIdResponseRecords) SetAuthorizationState(v string) *GetUserDetailByPolicyIdResponseRecords {
	s.AuthorizationState = &v
	return s
}

func (s *GetUserDetailByPolicyIdResponseRecords) SetUserName(v string) *GetUserDetailByPolicyIdResponseRecords {
	s.UserName = &v
	return s
}

func (s *GetUserDetailByPolicyIdResponseRecords) SetStatus(v string) *GetUserDetailByPolicyIdResponseRecords {
	s.Status = &v
	return s
}


type GetUserDetailByPolicyIdResponseRecordsBuilder struct {
	s *GetUserDetailByPolicyIdResponseRecords
}

func NewGetUserDetailByPolicyIdResponseRecordsBuilder() *GetUserDetailByPolicyIdResponseRecordsBuilder {
	s := &GetUserDetailByPolicyIdResponseRecords{}
	b := &GetUserDetailByPolicyIdResponseRecordsBuilder{s: s}
	return b
}

func (b *GetUserDetailByPolicyIdResponseRecordsBuilder) ActivationType(v string) *GetUserDetailByPolicyIdResponseRecordsBuilder {
	b.s.ActivationType = &v
	return b
}

func (b *GetUserDetailByPolicyIdResponseRecordsBuilder) OrgName(v string) *GetUserDetailByPolicyIdResponseRecordsBuilder {
	b.s.OrgName = &v
	return b
}

func (b *GetUserDetailByPolicyIdResponseRecordsBuilder) Mail(v string) *GetUserDetailByPolicyIdResponseRecordsBuilder {
	b.s.Mail = &v
	return b
}

func (b *GetUserDetailByPolicyIdResponseRecordsBuilder) BindCount(v int32) *GetUserDetailByPolicyIdResponseRecordsBuilder {
	b.s.BindCount = &v
	return b
}

func (b *GetUserDetailByPolicyIdResponseRecordsBuilder) Mobile(v string) *GetUserDetailByPolicyIdResponseRecordsBuilder {
	b.s.Mobile = &v
	return b
}

func (b *GetUserDetailByPolicyIdResponseRecordsBuilder) AssociationState(v int32) *GetUserDetailByPolicyIdResponseRecordsBuilder {
	b.s.AssociationState = &v
	return b
}

func (b *GetUserDetailByPolicyIdResponseRecordsBuilder) Description(v string) *GetUserDetailByPolicyIdResponseRecordsBuilder {
	b.s.Description = &v
	return b
}

func (b *GetUserDetailByPolicyIdResponseRecordsBuilder) MopUserName(v string) *GetUserDetailByPolicyIdResponseRecordsBuilder {
	b.s.MopUserName = &v
	return b
}

func (b *GetUserDetailByPolicyIdResponseRecordsBuilder) Id(v string) *GetUserDetailByPolicyIdResponseRecordsBuilder {
	b.s.Id = &v
	return b
}

func (b *GetUserDetailByPolicyIdResponseRecordsBuilder) AuthorizationState(v string) *GetUserDetailByPolicyIdResponseRecordsBuilder {
	b.s.AuthorizationState = &v
	return b
}

func (b *GetUserDetailByPolicyIdResponseRecordsBuilder) UserName(v string) *GetUserDetailByPolicyIdResponseRecordsBuilder {
	b.s.UserName = &v
	return b
}

func (b *GetUserDetailByPolicyIdResponseRecordsBuilder) Status(v string) *GetUserDetailByPolicyIdResponseRecordsBuilder {
	b.s.Status = &v
	return b
}

func (b *GetUserDetailByPolicyIdResponseRecordsBuilder) Build() *GetUserDetailByPolicyIdResponseRecords {
	return b.s
}


