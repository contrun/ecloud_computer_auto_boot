// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type DeleteSNatRuleResponseStateEnum string

// List of State
const (
    DeleteSNatRuleResponseStateEnumOk DeleteSNatRuleResponseStateEnum = "OK"
    DeleteSNatRuleResponseStateEnumError DeleteSNatRuleResponseStateEnum = "ERROR"
    DeleteSNatRuleResponseStateEnumException DeleteSNatRuleResponseStateEnum = "EXCEPTION"
    DeleteSNatRuleResponseStateEnumForbidden DeleteSNatRuleResponseStateEnum = "FORBIDDEN"
    DeleteSNatRuleResponseStateEnumRemaining DeleteSNatRuleResponseStateEnum = "REMAINING"
    DeleteSNatRuleResponseStateEnumTimeout DeleteSNatRuleResponseStateEnum = "TIMEOUT"
)

type DeleteSNatRuleResponse struct {


	RequestId *string `json:"requestId,omitempty"`

	ErrorMessage *string `json:"errorMessage,omitempty"`

	ErrorCode *string `json:"errorCode,omitempty"`

	State *DeleteSNatRuleResponseStateEnum `json:"state,omitempty"`

	Body *string `json:"body,omitempty"`
}

func (s DeleteSNatRuleResponse) String() string {
	return utils.Beautify(s)
}

func (s DeleteSNatRuleResponse) GoString() string {
	return s.String()
}

func (s DeleteSNatRuleResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteSNatRuleResponse) SetRequestId(v string) *DeleteSNatRuleResponse {
	s.RequestId = &v
	return s
}

func (s *DeleteSNatRuleResponse) SetErrorMessage(v string) *DeleteSNatRuleResponse {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteSNatRuleResponse) SetErrorCode(v string) *DeleteSNatRuleResponse {
	s.ErrorCode = &v
	return s
}

func (s *DeleteSNatRuleResponse) SetState(v DeleteSNatRuleResponseStateEnum) *DeleteSNatRuleResponse {
	s.State = &v
	return s
}

func (s *DeleteSNatRuleResponse) SetBody(v string) *DeleteSNatRuleResponse {
	s.Body = &v
	return s
}


type DeleteSNatRuleResponseBuilder struct {
	s *DeleteSNatRuleResponse
}

func NewDeleteSNatRuleResponseBuilder() *DeleteSNatRuleResponseBuilder {
	s := &DeleteSNatRuleResponse{}
	b := &DeleteSNatRuleResponseBuilder{s: s}
	return b
}

func (b *DeleteSNatRuleResponseBuilder) RequestId(v string) *DeleteSNatRuleResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *DeleteSNatRuleResponseBuilder) ErrorMessage(v string) *DeleteSNatRuleResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *DeleteSNatRuleResponseBuilder) ErrorCode(v string) *DeleteSNatRuleResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *DeleteSNatRuleResponseBuilder) State(v DeleteSNatRuleResponseStateEnum) *DeleteSNatRuleResponseBuilder {
	b.s.State = &v
	return b
}

func (b *DeleteSNatRuleResponseBuilder) Body(v string) *DeleteSNatRuleResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *DeleteSNatRuleResponseBuilder) Build() *DeleteSNatRuleResponse {
	return b.s
}


