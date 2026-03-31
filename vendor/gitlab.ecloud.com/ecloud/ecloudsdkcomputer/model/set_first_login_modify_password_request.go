// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type SetFirstLoginModifyPasswordRequest struct {


	SetFirstLoginModifyPasswordBody *SetFirstLoginModifyPasswordBody `json:"setFirstLoginModifyPasswordBody,omitempty"`
}

func (s SetFirstLoginModifyPasswordRequest) String() string {
	return utils.Beautify(s)
}

func (s SetFirstLoginModifyPasswordRequest) GoString() string {
	return s.String()
}

func (s SetFirstLoginModifyPasswordRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *SetFirstLoginModifyPasswordRequest) SetSetFirstLoginModifyPasswordBody(v *SetFirstLoginModifyPasswordBody) *SetFirstLoginModifyPasswordRequest {
	s.SetFirstLoginModifyPasswordBody = v
	return s
}


type SetFirstLoginModifyPasswordRequestBuilder struct {
	s *SetFirstLoginModifyPasswordRequest
}

func NewSetFirstLoginModifyPasswordRequestBuilder() *SetFirstLoginModifyPasswordRequestBuilder {
	s := &SetFirstLoginModifyPasswordRequest{}
	b := &SetFirstLoginModifyPasswordRequestBuilder{s: s}
	return b
}

func (b *SetFirstLoginModifyPasswordRequestBuilder) SetFirstLoginModifyPasswordBody(v *SetFirstLoginModifyPasswordBody) *SetFirstLoginModifyPasswordRequestBuilder {
	b.s.SetFirstLoginModifyPasswordBody = v
	return b
}

func (b *SetFirstLoginModifyPasswordRequestBuilder) Build() *SetFirstLoginModifyPasswordRequest {
	return b.s
}


