// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type OperateMachineByReloadResponseStateEnum string

// List of State
const (
    OperateMachineByReloadResponseStateEnumOk OperateMachineByReloadResponseStateEnum = "OK"
    OperateMachineByReloadResponseStateEnumError OperateMachineByReloadResponseStateEnum = "ERROR"
    OperateMachineByReloadResponseStateEnumException OperateMachineByReloadResponseStateEnum = "EXCEPTION"
    OperateMachineByReloadResponseStateEnumForbidden OperateMachineByReloadResponseStateEnum = "FORBIDDEN"
    OperateMachineByReloadResponseStateEnumRemaining OperateMachineByReloadResponseStateEnum = "REMAINING"
    OperateMachineByReloadResponseStateEnumTimeout OperateMachineByReloadResponseStateEnum = "TIMEOUT"
)

type OperateMachineByReloadResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *OperateMachineByReloadResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body interface{} `json:"body,omitempty"`
}

func (s OperateMachineByReloadResponse) String() string {
	return utils.Beautify(s)
}

func (s OperateMachineByReloadResponse) GoString() string {
	return s.String()
}

func (s OperateMachineByReloadResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *OperateMachineByReloadResponse) SetRequestId(v string) *OperateMachineByReloadResponse {
	s.RequestId = &v
	return s
}

func (s *OperateMachineByReloadResponse) SetErrorMessage(v string) *OperateMachineByReloadResponse {
	s.ErrorMessage = &v
	return s
}

func (s *OperateMachineByReloadResponse) SetErrorCode(v string) *OperateMachineByReloadResponse {
	s.ErrorCode = &v
	return s
}

func (s *OperateMachineByReloadResponse) SetState(v OperateMachineByReloadResponseStateEnum) *OperateMachineByReloadResponse {
	s.State = &v
	return s
}

func (s *OperateMachineByReloadResponse) SetBody(v interface{}) *OperateMachineByReloadResponse {
	s.Body = v
	return s
}


type OperateMachineByReloadResponseBuilder struct {
	s *OperateMachineByReloadResponse
}

func NewOperateMachineByReloadResponseBuilder() *OperateMachineByReloadResponseBuilder {
	s := &OperateMachineByReloadResponse{}
	b := &OperateMachineByReloadResponseBuilder{s: s}
	return b
}

func (b *OperateMachineByReloadResponseBuilder) RequestId(v string) *OperateMachineByReloadResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *OperateMachineByReloadResponseBuilder) ErrorMessage(v string) *OperateMachineByReloadResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *OperateMachineByReloadResponseBuilder) ErrorCode(v string) *OperateMachineByReloadResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *OperateMachineByReloadResponseBuilder) State(v OperateMachineByReloadResponseStateEnum) *OperateMachineByReloadResponseBuilder {
	b.s.State = &v
	return b
}

func (b *OperateMachineByReloadResponseBuilder) Body(v interface{}) *OperateMachineByReloadResponseBuilder {
	b.s.Body = v
	return b
}

func (b *OperateMachineByReloadResponseBuilder) Build() *OperateMachineByReloadResponse {
	return b.s
}


