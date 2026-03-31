// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type SetUserModifyPasswordRequest struct {


	SetUserModifyPasswordBody *SetUserModifyPasswordBody `json:"setUserModifyPasswordBody,omitempty"`
}

func (s SetUserModifyPasswordRequest) String() string {
	return utils.Beautify(s)
}

func (s SetUserModifyPasswordRequest) GoString() string {
	return s.String()
}

func (s SetUserModifyPasswordRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *SetUserModifyPasswordRequest) SetSetUserModifyPasswordBody(v *SetUserModifyPasswordBody) *SetUserModifyPasswordRequest {
	s.SetUserModifyPasswordBody = v
	return s
}


type SetUserModifyPasswordRequestBuilder struct {
	s *SetUserModifyPasswordRequest
}

func NewSetUserModifyPasswordRequestBuilder() *SetUserModifyPasswordRequestBuilder {
	s := &SetUserModifyPasswordRequest{}
	b := &SetUserModifyPasswordRequestBuilder{s: s}
	return b
}

func (b *SetUserModifyPasswordRequestBuilder) SetUserModifyPasswordBody(v *SetUserModifyPasswordBody) *SetUserModifyPasswordRequestBuilder {
	b.s.SetUserModifyPasswordBody = v
	return b
}

func (b *SetUserModifyPasswordRequestBuilder) Build() *SetUserModifyPasswordRequest {
	return b.s
}


