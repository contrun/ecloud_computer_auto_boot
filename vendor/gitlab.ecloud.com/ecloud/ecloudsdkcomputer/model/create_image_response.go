// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type CreateImageResponseStateEnum string

// List of State
const (
    CreateImageResponseStateEnumOk CreateImageResponseStateEnum = "OK"
    CreateImageResponseStateEnumError CreateImageResponseStateEnum = "ERROR"
    CreateImageResponseStateEnumException CreateImageResponseStateEnum = "EXCEPTION"
    CreateImageResponseStateEnumForbidden CreateImageResponseStateEnum = "FORBIDDEN"
    CreateImageResponseStateEnumRemaining CreateImageResponseStateEnum = "REMAINING"
    CreateImageResponseStateEnumTimeout CreateImageResponseStateEnum = "TIMEOUT"
)

type CreateImageResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *CreateImageResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *CreateImageResponseBody `json:"body,omitempty"`
}

func (s CreateImageResponse) String() string {
	return utils.Beautify(s)
}

func (s CreateImageResponse) GoString() string {
	return s.String()
}

func (s CreateImageResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateImageResponse) SetRequestId(v string) *CreateImageResponse {
	s.RequestId = &v
	return s
}

func (s *CreateImageResponse) SetErrorMessage(v string) *CreateImageResponse {
	s.ErrorMessage = &v
	return s
}

func (s *CreateImageResponse) SetErrorCode(v string) *CreateImageResponse {
	s.ErrorCode = &v
	return s
}

func (s *CreateImageResponse) SetState(v CreateImageResponseStateEnum) *CreateImageResponse {
	s.State = &v
	return s
}

func (s *CreateImageResponse) SetBody(v *CreateImageResponseBody) *CreateImageResponse {
	s.Body = v
	return s
}


type CreateImageResponseBuilder struct {
	s *CreateImageResponse
}

func NewCreateImageResponseBuilder() *CreateImageResponseBuilder {
	s := &CreateImageResponse{}
	b := &CreateImageResponseBuilder{s: s}
	return b
}

func (b *CreateImageResponseBuilder) RequestId(v string) *CreateImageResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *CreateImageResponseBuilder) ErrorMessage(v string) *CreateImageResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *CreateImageResponseBuilder) ErrorCode(v string) *CreateImageResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *CreateImageResponseBuilder) State(v CreateImageResponseStateEnum) *CreateImageResponseBuilder {
	b.s.State = &v
	return b
}

func (b *CreateImageResponseBuilder) Body(v *CreateImageResponseBody) *CreateImageResponseBuilder {
	b.s.Body = v
	return b
}

func (b *CreateImageResponseBuilder) Build() *CreateImageResponse {
	return b.s
}


