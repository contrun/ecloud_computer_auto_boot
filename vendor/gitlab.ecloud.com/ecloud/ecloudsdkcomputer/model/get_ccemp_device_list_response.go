// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type GetCcempDeviceListResponseStateEnum string

// List of State
const (
    GetCcempDeviceListResponseStateEnumOk GetCcempDeviceListResponseStateEnum = "OK"
    GetCcempDeviceListResponseStateEnumError GetCcempDeviceListResponseStateEnum = "ERROR"
    GetCcempDeviceListResponseStateEnumException GetCcempDeviceListResponseStateEnum = "EXCEPTION"
    GetCcempDeviceListResponseStateEnumForbidden GetCcempDeviceListResponseStateEnum = "FORBIDDEN"
    GetCcempDeviceListResponseStateEnumRemaining GetCcempDeviceListResponseStateEnum = "REMAINING"
    GetCcempDeviceListResponseStateEnumTimeout GetCcempDeviceListResponseStateEnum = "TIMEOUT"
)

type GetCcempDeviceListResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *GetCcempDeviceListResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *GetCcempDeviceListResponseBody `json:"body,omitempty"`
}

func (s GetCcempDeviceListResponse) String() string {
	return utils.Beautify(s)
}

func (s GetCcempDeviceListResponse) GoString() string {
	return s.String()
}

func (s GetCcempDeviceListResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetCcempDeviceListResponse) SetRequestId(v string) *GetCcempDeviceListResponse {
	s.RequestId = &v
	return s
}

func (s *GetCcempDeviceListResponse) SetErrorMessage(v string) *GetCcempDeviceListResponse {
	s.ErrorMessage = &v
	return s
}

func (s *GetCcempDeviceListResponse) SetErrorCode(v string) *GetCcempDeviceListResponse {
	s.ErrorCode = &v
	return s
}

func (s *GetCcempDeviceListResponse) SetState(v GetCcempDeviceListResponseStateEnum) *GetCcempDeviceListResponse {
	s.State = &v
	return s
}

func (s *GetCcempDeviceListResponse) SetBody(v *GetCcempDeviceListResponseBody) *GetCcempDeviceListResponse {
	s.Body = v
	return s
}


type GetCcempDeviceListResponseBuilder struct {
	s *GetCcempDeviceListResponse
}

func NewGetCcempDeviceListResponseBuilder() *GetCcempDeviceListResponseBuilder {
	s := &GetCcempDeviceListResponse{}
	b := &GetCcempDeviceListResponseBuilder{s: s}
	return b
}

func (b *GetCcempDeviceListResponseBuilder) RequestId(v string) *GetCcempDeviceListResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *GetCcempDeviceListResponseBuilder) ErrorMessage(v string) *GetCcempDeviceListResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *GetCcempDeviceListResponseBuilder) ErrorCode(v string) *GetCcempDeviceListResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *GetCcempDeviceListResponseBuilder) State(v GetCcempDeviceListResponseStateEnum) *GetCcempDeviceListResponseBuilder {
	b.s.State = &v
	return b
}

func (b *GetCcempDeviceListResponseBuilder) Body(v *GetCcempDeviceListResponseBody) *GetCcempDeviceListResponseBuilder {
	b.s.Body = v
	return b
}

func (b *GetCcempDeviceListResponseBuilder) Build() *GetCcempDeviceListResponse {
	return b.s
}


