// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type DeleteSecurityGroupResponseStateEnum string

// List of State
const (
    DeleteSecurityGroupResponseStateEnumOk DeleteSecurityGroupResponseStateEnum = "OK"
    DeleteSecurityGroupResponseStateEnumError DeleteSecurityGroupResponseStateEnum = "ERROR"
    DeleteSecurityGroupResponseStateEnumException DeleteSecurityGroupResponseStateEnum = "EXCEPTION"
    DeleteSecurityGroupResponseStateEnumForbidden DeleteSecurityGroupResponseStateEnum = "FORBIDDEN"
    DeleteSecurityGroupResponseStateEnumRemaining DeleteSecurityGroupResponseStateEnum = "REMAINING"
    DeleteSecurityGroupResponseStateEnumTimeout DeleteSecurityGroupResponseStateEnum = "TIMEOUT"
)

type DeleteSecurityGroupResponse struct {


	RequestId *string `json:"requestId,omitempty"`

	ErrorMessage *string `json:"errorMessage,omitempty"`

	ErrorCode *string `json:"errorCode,omitempty"`

	State *DeleteSecurityGroupResponseStateEnum `json:"state,omitempty"`

	Body *string `json:"body,omitempty"`
}

func (s DeleteSecurityGroupResponse) String() string {
	return utils.Beautify(s)
}

func (s DeleteSecurityGroupResponse) GoString() string {
	return s.String()
}

func (s DeleteSecurityGroupResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteSecurityGroupResponse) SetRequestId(v string) *DeleteSecurityGroupResponse {
	s.RequestId = &v
	return s
}

func (s *DeleteSecurityGroupResponse) SetErrorMessage(v string) *DeleteSecurityGroupResponse {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteSecurityGroupResponse) SetErrorCode(v string) *DeleteSecurityGroupResponse {
	s.ErrorCode = &v
	return s
}

func (s *DeleteSecurityGroupResponse) SetState(v DeleteSecurityGroupResponseStateEnum) *DeleteSecurityGroupResponse {
	s.State = &v
	return s
}

func (s *DeleteSecurityGroupResponse) SetBody(v string) *DeleteSecurityGroupResponse {
	s.Body = &v
	return s
}


type DeleteSecurityGroupResponseBuilder struct {
	s *DeleteSecurityGroupResponse
}

func NewDeleteSecurityGroupResponseBuilder() *DeleteSecurityGroupResponseBuilder {
	s := &DeleteSecurityGroupResponse{}
	b := &DeleteSecurityGroupResponseBuilder{s: s}
	return b
}

func (b *DeleteSecurityGroupResponseBuilder) RequestId(v string) *DeleteSecurityGroupResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *DeleteSecurityGroupResponseBuilder) ErrorMessage(v string) *DeleteSecurityGroupResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *DeleteSecurityGroupResponseBuilder) ErrorCode(v string) *DeleteSecurityGroupResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *DeleteSecurityGroupResponseBuilder) State(v DeleteSecurityGroupResponseStateEnum) *DeleteSecurityGroupResponseBuilder {
	b.s.State = &v
	return b
}

func (b *DeleteSecurityGroupResponseBuilder) Body(v string) *DeleteSecurityGroupResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *DeleteSecurityGroupResponseBuilder) Build() *DeleteSecurityGroupResponse {
	return b.s
}


