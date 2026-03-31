// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreateSecurityGroupRequest struct {


	CreateSecurityGroupBody *CreateSecurityGroupBody `json:"createSecurityGroupBody,omitempty"`
}

func (s CreateSecurityGroupRequest) String() string {
	return utils.Beautify(s)
}

func (s CreateSecurityGroupRequest) GoString() string {
	return s.String()
}

func (s CreateSecurityGroupRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateSecurityGroupRequest) SetCreateSecurityGroupBody(v *CreateSecurityGroupBody) *CreateSecurityGroupRequest {
	s.CreateSecurityGroupBody = v
	return s
}


type CreateSecurityGroupRequestBuilder struct {
	s *CreateSecurityGroupRequest
}

func NewCreateSecurityGroupRequestBuilder() *CreateSecurityGroupRequestBuilder {
	s := &CreateSecurityGroupRequest{}
	b := &CreateSecurityGroupRequestBuilder{s: s}
	return b
}

func (b *CreateSecurityGroupRequestBuilder) CreateSecurityGroupBody(v *CreateSecurityGroupBody) *CreateSecurityGroupRequestBuilder {
	b.s.CreateSecurityGroupBody = v
	return b
}

func (b *CreateSecurityGroupRequestBuilder) Build() *CreateSecurityGroupRequest {
	return b.s
}


