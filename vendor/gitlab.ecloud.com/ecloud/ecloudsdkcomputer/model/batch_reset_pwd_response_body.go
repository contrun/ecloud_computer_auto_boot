// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BatchResetPwdResponseBody struct {

    // 批量重置密码成功集合
	SuccessList *[]BatchResetPwdResponseSuccessList `json:"successList,omitempty"`
    // 重置成功数量
	SuccessSize *string `json:"successSize,omitempty"`
    // 重置失败数量
	FailSize *string `json:"failSize,omitempty"`
    // 批量重置密码失败集合
	FailList *[]BatchResetPwdResponseFailList `json:"failList,omitempty"`
}

func (s BatchResetPwdResponseBody) String() string {
	return utils.Beautify(s)
}

func (s BatchResetPwdResponseBody) GoString() string {
	return s.String()
}

func (s BatchResetPwdResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BatchResetPwdResponseBody) SetSuccessList(v []BatchResetPwdResponseSuccessList) *BatchResetPwdResponseBody {
	s.SuccessList = &v
	return s
}

func (s *BatchResetPwdResponseBody) SetSuccessSize(v string) *BatchResetPwdResponseBody {
	s.SuccessSize = &v
	return s
}

func (s *BatchResetPwdResponseBody) SetFailSize(v string) *BatchResetPwdResponseBody {
	s.FailSize = &v
	return s
}

func (s *BatchResetPwdResponseBody) SetFailList(v []BatchResetPwdResponseFailList) *BatchResetPwdResponseBody {
	s.FailList = &v
	return s
}


type BatchResetPwdResponseBodyBuilder struct {
	s *BatchResetPwdResponseBody
}

func NewBatchResetPwdResponseBodyBuilder() *BatchResetPwdResponseBodyBuilder {
	s := &BatchResetPwdResponseBody{}
	b := &BatchResetPwdResponseBodyBuilder{s: s}
	return b
}

func (b *BatchResetPwdResponseBodyBuilder) SuccessList(v []BatchResetPwdResponseSuccessList) *BatchResetPwdResponseBodyBuilder {
	b.s.SuccessList = &v
	return b
}

func (b *BatchResetPwdResponseBodyBuilder) SuccessSize(v string) *BatchResetPwdResponseBodyBuilder {
	b.s.SuccessSize = &v
	return b
}

func (b *BatchResetPwdResponseBodyBuilder) FailSize(v string) *BatchResetPwdResponseBodyBuilder {
	b.s.FailSize = &v
	return b
}

func (b *BatchResetPwdResponseBodyBuilder) FailList(v []BatchResetPwdResponseFailList) *BatchResetPwdResponseBodyBuilder {
	b.s.FailList = &v
	return b
}

func (b *BatchResetPwdResponseBodyBuilder) Build() *BatchResetPwdResponseBody {
	return b.s
}


