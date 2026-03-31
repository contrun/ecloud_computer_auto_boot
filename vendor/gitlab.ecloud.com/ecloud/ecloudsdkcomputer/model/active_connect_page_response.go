// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type ActiveConnectPageResponseStateEnum string

// List of State
const (
    ActiveConnectPageResponseStateEnumOk ActiveConnectPageResponseStateEnum = "OK"
    ActiveConnectPageResponseStateEnumError ActiveConnectPageResponseStateEnum = "ERROR"
    ActiveConnectPageResponseStateEnumException ActiveConnectPageResponseStateEnum = "EXCEPTION"
    ActiveConnectPageResponseStateEnumForbidden ActiveConnectPageResponseStateEnum = "FORBIDDEN"
    ActiveConnectPageResponseStateEnumRemaining ActiveConnectPageResponseStateEnum = "REMAINING"
    ActiveConnectPageResponseStateEnumTimeout ActiveConnectPageResponseStateEnum = "TIMEOUT"
)

type ActiveConnectPageResponse struct {


	RequestId *string `json:"requestId,omitempty"`

	ErrorMessage *string `json:"errorMessage,omitempty"`

	ErrorCode *string `json:"errorCode,omitempty"`

	State *ActiveConnectPageResponseStateEnum `json:"state,omitempty"`

	Body *ActiveConnectPageResponseBody `json:"body,omitempty"`
}

func (s ActiveConnectPageResponse) String() string {
	return utils.Beautify(s)
}

func (s ActiveConnectPageResponse) GoString() string {
	return s.String()
}

func (s ActiveConnectPageResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ActiveConnectPageResponse) SetRequestId(v string) *ActiveConnectPageResponse {
	s.RequestId = &v
	return s
}

func (s *ActiveConnectPageResponse) SetErrorMessage(v string) *ActiveConnectPageResponse {
	s.ErrorMessage = &v
	return s
}

func (s *ActiveConnectPageResponse) SetErrorCode(v string) *ActiveConnectPageResponse {
	s.ErrorCode = &v
	return s
}

func (s *ActiveConnectPageResponse) SetState(v ActiveConnectPageResponseStateEnum) *ActiveConnectPageResponse {
	s.State = &v
	return s
}

func (s *ActiveConnectPageResponse) SetBody(v *ActiveConnectPageResponseBody) *ActiveConnectPageResponse {
	s.Body = v
	return s
}


type ActiveConnectPageResponseBuilder struct {
	s *ActiveConnectPageResponse
}

func NewActiveConnectPageResponseBuilder() *ActiveConnectPageResponseBuilder {
	s := &ActiveConnectPageResponse{}
	b := &ActiveConnectPageResponseBuilder{s: s}
	return b
}

func (b *ActiveConnectPageResponseBuilder) RequestId(v string) *ActiveConnectPageResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *ActiveConnectPageResponseBuilder) ErrorMessage(v string) *ActiveConnectPageResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *ActiveConnectPageResponseBuilder) ErrorCode(v string) *ActiveConnectPageResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *ActiveConnectPageResponseBuilder) State(v ActiveConnectPageResponseStateEnum) *ActiveConnectPageResponseBuilder {
	b.s.State = &v
	return b
}

func (b *ActiveConnectPageResponseBuilder) Body(v *ActiveConnectPageResponseBody) *ActiveConnectPageResponseBuilder {
	b.s.Body = v
	return b
}

func (b *ActiveConnectPageResponseBuilder) Build() *ActiveConnectPageResponse {
	return b.s
}


