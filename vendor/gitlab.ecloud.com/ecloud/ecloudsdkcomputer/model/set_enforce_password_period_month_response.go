// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type SetEnforcePasswordPeriodMonthResponseStateEnum string

// List of State
const (
    SetEnforcePasswordPeriodMonthResponseStateEnumOk SetEnforcePasswordPeriodMonthResponseStateEnum = "OK"
    SetEnforcePasswordPeriodMonthResponseStateEnumError SetEnforcePasswordPeriodMonthResponseStateEnum = "ERROR"
    SetEnforcePasswordPeriodMonthResponseStateEnumException SetEnforcePasswordPeriodMonthResponseStateEnum = "EXCEPTION"
    SetEnforcePasswordPeriodMonthResponseStateEnumForbidden SetEnforcePasswordPeriodMonthResponseStateEnum = "FORBIDDEN"
    SetEnforcePasswordPeriodMonthResponseStateEnumRemaining SetEnforcePasswordPeriodMonthResponseStateEnum = "REMAINING"
    SetEnforcePasswordPeriodMonthResponseStateEnumTimeout SetEnforcePasswordPeriodMonthResponseStateEnum = "TIMEOUT"
)

type SetEnforcePasswordPeriodMonthResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *SetEnforcePasswordPeriodMonthResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *bool `json:"body,omitempty"`
}

func (s SetEnforcePasswordPeriodMonthResponse) String() string {
	return utils.Beautify(s)
}

func (s SetEnforcePasswordPeriodMonthResponse) GoString() string {
	return s.String()
}

func (s SetEnforcePasswordPeriodMonthResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *SetEnforcePasswordPeriodMonthResponse) SetRequestId(v string) *SetEnforcePasswordPeriodMonthResponse {
	s.RequestId = &v
	return s
}

func (s *SetEnforcePasswordPeriodMonthResponse) SetErrorMessage(v string) *SetEnforcePasswordPeriodMonthResponse {
	s.ErrorMessage = &v
	return s
}

func (s *SetEnforcePasswordPeriodMonthResponse) SetErrorCode(v string) *SetEnforcePasswordPeriodMonthResponse {
	s.ErrorCode = &v
	return s
}

func (s *SetEnforcePasswordPeriodMonthResponse) SetState(v SetEnforcePasswordPeriodMonthResponseStateEnum) *SetEnforcePasswordPeriodMonthResponse {
	s.State = &v
	return s
}

func (s *SetEnforcePasswordPeriodMonthResponse) SetBody(v bool) *SetEnforcePasswordPeriodMonthResponse {
	s.Body = &v
	return s
}


type SetEnforcePasswordPeriodMonthResponseBuilder struct {
	s *SetEnforcePasswordPeriodMonthResponse
}

func NewSetEnforcePasswordPeriodMonthResponseBuilder() *SetEnforcePasswordPeriodMonthResponseBuilder {
	s := &SetEnforcePasswordPeriodMonthResponse{}
	b := &SetEnforcePasswordPeriodMonthResponseBuilder{s: s}
	return b
}

func (b *SetEnforcePasswordPeriodMonthResponseBuilder) RequestId(v string) *SetEnforcePasswordPeriodMonthResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *SetEnforcePasswordPeriodMonthResponseBuilder) ErrorMessage(v string) *SetEnforcePasswordPeriodMonthResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *SetEnforcePasswordPeriodMonthResponseBuilder) ErrorCode(v string) *SetEnforcePasswordPeriodMonthResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *SetEnforcePasswordPeriodMonthResponseBuilder) State(v SetEnforcePasswordPeriodMonthResponseStateEnum) *SetEnforcePasswordPeriodMonthResponseBuilder {
	b.s.State = &v
	return b
}

func (b *SetEnforcePasswordPeriodMonthResponseBuilder) Body(v bool) *SetEnforcePasswordPeriodMonthResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *SetEnforcePasswordPeriodMonthResponseBuilder) Build() *SetEnforcePasswordPeriodMonthResponse {
	return b.s
}


