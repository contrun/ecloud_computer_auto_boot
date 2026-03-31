// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type ChangeVpcResponseStateEnum string

// List of State
const (
    ChangeVpcResponseStateEnumOk ChangeVpcResponseStateEnum = "OK"
    ChangeVpcResponseStateEnumError ChangeVpcResponseStateEnum = "ERROR"
    ChangeVpcResponseStateEnumException ChangeVpcResponseStateEnum = "EXCEPTION"
    ChangeVpcResponseStateEnumForbidden ChangeVpcResponseStateEnum = "FORBIDDEN"
    ChangeVpcResponseStateEnumRemaining ChangeVpcResponseStateEnum = "REMAINING"
    ChangeVpcResponseStateEnumTimeout ChangeVpcResponseStateEnum = "TIMEOUT"
)

type ChangeVpcResponse struct {


	RequestId *string `json:"requestId,omitempty"`

	ErrorMessage *string `json:"errorMessage,omitempty"`

	ErrorCode *string `json:"errorCode,omitempty"`

	State *ChangeVpcResponseStateEnum `json:"state,omitempty"`

	Body *string `json:"body,omitempty"`
}

func (s ChangeVpcResponse) String() string {
	return utils.Beautify(s)
}

func (s ChangeVpcResponse) GoString() string {
	return s.String()
}

func (s ChangeVpcResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ChangeVpcResponse) SetRequestId(v string) *ChangeVpcResponse {
	s.RequestId = &v
	return s
}

func (s *ChangeVpcResponse) SetErrorMessage(v string) *ChangeVpcResponse {
	s.ErrorMessage = &v
	return s
}

func (s *ChangeVpcResponse) SetErrorCode(v string) *ChangeVpcResponse {
	s.ErrorCode = &v
	return s
}

func (s *ChangeVpcResponse) SetState(v ChangeVpcResponseStateEnum) *ChangeVpcResponse {
	s.State = &v
	return s
}

func (s *ChangeVpcResponse) SetBody(v string) *ChangeVpcResponse {
	s.Body = &v
	return s
}


type ChangeVpcResponseBuilder struct {
	s *ChangeVpcResponse
}

func NewChangeVpcResponseBuilder() *ChangeVpcResponseBuilder {
	s := &ChangeVpcResponse{}
	b := &ChangeVpcResponseBuilder{s: s}
	return b
}

func (b *ChangeVpcResponseBuilder) RequestId(v string) *ChangeVpcResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *ChangeVpcResponseBuilder) ErrorMessage(v string) *ChangeVpcResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *ChangeVpcResponseBuilder) ErrorCode(v string) *ChangeVpcResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *ChangeVpcResponseBuilder) State(v ChangeVpcResponseStateEnum) *ChangeVpcResponseBuilder {
	b.s.State = &v
	return b
}

func (b *ChangeVpcResponseBuilder) Body(v string) *ChangeVpcResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *ChangeVpcResponseBuilder) Build() *ChangeVpcResponse {
	return b.s
}


