// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateRequest struct {


	UpdateBody *UpdateBody `json:"updateBody,omitempty"`
}

func (s UpdateRequest) String() string {
	return utils.Beautify(s)
}

func (s UpdateRequest) GoString() string {
	return s.String()
}

func (s UpdateRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateRequest) SetUpdateBody(v *UpdateBody) *UpdateRequest {
	s.UpdateBody = v
	return s
}


type UpdateRequestBuilder struct {
	s *UpdateRequest
}

func NewUpdateRequestBuilder() *UpdateRequestBuilder {
	s := &UpdateRequest{}
	b := &UpdateRequestBuilder{s: s}
	return b
}

func (b *UpdateRequestBuilder) UpdateBody(v *UpdateBody) *UpdateRequestBuilder {
	b.s.UpdateBody = v
	return b
}

func (b *UpdateRequestBuilder) Build() *UpdateRequest {
	return b.s
}


