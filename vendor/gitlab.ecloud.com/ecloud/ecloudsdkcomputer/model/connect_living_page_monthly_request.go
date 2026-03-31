// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ConnectLivingPageMonthlyRequest struct {


	ConnectLivingPageMonthlyBody *ConnectLivingPageMonthlyBody `json:"connectLivingPageMonthlyBody,omitempty"`
}

func (s ConnectLivingPageMonthlyRequest) String() string {
	return utils.Beautify(s)
}

func (s ConnectLivingPageMonthlyRequest) GoString() string {
	return s.String()
}

func (s ConnectLivingPageMonthlyRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ConnectLivingPageMonthlyRequest) SetConnectLivingPageMonthlyBody(v *ConnectLivingPageMonthlyBody) *ConnectLivingPageMonthlyRequest {
	s.ConnectLivingPageMonthlyBody = v
	return s
}


type ConnectLivingPageMonthlyRequestBuilder struct {
	s *ConnectLivingPageMonthlyRequest
}

func NewConnectLivingPageMonthlyRequestBuilder() *ConnectLivingPageMonthlyRequestBuilder {
	s := &ConnectLivingPageMonthlyRequest{}
	b := &ConnectLivingPageMonthlyRequestBuilder{s: s}
	return b
}

func (b *ConnectLivingPageMonthlyRequestBuilder) ConnectLivingPageMonthlyBody(v *ConnectLivingPageMonthlyBody) *ConnectLivingPageMonthlyRequestBuilder {
	b.s.ConnectLivingPageMonthlyBody = v
	return b
}

func (b *ConnectLivingPageMonthlyRequestBuilder) Build() *ConnectLivingPageMonthlyRequest {
	return b.s
}


