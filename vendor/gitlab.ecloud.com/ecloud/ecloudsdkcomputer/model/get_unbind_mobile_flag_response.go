// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type GetUnbindMobileFlagResponseStateEnum string

// List of State
const (
    GetUnbindMobileFlagResponseStateEnumOk GetUnbindMobileFlagResponseStateEnum = "OK"
    GetUnbindMobileFlagResponseStateEnumError GetUnbindMobileFlagResponseStateEnum = "ERROR"
    GetUnbindMobileFlagResponseStateEnumException GetUnbindMobileFlagResponseStateEnum = "EXCEPTION"
    GetUnbindMobileFlagResponseStateEnumForbidden GetUnbindMobileFlagResponseStateEnum = "FORBIDDEN"
    GetUnbindMobileFlagResponseStateEnumRemaining GetUnbindMobileFlagResponseStateEnum = "REMAINING"
    GetUnbindMobileFlagResponseStateEnumTimeout GetUnbindMobileFlagResponseStateEnum = "TIMEOUT"
)

type GetUnbindMobileFlagResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *GetUnbindMobileFlagResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *GetUnbindMobileFlagResponseBody `json:"body,omitempty"`
}

func (s GetUnbindMobileFlagResponse) String() string {
	return utils.Beautify(s)
}

func (s GetUnbindMobileFlagResponse) GoString() string {
	return s.String()
}

func (s GetUnbindMobileFlagResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetUnbindMobileFlagResponse) SetRequestId(v string) *GetUnbindMobileFlagResponse {
	s.RequestId = &v
	return s
}

func (s *GetUnbindMobileFlagResponse) SetErrorMessage(v string) *GetUnbindMobileFlagResponse {
	s.ErrorMessage = &v
	return s
}

func (s *GetUnbindMobileFlagResponse) SetErrorCode(v string) *GetUnbindMobileFlagResponse {
	s.ErrorCode = &v
	return s
}

func (s *GetUnbindMobileFlagResponse) SetState(v GetUnbindMobileFlagResponseStateEnum) *GetUnbindMobileFlagResponse {
	s.State = &v
	return s
}

func (s *GetUnbindMobileFlagResponse) SetBody(v *GetUnbindMobileFlagResponseBody) *GetUnbindMobileFlagResponse {
	s.Body = v
	return s
}


type GetUnbindMobileFlagResponseBuilder struct {
	s *GetUnbindMobileFlagResponse
}

func NewGetUnbindMobileFlagResponseBuilder() *GetUnbindMobileFlagResponseBuilder {
	s := &GetUnbindMobileFlagResponse{}
	b := &GetUnbindMobileFlagResponseBuilder{s: s}
	return b
}

func (b *GetUnbindMobileFlagResponseBuilder) RequestId(v string) *GetUnbindMobileFlagResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *GetUnbindMobileFlagResponseBuilder) ErrorMessage(v string) *GetUnbindMobileFlagResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *GetUnbindMobileFlagResponseBuilder) ErrorCode(v string) *GetUnbindMobileFlagResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *GetUnbindMobileFlagResponseBuilder) State(v GetUnbindMobileFlagResponseStateEnum) *GetUnbindMobileFlagResponseBuilder {
	b.s.State = &v
	return b
}

func (b *GetUnbindMobileFlagResponseBuilder) Body(v *GetUnbindMobileFlagResponseBody) *GetUnbindMobileFlagResponseBuilder {
	b.s.Body = v
	return b
}

func (b *GetUnbindMobileFlagResponseBuilder) Build() *GetUnbindMobileFlagResponse {
	return b.s
}


