// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type ConnectLivingPageMonthlyResponseStateEnum string

// List of State
const (
    ConnectLivingPageMonthlyResponseStateEnumOk ConnectLivingPageMonthlyResponseStateEnum = "OK"
    ConnectLivingPageMonthlyResponseStateEnumError ConnectLivingPageMonthlyResponseStateEnum = "ERROR"
    ConnectLivingPageMonthlyResponseStateEnumException ConnectLivingPageMonthlyResponseStateEnum = "EXCEPTION"
    ConnectLivingPageMonthlyResponseStateEnumForbidden ConnectLivingPageMonthlyResponseStateEnum = "FORBIDDEN"
    ConnectLivingPageMonthlyResponseStateEnumRemaining ConnectLivingPageMonthlyResponseStateEnum = "REMAINING"
    ConnectLivingPageMonthlyResponseStateEnumTimeout ConnectLivingPageMonthlyResponseStateEnum = "TIMEOUT"
)

type ConnectLivingPageMonthlyResponse struct {


	RequestId *string `json:"requestId,omitempty"`

	ErrorMessage *string `json:"errorMessage,omitempty"`

	ErrorCode *string `json:"errorCode,omitempty"`

	State *ConnectLivingPageMonthlyResponseStateEnum `json:"state,omitempty"`

	Body *ConnectLivingPageMonthlyResponseBody `json:"body,omitempty"`
}

func (s ConnectLivingPageMonthlyResponse) String() string {
	return utils.Beautify(s)
}

func (s ConnectLivingPageMonthlyResponse) GoString() string {
	return s.String()
}

func (s ConnectLivingPageMonthlyResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ConnectLivingPageMonthlyResponse) SetRequestId(v string) *ConnectLivingPageMonthlyResponse {
	s.RequestId = &v
	return s
}

func (s *ConnectLivingPageMonthlyResponse) SetErrorMessage(v string) *ConnectLivingPageMonthlyResponse {
	s.ErrorMessage = &v
	return s
}

func (s *ConnectLivingPageMonthlyResponse) SetErrorCode(v string) *ConnectLivingPageMonthlyResponse {
	s.ErrorCode = &v
	return s
}

func (s *ConnectLivingPageMonthlyResponse) SetState(v ConnectLivingPageMonthlyResponseStateEnum) *ConnectLivingPageMonthlyResponse {
	s.State = &v
	return s
}

func (s *ConnectLivingPageMonthlyResponse) SetBody(v *ConnectLivingPageMonthlyResponseBody) *ConnectLivingPageMonthlyResponse {
	s.Body = v
	return s
}


type ConnectLivingPageMonthlyResponseBuilder struct {
	s *ConnectLivingPageMonthlyResponse
}

func NewConnectLivingPageMonthlyResponseBuilder() *ConnectLivingPageMonthlyResponseBuilder {
	s := &ConnectLivingPageMonthlyResponse{}
	b := &ConnectLivingPageMonthlyResponseBuilder{s: s}
	return b
}

func (b *ConnectLivingPageMonthlyResponseBuilder) RequestId(v string) *ConnectLivingPageMonthlyResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *ConnectLivingPageMonthlyResponseBuilder) ErrorMessage(v string) *ConnectLivingPageMonthlyResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *ConnectLivingPageMonthlyResponseBuilder) ErrorCode(v string) *ConnectLivingPageMonthlyResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *ConnectLivingPageMonthlyResponseBuilder) State(v ConnectLivingPageMonthlyResponseStateEnum) *ConnectLivingPageMonthlyResponseBuilder {
	b.s.State = &v
	return b
}

func (b *ConnectLivingPageMonthlyResponseBuilder) Body(v *ConnectLivingPageMonthlyResponseBody) *ConnectLivingPageMonthlyResponseBuilder {
	b.s.Body = v
	return b
}

func (b *ConnectLivingPageMonthlyResponseBuilder) Build() *ConnectLivingPageMonthlyResponse {
	return b.s
}


