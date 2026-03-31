// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BatchResetPwdResponseSuccessList struct {

    // 重置结果
	Result *string `json:"result,omitempty"`
    // 邮箱
	Mail *string `json:"mail,omitempty"`
    // 手机号
	Mobile *string `json:"mobile,omitempty"`
    // 租户id
	TenantId *string `json:"tenantId,omitempty"`
    // 云电脑用户id
	UserId *string `json:"userId,omitempty"`
}

func (s BatchResetPwdResponseSuccessList) String() string {
	return utils.Beautify(s)
}

func (s BatchResetPwdResponseSuccessList) GoString() string {
	return s.String()
}

func (s BatchResetPwdResponseSuccessList) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BatchResetPwdResponseSuccessList) SetResult(v string) *BatchResetPwdResponseSuccessList {
	s.Result = &v
	return s
}

func (s *BatchResetPwdResponseSuccessList) SetMail(v string) *BatchResetPwdResponseSuccessList {
	s.Mail = &v
	return s
}

func (s *BatchResetPwdResponseSuccessList) SetMobile(v string) *BatchResetPwdResponseSuccessList {
	s.Mobile = &v
	return s
}

func (s *BatchResetPwdResponseSuccessList) SetTenantId(v string) *BatchResetPwdResponseSuccessList {
	s.TenantId = &v
	return s
}

func (s *BatchResetPwdResponseSuccessList) SetUserId(v string) *BatchResetPwdResponseSuccessList {
	s.UserId = &v
	return s
}


type BatchResetPwdResponseSuccessListBuilder struct {
	s *BatchResetPwdResponseSuccessList
}

func NewBatchResetPwdResponseSuccessListBuilder() *BatchResetPwdResponseSuccessListBuilder {
	s := &BatchResetPwdResponseSuccessList{}
	b := &BatchResetPwdResponseSuccessListBuilder{s: s}
	return b
}

func (b *BatchResetPwdResponseSuccessListBuilder) Result(v string) *BatchResetPwdResponseSuccessListBuilder {
	b.s.Result = &v
	return b
}

func (b *BatchResetPwdResponseSuccessListBuilder) Mail(v string) *BatchResetPwdResponseSuccessListBuilder {
	b.s.Mail = &v
	return b
}

func (b *BatchResetPwdResponseSuccessListBuilder) Mobile(v string) *BatchResetPwdResponseSuccessListBuilder {
	b.s.Mobile = &v
	return b
}

func (b *BatchResetPwdResponseSuccessListBuilder) TenantId(v string) *BatchResetPwdResponseSuccessListBuilder {
	b.s.TenantId = &v
	return b
}

func (b *BatchResetPwdResponseSuccessListBuilder) UserId(v string) *BatchResetPwdResponseSuccessListBuilder {
	b.s.UserId = &v
	return b
}

func (b *BatchResetPwdResponseSuccessListBuilder) Build() *BatchResetPwdResponseSuccessList {
	return b.s
}


