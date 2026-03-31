// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type CreateSecurityGroupResponseStateEnum string

// List of State
const (
    CreateSecurityGroupResponseStateEnumOk CreateSecurityGroupResponseStateEnum = "OK"
    CreateSecurityGroupResponseStateEnumError CreateSecurityGroupResponseStateEnum = "ERROR"
    CreateSecurityGroupResponseStateEnumException CreateSecurityGroupResponseStateEnum = "EXCEPTION"
    CreateSecurityGroupResponseStateEnumForbidden CreateSecurityGroupResponseStateEnum = "FORBIDDEN"
    CreateSecurityGroupResponseStateEnumRemaining CreateSecurityGroupResponseStateEnum = "REMAINING"
    CreateSecurityGroupResponseStateEnumTimeout CreateSecurityGroupResponseStateEnum = "TIMEOUT"
)

type CreateSecurityGroupResponse struct {


	RequestId *string `json:"requestId,omitempty"`

	ErrorMessage *string `json:"errorMessage,omitempty"`

	ErrorCode *string `json:"errorCode,omitempty"`

	State *CreateSecurityGroupResponseStateEnum `json:"state,omitempty"`

	Body *string `json:"body,omitempty"`
}

func (s CreateSecurityGroupResponse) String() string {
	return utils.Beautify(s)
}

func (s CreateSecurityGroupResponse) GoString() string {
	return s.String()
}

func (s CreateSecurityGroupResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateSecurityGroupResponse) SetRequestId(v string) *CreateSecurityGroupResponse {
	s.RequestId = &v
	return s
}

func (s *CreateSecurityGroupResponse) SetErrorMessage(v string) *CreateSecurityGroupResponse {
	s.ErrorMessage = &v
	return s
}

func (s *CreateSecurityGroupResponse) SetErrorCode(v string) *CreateSecurityGroupResponse {
	s.ErrorCode = &v
	return s
}

func (s *CreateSecurityGroupResponse) SetState(v CreateSecurityGroupResponseStateEnum) *CreateSecurityGroupResponse {
	s.State = &v
	return s
}

func (s *CreateSecurityGroupResponse) SetBody(v string) *CreateSecurityGroupResponse {
	s.Body = &v
	return s
}


type CreateSecurityGroupResponseBuilder struct {
	s *CreateSecurityGroupResponse
}

func NewCreateSecurityGroupResponseBuilder() *CreateSecurityGroupResponseBuilder {
	s := &CreateSecurityGroupResponse{}
	b := &CreateSecurityGroupResponseBuilder{s: s}
	return b
}

func (b *CreateSecurityGroupResponseBuilder) RequestId(v string) *CreateSecurityGroupResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *CreateSecurityGroupResponseBuilder) ErrorMessage(v string) *CreateSecurityGroupResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *CreateSecurityGroupResponseBuilder) ErrorCode(v string) *CreateSecurityGroupResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *CreateSecurityGroupResponseBuilder) State(v CreateSecurityGroupResponseStateEnum) *CreateSecurityGroupResponseBuilder {
	b.s.State = &v
	return b
}

func (b *CreateSecurityGroupResponseBuilder) Body(v string) *CreateSecurityGroupResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *CreateSecurityGroupResponseBuilder) Build() *CreateSecurityGroupResponse {
	return b.s
}


