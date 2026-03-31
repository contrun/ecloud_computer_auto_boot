// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ConnectLivingPageDailyRequest struct {


	ConnectLivingPageDailyBody *ConnectLivingPageDailyBody `json:"connectLivingPageDailyBody,omitempty"`
}

func (s ConnectLivingPageDailyRequest) String() string {
	return utils.Beautify(s)
}

func (s ConnectLivingPageDailyRequest) GoString() string {
	return s.String()
}

func (s ConnectLivingPageDailyRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ConnectLivingPageDailyRequest) SetConnectLivingPageDailyBody(v *ConnectLivingPageDailyBody) *ConnectLivingPageDailyRequest {
	s.ConnectLivingPageDailyBody = v
	return s
}


type ConnectLivingPageDailyRequestBuilder struct {
	s *ConnectLivingPageDailyRequest
}

func NewConnectLivingPageDailyRequestBuilder() *ConnectLivingPageDailyRequestBuilder {
	s := &ConnectLivingPageDailyRequest{}
	b := &ConnectLivingPageDailyRequestBuilder{s: s}
	return b
}

func (b *ConnectLivingPageDailyRequestBuilder) ConnectLivingPageDailyBody(v *ConnectLivingPageDailyBody) *ConnectLivingPageDailyRequestBuilder {
	b.s.ConnectLivingPageDailyBody = v
	return b
}

func (b *ConnectLivingPageDailyRequestBuilder) Build() *ConnectLivingPageDailyRequest {
	return b.s
}


