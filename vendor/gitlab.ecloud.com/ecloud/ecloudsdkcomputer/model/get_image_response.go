// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type GetImageResponseStateEnum string

// List of State
const (
    GetImageResponseStateEnumOk GetImageResponseStateEnum = "OK"
    GetImageResponseStateEnumError GetImageResponseStateEnum = "ERROR"
    GetImageResponseStateEnumException GetImageResponseStateEnum = "EXCEPTION"
    GetImageResponseStateEnumForbidden GetImageResponseStateEnum = "FORBIDDEN"
    GetImageResponseStateEnumRemaining GetImageResponseStateEnum = "REMAINING"
    GetImageResponseStateEnumTimeout GetImageResponseStateEnum = "TIMEOUT"
)

type GetImageResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *GetImageResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *GetImageResponseBody `json:"body,omitempty"`
}

func (s GetImageResponse) String() string {
	return utils.Beautify(s)
}

func (s GetImageResponse) GoString() string {
	return s.String()
}

func (s GetImageResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetImageResponse) SetRequestId(v string) *GetImageResponse {
	s.RequestId = &v
	return s
}

func (s *GetImageResponse) SetErrorMessage(v string) *GetImageResponse {
	s.ErrorMessage = &v
	return s
}

func (s *GetImageResponse) SetErrorCode(v string) *GetImageResponse {
	s.ErrorCode = &v
	return s
}

func (s *GetImageResponse) SetState(v GetImageResponseStateEnum) *GetImageResponse {
	s.State = &v
	return s
}

func (s *GetImageResponse) SetBody(v *GetImageResponseBody) *GetImageResponse {
	s.Body = v
	return s
}


type GetImageResponseBuilder struct {
	s *GetImageResponse
}

func NewGetImageResponseBuilder() *GetImageResponseBuilder {
	s := &GetImageResponse{}
	b := &GetImageResponseBuilder{s: s}
	return b
}

func (b *GetImageResponseBuilder) RequestId(v string) *GetImageResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *GetImageResponseBuilder) ErrorMessage(v string) *GetImageResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *GetImageResponseBuilder) ErrorCode(v string) *GetImageResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *GetImageResponseBuilder) State(v GetImageResponseStateEnum) *GetImageResponseBuilder {
	b.s.State = &v
	return b
}

func (b *GetImageResponseBuilder) Body(v *GetImageResponseBody) *GetImageResponseBuilder {
	b.s.Body = v
	return b
}

func (b *GetImageResponseBuilder) Build() *GetImageResponse {
	return b.s
}


