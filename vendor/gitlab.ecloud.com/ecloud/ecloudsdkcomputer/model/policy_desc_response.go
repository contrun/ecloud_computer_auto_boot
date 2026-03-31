// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type PolicyDescResponseStateEnum string

// List of State
const (
    PolicyDescResponseStateEnumOk PolicyDescResponseStateEnum = "OK"
    PolicyDescResponseStateEnumError PolicyDescResponseStateEnum = "ERROR"
    PolicyDescResponseStateEnumException PolicyDescResponseStateEnum = "EXCEPTION"
    PolicyDescResponseStateEnumForbidden PolicyDescResponseStateEnum = "FORBIDDEN"
    PolicyDescResponseStateEnumRemaining PolicyDescResponseStateEnum = "REMAINING"
    PolicyDescResponseStateEnumTimeout PolicyDescResponseStateEnum = "TIMEOUT"
)

type PolicyDescResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *PolicyDescResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *PolicyDescResponseBody `json:"body,omitempty"`
}

func (s PolicyDescResponse) String() string {
	return utils.Beautify(s)
}

func (s PolicyDescResponse) GoString() string {
	return s.String()
}

func (s PolicyDescResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *PolicyDescResponse) SetRequestId(v string) *PolicyDescResponse {
	s.RequestId = &v
	return s
}

func (s *PolicyDescResponse) SetErrorMessage(v string) *PolicyDescResponse {
	s.ErrorMessage = &v
	return s
}

func (s *PolicyDescResponse) SetErrorCode(v string) *PolicyDescResponse {
	s.ErrorCode = &v
	return s
}

func (s *PolicyDescResponse) SetState(v PolicyDescResponseStateEnum) *PolicyDescResponse {
	s.State = &v
	return s
}

func (s *PolicyDescResponse) SetBody(v *PolicyDescResponseBody) *PolicyDescResponse {
	s.Body = v
	return s
}


type PolicyDescResponseBuilder struct {
	s *PolicyDescResponse
}

func NewPolicyDescResponseBuilder() *PolicyDescResponseBuilder {
	s := &PolicyDescResponse{}
	b := &PolicyDescResponseBuilder{s: s}
	return b
}

func (b *PolicyDescResponseBuilder) RequestId(v string) *PolicyDescResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *PolicyDescResponseBuilder) ErrorMessage(v string) *PolicyDescResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *PolicyDescResponseBuilder) ErrorCode(v string) *PolicyDescResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *PolicyDescResponseBuilder) State(v PolicyDescResponseStateEnum) *PolicyDescResponseBuilder {
	b.s.State = &v
	return b
}

func (b *PolicyDescResponseBuilder) Body(v *PolicyDescResponseBody) *PolicyDescResponseBuilder {
	b.s.Body = v
	return b
}

func (b *PolicyDescResponseBuilder) Build() *PolicyDescResponse {
	return b.s
}


