// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ModifyUserInfoResponseSuccessList struct {

    // 验证结果
	Result *string `json:"result,omitempty"`
    // 邮箱
	Mail *string `json:"mail,omitempty"`
    // 手机号
	Mobile *string `json:"mobile,omitempty"`
    // 用户登录账号
	UserName *string `json:"userName,omitempty"`
}

func (s ModifyUserInfoResponseSuccessList) String() string {
	return utils.Beautify(s)
}

func (s ModifyUserInfoResponseSuccessList) GoString() string {
	return s.String()
}

func (s ModifyUserInfoResponseSuccessList) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ModifyUserInfoResponseSuccessList) SetResult(v string) *ModifyUserInfoResponseSuccessList {
	s.Result = &v
	return s
}

func (s *ModifyUserInfoResponseSuccessList) SetMail(v string) *ModifyUserInfoResponseSuccessList {
	s.Mail = &v
	return s
}

func (s *ModifyUserInfoResponseSuccessList) SetMobile(v string) *ModifyUserInfoResponseSuccessList {
	s.Mobile = &v
	return s
}

func (s *ModifyUserInfoResponseSuccessList) SetUserName(v string) *ModifyUserInfoResponseSuccessList {
	s.UserName = &v
	return s
}


type ModifyUserInfoResponseSuccessListBuilder struct {
	s *ModifyUserInfoResponseSuccessList
}

func NewModifyUserInfoResponseSuccessListBuilder() *ModifyUserInfoResponseSuccessListBuilder {
	s := &ModifyUserInfoResponseSuccessList{}
	b := &ModifyUserInfoResponseSuccessListBuilder{s: s}
	return b
}

func (b *ModifyUserInfoResponseSuccessListBuilder) Result(v string) *ModifyUserInfoResponseSuccessListBuilder {
	b.s.Result = &v
	return b
}

func (b *ModifyUserInfoResponseSuccessListBuilder) Mail(v string) *ModifyUserInfoResponseSuccessListBuilder {
	b.s.Mail = &v
	return b
}

func (b *ModifyUserInfoResponseSuccessListBuilder) Mobile(v string) *ModifyUserInfoResponseSuccessListBuilder {
	b.s.Mobile = &v
	return b
}

func (b *ModifyUserInfoResponseSuccessListBuilder) UserName(v string) *ModifyUserInfoResponseSuccessListBuilder {
	b.s.UserName = &v
	return b
}

func (b *ModifyUserInfoResponseSuccessListBuilder) Build() *ModifyUserInfoResponseSuccessList {
	return b.s
}


