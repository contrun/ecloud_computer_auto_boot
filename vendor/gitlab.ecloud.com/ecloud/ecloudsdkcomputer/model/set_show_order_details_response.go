// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type SetShowOrderDetailsResponseStateEnum string

// List of State
const (
    SetShowOrderDetailsResponseStateEnumOk SetShowOrderDetailsResponseStateEnum = "OK"
    SetShowOrderDetailsResponseStateEnumError SetShowOrderDetailsResponseStateEnum = "ERROR"
    SetShowOrderDetailsResponseStateEnumException SetShowOrderDetailsResponseStateEnum = "EXCEPTION"
    SetShowOrderDetailsResponseStateEnumForbidden SetShowOrderDetailsResponseStateEnum = "FORBIDDEN"
    SetShowOrderDetailsResponseStateEnumRemaining SetShowOrderDetailsResponseStateEnum = "REMAINING"
    SetShowOrderDetailsResponseStateEnumTimeout SetShowOrderDetailsResponseStateEnum = "TIMEOUT"
)

type SetShowOrderDetailsResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *SetShowOrderDetailsResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *bool `json:"body,omitempty"`
}

func (s SetShowOrderDetailsResponse) String() string {
	return utils.Beautify(s)
}

func (s SetShowOrderDetailsResponse) GoString() string {
	return s.String()
}

func (s SetShowOrderDetailsResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *SetShowOrderDetailsResponse) SetRequestId(v string) *SetShowOrderDetailsResponse {
	s.RequestId = &v
	return s
}

func (s *SetShowOrderDetailsResponse) SetErrorMessage(v string) *SetShowOrderDetailsResponse {
	s.ErrorMessage = &v
	return s
}

func (s *SetShowOrderDetailsResponse) SetErrorCode(v string) *SetShowOrderDetailsResponse {
	s.ErrorCode = &v
	return s
}

func (s *SetShowOrderDetailsResponse) SetState(v SetShowOrderDetailsResponseStateEnum) *SetShowOrderDetailsResponse {
	s.State = &v
	return s
}

func (s *SetShowOrderDetailsResponse) SetBody(v bool) *SetShowOrderDetailsResponse {
	s.Body = &v
	return s
}


type SetShowOrderDetailsResponseBuilder struct {
	s *SetShowOrderDetailsResponse
}

func NewSetShowOrderDetailsResponseBuilder() *SetShowOrderDetailsResponseBuilder {
	s := &SetShowOrderDetailsResponse{}
	b := &SetShowOrderDetailsResponseBuilder{s: s}
	return b
}

func (b *SetShowOrderDetailsResponseBuilder) RequestId(v string) *SetShowOrderDetailsResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *SetShowOrderDetailsResponseBuilder) ErrorMessage(v string) *SetShowOrderDetailsResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *SetShowOrderDetailsResponseBuilder) ErrorCode(v string) *SetShowOrderDetailsResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *SetShowOrderDetailsResponseBuilder) State(v SetShowOrderDetailsResponseStateEnum) *SetShowOrderDetailsResponseBuilder {
	b.s.State = &v
	return b
}

func (b *SetShowOrderDetailsResponseBuilder) Body(v bool) *SetShowOrderDetailsResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *SetShowOrderDetailsResponseBuilder) Build() *SetShowOrderDetailsResponse {
	return b.s
}


