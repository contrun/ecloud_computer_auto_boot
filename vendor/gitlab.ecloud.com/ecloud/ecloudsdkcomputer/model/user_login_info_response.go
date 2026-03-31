// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type UserLoginInfoResponseStateEnum string

// List of State
const (
    UserLoginInfoResponseStateEnumOk UserLoginInfoResponseStateEnum = "OK"
    UserLoginInfoResponseStateEnumError UserLoginInfoResponseStateEnum = "ERROR"
    UserLoginInfoResponseStateEnumException UserLoginInfoResponseStateEnum = "EXCEPTION"
    UserLoginInfoResponseStateEnumForbidden UserLoginInfoResponseStateEnum = "FORBIDDEN"
    UserLoginInfoResponseStateEnumRemaining UserLoginInfoResponseStateEnum = "REMAINING"
    UserLoginInfoResponseStateEnumTimeout UserLoginInfoResponseStateEnum = "TIMEOUT"
)

type UserLoginInfoResponse struct {


	RequestId *string `json:"requestId,omitempty"`

	ErrorMessage *string `json:"errorMessage,omitempty"`

	ErrorCode *string `json:"errorCode,omitempty"`

	State *UserLoginInfoResponseStateEnum `json:"state,omitempty"`

	Body *UserLoginInfoResponseBody `json:"body,omitempty"`
}

func (s UserLoginInfoResponse) String() string {
	return utils.Beautify(s)
}

func (s UserLoginInfoResponse) GoString() string {
	return s.String()
}

func (s UserLoginInfoResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UserLoginInfoResponse) SetRequestId(v string) *UserLoginInfoResponse {
	s.RequestId = &v
	return s
}

func (s *UserLoginInfoResponse) SetErrorMessage(v string) *UserLoginInfoResponse {
	s.ErrorMessage = &v
	return s
}

func (s *UserLoginInfoResponse) SetErrorCode(v string) *UserLoginInfoResponse {
	s.ErrorCode = &v
	return s
}

func (s *UserLoginInfoResponse) SetState(v UserLoginInfoResponseStateEnum) *UserLoginInfoResponse {
	s.State = &v
	return s
}

func (s *UserLoginInfoResponse) SetBody(v *UserLoginInfoResponseBody) *UserLoginInfoResponse {
	s.Body = v
	return s
}


type UserLoginInfoResponseBuilder struct {
	s *UserLoginInfoResponse
}

func NewUserLoginInfoResponseBuilder() *UserLoginInfoResponseBuilder {
	s := &UserLoginInfoResponse{}
	b := &UserLoginInfoResponseBuilder{s: s}
	return b
}

func (b *UserLoginInfoResponseBuilder) RequestId(v string) *UserLoginInfoResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *UserLoginInfoResponseBuilder) ErrorMessage(v string) *UserLoginInfoResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *UserLoginInfoResponseBuilder) ErrorCode(v string) *UserLoginInfoResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *UserLoginInfoResponseBuilder) State(v UserLoginInfoResponseStateEnum) *UserLoginInfoResponseBuilder {
	b.s.State = &v
	return b
}

func (b *UserLoginInfoResponseBuilder) Body(v *UserLoginInfoResponseBody) *UserLoginInfoResponseBuilder {
	b.s.Body = v
	return b
}

func (b *UserLoginInfoResponseBuilder) Build() *UserLoginInfoResponse {
	return b.s
}


