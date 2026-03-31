// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type GetPolicyInfoByIdResponseStateEnum string

// List of State
const (
    GetPolicyInfoByIdResponseStateEnumOk GetPolicyInfoByIdResponseStateEnum = "OK"
    GetPolicyInfoByIdResponseStateEnumError GetPolicyInfoByIdResponseStateEnum = "ERROR"
    GetPolicyInfoByIdResponseStateEnumException GetPolicyInfoByIdResponseStateEnum = "EXCEPTION"
    GetPolicyInfoByIdResponseStateEnumForbidden GetPolicyInfoByIdResponseStateEnum = "FORBIDDEN"
    GetPolicyInfoByIdResponseStateEnumRemaining GetPolicyInfoByIdResponseStateEnum = "REMAINING"
    GetPolicyInfoByIdResponseStateEnumTimeout GetPolicyInfoByIdResponseStateEnum = "TIMEOUT"
)

type GetPolicyInfoByIdResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *GetPolicyInfoByIdResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *GetPolicyInfoByIdResponseBody `json:"body,omitempty"`
}

func (s GetPolicyInfoByIdResponse) String() string {
	return utils.Beautify(s)
}

func (s GetPolicyInfoByIdResponse) GoString() string {
	return s.String()
}

func (s GetPolicyInfoByIdResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetPolicyInfoByIdResponse) SetRequestId(v string) *GetPolicyInfoByIdResponse {
	s.RequestId = &v
	return s
}

func (s *GetPolicyInfoByIdResponse) SetErrorMessage(v string) *GetPolicyInfoByIdResponse {
	s.ErrorMessage = &v
	return s
}

func (s *GetPolicyInfoByIdResponse) SetErrorCode(v string) *GetPolicyInfoByIdResponse {
	s.ErrorCode = &v
	return s
}

func (s *GetPolicyInfoByIdResponse) SetState(v GetPolicyInfoByIdResponseStateEnum) *GetPolicyInfoByIdResponse {
	s.State = &v
	return s
}

func (s *GetPolicyInfoByIdResponse) SetBody(v *GetPolicyInfoByIdResponseBody) *GetPolicyInfoByIdResponse {
	s.Body = v
	return s
}


type GetPolicyInfoByIdResponseBuilder struct {
	s *GetPolicyInfoByIdResponse
}

func NewGetPolicyInfoByIdResponseBuilder() *GetPolicyInfoByIdResponseBuilder {
	s := &GetPolicyInfoByIdResponse{}
	b := &GetPolicyInfoByIdResponseBuilder{s: s}
	return b
}

func (b *GetPolicyInfoByIdResponseBuilder) RequestId(v string) *GetPolicyInfoByIdResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *GetPolicyInfoByIdResponseBuilder) ErrorMessage(v string) *GetPolicyInfoByIdResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *GetPolicyInfoByIdResponseBuilder) ErrorCode(v string) *GetPolicyInfoByIdResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *GetPolicyInfoByIdResponseBuilder) State(v GetPolicyInfoByIdResponseStateEnum) *GetPolicyInfoByIdResponseBuilder {
	b.s.State = &v
	return b
}

func (b *GetPolicyInfoByIdResponseBuilder) Body(v *GetPolicyInfoByIdResponseBody) *GetPolicyInfoByIdResponseBuilder {
	b.s.Body = v
	return b
}

func (b *GetPolicyInfoByIdResponseBuilder) Build() *GetPolicyInfoByIdResponse {
	return b.s
}


