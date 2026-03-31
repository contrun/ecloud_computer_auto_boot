// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type UpdateResponseStateEnum string

// List of State
const (
    UpdateResponseStateEnumOk UpdateResponseStateEnum = "OK"
    UpdateResponseStateEnumError UpdateResponseStateEnum = "ERROR"
    UpdateResponseStateEnumException UpdateResponseStateEnum = "EXCEPTION"
    UpdateResponseStateEnumForbidden UpdateResponseStateEnum = "FORBIDDEN"
    UpdateResponseStateEnumRemaining UpdateResponseStateEnum = "REMAINING"
    UpdateResponseStateEnumTimeout UpdateResponseStateEnum = "TIMEOUT"
)

type UpdateResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *UpdateResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *string `json:"body,omitempty"`
}

func (s UpdateResponse) String() string {
	return utils.Beautify(s)
}

func (s UpdateResponse) GoString() string {
	return s.String()
}

func (s UpdateResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateResponse) SetRequestId(v string) *UpdateResponse {
	s.RequestId = &v
	return s
}

func (s *UpdateResponse) SetErrorMessage(v string) *UpdateResponse {
	s.ErrorMessage = &v
	return s
}

func (s *UpdateResponse) SetErrorCode(v string) *UpdateResponse {
	s.ErrorCode = &v
	return s
}

func (s *UpdateResponse) SetState(v UpdateResponseStateEnum) *UpdateResponse {
	s.State = &v
	return s
}

func (s *UpdateResponse) SetBody(v string) *UpdateResponse {
	s.Body = &v
	return s
}


type UpdateResponseBuilder struct {
	s *UpdateResponse
}

func NewUpdateResponseBuilder() *UpdateResponseBuilder {
	s := &UpdateResponse{}
	b := &UpdateResponseBuilder{s: s}
	return b
}

func (b *UpdateResponseBuilder) RequestId(v string) *UpdateResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *UpdateResponseBuilder) ErrorMessage(v string) *UpdateResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *UpdateResponseBuilder) ErrorCode(v string) *UpdateResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *UpdateResponseBuilder) State(v UpdateResponseStateEnum) *UpdateResponseBuilder {
	b.s.State = &v
	return b
}

func (b *UpdateResponseBuilder) Body(v string) *UpdateResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *UpdateResponseBuilder) Build() *UpdateResponse {
	return b.s
}


