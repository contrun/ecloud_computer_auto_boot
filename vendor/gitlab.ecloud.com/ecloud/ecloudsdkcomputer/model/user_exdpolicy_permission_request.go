// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UserExdpolicyPermissionRequest struct {


	UserExdpolicyPermissionBody *UserExdpolicyPermissionBody `json:"userExdpolicyPermissionBody,omitempty"`
}

func (s UserExdpolicyPermissionRequest) String() string {
	return utils.Beautify(s)
}

func (s UserExdpolicyPermissionRequest) GoString() string {
	return s.String()
}

func (s UserExdpolicyPermissionRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UserExdpolicyPermissionRequest) SetUserExdpolicyPermissionBody(v *UserExdpolicyPermissionBody) *UserExdpolicyPermissionRequest {
	s.UserExdpolicyPermissionBody = v
	return s
}


type UserExdpolicyPermissionRequestBuilder struct {
	s *UserExdpolicyPermissionRequest
}

func NewUserExdpolicyPermissionRequestBuilder() *UserExdpolicyPermissionRequestBuilder {
	s := &UserExdpolicyPermissionRequest{}
	b := &UserExdpolicyPermissionRequestBuilder{s: s}
	return b
}

func (b *UserExdpolicyPermissionRequestBuilder) UserExdpolicyPermissionBody(v *UserExdpolicyPermissionBody) *UserExdpolicyPermissionRequestBuilder {
	b.s.UserExdpolicyPermissionBody = v
	return b
}

func (b *UserExdpolicyPermissionRequestBuilder) Build() *UserExdpolicyPermissionRequest {
	return b.s
}


