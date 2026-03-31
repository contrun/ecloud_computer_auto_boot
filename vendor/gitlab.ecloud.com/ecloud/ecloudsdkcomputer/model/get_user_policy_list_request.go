// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetUserPolicyListRequest struct {


	GetUserPolicyListBody *GetUserPolicyListBody `json:"getUserPolicyListBody,omitempty"`
}

func (s GetUserPolicyListRequest) String() string {
	return utils.Beautify(s)
}

func (s GetUserPolicyListRequest) GoString() string {
	return s.String()
}

func (s GetUserPolicyListRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetUserPolicyListRequest) SetGetUserPolicyListBody(v *GetUserPolicyListBody) *GetUserPolicyListRequest {
	s.GetUserPolicyListBody = v
	return s
}


type GetUserPolicyListRequestBuilder struct {
	s *GetUserPolicyListRequest
}

func NewGetUserPolicyListRequestBuilder() *GetUserPolicyListRequestBuilder {
	s := &GetUserPolicyListRequest{}
	b := &GetUserPolicyListRequestBuilder{s: s}
	return b
}

func (b *GetUserPolicyListRequestBuilder) GetUserPolicyListBody(v *GetUserPolicyListBody) *GetUserPolicyListRequestBuilder {
	b.s.GetUserPolicyListBody = v
	return b
}

func (b *GetUserPolicyListRequestBuilder) Build() *GetUserPolicyListRequest {
	return b.s
}


