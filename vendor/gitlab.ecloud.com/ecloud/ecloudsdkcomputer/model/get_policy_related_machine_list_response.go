// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type GetPolicyRelatedMachineListResponseStateEnum string

// List of State
const (
    GetPolicyRelatedMachineListResponseStateEnumOk GetPolicyRelatedMachineListResponseStateEnum = "OK"
    GetPolicyRelatedMachineListResponseStateEnumError GetPolicyRelatedMachineListResponseStateEnum = "ERROR"
    GetPolicyRelatedMachineListResponseStateEnumException GetPolicyRelatedMachineListResponseStateEnum = "EXCEPTION"
    GetPolicyRelatedMachineListResponseStateEnumForbidden GetPolicyRelatedMachineListResponseStateEnum = "FORBIDDEN"
    GetPolicyRelatedMachineListResponseStateEnumRemaining GetPolicyRelatedMachineListResponseStateEnum = "REMAINING"
    GetPolicyRelatedMachineListResponseStateEnumTimeout GetPolicyRelatedMachineListResponseStateEnum = "TIMEOUT"
)

type GetPolicyRelatedMachineListResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *GetPolicyRelatedMachineListResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *GetPolicyRelatedMachineListResponseBody `json:"body,omitempty"`
}

func (s GetPolicyRelatedMachineListResponse) String() string {
	return utils.Beautify(s)
}

func (s GetPolicyRelatedMachineListResponse) GoString() string {
	return s.String()
}

func (s GetPolicyRelatedMachineListResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetPolicyRelatedMachineListResponse) SetRequestId(v string) *GetPolicyRelatedMachineListResponse {
	s.RequestId = &v
	return s
}

func (s *GetPolicyRelatedMachineListResponse) SetErrorMessage(v string) *GetPolicyRelatedMachineListResponse {
	s.ErrorMessage = &v
	return s
}

func (s *GetPolicyRelatedMachineListResponse) SetErrorCode(v string) *GetPolicyRelatedMachineListResponse {
	s.ErrorCode = &v
	return s
}

func (s *GetPolicyRelatedMachineListResponse) SetState(v GetPolicyRelatedMachineListResponseStateEnum) *GetPolicyRelatedMachineListResponse {
	s.State = &v
	return s
}

func (s *GetPolicyRelatedMachineListResponse) SetBody(v *GetPolicyRelatedMachineListResponseBody) *GetPolicyRelatedMachineListResponse {
	s.Body = v
	return s
}


type GetPolicyRelatedMachineListResponseBuilder struct {
	s *GetPolicyRelatedMachineListResponse
}

func NewGetPolicyRelatedMachineListResponseBuilder() *GetPolicyRelatedMachineListResponseBuilder {
	s := &GetPolicyRelatedMachineListResponse{}
	b := &GetPolicyRelatedMachineListResponseBuilder{s: s}
	return b
}

func (b *GetPolicyRelatedMachineListResponseBuilder) RequestId(v string) *GetPolicyRelatedMachineListResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *GetPolicyRelatedMachineListResponseBuilder) ErrorMessage(v string) *GetPolicyRelatedMachineListResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *GetPolicyRelatedMachineListResponseBuilder) ErrorCode(v string) *GetPolicyRelatedMachineListResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *GetPolicyRelatedMachineListResponseBuilder) State(v GetPolicyRelatedMachineListResponseStateEnum) *GetPolicyRelatedMachineListResponseBuilder {
	b.s.State = &v
	return b
}

func (b *GetPolicyRelatedMachineListResponseBuilder) Body(v *GetPolicyRelatedMachineListResponseBody) *GetPolicyRelatedMachineListResponseBuilder {
	b.s.Body = v
	return b
}

func (b *GetPolicyRelatedMachineListResponseBuilder) Build() *GetPolicyRelatedMachineListResponse {
	return b.s
}


