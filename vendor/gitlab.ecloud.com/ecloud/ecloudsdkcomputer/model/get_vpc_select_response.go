// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type GetVpcSelectResponseStateEnum string

// List of State
const (
    GetVpcSelectResponseStateEnumOk GetVpcSelectResponseStateEnum = "OK"
    GetVpcSelectResponseStateEnumError GetVpcSelectResponseStateEnum = "ERROR"
    GetVpcSelectResponseStateEnumException GetVpcSelectResponseStateEnum = "EXCEPTION"
    GetVpcSelectResponseStateEnumForbidden GetVpcSelectResponseStateEnum = "FORBIDDEN"
    GetVpcSelectResponseStateEnumRemaining GetVpcSelectResponseStateEnum = "REMAINING"
    GetVpcSelectResponseStateEnumTimeout GetVpcSelectResponseStateEnum = "TIMEOUT"
)

type GetVpcSelectResponse struct {


	RequestId *string `json:"requestId,omitempty"`

	ErrorMessage *string `json:"errorMessage,omitempty"`

	ErrorCode *string `json:"errorCode,omitempty"`

	State *GetVpcSelectResponseStateEnum `json:"state,omitempty"`

	Body *GetVpcSelectResponseBody `json:"body,omitempty"`
}

func (s GetVpcSelectResponse) String() string {
	return utils.Beautify(s)
}

func (s GetVpcSelectResponse) GoString() string {
	return s.String()
}

func (s GetVpcSelectResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetVpcSelectResponse) SetRequestId(v string) *GetVpcSelectResponse {
	s.RequestId = &v
	return s
}

func (s *GetVpcSelectResponse) SetErrorMessage(v string) *GetVpcSelectResponse {
	s.ErrorMessage = &v
	return s
}

func (s *GetVpcSelectResponse) SetErrorCode(v string) *GetVpcSelectResponse {
	s.ErrorCode = &v
	return s
}

func (s *GetVpcSelectResponse) SetState(v GetVpcSelectResponseStateEnum) *GetVpcSelectResponse {
	s.State = &v
	return s
}

func (s *GetVpcSelectResponse) SetBody(v *GetVpcSelectResponseBody) *GetVpcSelectResponse {
	s.Body = v
	return s
}


type GetVpcSelectResponseBuilder struct {
	s *GetVpcSelectResponse
}

func NewGetVpcSelectResponseBuilder() *GetVpcSelectResponseBuilder {
	s := &GetVpcSelectResponse{}
	b := &GetVpcSelectResponseBuilder{s: s}
	return b
}

func (b *GetVpcSelectResponseBuilder) RequestId(v string) *GetVpcSelectResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *GetVpcSelectResponseBuilder) ErrorMessage(v string) *GetVpcSelectResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *GetVpcSelectResponseBuilder) ErrorCode(v string) *GetVpcSelectResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *GetVpcSelectResponseBuilder) State(v GetVpcSelectResponseStateEnum) *GetVpcSelectResponseBuilder {
	b.s.State = &v
	return b
}

func (b *GetVpcSelectResponseBuilder) Body(v *GetVpcSelectResponseBody) *GetVpcSelectResponseBuilder {
	b.s.Body = v
	return b
}

func (b *GetVpcSelectResponseBuilder) Build() *GetVpcSelectResponse {
	return b.s
}


