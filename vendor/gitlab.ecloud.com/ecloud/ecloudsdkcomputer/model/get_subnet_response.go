// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type GetSubnetResponseStateEnum string

// List of State
const (
    GetSubnetResponseStateEnumOk GetSubnetResponseStateEnum = "OK"
    GetSubnetResponseStateEnumError GetSubnetResponseStateEnum = "ERROR"
    GetSubnetResponseStateEnumException GetSubnetResponseStateEnum = "EXCEPTION"
    GetSubnetResponseStateEnumForbidden GetSubnetResponseStateEnum = "FORBIDDEN"
    GetSubnetResponseStateEnumRemaining GetSubnetResponseStateEnum = "REMAINING"
    GetSubnetResponseStateEnumTimeout GetSubnetResponseStateEnum = "TIMEOUT"
)

type GetSubnetResponse struct {


	RequestId *string `json:"requestId,omitempty"`

	ErrorMessage *string `json:"errorMessage,omitempty"`

	ErrorCode *string `json:"errorCode,omitempty"`

	State *GetSubnetResponseStateEnum `json:"state,omitempty"`

	Body *[]GetSubnetResponseBody `json:"body,omitempty"`
}

func (s GetSubnetResponse) String() string {
	return utils.Beautify(s)
}

func (s GetSubnetResponse) GoString() string {
	return s.String()
}

func (s GetSubnetResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetSubnetResponse) SetRequestId(v string) *GetSubnetResponse {
	s.RequestId = &v
	return s
}

func (s *GetSubnetResponse) SetErrorMessage(v string) *GetSubnetResponse {
	s.ErrorMessage = &v
	return s
}

func (s *GetSubnetResponse) SetErrorCode(v string) *GetSubnetResponse {
	s.ErrorCode = &v
	return s
}

func (s *GetSubnetResponse) SetState(v GetSubnetResponseStateEnum) *GetSubnetResponse {
	s.State = &v
	return s
}

func (s *GetSubnetResponse) SetBody(v []GetSubnetResponseBody) *GetSubnetResponse {
	s.Body = &v
	return s
}


type GetSubnetResponseBuilder struct {
	s *GetSubnetResponse
}

func NewGetSubnetResponseBuilder() *GetSubnetResponseBuilder {
	s := &GetSubnetResponse{}
	b := &GetSubnetResponseBuilder{s: s}
	return b
}

func (b *GetSubnetResponseBuilder) RequestId(v string) *GetSubnetResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *GetSubnetResponseBuilder) ErrorMessage(v string) *GetSubnetResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *GetSubnetResponseBuilder) ErrorCode(v string) *GetSubnetResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *GetSubnetResponseBuilder) State(v GetSubnetResponseStateEnum) *GetSubnetResponseBuilder {
	b.s.State = &v
	return b
}

func (b *GetSubnetResponseBuilder) Body(v []GetSubnetResponseBody) *GetSubnetResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *GetSubnetResponseBuilder) Build() *GetSubnetResponse {
	return b.s
}


