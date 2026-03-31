// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type LoginUsageStatisticsPageRequest struct {


	LoginUsageStatisticsPageBody *LoginUsageStatisticsPageBody `json:"loginUsageStatisticsPageBody,omitempty"`
}

func (s LoginUsageStatisticsPageRequest) String() string {
	return utils.Beautify(s)
}

func (s LoginUsageStatisticsPageRequest) GoString() string {
	return s.String()
}

func (s LoginUsageStatisticsPageRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *LoginUsageStatisticsPageRequest) SetLoginUsageStatisticsPageBody(v *LoginUsageStatisticsPageBody) *LoginUsageStatisticsPageRequest {
	s.LoginUsageStatisticsPageBody = v
	return s
}


type LoginUsageStatisticsPageRequestBuilder struct {
	s *LoginUsageStatisticsPageRequest
}

func NewLoginUsageStatisticsPageRequestBuilder() *LoginUsageStatisticsPageRequestBuilder {
	s := &LoginUsageStatisticsPageRequest{}
	b := &LoginUsageStatisticsPageRequestBuilder{s: s}
	return b
}

func (b *LoginUsageStatisticsPageRequestBuilder) LoginUsageStatisticsPageBody(v *LoginUsageStatisticsPageBody) *LoginUsageStatisticsPageRequestBuilder {
	b.s.LoginUsageStatisticsPageBody = v
	return b
}

func (b *LoginUsageStatisticsPageRequestBuilder) Build() *LoginUsageStatisticsPageRequest {
	return b.s
}


