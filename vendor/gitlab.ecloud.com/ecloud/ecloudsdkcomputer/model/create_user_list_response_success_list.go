// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreateUserListResponseSuccessList struct {

    // 邮箱
	Mail *string `json:"mail,omitempty"`
    // 用户名
	UName *string `json:"uName,omitempty"`
    // 所属组织
	Organization *string `json:"organization,omitempty"`
    // 手机号
	Mobile *string `json:"mobile,omitempty"`
    // 备注
	Description *string `json:"description,omitempty"`
    // 用户ID
	Id *string `json:"id,omitempty"`
    // 登录账号
	UserName *string `json:"userName,omitempty"`
}

func (s CreateUserListResponseSuccessList) String() string {
	return utils.Beautify(s)
}

func (s CreateUserListResponseSuccessList) GoString() string {
	return s.String()
}

func (s CreateUserListResponseSuccessList) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateUserListResponseSuccessList) SetMail(v string) *CreateUserListResponseSuccessList {
	s.Mail = &v
	return s
}

func (s *CreateUserListResponseSuccessList) SetUName(v string) *CreateUserListResponseSuccessList {
	s.UName = &v
	return s
}

func (s *CreateUserListResponseSuccessList) SetOrganization(v string) *CreateUserListResponseSuccessList {
	s.Organization = &v
	return s
}

func (s *CreateUserListResponseSuccessList) SetMobile(v string) *CreateUserListResponseSuccessList {
	s.Mobile = &v
	return s
}

func (s *CreateUserListResponseSuccessList) SetDescription(v string) *CreateUserListResponseSuccessList {
	s.Description = &v
	return s
}

func (s *CreateUserListResponseSuccessList) SetId(v string) *CreateUserListResponseSuccessList {
	s.Id = &v
	return s
}

func (s *CreateUserListResponseSuccessList) SetUserName(v string) *CreateUserListResponseSuccessList {
	s.UserName = &v
	return s
}


type CreateUserListResponseSuccessListBuilder struct {
	s *CreateUserListResponseSuccessList
}

func NewCreateUserListResponseSuccessListBuilder() *CreateUserListResponseSuccessListBuilder {
	s := &CreateUserListResponseSuccessList{}
	b := &CreateUserListResponseSuccessListBuilder{s: s}
	return b
}

func (b *CreateUserListResponseSuccessListBuilder) Mail(v string) *CreateUserListResponseSuccessListBuilder {
	b.s.Mail = &v
	return b
}

func (b *CreateUserListResponseSuccessListBuilder) UName(v string) *CreateUserListResponseSuccessListBuilder {
	b.s.UName = &v
	return b
}

func (b *CreateUserListResponseSuccessListBuilder) Organization(v string) *CreateUserListResponseSuccessListBuilder {
	b.s.Organization = &v
	return b
}

func (b *CreateUserListResponseSuccessListBuilder) Mobile(v string) *CreateUserListResponseSuccessListBuilder {
	b.s.Mobile = &v
	return b
}

func (b *CreateUserListResponseSuccessListBuilder) Description(v string) *CreateUserListResponseSuccessListBuilder {
	b.s.Description = &v
	return b
}

func (b *CreateUserListResponseSuccessListBuilder) Id(v string) *CreateUserListResponseSuccessListBuilder {
	b.s.Id = &v
	return b
}

func (b *CreateUserListResponseSuccessListBuilder) UserName(v string) *CreateUserListResponseSuccessListBuilder {
	b.s.UserName = &v
	return b
}

func (b *CreateUserListResponseSuccessListBuilder) Build() *CreateUserListResponseSuccessList {
	return b.s
}


