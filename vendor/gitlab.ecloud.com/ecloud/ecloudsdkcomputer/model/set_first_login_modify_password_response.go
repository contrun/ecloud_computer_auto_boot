// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type SetFirstLoginModifyPasswordResponseStateEnum string

// List of State
const (
    SetFirstLoginModifyPasswordResponseStateEnumOk SetFirstLoginModifyPasswordResponseStateEnum = "OK"
    SetFirstLoginModifyPasswordResponseStateEnumError SetFirstLoginModifyPasswordResponseStateEnum = "ERROR"
    SetFirstLoginModifyPasswordResponseStateEnumException SetFirstLoginModifyPasswordResponseStateEnum = "EXCEPTION"
    SetFirstLoginModifyPasswordResponseStateEnumForbidden SetFirstLoginModifyPasswordResponseStateEnum = "FORBIDDEN"
    SetFirstLoginModifyPasswordResponseStateEnumRemaining SetFirstLoginModifyPasswordResponseStateEnum = "REMAINING"
    SetFirstLoginModifyPasswordResponseStateEnumTimeout SetFirstLoginModifyPasswordResponseStateEnum = "TIMEOUT"
)

type SetFirstLoginModifyPasswordResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *SetFirstLoginModifyPasswordResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *bool `json:"body,omitempty"`
}

func (s SetFirstLoginModifyPasswordResponse) String() string {
	return utils.Beautify(s)
}

func (s SetFirstLoginModifyPasswordResponse) GoString() string {
	return s.String()
}

func (s SetFirstLoginModifyPasswordResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *SetFirstLoginModifyPasswordResponse) SetRequestId(v string) *SetFirstLoginModifyPasswordResponse {
	s.RequestId = &v
	return s
}

func (s *SetFirstLoginModifyPasswordResponse) SetErrorMessage(v string) *SetFirstLoginModifyPasswordResponse {
	s.ErrorMessage = &v
	return s
}

func (s *SetFirstLoginModifyPasswordResponse) SetErrorCode(v string) *SetFirstLoginModifyPasswordResponse {
	s.ErrorCode = &v
	return s
}

func (s *SetFirstLoginModifyPasswordResponse) SetState(v SetFirstLoginModifyPasswordResponseStateEnum) *SetFirstLoginModifyPasswordResponse {
	s.State = &v
	return s
}

func (s *SetFirstLoginModifyPasswordResponse) SetBody(v bool) *SetFirstLoginModifyPasswordResponse {
	s.Body = &v
	return s
}


type SetFirstLoginModifyPasswordResponseBuilder struct {
	s *SetFirstLoginModifyPasswordResponse
}

func NewSetFirstLoginModifyPasswordResponseBuilder() *SetFirstLoginModifyPasswordResponseBuilder {
	s := &SetFirstLoginModifyPasswordResponse{}
	b := &SetFirstLoginModifyPasswordResponseBuilder{s: s}
	return b
}

func (b *SetFirstLoginModifyPasswordResponseBuilder) RequestId(v string) *SetFirstLoginModifyPasswordResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *SetFirstLoginModifyPasswordResponseBuilder) ErrorMessage(v string) *SetFirstLoginModifyPasswordResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *SetFirstLoginModifyPasswordResponseBuilder) ErrorCode(v string) *SetFirstLoginModifyPasswordResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *SetFirstLoginModifyPasswordResponseBuilder) State(v SetFirstLoginModifyPasswordResponseStateEnum) *SetFirstLoginModifyPasswordResponseBuilder {
	b.s.State = &v
	return b
}

func (b *SetFirstLoginModifyPasswordResponseBuilder) Body(v bool) *SetFirstLoginModifyPasswordResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *SetFirstLoginModifyPasswordResponseBuilder) Build() *SetFirstLoginModifyPasswordResponse {
	return b.s
}


