// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type GetVpcResponseStateEnum string

// List of State
const (
    GetVpcResponseStateEnumOk GetVpcResponseStateEnum = "OK"
    GetVpcResponseStateEnumError GetVpcResponseStateEnum = "ERROR"
    GetVpcResponseStateEnumException GetVpcResponseStateEnum = "EXCEPTION"
    GetVpcResponseStateEnumForbidden GetVpcResponseStateEnum = "FORBIDDEN"
    GetVpcResponseStateEnumRemaining GetVpcResponseStateEnum = "REMAINING"
    GetVpcResponseStateEnumTimeout GetVpcResponseStateEnum = "TIMEOUT"
)

type GetVpcResponse struct {


	RequestId *string `json:"requestId,omitempty"`

	ErrorMessage *string `json:"errorMessage,omitempty"`

	ErrorCode *string `json:"errorCode,omitempty"`

	State *GetVpcResponseStateEnum `json:"state,omitempty"`

	Body *GetVpcResponseBody `json:"body,omitempty"`
}

func (s GetVpcResponse) String() string {
	return utils.Beautify(s)
}

func (s GetVpcResponse) GoString() string {
	return s.String()
}

func (s GetVpcResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetVpcResponse) SetRequestId(v string) *GetVpcResponse {
	s.RequestId = &v
	return s
}

func (s *GetVpcResponse) SetErrorMessage(v string) *GetVpcResponse {
	s.ErrorMessage = &v
	return s
}

func (s *GetVpcResponse) SetErrorCode(v string) *GetVpcResponse {
	s.ErrorCode = &v
	return s
}

func (s *GetVpcResponse) SetState(v GetVpcResponseStateEnum) *GetVpcResponse {
	s.State = &v
	return s
}

func (s *GetVpcResponse) SetBody(v *GetVpcResponseBody) *GetVpcResponse {
	s.Body = v
	return s
}


type GetVpcResponseBuilder struct {
	s *GetVpcResponse
}

func NewGetVpcResponseBuilder() *GetVpcResponseBuilder {
	s := &GetVpcResponse{}
	b := &GetVpcResponseBuilder{s: s}
	return b
}

func (b *GetVpcResponseBuilder) RequestId(v string) *GetVpcResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *GetVpcResponseBuilder) ErrorMessage(v string) *GetVpcResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *GetVpcResponseBuilder) ErrorCode(v string) *GetVpcResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *GetVpcResponseBuilder) State(v GetVpcResponseStateEnum) *GetVpcResponseBuilder {
	b.s.State = &v
	return b
}

func (b *GetVpcResponseBuilder) Body(v *GetVpcResponseBody) *GetVpcResponseBuilder {
	b.s.Body = v
	return b
}

func (b *GetVpcResponseBuilder) Build() *GetVpcResponse {
	return b.s
}


