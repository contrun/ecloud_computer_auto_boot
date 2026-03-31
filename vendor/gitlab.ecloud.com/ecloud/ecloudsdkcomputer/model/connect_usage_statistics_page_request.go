// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ConnectUsageStatisticsPageRequest struct {


	ConnectUsageStatisticsPageBody *ConnectUsageStatisticsPageBody `json:"connectUsageStatisticsPageBody,omitempty"`
}

func (s ConnectUsageStatisticsPageRequest) String() string {
	return utils.Beautify(s)
}

func (s ConnectUsageStatisticsPageRequest) GoString() string {
	return s.String()
}

func (s ConnectUsageStatisticsPageRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ConnectUsageStatisticsPageRequest) SetConnectUsageStatisticsPageBody(v *ConnectUsageStatisticsPageBody) *ConnectUsageStatisticsPageRequest {
	s.ConnectUsageStatisticsPageBody = v
	return s
}


type ConnectUsageStatisticsPageRequestBuilder struct {
	s *ConnectUsageStatisticsPageRequest
}

func NewConnectUsageStatisticsPageRequestBuilder() *ConnectUsageStatisticsPageRequestBuilder {
	s := &ConnectUsageStatisticsPageRequest{}
	b := &ConnectUsageStatisticsPageRequestBuilder{s: s}
	return b
}

func (b *ConnectUsageStatisticsPageRequestBuilder) ConnectUsageStatisticsPageBody(v *ConnectUsageStatisticsPageBody) *ConnectUsageStatisticsPageRequestBuilder {
	b.s.ConnectUsageStatisticsPageBody = v
	return b
}

func (b *ConnectUsageStatisticsPageRequestBuilder) Build() *ConnectUsageStatisticsPageRequest {
	return b.s
}


