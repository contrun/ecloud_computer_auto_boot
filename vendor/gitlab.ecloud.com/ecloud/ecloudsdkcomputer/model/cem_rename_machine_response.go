// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type CemRenameMachineResponseStateEnum string

// List of State
const (
    CemRenameMachineResponseStateEnumOk CemRenameMachineResponseStateEnum = "OK"
    CemRenameMachineResponseStateEnumError CemRenameMachineResponseStateEnum = "ERROR"
    CemRenameMachineResponseStateEnumException CemRenameMachineResponseStateEnum = "EXCEPTION"
    CemRenameMachineResponseStateEnumForbidden CemRenameMachineResponseStateEnum = "FORBIDDEN"
    CemRenameMachineResponseStateEnumRemaining CemRenameMachineResponseStateEnum = "REMAINING"
    CemRenameMachineResponseStateEnumTimeout CemRenameMachineResponseStateEnum = "TIMEOUT"
)

type CemRenameMachineResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *CemRenameMachineResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body interface{} `json:"body,omitempty"`
}

func (s CemRenameMachineResponse) String() string {
	return utils.Beautify(s)
}

func (s CemRenameMachineResponse) GoString() string {
	return s.String()
}

func (s CemRenameMachineResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CemRenameMachineResponse) SetRequestId(v string) *CemRenameMachineResponse {
	s.RequestId = &v
	return s
}

func (s *CemRenameMachineResponse) SetErrorMessage(v string) *CemRenameMachineResponse {
	s.ErrorMessage = &v
	return s
}

func (s *CemRenameMachineResponse) SetErrorCode(v string) *CemRenameMachineResponse {
	s.ErrorCode = &v
	return s
}

func (s *CemRenameMachineResponse) SetState(v CemRenameMachineResponseStateEnum) *CemRenameMachineResponse {
	s.State = &v
	return s
}

func (s *CemRenameMachineResponse) SetBody(v interface{}) *CemRenameMachineResponse {
	s.Body = v
	return s
}


type CemRenameMachineResponseBuilder struct {
	s *CemRenameMachineResponse
}

func NewCemRenameMachineResponseBuilder() *CemRenameMachineResponseBuilder {
	s := &CemRenameMachineResponse{}
	b := &CemRenameMachineResponseBuilder{s: s}
	return b
}

func (b *CemRenameMachineResponseBuilder) RequestId(v string) *CemRenameMachineResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *CemRenameMachineResponseBuilder) ErrorMessage(v string) *CemRenameMachineResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *CemRenameMachineResponseBuilder) ErrorCode(v string) *CemRenameMachineResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *CemRenameMachineResponseBuilder) State(v CemRenameMachineResponseStateEnum) *CemRenameMachineResponseBuilder {
	b.s.State = &v
	return b
}

func (b *CemRenameMachineResponseBuilder) Body(v interface{}) *CemRenameMachineResponseBuilder {
	b.s.Body = v
	return b
}

func (b *CemRenameMachineResponseBuilder) Build() *CemRenameMachineResponse {
	return b.s
}


