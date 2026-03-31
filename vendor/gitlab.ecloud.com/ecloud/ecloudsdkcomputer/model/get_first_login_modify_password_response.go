// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type GetFirstLoginModifyPasswordResponseStateEnum string

// List of State
const (
    GetFirstLoginModifyPasswordResponseStateEnumOk GetFirstLoginModifyPasswordResponseStateEnum = "OK"
    GetFirstLoginModifyPasswordResponseStateEnumError GetFirstLoginModifyPasswordResponseStateEnum = "ERROR"
    GetFirstLoginModifyPasswordResponseStateEnumException GetFirstLoginModifyPasswordResponseStateEnum = "EXCEPTION"
    GetFirstLoginModifyPasswordResponseStateEnumForbidden GetFirstLoginModifyPasswordResponseStateEnum = "FORBIDDEN"
    GetFirstLoginModifyPasswordResponseStateEnumRemaining GetFirstLoginModifyPasswordResponseStateEnum = "REMAINING"
    GetFirstLoginModifyPasswordResponseStateEnumTimeout GetFirstLoginModifyPasswordResponseStateEnum = "TIMEOUT"
)

type GetFirstLoginModifyPasswordResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *GetFirstLoginModifyPasswordResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *GetFirstLoginModifyPasswordResponseBody `json:"body,omitempty"`
}

func (s GetFirstLoginModifyPasswordResponse) String() string {
	return utils.Beautify(s)
}

func (s GetFirstLoginModifyPasswordResponse) GoString() string {
	return s.String()
}

func (s GetFirstLoginModifyPasswordResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetFirstLoginModifyPasswordResponse) SetRequestId(v string) *GetFirstLoginModifyPasswordResponse {
	s.RequestId = &v
	return s
}

func (s *GetFirstLoginModifyPasswordResponse) SetErrorMessage(v string) *GetFirstLoginModifyPasswordResponse {
	s.ErrorMessage = &v
	return s
}

func (s *GetFirstLoginModifyPasswordResponse) SetErrorCode(v string) *GetFirstLoginModifyPasswordResponse {
	s.ErrorCode = &v
	return s
}

func (s *GetFirstLoginModifyPasswordResponse) SetState(v GetFirstLoginModifyPasswordResponseStateEnum) *GetFirstLoginModifyPasswordResponse {
	s.State = &v
	return s
}

func (s *GetFirstLoginModifyPasswordResponse) SetBody(v *GetFirstLoginModifyPasswordResponseBody) *GetFirstLoginModifyPasswordResponse {
	s.Body = v
	return s
}


type GetFirstLoginModifyPasswordResponseBuilder struct {
	s *GetFirstLoginModifyPasswordResponse
}

func NewGetFirstLoginModifyPasswordResponseBuilder() *GetFirstLoginModifyPasswordResponseBuilder {
	s := &GetFirstLoginModifyPasswordResponse{}
	b := &GetFirstLoginModifyPasswordResponseBuilder{s: s}
	return b
}

func (b *GetFirstLoginModifyPasswordResponseBuilder) RequestId(v string) *GetFirstLoginModifyPasswordResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *GetFirstLoginModifyPasswordResponseBuilder) ErrorMessage(v string) *GetFirstLoginModifyPasswordResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *GetFirstLoginModifyPasswordResponseBuilder) ErrorCode(v string) *GetFirstLoginModifyPasswordResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *GetFirstLoginModifyPasswordResponseBuilder) State(v GetFirstLoginModifyPasswordResponseStateEnum) *GetFirstLoginModifyPasswordResponseBuilder {
	b.s.State = &v
	return b
}

func (b *GetFirstLoginModifyPasswordResponseBuilder) Body(v *GetFirstLoginModifyPasswordResponseBody) *GetFirstLoginModifyPasswordResponseBuilder {
	b.s.Body = v
	return b
}

func (b *GetFirstLoginModifyPasswordResponseBuilder) Build() *GetFirstLoginModifyPasswordResponse {
	return b.s
}


