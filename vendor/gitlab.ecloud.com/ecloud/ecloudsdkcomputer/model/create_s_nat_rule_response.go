// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type CreateSNatRuleResponseStateEnum string

// List of State
const (
    CreateSNatRuleResponseStateEnumOk CreateSNatRuleResponseStateEnum = "OK"
    CreateSNatRuleResponseStateEnumError CreateSNatRuleResponseStateEnum = "ERROR"
    CreateSNatRuleResponseStateEnumException CreateSNatRuleResponseStateEnum = "EXCEPTION"
    CreateSNatRuleResponseStateEnumForbidden CreateSNatRuleResponseStateEnum = "FORBIDDEN"
    CreateSNatRuleResponseStateEnumRemaining CreateSNatRuleResponseStateEnum = "REMAINING"
    CreateSNatRuleResponseStateEnumTimeout CreateSNatRuleResponseStateEnum = "TIMEOUT"
)

type CreateSNatRuleResponse struct {


	RequestId *string `json:"requestId,omitempty"`

	ErrorMessage *string `json:"errorMessage,omitempty"`

	ErrorCode *string `json:"errorCode,omitempty"`

	State *CreateSNatRuleResponseStateEnum `json:"state,omitempty"`

	Body *string `json:"body,omitempty"`
}

func (s CreateSNatRuleResponse) String() string {
	return utils.Beautify(s)
}

func (s CreateSNatRuleResponse) GoString() string {
	return s.String()
}

func (s CreateSNatRuleResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateSNatRuleResponse) SetRequestId(v string) *CreateSNatRuleResponse {
	s.RequestId = &v
	return s
}

func (s *CreateSNatRuleResponse) SetErrorMessage(v string) *CreateSNatRuleResponse {
	s.ErrorMessage = &v
	return s
}

func (s *CreateSNatRuleResponse) SetErrorCode(v string) *CreateSNatRuleResponse {
	s.ErrorCode = &v
	return s
}

func (s *CreateSNatRuleResponse) SetState(v CreateSNatRuleResponseStateEnum) *CreateSNatRuleResponse {
	s.State = &v
	return s
}

func (s *CreateSNatRuleResponse) SetBody(v string) *CreateSNatRuleResponse {
	s.Body = &v
	return s
}


type CreateSNatRuleResponseBuilder struct {
	s *CreateSNatRuleResponse
}

func NewCreateSNatRuleResponseBuilder() *CreateSNatRuleResponseBuilder {
	s := &CreateSNatRuleResponse{}
	b := &CreateSNatRuleResponseBuilder{s: s}
	return b
}

func (b *CreateSNatRuleResponseBuilder) RequestId(v string) *CreateSNatRuleResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *CreateSNatRuleResponseBuilder) ErrorMessage(v string) *CreateSNatRuleResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *CreateSNatRuleResponseBuilder) ErrorCode(v string) *CreateSNatRuleResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *CreateSNatRuleResponseBuilder) State(v CreateSNatRuleResponseStateEnum) *CreateSNatRuleResponseBuilder {
	b.s.State = &v
	return b
}

func (b *CreateSNatRuleResponseBuilder) Body(v string) *CreateSNatRuleResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *CreateSNatRuleResponseBuilder) Build() *CreateSNatRuleResponse {
	return b.s
}


