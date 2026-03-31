// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type AddMacAddressResponseStateEnum string

// List of State
const (
    AddMacAddressResponseStateEnumOk AddMacAddressResponseStateEnum = "OK"
    AddMacAddressResponseStateEnumError AddMacAddressResponseStateEnum = "ERROR"
    AddMacAddressResponseStateEnumException AddMacAddressResponseStateEnum = "EXCEPTION"
    AddMacAddressResponseStateEnumForbidden AddMacAddressResponseStateEnum = "FORBIDDEN"
    AddMacAddressResponseStateEnumRemaining AddMacAddressResponseStateEnum = "REMAINING"
    AddMacAddressResponseStateEnumTimeout AddMacAddressResponseStateEnum = "TIMEOUT"
)

type AddMacAddressResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *AddMacAddressResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *bool `json:"body,omitempty"`
}

func (s AddMacAddressResponse) String() string {
	return utils.Beautify(s)
}

func (s AddMacAddressResponse) GoString() string {
	return s.String()
}

func (s AddMacAddressResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *AddMacAddressResponse) SetRequestId(v string) *AddMacAddressResponse {
	s.RequestId = &v
	return s
}

func (s *AddMacAddressResponse) SetErrorMessage(v string) *AddMacAddressResponse {
	s.ErrorMessage = &v
	return s
}

func (s *AddMacAddressResponse) SetErrorCode(v string) *AddMacAddressResponse {
	s.ErrorCode = &v
	return s
}

func (s *AddMacAddressResponse) SetState(v AddMacAddressResponseStateEnum) *AddMacAddressResponse {
	s.State = &v
	return s
}

func (s *AddMacAddressResponse) SetBody(v bool) *AddMacAddressResponse {
	s.Body = &v
	return s
}


type AddMacAddressResponseBuilder struct {
	s *AddMacAddressResponse
}

func NewAddMacAddressResponseBuilder() *AddMacAddressResponseBuilder {
	s := &AddMacAddressResponse{}
	b := &AddMacAddressResponseBuilder{s: s}
	return b
}

func (b *AddMacAddressResponseBuilder) RequestId(v string) *AddMacAddressResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *AddMacAddressResponseBuilder) ErrorMessage(v string) *AddMacAddressResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *AddMacAddressResponseBuilder) ErrorCode(v string) *AddMacAddressResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *AddMacAddressResponseBuilder) State(v AddMacAddressResponseStateEnum) *AddMacAddressResponseBuilder {
	b.s.State = &v
	return b
}

func (b *AddMacAddressResponseBuilder) Body(v bool) *AddMacAddressResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *AddMacAddressResponseBuilder) Build() *AddMacAddressResponse {
	return b.s
}


