// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type UpdateVpcResponseStateEnum string

// List of State
const (
    UpdateVpcResponseStateEnumOk UpdateVpcResponseStateEnum = "OK"
    UpdateVpcResponseStateEnumError UpdateVpcResponseStateEnum = "ERROR"
    UpdateVpcResponseStateEnumException UpdateVpcResponseStateEnum = "EXCEPTION"
    UpdateVpcResponseStateEnumForbidden UpdateVpcResponseStateEnum = "FORBIDDEN"
    UpdateVpcResponseStateEnumRemaining UpdateVpcResponseStateEnum = "REMAINING"
    UpdateVpcResponseStateEnumTimeout UpdateVpcResponseStateEnum = "TIMEOUT"
)

type UpdateVpcResponse struct {


	RequestId *string `json:"requestId,omitempty"`

	ErrorMessage *string `json:"errorMessage,omitempty"`

	ErrorCode *string `json:"errorCode,omitempty"`

	State *UpdateVpcResponseStateEnum `json:"state,omitempty"`

	Body *string `json:"body,omitempty"`
}

func (s UpdateVpcResponse) String() string {
	return utils.Beautify(s)
}

func (s UpdateVpcResponse) GoString() string {
	return s.String()
}

func (s UpdateVpcResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateVpcResponse) SetRequestId(v string) *UpdateVpcResponse {
	s.RequestId = &v
	return s
}

func (s *UpdateVpcResponse) SetErrorMessage(v string) *UpdateVpcResponse {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateVpcResponse) SetErrorCode(v string) *UpdateVpcResponse {
	s.ErrorCode = &v
	return s
}

func (s *UpdateVpcResponse) SetState(v UpdateVpcResponseStateEnum) *UpdateVpcResponse {
	s.State = &v
	return s
}

func (s *UpdateVpcResponse) SetBody(v string) *UpdateVpcResponse {
	s.Body = &v
	return s
}


type UpdateVpcResponseBuilder struct {
	s *UpdateVpcResponse
}

func NewUpdateVpcResponseBuilder() *UpdateVpcResponseBuilder {
	s := &UpdateVpcResponse{}
	b := &UpdateVpcResponseBuilder{s: s}
	return b
}

func (b *UpdateVpcResponseBuilder) RequestId(v string) *UpdateVpcResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *UpdateVpcResponseBuilder) ErrorMessage(v string) *UpdateVpcResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *UpdateVpcResponseBuilder) ErrorCode(v string) *UpdateVpcResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *UpdateVpcResponseBuilder) State(v UpdateVpcResponseStateEnum) *UpdateVpcResponseBuilder {
	b.s.State = &v
	return b
}

func (b *UpdateVpcResponseBuilder) Body(v string) *UpdateVpcResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *UpdateVpcResponseBuilder) Build() *UpdateVpcResponse {
	return b.s
}


