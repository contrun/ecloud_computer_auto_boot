// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type SetUserModifyPasswordResponseStateEnum string

// List of State
const (
    SetUserModifyPasswordResponseStateEnumOk SetUserModifyPasswordResponseStateEnum = "OK"
    SetUserModifyPasswordResponseStateEnumError SetUserModifyPasswordResponseStateEnum = "ERROR"
    SetUserModifyPasswordResponseStateEnumException SetUserModifyPasswordResponseStateEnum = "EXCEPTION"
    SetUserModifyPasswordResponseStateEnumForbidden SetUserModifyPasswordResponseStateEnum = "FORBIDDEN"
    SetUserModifyPasswordResponseStateEnumRemaining SetUserModifyPasswordResponseStateEnum = "REMAINING"
    SetUserModifyPasswordResponseStateEnumTimeout SetUserModifyPasswordResponseStateEnum = "TIMEOUT"
)

type SetUserModifyPasswordResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *SetUserModifyPasswordResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *bool `json:"body,omitempty"`
}

func (s SetUserModifyPasswordResponse) String() string {
	return utils.Beautify(s)
}

func (s SetUserModifyPasswordResponse) GoString() string {
	return s.String()
}

func (s SetUserModifyPasswordResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *SetUserModifyPasswordResponse) SetRequestId(v string) *SetUserModifyPasswordResponse {
	s.RequestId = &v
	return s
}

func (s *SetUserModifyPasswordResponse) SetErrorMessage(v string) *SetUserModifyPasswordResponse {
	s.ErrorMessage = &v
	return s
}

func (s *SetUserModifyPasswordResponse) SetErrorCode(v string) *SetUserModifyPasswordResponse {
	s.ErrorCode = &v
	return s
}

func (s *SetUserModifyPasswordResponse) SetState(v SetUserModifyPasswordResponseStateEnum) *SetUserModifyPasswordResponse {
	s.State = &v
	return s
}

func (s *SetUserModifyPasswordResponse) SetBody(v bool) *SetUserModifyPasswordResponse {
	s.Body = &v
	return s
}


type SetUserModifyPasswordResponseBuilder struct {
	s *SetUserModifyPasswordResponse
}

func NewSetUserModifyPasswordResponseBuilder() *SetUserModifyPasswordResponseBuilder {
	s := &SetUserModifyPasswordResponse{}
	b := &SetUserModifyPasswordResponseBuilder{s: s}
	return b
}

func (b *SetUserModifyPasswordResponseBuilder) RequestId(v string) *SetUserModifyPasswordResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *SetUserModifyPasswordResponseBuilder) ErrorMessage(v string) *SetUserModifyPasswordResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *SetUserModifyPasswordResponseBuilder) ErrorCode(v string) *SetUserModifyPasswordResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *SetUserModifyPasswordResponseBuilder) State(v SetUserModifyPasswordResponseStateEnum) *SetUserModifyPasswordResponseBuilder {
	b.s.State = &v
	return b
}

func (b *SetUserModifyPasswordResponseBuilder) Body(v bool) *SetUserModifyPasswordResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *SetUserModifyPasswordResponseBuilder) Build() *SetUserModifyPasswordResponse {
	return b.s
}


