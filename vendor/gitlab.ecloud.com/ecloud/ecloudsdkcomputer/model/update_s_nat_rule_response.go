// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type UpdateSNatRuleResponseStateEnum string

// List of State
const (
    UpdateSNatRuleResponseStateEnumOk UpdateSNatRuleResponseStateEnum = "OK"
    UpdateSNatRuleResponseStateEnumError UpdateSNatRuleResponseStateEnum = "ERROR"
    UpdateSNatRuleResponseStateEnumException UpdateSNatRuleResponseStateEnum = "EXCEPTION"
    UpdateSNatRuleResponseStateEnumForbidden UpdateSNatRuleResponseStateEnum = "FORBIDDEN"
    UpdateSNatRuleResponseStateEnumRemaining UpdateSNatRuleResponseStateEnum = "REMAINING"
    UpdateSNatRuleResponseStateEnumTimeout UpdateSNatRuleResponseStateEnum = "TIMEOUT"
)

type UpdateSNatRuleResponse struct {


	RequestId *string `json:"requestId,omitempty"`

	ErrorMessage *string `json:"errorMessage,omitempty"`

	ErrorCode *string `json:"errorCode,omitempty"`

	State *UpdateSNatRuleResponseStateEnum `json:"state,omitempty"`

	Body *string `json:"body,omitempty"`
}

func (s UpdateSNatRuleResponse) String() string {
	return utils.Beautify(s)
}

func (s UpdateSNatRuleResponse) GoString() string {
	return s.String()
}

func (s UpdateSNatRuleResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateSNatRuleResponse) SetRequestId(v string) *UpdateSNatRuleResponse {
	s.RequestId = &v
	return s
}

func (s *UpdateSNatRuleResponse) SetErrorMessage(v string) *UpdateSNatRuleResponse {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateSNatRuleResponse) SetErrorCode(v string) *UpdateSNatRuleResponse {
	s.ErrorCode = &v
	return s
}

func (s *UpdateSNatRuleResponse) SetState(v UpdateSNatRuleResponseStateEnum) *UpdateSNatRuleResponse {
	s.State = &v
	return s
}

func (s *UpdateSNatRuleResponse) SetBody(v string) *UpdateSNatRuleResponse {
	s.Body = &v
	return s
}


type UpdateSNatRuleResponseBuilder struct {
	s *UpdateSNatRuleResponse
}

func NewUpdateSNatRuleResponseBuilder() *UpdateSNatRuleResponseBuilder {
	s := &UpdateSNatRuleResponse{}
	b := &UpdateSNatRuleResponseBuilder{s: s}
	return b
}

func (b *UpdateSNatRuleResponseBuilder) RequestId(v string) *UpdateSNatRuleResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *UpdateSNatRuleResponseBuilder) ErrorMessage(v string) *UpdateSNatRuleResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *UpdateSNatRuleResponseBuilder) ErrorCode(v string) *UpdateSNatRuleResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *UpdateSNatRuleResponseBuilder) State(v UpdateSNatRuleResponseStateEnum) *UpdateSNatRuleResponseBuilder {
	b.s.State = &v
	return b
}

func (b *UpdateSNatRuleResponseBuilder) Body(v string) *UpdateSNatRuleResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *UpdateSNatRuleResponseBuilder) Build() *UpdateSNatRuleResponse {
	return b.s
}


