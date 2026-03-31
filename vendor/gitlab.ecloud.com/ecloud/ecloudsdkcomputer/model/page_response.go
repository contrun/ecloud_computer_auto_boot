// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type PageResponseStateEnum string

// List of State
const (
    PageResponseStateEnumOk PageResponseStateEnum = "OK"
    PageResponseStateEnumError PageResponseStateEnum = "ERROR"
    PageResponseStateEnumException PageResponseStateEnum = "EXCEPTION"
    PageResponseStateEnumForbidden PageResponseStateEnum = "FORBIDDEN"
    PageResponseStateEnumRemaining PageResponseStateEnum = "REMAINING"
    PageResponseStateEnumTimeout PageResponseStateEnum = "TIMEOUT"
)

type PageResponse struct {


	RequestId *string `json:"requestId,omitempty"`

	ErrorMessage *string `json:"errorMessage,omitempty"`

	ErrorCode *string `json:"errorCode,omitempty"`

	State *PageResponseStateEnum `json:"state,omitempty"`

	Body *PageResponseBody `json:"body,omitempty"`
}

func (s PageResponse) String() string {
	return utils.Beautify(s)
}

func (s PageResponse) GoString() string {
	return s.String()
}

func (s PageResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *PageResponse) SetRequestId(v string) *PageResponse {
	s.RequestId = &v
	return s
}

func (s *PageResponse) SetErrorMessage(v string) *PageResponse {
	s.ErrorMessage = &v
	return s
}

func (s *PageResponse) SetErrorCode(v string) *PageResponse {
	s.ErrorCode = &v
	return s
}

func (s *PageResponse) SetState(v PageResponseStateEnum) *PageResponse {
	s.State = &v
	return s
}

func (s *PageResponse) SetBody(v *PageResponseBody) *PageResponse {
	s.Body = v
	return s
}


type PageResponseBuilder struct {
	s *PageResponse
}

func NewPageResponseBuilder() *PageResponseBuilder {
	s := &PageResponse{}
	b := &PageResponseBuilder{s: s}
	return b
}

func (b *PageResponseBuilder) RequestId(v string) *PageResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *PageResponseBuilder) ErrorMessage(v string) *PageResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *PageResponseBuilder) ErrorCode(v string) *PageResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *PageResponseBuilder) State(v PageResponseStateEnum) *PageResponseBuilder {
	b.s.State = &v
	return b
}

func (b *PageResponseBuilder) Body(v *PageResponseBody) *PageResponseBuilder {
	b.s.Body = v
	return b
}

func (b *PageResponseBuilder) Build() *PageResponse {
	return b.s
}


