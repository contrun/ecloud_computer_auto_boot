// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CopySecurityGroupRequest struct {


	CopySecurityGroupBody *CopySecurityGroupBody `json:"copySecurityGroupBody,omitempty"`
}

func (s CopySecurityGroupRequest) String() string {
	return utils.Beautify(s)
}

func (s CopySecurityGroupRequest) GoString() string {
	return s.String()
}

func (s CopySecurityGroupRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CopySecurityGroupRequest) SetCopySecurityGroupBody(v *CopySecurityGroupBody) *CopySecurityGroupRequest {
	s.CopySecurityGroupBody = v
	return s
}


type CopySecurityGroupRequestBuilder struct {
	s *CopySecurityGroupRequest
}

func NewCopySecurityGroupRequestBuilder() *CopySecurityGroupRequestBuilder {
	s := &CopySecurityGroupRequest{}
	b := &CopySecurityGroupRequestBuilder{s: s}
	return b
}

func (b *CopySecurityGroupRequestBuilder) CopySecurityGroupBody(v *CopySecurityGroupBody) *CopySecurityGroupRequestBuilder {
	b.s.CopySecurityGroupBody = v
	return b
}

func (b *CopySecurityGroupRequestBuilder) Build() *CopySecurityGroupRequest {
	return b.s
}


