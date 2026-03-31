// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type EditOrgRequest struct {


	EditOrgBody *EditOrgBody `json:"editOrgBody,omitempty"`
}

func (s EditOrgRequest) String() string {
	return utils.Beautify(s)
}

func (s EditOrgRequest) GoString() string {
	return s.String()
}

func (s EditOrgRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *EditOrgRequest) SetEditOrgBody(v *EditOrgBody) *EditOrgRequest {
	s.EditOrgBody = v
	return s
}


type EditOrgRequestBuilder struct {
	s *EditOrgRequest
}

func NewEditOrgRequestBuilder() *EditOrgRequestBuilder {
	s := &EditOrgRequest{}
	b := &EditOrgRequestBuilder{s: s}
	return b
}

func (b *EditOrgRequestBuilder) EditOrgBody(v *EditOrgBody) *EditOrgRequestBuilder {
	b.s.EditOrgBody = v
	return b
}

func (b *EditOrgRequestBuilder) Build() *EditOrgRequest {
	return b.s
}


