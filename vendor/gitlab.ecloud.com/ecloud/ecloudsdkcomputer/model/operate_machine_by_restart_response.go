// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type OperateMachineByRestartResponseStateEnum string

// List of State
const (
    OperateMachineByRestartResponseStateEnumOk OperateMachineByRestartResponseStateEnum = "OK"
    OperateMachineByRestartResponseStateEnumError OperateMachineByRestartResponseStateEnum = "ERROR"
    OperateMachineByRestartResponseStateEnumException OperateMachineByRestartResponseStateEnum = "EXCEPTION"
    OperateMachineByRestartResponseStateEnumForbidden OperateMachineByRestartResponseStateEnum = "FORBIDDEN"
    OperateMachineByRestartResponseStateEnumRemaining OperateMachineByRestartResponseStateEnum = "REMAINING"
    OperateMachineByRestartResponseStateEnumTimeout OperateMachineByRestartResponseStateEnum = "TIMEOUT"
)

type OperateMachineByRestartResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *OperateMachineByRestartResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body interface{} `json:"body,omitempty"`
}

func (s OperateMachineByRestartResponse) String() string {
	return utils.Beautify(s)
}

func (s OperateMachineByRestartResponse) GoString() string {
	return s.String()
}

func (s OperateMachineByRestartResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *OperateMachineByRestartResponse) SetRequestId(v string) *OperateMachineByRestartResponse {
	s.RequestId = &v
	return s
}

func (s *OperateMachineByRestartResponse) SetErrorMessage(v string) *OperateMachineByRestartResponse {
	s.ErrorMessage = &v
	return s
}

func (s *OperateMachineByRestartResponse) SetErrorCode(v string) *OperateMachineByRestartResponse {
	s.ErrorCode = &v
	return s
}

func (s *OperateMachineByRestartResponse) SetState(v OperateMachineByRestartResponseStateEnum) *OperateMachineByRestartResponse {
	s.State = &v
	return s
}

func (s *OperateMachineByRestartResponse) SetBody(v interface{}) *OperateMachineByRestartResponse {
	s.Body = v
	return s
}


type OperateMachineByRestartResponseBuilder struct {
	s *OperateMachineByRestartResponse
}

func NewOperateMachineByRestartResponseBuilder() *OperateMachineByRestartResponseBuilder {
	s := &OperateMachineByRestartResponse{}
	b := &OperateMachineByRestartResponseBuilder{s: s}
	return b
}

func (b *OperateMachineByRestartResponseBuilder) RequestId(v string) *OperateMachineByRestartResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *OperateMachineByRestartResponseBuilder) ErrorMessage(v string) *OperateMachineByRestartResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *OperateMachineByRestartResponseBuilder) ErrorCode(v string) *OperateMachineByRestartResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *OperateMachineByRestartResponseBuilder) State(v OperateMachineByRestartResponseStateEnum) *OperateMachineByRestartResponseBuilder {
	b.s.State = &v
	return b
}

func (b *OperateMachineByRestartResponseBuilder) Body(v interface{}) *OperateMachineByRestartResponseBuilder {
	b.s.Body = v
	return b
}

func (b *OperateMachineByRestartResponseBuilder) Build() *OperateMachineByRestartResponse {
	return b.s
}


