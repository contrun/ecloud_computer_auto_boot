// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type HistoryConnectPageResponseStateEnum string

// List of State
const (
    HistoryConnectPageResponseStateEnumOk HistoryConnectPageResponseStateEnum = "OK"
    HistoryConnectPageResponseStateEnumError HistoryConnectPageResponseStateEnum = "ERROR"
    HistoryConnectPageResponseStateEnumException HistoryConnectPageResponseStateEnum = "EXCEPTION"
    HistoryConnectPageResponseStateEnumForbidden HistoryConnectPageResponseStateEnum = "FORBIDDEN"
    HistoryConnectPageResponseStateEnumRemaining HistoryConnectPageResponseStateEnum = "REMAINING"
    HistoryConnectPageResponseStateEnumTimeout HistoryConnectPageResponseStateEnum = "TIMEOUT"
)

type HistoryConnectPageResponse struct {


	RequestId *string `json:"requestId,omitempty"`

	ErrorMessage *string `json:"errorMessage,omitempty"`

	ErrorCode *string `json:"errorCode,omitempty"`

	State *HistoryConnectPageResponseStateEnum `json:"state,omitempty"`

	Body *HistoryConnectPageResponseBody `json:"body,omitempty"`
}

func (s HistoryConnectPageResponse) String() string {
	return utils.Beautify(s)
}

func (s HistoryConnectPageResponse) GoString() string {
	return s.String()
}

func (s HistoryConnectPageResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *HistoryConnectPageResponse) SetRequestId(v string) *HistoryConnectPageResponse {
	s.RequestId = &v
	return s
}

func (s *HistoryConnectPageResponse) SetErrorMessage(v string) *HistoryConnectPageResponse {
	s.ErrorMessage = &v
	return s
}

func (s *HistoryConnectPageResponse) SetErrorCode(v string) *HistoryConnectPageResponse {
	s.ErrorCode = &v
	return s
}

func (s *HistoryConnectPageResponse) SetState(v HistoryConnectPageResponseStateEnum) *HistoryConnectPageResponse {
	s.State = &v
	return s
}

func (s *HistoryConnectPageResponse) SetBody(v *HistoryConnectPageResponseBody) *HistoryConnectPageResponse {
	s.Body = v
	return s
}


type HistoryConnectPageResponseBuilder struct {
	s *HistoryConnectPageResponse
}

func NewHistoryConnectPageResponseBuilder() *HistoryConnectPageResponseBuilder {
	s := &HistoryConnectPageResponse{}
	b := &HistoryConnectPageResponseBuilder{s: s}
	return b
}

func (b *HistoryConnectPageResponseBuilder) RequestId(v string) *HistoryConnectPageResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *HistoryConnectPageResponseBuilder) ErrorMessage(v string) *HistoryConnectPageResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *HistoryConnectPageResponseBuilder) ErrorCode(v string) *HistoryConnectPageResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *HistoryConnectPageResponseBuilder) State(v HistoryConnectPageResponseStateEnum) *HistoryConnectPageResponseBuilder {
	b.s.State = &v
	return b
}

func (b *HistoryConnectPageResponseBuilder) Body(v *HistoryConnectPageResponseBody) *HistoryConnectPageResponseBuilder {
	b.s.Body = v
	return b
}

func (b *HistoryConnectPageResponseBuilder) Build() *HistoryConnectPageResponse {
	return b.s
}


