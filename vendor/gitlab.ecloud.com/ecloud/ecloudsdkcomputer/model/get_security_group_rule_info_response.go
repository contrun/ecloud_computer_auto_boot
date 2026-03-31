// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type GetSecurityGroupRuleInfoResponseStateEnum string

// List of State
const (
    GetSecurityGroupRuleInfoResponseStateEnumOk GetSecurityGroupRuleInfoResponseStateEnum = "OK"
    GetSecurityGroupRuleInfoResponseStateEnumError GetSecurityGroupRuleInfoResponseStateEnum = "ERROR"
    GetSecurityGroupRuleInfoResponseStateEnumException GetSecurityGroupRuleInfoResponseStateEnum = "EXCEPTION"
    GetSecurityGroupRuleInfoResponseStateEnumForbidden GetSecurityGroupRuleInfoResponseStateEnum = "FORBIDDEN"
    GetSecurityGroupRuleInfoResponseStateEnumRemaining GetSecurityGroupRuleInfoResponseStateEnum = "REMAINING"
    GetSecurityGroupRuleInfoResponseStateEnumTimeout GetSecurityGroupRuleInfoResponseStateEnum = "TIMEOUT"
)

type GetSecurityGroupRuleInfoResponse struct {


	RequestId *string `json:"requestId,omitempty"`

	ErrorMessage *string `json:"errorMessage,omitempty"`

	ErrorCode *string `json:"errorCode,omitempty"`

	State *GetSecurityGroupRuleInfoResponseStateEnum `json:"state,omitempty"`

	Body *GetSecurityGroupRuleInfoResponseBody `json:"body,omitempty"`
}

func (s GetSecurityGroupRuleInfoResponse) String() string {
	return utils.Beautify(s)
}

func (s GetSecurityGroupRuleInfoResponse) GoString() string {
	return s.String()
}

func (s GetSecurityGroupRuleInfoResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetSecurityGroupRuleInfoResponse) SetRequestId(v string) *GetSecurityGroupRuleInfoResponse {
	s.RequestId = &v
	return s
}

func (s *GetSecurityGroupRuleInfoResponse) SetErrorMessage(v string) *GetSecurityGroupRuleInfoResponse {
	s.ErrorMessage = &v
	return s
}

func (s *GetSecurityGroupRuleInfoResponse) SetErrorCode(v string) *GetSecurityGroupRuleInfoResponse {
	s.ErrorCode = &v
	return s
}

func (s *GetSecurityGroupRuleInfoResponse) SetState(v GetSecurityGroupRuleInfoResponseStateEnum) *GetSecurityGroupRuleInfoResponse {
	s.State = &v
	return s
}

func (s *GetSecurityGroupRuleInfoResponse) SetBody(v *GetSecurityGroupRuleInfoResponseBody) *GetSecurityGroupRuleInfoResponse {
	s.Body = v
	return s
}


type GetSecurityGroupRuleInfoResponseBuilder struct {
	s *GetSecurityGroupRuleInfoResponse
}

func NewGetSecurityGroupRuleInfoResponseBuilder() *GetSecurityGroupRuleInfoResponseBuilder {
	s := &GetSecurityGroupRuleInfoResponse{}
	b := &GetSecurityGroupRuleInfoResponseBuilder{s: s}
	return b
}

func (b *GetSecurityGroupRuleInfoResponseBuilder) RequestId(v string) *GetSecurityGroupRuleInfoResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *GetSecurityGroupRuleInfoResponseBuilder) ErrorMessage(v string) *GetSecurityGroupRuleInfoResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *GetSecurityGroupRuleInfoResponseBuilder) ErrorCode(v string) *GetSecurityGroupRuleInfoResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *GetSecurityGroupRuleInfoResponseBuilder) State(v GetSecurityGroupRuleInfoResponseStateEnum) *GetSecurityGroupRuleInfoResponseBuilder {
	b.s.State = &v
	return b
}

func (b *GetSecurityGroupRuleInfoResponseBuilder) Body(v *GetSecurityGroupRuleInfoResponseBody) *GetSecurityGroupRuleInfoResponseBuilder {
	b.s.Body = v
	return b
}

func (b *GetSecurityGroupRuleInfoResponseBuilder) Build() *GetSecurityGroupRuleInfoResponse {
	return b.s
}


