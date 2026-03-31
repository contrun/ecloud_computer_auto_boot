// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type GetPoolListResponseStateEnum string

// List of State
const (
    GetPoolListResponseStateEnumOk GetPoolListResponseStateEnum = "OK"
    GetPoolListResponseStateEnumError GetPoolListResponseStateEnum = "ERROR"
    GetPoolListResponseStateEnumException GetPoolListResponseStateEnum = "EXCEPTION"
    GetPoolListResponseStateEnumForbidden GetPoolListResponseStateEnum = "FORBIDDEN"
    GetPoolListResponseStateEnumRemaining GetPoolListResponseStateEnum = "REMAINING"
    GetPoolListResponseStateEnumTimeout GetPoolListResponseStateEnum = "TIMEOUT"
)

type GetPoolListResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *GetPoolListResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *[]GetPoolListResponseBody `json:"body,omitempty"`
}

func (s GetPoolListResponse) String() string {
	return utils.Beautify(s)
}

func (s GetPoolListResponse) GoString() string {
	return s.String()
}

func (s GetPoolListResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetPoolListResponse) SetRequestId(v string) *GetPoolListResponse {
	s.RequestId = &v
	return s
}

func (s *GetPoolListResponse) SetErrorMessage(v string) *GetPoolListResponse {
	s.ErrorMessage = &v
	return s
}

func (s *GetPoolListResponse) SetErrorCode(v string) *GetPoolListResponse {
	s.ErrorCode = &v
	return s
}

func (s *GetPoolListResponse) SetState(v GetPoolListResponseStateEnum) *GetPoolListResponse {
	s.State = &v
	return s
}

func (s *GetPoolListResponse) SetBody(v []GetPoolListResponseBody) *GetPoolListResponse {
	s.Body = &v
	return s
}


type GetPoolListResponseBuilder struct {
	s *GetPoolListResponse
}

func NewGetPoolListResponseBuilder() *GetPoolListResponseBuilder {
	s := &GetPoolListResponse{}
	b := &GetPoolListResponseBuilder{s: s}
	return b
}

func (b *GetPoolListResponseBuilder) RequestId(v string) *GetPoolListResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *GetPoolListResponseBuilder) ErrorMessage(v string) *GetPoolListResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *GetPoolListResponseBuilder) ErrorCode(v string) *GetPoolListResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *GetPoolListResponseBuilder) State(v GetPoolListResponseStateEnum) *GetPoolListResponseBuilder {
	b.s.State = &v
	return b
}

func (b *GetPoolListResponseBuilder) Body(v []GetPoolListResponseBody) *GetPoolListResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *GetPoolListResponseBuilder) Build() *GetPoolListResponse {
	return b.s
}


