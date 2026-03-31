// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type ListSNatRuleResponseStateEnum string

// List of State
const (
    ListSNatRuleResponseStateEnumOk ListSNatRuleResponseStateEnum = "OK"
    ListSNatRuleResponseStateEnumError ListSNatRuleResponseStateEnum = "ERROR"
    ListSNatRuleResponseStateEnumException ListSNatRuleResponseStateEnum = "EXCEPTION"
    ListSNatRuleResponseStateEnumForbidden ListSNatRuleResponseStateEnum = "FORBIDDEN"
    ListSNatRuleResponseStateEnumRemaining ListSNatRuleResponseStateEnum = "REMAINING"
    ListSNatRuleResponseStateEnumTimeout ListSNatRuleResponseStateEnum = "TIMEOUT"
)

type ListSNatRuleResponse struct {


	RequestId *string `json:"requestId,omitempty"`

	ErrorMessage *string `json:"errorMessage,omitempty"`

	ErrorCode *string `json:"errorCode,omitempty"`

	State *ListSNatRuleResponseStateEnum `json:"state,omitempty"`

	Body *ListSNatRuleResponseBody `json:"body,omitempty"`
}

func (s ListSNatRuleResponse) String() string {
	return utils.Beautify(s)
}

func (s ListSNatRuleResponse) GoString() string {
	return s.String()
}

func (s ListSNatRuleResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListSNatRuleResponse) SetRequestId(v string) *ListSNatRuleResponse {
	s.RequestId = &v
	return s
}

func (s *ListSNatRuleResponse) SetErrorMessage(v string) *ListSNatRuleResponse {
	s.ErrorMessage = &v
	return s
}

func (s *ListSNatRuleResponse) SetErrorCode(v string) *ListSNatRuleResponse {
	s.ErrorCode = &v
	return s
}

func (s *ListSNatRuleResponse) SetState(v ListSNatRuleResponseStateEnum) *ListSNatRuleResponse {
	s.State = &v
	return s
}

func (s *ListSNatRuleResponse) SetBody(v *ListSNatRuleResponseBody) *ListSNatRuleResponse {
	s.Body = v
	return s
}


type ListSNatRuleResponseBuilder struct {
	s *ListSNatRuleResponse
}

func NewListSNatRuleResponseBuilder() *ListSNatRuleResponseBuilder {
	s := &ListSNatRuleResponse{}
	b := &ListSNatRuleResponseBuilder{s: s}
	return b
}

func (b *ListSNatRuleResponseBuilder) RequestId(v string) *ListSNatRuleResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *ListSNatRuleResponseBuilder) ErrorMessage(v string) *ListSNatRuleResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *ListSNatRuleResponseBuilder) ErrorCode(v string) *ListSNatRuleResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *ListSNatRuleResponseBuilder) State(v ListSNatRuleResponseStateEnum) *ListSNatRuleResponseBuilder {
	b.s.State = &v
	return b
}

func (b *ListSNatRuleResponseBuilder) Body(v *ListSNatRuleResponseBody) *ListSNatRuleResponseBuilder {
	b.s.Body = v
	return b
}

func (b *ListSNatRuleResponseBuilder) Build() *ListSNatRuleResponse {
	return b.s
}


