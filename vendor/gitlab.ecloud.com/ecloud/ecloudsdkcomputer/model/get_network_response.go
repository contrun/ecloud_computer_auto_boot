// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type GetNetworkResponseStateEnum string

// List of State
const (
    GetNetworkResponseStateEnumOk GetNetworkResponseStateEnum = "OK"
    GetNetworkResponseStateEnumError GetNetworkResponseStateEnum = "ERROR"
    GetNetworkResponseStateEnumException GetNetworkResponseStateEnum = "EXCEPTION"
    GetNetworkResponseStateEnumForbidden GetNetworkResponseStateEnum = "FORBIDDEN"
    GetNetworkResponseStateEnumRemaining GetNetworkResponseStateEnum = "REMAINING"
    GetNetworkResponseStateEnumTimeout GetNetworkResponseStateEnum = "TIMEOUT"
)

type GetNetworkResponse struct {


	RequestId *string `json:"requestId,omitempty"`

	ErrorMessage *string `json:"errorMessage,omitempty"`

	ErrorCode *string `json:"errorCode,omitempty"`

	State *GetNetworkResponseStateEnum `json:"state,omitempty"`

	Body *GetNetworkResponseBody `json:"body,omitempty"`
}

func (s GetNetworkResponse) String() string {
	return utils.Beautify(s)
}

func (s GetNetworkResponse) GoString() string {
	return s.String()
}

func (s GetNetworkResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetNetworkResponse) SetRequestId(v string) *GetNetworkResponse {
	s.RequestId = &v
	return s
}

func (s *GetNetworkResponse) SetErrorMessage(v string) *GetNetworkResponse {
	s.ErrorMessage = &v
	return s
}

func (s *GetNetworkResponse) SetErrorCode(v string) *GetNetworkResponse {
	s.ErrorCode = &v
	return s
}

func (s *GetNetworkResponse) SetState(v GetNetworkResponseStateEnum) *GetNetworkResponse {
	s.State = &v
	return s
}

func (s *GetNetworkResponse) SetBody(v *GetNetworkResponseBody) *GetNetworkResponse {
	s.Body = v
	return s
}


type GetNetworkResponseBuilder struct {
	s *GetNetworkResponse
}

func NewGetNetworkResponseBuilder() *GetNetworkResponseBuilder {
	s := &GetNetworkResponse{}
	b := &GetNetworkResponseBuilder{s: s}
	return b
}

func (b *GetNetworkResponseBuilder) RequestId(v string) *GetNetworkResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *GetNetworkResponseBuilder) ErrorMessage(v string) *GetNetworkResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *GetNetworkResponseBuilder) ErrorCode(v string) *GetNetworkResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *GetNetworkResponseBuilder) State(v GetNetworkResponseStateEnum) *GetNetworkResponseBuilder {
	b.s.State = &v
	return b
}

func (b *GetNetworkResponseBuilder) Body(v *GetNetworkResponseBody) *GetNetworkResponseBuilder {
	b.s.Body = v
	return b
}

func (b *GetNetworkResponseBuilder) Build() *GetNetworkResponse {
	return b.s
}


