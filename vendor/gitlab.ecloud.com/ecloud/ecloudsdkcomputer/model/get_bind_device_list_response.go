// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type GetBindDeviceListResponseStateEnum string

// List of State
const (
    GetBindDeviceListResponseStateEnumOk GetBindDeviceListResponseStateEnum = "OK"
    GetBindDeviceListResponseStateEnumError GetBindDeviceListResponseStateEnum = "ERROR"
    GetBindDeviceListResponseStateEnumException GetBindDeviceListResponseStateEnum = "EXCEPTION"
    GetBindDeviceListResponseStateEnumForbidden GetBindDeviceListResponseStateEnum = "FORBIDDEN"
    GetBindDeviceListResponseStateEnumRemaining GetBindDeviceListResponseStateEnum = "REMAINING"
    GetBindDeviceListResponseStateEnumTimeout GetBindDeviceListResponseStateEnum = "TIMEOUT"
)

type GetBindDeviceListResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *GetBindDeviceListResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *GetBindDeviceListResponseBody `json:"body,omitempty"`
}

func (s GetBindDeviceListResponse) String() string {
	return utils.Beautify(s)
}

func (s GetBindDeviceListResponse) GoString() string {
	return s.String()
}

func (s GetBindDeviceListResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetBindDeviceListResponse) SetRequestId(v string) *GetBindDeviceListResponse {
	s.RequestId = &v
	return s
}

func (s *GetBindDeviceListResponse) SetErrorMessage(v string) *GetBindDeviceListResponse {
	s.ErrorMessage = &v
	return s
}

func (s *GetBindDeviceListResponse) SetErrorCode(v string) *GetBindDeviceListResponse {
	s.ErrorCode = &v
	return s
}

func (s *GetBindDeviceListResponse) SetState(v GetBindDeviceListResponseStateEnum) *GetBindDeviceListResponse {
	s.State = &v
	return s
}

func (s *GetBindDeviceListResponse) SetBody(v *GetBindDeviceListResponseBody) *GetBindDeviceListResponse {
	s.Body = v
	return s
}


type GetBindDeviceListResponseBuilder struct {
	s *GetBindDeviceListResponse
}

func NewGetBindDeviceListResponseBuilder() *GetBindDeviceListResponseBuilder {
	s := &GetBindDeviceListResponse{}
	b := &GetBindDeviceListResponseBuilder{s: s}
	return b
}

func (b *GetBindDeviceListResponseBuilder) RequestId(v string) *GetBindDeviceListResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *GetBindDeviceListResponseBuilder) ErrorMessage(v string) *GetBindDeviceListResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *GetBindDeviceListResponseBuilder) ErrorCode(v string) *GetBindDeviceListResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *GetBindDeviceListResponseBuilder) State(v GetBindDeviceListResponseStateEnum) *GetBindDeviceListResponseBuilder {
	b.s.State = &v
	return b
}

func (b *GetBindDeviceListResponseBuilder) Body(v *GetBindDeviceListResponseBody) *GetBindDeviceListResponseBuilder {
	b.s.Body = v
	return b
}

func (b *GetBindDeviceListResponseBuilder) Build() *GetBindDeviceListResponse {
	return b.s
}


