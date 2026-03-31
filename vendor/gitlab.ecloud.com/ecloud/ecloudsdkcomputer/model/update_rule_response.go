// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type UpdateRuleResponseStateEnum string

// List of State
const (
    UpdateRuleResponseStateEnumOk UpdateRuleResponseStateEnum = "OK"
    UpdateRuleResponseStateEnumError UpdateRuleResponseStateEnum = "ERROR"
    UpdateRuleResponseStateEnumException UpdateRuleResponseStateEnum = "EXCEPTION"
    UpdateRuleResponseStateEnumForbidden UpdateRuleResponseStateEnum = "FORBIDDEN"
    UpdateRuleResponseStateEnumRemaining UpdateRuleResponseStateEnum = "REMAINING"
    UpdateRuleResponseStateEnumTimeout UpdateRuleResponseStateEnum = "TIMEOUT"
)

type UpdateRuleResponse struct {


	RequestId *string `json:"requestId,omitempty"`

	ErrorMessage *string `json:"errorMessage,omitempty"`

	ErrorCode *string `json:"errorCode,omitempty"`

	State *UpdateRuleResponseStateEnum `json:"state,omitempty"`

	Body *string `json:"body,omitempty"`
}

func (s UpdateRuleResponse) String() string {
	return utils.Beautify(s)
}

func (s UpdateRuleResponse) GoString() string {
	return s.String()
}

func (s UpdateRuleResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateRuleResponse) SetRequestId(v string) *UpdateRuleResponse {
	s.RequestId = &v
	return s
}

func (s *UpdateRuleResponse) SetErrorMessage(v string) *UpdateRuleResponse {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateRuleResponse) SetErrorCode(v string) *UpdateRuleResponse {
	s.ErrorCode = &v
	return s
}

func (s *UpdateRuleResponse) SetState(v UpdateRuleResponseStateEnum) *UpdateRuleResponse {
	s.State = &v
	return s
}

func (s *UpdateRuleResponse) SetBody(v string) *UpdateRuleResponse {
	s.Body = &v
	return s
}


type UpdateRuleResponseBuilder struct {
	s *UpdateRuleResponse
}

func NewUpdateRuleResponseBuilder() *UpdateRuleResponseBuilder {
	s := &UpdateRuleResponse{}
	b := &UpdateRuleResponseBuilder{s: s}
	return b
}

func (b *UpdateRuleResponseBuilder) RequestId(v string) *UpdateRuleResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *UpdateRuleResponseBuilder) ErrorMessage(v string) *UpdateRuleResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *UpdateRuleResponseBuilder) ErrorCode(v string) *UpdateRuleResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *UpdateRuleResponseBuilder) State(v UpdateRuleResponseStateEnum) *UpdateRuleResponseBuilder {
	b.s.State = &v
	return b
}

func (b *UpdateRuleResponseBuilder) Body(v string) *UpdateRuleResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *UpdateRuleResponseBuilder) Build() *UpdateRuleResponse {
	return b.s
}


