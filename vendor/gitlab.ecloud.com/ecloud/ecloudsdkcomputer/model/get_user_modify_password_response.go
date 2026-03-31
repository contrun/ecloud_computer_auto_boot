// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type GetUserModifyPasswordResponseStateEnum string

// List of State
const (
    GetUserModifyPasswordResponseStateEnumOk GetUserModifyPasswordResponseStateEnum = "OK"
    GetUserModifyPasswordResponseStateEnumError GetUserModifyPasswordResponseStateEnum = "ERROR"
    GetUserModifyPasswordResponseStateEnumException GetUserModifyPasswordResponseStateEnum = "EXCEPTION"
    GetUserModifyPasswordResponseStateEnumForbidden GetUserModifyPasswordResponseStateEnum = "FORBIDDEN"
    GetUserModifyPasswordResponseStateEnumRemaining GetUserModifyPasswordResponseStateEnum = "REMAINING"
    GetUserModifyPasswordResponseStateEnumTimeout GetUserModifyPasswordResponseStateEnum = "TIMEOUT"
)

type GetUserModifyPasswordResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *GetUserModifyPasswordResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *GetUserModifyPasswordResponseBody `json:"body,omitempty"`
}

func (s GetUserModifyPasswordResponse) String() string {
	return utils.Beautify(s)
}

func (s GetUserModifyPasswordResponse) GoString() string {
	return s.String()
}

func (s GetUserModifyPasswordResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetUserModifyPasswordResponse) SetRequestId(v string) *GetUserModifyPasswordResponse {
	s.RequestId = &v
	return s
}

func (s *GetUserModifyPasswordResponse) SetErrorMessage(v string) *GetUserModifyPasswordResponse {
	s.ErrorMessage = &v
	return s
}

func (s *GetUserModifyPasswordResponse) SetErrorCode(v string) *GetUserModifyPasswordResponse {
	s.ErrorCode = &v
	return s
}

func (s *GetUserModifyPasswordResponse) SetState(v GetUserModifyPasswordResponseStateEnum) *GetUserModifyPasswordResponse {
	s.State = &v
	return s
}

func (s *GetUserModifyPasswordResponse) SetBody(v *GetUserModifyPasswordResponseBody) *GetUserModifyPasswordResponse {
	s.Body = v
	return s
}


type GetUserModifyPasswordResponseBuilder struct {
	s *GetUserModifyPasswordResponse
}

func NewGetUserModifyPasswordResponseBuilder() *GetUserModifyPasswordResponseBuilder {
	s := &GetUserModifyPasswordResponse{}
	b := &GetUserModifyPasswordResponseBuilder{s: s}
	return b
}

func (b *GetUserModifyPasswordResponseBuilder) RequestId(v string) *GetUserModifyPasswordResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *GetUserModifyPasswordResponseBuilder) ErrorMessage(v string) *GetUserModifyPasswordResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *GetUserModifyPasswordResponseBuilder) ErrorCode(v string) *GetUserModifyPasswordResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *GetUserModifyPasswordResponseBuilder) State(v GetUserModifyPasswordResponseStateEnum) *GetUserModifyPasswordResponseBuilder {
	b.s.State = &v
	return b
}

func (b *GetUserModifyPasswordResponseBuilder) Body(v *GetUserModifyPasswordResponseBody) *GetUserModifyPasswordResponseBuilder {
	b.s.Body = v
	return b
}

func (b *GetUserModifyPasswordResponseBuilder) Build() *GetUserModifyPasswordResponse {
	return b.s
}


