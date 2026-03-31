// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UserLoginStatisticsRequest struct {


	UserLoginStatisticsBody *UserLoginStatisticsBody `json:"userLoginStatisticsBody,omitempty"`
}

func (s UserLoginStatisticsRequest) String() string {
	return utils.Beautify(s)
}

func (s UserLoginStatisticsRequest) GoString() string {
	return s.String()
}

func (s UserLoginStatisticsRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UserLoginStatisticsRequest) SetUserLoginStatisticsBody(v *UserLoginStatisticsBody) *UserLoginStatisticsRequest {
	s.UserLoginStatisticsBody = v
	return s
}


type UserLoginStatisticsRequestBuilder struct {
	s *UserLoginStatisticsRequest
}

func NewUserLoginStatisticsRequestBuilder() *UserLoginStatisticsRequestBuilder {
	s := &UserLoginStatisticsRequest{}
	b := &UserLoginStatisticsRequestBuilder{s: s}
	return b
}

func (b *UserLoginStatisticsRequestBuilder) UserLoginStatisticsBody(v *UserLoginStatisticsBody) *UserLoginStatisticsRequestBuilder {
	b.s.UserLoginStatisticsBody = v
	return b
}

func (b *UserLoginStatisticsRequestBuilder) Build() *UserLoginStatisticsRequest {
	return b.s
}


