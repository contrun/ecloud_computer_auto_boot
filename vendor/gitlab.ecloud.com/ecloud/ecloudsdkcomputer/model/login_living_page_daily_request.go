// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type LoginLivingPageDailyRequest struct {


	LoginLivingPageDailyBody *LoginLivingPageDailyBody `json:"loginLivingPageDailyBody,omitempty"`
}

func (s LoginLivingPageDailyRequest) String() string {
	return utils.Beautify(s)
}

func (s LoginLivingPageDailyRequest) GoString() string {
	return s.String()
}

func (s LoginLivingPageDailyRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *LoginLivingPageDailyRequest) SetLoginLivingPageDailyBody(v *LoginLivingPageDailyBody) *LoginLivingPageDailyRequest {
	s.LoginLivingPageDailyBody = v
	return s
}


type LoginLivingPageDailyRequestBuilder struct {
	s *LoginLivingPageDailyRequest
}

func NewLoginLivingPageDailyRequestBuilder() *LoginLivingPageDailyRequestBuilder {
	s := &LoginLivingPageDailyRequest{}
	b := &LoginLivingPageDailyRequestBuilder{s: s}
	return b
}

func (b *LoginLivingPageDailyRequestBuilder) LoginLivingPageDailyBody(v *LoginLivingPageDailyBody) *LoginLivingPageDailyRequestBuilder {
	b.s.LoginLivingPageDailyBody = v
	return b
}

func (b *LoginLivingPageDailyRequestBuilder) Build() *LoginLivingPageDailyRequest {
	return b.s
}


