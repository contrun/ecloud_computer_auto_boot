// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type ShareImageOperationResponseStateEnum string

// List of State
const (
    ShareImageOperationResponseStateEnumOk ShareImageOperationResponseStateEnum = "OK"
    ShareImageOperationResponseStateEnumError ShareImageOperationResponseStateEnum = "ERROR"
    ShareImageOperationResponseStateEnumException ShareImageOperationResponseStateEnum = "EXCEPTION"
    ShareImageOperationResponseStateEnumForbidden ShareImageOperationResponseStateEnum = "FORBIDDEN"
    ShareImageOperationResponseStateEnumRemaining ShareImageOperationResponseStateEnum = "REMAINING"
    ShareImageOperationResponseStateEnumTimeout ShareImageOperationResponseStateEnum = "TIMEOUT"
)

type ShareImageOperationResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *ShareImageOperationResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *bool `json:"body,omitempty"`
}

func (s ShareImageOperationResponse) String() string {
	return utils.Beautify(s)
}

func (s ShareImageOperationResponse) GoString() string {
	return s.String()
}

func (s ShareImageOperationResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ShareImageOperationResponse) SetRequestId(v string) *ShareImageOperationResponse {
	s.RequestId = &v
	return s
}

func (s *ShareImageOperationResponse) SetErrorMessage(v string) *ShareImageOperationResponse {
	s.ErrorMessage = &v
	return s
}

func (s *ShareImageOperationResponse) SetErrorCode(v string) *ShareImageOperationResponse {
	s.ErrorCode = &v
	return s
}

func (s *ShareImageOperationResponse) SetState(v ShareImageOperationResponseStateEnum) *ShareImageOperationResponse {
	s.State = &v
	return s
}

func (s *ShareImageOperationResponse) SetBody(v bool) *ShareImageOperationResponse {
	s.Body = &v
	return s
}


type ShareImageOperationResponseBuilder struct {
	s *ShareImageOperationResponse
}

func NewShareImageOperationResponseBuilder() *ShareImageOperationResponseBuilder {
	s := &ShareImageOperationResponse{}
	b := &ShareImageOperationResponseBuilder{s: s}
	return b
}

func (b *ShareImageOperationResponseBuilder) RequestId(v string) *ShareImageOperationResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *ShareImageOperationResponseBuilder) ErrorMessage(v string) *ShareImageOperationResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *ShareImageOperationResponseBuilder) ErrorCode(v string) *ShareImageOperationResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *ShareImageOperationResponseBuilder) State(v ShareImageOperationResponseStateEnum) *ShareImageOperationResponseBuilder {
	b.s.State = &v
	return b
}

func (b *ShareImageOperationResponseBuilder) Body(v bool) *ShareImageOperationResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *ShareImageOperationResponseBuilder) Build() *ShareImageOperationResponse {
	return b.s
}


