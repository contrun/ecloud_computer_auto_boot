// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type GetUserDetailByPolicyIdResponseStateEnum string

// List of State
const (
    GetUserDetailByPolicyIdResponseStateEnumOk GetUserDetailByPolicyIdResponseStateEnum = "OK"
    GetUserDetailByPolicyIdResponseStateEnumError GetUserDetailByPolicyIdResponseStateEnum = "ERROR"
    GetUserDetailByPolicyIdResponseStateEnumException GetUserDetailByPolicyIdResponseStateEnum = "EXCEPTION"
    GetUserDetailByPolicyIdResponseStateEnumForbidden GetUserDetailByPolicyIdResponseStateEnum = "FORBIDDEN"
    GetUserDetailByPolicyIdResponseStateEnumRemaining GetUserDetailByPolicyIdResponseStateEnum = "REMAINING"
    GetUserDetailByPolicyIdResponseStateEnumTimeout GetUserDetailByPolicyIdResponseStateEnum = "TIMEOUT"
)

type GetUserDetailByPolicyIdResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *GetUserDetailByPolicyIdResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *GetUserDetailByPolicyIdResponseBody `json:"body,omitempty"`
}

func (s GetUserDetailByPolicyIdResponse) String() string {
	return utils.Beautify(s)
}

func (s GetUserDetailByPolicyIdResponse) GoString() string {
	return s.String()
}

func (s GetUserDetailByPolicyIdResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetUserDetailByPolicyIdResponse) SetRequestId(v string) *GetUserDetailByPolicyIdResponse {
	s.RequestId = &v
	return s
}

func (s *GetUserDetailByPolicyIdResponse) SetErrorMessage(v string) *GetUserDetailByPolicyIdResponse {
	s.ErrorMessage = &v
	return s
}

func (s *GetUserDetailByPolicyIdResponse) SetErrorCode(v string) *GetUserDetailByPolicyIdResponse {
	s.ErrorCode = &v
	return s
}

func (s *GetUserDetailByPolicyIdResponse) SetState(v GetUserDetailByPolicyIdResponseStateEnum) *GetUserDetailByPolicyIdResponse {
	s.State = &v
	return s
}

func (s *GetUserDetailByPolicyIdResponse) SetBody(v *GetUserDetailByPolicyIdResponseBody) *GetUserDetailByPolicyIdResponse {
	s.Body = v
	return s
}


type GetUserDetailByPolicyIdResponseBuilder struct {
	s *GetUserDetailByPolicyIdResponse
}

func NewGetUserDetailByPolicyIdResponseBuilder() *GetUserDetailByPolicyIdResponseBuilder {
	s := &GetUserDetailByPolicyIdResponse{}
	b := &GetUserDetailByPolicyIdResponseBuilder{s: s}
	return b
}

func (b *GetUserDetailByPolicyIdResponseBuilder) RequestId(v string) *GetUserDetailByPolicyIdResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *GetUserDetailByPolicyIdResponseBuilder) ErrorMessage(v string) *GetUserDetailByPolicyIdResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *GetUserDetailByPolicyIdResponseBuilder) ErrorCode(v string) *GetUserDetailByPolicyIdResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *GetUserDetailByPolicyIdResponseBuilder) State(v GetUserDetailByPolicyIdResponseStateEnum) *GetUserDetailByPolicyIdResponseBuilder {
	b.s.State = &v
	return b
}

func (b *GetUserDetailByPolicyIdResponseBuilder) Body(v *GetUserDetailByPolicyIdResponseBody) *GetUserDetailByPolicyIdResponseBuilder {
	b.s.Body = v
	return b
}

func (b *GetUserDetailByPolicyIdResponseBuilder) Build() *GetUserDetailByPolicyIdResponse {
	return b.s
}


