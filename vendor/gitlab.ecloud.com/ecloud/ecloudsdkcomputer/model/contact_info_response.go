// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type ContactInfoResponseStateEnum string

// List of State
const (
    ContactInfoResponseStateEnumOk ContactInfoResponseStateEnum = "OK"
    ContactInfoResponseStateEnumError ContactInfoResponseStateEnum = "ERROR"
    ContactInfoResponseStateEnumException ContactInfoResponseStateEnum = "EXCEPTION"
    ContactInfoResponseStateEnumForbidden ContactInfoResponseStateEnum = "FORBIDDEN"
    ContactInfoResponseStateEnumRemaining ContactInfoResponseStateEnum = "REMAINING"
    ContactInfoResponseStateEnumTimeout ContactInfoResponseStateEnum = "TIMEOUT"
)

type ContactInfoResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *ContactInfoResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *ContactInfoResponseBody `json:"body,omitempty"`
}

func (s ContactInfoResponse) String() string {
	return utils.Beautify(s)
}

func (s ContactInfoResponse) GoString() string {
	return s.String()
}

func (s ContactInfoResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ContactInfoResponse) SetRequestId(v string) *ContactInfoResponse {
	s.RequestId = &v
	return s
}

func (s *ContactInfoResponse) SetErrorMessage(v string) *ContactInfoResponse {
	s.ErrorMessage = &v
	return s
}

func (s *ContactInfoResponse) SetErrorCode(v string) *ContactInfoResponse {
	s.ErrorCode = &v
	return s
}

func (s *ContactInfoResponse) SetState(v ContactInfoResponseStateEnum) *ContactInfoResponse {
	s.State = &v
	return s
}

func (s *ContactInfoResponse) SetBody(v *ContactInfoResponseBody) *ContactInfoResponse {
	s.Body = v
	return s
}


type ContactInfoResponseBuilder struct {
	s *ContactInfoResponse
}

func NewContactInfoResponseBuilder() *ContactInfoResponseBuilder {
	s := &ContactInfoResponse{}
	b := &ContactInfoResponseBuilder{s: s}
	return b
}

func (b *ContactInfoResponseBuilder) RequestId(v string) *ContactInfoResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *ContactInfoResponseBuilder) ErrorMessage(v string) *ContactInfoResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *ContactInfoResponseBuilder) ErrorCode(v string) *ContactInfoResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *ContactInfoResponseBuilder) State(v ContactInfoResponseStateEnum) *ContactInfoResponseBuilder {
	b.s.State = &v
	return b
}

func (b *ContactInfoResponseBuilder) Body(v *ContactInfoResponseBody) *ContactInfoResponseBuilder {
	b.s.Body = v
	return b
}

func (b *ContactInfoResponseBuilder) Build() *ContactInfoResponse {
	return b.s
}


