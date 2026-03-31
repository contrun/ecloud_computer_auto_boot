// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type UserExdpolicyPermissionResponseStateEnum string

// List of State
const (
    UserExdpolicyPermissionResponseStateEnumOk UserExdpolicyPermissionResponseStateEnum = "OK"
    UserExdpolicyPermissionResponseStateEnumError UserExdpolicyPermissionResponseStateEnum = "ERROR"
    UserExdpolicyPermissionResponseStateEnumException UserExdpolicyPermissionResponseStateEnum = "EXCEPTION"
    UserExdpolicyPermissionResponseStateEnumForbidden UserExdpolicyPermissionResponseStateEnum = "FORBIDDEN"
    UserExdpolicyPermissionResponseStateEnumRemaining UserExdpolicyPermissionResponseStateEnum = "REMAINING"
    UserExdpolicyPermissionResponseStateEnumTimeout UserExdpolicyPermissionResponseStateEnum = "TIMEOUT"
)

type UserExdpolicyPermissionResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *UserExdpolicyPermissionResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *UserExdpolicyPermissionResponseBody `json:"body,omitempty"`
}

func (s UserExdpolicyPermissionResponse) String() string {
	return utils.Beautify(s)
}

func (s UserExdpolicyPermissionResponse) GoString() string {
	return s.String()
}

func (s UserExdpolicyPermissionResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UserExdpolicyPermissionResponse) SetRequestId(v string) *UserExdpolicyPermissionResponse {
	s.RequestId = &v
	return s
}

func (s *UserExdpolicyPermissionResponse) SetErrorMessage(v string) *UserExdpolicyPermissionResponse {
	s.ErrorMessage = &v
	return s
}

func (s *UserExdpolicyPermissionResponse) SetErrorCode(v string) *UserExdpolicyPermissionResponse {
	s.ErrorCode = &v
	return s
}

func (s *UserExdpolicyPermissionResponse) SetState(v UserExdpolicyPermissionResponseStateEnum) *UserExdpolicyPermissionResponse {
	s.State = &v
	return s
}

func (s *UserExdpolicyPermissionResponse) SetBody(v *UserExdpolicyPermissionResponseBody) *UserExdpolicyPermissionResponse {
	s.Body = v
	return s
}


type UserExdpolicyPermissionResponseBuilder struct {
	s *UserExdpolicyPermissionResponse
}

func NewUserExdpolicyPermissionResponseBuilder() *UserExdpolicyPermissionResponseBuilder {
	s := &UserExdpolicyPermissionResponse{}
	b := &UserExdpolicyPermissionResponseBuilder{s: s}
	return b
}

func (b *UserExdpolicyPermissionResponseBuilder) RequestId(v string) *UserExdpolicyPermissionResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *UserExdpolicyPermissionResponseBuilder) ErrorMessage(v string) *UserExdpolicyPermissionResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *UserExdpolicyPermissionResponseBuilder) ErrorCode(v string) *UserExdpolicyPermissionResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *UserExdpolicyPermissionResponseBuilder) State(v UserExdpolicyPermissionResponseStateEnum) *UserExdpolicyPermissionResponseBuilder {
	b.s.State = &v
	return b
}

func (b *UserExdpolicyPermissionResponseBuilder) Body(v *UserExdpolicyPermissionResponseBody) *UserExdpolicyPermissionResponseBuilder {
	b.s.Body = v
	return b
}

func (b *UserExdpolicyPermissionResponseBuilder) Build() *UserExdpolicyPermissionResponse {
	return b.s
}


