// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type StopAutoRenewRequest struct {


	StopAutoRenewBody *StopAutoRenewBody `json:"stopAutoRenewBody,omitempty"`
}

func (s StopAutoRenewRequest) String() string {
	return utils.Beautify(s)
}

func (s StopAutoRenewRequest) GoString() string {
	return s.String()
}

func (s StopAutoRenewRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *StopAutoRenewRequest) SetStopAutoRenewBody(v *StopAutoRenewBody) *StopAutoRenewRequest {
	s.StopAutoRenewBody = v
	return s
}


type StopAutoRenewRequestBuilder struct {
	s *StopAutoRenewRequest
}

func NewStopAutoRenewRequestBuilder() *StopAutoRenewRequestBuilder {
	s := &StopAutoRenewRequest{}
	b := &StopAutoRenewRequestBuilder{s: s}
	return b
}

func (b *StopAutoRenewRequestBuilder) StopAutoRenewBody(v *StopAutoRenewBody) *StopAutoRenewRequestBuilder {
	b.s.StopAutoRenewBody = v
	return b
}

func (b *StopAutoRenewRequestBuilder) Build() *StopAutoRenewRequest {
	return b.s
}


