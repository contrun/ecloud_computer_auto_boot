// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type VerifyAccessTicketResponseStateEnum string

// List of State
const (
    VerifyAccessTicketResponseStateEnumOk VerifyAccessTicketResponseStateEnum = "OK"
    VerifyAccessTicketResponseStateEnumError VerifyAccessTicketResponseStateEnum = "ERROR"
    VerifyAccessTicketResponseStateEnumException VerifyAccessTicketResponseStateEnum = "EXCEPTION"
    VerifyAccessTicketResponseStateEnumForbidden VerifyAccessTicketResponseStateEnum = "FORBIDDEN"
    VerifyAccessTicketResponseStateEnumRemaining VerifyAccessTicketResponseStateEnum = "REMAINING"
    VerifyAccessTicketResponseStateEnumTimeout VerifyAccessTicketResponseStateEnum = "TIMEOUT"
)

type VerifyAccessTicketResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *VerifyAccessTicketResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *VerifyAccessTicketResponseBody `json:"body,omitempty"`
}

func (s VerifyAccessTicketResponse) String() string {
	return utils.Beautify(s)
}

func (s VerifyAccessTicketResponse) GoString() string {
	return s.String()
}

func (s VerifyAccessTicketResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *VerifyAccessTicketResponse) SetRequestId(v string) *VerifyAccessTicketResponse {
	s.RequestId = &v
	return s
}

func (s *VerifyAccessTicketResponse) SetErrorMessage(v string) *VerifyAccessTicketResponse {
	s.ErrorMessage = &v
	return s
}

func (s *VerifyAccessTicketResponse) SetErrorCode(v string) *VerifyAccessTicketResponse {
	s.ErrorCode = &v
	return s
}

func (s *VerifyAccessTicketResponse) SetState(v VerifyAccessTicketResponseStateEnum) *VerifyAccessTicketResponse {
	s.State = &v
	return s
}

func (s *VerifyAccessTicketResponse) SetBody(v *VerifyAccessTicketResponseBody) *VerifyAccessTicketResponse {
	s.Body = v
	return s
}


type VerifyAccessTicketResponseBuilder struct {
	s *VerifyAccessTicketResponse
}

func NewVerifyAccessTicketResponseBuilder() *VerifyAccessTicketResponseBuilder {
	s := &VerifyAccessTicketResponse{}
	b := &VerifyAccessTicketResponseBuilder{s: s}
	return b
}

func (b *VerifyAccessTicketResponseBuilder) RequestId(v string) *VerifyAccessTicketResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *VerifyAccessTicketResponseBuilder) ErrorMessage(v string) *VerifyAccessTicketResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *VerifyAccessTicketResponseBuilder) ErrorCode(v string) *VerifyAccessTicketResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *VerifyAccessTicketResponseBuilder) State(v VerifyAccessTicketResponseStateEnum) *VerifyAccessTicketResponseBuilder {
	b.s.State = &v
	return b
}

func (b *VerifyAccessTicketResponseBuilder) Body(v *VerifyAccessTicketResponseBody) *VerifyAccessTicketResponseBuilder {
	b.s.Body = v
	return b
}

func (b *VerifyAccessTicketResponseBuilder) Build() *VerifyAccessTicketResponse {
	return b.s
}


