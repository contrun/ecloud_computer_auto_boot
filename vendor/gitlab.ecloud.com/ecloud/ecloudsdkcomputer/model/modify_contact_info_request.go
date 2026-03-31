// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ModifyContactInfoRequest struct {


	ModifyContactInfoBody *ModifyContactInfoBody `json:"modifyContactInfoBody,omitempty"`
}

func (s ModifyContactInfoRequest) String() string {
	return utils.Beautify(s)
}

func (s ModifyContactInfoRequest) GoString() string {
	return s.String()
}

func (s ModifyContactInfoRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ModifyContactInfoRequest) SetModifyContactInfoBody(v *ModifyContactInfoBody) *ModifyContactInfoRequest {
	s.ModifyContactInfoBody = v
	return s
}


type ModifyContactInfoRequestBuilder struct {
	s *ModifyContactInfoRequest
}

func NewModifyContactInfoRequestBuilder() *ModifyContactInfoRequestBuilder {
	s := &ModifyContactInfoRequest{}
	b := &ModifyContactInfoRequestBuilder{s: s}
	return b
}

func (b *ModifyContactInfoRequestBuilder) ModifyContactInfoBody(v *ModifyContactInfoBody) *ModifyContactInfoRequestBuilder {
	b.s.ModifyContactInfoBody = v
	return b
}

func (b *ModifyContactInfoRequestBuilder) Build() *ModifyContactInfoRequest {
	return b.s
}


