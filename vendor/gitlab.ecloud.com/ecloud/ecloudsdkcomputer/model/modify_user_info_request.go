// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ModifyUserInfoRequest struct {


	ModifyUserInfoBody *ModifyUserInfoBody `json:"modifyUserInfoBody,omitempty"`
}

func (s ModifyUserInfoRequest) String() string {
	return utils.Beautify(s)
}

func (s ModifyUserInfoRequest) GoString() string {
	return s.String()
}

func (s ModifyUserInfoRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ModifyUserInfoRequest) SetModifyUserInfoBody(v *ModifyUserInfoBody) *ModifyUserInfoRequest {
	s.ModifyUserInfoBody = v
	return s
}


type ModifyUserInfoRequestBuilder struct {
	s *ModifyUserInfoRequest
}

func NewModifyUserInfoRequestBuilder() *ModifyUserInfoRequestBuilder {
	s := &ModifyUserInfoRequest{}
	b := &ModifyUserInfoRequestBuilder{s: s}
	return b
}

func (b *ModifyUserInfoRequestBuilder) ModifyUserInfoBody(v *ModifyUserInfoBody) *ModifyUserInfoRequestBuilder {
	b.s.ModifyUserInfoBody = v
	return b
}

func (b *ModifyUserInfoRequestBuilder) Build() *ModifyUserInfoRequest {
	return b.s
}


