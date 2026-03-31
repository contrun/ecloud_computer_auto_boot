// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type UnbindUserPolicyResponseStateEnum string

// List of State
const (
    UnbindUserPolicyResponseStateEnumOk UnbindUserPolicyResponseStateEnum = "OK"
    UnbindUserPolicyResponseStateEnumError UnbindUserPolicyResponseStateEnum = "ERROR"
    UnbindUserPolicyResponseStateEnumException UnbindUserPolicyResponseStateEnum = "EXCEPTION"
    UnbindUserPolicyResponseStateEnumForbidden UnbindUserPolicyResponseStateEnum = "FORBIDDEN"
    UnbindUserPolicyResponseStateEnumRemaining UnbindUserPolicyResponseStateEnum = "REMAINING"
    UnbindUserPolicyResponseStateEnumTimeout UnbindUserPolicyResponseStateEnum = "TIMEOUT"
)

type UnbindUserPolicyResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *UnbindUserPolicyResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *string `json:"body,omitempty"`
}

func (s UnbindUserPolicyResponse) String() string {
	return utils.Beautify(s)
}

func (s UnbindUserPolicyResponse) GoString() string {
	return s.String()
}

func (s UnbindUserPolicyResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UnbindUserPolicyResponse) SetRequestId(v string) *UnbindUserPolicyResponse {
	s.RequestId = &v
	return s
}

func (s *UnbindUserPolicyResponse) SetErrorMessage(v string) *UnbindUserPolicyResponse {
	s.ErrorMessage = &v
	return s
}

func (s *UnbindUserPolicyResponse) SetErrorCode(v string) *UnbindUserPolicyResponse {
	s.ErrorCode = &v
	return s
}

func (s *UnbindUserPolicyResponse) SetState(v UnbindUserPolicyResponseStateEnum) *UnbindUserPolicyResponse {
	s.State = &v
	return s
}

func (s *UnbindUserPolicyResponse) SetBody(v string) *UnbindUserPolicyResponse {
	s.Body = &v
	return s
}


type UnbindUserPolicyResponseBuilder struct {
	s *UnbindUserPolicyResponse
}

func NewUnbindUserPolicyResponseBuilder() *UnbindUserPolicyResponseBuilder {
	s := &UnbindUserPolicyResponse{}
	b := &UnbindUserPolicyResponseBuilder{s: s}
	return b
}

func (b *UnbindUserPolicyResponseBuilder) RequestId(v string) *UnbindUserPolicyResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *UnbindUserPolicyResponseBuilder) ErrorMessage(v string) *UnbindUserPolicyResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *UnbindUserPolicyResponseBuilder) ErrorCode(v string) *UnbindUserPolicyResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *UnbindUserPolicyResponseBuilder) State(v UnbindUserPolicyResponseStateEnum) *UnbindUserPolicyResponseBuilder {
	b.s.State = &v
	return b
}

func (b *UnbindUserPolicyResponseBuilder) Body(v string) *UnbindUserPolicyResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *UnbindUserPolicyResponseBuilder) Build() *UnbindUserPolicyResponse {
	return b.s
}


