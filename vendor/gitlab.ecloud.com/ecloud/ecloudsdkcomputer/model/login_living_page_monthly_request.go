// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type LoginLivingPageMonthlyRequest struct {


	LoginLivingPageMonthlyBody *LoginLivingPageMonthlyBody `json:"loginLivingPageMonthlyBody,omitempty"`
}

func (s LoginLivingPageMonthlyRequest) String() string {
	return utils.Beautify(s)
}

func (s LoginLivingPageMonthlyRequest) GoString() string {
	return s.String()
}

func (s LoginLivingPageMonthlyRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *LoginLivingPageMonthlyRequest) SetLoginLivingPageMonthlyBody(v *LoginLivingPageMonthlyBody) *LoginLivingPageMonthlyRequest {
	s.LoginLivingPageMonthlyBody = v
	return s
}


type LoginLivingPageMonthlyRequestBuilder struct {
	s *LoginLivingPageMonthlyRequest
}

func NewLoginLivingPageMonthlyRequestBuilder() *LoginLivingPageMonthlyRequestBuilder {
	s := &LoginLivingPageMonthlyRequest{}
	b := &LoginLivingPageMonthlyRequestBuilder{s: s}
	return b
}

func (b *LoginLivingPageMonthlyRequestBuilder) LoginLivingPageMonthlyBody(v *LoginLivingPageMonthlyBody) *LoginLivingPageMonthlyRequestBuilder {
	b.s.LoginLivingPageMonthlyBody = v
	return b
}

func (b *LoginLivingPageMonthlyRequestBuilder) Build() *LoginLivingPageMonthlyRequest {
	return b.s
}


