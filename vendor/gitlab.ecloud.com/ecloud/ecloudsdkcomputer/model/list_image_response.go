// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type ListImageResponseStateEnum string

// List of State
const (
    ListImageResponseStateEnumOk ListImageResponseStateEnum = "OK"
    ListImageResponseStateEnumError ListImageResponseStateEnum = "ERROR"
    ListImageResponseStateEnumException ListImageResponseStateEnum = "EXCEPTION"
    ListImageResponseStateEnumForbidden ListImageResponseStateEnum = "FORBIDDEN"
    ListImageResponseStateEnumRemaining ListImageResponseStateEnum = "REMAINING"
    ListImageResponseStateEnumTimeout ListImageResponseStateEnum = "TIMEOUT"
)

type ListImageResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *ListImageResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *ListImageResponseBody `json:"body,omitempty"`
}

func (s ListImageResponse) String() string {
	return utils.Beautify(s)
}

func (s ListImageResponse) GoString() string {
	return s.String()
}

func (s ListImageResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListImageResponse) SetRequestId(v string) *ListImageResponse {
	s.RequestId = &v
	return s
}

func (s *ListImageResponse) SetErrorMessage(v string) *ListImageResponse {
	s.ErrorMessage = &v
	return s
}

func (s *ListImageResponse) SetErrorCode(v string) *ListImageResponse {
	s.ErrorCode = &v
	return s
}

func (s *ListImageResponse) SetState(v ListImageResponseStateEnum) *ListImageResponse {
	s.State = &v
	return s
}

func (s *ListImageResponse) SetBody(v *ListImageResponseBody) *ListImageResponse {
	s.Body = v
	return s
}


type ListImageResponseBuilder struct {
	s *ListImageResponse
}

func NewListImageResponseBuilder() *ListImageResponseBuilder {
	s := &ListImageResponse{}
	b := &ListImageResponseBuilder{s: s}
	return b
}

func (b *ListImageResponseBuilder) RequestId(v string) *ListImageResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *ListImageResponseBuilder) ErrorMessage(v string) *ListImageResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *ListImageResponseBuilder) ErrorCode(v string) *ListImageResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *ListImageResponseBuilder) State(v ListImageResponseStateEnum) *ListImageResponseBuilder {
	b.s.State = &v
	return b
}

func (b *ListImageResponseBuilder) Body(v *ListImageResponseBody) *ListImageResponseBuilder {
	b.s.Body = v
	return b
}

func (b *ListImageResponseBuilder) Build() *ListImageResponse {
	return b.s
}


