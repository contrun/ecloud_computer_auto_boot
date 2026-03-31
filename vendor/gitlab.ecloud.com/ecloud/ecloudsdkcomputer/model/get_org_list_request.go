// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetOrgListRequest struct {


	GetOrgListBody *GetOrgListBody `json:"getOrgListBody,omitempty"`
}

func (s GetOrgListRequest) String() string {
	return utils.Beautify(s)
}

func (s GetOrgListRequest) GoString() string {
	return s.String()
}

func (s GetOrgListRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetOrgListRequest) SetGetOrgListBody(v *GetOrgListBody) *GetOrgListRequest {
	s.GetOrgListBody = v
	return s
}


type GetOrgListRequestBuilder struct {
	s *GetOrgListRequest
}

func NewGetOrgListRequestBuilder() *GetOrgListRequestBuilder {
	s := &GetOrgListRequest{}
	b := &GetOrgListRequestBuilder{s: s}
	return b
}

func (b *GetOrgListRequestBuilder) GetOrgListBody(v *GetOrgListBody) *GetOrgListRequestBuilder {
	b.s.GetOrgListBody = v
	return b
}

func (b *GetOrgListRequestBuilder) Build() *GetOrgListRequest {
	return b.s
}


