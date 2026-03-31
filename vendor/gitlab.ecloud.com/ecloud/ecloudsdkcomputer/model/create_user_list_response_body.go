// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreateUserListResponseBody struct {

    // 成功集合
	SuccessList *[]CreateUserListResponseSuccessList `json:"successList,omitempty"`
    // 失败集合
	FailList *[]CreateUserListResponseFailList `json:"failList,omitempty"`
}

func (s CreateUserListResponseBody) String() string {
	return utils.Beautify(s)
}

func (s CreateUserListResponseBody) GoString() string {
	return s.String()
}

func (s CreateUserListResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateUserListResponseBody) SetSuccessList(v []CreateUserListResponseSuccessList) *CreateUserListResponseBody {
	s.SuccessList = &v
	return s
}

func (s *CreateUserListResponseBody) SetFailList(v []CreateUserListResponseFailList) *CreateUserListResponseBody {
	s.FailList = &v
	return s
}


type CreateUserListResponseBodyBuilder struct {
	s *CreateUserListResponseBody
}

func NewCreateUserListResponseBodyBuilder() *CreateUserListResponseBodyBuilder {
	s := &CreateUserListResponseBody{}
	b := &CreateUserListResponseBodyBuilder{s: s}
	return b
}

func (b *CreateUserListResponseBodyBuilder) SuccessList(v []CreateUserListResponseSuccessList) *CreateUserListResponseBodyBuilder {
	b.s.SuccessList = &v
	return b
}

func (b *CreateUserListResponseBodyBuilder) FailList(v []CreateUserListResponseFailList) *CreateUserListResponseBodyBuilder {
	b.s.FailList = &v
	return b
}

func (b *CreateUserListResponseBodyBuilder) Build() *CreateUserListResponseBody {
	return b.s
}


