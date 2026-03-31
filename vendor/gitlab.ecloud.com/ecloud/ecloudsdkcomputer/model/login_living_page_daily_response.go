// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type LoginLivingPageDailyResponseStateEnum string

// List of State
const (
    LoginLivingPageDailyResponseStateEnumOk LoginLivingPageDailyResponseStateEnum = "OK"
    LoginLivingPageDailyResponseStateEnumError LoginLivingPageDailyResponseStateEnum = "ERROR"
    LoginLivingPageDailyResponseStateEnumException LoginLivingPageDailyResponseStateEnum = "EXCEPTION"
    LoginLivingPageDailyResponseStateEnumForbidden LoginLivingPageDailyResponseStateEnum = "FORBIDDEN"
    LoginLivingPageDailyResponseStateEnumRemaining LoginLivingPageDailyResponseStateEnum = "REMAINING"
    LoginLivingPageDailyResponseStateEnumTimeout LoginLivingPageDailyResponseStateEnum = "TIMEOUT"
)

type LoginLivingPageDailyResponse struct {


	RequestId *string `json:"requestId,omitempty"`

	ErrorMessage *string `json:"errorMessage,omitempty"`

	ErrorCode *string `json:"errorCode,omitempty"`

	State *LoginLivingPageDailyResponseStateEnum `json:"state,omitempty"`

	Body *LoginLivingPageDailyResponseBody `json:"body,omitempty"`
}

func (s LoginLivingPageDailyResponse) String() string {
	return utils.Beautify(s)
}

func (s LoginLivingPageDailyResponse) GoString() string {
	return s.String()
}

func (s LoginLivingPageDailyResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *LoginLivingPageDailyResponse) SetRequestId(v string) *LoginLivingPageDailyResponse {
	s.RequestId = &v
	return s
}

func (s *LoginLivingPageDailyResponse) SetErrorMessage(v string) *LoginLivingPageDailyResponse {
	s.ErrorMessage = &v
	return s
}

func (s *LoginLivingPageDailyResponse) SetErrorCode(v string) *LoginLivingPageDailyResponse {
	s.ErrorCode = &v
	return s
}

func (s *LoginLivingPageDailyResponse) SetState(v LoginLivingPageDailyResponseStateEnum) *LoginLivingPageDailyResponse {
	s.State = &v
	return s
}

func (s *LoginLivingPageDailyResponse) SetBody(v *LoginLivingPageDailyResponseBody) *LoginLivingPageDailyResponse {
	s.Body = v
	return s
}


type LoginLivingPageDailyResponseBuilder struct {
	s *LoginLivingPageDailyResponse
}

func NewLoginLivingPageDailyResponseBuilder() *LoginLivingPageDailyResponseBuilder {
	s := &LoginLivingPageDailyResponse{}
	b := &LoginLivingPageDailyResponseBuilder{s: s}
	return b
}

func (b *LoginLivingPageDailyResponseBuilder) RequestId(v string) *LoginLivingPageDailyResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *LoginLivingPageDailyResponseBuilder) ErrorMessage(v string) *LoginLivingPageDailyResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *LoginLivingPageDailyResponseBuilder) ErrorCode(v string) *LoginLivingPageDailyResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *LoginLivingPageDailyResponseBuilder) State(v LoginLivingPageDailyResponseStateEnum) *LoginLivingPageDailyResponseBuilder {
	b.s.State = &v
	return b
}

func (b *LoginLivingPageDailyResponseBuilder) Body(v *LoginLivingPageDailyResponseBody) *LoginLivingPageDailyResponseBuilder {
	b.s.Body = v
	return b
}

func (b *LoginLivingPageDailyResponseBuilder) Build() *LoginLivingPageDailyResponse {
	return b.s
}


