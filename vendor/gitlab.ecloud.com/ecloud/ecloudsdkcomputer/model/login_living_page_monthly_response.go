// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type LoginLivingPageMonthlyResponseStateEnum string

// List of State
const (
    LoginLivingPageMonthlyResponseStateEnumOk LoginLivingPageMonthlyResponseStateEnum = "OK"
    LoginLivingPageMonthlyResponseStateEnumError LoginLivingPageMonthlyResponseStateEnum = "ERROR"
    LoginLivingPageMonthlyResponseStateEnumException LoginLivingPageMonthlyResponseStateEnum = "EXCEPTION"
    LoginLivingPageMonthlyResponseStateEnumForbidden LoginLivingPageMonthlyResponseStateEnum = "FORBIDDEN"
    LoginLivingPageMonthlyResponseStateEnumRemaining LoginLivingPageMonthlyResponseStateEnum = "REMAINING"
    LoginLivingPageMonthlyResponseStateEnumTimeout LoginLivingPageMonthlyResponseStateEnum = "TIMEOUT"
)

type LoginLivingPageMonthlyResponse struct {


	RequestId *string `json:"requestId,omitempty"`

	ErrorMessage *string `json:"errorMessage,omitempty"`

	ErrorCode *string `json:"errorCode,omitempty"`

	State *LoginLivingPageMonthlyResponseStateEnum `json:"state,omitempty"`

	Body *LoginLivingPageMonthlyResponseBody `json:"body,omitempty"`
}

func (s LoginLivingPageMonthlyResponse) String() string {
	return utils.Beautify(s)
}

func (s LoginLivingPageMonthlyResponse) GoString() string {
	return s.String()
}

func (s LoginLivingPageMonthlyResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *LoginLivingPageMonthlyResponse) SetRequestId(v string) *LoginLivingPageMonthlyResponse {
	s.RequestId = &v
	return s
}

func (s *LoginLivingPageMonthlyResponse) SetErrorMessage(v string) *LoginLivingPageMonthlyResponse {
	s.ErrorMessage = &v
	return s
}

func (s *LoginLivingPageMonthlyResponse) SetErrorCode(v string) *LoginLivingPageMonthlyResponse {
	s.ErrorCode = &v
	return s
}

func (s *LoginLivingPageMonthlyResponse) SetState(v LoginLivingPageMonthlyResponseStateEnum) *LoginLivingPageMonthlyResponse {
	s.State = &v
	return s
}

func (s *LoginLivingPageMonthlyResponse) SetBody(v *LoginLivingPageMonthlyResponseBody) *LoginLivingPageMonthlyResponse {
	s.Body = v
	return s
}


type LoginLivingPageMonthlyResponseBuilder struct {
	s *LoginLivingPageMonthlyResponse
}

func NewLoginLivingPageMonthlyResponseBuilder() *LoginLivingPageMonthlyResponseBuilder {
	s := &LoginLivingPageMonthlyResponse{}
	b := &LoginLivingPageMonthlyResponseBuilder{s: s}
	return b
}

func (b *LoginLivingPageMonthlyResponseBuilder) RequestId(v string) *LoginLivingPageMonthlyResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *LoginLivingPageMonthlyResponseBuilder) ErrorMessage(v string) *LoginLivingPageMonthlyResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *LoginLivingPageMonthlyResponseBuilder) ErrorCode(v string) *LoginLivingPageMonthlyResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *LoginLivingPageMonthlyResponseBuilder) State(v LoginLivingPageMonthlyResponseStateEnum) *LoginLivingPageMonthlyResponseBuilder {
	b.s.State = &v
	return b
}

func (b *LoginLivingPageMonthlyResponseBuilder) Body(v *LoginLivingPageMonthlyResponseBody) *LoginLivingPageMonthlyResponseBuilder {
	b.s.Body = v
	return b
}

func (b *LoginLivingPageMonthlyResponseBuilder) Build() *LoginLivingPageMonthlyResponse {
	return b.s
}


