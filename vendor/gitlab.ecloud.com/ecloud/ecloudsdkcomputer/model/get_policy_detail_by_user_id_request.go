// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetPolicyDetailByUserIdRequest struct {


	GetPolicyDetailByUserIdQuery *GetPolicyDetailByUserIdQuery `json:"getPolicyDetailByUserIdQuery,omitempty"`
}

func (s GetPolicyDetailByUserIdRequest) String() string {
	return utils.Beautify(s)
}

func (s GetPolicyDetailByUserIdRequest) GoString() string {
	return s.String()
}

func (s GetPolicyDetailByUserIdRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetPolicyDetailByUserIdRequest) SetGetPolicyDetailByUserIdQuery(v *GetPolicyDetailByUserIdQuery) *GetPolicyDetailByUserIdRequest {
	s.GetPolicyDetailByUserIdQuery = v
	return s
}


type GetPolicyDetailByUserIdRequestBuilder struct {
	s *GetPolicyDetailByUserIdRequest
}

func NewGetPolicyDetailByUserIdRequestBuilder() *GetPolicyDetailByUserIdRequestBuilder {
	s := &GetPolicyDetailByUserIdRequest{}
	b := &GetPolicyDetailByUserIdRequestBuilder{s: s}
	return b
}

func (b *GetPolicyDetailByUserIdRequestBuilder) GetPolicyDetailByUserIdQuery(v *GetPolicyDetailByUserIdQuery) *GetPolicyDetailByUserIdRequestBuilder {
	b.s.GetPolicyDetailByUserIdQuery = v
	return b
}

func (b *GetPolicyDetailByUserIdRequestBuilder) Build() *GetPolicyDetailByUserIdRequest {
	return b.s
}


