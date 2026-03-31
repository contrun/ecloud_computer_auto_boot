// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type StartOrStopAutoRenewBwRequest struct {


	StartOrStopAutoRenewBwBody *StartOrStopAutoRenewBwBody `json:"startOrStopAutoRenewBwBody,omitempty"`
}

func (s StartOrStopAutoRenewBwRequest) String() string {
	return utils.Beautify(s)
}

func (s StartOrStopAutoRenewBwRequest) GoString() string {
	return s.String()
}

func (s StartOrStopAutoRenewBwRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *StartOrStopAutoRenewBwRequest) SetStartOrStopAutoRenewBwBody(v *StartOrStopAutoRenewBwBody) *StartOrStopAutoRenewBwRequest {
	s.StartOrStopAutoRenewBwBody = v
	return s
}


type StartOrStopAutoRenewBwRequestBuilder struct {
	s *StartOrStopAutoRenewBwRequest
}

func NewStartOrStopAutoRenewBwRequestBuilder() *StartOrStopAutoRenewBwRequestBuilder {
	s := &StartOrStopAutoRenewBwRequest{}
	b := &StartOrStopAutoRenewBwRequestBuilder{s: s}
	return b
}

func (b *StartOrStopAutoRenewBwRequestBuilder) StartOrStopAutoRenewBwBody(v *StartOrStopAutoRenewBwBody) *StartOrStopAutoRenewBwRequestBuilder {
	b.s.StartOrStopAutoRenewBwBody = v
	return b
}

func (b *StartOrStopAutoRenewBwRequestBuilder) Build() *StartOrStopAutoRenewBwRequest {
	return b.s
}


