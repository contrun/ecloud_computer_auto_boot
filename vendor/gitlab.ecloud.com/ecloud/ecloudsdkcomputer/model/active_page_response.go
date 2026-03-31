// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type ActivePageResponseStateEnum string

// List of State
const (
    ActivePageResponseStateEnumOk ActivePageResponseStateEnum = "OK"
    ActivePageResponseStateEnumError ActivePageResponseStateEnum = "ERROR"
    ActivePageResponseStateEnumException ActivePageResponseStateEnum = "EXCEPTION"
    ActivePageResponseStateEnumForbidden ActivePageResponseStateEnum = "FORBIDDEN"
    ActivePageResponseStateEnumRemaining ActivePageResponseStateEnum = "REMAINING"
    ActivePageResponseStateEnumTimeout ActivePageResponseStateEnum = "TIMEOUT"
)

type ActivePageResponse struct {


	RequestId *string `json:"requestId,omitempty"`

	ErrorMessage *string `json:"errorMessage,omitempty"`

	ErrorCode *string `json:"errorCode,omitempty"`

	State *ActivePageResponseStateEnum `json:"state,omitempty"`

	Body *ActivePageResponseBody `json:"body,omitempty"`
}

func (s ActivePageResponse) String() string {
	return utils.Beautify(s)
}

func (s ActivePageResponse) GoString() string {
	return s.String()
}

func (s ActivePageResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ActivePageResponse) SetRequestId(v string) *ActivePageResponse {
	s.RequestId = &v
	return s
}

func (s *ActivePageResponse) SetErrorMessage(v string) *ActivePageResponse {
	s.ErrorMessage = &v
	return s
}

func (s *ActivePageResponse) SetErrorCode(v string) *ActivePageResponse {
	s.ErrorCode = &v
	return s
}

func (s *ActivePageResponse) SetState(v ActivePageResponseStateEnum) *ActivePageResponse {
	s.State = &v
	return s
}

func (s *ActivePageResponse) SetBody(v *ActivePageResponseBody) *ActivePageResponse {
	s.Body = v
	return s
}


type ActivePageResponseBuilder struct {
	s *ActivePageResponse
}

func NewActivePageResponseBuilder() *ActivePageResponseBuilder {
	s := &ActivePageResponse{}
	b := &ActivePageResponseBuilder{s: s}
	return b
}

func (b *ActivePageResponseBuilder) RequestId(v string) *ActivePageResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *ActivePageResponseBuilder) ErrorMessage(v string) *ActivePageResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *ActivePageResponseBuilder) ErrorCode(v string) *ActivePageResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *ActivePageResponseBuilder) State(v ActivePageResponseStateEnum) *ActivePageResponseBuilder {
	b.s.State = &v
	return b
}

func (b *ActivePageResponseBuilder) Body(v *ActivePageResponseBody) *ActivePageResponseBuilder {
	b.s.Body = v
	return b
}

func (b *ActivePageResponseBuilder) Build() *ActivePageResponse {
	return b.s
}


