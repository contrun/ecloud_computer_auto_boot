// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetPolicyInfoByIdRequest struct {


	GetPolicyInfoByIdQuery *GetPolicyInfoByIdQuery `json:"getPolicyInfoByIdQuery,omitempty"`
}

func (s GetPolicyInfoByIdRequest) String() string {
	return utils.Beautify(s)
}

func (s GetPolicyInfoByIdRequest) GoString() string {
	return s.String()
}

func (s GetPolicyInfoByIdRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetPolicyInfoByIdRequest) SetGetPolicyInfoByIdQuery(v *GetPolicyInfoByIdQuery) *GetPolicyInfoByIdRequest {
	s.GetPolicyInfoByIdQuery = v
	return s
}


type GetPolicyInfoByIdRequestBuilder struct {
	s *GetPolicyInfoByIdRequest
}

func NewGetPolicyInfoByIdRequestBuilder() *GetPolicyInfoByIdRequestBuilder {
	s := &GetPolicyInfoByIdRequest{}
	b := &GetPolicyInfoByIdRequestBuilder{s: s}
	return b
}

func (b *GetPolicyInfoByIdRequestBuilder) GetPolicyInfoByIdQuery(v *GetPolicyInfoByIdQuery) *GetPolicyInfoByIdRequestBuilder {
	b.s.GetPolicyInfoByIdQuery = v
	return b
}

func (b *GetPolicyInfoByIdRequestBuilder) Build() *GetPolicyInfoByIdRequest {
	return b.s
}


