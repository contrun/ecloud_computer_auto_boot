// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type ModifyUserInfoResponseStateEnum string

// List of State
const (
    ModifyUserInfoResponseStateEnumOk ModifyUserInfoResponseStateEnum = "OK"
    ModifyUserInfoResponseStateEnumError ModifyUserInfoResponseStateEnum = "ERROR"
    ModifyUserInfoResponseStateEnumException ModifyUserInfoResponseStateEnum = "EXCEPTION"
    ModifyUserInfoResponseStateEnumForbidden ModifyUserInfoResponseStateEnum = "FORBIDDEN"
    ModifyUserInfoResponseStateEnumRemaining ModifyUserInfoResponseStateEnum = "REMAINING"
    ModifyUserInfoResponseStateEnumTimeout ModifyUserInfoResponseStateEnum = "TIMEOUT"
)

type ModifyUserInfoResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *ModifyUserInfoResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *ModifyUserInfoResponseBody `json:"body,omitempty"`
}

func (s ModifyUserInfoResponse) String() string {
	return utils.Beautify(s)
}

func (s ModifyUserInfoResponse) GoString() string {
	return s.String()
}

func (s ModifyUserInfoResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ModifyUserInfoResponse) SetRequestId(v string) *ModifyUserInfoResponse {
	s.RequestId = &v
	return s
}

func (s *ModifyUserInfoResponse) SetErrorMessage(v string) *ModifyUserInfoResponse {
	s.ErrorMessage = &v
	return s
}

func (s *ModifyUserInfoResponse) SetErrorCode(v string) *ModifyUserInfoResponse {
	s.ErrorCode = &v
	return s
}

func (s *ModifyUserInfoResponse) SetState(v ModifyUserInfoResponseStateEnum) *ModifyUserInfoResponse {
	s.State = &v
	return s
}

func (s *ModifyUserInfoResponse) SetBody(v *ModifyUserInfoResponseBody) *ModifyUserInfoResponse {
	s.Body = v
	return s
}


type ModifyUserInfoResponseBuilder struct {
	s *ModifyUserInfoResponse
}

func NewModifyUserInfoResponseBuilder() *ModifyUserInfoResponseBuilder {
	s := &ModifyUserInfoResponse{}
	b := &ModifyUserInfoResponseBuilder{s: s}
	return b
}

func (b *ModifyUserInfoResponseBuilder) RequestId(v string) *ModifyUserInfoResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *ModifyUserInfoResponseBuilder) ErrorMessage(v string) *ModifyUserInfoResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *ModifyUserInfoResponseBuilder) ErrorCode(v string) *ModifyUserInfoResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *ModifyUserInfoResponseBuilder) State(v ModifyUserInfoResponseStateEnum) *ModifyUserInfoResponseBuilder {
	b.s.State = &v
	return b
}

func (b *ModifyUserInfoResponseBuilder) Body(v *ModifyUserInfoResponseBody) *ModifyUserInfoResponseBuilder {
	b.s.Body = v
	return b
}

func (b *ModifyUserInfoResponseBuilder) Build() *ModifyUserInfoResponse {
	return b.s
}


