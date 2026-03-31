// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type DeleteResponseStateEnum string

// List of State
const (
    DeleteResponseStateEnumOk DeleteResponseStateEnum = "OK"
    DeleteResponseStateEnumError DeleteResponseStateEnum = "ERROR"
    DeleteResponseStateEnumException DeleteResponseStateEnum = "EXCEPTION"
    DeleteResponseStateEnumForbidden DeleteResponseStateEnum = "FORBIDDEN"
    DeleteResponseStateEnumRemaining DeleteResponseStateEnum = "REMAINING"
    DeleteResponseStateEnumTimeout DeleteResponseStateEnum = "TIMEOUT"
)

type DeleteResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *DeleteResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *string `json:"body,omitempty"`
}

func (s DeleteResponse) String() string {
	return utils.Beautify(s)
}

func (s DeleteResponse) GoString() string {
	return s.String()
}

func (s DeleteResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteResponse) SetRequestId(v string) *DeleteResponse {
	s.RequestId = &v
	return s
}

func (s *DeleteResponse) SetErrorMessage(v string) *DeleteResponse {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteResponse) SetErrorCode(v string) *DeleteResponse {
	s.ErrorCode = &v
	return s
}

func (s *DeleteResponse) SetState(v DeleteResponseStateEnum) *DeleteResponse {
	s.State = &v
	return s
}

func (s *DeleteResponse) SetBody(v string) *DeleteResponse {
	s.Body = &v
	return s
}


type DeleteResponseBuilder struct {
	s *DeleteResponse
}

func NewDeleteResponseBuilder() *DeleteResponseBuilder {
	s := &DeleteResponse{}
	b := &DeleteResponseBuilder{s: s}
	return b
}

func (b *DeleteResponseBuilder) RequestId(v string) *DeleteResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *DeleteResponseBuilder) ErrorMessage(v string) *DeleteResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *DeleteResponseBuilder) ErrorCode(v string) *DeleteResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *DeleteResponseBuilder) State(v DeleteResponseStateEnum) *DeleteResponseBuilder {
	b.s.State = &v
	return b
}

func (b *DeleteResponseBuilder) Body(v string) *DeleteResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *DeleteResponseBuilder) Build() *DeleteResponse {
	return b.s
}


