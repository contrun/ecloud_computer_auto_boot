// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type GetUserPolicyDetailsResponseStateEnum string

// List of State
const (
    GetUserPolicyDetailsResponseStateEnumOk GetUserPolicyDetailsResponseStateEnum = "OK"
    GetUserPolicyDetailsResponseStateEnumError GetUserPolicyDetailsResponseStateEnum = "ERROR"
    GetUserPolicyDetailsResponseStateEnumException GetUserPolicyDetailsResponseStateEnum = "EXCEPTION"
    GetUserPolicyDetailsResponseStateEnumForbidden GetUserPolicyDetailsResponseStateEnum = "FORBIDDEN"
    GetUserPolicyDetailsResponseStateEnumRemaining GetUserPolicyDetailsResponseStateEnum = "REMAINING"
    GetUserPolicyDetailsResponseStateEnumTimeout GetUserPolicyDetailsResponseStateEnum = "TIMEOUT"
)

type GetUserPolicyDetailsResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *GetUserPolicyDetailsResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *GetUserPolicyDetailsResponseBody `json:"body,omitempty"`
}

func (s GetUserPolicyDetailsResponse) String() string {
	return utils.Beautify(s)
}

func (s GetUserPolicyDetailsResponse) GoString() string {
	return s.String()
}

func (s GetUserPolicyDetailsResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetUserPolicyDetailsResponse) SetRequestId(v string) *GetUserPolicyDetailsResponse {
	s.RequestId = &v
	return s
}

func (s *GetUserPolicyDetailsResponse) SetErrorMessage(v string) *GetUserPolicyDetailsResponse {
	s.ErrorMessage = &v
	return s
}

func (s *GetUserPolicyDetailsResponse) SetErrorCode(v string) *GetUserPolicyDetailsResponse {
	s.ErrorCode = &v
	return s
}

func (s *GetUserPolicyDetailsResponse) SetState(v GetUserPolicyDetailsResponseStateEnum) *GetUserPolicyDetailsResponse {
	s.State = &v
	return s
}

func (s *GetUserPolicyDetailsResponse) SetBody(v *GetUserPolicyDetailsResponseBody) *GetUserPolicyDetailsResponse {
	s.Body = v
	return s
}


type GetUserPolicyDetailsResponseBuilder struct {
	s *GetUserPolicyDetailsResponse
}

func NewGetUserPolicyDetailsResponseBuilder() *GetUserPolicyDetailsResponseBuilder {
	s := &GetUserPolicyDetailsResponse{}
	b := &GetUserPolicyDetailsResponseBuilder{s: s}
	return b
}

func (b *GetUserPolicyDetailsResponseBuilder) RequestId(v string) *GetUserPolicyDetailsResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *GetUserPolicyDetailsResponseBuilder) ErrorMessage(v string) *GetUserPolicyDetailsResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *GetUserPolicyDetailsResponseBuilder) ErrorCode(v string) *GetUserPolicyDetailsResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *GetUserPolicyDetailsResponseBuilder) State(v GetUserPolicyDetailsResponseStateEnum) *GetUserPolicyDetailsResponseBuilder {
	b.s.State = &v
	return b
}

func (b *GetUserPolicyDetailsResponseBuilder) Body(v *GetUserPolicyDetailsResponseBody) *GetUserPolicyDetailsResponseBuilder {
	b.s.Body = v
	return b
}

func (b *GetUserPolicyDetailsResponseBuilder) Build() *GetUserPolicyDetailsResponse {
	return b.s
}


