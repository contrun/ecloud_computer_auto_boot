// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type CopySecurityGroupResponseStateEnum string

// List of State
const (
    CopySecurityGroupResponseStateEnumOk CopySecurityGroupResponseStateEnum = "OK"
    CopySecurityGroupResponseStateEnumError CopySecurityGroupResponseStateEnum = "ERROR"
    CopySecurityGroupResponseStateEnumException CopySecurityGroupResponseStateEnum = "EXCEPTION"
    CopySecurityGroupResponseStateEnumForbidden CopySecurityGroupResponseStateEnum = "FORBIDDEN"
    CopySecurityGroupResponseStateEnumRemaining CopySecurityGroupResponseStateEnum = "REMAINING"
    CopySecurityGroupResponseStateEnumTimeout CopySecurityGroupResponseStateEnum = "TIMEOUT"
)

type CopySecurityGroupResponse struct {


	RequestId *string `json:"requestId,omitempty"`

	ErrorMessage *string `json:"errorMessage,omitempty"`

	ErrorCode *string `json:"errorCode,omitempty"`

	State *CopySecurityGroupResponseStateEnum `json:"state,omitempty"`

	Body *string `json:"body,omitempty"`
}

func (s CopySecurityGroupResponse) String() string {
	return utils.Beautify(s)
}

func (s CopySecurityGroupResponse) GoString() string {
	return s.String()
}

func (s CopySecurityGroupResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CopySecurityGroupResponse) SetRequestId(v string) *CopySecurityGroupResponse {
	s.RequestId = &v
	return s
}

func (s *CopySecurityGroupResponse) SetErrorMessage(v string) *CopySecurityGroupResponse {
	s.ErrorMessage = &v
	return s
}

func (s *CopySecurityGroupResponse) SetErrorCode(v string) *CopySecurityGroupResponse {
	s.ErrorCode = &v
	return s
}

func (s *CopySecurityGroupResponse) SetState(v CopySecurityGroupResponseStateEnum) *CopySecurityGroupResponse {
	s.State = &v
	return s
}

func (s *CopySecurityGroupResponse) SetBody(v string) *CopySecurityGroupResponse {
	s.Body = &v
	return s
}


type CopySecurityGroupResponseBuilder struct {
	s *CopySecurityGroupResponse
}

func NewCopySecurityGroupResponseBuilder() *CopySecurityGroupResponseBuilder {
	s := &CopySecurityGroupResponse{}
	b := &CopySecurityGroupResponseBuilder{s: s}
	return b
}

func (b *CopySecurityGroupResponseBuilder) RequestId(v string) *CopySecurityGroupResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *CopySecurityGroupResponseBuilder) ErrorMessage(v string) *CopySecurityGroupResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *CopySecurityGroupResponseBuilder) ErrorCode(v string) *CopySecurityGroupResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *CopySecurityGroupResponseBuilder) State(v CopySecurityGroupResponseStateEnum) *CopySecurityGroupResponseBuilder {
	b.s.State = &v
	return b
}

func (b *CopySecurityGroupResponseBuilder) Body(v string) *CopySecurityGroupResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *CopySecurityGroupResponseBuilder) Build() *CopySecurityGroupResponse {
	return b.s
}


