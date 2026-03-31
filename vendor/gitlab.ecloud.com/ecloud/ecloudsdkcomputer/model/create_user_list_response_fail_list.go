// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreateUserListResponseFailList struct {

    // 异常原因
	Reason *string `json:"reason,omitempty"`
    // 邮箱
	Mail *string `json:"mail,omitempty"`
    // 用户名
	UName *string `json:"uName,omitempty"`
    // 所属组织
	Organization *string `json:"organization,omitempty"`
    // 手机号
	Mobile *string `json:"mobile,omitempty"`
}

func (s CreateUserListResponseFailList) String() string {
	return utils.Beautify(s)
}

func (s CreateUserListResponseFailList) GoString() string {
	return s.String()
}

func (s CreateUserListResponseFailList) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateUserListResponseFailList) SetReason(v string) *CreateUserListResponseFailList {
	s.Reason = &v
	return s
}

func (s *CreateUserListResponseFailList) SetMail(v string) *CreateUserListResponseFailList {
	s.Mail = &v
	return s
}

func (s *CreateUserListResponseFailList) SetUName(v string) *CreateUserListResponseFailList {
	s.UName = &v
	return s
}

func (s *CreateUserListResponseFailList) SetOrganization(v string) *CreateUserListResponseFailList {
	s.Organization = &v
	return s
}

func (s *CreateUserListResponseFailList) SetMobile(v string) *CreateUserListResponseFailList {
	s.Mobile = &v
	return s
}


type CreateUserListResponseFailListBuilder struct {
	s *CreateUserListResponseFailList
}

func NewCreateUserListResponseFailListBuilder() *CreateUserListResponseFailListBuilder {
	s := &CreateUserListResponseFailList{}
	b := &CreateUserListResponseFailListBuilder{s: s}
	return b
}

func (b *CreateUserListResponseFailListBuilder) Reason(v string) *CreateUserListResponseFailListBuilder {
	b.s.Reason = &v
	return b
}

func (b *CreateUserListResponseFailListBuilder) Mail(v string) *CreateUserListResponseFailListBuilder {
	b.s.Mail = &v
	return b
}

func (b *CreateUserListResponseFailListBuilder) UName(v string) *CreateUserListResponseFailListBuilder {
	b.s.UName = &v
	return b
}

func (b *CreateUserListResponseFailListBuilder) Organization(v string) *CreateUserListResponseFailListBuilder {
	b.s.Organization = &v
	return b
}

func (b *CreateUserListResponseFailListBuilder) Mobile(v string) *CreateUserListResponseFailListBuilder {
	b.s.Mobile = &v
	return b
}

func (b *CreateUserListResponseFailListBuilder) Build() *CreateUserListResponseFailList {
	return b.s
}


