// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ModifyUserInfoResponseFailList struct {

    // 验证结果
	Result *string `json:"result,omitempty"`
    // 失败原因
	Reason *string `json:"reason,omitempty"`
    // 邮箱
	Mail *string `json:"mail,omitempty"`
    // 手机号
	Mobile *string `json:"mobile,omitempty"`
    // 用户登录账号
	UserName *string `json:"userName,omitempty"`
}

func (s ModifyUserInfoResponseFailList) String() string {
	return utils.Beautify(s)
}

func (s ModifyUserInfoResponseFailList) GoString() string {
	return s.String()
}

func (s ModifyUserInfoResponseFailList) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ModifyUserInfoResponseFailList) SetResult(v string) *ModifyUserInfoResponseFailList {
	s.Result = &v
	return s
}

func (s *ModifyUserInfoResponseFailList) SetReason(v string) *ModifyUserInfoResponseFailList {
	s.Reason = &v
	return s
}

func (s *ModifyUserInfoResponseFailList) SetMail(v string) *ModifyUserInfoResponseFailList {
	s.Mail = &v
	return s
}

func (s *ModifyUserInfoResponseFailList) SetMobile(v string) *ModifyUserInfoResponseFailList {
	s.Mobile = &v
	return s
}

func (s *ModifyUserInfoResponseFailList) SetUserName(v string) *ModifyUserInfoResponseFailList {
	s.UserName = &v
	return s
}


type ModifyUserInfoResponseFailListBuilder struct {
	s *ModifyUserInfoResponseFailList
}

func NewModifyUserInfoResponseFailListBuilder() *ModifyUserInfoResponseFailListBuilder {
	s := &ModifyUserInfoResponseFailList{}
	b := &ModifyUserInfoResponseFailListBuilder{s: s}
	return b
}

func (b *ModifyUserInfoResponseFailListBuilder) Result(v string) *ModifyUserInfoResponseFailListBuilder {
	b.s.Result = &v
	return b
}

func (b *ModifyUserInfoResponseFailListBuilder) Reason(v string) *ModifyUserInfoResponseFailListBuilder {
	b.s.Reason = &v
	return b
}

func (b *ModifyUserInfoResponseFailListBuilder) Mail(v string) *ModifyUserInfoResponseFailListBuilder {
	b.s.Mail = &v
	return b
}

func (b *ModifyUserInfoResponseFailListBuilder) Mobile(v string) *ModifyUserInfoResponseFailListBuilder {
	b.s.Mobile = &v
	return b
}

func (b *ModifyUserInfoResponseFailListBuilder) UserName(v string) *ModifyUserInfoResponseFailListBuilder {
	b.s.UserName = &v
	return b
}

func (b *ModifyUserInfoResponseFailListBuilder) Build() *ModifyUserInfoResponseFailList {
	return b.s
}


