// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type GetResourceDiskInfoResponseStateEnum string

// List of State
const (
    GetResourceDiskInfoResponseStateEnumOk GetResourceDiskInfoResponseStateEnum = "OK"
    GetResourceDiskInfoResponseStateEnumError GetResourceDiskInfoResponseStateEnum = "ERROR"
    GetResourceDiskInfoResponseStateEnumException GetResourceDiskInfoResponseStateEnum = "EXCEPTION"
    GetResourceDiskInfoResponseStateEnumForbidden GetResourceDiskInfoResponseStateEnum = "FORBIDDEN"
    GetResourceDiskInfoResponseStateEnumRemaining GetResourceDiskInfoResponseStateEnum = "REMAINING"
    GetResourceDiskInfoResponseStateEnumTimeout GetResourceDiskInfoResponseStateEnum = "TIMEOUT"
)

type GetResourceDiskInfoResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *GetResourceDiskInfoResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *[]GetResourceDiskInfoResponseBody `json:"body,omitempty"`
}

func (s GetResourceDiskInfoResponse) String() string {
	return utils.Beautify(s)
}

func (s GetResourceDiskInfoResponse) GoString() string {
	return s.String()
}

func (s GetResourceDiskInfoResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetResourceDiskInfoResponse) SetRequestId(v string) *GetResourceDiskInfoResponse {
	s.RequestId = &v
	return s
}

func (s *GetResourceDiskInfoResponse) SetErrorMessage(v string) *GetResourceDiskInfoResponse {
	s.ErrorMessage = &v
	return s
}

func (s *GetResourceDiskInfoResponse) SetErrorCode(v string) *GetResourceDiskInfoResponse {
	s.ErrorCode = &v
	return s
}

func (s *GetResourceDiskInfoResponse) SetState(v GetResourceDiskInfoResponseStateEnum) *GetResourceDiskInfoResponse {
	s.State = &v
	return s
}

func (s *GetResourceDiskInfoResponse) SetBody(v []GetResourceDiskInfoResponseBody) *GetResourceDiskInfoResponse {
	s.Body = &v
	return s
}


type GetResourceDiskInfoResponseBuilder struct {
	s *GetResourceDiskInfoResponse
}

func NewGetResourceDiskInfoResponseBuilder() *GetResourceDiskInfoResponseBuilder {
	s := &GetResourceDiskInfoResponse{}
	b := &GetResourceDiskInfoResponseBuilder{s: s}
	return b
}

func (b *GetResourceDiskInfoResponseBuilder) RequestId(v string) *GetResourceDiskInfoResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *GetResourceDiskInfoResponseBuilder) ErrorMessage(v string) *GetResourceDiskInfoResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *GetResourceDiskInfoResponseBuilder) ErrorCode(v string) *GetResourceDiskInfoResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *GetResourceDiskInfoResponseBuilder) State(v GetResourceDiskInfoResponseStateEnum) *GetResourceDiskInfoResponseBuilder {
	b.s.State = &v
	return b
}

func (b *GetResourceDiskInfoResponseBuilder) Body(v []GetResourceDiskInfoResponseBody) *GetResourceDiskInfoResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *GetResourceDiskInfoResponseBuilder) Build() *GetResourceDiskInfoResponse {
	return b.s
}


