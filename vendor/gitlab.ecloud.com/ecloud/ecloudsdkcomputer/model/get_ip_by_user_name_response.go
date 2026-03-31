// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type GetIPByUserNameResponseStateEnum string

// List of State
const (
    GetIPByUserNameResponseStateEnumOk GetIPByUserNameResponseStateEnum = "OK"
    GetIPByUserNameResponseStateEnumError GetIPByUserNameResponseStateEnum = "ERROR"
    GetIPByUserNameResponseStateEnumException GetIPByUserNameResponseStateEnum = "EXCEPTION"
    GetIPByUserNameResponseStateEnumForbidden GetIPByUserNameResponseStateEnum = "FORBIDDEN"
    GetIPByUserNameResponseStateEnumRemaining GetIPByUserNameResponseStateEnum = "REMAINING"
    GetIPByUserNameResponseStateEnumTimeout GetIPByUserNameResponseStateEnum = "TIMEOUT"
)

type GetIPByUserNameResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *GetIPByUserNameResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *string `json:"body,omitempty"`
}

func (s GetIPByUserNameResponse) String() string {
	return utils.Beautify(s)
}

func (s GetIPByUserNameResponse) GoString() string {
	return s.String()
}

func (s GetIPByUserNameResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetIPByUserNameResponse) SetRequestId(v string) *GetIPByUserNameResponse {
	s.RequestId = &v
	return s
}

func (s *GetIPByUserNameResponse) SetErrorMessage(v string) *GetIPByUserNameResponse {
	s.ErrorMessage = &v
	return s
}

func (s *GetIPByUserNameResponse) SetErrorCode(v string) *GetIPByUserNameResponse {
	s.ErrorCode = &v
	return s
}

func (s *GetIPByUserNameResponse) SetState(v GetIPByUserNameResponseStateEnum) *GetIPByUserNameResponse {
	s.State = &v
	return s
}

func (s *GetIPByUserNameResponse) SetBody(v string) *GetIPByUserNameResponse {
	s.Body = &v
	return s
}


type GetIPByUserNameResponseBuilder struct {
	s *GetIPByUserNameResponse
}

func NewGetIPByUserNameResponseBuilder() *GetIPByUserNameResponseBuilder {
	s := &GetIPByUserNameResponse{}
	b := &GetIPByUserNameResponseBuilder{s: s}
	return b
}

func (b *GetIPByUserNameResponseBuilder) RequestId(v string) *GetIPByUserNameResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *GetIPByUserNameResponseBuilder) ErrorMessage(v string) *GetIPByUserNameResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *GetIPByUserNameResponseBuilder) ErrorCode(v string) *GetIPByUserNameResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *GetIPByUserNameResponseBuilder) State(v GetIPByUserNameResponseStateEnum) *GetIPByUserNameResponseBuilder {
	b.s.State = &v
	return b
}

func (b *GetIPByUserNameResponseBuilder) Body(v string) *GetIPByUserNameResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *GetIPByUserNameResponseBuilder) Build() *GetIPByUserNameResponse {
	return b.s
}


