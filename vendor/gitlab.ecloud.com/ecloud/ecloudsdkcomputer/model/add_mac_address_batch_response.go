// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type AddMacAddressBatchResponseStateEnum string

// List of State
const (
    AddMacAddressBatchResponseStateEnumOk AddMacAddressBatchResponseStateEnum = "OK"
    AddMacAddressBatchResponseStateEnumError AddMacAddressBatchResponseStateEnum = "ERROR"
    AddMacAddressBatchResponseStateEnumException AddMacAddressBatchResponseStateEnum = "EXCEPTION"
    AddMacAddressBatchResponseStateEnumForbidden AddMacAddressBatchResponseStateEnum = "FORBIDDEN"
    AddMacAddressBatchResponseStateEnumRemaining AddMacAddressBatchResponseStateEnum = "REMAINING"
    AddMacAddressBatchResponseStateEnumTimeout AddMacAddressBatchResponseStateEnum = "TIMEOUT"
)

type AddMacAddressBatchResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *AddMacAddressBatchResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *bool `json:"body,omitempty"`
}

func (s AddMacAddressBatchResponse) String() string {
	return utils.Beautify(s)
}

func (s AddMacAddressBatchResponse) GoString() string {
	return s.String()
}

func (s AddMacAddressBatchResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *AddMacAddressBatchResponse) SetRequestId(v string) *AddMacAddressBatchResponse {
	s.RequestId = &v
	return s
}

func (s *AddMacAddressBatchResponse) SetErrorMessage(v string) *AddMacAddressBatchResponse {
	s.ErrorMessage = &v
	return s
}

func (s *AddMacAddressBatchResponse) SetErrorCode(v string) *AddMacAddressBatchResponse {
	s.ErrorCode = &v
	return s
}

func (s *AddMacAddressBatchResponse) SetState(v AddMacAddressBatchResponseStateEnum) *AddMacAddressBatchResponse {
	s.State = &v
	return s
}

func (s *AddMacAddressBatchResponse) SetBody(v bool) *AddMacAddressBatchResponse {
	s.Body = &v
	return s
}


type AddMacAddressBatchResponseBuilder struct {
	s *AddMacAddressBatchResponse
}

func NewAddMacAddressBatchResponseBuilder() *AddMacAddressBatchResponseBuilder {
	s := &AddMacAddressBatchResponse{}
	b := &AddMacAddressBatchResponseBuilder{s: s}
	return b
}

func (b *AddMacAddressBatchResponseBuilder) RequestId(v string) *AddMacAddressBatchResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *AddMacAddressBatchResponseBuilder) ErrorMessage(v string) *AddMacAddressBatchResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *AddMacAddressBatchResponseBuilder) ErrorCode(v string) *AddMacAddressBatchResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *AddMacAddressBatchResponseBuilder) State(v AddMacAddressBatchResponseStateEnum) *AddMacAddressBatchResponseBuilder {
	b.s.State = &v
	return b
}

func (b *AddMacAddressBatchResponseBuilder) Body(v bool) *AddMacAddressBatchResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *AddMacAddressBatchResponseBuilder) Build() *AddMacAddressBatchResponse {
	return b.s
}


