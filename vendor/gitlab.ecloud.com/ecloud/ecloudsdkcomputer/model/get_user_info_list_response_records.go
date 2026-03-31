// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetUserInfoListResponseRecords struct {

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

func (s GetUserInfoListResponseRecords) String() string {
	return utils.Beautify(s)
}

func (s GetUserInfoListResponseRecords) GoString() string {
	return s.String()
}

func (s GetUserInfoListResponseRecords) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetUserInfoListResponseRecords) SetActivationType(v string) *GetUserInfoListResponseRecords {
	s.ActivationType = &v
	return s
}

func (s *GetUserInfoListResponseRecords) SetOrgName(v string) *GetUserInfoListResponseRecords {
	s.OrgName = &v
	return s
}

func (s *GetUserInfoListResponseRecords) SetMail(v string) *GetUserInfoListResponseRecords {
	s.Mail = &v
	return s
}

func (s *GetUserInfoListResponseRecords) SetBindCount(v int32) *GetUserInfoListResponseRecords {
	s.BindCount = &v
	return s
}

func (s *GetUserInfoListResponseRecords) SetMobile(v string) *GetUserInfoListResponseRecords {
	s.Mobile = &v
	return s
}

func (s *GetUserInfoListResponseRecords) SetDescription(v string) *GetUserInfoListResponseRecords {
	s.Description = &v
	return s
}

func (s *GetUserInfoListResponseRecords) SetMopUserName(v string) *GetUserInfoListResponseRecords {
	s.MopUserName = &v
	return s
}

func (s *GetUserInfoListResponseRecords) SetId(v string) *GetUserInfoListResponseRecords {
	s.Id = &v
	return s
}

func (s *GetUserInfoListResponseRecords) SetAuthorizationState(v string) *GetUserInfoListResponseRecords {
	s.AuthorizationState = &v
	return s
}

func (s *GetUserInfoListResponseRecords) SetUserName(v string) *GetUserInfoListResponseRecords {
	s.UserName = &v
	return s
}

func (s *GetUserInfoListResponseRecords) SetStatus(v string) *GetUserInfoListResponseRecords {
	s.Status = &v
	return s
}


type GetUserInfoListResponseRecordsBuilder struct {
	s *GetUserInfoListResponseRecords
}

func NewGetUserInfoListResponseRecordsBuilder() *GetUserInfoListResponseRecordsBuilder {
	s := &GetUserInfoListResponseRecords{}
	b := &GetUserInfoListResponseRecordsBuilder{s: s}
	return b
}

func (b *GetUserInfoListResponseRecordsBuilder) ActivationType(v string) *GetUserInfoListResponseRecordsBuilder {
	b.s.ActivationType = &v
	return b
}

func (b *GetUserInfoListResponseRecordsBuilder) OrgName(v string) *GetUserInfoListResponseRecordsBuilder {
	b.s.OrgName = &v
	return b
}

func (b *GetUserInfoListResponseRecordsBuilder) Mail(v string) *GetUserInfoListResponseRecordsBuilder {
	b.s.Mail = &v
	return b
}

func (b *GetUserInfoListResponseRecordsBuilder) BindCount(v int32) *GetUserInfoListResponseRecordsBuilder {
	b.s.BindCount = &v
	return b
}

func (b *GetUserInfoListResponseRecordsBuilder) Mobile(v string) *GetUserInfoListResponseRecordsBuilder {
	b.s.Mobile = &v
	return b
}

func (b *GetUserInfoListResponseRecordsBuilder) Description(v string) *GetUserInfoListResponseRecordsBuilder {
	b.s.Description = &v
	return b
}

func (b *GetUserInfoListResponseRecordsBuilder) MopUserName(v string) *GetUserInfoListResponseRecordsBuilder {
	b.s.MopUserName = &v
	return b
}

func (b *GetUserInfoListResponseRecordsBuilder) Id(v string) *GetUserInfoListResponseRecordsBuilder {
	b.s.Id = &v
	return b
}

func (b *GetUserInfoListResponseRecordsBuilder) AuthorizationState(v string) *GetUserInfoListResponseRecordsBuilder {
	b.s.AuthorizationState = &v
	return b
}

func (b *GetUserInfoListResponseRecordsBuilder) UserName(v string) *GetUserInfoListResponseRecordsBuilder {
	b.s.UserName = &v
	return b
}

func (b *GetUserInfoListResponseRecordsBuilder) Status(v string) *GetUserInfoListResponseRecordsBuilder {
	b.s.Status = &v
	return b
}

func (b *GetUserInfoListResponseRecordsBuilder) Build() *GetUserInfoListResponseRecords {
	return b.s
}


