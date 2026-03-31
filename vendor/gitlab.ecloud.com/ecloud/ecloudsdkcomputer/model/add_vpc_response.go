// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type AddVpcResponseStateEnum string

// List of State
const (
    AddVpcResponseStateEnumOk AddVpcResponseStateEnum = "OK"
    AddVpcResponseStateEnumError AddVpcResponseStateEnum = "ERROR"
    AddVpcResponseStateEnumException AddVpcResponseStateEnum = "EXCEPTION"
    AddVpcResponseStateEnumForbidden AddVpcResponseStateEnum = "FORBIDDEN"
    AddVpcResponseStateEnumRemaining AddVpcResponseStateEnum = "REMAINING"
    AddVpcResponseStateEnumTimeout AddVpcResponseStateEnum = "TIMEOUT"
)

type AddVpcResponse struct {


	RequestId *string `json:"requestId,omitempty"`

	ErrorMessage *string `json:"errorMessage,omitempty"`

	ErrorCode *string `json:"errorCode,omitempty"`

	State *AddVpcResponseStateEnum `json:"state,omitempty"`

	Body *string `json:"body,omitempty"`
}

func (s AddVpcResponse) String() string {
	return utils.Beautify(s)
}

func (s AddVpcResponse) GoString() string {
	return s.String()
}

func (s AddVpcResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *AddVpcResponse) SetRequestId(v string) *AddVpcResponse {
	s.RequestId = &v
	return s
}

func (s *AddVpcResponse) SetErrorMessage(v string) *AddVpcResponse {
	s.ErrorMessage = &v
	return s
}

func (s *AddVpcResponse) SetErrorCode(v string) *AddVpcResponse {
	s.ErrorCode = &v
	return s
}

func (s *AddVpcResponse) SetState(v AddVpcResponseStateEnum) *AddVpcResponse {
	s.State = &v
	return s
}

func (s *AddVpcResponse) SetBody(v string) *AddVpcResponse {
	s.Body = &v
	return s
}


type AddVpcResponseBuilder struct {
	s *AddVpcResponse
}

func NewAddVpcResponseBuilder() *AddVpcResponseBuilder {
	s := &AddVpcResponse{}
	b := &AddVpcResponseBuilder{s: s}
	return b
}

func (b *AddVpcResponseBuilder) RequestId(v string) *AddVpcResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *AddVpcResponseBuilder) ErrorMessage(v string) *AddVpcResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *AddVpcResponseBuilder) ErrorCode(v string) *AddVpcResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *AddVpcResponseBuilder) State(v AddVpcResponseStateEnum) *AddVpcResponseBuilder {
	b.s.State = &v
	return b
}

func (b *AddVpcResponseBuilder) Body(v string) *AddVpcResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *AddVpcResponseBuilder) Build() *AddVpcResponse {
	return b.s
}


