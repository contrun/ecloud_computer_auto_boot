// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UserLoginInfoRequest struct {


	UserLoginInfoBody *UserLoginInfoBody `json:"userLoginInfoBody,omitempty"`
}

func (s UserLoginInfoRequest) String() string {
	return utils.Beautify(s)
}

func (s UserLoginInfoRequest) GoString() string {
	return s.String()
}

func (s UserLoginInfoRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UserLoginInfoRequest) SetUserLoginInfoBody(v *UserLoginInfoBody) *UserLoginInfoRequest {
	s.UserLoginInfoBody = v
	return s
}


type UserLoginInfoRequestBuilder struct {
	s *UserLoginInfoRequest
}

func NewUserLoginInfoRequestBuilder() *UserLoginInfoRequestBuilder {
	s := &UserLoginInfoRequest{}
	b := &UserLoginInfoRequestBuilder{s: s}
	return b
}

func (b *UserLoginInfoRequestBuilder) UserLoginInfoBody(v *UserLoginInfoBody) *UserLoginInfoRequestBuilder {
	b.s.UserLoginInfoBody = v
	return b
}

func (b *UserLoginInfoRequestBuilder) Build() *UserLoginInfoRequest {
	return b.s
}


