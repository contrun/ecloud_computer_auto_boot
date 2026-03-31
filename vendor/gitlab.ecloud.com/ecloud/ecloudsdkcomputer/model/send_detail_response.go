// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type SendDetailResponseStateEnum string

// List of State
const (
    SendDetailResponseStateEnumOk SendDetailResponseStateEnum = "OK"
    SendDetailResponseStateEnumError SendDetailResponseStateEnum = "ERROR"
    SendDetailResponseStateEnumException SendDetailResponseStateEnum = "EXCEPTION"
    SendDetailResponseStateEnumForbidden SendDetailResponseStateEnum = "FORBIDDEN"
    SendDetailResponseStateEnumRemaining SendDetailResponseStateEnum = "REMAINING"
    SendDetailResponseStateEnumTimeout SendDetailResponseStateEnum = "TIMEOUT"
)

type SendDetailResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *SendDetailResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *SendDetailResponseBody `json:"body,omitempty"`
}

func (s SendDetailResponse) String() string {
	return utils.Beautify(s)
}

func (s SendDetailResponse) GoString() string {
	return s.String()
}

func (s SendDetailResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *SendDetailResponse) SetRequestId(v string) *SendDetailResponse {
	s.RequestId = &v
	return s
}

func (s *SendDetailResponse) SetErrorMessage(v string) *SendDetailResponse {
	s.ErrorMessage = &v
	return s
}

func (s *SendDetailResponse) SetErrorCode(v string) *SendDetailResponse {
	s.ErrorCode = &v
	return s
}

func (s *SendDetailResponse) SetState(v SendDetailResponseStateEnum) *SendDetailResponse {
	s.State = &v
	return s
}

func (s *SendDetailResponse) SetBody(v *SendDetailResponseBody) *SendDetailResponse {
	s.Body = v
	return s
}


type SendDetailResponseBuilder struct {
	s *SendDetailResponse
}

func NewSendDetailResponseBuilder() *SendDetailResponseBuilder {
	s := &SendDetailResponse{}
	b := &SendDetailResponseBuilder{s: s}
	return b
}

func (b *SendDetailResponseBuilder) RequestId(v string) *SendDetailResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *SendDetailResponseBuilder) ErrorMessage(v string) *SendDetailResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *SendDetailResponseBuilder) ErrorCode(v string) *SendDetailResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *SendDetailResponseBuilder) State(v SendDetailResponseStateEnum) *SendDetailResponseBuilder {
	b.s.State = &v
	return b
}

func (b *SendDetailResponseBuilder) Body(v *SendDetailResponseBody) *SendDetailResponseBuilder {
	b.s.Body = v
	return b
}

func (b *SendDetailResponseBuilder) Build() *SendDetailResponse {
	return b.s
}


