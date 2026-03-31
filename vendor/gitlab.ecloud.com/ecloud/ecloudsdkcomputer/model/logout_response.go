// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type LogoutResponseStateEnum string

// List of State
const (
    LogoutResponseStateEnumOk LogoutResponseStateEnum = "OK"
    LogoutResponseStateEnumError LogoutResponseStateEnum = "ERROR"
    LogoutResponseStateEnumException LogoutResponseStateEnum = "EXCEPTION"
    LogoutResponseStateEnumForbidden LogoutResponseStateEnum = "FORBIDDEN"
    LogoutResponseStateEnumRemaining LogoutResponseStateEnum = "REMAINING"
    LogoutResponseStateEnumTimeout LogoutResponseStateEnum = "TIMEOUT"
)

type LogoutResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *LogoutResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *string `json:"body,omitempty"`
}

func (s LogoutResponse) String() string {
	return utils.Beautify(s)
}

func (s LogoutResponse) GoString() string {
	return s.String()
}

func (s LogoutResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *LogoutResponse) SetRequestId(v string) *LogoutResponse {
	s.RequestId = &v
	return s
}

func (s *LogoutResponse) SetErrorMessage(v string) *LogoutResponse {
	s.ErrorMessage = &v
	return s
}

func (s *LogoutResponse) SetErrorCode(v string) *LogoutResponse {
	s.ErrorCode = &v
	return s
}

func (s *LogoutResponse) SetState(v LogoutResponseStateEnum) *LogoutResponse {
	s.State = &v
	return s
}

func (s *LogoutResponse) SetBody(v string) *LogoutResponse {
	s.Body = &v
	return s
}


type LogoutResponseBuilder struct {
	s *LogoutResponse
}

func NewLogoutResponseBuilder() *LogoutResponseBuilder {
	s := &LogoutResponse{}
	b := &LogoutResponseBuilder{s: s}
	return b
}

func (b *LogoutResponseBuilder) RequestId(v string) *LogoutResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *LogoutResponseBuilder) ErrorMessage(v string) *LogoutResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *LogoutResponseBuilder) ErrorCode(v string) *LogoutResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *LogoutResponseBuilder) State(v LogoutResponseStateEnum) *LogoutResponseBuilder {
	b.s.State = &v
	return b
}

func (b *LogoutResponseBuilder) Body(v string) *LogoutResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *LogoutResponseBuilder) Build() *LogoutResponse {
	return b.s
}


