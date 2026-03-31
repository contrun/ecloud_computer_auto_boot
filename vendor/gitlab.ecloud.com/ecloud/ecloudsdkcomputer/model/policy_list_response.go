// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type PolicyListResponseStateEnum string

// List of State
const (
    PolicyListResponseStateEnumOk PolicyListResponseStateEnum = "OK"
    PolicyListResponseStateEnumError PolicyListResponseStateEnum = "ERROR"
    PolicyListResponseStateEnumException PolicyListResponseStateEnum = "EXCEPTION"
    PolicyListResponseStateEnumForbidden PolicyListResponseStateEnum = "FORBIDDEN"
    PolicyListResponseStateEnumRemaining PolicyListResponseStateEnum = "REMAINING"
    PolicyListResponseStateEnumTimeout PolicyListResponseStateEnum = "TIMEOUT"
)

type PolicyListResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *PolicyListResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *PolicyListResponseBody `json:"body,omitempty"`
}

func (s PolicyListResponse) String() string {
	return utils.Beautify(s)
}

func (s PolicyListResponse) GoString() string {
	return s.String()
}

func (s PolicyListResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *PolicyListResponse) SetRequestId(v string) *PolicyListResponse {
	s.RequestId = &v
	return s
}

func (s *PolicyListResponse) SetErrorMessage(v string) *PolicyListResponse {
	s.ErrorMessage = &v
	return s
}

func (s *PolicyListResponse) SetErrorCode(v string) *PolicyListResponse {
	s.ErrorCode = &v
	return s
}

func (s *PolicyListResponse) SetState(v PolicyListResponseStateEnum) *PolicyListResponse {
	s.State = &v
	return s
}

func (s *PolicyListResponse) SetBody(v *PolicyListResponseBody) *PolicyListResponse {
	s.Body = v
	return s
}


type PolicyListResponseBuilder struct {
	s *PolicyListResponse
}

func NewPolicyListResponseBuilder() *PolicyListResponseBuilder {
	s := &PolicyListResponse{}
	b := &PolicyListResponseBuilder{s: s}
	return b
}

func (b *PolicyListResponseBuilder) RequestId(v string) *PolicyListResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *PolicyListResponseBuilder) ErrorMessage(v string) *PolicyListResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *PolicyListResponseBuilder) ErrorCode(v string) *PolicyListResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *PolicyListResponseBuilder) State(v PolicyListResponseStateEnum) *PolicyListResponseBuilder {
	b.s.State = &v
	return b
}

func (b *PolicyListResponseBuilder) Body(v *PolicyListResponseBody) *PolicyListResponseBuilder {
	b.s.Body = v
	return b
}

func (b *PolicyListResponseBuilder) Build() *PolicyListResponse {
	return b.s
}


