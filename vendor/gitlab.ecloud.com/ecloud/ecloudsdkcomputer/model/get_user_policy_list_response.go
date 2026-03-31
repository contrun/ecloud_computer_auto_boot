// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type GetUserPolicyListResponseStateEnum string

// List of State
const (
    GetUserPolicyListResponseStateEnumOk GetUserPolicyListResponseStateEnum = "OK"
    GetUserPolicyListResponseStateEnumError GetUserPolicyListResponseStateEnum = "ERROR"
    GetUserPolicyListResponseStateEnumException GetUserPolicyListResponseStateEnum = "EXCEPTION"
    GetUserPolicyListResponseStateEnumForbidden GetUserPolicyListResponseStateEnum = "FORBIDDEN"
    GetUserPolicyListResponseStateEnumRemaining GetUserPolicyListResponseStateEnum = "REMAINING"
    GetUserPolicyListResponseStateEnumTimeout GetUserPolicyListResponseStateEnum = "TIMEOUT"
)

type GetUserPolicyListResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *GetUserPolicyListResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *GetUserPolicyListResponseBody `json:"body,omitempty"`
}

func (s GetUserPolicyListResponse) String() string {
	return utils.Beautify(s)
}

func (s GetUserPolicyListResponse) GoString() string {
	return s.String()
}

func (s GetUserPolicyListResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetUserPolicyListResponse) SetRequestId(v string) *GetUserPolicyListResponse {
	s.RequestId = &v
	return s
}

func (s *GetUserPolicyListResponse) SetErrorMessage(v string) *GetUserPolicyListResponse {
	s.ErrorMessage = &v
	return s
}

func (s *GetUserPolicyListResponse) SetErrorCode(v string) *GetUserPolicyListResponse {
	s.ErrorCode = &v
	return s
}

func (s *GetUserPolicyListResponse) SetState(v GetUserPolicyListResponseStateEnum) *GetUserPolicyListResponse {
	s.State = &v
	return s
}

func (s *GetUserPolicyListResponse) SetBody(v *GetUserPolicyListResponseBody) *GetUserPolicyListResponse {
	s.Body = v
	return s
}


type GetUserPolicyListResponseBuilder struct {
	s *GetUserPolicyListResponse
}

func NewGetUserPolicyListResponseBuilder() *GetUserPolicyListResponseBuilder {
	s := &GetUserPolicyListResponse{}
	b := &GetUserPolicyListResponseBuilder{s: s}
	return b
}

func (b *GetUserPolicyListResponseBuilder) RequestId(v string) *GetUserPolicyListResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *GetUserPolicyListResponseBuilder) ErrorMessage(v string) *GetUserPolicyListResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *GetUserPolicyListResponseBuilder) ErrorCode(v string) *GetUserPolicyListResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *GetUserPolicyListResponseBuilder) State(v GetUserPolicyListResponseStateEnum) *GetUserPolicyListResponseBuilder {
	b.s.State = &v
	return b
}

func (b *GetUserPolicyListResponseBuilder) Body(v *GetUserPolicyListResponseBody) *GetUserPolicyListResponseBuilder {
	b.s.Body = v
	return b
}

func (b *GetUserPolicyListResponseBuilder) Build() *GetUserPolicyListResponse {
	return b.s
}


