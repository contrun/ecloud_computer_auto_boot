// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type UpdateComputerPolicyResponseStateEnum string

// List of State
const (
    UpdateComputerPolicyResponseStateEnumOk UpdateComputerPolicyResponseStateEnum = "OK"
    UpdateComputerPolicyResponseStateEnumError UpdateComputerPolicyResponseStateEnum = "ERROR"
    UpdateComputerPolicyResponseStateEnumException UpdateComputerPolicyResponseStateEnum = "EXCEPTION"
    UpdateComputerPolicyResponseStateEnumForbidden UpdateComputerPolicyResponseStateEnum = "FORBIDDEN"
    UpdateComputerPolicyResponseStateEnumRemaining UpdateComputerPolicyResponseStateEnum = "REMAINING"
    UpdateComputerPolicyResponseStateEnumTimeout UpdateComputerPolicyResponseStateEnum = "TIMEOUT"
)

type UpdateComputerPolicyResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *UpdateComputerPolicyResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *bool `json:"body,omitempty"`
}

func (s UpdateComputerPolicyResponse) String() string {
	return utils.Beautify(s)
}

func (s UpdateComputerPolicyResponse) GoString() string {
	return s.String()
}

func (s UpdateComputerPolicyResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateComputerPolicyResponse) SetRequestId(v string) *UpdateComputerPolicyResponse {
	s.RequestId = &v
	return s
}

func (s *UpdateComputerPolicyResponse) SetErrorMessage(v string) *UpdateComputerPolicyResponse {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateComputerPolicyResponse) SetErrorCode(v string) *UpdateComputerPolicyResponse {
	s.ErrorCode = &v
	return s
}

func (s *UpdateComputerPolicyResponse) SetState(v UpdateComputerPolicyResponseStateEnum) *UpdateComputerPolicyResponse {
	s.State = &v
	return s
}

func (s *UpdateComputerPolicyResponse) SetBody(v bool) *UpdateComputerPolicyResponse {
	s.Body = &v
	return s
}


type UpdateComputerPolicyResponseBuilder struct {
	s *UpdateComputerPolicyResponse
}

func NewUpdateComputerPolicyResponseBuilder() *UpdateComputerPolicyResponseBuilder {
	s := &UpdateComputerPolicyResponse{}
	b := &UpdateComputerPolicyResponseBuilder{s: s}
	return b
}

func (b *UpdateComputerPolicyResponseBuilder) RequestId(v string) *UpdateComputerPolicyResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *UpdateComputerPolicyResponseBuilder) ErrorMessage(v string) *UpdateComputerPolicyResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *UpdateComputerPolicyResponseBuilder) ErrorCode(v string) *UpdateComputerPolicyResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *UpdateComputerPolicyResponseBuilder) State(v UpdateComputerPolicyResponseStateEnum) *UpdateComputerPolicyResponseBuilder {
	b.s.State = &v
	return b
}

func (b *UpdateComputerPolicyResponseBuilder) Body(v bool) *UpdateComputerPolicyResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *UpdateComputerPolicyResponseBuilder) Build() *UpdateComputerPolicyResponse {
	return b.s
}


