// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type SecurityGroupManagementListRequest struct {


	SecurityGroupManagementListBody *SecurityGroupManagementListBody `json:"securityGroupManagementListBody,omitempty"`
}

func (s SecurityGroupManagementListRequest) String() string {
	return utils.Beautify(s)
}

func (s SecurityGroupManagementListRequest) GoString() string {
	return s.String()
}

func (s SecurityGroupManagementListRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *SecurityGroupManagementListRequest) SetSecurityGroupManagementListBody(v *SecurityGroupManagementListBody) *SecurityGroupManagementListRequest {
	s.SecurityGroupManagementListBody = v
	return s
}


type SecurityGroupManagementListRequestBuilder struct {
	s *SecurityGroupManagementListRequest
}

func NewSecurityGroupManagementListRequestBuilder() *SecurityGroupManagementListRequestBuilder {
	s := &SecurityGroupManagementListRequest{}
	b := &SecurityGroupManagementListRequestBuilder{s: s}
	return b
}

func (b *SecurityGroupManagementListRequestBuilder) SecurityGroupManagementListBody(v *SecurityGroupManagementListBody) *SecurityGroupManagementListRequestBuilder {
	b.s.SecurityGroupManagementListBody = v
	return b
}

func (b *SecurityGroupManagementListRequestBuilder) Build() *SecurityGroupManagementListRequest {
	return b.s
}


