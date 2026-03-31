// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetUserDetailByPolicyIdRequest struct {


	GetUserDetailByPolicyIdBody *GetUserDetailByPolicyIdBody `json:"getUserDetailByPolicyIdBody,omitempty"`
}

func (s GetUserDetailByPolicyIdRequest) String() string {
	return utils.Beautify(s)
}

func (s GetUserDetailByPolicyIdRequest) GoString() string {
	return s.String()
}

func (s GetUserDetailByPolicyIdRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetUserDetailByPolicyIdRequest) SetGetUserDetailByPolicyIdBody(v *GetUserDetailByPolicyIdBody) *GetUserDetailByPolicyIdRequest {
	s.GetUserDetailByPolicyIdBody = v
	return s
}


type GetUserDetailByPolicyIdRequestBuilder struct {
	s *GetUserDetailByPolicyIdRequest
}

func NewGetUserDetailByPolicyIdRequestBuilder() *GetUserDetailByPolicyIdRequestBuilder {
	s := &GetUserDetailByPolicyIdRequest{}
	b := &GetUserDetailByPolicyIdRequestBuilder{s: s}
	return b
}

func (b *GetUserDetailByPolicyIdRequestBuilder) GetUserDetailByPolicyIdBody(v *GetUserDetailByPolicyIdBody) *GetUserDetailByPolicyIdRequestBuilder {
	b.s.GetUserDetailByPolicyIdBody = v
	return b
}

func (b *GetUserDetailByPolicyIdRequestBuilder) Build() *GetUserDetailByPolicyIdRequest {
	return b.s
}


