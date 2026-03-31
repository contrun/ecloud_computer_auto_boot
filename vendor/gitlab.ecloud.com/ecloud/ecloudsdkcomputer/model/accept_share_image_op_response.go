// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type AcceptShareImageOpResponseStateEnum string

// List of State
const (
    AcceptShareImageOpResponseStateEnumOk AcceptShareImageOpResponseStateEnum = "OK"
    AcceptShareImageOpResponseStateEnumError AcceptShareImageOpResponseStateEnum = "ERROR"
    AcceptShareImageOpResponseStateEnumException AcceptShareImageOpResponseStateEnum = "EXCEPTION"
    AcceptShareImageOpResponseStateEnumForbidden AcceptShareImageOpResponseStateEnum = "FORBIDDEN"
    AcceptShareImageOpResponseStateEnumRemaining AcceptShareImageOpResponseStateEnum = "REMAINING"
    AcceptShareImageOpResponseStateEnumTimeout AcceptShareImageOpResponseStateEnum = "TIMEOUT"
)

type AcceptShareImageOpResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *AcceptShareImageOpResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *string `json:"body,omitempty"`
}

func (s AcceptShareImageOpResponse) String() string {
	return utils.Beautify(s)
}

func (s AcceptShareImageOpResponse) GoString() string {
	return s.String()
}

func (s AcceptShareImageOpResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *AcceptShareImageOpResponse) SetRequestId(v string) *AcceptShareImageOpResponse {
	s.RequestId = &v
	return s
}

func (s *AcceptShareImageOpResponse) SetErrorMessage(v string) *AcceptShareImageOpResponse {
	s.ErrorMessage = &v
	return s
}

func (s *AcceptShareImageOpResponse) SetErrorCode(v string) *AcceptShareImageOpResponse {
	s.ErrorCode = &v
	return s
}

func (s *AcceptShareImageOpResponse) SetState(v AcceptShareImageOpResponseStateEnum) *AcceptShareImageOpResponse {
	s.State = &v
	return s
}

func (s *AcceptShareImageOpResponse) SetBody(v string) *AcceptShareImageOpResponse {
	s.Body = &v
	return s
}


type AcceptShareImageOpResponseBuilder struct {
	s *AcceptShareImageOpResponse
}

func NewAcceptShareImageOpResponseBuilder() *AcceptShareImageOpResponseBuilder {
	s := &AcceptShareImageOpResponse{}
	b := &AcceptShareImageOpResponseBuilder{s: s}
	return b
}

func (b *AcceptShareImageOpResponseBuilder) RequestId(v string) *AcceptShareImageOpResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *AcceptShareImageOpResponseBuilder) ErrorMessage(v string) *AcceptShareImageOpResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *AcceptShareImageOpResponseBuilder) ErrorCode(v string) *AcceptShareImageOpResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *AcceptShareImageOpResponseBuilder) State(v AcceptShareImageOpResponseStateEnum) *AcceptShareImageOpResponseBuilder {
	b.s.State = &v
	return b
}

func (b *AcceptShareImageOpResponseBuilder) Body(v string) *AcceptShareImageOpResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *AcceptShareImageOpResponseBuilder) Build() *AcceptShareImageOpResponse {
	return b.s
}


