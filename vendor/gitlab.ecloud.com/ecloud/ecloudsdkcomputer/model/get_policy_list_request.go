// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetPolicyListRequest struct {


	GetPolicyListBody *GetPolicyListBody `json:"getPolicyListBody,omitempty"`
}

func (s GetPolicyListRequest) String() string {
	return utils.Beautify(s)
}

func (s GetPolicyListRequest) GoString() string {
	return s.String()
}

func (s GetPolicyListRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetPolicyListRequest) SetGetPolicyListBody(v *GetPolicyListBody) *GetPolicyListRequest {
	s.GetPolicyListBody = v
	return s
}


type GetPolicyListRequestBuilder struct {
	s *GetPolicyListRequest
}

func NewGetPolicyListRequestBuilder() *GetPolicyListRequestBuilder {
	s := &GetPolicyListRequest{}
	b := &GetPolicyListRequestBuilder{s: s}
	return b
}

func (b *GetPolicyListRequestBuilder) GetPolicyListBody(v *GetPolicyListBody) *GetPolicyListRequestBuilder {
	b.s.GetPolicyListBody = v
	return b
}

func (b *GetPolicyListRequestBuilder) Build() *GetPolicyListRequest {
	return b.s
}


