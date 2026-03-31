// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type GetShowOrderDetailsResponseStateEnum string

// List of State
const (
    GetShowOrderDetailsResponseStateEnumOk GetShowOrderDetailsResponseStateEnum = "OK"
    GetShowOrderDetailsResponseStateEnumError GetShowOrderDetailsResponseStateEnum = "ERROR"
    GetShowOrderDetailsResponseStateEnumException GetShowOrderDetailsResponseStateEnum = "EXCEPTION"
    GetShowOrderDetailsResponseStateEnumForbidden GetShowOrderDetailsResponseStateEnum = "FORBIDDEN"
    GetShowOrderDetailsResponseStateEnumRemaining GetShowOrderDetailsResponseStateEnum = "REMAINING"
    GetShowOrderDetailsResponseStateEnumTimeout GetShowOrderDetailsResponseStateEnum = "TIMEOUT"
)

type GetShowOrderDetailsResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *GetShowOrderDetailsResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *GetShowOrderDetailsResponseBody `json:"body,omitempty"`
}

func (s GetShowOrderDetailsResponse) String() string {
	return utils.Beautify(s)
}

func (s GetShowOrderDetailsResponse) GoString() string {
	return s.String()
}

func (s GetShowOrderDetailsResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetShowOrderDetailsResponse) SetRequestId(v string) *GetShowOrderDetailsResponse {
	s.RequestId = &v
	return s
}

func (s *GetShowOrderDetailsResponse) SetErrorMessage(v string) *GetShowOrderDetailsResponse {
	s.ErrorMessage = &v
	return s
}

func (s *GetShowOrderDetailsResponse) SetErrorCode(v string) *GetShowOrderDetailsResponse {
	s.ErrorCode = &v
	return s
}

func (s *GetShowOrderDetailsResponse) SetState(v GetShowOrderDetailsResponseStateEnum) *GetShowOrderDetailsResponse {
	s.State = &v
	return s
}

func (s *GetShowOrderDetailsResponse) SetBody(v *GetShowOrderDetailsResponseBody) *GetShowOrderDetailsResponse {
	s.Body = v
	return s
}


type GetShowOrderDetailsResponseBuilder struct {
	s *GetShowOrderDetailsResponse
}

func NewGetShowOrderDetailsResponseBuilder() *GetShowOrderDetailsResponseBuilder {
	s := &GetShowOrderDetailsResponse{}
	b := &GetShowOrderDetailsResponseBuilder{s: s}
	return b
}

func (b *GetShowOrderDetailsResponseBuilder) RequestId(v string) *GetShowOrderDetailsResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *GetShowOrderDetailsResponseBuilder) ErrorMessage(v string) *GetShowOrderDetailsResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *GetShowOrderDetailsResponseBuilder) ErrorCode(v string) *GetShowOrderDetailsResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *GetShowOrderDetailsResponseBuilder) State(v GetShowOrderDetailsResponseStateEnum) *GetShowOrderDetailsResponseBuilder {
	b.s.State = &v
	return b
}

func (b *GetShowOrderDetailsResponseBuilder) Body(v *GetShowOrderDetailsResponseBody) *GetShowOrderDetailsResponseBuilder {
	b.s.Body = v
	return b
}

func (b *GetShowOrderDetailsResponseBuilder) Build() *GetShowOrderDetailsResponse {
	return b.s
}


