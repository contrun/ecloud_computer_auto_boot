// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type GetVmListResponseStateEnum string

// List of State
const (
    GetVmListResponseStateEnumOk GetVmListResponseStateEnum = "OK"
    GetVmListResponseStateEnumError GetVmListResponseStateEnum = "ERROR"
    GetVmListResponseStateEnumException GetVmListResponseStateEnum = "EXCEPTION"
    GetVmListResponseStateEnumForbidden GetVmListResponseStateEnum = "FORBIDDEN"
    GetVmListResponseStateEnumRemaining GetVmListResponseStateEnum = "REMAINING"
    GetVmListResponseStateEnumTimeout GetVmListResponseStateEnum = "TIMEOUT"
)

type GetVmListResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *GetVmListResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *GetVmListResponseBody `json:"body,omitempty"`
}

func (s GetVmListResponse) String() string {
	return utils.Beautify(s)
}

func (s GetVmListResponse) GoString() string {
	return s.String()
}

func (s GetVmListResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetVmListResponse) SetRequestId(v string) *GetVmListResponse {
	s.RequestId = &v
	return s
}

func (s *GetVmListResponse) SetErrorMessage(v string) *GetVmListResponse {
	s.ErrorMessage = &v
	return s
}

func (s *GetVmListResponse) SetErrorCode(v string) *GetVmListResponse {
	s.ErrorCode = &v
	return s
}

func (s *GetVmListResponse) SetState(v GetVmListResponseStateEnum) *GetVmListResponse {
	s.State = &v
	return s
}

func (s *GetVmListResponse) SetBody(v *GetVmListResponseBody) *GetVmListResponse {
	s.Body = v
	return s
}


type GetVmListResponseBuilder struct {
	s *GetVmListResponse
}

func NewGetVmListResponseBuilder() *GetVmListResponseBuilder {
	s := &GetVmListResponse{}
	b := &GetVmListResponseBuilder{s: s}
	return b
}

func (b *GetVmListResponseBuilder) RequestId(v string) *GetVmListResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *GetVmListResponseBuilder) ErrorMessage(v string) *GetVmListResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *GetVmListResponseBuilder) ErrorCode(v string) *GetVmListResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *GetVmListResponseBuilder) State(v GetVmListResponseStateEnum) *GetVmListResponseBuilder {
	b.s.State = &v
	return b
}

func (b *GetVmListResponseBuilder) Body(v *GetVmListResponseBody) *GetVmListResponseBuilder {
	b.s.Body = v
	return b
}

func (b *GetVmListResponseBuilder) Build() *GetVmListResponse {
	return b.s
}


