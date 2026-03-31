// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type HistoryPageResponseStateEnum string

// List of State
const (
    HistoryPageResponseStateEnumOk HistoryPageResponseStateEnum = "OK"
    HistoryPageResponseStateEnumError HistoryPageResponseStateEnum = "ERROR"
    HistoryPageResponseStateEnumException HistoryPageResponseStateEnum = "EXCEPTION"
    HistoryPageResponseStateEnumForbidden HistoryPageResponseStateEnum = "FORBIDDEN"
    HistoryPageResponseStateEnumRemaining HistoryPageResponseStateEnum = "REMAINING"
    HistoryPageResponseStateEnumTimeout HistoryPageResponseStateEnum = "TIMEOUT"
)

type HistoryPageResponse struct {


	RequestId *string `json:"requestId,omitempty"`

	ErrorMessage *string `json:"errorMessage,omitempty"`

	ErrorCode *string `json:"errorCode,omitempty"`

	State *HistoryPageResponseStateEnum `json:"state,omitempty"`

	Body *HistoryPageResponseBody `json:"body,omitempty"`
}

func (s HistoryPageResponse) String() string {
	return utils.Beautify(s)
}

func (s HistoryPageResponse) GoString() string {
	return s.String()
}

func (s HistoryPageResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *HistoryPageResponse) SetRequestId(v string) *HistoryPageResponse {
	s.RequestId = &v
	return s
}

func (s *HistoryPageResponse) SetErrorMessage(v string) *HistoryPageResponse {
	s.ErrorMessage = &v
	return s
}

func (s *HistoryPageResponse) SetErrorCode(v string) *HistoryPageResponse {
	s.ErrorCode = &v
	return s
}

func (s *HistoryPageResponse) SetState(v HistoryPageResponseStateEnum) *HistoryPageResponse {
	s.State = &v
	return s
}

func (s *HistoryPageResponse) SetBody(v *HistoryPageResponseBody) *HistoryPageResponse {
	s.Body = v
	return s
}


type HistoryPageResponseBuilder struct {
	s *HistoryPageResponse
}

func NewHistoryPageResponseBuilder() *HistoryPageResponseBuilder {
	s := &HistoryPageResponse{}
	b := &HistoryPageResponseBuilder{s: s}
	return b
}

func (b *HistoryPageResponseBuilder) RequestId(v string) *HistoryPageResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *HistoryPageResponseBuilder) ErrorMessage(v string) *HistoryPageResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *HistoryPageResponseBuilder) ErrorCode(v string) *HistoryPageResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *HistoryPageResponseBuilder) State(v HistoryPageResponseStateEnum) *HistoryPageResponseBuilder {
	b.s.State = &v
	return b
}

func (b *HistoryPageResponseBuilder) Body(v *HistoryPageResponseBody) *HistoryPageResponseBuilder {
	b.s.Body = v
	return b
}

func (b *HistoryPageResponseBuilder) Build() *HistoryPageResponse {
	return b.s
}


