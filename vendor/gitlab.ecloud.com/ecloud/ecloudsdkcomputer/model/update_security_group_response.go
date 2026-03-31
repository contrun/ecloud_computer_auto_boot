// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type UpdateSecurityGroupResponseStateEnum string

// List of State
const (
    UpdateSecurityGroupResponseStateEnumOk UpdateSecurityGroupResponseStateEnum = "OK"
    UpdateSecurityGroupResponseStateEnumError UpdateSecurityGroupResponseStateEnum = "ERROR"
    UpdateSecurityGroupResponseStateEnumException UpdateSecurityGroupResponseStateEnum = "EXCEPTION"
    UpdateSecurityGroupResponseStateEnumForbidden UpdateSecurityGroupResponseStateEnum = "FORBIDDEN"
    UpdateSecurityGroupResponseStateEnumRemaining UpdateSecurityGroupResponseStateEnum = "REMAINING"
    UpdateSecurityGroupResponseStateEnumTimeout UpdateSecurityGroupResponseStateEnum = "TIMEOUT"
)

type UpdateSecurityGroupResponse struct {


	RequestId *string `json:"requestId,omitempty"`

	ErrorMessage *string `json:"errorMessage,omitempty"`

	ErrorCode *string `json:"errorCode,omitempty"`

	State *UpdateSecurityGroupResponseStateEnum `json:"state,omitempty"`

	Body *string `json:"body,omitempty"`
}

func (s UpdateSecurityGroupResponse) String() string {
	return utils.Beautify(s)
}

func (s UpdateSecurityGroupResponse) GoString() string {
	return s.String()
}

func (s UpdateSecurityGroupResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateSecurityGroupResponse) SetRequestId(v string) *UpdateSecurityGroupResponse {
	s.RequestId = &v
	return s
}

func (s *UpdateSecurityGroupResponse) SetErrorMessage(v string) *UpdateSecurityGroupResponse {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateSecurityGroupResponse) SetErrorCode(v string) *UpdateSecurityGroupResponse {
	s.ErrorCode = &v
	return s
}

func (s *UpdateSecurityGroupResponse) SetState(v UpdateSecurityGroupResponseStateEnum) *UpdateSecurityGroupResponse {
	s.State = &v
	return s
}

func (s *UpdateSecurityGroupResponse) SetBody(v string) *UpdateSecurityGroupResponse {
	s.Body = &v
	return s
}


type UpdateSecurityGroupResponseBuilder struct {
	s *UpdateSecurityGroupResponse
}

func NewUpdateSecurityGroupResponseBuilder() *UpdateSecurityGroupResponseBuilder {
	s := &UpdateSecurityGroupResponse{}
	b := &UpdateSecurityGroupResponseBuilder{s: s}
	return b
}

func (b *UpdateSecurityGroupResponseBuilder) RequestId(v string) *UpdateSecurityGroupResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *UpdateSecurityGroupResponseBuilder) ErrorMessage(v string) *UpdateSecurityGroupResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *UpdateSecurityGroupResponseBuilder) ErrorCode(v string) *UpdateSecurityGroupResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *UpdateSecurityGroupResponseBuilder) State(v UpdateSecurityGroupResponseStateEnum) *UpdateSecurityGroupResponseBuilder {
	b.s.State = &v
	return b
}

func (b *UpdateSecurityGroupResponseBuilder) Body(v string) *UpdateSecurityGroupResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *UpdateSecurityGroupResponseBuilder) Build() *UpdateSecurityGroupResponse {
	return b.s
}


