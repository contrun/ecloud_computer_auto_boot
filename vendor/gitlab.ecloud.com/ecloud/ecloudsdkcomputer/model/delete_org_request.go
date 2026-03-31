// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeleteOrgRequest struct {


	DeleteOrgBody *DeleteOrgBody `json:"deleteOrgBody,omitempty"`
}

func (s DeleteOrgRequest) String() string {
	return utils.Beautify(s)
}

func (s DeleteOrgRequest) GoString() string {
	return s.String()
}

func (s DeleteOrgRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteOrgRequest) SetDeleteOrgBody(v *DeleteOrgBody) *DeleteOrgRequest {
	s.DeleteOrgBody = v
	return s
}


type DeleteOrgRequestBuilder struct {
	s *DeleteOrgRequest
}

func NewDeleteOrgRequestBuilder() *DeleteOrgRequestBuilder {
	s := &DeleteOrgRequest{}
	b := &DeleteOrgRequestBuilder{s: s}
	return b
}

func (b *DeleteOrgRequestBuilder) DeleteOrgBody(v *DeleteOrgBody) *DeleteOrgRequestBuilder {
	b.s.DeleteOrgBody = v
	return b
}

func (b *DeleteOrgRequestBuilder) Build() *DeleteOrgRequest {
	return b.s
}


