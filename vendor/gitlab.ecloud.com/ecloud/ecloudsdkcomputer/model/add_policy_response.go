// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type AddPolicyResponseStateEnum string

// List of State
const (
    AddPolicyResponseStateEnumOk AddPolicyResponseStateEnum = "OK"
    AddPolicyResponseStateEnumError AddPolicyResponseStateEnum = "ERROR"
    AddPolicyResponseStateEnumException AddPolicyResponseStateEnum = "EXCEPTION"
    AddPolicyResponseStateEnumForbidden AddPolicyResponseStateEnum = "FORBIDDEN"
    AddPolicyResponseStateEnumRemaining AddPolicyResponseStateEnum = "REMAINING"
    AddPolicyResponseStateEnumTimeout AddPolicyResponseStateEnum = "TIMEOUT"
)

type AddPolicyResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *AddPolicyResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *bool `json:"body,omitempty"`
}

func (s AddPolicyResponse) String() string {
	return utils.Beautify(s)
}

func (s AddPolicyResponse) GoString() string {
	return s.String()
}

func (s AddPolicyResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *AddPolicyResponse) SetRequestId(v string) *AddPolicyResponse {
	s.RequestId = &v
	return s
}

func (s *AddPolicyResponse) SetErrorMessage(v string) *AddPolicyResponse {
	s.ErrorMessage = &v
	return s
}

func (s *AddPolicyResponse) SetErrorCode(v string) *AddPolicyResponse {
	s.ErrorCode = &v
	return s
}

func (s *AddPolicyResponse) SetState(v AddPolicyResponseStateEnum) *AddPolicyResponse {
	s.State = &v
	return s
}

func (s *AddPolicyResponse) SetBody(v bool) *AddPolicyResponse {
	s.Body = &v
	return s
}


type AddPolicyResponseBuilder struct {
	s *AddPolicyResponse
}

func NewAddPolicyResponseBuilder() *AddPolicyResponseBuilder {
	s := &AddPolicyResponse{}
	b := &AddPolicyResponseBuilder{s: s}
	return b
}

func (b *AddPolicyResponseBuilder) RequestId(v string) *AddPolicyResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *AddPolicyResponseBuilder) ErrorMessage(v string) *AddPolicyResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *AddPolicyResponseBuilder) ErrorCode(v string) *AddPolicyResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *AddPolicyResponseBuilder) State(v AddPolicyResponseStateEnum) *AddPolicyResponseBuilder {
	b.s.State = &v
	return b
}

func (b *AddPolicyResponseBuilder) Body(v bool) *AddPolicyResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *AddPolicyResponseBuilder) Build() *AddPolicyResponse {
	return b.s
}


