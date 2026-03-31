// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type GetNetworkSelectResponseStateEnum string

// List of State
const (
    GetNetworkSelectResponseStateEnumOk GetNetworkSelectResponseStateEnum = "OK"
    GetNetworkSelectResponseStateEnumError GetNetworkSelectResponseStateEnum = "ERROR"
    GetNetworkSelectResponseStateEnumException GetNetworkSelectResponseStateEnum = "EXCEPTION"
    GetNetworkSelectResponseStateEnumForbidden GetNetworkSelectResponseStateEnum = "FORBIDDEN"
    GetNetworkSelectResponseStateEnumRemaining GetNetworkSelectResponseStateEnum = "REMAINING"
    GetNetworkSelectResponseStateEnumTimeout GetNetworkSelectResponseStateEnum = "TIMEOUT"
)

type GetNetworkSelectResponse struct {


	RequestId *string `json:"requestId,omitempty"`

	ErrorMessage *string `json:"errorMessage,omitempty"`

	ErrorCode *string `json:"errorCode,omitempty"`

	State *GetNetworkSelectResponseStateEnum `json:"state,omitempty"`

	Body *GetNetworkSelectResponseBody `json:"body,omitempty"`
}

func (s GetNetworkSelectResponse) String() string {
	return utils.Beautify(s)
}

func (s GetNetworkSelectResponse) GoString() string {
	return s.String()
}

func (s GetNetworkSelectResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetNetworkSelectResponse) SetRequestId(v string) *GetNetworkSelectResponse {
	s.RequestId = &v
	return s
}

func (s *GetNetworkSelectResponse) SetErrorMessage(v string) *GetNetworkSelectResponse {
	s.ErrorMessage = &v
	return s
}

func (s *GetNetworkSelectResponse) SetErrorCode(v string) *GetNetworkSelectResponse {
	s.ErrorCode = &v
	return s
}

func (s *GetNetworkSelectResponse) SetState(v GetNetworkSelectResponseStateEnum) *GetNetworkSelectResponse {
	s.State = &v
	return s
}

func (s *GetNetworkSelectResponse) SetBody(v *GetNetworkSelectResponseBody) *GetNetworkSelectResponse {
	s.Body = v
	return s
}


type GetNetworkSelectResponseBuilder struct {
	s *GetNetworkSelectResponse
}

func NewGetNetworkSelectResponseBuilder() *GetNetworkSelectResponseBuilder {
	s := &GetNetworkSelectResponse{}
	b := &GetNetworkSelectResponseBuilder{s: s}
	return b
}

func (b *GetNetworkSelectResponseBuilder) RequestId(v string) *GetNetworkSelectResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *GetNetworkSelectResponseBuilder) ErrorMessage(v string) *GetNetworkSelectResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *GetNetworkSelectResponseBuilder) ErrorCode(v string) *GetNetworkSelectResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *GetNetworkSelectResponseBuilder) State(v GetNetworkSelectResponseStateEnum) *GetNetworkSelectResponseBuilder {
	b.s.State = &v
	return b
}

func (b *GetNetworkSelectResponseBuilder) Body(v *GetNetworkSelectResponseBody) *GetNetworkSelectResponseBuilder {
	b.s.Body = v
	return b
}

func (b *GetNetworkSelectResponseBuilder) Build() *GetNetworkSelectResponse {
	return b.s
}


