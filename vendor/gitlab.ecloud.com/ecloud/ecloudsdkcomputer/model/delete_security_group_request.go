// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeleteSecurityGroupRequest struct {


	DeleteSecurityGroupBody *DeleteSecurityGroupBody `json:"deleteSecurityGroupBody,omitempty"`
}

func (s DeleteSecurityGroupRequest) String() string {
	return utils.Beautify(s)
}

func (s DeleteSecurityGroupRequest) GoString() string {
	return s.String()
}

func (s DeleteSecurityGroupRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteSecurityGroupRequest) SetDeleteSecurityGroupBody(v *DeleteSecurityGroupBody) *DeleteSecurityGroupRequest {
	s.DeleteSecurityGroupBody = v
	return s
}


type DeleteSecurityGroupRequestBuilder struct {
	s *DeleteSecurityGroupRequest
}

func NewDeleteSecurityGroupRequestBuilder() *DeleteSecurityGroupRequestBuilder {
	s := &DeleteSecurityGroupRequest{}
	b := &DeleteSecurityGroupRequestBuilder{s: s}
	return b
}

func (b *DeleteSecurityGroupRequestBuilder) DeleteSecurityGroupBody(v *DeleteSecurityGroupBody) *DeleteSecurityGroupRequestBuilder {
	b.s.DeleteSecurityGroupBody = v
	return b
}

func (b *DeleteSecurityGroupRequestBuilder) Build() *DeleteSecurityGroupRequest {
	return b.s
}


