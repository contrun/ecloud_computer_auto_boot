// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type LogoutRequest struct {


	LogoutBody *LogoutBody `json:"logoutBody,omitempty"`
}

func (s LogoutRequest) String() string {
	return utils.Beautify(s)
}

func (s LogoutRequest) GoString() string {
	return s.String()
}

func (s LogoutRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *LogoutRequest) SetLogoutBody(v *LogoutBody) *LogoutRequest {
	s.LogoutBody = v
	return s
}


type LogoutRequestBuilder struct {
	s *LogoutRequest
}

func NewLogoutRequestBuilder() *LogoutRequestBuilder {
	s := &LogoutRequest{}
	b := &LogoutRequestBuilder{s: s}
	return b
}

func (b *LogoutRequestBuilder) LogoutBody(v *LogoutBody) *LogoutRequestBuilder {
	b.s.LogoutBody = v
	return b
}

func (b *LogoutRequestBuilder) Build() *LogoutRequest {
	return b.s
}


