// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type GetTwoFactorAuthResponseStateEnum string

// List of State
const (
    GetTwoFactorAuthResponseStateEnumOk GetTwoFactorAuthResponseStateEnum = "OK"
    GetTwoFactorAuthResponseStateEnumError GetTwoFactorAuthResponseStateEnum = "ERROR"
    GetTwoFactorAuthResponseStateEnumException GetTwoFactorAuthResponseStateEnum = "EXCEPTION"
    GetTwoFactorAuthResponseStateEnumForbidden GetTwoFactorAuthResponseStateEnum = "FORBIDDEN"
    GetTwoFactorAuthResponseStateEnumRemaining GetTwoFactorAuthResponseStateEnum = "REMAINING"
    GetTwoFactorAuthResponseStateEnumTimeout GetTwoFactorAuthResponseStateEnum = "TIMEOUT"
)

type GetTwoFactorAuthResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *GetTwoFactorAuthResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *GetTwoFactorAuthResponseBody `json:"body,omitempty"`
}

func (s GetTwoFactorAuthResponse) String() string {
	return utils.Beautify(s)
}

func (s GetTwoFactorAuthResponse) GoString() string {
	return s.String()
}

func (s GetTwoFactorAuthResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetTwoFactorAuthResponse) SetRequestId(v string) *GetTwoFactorAuthResponse {
	s.RequestId = &v
	return s
}

func (s *GetTwoFactorAuthResponse) SetErrorMessage(v string) *GetTwoFactorAuthResponse {
	s.ErrorMessage = &v
	return s
}

func (s *GetTwoFactorAuthResponse) SetErrorCode(v string) *GetTwoFactorAuthResponse {
	s.ErrorCode = &v
	return s
}

func (s *GetTwoFactorAuthResponse) SetState(v GetTwoFactorAuthResponseStateEnum) *GetTwoFactorAuthResponse {
	s.State = &v
	return s
}

func (s *GetTwoFactorAuthResponse) SetBody(v *GetTwoFactorAuthResponseBody) *GetTwoFactorAuthResponse {
	s.Body = v
	return s
}


type GetTwoFactorAuthResponseBuilder struct {
	s *GetTwoFactorAuthResponse
}

func NewGetTwoFactorAuthResponseBuilder() *GetTwoFactorAuthResponseBuilder {
	s := &GetTwoFactorAuthResponse{}
	b := &GetTwoFactorAuthResponseBuilder{s: s}
	return b
}

func (b *GetTwoFactorAuthResponseBuilder) RequestId(v string) *GetTwoFactorAuthResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *GetTwoFactorAuthResponseBuilder) ErrorMessage(v string) *GetTwoFactorAuthResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *GetTwoFactorAuthResponseBuilder) ErrorCode(v string) *GetTwoFactorAuthResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *GetTwoFactorAuthResponseBuilder) State(v GetTwoFactorAuthResponseStateEnum) *GetTwoFactorAuthResponseBuilder {
	b.s.State = &v
	return b
}

func (b *GetTwoFactorAuthResponseBuilder) Body(v *GetTwoFactorAuthResponseBody) *GetTwoFactorAuthResponseBuilder {
	b.s.Body = v
	return b
}

func (b *GetTwoFactorAuthResponseBuilder) Build() *GetTwoFactorAuthResponse {
	return b.s
}


