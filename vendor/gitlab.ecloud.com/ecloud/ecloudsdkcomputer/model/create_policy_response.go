// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type CreatePolicyResponseStateEnum string

// List of State
const (
    CreatePolicyResponseStateEnumOk CreatePolicyResponseStateEnum = "OK"
    CreatePolicyResponseStateEnumError CreatePolicyResponseStateEnum = "ERROR"
    CreatePolicyResponseStateEnumException CreatePolicyResponseStateEnum = "EXCEPTION"
    CreatePolicyResponseStateEnumForbidden CreatePolicyResponseStateEnum = "FORBIDDEN"
    CreatePolicyResponseStateEnumRemaining CreatePolicyResponseStateEnum = "REMAINING"
    CreatePolicyResponseStateEnumTimeout CreatePolicyResponseStateEnum = "TIMEOUT"
)

type CreatePolicyResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *CreatePolicyResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *bool `json:"body,omitempty"`
}

func (s CreatePolicyResponse) String() string {
	return utils.Beautify(s)
}

func (s CreatePolicyResponse) GoString() string {
	return s.String()
}

func (s CreatePolicyResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreatePolicyResponse) SetRequestId(v string) *CreatePolicyResponse {
	s.RequestId = &v
	return s
}

func (s *CreatePolicyResponse) SetErrorMessage(v string) *CreatePolicyResponse {
	s.ErrorMessage = &v
	return s
}

func (s *CreatePolicyResponse) SetErrorCode(v string) *CreatePolicyResponse {
	s.ErrorCode = &v
	return s
}

func (s *CreatePolicyResponse) SetState(v CreatePolicyResponseStateEnum) *CreatePolicyResponse {
	s.State = &v
	return s
}

func (s *CreatePolicyResponse) SetBody(v bool) *CreatePolicyResponse {
	s.Body = &v
	return s
}


type CreatePolicyResponseBuilder struct {
	s *CreatePolicyResponse
}

func NewCreatePolicyResponseBuilder() *CreatePolicyResponseBuilder {
	s := &CreatePolicyResponse{}
	b := &CreatePolicyResponseBuilder{s: s}
	return b
}

func (b *CreatePolicyResponseBuilder) RequestId(v string) *CreatePolicyResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *CreatePolicyResponseBuilder) ErrorMessage(v string) *CreatePolicyResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *CreatePolicyResponseBuilder) ErrorCode(v string) *CreatePolicyResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *CreatePolicyResponseBuilder) State(v CreatePolicyResponseStateEnum) *CreatePolicyResponseBuilder {
	b.s.State = &v
	return b
}

func (b *CreatePolicyResponseBuilder) Body(v bool) *CreatePolicyResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *CreatePolicyResponseBuilder) Build() *CreatePolicyResponse {
	return b.s
}


