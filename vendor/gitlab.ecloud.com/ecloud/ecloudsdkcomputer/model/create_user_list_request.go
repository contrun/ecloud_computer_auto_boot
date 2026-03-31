// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreateUserListRequest struct {


	CreateUserListBody *CreateUserListBody `json:"createUserListBody,omitempty"`
}

func (s CreateUserListRequest) String() string {
	return utils.Beautify(s)
}

func (s CreateUserListRequest) GoString() string {
	return s.String()
}

func (s CreateUserListRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateUserListRequest) SetCreateUserListBody(v *CreateUserListBody) *CreateUserListRequest {
	s.CreateUserListBody = v
	return s
}


type CreateUserListRequestBuilder struct {
	s *CreateUserListRequest
}

func NewCreateUserListRequestBuilder() *CreateUserListRequestBuilder {
	s := &CreateUserListRequest{}
	b := &CreateUserListRequestBuilder{s: s}
	return b
}

func (b *CreateUserListRequestBuilder) CreateUserListBody(v *CreateUserListBody) *CreateUserListRequestBuilder {
	b.s.CreateUserListBody = v
	return b
}

func (b *CreateUserListRequestBuilder) Build() *CreateUserListRequest {
	return b.s
}


