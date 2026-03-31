// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreateUserListRequestAddUserVoList struct {

    // 用户类型 0: 管理员激活(默认管理激活) 1: 用户激活
	ActivationType *string `json:"activationType,omitempty"`
    // 密码
	Password *string `json:"password,omitempty"`
    // 邮箱 用户激活时,手机号和邮箱二选一
	Mail *string `json:"mail,omitempty"`
    // 用户名
	UName *string `json:"uName,omitempty"`
    // 组织id 默认二级组织
	CcempOrganizeId *string `json:"ccempOrganizeId,omitempty"`
    // 手机号 用户激活时,手机号和邮箱二选一
	Mobile *string `json:"mobile,omitempty"`
    // 备注
	Description *string `json:"description,omitempty"`
}

func (s CreateUserListRequestAddUserVoList) String() string {
	return utils.Beautify(s)
}

func (s CreateUserListRequestAddUserVoList) GoString() string {
	return s.String()
}

func (s CreateUserListRequestAddUserVoList) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateUserListRequestAddUserVoList) SetActivationType(v string) *CreateUserListRequestAddUserVoList {
	s.ActivationType = &v
	return s
}

func (s *CreateUserListRequestAddUserVoList) SetPassword(v string) *CreateUserListRequestAddUserVoList {
	s.Password = &v
	return s
}

func (s *CreateUserListRequestAddUserVoList) SetMail(v string) *CreateUserListRequestAddUserVoList {
	s.Mail = &v
	return s
}

func (s *CreateUserListRequestAddUserVoList) SetUName(v string) *CreateUserListRequestAddUserVoList {
	s.UName = &v
	return s
}

func (s *CreateUserListRequestAddUserVoList) SetCcempOrganizeId(v string) *CreateUserListRequestAddUserVoList {
	s.CcempOrganizeId = &v
	return s
}

func (s *CreateUserListRequestAddUserVoList) SetMobile(v string) *CreateUserListRequestAddUserVoList {
	s.Mobile = &v
	return s
}

func (s *CreateUserListRequestAddUserVoList) SetDescription(v string) *CreateUserListRequestAddUserVoList {
	s.Description = &v
	return s
}


type CreateUserListRequestAddUserVoListBuilder struct {
	s *CreateUserListRequestAddUserVoList
}

func NewCreateUserListRequestAddUserVoListBuilder() *CreateUserListRequestAddUserVoListBuilder {
	s := &CreateUserListRequestAddUserVoList{}
	b := &CreateUserListRequestAddUserVoListBuilder{s: s}
	return b
}

func (b *CreateUserListRequestAddUserVoListBuilder) ActivationType(v string) *CreateUserListRequestAddUserVoListBuilder {
	b.s.ActivationType = &v
	return b
}

func (b *CreateUserListRequestAddUserVoListBuilder) Password(v string) *CreateUserListRequestAddUserVoListBuilder {
	b.s.Password = &v
	return b
}

func (b *CreateUserListRequestAddUserVoListBuilder) Mail(v string) *CreateUserListRequestAddUserVoListBuilder {
	b.s.Mail = &v
	return b
}

func (b *CreateUserListRequestAddUserVoListBuilder) UName(v string) *CreateUserListRequestAddUserVoListBuilder {
	b.s.UName = &v
	return b
}

func (b *CreateUserListRequestAddUserVoListBuilder) CcempOrganizeId(v string) *CreateUserListRequestAddUserVoListBuilder {
	b.s.CcempOrganizeId = &v
	return b
}

func (b *CreateUserListRequestAddUserVoListBuilder) Mobile(v string) *CreateUserListRequestAddUserVoListBuilder {
	b.s.Mobile = &v
	return b
}

func (b *CreateUserListRequestAddUserVoListBuilder) Description(v string) *CreateUserListRequestAddUserVoListBuilder {
	b.s.Description = &v
	return b
}

func (b *CreateUserListRequestAddUserVoListBuilder) Build() *CreateUserListRequestAddUserVoList {
	return b.s
}


