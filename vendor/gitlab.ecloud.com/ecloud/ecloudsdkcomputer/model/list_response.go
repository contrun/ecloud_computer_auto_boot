// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type ListResponseStateEnum string

// List of State
const (
    ListResponseStateEnumOk ListResponseStateEnum = "OK"
    ListResponseStateEnumError ListResponseStateEnum = "ERROR"
    ListResponseStateEnumException ListResponseStateEnum = "EXCEPTION"
    ListResponseStateEnumForbidden ListResponseStateEnum = "FORBIDDEN"
    ListResponseStateEnumRemaining ListResponseStateEnum = "REMAINING"
    ListResponseStateEnumTimeout ListResponseStateEnum = "TIMEOUT"
)

type ListResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *ListResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *ListResponseBody `json:"body,omitempty"`
}

func (s ListResponse) String() string {
	return utils.Beautify(s)
}

func (s ListResponse) GoString() string {
	return s.String()
}

func (s ListResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListResponse) SetRequestId(v string) *ListResponse {
	s.RequestId = &v
	return s
}

func (s *ListResponse) SetErrorMessage(v string) *ListResponse {
	s.ErrorMessage = &v
	return s
}

func (s *ListResponse) SetErrorCode(v string) *ListResponse {
	s.ErrorCode = &v
	return s
}

func (s *ListResponse) SetState(v ListResponseStateEnum) *ListResponse {
	s.State = &v
	return s
}

func (s *ListResponse) SetBody(v *ListResponseBody) *ListResponse {
	s.Body = v
	return s
}


type ListResponseBuilder struct {
	s *ListResponse
}

func NewListResponseBuilder() *ListResponseBuilder {
	s := &ListResponse{}
	b := &ListResponseBuilder{s: s}
	return b
}

func (b *ListResponseBuilder) RequestId(v string) *ListResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *ListResponseBuilder) ErrorMessage(v string) *ListResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *ListResponseBuilder) ErrorCode(v string) *ListResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *ListResponseBuilder) State(v ListResponseStateEnum) *ListResponseBuilder {
	b.s.State = &v
	return b
}

func (b *ListResponseBuilder) Body(v *ListResponseBody) *ListResponseBuilder {
	b.s.Body = v
	return b
}

func (b *ListResponseBuilder) Build() *ListResponse {
	return b.s
}


