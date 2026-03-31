// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type GetCloudbandwidthResponseStateEnum string

// List of State
const (
    GetCloudbandwidthResponseStateEnumOk GetCloudbandwidthResponseStateEnum = "OK"
    GetCloudbandwidthResponseStateEnumError GetCloudbandwidthResponseStateEnum = "ERROR"
    GetCloudbandwidthResponseStateEnumException GetCloudbandwidthResponseStateEnum = "EXCEPTION"
    GetCloudbandwidthResponseStateEnumForbidden GetCloudbandwidthResponseStateEnum = "FORBIDDEN"
    GetCloudbandwidthResponseStateEnumRemaining GetCloudbandwidthResponseStateEnum = "REMAINING"
    GetCloudbandwidthResponseStateEnumTimeout GetCloudbandwidthResponseStateEnum = "TIMEOUT"
)

type GetCloudbandwidthResponse struct {


	RequestId *string `json:"requestId,omitempty"`

	ErrorMessage *string `json:"errorMessage,omitempty"`

	ErrorCode *string `json:"errorCode,omitempty"`

	State *GetCloudbandwidthResponseStateEnum `json:"state,omitempty"`

	Body *[]GetCloudbandwidthResponseBody `json:"body,omitempty"`
}

func (s GetCloudbandwidthResponse) String() string {
	return utils.Beautify(s)
}

func (s GetCloudbandwidthResponse) GoString() string {
	return s.String()
}

func (s GetCloudbandwidthResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetCloudbandwidthResponse) SetRequestId(v string) *GetCloudbandwidthResponse {
	s.RequestId = &v
	return s
}

func (s *GetCloudbandwidthResponse) SetErrorMessage(v string) *GetCloudbandwidthResponse {
	s.ErrorMessage = &v
	return s
}

func (s *GetCloudbandwidthResponse) SetErrorCode(v string) *GetCloudbandwidthResponse {
	s.ErrorCode = &v
	return s
}

func (s *GetCloudbandwidthResponse) SetState(v GetCloudbandwidthResponseStateEnum) *GetCloudbandwidthResponse {
	s.State = &v
	return s
}

func (s *GetCloudbandwidthResponse) SetBody(v []GetCloudbandwidthResponseBody) *GetCloudbandwidthResponse {
	s.Body = &v
	return s
}


type GetCloudbandwidthResponseBuilder struct {
	s *GetCloudbandwidthResponse
}

func NewGetCloudbandwidthResponseBuilder() *GetCloudbandwidthResponseBuilder {
	s := &GetCloudbandwidthResponse{}
	b := &GetCloudbandwidthResponseBuilder{s: s}
	return b
}

func (b *GetCloudbandwidthResponseBuilder) RequestId(v string) *GetCloudbandwidthResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *GetCloudbandwidthResponseBuilder) ErrorMessage(v string) *GetCloudbandwidthResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *GetCloudbandwidthResponseBuilder) ErrorCode(v string) *GetCloudbandwidthResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *GetCloudbandwidthResponseBuilder) State(v GetCloudbandwidthResponseStateEnum) *GetCloudbandwidthResponseBuilder {
	b.s.State = &v
	return b
}

func (b *GetCloudbandwidthResponseBuilder) Body(v []GetCloudbandwidthResponseBody) *GetCloudbandwidthResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *GetCloudbandwidthResponseBuilder) Build() *GetCloudbandwidthResponse {
	return b.s
}


