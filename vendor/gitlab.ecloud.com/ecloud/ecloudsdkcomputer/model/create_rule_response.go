// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type CreateRuleResponseStateEnum string

// List of State
const (
    CreateRuleResponseStateEnumOk CreateRuleResponseStateEnum = "OK"
    CreateRuleResponseStateEnumError CreateRuleResponseStateEnum = "ERROR"
    CreateRuleResponseStateEnumException CreateRuleResponseStateEnum = "EXCEPTION"
    CreateRuleResponseStateEnumForbidden CreateRuleResponseStateEnum = "FORBIDDEN"
    CreateRuleResponseStateEnumRemaining CreateRuleResponseStateEnum = "REMAINING"
    CreateRuleResponseStateEnumTimeout CreateRuleResponseStateEnum = "TIMEOUT"
)

type CreateRuleResponse struct {


	RequestId *string `json:"requestId,omitempty"`

	ErrorMessage *string `json:"errorMessage,omitempty"`

	ErrorCode *string `json:"errorCode,omitempty"`

	State *CreateRuleResponseStateEnum `json:"state,omitempty"`

	Body *string `json:"body,omitempty"`
}

func (s CreateRuleResponse) String() string {
	return utils.Beautify(s)
}

func (s CreateRuleResponse) GoString() string {
	return s.String()
}

func (s CreateRuleResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateRuleResponse) SetRequestId(v string) *CreateRuleResponse {
	s.RequestId = &v
	return s
}

func (s *CreateRuleResponse) SetErrorMessage(v string) *CreateRuleResponse {
	s.ErrorMessage = &v
	return s
}

func (s *CreateRuleResponse) SetErrorCode(v string) *CreateRuleResponse {
	s.ErrorCode = &v
	return s
}

func (s *CreateRuleResponse) SetState(v CreateRuleResponseStateEnum) *CreateRuleResponse {
	s.State = &v
	return s
}

func (s *CreateRuleResponse) SetBody(v string) *CreateRuleResponse {
	s.Body = &v
	return s
}


type CreateRuleResponseBuilder struct {
	s *CreateRuleResponse
}

func NewCreateRuleResponseBuilder() *CreateRuleResponseBuilder {
	s := &CreateRuleResponse{}
	b := &CreateRuleResponseBuilder{s: s}
	return b
}

func (b *CreateRuleResponseBuilder) RequestId(v string) *CreateRuleResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *CreateRuleResponseBuilder) ErrorMessage(v string) *CreateRuleResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *CreateRuleResponseBuilder) ErrorCode(v string) *CreateRuleResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *CreateRuleResponseBuilder) State(v CreateRuleResponseStateEnum) *CreateRuleResponseBuilder {
	b.s.State = &v
	return b
}

func (b *CreateRuleResponseBuilder) Body(v string) *CreateRuleResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *CreateRuleResponseBuilder) Build() *CreateRuleResponse {
	return b.s
}


