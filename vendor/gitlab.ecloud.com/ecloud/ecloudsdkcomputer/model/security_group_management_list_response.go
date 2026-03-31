// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type SecurityGroupManagementListResponseStateEnum string

// List of State
const (
    SecurityGroupManagementListResponseStateEnumOk SecurityGroupManagementListResponseStateEnum = "OK"
    SecurityGroupManagementListResponseStateEnumError SecurityGroupManagementListResponseStateEnum = "ERROR"
    SecurityGroupManagementListResponseStateEnumException SecurityGroupManagementListResponseStateEnum = "EXCEPTION"
    SecurityGroupManagementListResponseStateEnumForbidden SecurityGroupManagementListResponseStateEnum = "FORBIDDEN"
    SecurityGroupManagementListResponseStateEnumRemaining SecurityGroupManagementListResponseStateEnum = "REMAINING"
    SecurityGroupManagementListResponseStateEnumTimeout SecurityGroupManagementListResponseStateEnum = "TIMEOUT"
)

type SecurityGroupManagementListResponse struct {


	RequestId *string `json:"requestId,omitempty"`

	ErrorMessage *string `json:"errorMessage,omitempty"`

	ErrorCode *string `json:"errorCode,omitempty"`

	State *SecurityGroupManagementListResponseStateEnum `json:"state,omitempty"`

	Body *SecurityGroupManagementListResponseBody `json:"body,omitempty"`
}

func (s SecurityGroupManagementListResponse) String() string {
	return utils.Beautify(s)
}

func (s SecurityGroupManagementListResponse) GoString() string {
	return s.String()
}

func (s SecurityGroupManagementListResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *SecurityGroupManagementListResponse) SetRequestId(v string) *SecurityGroupManagementListResponse {
	s.RequestId = &v
	return s
}

func (s *SecurityGroupManagementListResponse) SetErrorMessage(v string) *SecurityGroupManagementListResponse {
	s.ErrorMessage = &v
	return s
}

func (s *SecurityGroupManagementListResponse) SetErrorCode(v string) *SecurityGroupManagementListResponse {
	s.ErrorCode = &v
	return s
}

func (s *SecurityGroupManagementListResponse) SetState(v SecurityGroupManagementListResponseStateEnum) *SecurityGroupManagementListResponse {
	s.State = &v
	return s
}

func (s *SecurityGroupManagementListResponse) SetBody(v *SecurityGroupManagementListResponseBody) *SecurityGroupManagementListResponse {
	s.Body = v
	return s
}


type SecurityGroupManagementListResponseBuilder struct {
	s *SecurityGroupManagementListResponse
}

func NewSecurityGroupManagementListResponseBuilder() *SecurityGroupManagementListResponseBuilder {
	s := &SecurityGroupManagementListResponse{}
	b := &SecurityGroupManagementListResponseBuilder{s: s}
	return b
}

func (b *SecurityGroupManagementListResponseBuilder) RequestId(v string) *SecurityGroupManagementListResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *SecurityGroupManagementListResponseBuilder) ErrorMessage(v string) *SecurityGroupManagementListResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *SecurityGroupManagementListResponseBuilder) ErrorCode(v string) *SecurityGroupManagementListResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *SecurityGroupManagementListResponseBuilder) State(v SecurityGroupManagementListResponseStateEnum) *SecurityGroupManagementListResponseBuilder {
	b.s.State = &v
	return b
}

func (b *SecurityGroupManagementListResponseBuilder) Body(v *SecurityGroupManagementListResponseBody) *SecurityGroupManagementListResponseBuilder {
	b.s.Body = v
	return b
}

func (b *SecurityGroupManagementListResponseBuilder) Build() *SecurityGroupManagementListResponse {
	return b.s
}


