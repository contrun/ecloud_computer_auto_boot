// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ModifyUserInfoResponseBody struct {

    // 成功集合
	SuccessList *[]ModifyUserInfoResponseSuccessList `json:"successList,omitempty"`
    // 验证成功数量
	SuccessSize *string `json:"successSize,omitempty"`
    // 验证失败数量
	FailSize *string `json:"failSize,omitempty"`
    // 失败集合
	FailList *[]ModifyUserInfoResponseFailList `json:"failList,omitempty"`
}

func (s ModifyUserInfoResponseBody) String() string {
	return utils.Beautify(s)
}

func (s ModifyUserInfoResponseBody) GoString() string {
	return s.String()
}

func (s ModifyUserInfoResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ModifyUserInfoResponseBody) SetSuccessList(v []ModifyUserInfoResponseSuccessList) *ModifyUserInfoResponseBody {
	s.SuccessList = &v
	return s
}

func (s *ModifyUserInfoResponseBody) SetSuccessSize(v string) *ModifyUserInfoResponseBody {
	s.SuccessSize = &v
	return s
}

func (s *ModifyUserInfoResponseBody) SetFailSize(v string) *ModifyUserInfoResponseBody {
	s.FailSize = &v
	return s
}

func (s *ModifyUserInfoResponseBody) SetFailList(v []ModifyUserInfoResponseFailList) *ModifyUserInfoResponseBody {
	s.FailList = &v
	return s
}


type ModifyUserInfoResponseBodyBuilder struct {
	s *ModifyUserInfoResponseBody
}

func NewModifyUserInfoResponseBodyBuilder() *ModifyUserInfoResponseBodyBuilder {
	s := &ModifyUserInfoResponseBody{}
	b := &ModifyUserInfoResponseBodyBuilder{s: s}
	return b
}

func (b *ModifyUserInfoResponseBodyBuilder) SuccessList(v []ModifyUserInfoResponseSuccessList) *ModifyUserInfoResponseBodyBuilder {
	b.s.SuccessList = &v
	return b
}

func (b *ModifyUserInfoResponseBodyBuilder) SuccessSize(v string) *ModifyUserInfoResponseBodyBuilder {
	b.s.SuccessSize = &v
	return b
}

func (b *ModifyUserInfoResponseBodyBuilder) FailSize(v string) *ModifyUserInfoResponseBodyBuilder {
	b.s.FailSize = &v
	return b
}

func (b *ModifyUserInfoResponseBodyBuilder) FailList(v []ModifyUserInfoResponseFailList) *ModifyUserInfoResponseBodyBuilder {
	b.s.FailList = &v
	return b
}

func (b *ModifyUserInfoResponseBodyBuilder) Build() *ModifyUserInfoResponseBody {
	return b.s
}


