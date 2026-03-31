// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type UpdatePolicyResponseStateEnum string

// List of State
const (
    UpdatePolicyResponseStateEnumOk UpdatePolicyResponseStateEnum = "OK"
    UpdatePolicyResponseStateEnumError UpdatePolicyResponseStateEnum = "ERROR"
    UpdatePolicyResponseStateEnumException UpdatePolicyResponseStateEnum = "EXCEPTION"
    UpdatePolicyResponseStateEnumForbidden UpdatePolicyResponseStateEnum = "FORBIDDEN"
    UpdatePolicyResponseStateEnumRemaining UpdatePolicyResponseStateEnum = "REMAINING"
    UpdatePolicyResponseStateEnumTimeout UpdatePolicyResponseStateEnum = "TIMEOUT"
)

type UpdatePolicyResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *UpdatePolicyResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *bool `json:"body,omitempty"`
}

func (s UpdatePolicyResponse) String() string {
	return utils.Beautify(s)
}

func (s UpdatePolicyResponse) GoString() string {
	return s.String()
}

func (s UpdatePolicyResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdatePolicyResponse) SetRequestId(v string) *UpdatePolicyResponse {
	s.RequestId = &v
	return s
}

func (s *UpdatePolicyResponse) SetErrorMessage(v string) *UpdatePolicyResponse {
	s.ErrorMessage = &v
	return s
}

func (s *UpdatePolicyResponse) SetErrorCode(v string) *UpdatePolicyResponse {
	s.ErrorCode = &v
	return s
}

func (s *UpdatePolicyResponse) SetState(v UpdatePolicyResponseStateEnum) *UpdatePolicyResponse {
	s.State = &v
	return s
}

func (s *UpdatePolicyResponse) SetBody(v bool) *UpdatePolicyResponse {
	s.Body = &v
	return s
}


type UpdatePolicyResponseBuilder struct {
	s *UpdatePolicyResponse
}

func NewUpdatePolicyResponseBuilder() *UpdatePolicyResponseBuilder {
	s := &UpdatePolicyResponse{}
	b := &UpdatePolicyResponseBuilder{s: s}
	return b
}

func (b *UpdatePolicyResponseBuilder) RequestId(v string) *UpdatePolicyResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *UpdatePolicyResponseBuilder) ErrorMessage(v string) *UpdatePolicyResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *UpdatePolicyResponseBuilder) ErrorCode(v string) *UpdatePolicyResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *UpdatePolicyResponseBuilder) State(v UpdatePolicyResponseStateEnum) *UpdatePolicyResponseBuilder {
	b.s.State = &v
	return b
}

func (b *UpdatePolicyResponseBuilder) Body(v bool) *UpdatePolicyResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *UpdatePolicyResponseBuilder) Build() *UpdatePolicyResponse {
	return b.s
}


