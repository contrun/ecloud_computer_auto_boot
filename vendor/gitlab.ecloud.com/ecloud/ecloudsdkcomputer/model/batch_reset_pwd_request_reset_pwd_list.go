// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BatchResetPwdRequestResetPwdList struct {

    // 密码
	Password *string `json:"password,omitempty"`
    // 云电脑用户id
	UserId *string `json:"userId,omitempty"`
}

func (s BatchResetPwdRequestResetPwdList) String() string {
	return utils.Beautify(s)
}

func (s BatchResetPwdRequestResetPwdList) GoString() string {
	return s.String()
}

func (s BatchResetPwdRequestResetPwdList) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BatchResetPwdRequestResetPwdList) SetPassword(v string) *BatchResetPwdRequestResetPwdList {
	s.Password = &v
	return s
}

func (s *BatchResetPwdRequestResetPwdList) SetUserId(v string) *BatchResetPwdRequestResetPwdList {
	s.UserId = &v
	return s
}


type BatchResetPwdRequestResetPwdListBuilder struct {
	s *BatchResetPwdRequestResetPwdList
}

func NewBatchResetPwdRequestResetPwdListBuilder() *BatchResetPwdRequestResetPwdListBuilder {
	s := &BatchResetPwdRequestResetPwdList{}
	b := &BatchResetPwdRequestResetPwdListBuilder{s: s}
	return b
}

func (b *BatchResetPwdRequestResetPwdListBuilder) Password(v string) *BatchResetPwdRequestResetPwdListBuilder {
	b.s.Password = &v
	return b
}

func (b *BatchResetPwdRequestResetPwdListBuilder) UserId(v string) *BatchResetPwdRequestResetPwdListBuilder {
	b.s.UserId = &v
	return b
}

func (b *BatchResetPwdRequestResetPwdListBuilder) Build() *BatchResetPwdRequestResetPwdList {
	return b.s
}


