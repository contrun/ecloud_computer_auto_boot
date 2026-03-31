// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateSecurityGroupRequest struct {


	UpdateSecurityGroupBody *UpdateSecurityGroupBody `json:"updateSecurityGroupBody,omitempty"`
}

func (s UpdateSecurityGroupRequest) String() string {
	return utils.Beautify(s)
}

func (s UpdateSecurityGroupRequest) GoString() string {
	return s.String()
}

func (s UpdateSecurityGroupRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateSecurityGroupRequest) SetUpdateSecurityGroupBody(v *UpdateSecurityGroupBody) *UpdateSecurityGroupRequest {
	s.UpdateSecurityGroupBody = v
	return s
}


type UpdateSecurityGroupRequestBuilder struct {
	s *UpdateSecurityGroupRequest
}

func NewUpdateSecurityGroupRequestBuilder() *UpdateSecurityGroupRequestBuilder {
	s := &UpdateSecurityGroupRequest{}
	b := &UpdateSecurityGroupRequestBuilder{s: s}
	return b
}

func (b *UpdateSecurityGroupRequestBuilder) UpdateSecurityGroupBody(v *UpdateSecurityGroupBody) *UpdateSecurityGroupRequestBuilder {
	b.s.UpdateSecurityGroupBody = v
	return b
}

func (b *UpdateSecurityGroupRequestBuilder) Build() *UpdateSecurityGroupRequest {
	return b.s
}


