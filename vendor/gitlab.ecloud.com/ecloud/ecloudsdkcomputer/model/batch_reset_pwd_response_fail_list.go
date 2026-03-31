// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BatchResetPwdResponseFailList struct {

    // 重置结果
	Result *string `json:"result,omitempty"`
    // 重置失败原因
	Reason *string `json:"reason,omitempty"`
    // 邮箱
	Mail *string `json:"mail,omitempty"`
    // 手机号
	Mobile *string `json:"mobile,omitempty"`
    // 租户id
	TenantId *string `json:"tenantId,omitempty"`
    // 云电脑用户id
	UserId *string `json:"userId,omitempty"`
}

func (s BatchResetPwdResponseFailList) String() string {
	return utils.Beautify(s)
}

func (s BatchResetPwdResponseFailList) GoString() string {
	return s.String()
}

func (s BatchResetPwdResponseFailList) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BatchResetPwdResponseFailList) SetResult(v string) *BatchResetPwdResponseFailList {
	s.Result = &v
	return s
}

func (s *BatchResetPwdResponseFailList) SetReason(v string) *BatchResetPwdResponseFailList {
	s.Reason = &v
	return s
}

func (s *BatchResetPwdResponseFailList) SetMail(v string) *BatchResetPwdResponseFailList {
	s.Mail = &v
	return s
}

func (s *BatchResetPwdResponseFailList) SetMobile(v string) *BatchResetPwdResponseFailList {
	s.Mobile = &v
	return s
}

func (s *BatchResetPwdResponseFailList) SetTenantId(v string) *BatchResetPwdResponseFailList {
	s.TenantId = &v
	return s
}

func (s *BatchResetPwdResponseFailList) SetUserId(v string) *BatchResetPwdResponseFailList {
	s.UserId = &v
	return s
}


type BatchResetPwdResponseFailListBuilder struct {
	s *BatchResetPwdResponseFailList
}

func NewBatchResetPwdResponseFailListBuilder() *BatchResetPwdResponseFailListBuilder {
	s := &BatchResetPwdResponseFailList{}
	b := &BatchResetPwdResponseFailListBuilder{s: s}
	return b
}

func (b *BatchResetPwdResponseFailListBuilder) Result(v string) *BatchResetPwdResponseFailListBuilder {
	b.s.Result = &v
	return b
}

func (b *BatchResetPwdResponseFailListBuilder) Reason(v string) *BatchResetPwdResponseFailListBuilder {
	b.s.Reason = &v
	return b
}

func (b *BatchResetPwdResponseFailListBuilder) Mail(v string) *BatchResetPwdResponseFailListBuilder {
	b.s.Mail = &v
	return b
}

func (b *BatchResetPwdResponseFailListBuilder) Mobile(v string) *BatchResetPwdResponseFailListBuilder {
	b.s.Mobile = &v
	return b
}

func (b *BatchResetPwdResponseFailListBuilder) TenantId(v string) *BatchResetPwdResponseFailListBuilder {
	b.s.TenantId = &v
	return b
}

func (b *BatchResetPwdResponseFailListBuilder) UserId(v string) *BatchResetPwdResponseFailListBuilder {
	b.s.UserId = &v
	return b
}

func (b *BatchResetPwdResponseFailListBuilder) Build() *BatchResetPwdResponseFailList {
	return b.s
}


