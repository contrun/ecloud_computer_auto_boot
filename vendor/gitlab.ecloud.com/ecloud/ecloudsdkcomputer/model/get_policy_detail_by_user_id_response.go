// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type GetPolicyDetailByUserIdResponseStateEnum string

// List of State
const (
    GetPolicyDetailByUserIdResponseStateEnumOk GetPolicyDetailByUserIdResponseStateEnum = "OK"
    GetPolicyDetailByUserIdResponseStateEnumError GetPolicyDetailByUserIdResponseStateEnum = "ERROR"
    GetPolicyDetailByUserIdResponseStateEnumException GetPolicyDetailByUserIdResponseStateEnum = "EXCEPTION"
    GetPolicyDetailByUserIdResponseStateEnumForbidden GetPolicyDetailByUserIdResponseStateEnum = "FORBIDDEN"
    GetPolicyDetailByUserIdResponseStateEnumRemaining GetPolicyDetailByUserIdResponseStateEnum = "REMAINING"
    GetPolicyDetailByUserIdResponseStateEnumTimeout GetPolicyDetailByUserIdResponseStateEnum = "TIMEOUT"
)

type GetPolicyDetailByUserIdResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *GetPolicyDetailByUserIdResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *GetPolicyDetailByUserIdResponseBody `json:"body,omitempty"`
}

func (s GetPolicyDetailByUserIdResponse) String() string {
	return utils.Beautify(s)
}

func (s GetPolicyDetailByUserIdResponse) GoString() string {
	return s.String()
}

func (s GetPolicyDetailByUserIdResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetPolicyDetailByUserIdResponse) SetRequestId(v string) *GetPolicyDetailByUserIdResponse {
	s.RequestId = &v
	return s
}

func (s *GetPolicyDetailByUserIdResponse) SetErrorMessage(v string) *GetPolicyDetailByUserIdResponse {
	s.ErrorMessage = &v
	return s
}

func (s *GetPolicyDetailByUserIdResponse) SetErrorCode(v string) *GetPolicyDetailByUserIdResponse {
	s.ErrorCode = &v
	return s
}

func (s *GetPolicyDetailByUserIdResponse) SetState(v GetPolicyDetailByUserIdResponseStateEnum) *GetPolicyDetailByUserIdResponse {
	s.State = &v
	return s
}

func (s *GetPolicyDetailByUserIdResponse) SetBody(v *GetPolicyDetailByUserIdResponseBody) *GetPolicyDetailByUserIdResponse {
	s.Body = v
	return s
}


type GetPolicyDetailByUserIdResponseBuilder struct {
	s *GetPolicyDetailByUserIdResponse
}

func NewGetPolicyDetailByUserIdResponseBuilder() *GetPolicyDetailByUserIdResponseBuilder {
	s := &GetPolicyDetailByUserIdResponse{}
	b := &GetPolicyDetailByUserIdResponseBuilder{s: s}
	return b
}

func (b *GetPolicyDetailByUserIdResponseBuilder) RequestId(v string) *GetPolicyDetailByUserIdResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *GetPolicyDetailByUserIdResponseBuilder) ErrorMessage(v string) *GetPolicyDetailByUserIdResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *GetPolicyDetailByUserIdResponseBuilder) ErrorCode(v string) *GetPolicyDetailByUserIdResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *GetPolicyDetailByUserIdResponseBuilder) State(v GetPolicyDetailByUserIdResponseStateEnum) *GetPolicyDetailByUserIdResponseBuilder {
	b.s.State = &v
	return b
}

func (b *GetPolicyDetailByUserIdResponseBuilder) Body(v *GetPolicyDetailByUserIdResponseBody) *GetPolicyDetailByUserIdResponseBuilder {
	b.s.Body = v
	return b
}

func (b *GetPolicyDetailByUserIdResponseBuilder) Build() *GetPolicyDetailByUserIdResponse {
	return b.s
}


