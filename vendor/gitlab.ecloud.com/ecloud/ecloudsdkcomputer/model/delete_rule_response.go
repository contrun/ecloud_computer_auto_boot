// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type DeleteRuleResponseStateEnum string

// List of State
const (
    DeleteRuleResponseStateEnumOk DeleteRuleResponseStateEnum = "OK"
    DeleteRuleResponseStateEnumError DeleteRuleResponseStateEnum = "ERROR"
    DeleteRuleResponseStateEnumException DeleteRuleResponseStateEnum = "EXCEPTION"
    DeleteRuleResponseStateEnumForbidden DeleteRuleResponseStateEnum = "FORBIDDEN"
    DeleteRuleResponseStateEnumRemaining DeleteRuleResponseStateEnum = "REMAINING"
    DeleteRuleResponseStateEnumTimeout DeleteRuleResponseStateEnum = "TIMEOUT"
)

type DeleteRuleResponse struct {


	RequestId *string `json:"requestId,omitempty"`

	ErrorMessage *string `json:"errorMessage,omitempty"`

	ErrorCode *string `json:"errorCode,omitempty"`

	State *DeleteRuleResponseStateEnum `json:"state,omitempty"`

	Body *string `json:"body,omitempty"`
}

func (s DeleteRuleResponse) String() string {
	return utils.Beautify(s)
}

func (s DeleteRuleResponse) GoString() string {
	return s.String()
}

func (s DeleteRuleResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteRuleResponse) SetRequestId(v string) *DeleteRuleResponse {
	s.RequestId = &v
	return s
}

func (s *DeleteRuleResponse) SetErrorMessage(v string) *DeleteRuleResponse {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteRuleResponse) SetErrorCode(v string) *DeleteRuleResponse {
	s.ErrorCode = &v
	return s
}

func (s *DeleteRuleResponse) SetState(v DeleteRuleResponseStateEnum) *DeleteRuleResponse {
	s.State = &v
	return s
}

func (s *DeleteRuleResponse) SetBody(v string) *DeleteRuleResponse {
	s.Body = &v
	return s
}


type DeleteRuleResponseBuilder struct {
	s *DeleteRuleResponse
}

func NewDeleteRuleResponseBuilder() *DeleteRuleResponseBuilder {
	s := &DeleteRuleResponse{}
	b := &DeleteRuleResponseBuilder{s: s}
	return b
}

func (b *DeleteRuleResponseBuilder) RequestId(v string) *DeleteRuleResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *DeleteRuleResponseBuilder) ErrorMessage(v string) *DeleteRuleResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *DeleteRuleResponseBuilder) ErrorCode(v string) *DeleteRuleResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *DeleteRuleResponseBuilder) State(v DeleteRuleResponseStateEnum) *DeleteRuleResponseBuilder {
	b.s.State = &v
	return b
}

func (b *DeleteRuleResponseBuilder) Body(v string) *DeleteRuleResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *DeleteRuleResponseBuilder) Build() *DeleteRuleResponse {
	return b.s
}


