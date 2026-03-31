// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ModifyUserInfoRequestBatchReqList struct {

    // 邮箱
	Mail *string `json:"mail,omitempty"`
    // 手机号
	Mobile *string `json:"mobile,omitempty"`
    // 备注
	Description *string `json:"description,omitempty"`
    // 用户登录账号
	UserName *string `json:"userName,omitempty"`
}

func (s ModifyUserInfoRequestBatchReqList) String() string {
	return utils.Beautify(s)
}

func (s ModifyUserInfoRequestBatchReqList) GoString() string {
	return s.String()
}

func (s ModifyUserInfoRequestBatchReqList) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ModifyUserInfoRequestBatchReqList) SetMail(v string) *ModifyUserInfoRequestBatchReqList {
	s.Mail = &v
	return s
}

func (s *ModifyUserInfoRequestBatchReqList) SetMobile(v string) *ModifyUserInfoRequestBatchReqList {
	s.Mobile = &v
	return s
}

func (s *ModifyUserInfoRequestBatchReqList) SetDescription(v string) *ModifyUserInfoRequestBatchReqList {
	s.Description = &v
	return s
}

func (s *ModifyUserInfoRequestBatchReqList) SetUserName(v string) *ModifyUserInfoRequestBatchReqList {
	s.UserName = &v
	return s
}


type ModifyUserInfoRequestBatchReqListBuilder struct {
	s *ModifyUserInfoRequestBatchReqList
}

func NewModifyUserInfoRequestBatchReqListBuilder() *ModifyUserInfoRequestBatchReqListBuilder {
	s := &ModifyUserInfoRequestBatchReqList{}
	b := &ModifyUserInfoRequestBatchReqListBuilder{s: s}
	return b
}

func (b *ModifyUserInfoRequestBatchReqListBuilder) Mail(v string) *ModifyUserInfoRequestBatchReqListBuilder {
	b.s.Mail = &v
	return b
}

func (b *ModifyUserInfoRequestBatchReqListBuilder) Mobile(v string) *ModifyUserInfoRequestBatchReqListBuilder {
	b.s.Mobile = &v
	return b
}

func (b *ModifyUserInfoRequestBatchReqListBuilder) Description(v string) *ModifyUserInfoRequestBatchReqListBuilder {
	b.s.Description = &v
	return b
}

func (b *ModifyUserInfoRequestBatchReqListBuilder) UserName(v string) *ModifyUserInfoRequestBatchReqListBuilder {
	b.s.UserName = &v
	return b
}

func (b *ModifyUserInfoRequestBatchReqListBuilder) Build() *ModifyUserInfoRequestBatchReqList {
	return b.s
}


