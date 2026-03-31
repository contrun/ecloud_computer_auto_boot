// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type EditUserPolicyResponseStateEnum string

// List of State
const (
    EditUserPolicyResponseStateEnumOk EditUserPolicyResponseStateEnum = "OK"
    EditUserPolicyResponseStateEnumError EditUserPolicyResponseStateEnum = "ERROR"
    EditUserPolicyResponseStateEnumException EditUserPolicyResponseStateEnum = "EXCEPTION"
    EditUserPolicyResponseStateEnumForbidden EditUserPolicyResponseStateEnum = "FORBIDDEN"
    EditUserPolicyResponseStateEnumRemaining EditUserPolicyResponseStateEnum = "REMAINING"
    EditUserPolicyResponseStateEnumTimeout EditUserPolicyResponseStateEnum = "TIMEOUT"
)

type EditUserPolicyResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *EditUserPolicyResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *string `json:"body,omitempty"`
}

func (s EditUserPolicyResponse) String() string {
	return utils.Beautify(s)
}

func (s EditUserPolicyResponse) GoString() string {
	return s.String()
}

func (s EditUserPolicyResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *EditUserPolicyResponse) SetRequestId(v string) *EditUserPolicyResponse {
	s.RequestId = &v
	return s
}

func (s *EditUserPolicyResponse) SetErrorMessage(v string) *EditUserPolicyResponse {
	s.ErrorMessage = &v
	return s
}

func (s *EditUserPolicyResponse) SetErrorCode(v string) *EditUserPolicyResponse {
	s.ErrorCode = &v
	return s
}

func (s *EditUserPolicyResponse) SetState(v EditUserPolicyResponseStateEnum) *EditUserPolicyResponse {
	s.State = &v
	return s
}

func (s *EditUserPolicyResponse) SetBody(v string) *EditUserPolicyResponse {
	s.Body = &v
	return s
}


type EditUserPolicyResponseBuilder struct {
	s *EditUserPolicyResponse
}

func NewEditUserPolicyResponseBuilder() *EditUserPolicyResponseBuilder {
	s := &EditUserPolicyResponse{}
	b := &EditUserPolicyResponseBuilder{s: s}
	return b
}

func (b *EditUserPolicyResponseBuilder) RequestId(v string) *EditUserPolicyResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *EditUserPolicyResponseBuilder) ErrorMessage(v string) *EditUserPolicyResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *EditUserPolicyResponseBuilder) ErrorCode(v string) *EditUserPolicyResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *EditUserPolicyResponseBuilder) State(v EditUserPolicyResponseStateEnum) *EditUserPolicyResponseBuilder {
	b.s.State = &v
	return b
}

func (b *EditUserPolicyResponseBuilder) Body(v string) *EditUserPolicyResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *EditUserPolicyResponseBuilder) Build() *EditUserPolicyResponse {
	return b.s
}


