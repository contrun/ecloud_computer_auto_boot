// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type BindInstanceListResponseStateEnum string

// List of State
const (
    BindInstanceListResponseStateEnumOk BindInstanceListResponseStateEnum = "OK"
    BindInstanceListResponseStateEnumError BindInstanceListResponseStateEnum = "ERROR"
    BindInstanceListResponseStateEnumException BindInstanceListResponseStateEnum = "EXCEPTION"
    BindInstanceListResponseStateEnumForbidden BindInstanceListResponseStateEnum = "FORBIDDEN"
    BindInstanceListResponseStateEnumRemaining BindInstanceListResponseStateEnum = "REMAINING"
    BindInstanceListResponseStateEnumTimeout BindInstanceListResponseStateEnum = "TIMEOUT"
)

type BindInstanceListResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *BindInstanceListResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *BindInstanceListResponseBody `json:"body,omitempty"`
}

func (s BindInstanceListResponse) String() string {
	return utils.Beautify(s)
}

func (s BindInstanceListResponse) GoString() string {
	return s.String()
}

func (s BindInstanceListResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BindInstanceListResponse) SetRequestId(v string) *BindInstanceListResponse {
	s.RequestId = &v
	return s
}

func (s *BindInstanceListResponse) SetErrorMessage(v string) *BindInstanceListResponse {
	s.ErrorMessage = &v
	return s
}

func (s *BindInstanceListResponse) SetErrorCode(v string) *BindInstanceListResponse {
	s.ErrorCode = &v
	return s
}

func (s *BindInstanceListResponse) SetState(v BindInstanceListResponseStateEnum) *BindInstanceListResponse {
	s.State = &v
	return s
}

func (s *BindInstanceListResponse) SetBody(v *BindInstanceListResponseBody) *BindInstanceListResponse {
	s.Body = v
	return s
}


type BindInstanceListResponseBuilder struct {
	s *BindInstanceListResponse
}

func NewBindInstanceListResponseBuilder() *BindInstanceListResponseBuilder {
	s := &BindInstanceListResponse{}
	b := &BindInstanceListResponseBuilder{s: s}
	return b
}

func (b *BindInstanceListResponseBuilder) RequestId(v string) *BindInstanceListResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *BindInstanceListResponseBuilder) ErrorMessage(v string) *BindInstanceListResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *BindInstanceListResponseBuilder) ErrorCode(v string) *BindInstanceListResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *BindInstanceListResponseBuilder) State(v BindInstanceListResponseStateEnum) *BindInstanceListResponseBuilder {
	b.s.State = &v
	return b
}

func (b *BindInstanceListResponseBuilder) Body(v *BindInstanceListResponseBody) *BindInstanceListResponseBuilder {
	b.s.Body = v
	return b
}

func (b *BindInstanceListResponseBuilder) Build() *BindInstanceListResponse {
	return b.s
}


