// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type AddResponseStateEnum string

// List of State
const (
    AddResponseStateEnumOk AddResponseStateEnum = "OK"
    AddResponseStateEnumError AddResponseStateEnum = "ERROR"
    AddResponseStateEnumException AddResponseStateEnum = "EXCEPTION"
    AddResponseStateEnumForbidden AddResponseStateEnum = "FORBIDDEN"
    AddResponseStateEnumRemaining AddResponseStateEnum = "REMAINING"
    AddResponseStateEnumTimeout AddResponseStateEnum = "TIMEOUT"
)

type AddResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *AddResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *string `json:"body,omitempty"`
}

func (s AddResponse) String() string {
	return utils.Beautify(s)
}

func (s AddResponse) GoString() string {
	return s.String()
}

func (s AddResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *AddResponse) SetRequestId(v string) *AddResponse {
	s.RequestId = &v
	return s
}

func (s *AddResponse) SetErrorMessage(v string) *AddResponse {
	s.ErrorMessage = &v
	return s
}

func (s *AddResponse) SetErrorCode(v string) *AddResponse {
	s.ErrorCode = &v
	return s
}

func (s *AddResponse) SetState(v AddResponseStateEnum) *AddResponse {
	s.State = &v
	return s
}

func (s *AddResponse) SetBody(v string) *AddResponse {
	s.Body = &v
	return s
}


type AddResponseBuilder struct {
	s *AddResponse
}

func NewAddResponseBuilder() *AddResponseBuilder {
	s := &AddResponse{}
	b := &AddResponseBuilder{s: s}
	return b
}

func (b *AddResponseBuilder) RequestId(v string) *AddResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *AddResponseBuilder) ErrorMessage(v string) *AddResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *AddResponseBuilder) ErrorCode(v string) *AddResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *AddResponseBuilder) State(v AddResponseStateEnum) *AddResponseBuilder {
	b.s.State = &v
	return b
}

func (b *AddResponseBuilder) Body(v string) *AddResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *AddResponseBuilder) Build() *AddResponse {
	return b.s
}


