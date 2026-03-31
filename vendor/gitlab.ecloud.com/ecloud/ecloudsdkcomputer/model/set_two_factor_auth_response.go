// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type SetTwoFactorAuthResponseStateEnum string

// List of State
const (
    SetTwoFactorAuthResponseStateEnumOk SetTwoFactorAuthResponseStateEnum = "OK"
    SetTwoFactorAuthResponseStateEnumError SetTwoFactorAuthResponseStateEnum = "ERROR"
    SetTwoFactorAuthResponseStateEnumException SetTwoFactorAuthResponseStateEnum = "EXCEPTION"
    SetTwoFactorAuthResponseStateEnumForbidden SetTwoFactorAuthResponseStateEnum = "FORBIDDEN"
    SetTwoFactorAuthResponseStateEnumRemaining SetTwoFactorAuthResponseStateEnum = "REMAINING"
    SetTwoFactorAuthResponseStateEnumTimeout SetTwoFactorAuthResponseStateEnum = "TIMEOUT"
)

type SetTwoFactorAuthResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *SetTwoFactorAuthResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *bool `json:"body,omitempty"`
}

func (s SetTwoFactorAuthResponse) String() string {
	return utils.Beautify(s)
}

func (s SetTwoFactorAuthResponse) GoString() string {
	return s.String()
}

func (s SetTwoFactorAuthResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *SetTwoFactorAuthResponse) SetRequestId(v string) *SetTwoFactorAuthResponse {
	s.RequestId = &v
	return s
}

func (s *SetTwoFactorAuthResponse) SetErrorMessage(v string) *SetTwoFactorAuthResponse {
	s.ErrorMessage = &v
	return s
}

func (s *SetTwoFactorAuthResponse) SetErrorCode(v string) *SetTwoFactorAuthResponse {
	s.ErrorCode = &v
	return s
}

func (s *SetTwoFactorAuthResponse) SetState(v SetTwoFactorAuthResponseStateEnum) *SetTwoFactorAuthResponse {
	s.State = &v
	return s
}

func (s *SetTwoFactorAuthResponse) SetBody(v bool) *SetTwoFactorAuthResponse {
	s.Body = &v
	return s
}


type SetTwoFactorAuthResponseBuilder struct {
	s *SetTwoFactorAuthResponse
}

func NewSetTwoFactorAuthResponseBuilder() *SetTwoFactorAuthResponseBuilder {
	s := &SetTwoFactorAuthResponse{}
	b := &SetTwoFactorAuthResponseBuilder{s: s}
	return b
}

func (b *SetTwoFactorAuthResponseBuilder) RequestId(v string) *SetTwoFactorAuthResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *SetTwoFactorAuthResponseBuilder) ErrorMessage(v string) *SetTwoFactorAuthResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *SetTwoFactorAuthResponseBuilder) ErrorCode(v string) *SetTwoFactorAuthResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *SetTwoFactorAuthResponseBuilder) State(v SetTwoFactorAuthResponseStateEnum) *SetTwoFactorAuthResponseBuilder {
	b.s.State = &v
	return b
}

func (b *SetTwoFactorAuthResponseBuilder) Body(v bool) *SetTwoFactorAuthResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *SetTwoFactorAuthResponseBuilder) Build() *SetTwoFactorAuthResponse {
	return b.s
}


