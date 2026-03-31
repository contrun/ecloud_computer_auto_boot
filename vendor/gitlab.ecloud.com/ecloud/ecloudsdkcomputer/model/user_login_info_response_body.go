// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UserLoginInfoResponseBody struct {

    // 成功集合
	Records *[]UserLoginInfoResponseRecords `json:"records,omitempty"`
}

func (s UserLoginInfoResponseBody) String() string {
	return utils.Beautify(s)
}

func (s UserLoginInfoResponseBody) GoString() string {
	return s.String()
}

func (s UserLoginInfoResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UserLoginInfoResponseBody) SetRecords(v []UserLoginInfoResponseRecords) *UserLoginInfoResponseBody {
	s.Records = &v
	return s
}


type UserLoginInfoResponseBodyBuilder struct {
	s *UserLoginInfoResponseBody
}

func NewUserLoginInfoResponseBodyBuilder() *UserLoginInfoResponseBodyBuilder {
	s := &UserLoginInfoResponseBody{}
	b := &UserLoginInfoResponseBodyBuilder{s: s}
	return b
}

func (b *UserLoginInfoResponseBodyBuilder) Records(v []UserLoginInfoResponseRecords) *UserLoginInfoResponseBodyBuilder {
	b.s.Records = &v
	return b
}

func (b *UserLoginInfoResponseBodyBuilder) Build() *UserLoginInfoResponseBody {
	return b.s
}


