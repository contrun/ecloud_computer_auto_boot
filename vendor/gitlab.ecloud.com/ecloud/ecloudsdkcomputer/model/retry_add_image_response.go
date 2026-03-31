// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type RetryAddImageResponseStateEnum string

// List of State
const (
    RetryAddImageResponseStateEnumOk RetryAddImageResponseStateEnum = "OK"
    RetryAddImageResponseStateEnumError RetryAddImageResponseStateEnum = "ERROR"
    RetryAddImageResponseStateEnumException RetryAddImageResponseStateEnum = "EXCEPTION"
    RetryAddImageResponseStateEnumForbidden RetryAddImageResponseStateEnum = "FORBIDDEN"
    RetryAddImageResponseStateEnumRemaining RetryAddImageResponseStateEnum = "REMAINING"
    RetryAddImageResponseStateEnumTimeout RetryAddImageResponseStateEnum = "TIMEOUT"
)

type RetryAddImageResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *RetryAddImageResponseStateEnum `json:"state,omitempty"`
    //  请求成功时返回的数据
	Body *RetryAddImageResponseBody `json:"body,omitempty"`
}

func (s RetryAddImageResponse) String() string {
	return utils.Beautify(s)
}

func (s RetryAddImageResponse) GoString() string {
	return s.String()
}

func (s RetryAddImageResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *RetryAddImageResponse) SetRequestId(v string) *RetryAddImageResponse {
	s.RequestId = &v
	return s
}

func (s *RetryAddImageResponse) SetErrorMessage(v string) *RetryAddImageResponse {
	s.ErrorMessage = &v
	return s
}

func (s *RetryAddImageResponse) SetErrorCode(v string) *RetryAddImageResponse {
	s.ErrorCode = &v
	return s
}

func (s *RetryAddImageResponse) SetState(v RetryAddImageResponseStateEnum) *RetryAddImageResponse {
	s.State = &v
	return s
}

func (s *RetryAddImageResponse) SetBody(v *RetryAddImageResponseBody) *RetryAddImageResponse {
	s.Body = v
	return s
}


type RetryAddImageResponseBuilder struct {
	s *RetryAddImageResponse
}

func NewRetryAddImageResponseBuilder() *RetryAddImageResponseBuilder {
	s := &RetryAddImageResponse{}
	b := &RetryAddImageResponseBuilder{s: s}
	return b
}

func (b *RetryAddImageResponseBuilder) RequestId(v string) *RetryAddImageResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *RetryAddImageResponseBuilder) ErrorMessage(v string) *RetryAddImageResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *RetryAddImageResponseBuilder) ErrorCode(v string) *RetryAddImageResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *RetryAddImageResponseBuilder) State(v RetryAddImageResponseStateEnum) *RetryAddImageResponseBuilder {
	b.s.State = &v
	return b
}

func (b *RetryAddImageResponseBuilder) Body(v *RetryAddImageResponseBody) *RetryAddImageResponseBuilder {
	b.s.Body = v
	return b
}

func (b *RetryAddImageResponseBuilder) Build() *RetryAddImageResponse {
	return b.s
}


