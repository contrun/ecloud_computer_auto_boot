// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type CopyPolicyResponseStateEnum string

// List of State
const (
    CopyPolicyResponseStateEnumOk CopyPolicyResponseStateEnum = "OK"
    CopyPolicyResponseStateEnumError CopyPolicyResponseStateEnum = "ERROR"
    CopyPolicyResponseStateEnumException CopyPolicyResponseStateEnum = "EXCEPTION"
    CopyPolicyResponseStateEnumForbidden CopyPolicyResponseStateEnum = "FORBIDDEN"
    CopyPolicyResponseStateEnumRemaining CopyPolicyResponseStateEnum = "REMAINING"
    CopyPolicyResponseStateEnumTimeout CopyPolicyResponseStateEnum = "TIMEOUT"
)

type CopyPolicyResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *CopyPolicyResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *bool `json:"body,omitempty"`
}

func (s CopyPolicyResponse) String() string {
	return utils.Beautify(s)
}

func (s CopyPolicyResponse) GoString() string {
	return s.String()
}

func (s CopyPolicyResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CopyPolicyResponse) SetRequestId(v string) *CopyPolicyResponse {
	s.RequestId = &v
	return s
}

func (s *CopyPolicyResponse) SetErrorMessage(v string) *CopyPolicyResponse {
	s.ErrorMessage = &v
	return s
}

func (s *CopyPolicyResponse) SetErrorCode(v string) *CopyPolicyResponse {
	s.ErrorCode = &v
	return s
}

func (s *CopyPolicyResponse) SetState(v CopyPolicyResponseStateEnum) *CopyPolicyResponse {
	s.State = &v
	return s
}

func (s *CopyPolicyResponse) SetBody(v bool) *CopyPolicyResponse {
	s.Body = &v
	return s
}


type CopyPolicyResponseBuilder struct {
	s *CopyPolicyResponse
}

func NewCopyPolicyResponseBuilder() *CopyPolicyResponseBuilder {
	s := &CopyPolicyResponse{}
	b := &CopyPolicyResponseBuilder{s: s}
	return b
}

func (b *CopyPolicyResponseBuilder) RequestId(v string) *CopyPolicyResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *CopyPolicyResponseBuilder) ErrorMessage(v string) *CopyPolicyResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *CopyPolicyResponseBuilder) ErrorCode(v string) *CopyPolicyResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *CopyPolicyResponseBuilder) State(v CopyPolicyResponseStateEnum) *CopyPolicyResponseBuilder {
	b.s.State = &v
	return b
}

func (b *CopyPolicyResponseBuilder) Body(v bool) *CopyPolicyResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *CopyPolicyResponseBuilder) Build() *CopyPolicyResponse {
	return b.s
}


