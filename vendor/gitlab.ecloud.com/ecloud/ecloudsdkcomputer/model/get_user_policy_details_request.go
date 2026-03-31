// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetUserPolicyDetailsRequest struct {


	GetUserPolicyDetailsQuery *GetUserPolicyDetailsQuery `json:"getUserPolicyDetailsQuery,omitempty"`
}

func (s GetUserPolicyDetailsRequest) String() string {
	return utils.Beautify(s)
}

func (s GetUserPolicyDetailsRequest) GoString() string {
	return s.String()
}

func (s GetUserPolicyDetailsRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetUserPolicyDetailsRequest) SetGetUserPolicyDetailsQuery(v *GetUserPolicyDetailsQuery) *GetUserPolicyDetailsRequest {
	s.GetUserPolicyDetailsQuery = v
	return s
}


type GetUserPolicyDetailsRequestBuilder struct {
	s *GetUserPolicyDetailsRequest
}

func NewGetUserPolicyDetailsRequestBuilder() *GetUserPolicyDetailsRequestBuilder {
	s := &GetUserPolicyDetailsRequest{}
	b := &GetUserPolicyDetailsRequestBuilder{s: s}
	return b
}

func (b *GetUserPolicyDetailsRequestBuilder) GetUserPolicyDetailsQuery(v *GetUserPolicyDetailsQuery) *GetUserPolicyDetailsRequestBuilder {
	b.s.GetUserPolicyDetailsQuery = v
	return b
}

func (b *GetUserPolicyDetailsRequestBuilder) Build() *GetUserPolicyDetailsRequest {
	return b.s
}


