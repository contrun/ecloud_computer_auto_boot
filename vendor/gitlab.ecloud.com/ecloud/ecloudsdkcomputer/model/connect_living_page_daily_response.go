// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type ConnectLivingPageDailyResponseStateEnum string

// List of State
const (
    ConnectLivingPageDailyResponseStateEnumOk ConnectLivingPageDailyResponseStateEnum = "OK"
    ConnectLivingPageDailyResponseStateEnumError ConnectLivingPageDailyResponseStateEnum = "ERROR"
    ConnectLivingPageDailyResponseStateEnumException ConnectLivingPageDailyResponseStateEnum = "EXCEPTION"
    ConnectLivingPageDailyResponseStateEnumForbidden ConnectLivingPageDailyResponseStateEnum = "FORBIDDEN"
    ConnectLivingPageDailyResponseStateEnumRemaining ConnectLivingPageDailyResponseStateEnum = "REMAINING"
    ConnectLivingPageDailyResponseStateEnumTimeout ConnectLivingPageDailyResponseStateEnum = "TIMEOUT"
)

type ConnectLivingPageDailyResponse struct {


	RequestId *string `json:"requestId,omitempty"`

	ErrorMessage *string `json:"errorMessage,omitempty"`

	ErrorCode *string `json:"errorCode,omitempty"`

	State *ConnectLivingPageDailyResponseStateEnum `json:"state,omitempty"`

	Body *ConnectLivingPageDailyResponseBody `json:"body,omitempty"`
}

func (s ConnectLivingPageDailyResponse) String() string {
	return utils.Beautify(s)
}

func (s ConnectLivingPageDailyResponse) GoString() string {
	return s.String()
}

func (s ConnectLivingPageDailyResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ConnectLivingPageDailyResponse) SetRequestId(v string) *ConnectLivingPageDailyResponse {
	s.RequestId = &v
	return s
}

func (s *ConnectLivingPageDailyResponse) SetErrorMessage(v string) *ConnectLivingPageDailyResponse {
	s.ErrorMessage = &v
	return s
}

func (s *ConnectLivingPageDailyResponse) SetErrorCode(v string) *ConnectLivingPageDailyResponse {
	s.ErrorCode = &v
	return s
}

func (s *ConnectLivingPageDailyResponse) SetState(v ConnectLivingPageDailyResponseStateEnum) *ConnectLivingPageDailyResponse {
	s.State = &v
	return s
}

func (s *ConnectLivingPageDailyResponse) SetBody(v *ConnectLivingPageDailyResponseBody) *ConnectLivingPageDailyResponse {
	s.Body = v
	return s
}


type ConnectLivingPageDailyResponseBuilder struct {
	s *ConnectLivingPageDailyResponse
}

func NewConnectLivingPageDailyResponseBuilder() *ConnectLivingPageDailyResponseBuilder {
	s := &ConnectLivingPageDailyResponse{}
	b := &ConnectLivingPageDailyResponseBuilder{s: s}
	return b
}

func (b *ConnectLivingPageDailyResponseBuilder) RequestId(v string) *ConnectLivingPageDailyResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *ConnectLivingPageDailyResponseBuilder) ErrorMessage(v string) *ConnectLivingPageDailyResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *ConnectLivingPageDailyResponseBuilder) ErrorCode(v string) *ConnectLivingPageDailyResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *ConnectLivingPageDailyResponseBuilder) State(v ConnectLivingPageDailyResponseStateEnum) *ConnectLivingPageDailyResponseBuilder {
	b.s.State = &v
	return b
}

func (b *ConnectLivingPageDailyResponseBuilder) Body(v *ConnectLivingPageDailyResponseBody) *ConnectLivingPageDailyResponseBuilder {
	b.s.Body = v
	return b
}

func (b *ConnectLivingPageDailyResponseBuilder) Build() *ConnectLivingPageDailyResponse {
	return b.s
}


