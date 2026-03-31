// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type DeleteVpcResponseStateEnum string

// List of State
const (
    DeleteVpcResponseStateEnumOk DeleteVpcResponseStateEnum = "OK"
    DeleteVpcResponseStateEnumError DeleteVpcResponseStateEnum = "ERROR"
    DeleteVpcResponseStateEnumException DeleteVpcResponseStateEnum = "EXCEPTION"
    DeleteVpcResponseStateEnumForbidden DeleteVpcResponseStateEnum = "FORBIDDEN"
    DeleteVpcResponseStateEnumRemaining DeleteVpcResponseStateEnum = "REMAINING"
    DeleteVpcResponseStateEnumTimeout DeleteVpcResponseStateEnum = "TIMEOUT"
)

type DeleteVpcResponse struct {


	RequestId *string `json:"requestId,omitempty"`

	ErrorMessage *string `json:"errorMessage,omitempty"`

	ErrorCode *string `json:"errorCode,omitempty"`

	State *DeleteVpcResponseStateEnum `json:"state,omitempty"`

	Body *string `json:"body,omitempty"`
}

func (s DeleteVpcResponse) String() string {
	return utils.Beautify(s)
}

func (s DeleteVpcResponse) GoString() string {
	return s.String()
}

func (s DeleteVpcResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteVpcResponse) SetRequestId(v string) *DeleteVpcResponse {
	s.RequestId = &v
	return s
}

func (s *DeleteVpcResponse) SetErrorMessage(v string) *DeleteVpcResponse {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteVpcResponse) SetErrorCode(v string) *DeleteVpcResponse {
	s.ErrorCode = &v
	return s
}

func (s *DeleteVpcResponse) SetState(v DeleteVpcResponseStateEnum) *DeleteVpcResponse {
	s.State = &v
	return s
}

func (s *DeleteVpcResponse) SetBody(v string) *DeleteVpcResponse {
	s.Body = &v
	return s
}


type DeleteVpcResponseBuilder struct {
	s *DeleteVpcResponse
}

func NewDeleteVpcResponseBuilder() *DeleteVpcResponseBuilder {
	s := &DeleteVpcResponse{}
	b := &DeleteVpcResponseBuilder{s: s}
	return b
}

func (b *DeleteVpcResponseBuilder) RequestId(v string) *DeleteVpcResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *DeleteVpcResponseBuilder) ErrorMessage(v string) *DeleteVpcResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *DeleteVpcResponseBuilder) ErrorCode(v string) *DeleteVpcResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *DeleteVpcResponseBuilder) State(v DeleteVpcResponseStateEnum) *DeleteVpcResponseBuilder {
	b.s.State = &v
	return b
}

func (b *DeleteVpcResponseBuilder) Body(v string) *DeleteVpcResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *DeleteVpcResponseBuilder) Build() *DeleteVpcResponse {
	return b.s
}


