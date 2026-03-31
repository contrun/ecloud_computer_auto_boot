// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type RevokeResponseStateEnum string

// List of State
const (
    RevokeResponseStateEnumOk RevokeResponseStateEnum = "OK"
    RevokeResponseStateEnumError RevokeResponseStateEnum = "ERROR"
    RevokeResponseStateEnumException RevokeResponseStateEnum = "EXCEPTION"
    RevokeResponseStateEnumForbidden RevokeResponseStateEnum = "FORBIDDEN"
    RevokeResponseStateEnumRemaining RevokeResponseStateEnum = "REMAINING"
    RevokeResponseStateEnumTimeout RevokeResponseStateEnum = "TIMEOUT"
)

type RevokeResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *RevokeResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *string `json:"body,omitempty"`
}

func (s RevokeResponse) String() string {
	return utils.Beautify(s)
}

func (s RevokeResponse) GoString() string {
	return s.String()
}

func (s RevokeResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *RevokeResponse) SetRequestId(v string) *RevokeResponse {
	s.RequestId = &v
	return s
}

func (s *RevokeResponse) SetErrorMessage(v string) *RevokeResponse {
	s.ErrorMessage = &v
	return s
}

func (s *RevokeResponse) SetErrorCode(v string) *RevokeResponse {
	s.ErrorCode = &v
	return s
}

func (s *RevokeResponse) SetState(v RevokeResponseStateEnum) *RevokeResponse {
	s.State = &v
	return s
}

func (s *RevokeResponse) SetBody(v string) *RevokeResponse {
	s.Body = &v
	return s
}


type RevokeResponseBuilder struct {
	s *RevokeResponse
}

func NewRevokeResponseBuilder() *RevokeResponseBuilder {
	s := &RevokeResponse{}
	b := &RevokeResponseBuilder{s: s}
	return b
}

func (b *RevokeResponseBuilder) RequestId(v string) *RevokeResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *RevokeResponseBuilder) ErrorMessage(v string) *RevokeResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *RevokeResponseBuilder) ErrorCode(v string) *RevokeResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *RevokeResponseBuilder) State(v RevokeResponseStateEnum) *RevokeResponseBuilder {
	b.s.State = &v
	return b
}

func (b *RevokeResponseBuilder) Body(v string) *RevokeResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *RevokeResponseBuilder) Build() *RevokeResponse {
	return b.s
}


