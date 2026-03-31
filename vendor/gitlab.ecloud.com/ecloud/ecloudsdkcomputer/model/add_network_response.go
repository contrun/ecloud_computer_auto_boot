// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type AddNetworkResponseStateEnum string

// List of State
const (
    AddNetworkResponseStateEnumOk AddNetworkResponseStateEnum = "OK"
    AddNetworkResponseStateEnumError AddNetworkResponseStateEnum = "ERROR"
    AddNetworkResponseStateEnumException AddNetworkResponseStateEnum = "EXCEPTION"
    AddNetworkResponseStateEnumForbidden AddNetworkResponseStateEnum = "FORBIDDEN"
    AddNetworkResponseStateEnumRemaining AddNetworkResponseStateEnum = "REMAINING"
    AddNetworkResponseStateEnumTimeout AddNetworkResponseStateEnum = "TIMEOUT"
)

type AddNetworkResponse struct {


	RequestId *string `json:"requestId,omitempty"`

	ErrorMessage *string `json:"errorMessage,omitempty"`

	ErrorCode *string `json:"errorCode,omitempty"`

	State *AddNetworkResponseStateEnum `json:"state,omitempty"`

	Body *string `json:"body,omitempty"`
}

func (s AddNetworkResponse) String() string {
	return utils.Beautify(s)
}

func (s AddNetworkResponse) GoString() string {
	return s.String()
}

func (s AddNetworkResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *AddNetworkResponse) SetRequestId(v string) *AddNetworkResponse {
	s.RequestId = &v
	return s
}

func (s *AddNetworkResponse) SetErrorMessage(v string) *AddNetworkResponse {
	s.ErrorMessage = &v
	return s
}

func (s *AddNetworkResponse) SetErrorCode(v string) *AddNetworkResponse {
	s.ErrorCode = &v
	return s
}

func (s *AddNetworkResponse) SetState(v AddNetworkResponseStateEnum) *AddNetworkResponse {
	s.State = &v
	return s
}

func (s *AddNetworkResponse) SetBody(v string) *AddNetworkResponse {
	s.Body = &v
	return s
}


type AddNetworkResponseBuilder struct {
	s *AddNetworkResponse
}

func NewAddNetworkResponseBuilder() *AddNetworkResponseBuilder {
	s := &AddNetworkResponse{}
	b := &AddNetworkResponseBuilder{s: s}
	return b
}

func (b *AddNetworkResponseBuilder) RequestId(v string) *AddNetworkResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *AddNetworkResponseBuilder) ErrorMessage(v string) *AddNetworkResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *AddNetworkResponseBuilder) ErrorCode(v string) *AddNetworkResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *AddNetworkResponseBuilder) State(v AddNetworkResponseStateEnum) *AddNetworkResponseBuilder {
	b.s.State = &v
	return b
}

func (b *AddNetworkResponseBuilder) Body(v string) *AddNetworkResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *AddNetworkResponseBuilder) Build() *AddNetworkResponse {
	return b.s
}


