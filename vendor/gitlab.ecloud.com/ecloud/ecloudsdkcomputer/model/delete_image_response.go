// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type DeleteImageResponseStateEnum string

// List of State
const (
    DeleteImageResponseStateEnumOk DeleteImageResponseStateEnum = "OK"
    DeleteImageResponseStateEnumError DeleteImageResponseStateEnum = "ERROR"
    DeleteImageResponseStateEnumException DeleteImageResponseStateEnum = "EXCEPTION"
    DeleteImageResponseStateEnumForbidden DeleteImageResponseStateEnum = "FORBIDDEN"
    DeleteImageResponseStateEnumRemaining DeleteImageResponseStateEnum = "REMAINING"
    DeleteImageResponseStateEnumTimeout DeleteImageResponseStateEnum = "TIMEOUT"
)

type DeleteImageResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *DeleteImageResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *string `json:"body,omitempty"`
}

func (s DeleteImageResponse) String() string {
	return utils.Beautify(s)
}

func (s DeleteImageResponse) GoString() string {
	return s.String()
}

func (s DeleteImageResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteImageResponse) SetRequestId(v string) *DeleteImageResponse {
	s.RequestId = &v
	return s
}

func (s *DeleteImageResponse) SetErrorMessage(v string) *DeleteImageResponse {
	s.ErrorMessage = &v
	return s
}

func (s *DeleteImageResponse) SetErrorCode(v string) *DeleteImageResponse {
	s.ErrorCode = &v
	return s
}

func (s *DeleteImageResponse) SetState(v DeleteImageResponseStateEnum) *DeleteImageResponse {
	s.State = &v
	return s
}

func (s *DeleteImageResponse) SetBody(v string) *DeleteImageResponse {
	s.Body = &v
	return s
}


type DeleteImageResponseBuilder struct {
	s *DeleteImageResponse
}

func NewDeleteImageResponseBuilder() *DeleteImageResponseBuilder {
	s := &DeleteImageResponse{}
	b := &DeleteImageResponseBuilder{s: s}
	return b
}

func (b *DeleteImageResponseBuilder) RequestId(v string) *DeleteImageResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *DeleteImageResponseBuilder) ErrorMessage(v string) *DeleteImageResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *DeleteImageResponseBuilder) ErrorCode(v string) *DeleteImageResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *DeleteImageResponseBuilder) State(v DeleteImageResponseStateEnum) *DeleteImageResponseBuilder {
	b.s.State = &v
	return b
}

func (b *DeleteImageResponseBuilder) Body(v string) *DeleteImageResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *DeleteImageResponseBuilder) Build() *DeleteImageResponse {
	return b.s
}


