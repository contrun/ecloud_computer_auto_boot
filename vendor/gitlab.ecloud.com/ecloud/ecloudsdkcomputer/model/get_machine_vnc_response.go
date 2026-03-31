// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type GetMachineVNCResponseStateEnum string

// List of State
const (
    GetMachineVNCResponseStateEnumOk GetMachineVNCResponseStateEnum = "OK"
    GetMachineVNCResponseStateEnumError GetMachineVNCResponseStateEnum = "ERROR"
    GetMachineVNCResponseStateEnumException GetMachineVNCResponseStateEnum = "EXCEPTION"
    GetMachineVNCResponseStateEnumForbidden GetMachineVNCResponseStateEnum = "FORBIDDEN"
    GetMachineVNCResponseStateEnumRemaining GetMachineVNCResponseStateEnum = "REMAINING"
    GetMachineVNCResponseStateEnumTimeout GetMachineVNCResponseStateEnum = "TIMEOUT"
)

type GetMachineVNCResponse struct {


	RequestId *string `json:"requestId,omitempty"`

	ErrorMessage *string `json:"errorMessage,omitempty"`

	ErrorCode *string `json:"errorCode,omitempty"`

	State *GetMachineVNCResponseStateEnum `json:"state,omitempty"`

	Body *string `json:"body,omitempty"`
}

func (s GetMachineVNCResponse) String() string {
	return utils.Beautify(s)
}

func (s GetMachineVNCResponse) GoString() string {
	return s.String()
}

func (s GetMachineVNCResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetMachineVNCResponse) SetRequestId(v string) *GetMachineVNCResponse {
	s.RequestId = &v
	return s
}

func (s *GetMachineVNCResponse) SetErrorMessage(v string) *GetMachineVNCResponse {
	s.ErrorMessage = &v
	return s
}

func (s *GetMachineVNCResponse) SetErrorCode(v string) *GetMachineVNCResponse {
	s.ErrorCode = &v
	return s
}

func (s *GetMachineVNCResponse) SetState(v GetMachineVNCResponseStateEnum) *GetMachineVNCResponse {
	s.State = &v
	return s
}

func (s *GetMachineVNCResponse) SetBody(v string) *GetMachineVNCResponse {
	s.Body = &v
	return s
}


type GetMachineVNCResponseBuilder struct {
	s *GetMachineVNCResponse
}

func NewGetMachineVNCResponseBuilder() *GetMachineVNCResponseBuilder {
	s := &GetMachineVNCResponse{}
	b := &GetMachineVNCResponseBuilder{s: s}
	return b
}

func (b *GetMachineVNCResponseBuilder) RequestId(v string) *GetMachineVNCResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *GetMachineVNCResponseBuilder) ErrorMessage(v string) *GetMachineVNCResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *GetMachineVNCResponseBuilder) ErrorCode(v string) *GetMachineVNCResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *GetMachineVNCResponseBuilder) State(v GetMachineVNCResponseStateEnum) *GetMachineVNCResponseBuilder {
	b.s.State = &v
	return b
}

func (b *GetMachineVNCResponseBuilder) Body(v string) *GetMachineVNCResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *GetMachineVNCResponseBuilder) Build() *GetMachineVNCResponse {
	return b.s
}


