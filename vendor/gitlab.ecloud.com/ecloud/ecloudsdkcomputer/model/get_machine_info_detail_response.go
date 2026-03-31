// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type GetMachineInfoDetailResponseStateEnum string

// List of State
const (
    GetMachineInfoDetailResponseStateEnumOk GetMachineInfoDetailResponseStateEnum = "OK"
    GetMachineInfoDetailResponseStateEnumError GetMachineInfoDetailResponseStateEnum = "ERROR"
    GetMachineInfoDetailResponseStateEnumException GetMachineInfoDetailResponseStateEnum = "EXCEPTION"
    GetMachineInfoDetailResponseStateEnumForbidden GetMachineInfoDetailResponseStateEnum = "FORBIDDEN"
    GetMachineInfoDetailResponseStateEnumRemaining GetMachineInfoDetailResponseStateEnum = "REMAINING"
    GetMachineInfoDetailResponseStateEnumTimeout GetMachineInfoDetailResponseStateEnum = "TIMEOUT"
)

type GetMachineInfoDetailResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *GetMachineInfoDetailResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *GetMachineInfoDetailResponseBody `json:"body,omitempty"`
}

func (s GetMachineInfoDetailResponse) String() string {
	return utils.Beautify(s)
}

func (s GetMachineInfoDetailResponse) GoString() string {
	return s.String()
}

func (s GetMachineInfoDetailResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetMachineInfoDetailResponse) SetRequestId(v string) *GetMachineInfoDetailResponse {
	s.RequestId = &v
	return s
}

func (s *GetMachineInfoDetailResponse) SetErrorMessage(v string) *GetMachineInfoDetailResponse {
	s.ErrorMessage = &v
	return s
}

func (s *GetMachineInfoDetailResponse) SetErrorCode(v string) *GetMachineInfoDetailResponse {
	s.ErrorCode = &v
	return s
}

func (s *GetMachineInfoDetailResponse) SetState(v GetMachineInfoDetailResponseStateEnum) *GetMachineInfoDetailResponse {
	s.State = &v
	return s
}

func (s *GetMachineInfoDetailResponse) SetBody(v *GetMachineInfoDetailResponseBody) *GetMachineInfoDetailResponse {
	s.Body = v
	return s
}


type GetMachineInfoDetailResponseBuilder struct {
	s *GetMachineInfoDetailResponse
}

func NewGetMachineInfoDetailResponseBuilder() *GetMachineInfoDetailResponseBuilder {
	s := &GetMachineInfoDetailResponse{}
	b := &GetMachineInfoDetailResponseBuilder{s: s}
	return b
}

func (b *GetMachineInfoDetailResponseBuilder) RequestId(v string) *GetMachineInfoDetailResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *GetMachineInfoDetailResponseBuilder) ErrorMessage(v string) *GetMachineInfoDetailResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *GetMachineInfoDetailResponseBuilder) ErrorCode(v string) *GetMachineInfoDetailResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *GetMachineInfoDetailResponseBuilder) State(v GetMachineInfoDetailResponseStateEnum) *GetMachineInfoDetailResponseBuilder {
	b.s.State = &v
	return b
}

func (b *GetMachineInfoDetailResponseBuilder) Body(v *GetMachineInfoDetailResponseBody) *GetMachineInfoDetailResponseBuilder {
	b.s.Body = v
	return b
}

func (b *GetMachineInfoDetailResponseBuilder) Build() *GetMachineInfoDetailResponse {
	return b.s
}


