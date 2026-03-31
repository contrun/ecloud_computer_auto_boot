// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type GetResourceListResponseStateEnum string

// List of State
const (
    GetResourceListResponseStateEnumOk GetResourceListResponseStateEnum = "OK"
    GetResourceListResponseStateEnumError GetResourceListResponseStateEnum = "ERROR"
    GetResourceListResponseStateEnumException GetResourceListResponseStateEnum = "EXCEPTION"
    GetResourceListResponseStateEnumForbidden GetResourceListResponseStateEnum = "FORBIDDEN"
    GetResourceListResponseStateEnumRemaining GetResourceListResponseStateEnum = "REMAINING"
    GetResourceListResponseStateEnumTimeout GetResourceListResponseStateEnum = "TIMEOUT"
)

type GetResourceListResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *GetResourceListResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *GetResourceListResponseBody `json:"body,omitempty"`
}

func (s GetResourceListResponse) String() string {
	return utils.Beautify(s)
}

func (s GetResourceListResponse) GoString() string {
	return s.String()
}

func (s GetResourceListResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetResourceListResponse) SetRequestId(v string) *GetResourceListResponse {
	s.RequestId = &v
	return s
}

func (s *GetResourceListResponse) SetErrorMessage(v string) *GetResourceListResponse {
	s.ErrorMessage = &v
	return s
}

func (s *GetResourceListResponse) SetErrorCode(v string) *GetResourceListResponse {
	s.ErrorCode = &v
	return s
}

func (s *GetResourceListResponse) SetState(v GetResourceListResponseStateEnum) *GetResourceListResponse {
	s.State = &v
	return s
}

func (s *GetResourceListResponse) SetBody(v *GetResourceListResponseBody) *GetResourceListResponse {
	s.Body = v
	return s
}


type GetResourceListResponseBuilder struct {
	s *GetResourceListResponse
}

func NewGetResourceListResponseBuilder() *GetResourceListResponseBuilder {
	s := &GetResourceListResponse{}
	b := &GetResourceListResponseBuilder{s: s}
	return b
}

func (b *GetResourceListResponseBuilder) RequestId(v string) *GetResourceListResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *GetResourceListResponseBuilder) ErrorMessage(v string) *GetResourceListResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *GetResourceListResponseBuilder) ErrorCode(v string) *GetResourceListResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *GetResourceListResponseBuilder) State(v GetResourceListResponseStateEnum) *GetResourceListResponseBuilder {
	b.s.State = &v
	return b
}

func (b *GetResourceListResponseBuilder) Body(v *GetResourceListResponseBody) *GetResourceListResponseBuilder {
	b.s.Body = v
	return b
}

func (b *GetResourceListResponseBuilder) Build() *GetResourceListResponse {
	return b.s
}


