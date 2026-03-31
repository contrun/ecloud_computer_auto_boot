// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetPolicyInfoByInstanceIdRequest struct {


	GetPolicyInfoByInstanceIdQuery *GetPolicyInfoByInstanceIdQuery `json:"getPolicyInfoByInstanceIdQuery,omitempty"`
}

func (s GetPolicyInfoByInstanceIdRequest) String() string {
	return utils.Beautify(s)
}

func (s GetPolicyInfoByInstanceIdRequest) GoString() string {
	return s.String()
}

func (s GetPolicyInfoByInstanceIdRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetPolicyInfoByInstanceIdRequest) SetGetPolicyInfoByInstanceIdQuery(v *GetPolicyInfoByInstanceIdQuery) *GetPolicyInfoByInstanceIdRequest {
	s.GetPolicyInfoByInstanceIdQuery = v
	return s
}


type GetPolicyInfoByInstanceIdRequestBuilder struct {
	s *GetPolicyInfoByInstanceIdRequest
}

func NewGetPolicyInfoByInstanceIdRequestBuilder() *GetPolicyInfoByInstanceIdRequestBuilder {
	s := &GetPolicyInfoByInstanceIdRequest{}
	b := &GetPolicyInfoByInstanceIdRequestBuilder{s: s}
	return b
}

func (b *GetPolicyInfoByInstanceIdRequestBuilder) GetPolicyInfoByInstanceIdQuery(v *GetPolicyInfoByInstanceIdQuery) *GetPolicyInfoByInstanceIdRequestBuilder {
	b.s.GetPolicyInfoByInstanceIdQuery = v
	return b
}

func (b *GetPolicyInfoByInstanceIdRequestBuilder) Build() *GetPolicyInfoByInstanceIdRequest {
	return b.s
}


