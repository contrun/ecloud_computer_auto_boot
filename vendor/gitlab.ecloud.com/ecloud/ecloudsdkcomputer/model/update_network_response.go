// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type UpdateNetworkResponseStateEnum string

// List of State
const (
    UpdateNetworkResponseStateEnumOk UpdateNetworkResponseStateEnum = "OK"
    UpdateNetworkResponseStateEnumError UpdateNetworkResponseStateEnum = "ERROR"
    UpdateNetworkResponseStateEnumException UpdateNetworkResponseStateEnum = "EXCEPTION"
    UpdateNetworkResponseStateEnumForbidden UpdateNetworkResponseStateEnum = "FORBIDDEN"
    UpdateNetworkResponseStateEnumRemaining UpdateNetworkResponseStateEnum = "REMAINING"
    UpdateNetworkResponseStateEnumTimeout UpdateNetworkResponseStateEnum = "TIMEOUT"
)

type UpdateNetworkResponse struct {


	RequestId *string `json:"requestId,omitempty"`

	ErrorMessage *string `json:"errorMessage,omitempty"`

	ErrorCode *string `json:"errorCode,omitempty"`

	State *UpdateNetworkResponseStateEnum `json:"state,omitempty"`

	Body *string `json:"body,omitempty"`
}

func (s UpdateNetworkResponse) String() string {
	return utils.Beautify(s)
}

func (s UpdateNetworkResponse) GoString() string {
	return s.String()
}

func (s UpdateNetworkResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateNetworkResponse) SetRequestId(v string) *UpdateNetworkResponse {
	s.RequestId = &v
	return s
}

func (s *UpdateNetworkResponse) SetErrorMessage(v string) *UpdateNetworkResponse {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateNetworkResponse) SetErrorCode(v string) *UpdateNetworkResponse {
	s.ErrorCode = &v
	return s
}

func (s *UpdateNetworkResponse) SetState(v UpdateNetworkResponseStateEnum) *UpdateNetworkResponse {
	s.State = &v
	return s
}

func (s *UpdateNetworkResponse) SetBody(v string) *UpdateNetworkResponse {
	s.Body = &v
	return s
}


type UpdateNetworkResponseBuilder struct {
	s *UpdateNetworkResponse
}

func NewUpdateNetworkResponseBuilder() *UpdateNetworkResponseBuilder {
	s := &UpdateNetworkResponse{}
	b := &UpdateNetworkResponseBuilder{s: s}
	return b
}

func (b *UpdateNetworkResponseBuilder) RequestId(v string) *UpdateNetworkResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *UpdateNetworkResponseBuilder) ErrorMessage(v string) *UpdateNetworkResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *UpdateNetworkResponseBuilder) ErrorCode(v string) *UpdateNetworkResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *UpdateNetworkResponseBuilder) State(v UpdateNetworkResponseStateEnum) *UpdateNetworkResponseBuilder {
	b.s.State = &v
	return b
}

func (b *UpdateNetworkResponseBuilder) Body(v string) *UpdateNetworkResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *UpdateNetworkResponseBuilder) Build() *UpdateNetworkResponse {
	return b.s
}


