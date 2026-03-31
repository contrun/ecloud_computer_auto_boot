// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type SetUnbindMobileFlagResponseStateEnum string

// List of State
const (
    SetUnbindMobileFlagResponseStateEnumOk SetUnbindMobileFlagResponseStateEnum = "OK"
    SetUnbindMobileFlagResponseStateEnumError SetUnbindMobileFlagResponseStateEnum = "ERROR"
    SetUnbindMobileFlagResponseStateEnumException SetUnbindMobileFlagResponseStateEnum = "EXCEPTION"
    SetUnbindMobileFlagResponseStateEnumForbidden SetUnbindMobileFlagResponseStateEnum = "FORBIDDEN"
    SetUnbindMobileFlagResponseStateEnumRemaining SetUnbindMobileFlagResponseStateEnum = "REMAINING"
    SetUnbindMobileFlagResponseStateEnumTimeout SetUnbindMobileFlagResponseStateEnum = "TIMEOUT"
)

type SetUnbindMobileFlagResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *SetUnbindMobileFlagResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *bool `json:"body,omitempty"`
}

func (s SetUnbindMobileFlagResponse) String() string {
	return utils.Beautify(s)
}

func (s SetUnbindMobileFlagResponse) GoString() string {
	return s.String()
}

func (s SetUnbindMobileFlagResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *SetUnbindMobileFlagResponse) SetRequestId(v string) *SetUnbindMobileFlagResponse {
	s.RequestId = &v
	return s
}

func (s *SetUnbindMobileFlagResponse) SetErrorMessage(v string) *SetUnbindMobileFlagResponse {
	s.ErrorMessage = &v
	return s
}

func (s *SetUnbindMobileFlagResponse) SetErrorCode(v string) *SetUnbindMobileFlagResponse {
	s.ErrorCode = &v
	return s
}

func (s *SetUnbindMobileFlagResponse) SetState(v SetUnbindMobileFlagResponseStateEnum) *SetUnbindMobileFlagResponse {
	s.State = &v
	return s
}

func (s *SetUnbindMobileFlagResponse) SetBody(v bool) *SetUnbindMobileFlagResponse {
	s.Body = &v
	return s
}


type SetUnbindMobileFlagResponseBuilder struct {
	s *SetUnbindMobileFlagResponse
}

func NewSetUnbindMobileFlagResponseBuilder() *SetUnbindMobileFlagResponseBuilder {
	s := &SetUnbindMobileFlagResponse{}
	b := &SetUnbindMobileFlagResponseBuilder{s: s}
	return b
}

func (b *SetUnbindMobileFlagResponseBuilder) RequestId(v string) *SetUnbindMobileFlagResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *SetUnbindMobileFlagResponseBuilder) ErrorMessage(v string) *SetUnbindMobileFlagResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *SetUnbindMobileFlagResponseBuilder) ErrorCode(v string) *SetUnbindMobileFlagResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *SetUnbindMobileFlagResponseBuilder) State(v SetUnbindMobileFlagResponseStateEnum) *SetUnbindMobileFlagResponseBuilder {
	b.s.State = &v
	return b
}

func (b *SetUnbindMobileFlagResponseBuilder) Body(v bool) *SetUnbindMobileFlagResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *SetUnbindMobileFlagResponseBuilder) Build() *SetUnbindMobileFlagResponse {
	return b.s
}


