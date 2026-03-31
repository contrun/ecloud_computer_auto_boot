// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type DeleteNetworkResponseStateEnum string

// List of State
const (
    DeleteNetworkResponseStateEnumOk DeleteNetworkResponseStateEnum = "OK"
    DeleteNetworkResponseStateEnumError DeleteNetworkResponseStateEnum = "ERROR"
    DeleteNetworkResponseStateEnumException DeleteNetworkResponseStateEnum = "EXCEPTION"
    DeleteNetworkResponseStateEnumForbidden DeleteNetworkResponseStateEnum = "FORBIDDEN"
    DeleteNetworkResponseStateEnumRemaining DeleteNetworkResponseStateEnum = "REMAINING"
    DeleteNetworkResponseStateEnumTimeout DeleteNetworkResponseStateEnum = "TIMEOUT"
)

type DeleteNetworkResponse struct {


	RequestId *string `json:"requestId,omitempty"`

	ErrorMessage *string `json:"errorMessage,omitempty"`

	ErrorCode *string `json:"errorCode,omitempty"`

	State *DeleteNetworkResponseStateEnum `json:"state,omitempty"`

	Body *string `json:"body,omitempty"`
}

func (s DeleteNetworkResponse) String() string {
	return utils.Beautify(s)
}

func (s DeleteNetworkResponse) GoString() string {
	return s.String()
}

func (s DeleteNetworkResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteNetworkResponse) SetRequestId(v string) *DeleteNetworkResponse {
	s.RequestId = &v
	return s
}

func (s *DeleteNetworkResponse) SetErrorMessage(v string) *DeleteNetworkResponse {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteNetworkResponse) SetErrorCode(v string) *DeleteNetworkResponse {
	s.ErrorCode = &v
	return s
}

func (s *DeleteNetworkResponse) SetState(v DeleteNetworkResponseStateEnum) *DeleteNetworkResponse {
	s.State = &v
	return s
}

func (s *DeleteNetworkResponse) SetBody(v string) *DeleteNetworkResponse {
	s.Body = &v
	return s
}


type DeleteNetworkResponseBuilder struct {
	s *DeleteNetworkResponse
}

func NewDeleteNetworkResponseBuilder() *DeleteNetworkResponseBuilder {
	s := &DeleteNetworkResponse{}
	b := &DeleteNetworkResponseBuilder{s: s}
	return b
}

func (b *DeleteNetworkResponseBuilder) RequestId(v string) *DeleteNetworkResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *DeleteNetworkResponseBuilder) ErrorMessage(v string) *DeleteNetworkResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *DeleteNetworkResponseBuilder) ErrorCode(v string) *DeleteNetworkResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *DeleteNetworkResponseBuilder) State(v DeleteNetworkResponseStateEnum) *DeleteNetworkResponseBuilder {
	b.s.State = &v
	return b
}

func (b *DeleteNetworkResponseBuilder) Body(v string) *DeleteNetworkResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *DeleteNetworkResponseBuilder) Build() *DeleteNetworkResponse {
	return b.s
}


