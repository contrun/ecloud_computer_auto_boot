// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type ModifyDeviceTrustNumResponseStateEnum string

// List of State
const (
    ModifyDeviceTrustNumResponseStateEnumOk ModifyDeviceTrustNumResponseStateEnum = "OK"
    ModifyDeviceTrustNumResponseStateEnumError ModifyDeviceTrustNumResponseStateEnum = "ERROR"
    ModifyDeviceTrustNumResponseStateEnumException ModifyDeviceTrustNumResponseStateEnum = "EXCEPTION"
    ModifyDeviceTrustNumResponseStateEnumForbidden ModifyDeviceTrustNumResponseStateEnum = "FORBIDDEN"
    ModifyDeviceTrustNumResponseStateEnumRemaining ModifyDeviceTrustNumResponseStateEnum = "REMAINING"
    ModifyDeviceTrustNumResponseStateEnumTimeout ModifyDeviceTrustNumResponseStateEnum = "TIMEOUT"
)

type ModifyDeviceTrustNumResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *ModifyDeviceTrustNumResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *bool `json:"body,omitempty"`
}

func (s ModifyDeviceTrustNumResponse) String() string {
	return utils.Beautify(s)
}

func (s ModifyDeviceTrustNumResponse) GoString() string {
	return s.String()
}

func (s ModifyDeviceTrustNumResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ModifyDeviceTrustNumResponse) SetRequestId(v string) *ModifyDeviceTrustNumResponse {
	s.RequestId = &v
	return s
}

func (s *ModifyDeviceTrustNumResponse) SetErrorMessage(v string) *ModifyDeviceTrustNumResponse {
	s.ErrorMessage = &v
	return s
}

func (s *ModifyDeviceTrustNumResponse) SetErrorCode(v string) *ModifyDeviceTrustNumResponse {
	s.ErrorCode = &v
	return s
}

func (s *ModifyDeviceTrustNumResponse) SetState(v ModifyDeviceTrustNumResponseStateEnum) *ModifyDeviceTrustNumResponse {
	s.State = &v
	return s
}

func (s *ModifyDeviceTrustNumResponse) SetBody(v bool) *ModifyDeviceTrustNumResponse {
	s.Body = &v
	return s
}


type ModifyDeviceTrustNumResponseBuilder struct {
	s *ModifyDeviceTrustNumResponse
}

func NewModifyDeviceTrustNumResponseBuilder() *ModifyDeviceTrustNumResponseBuilder {
	s := &ModifyDeviceTrustNumResponse{}
	b := &ModifyDeviceTrustNumResponseBuilder{s: s}
	return b
}

func (b *ModifyDeviceTrustNumResponseBuilder) RequestId(v string) *ModifyDeviceTrustNumResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *ModifyDeviceTrustNumResponseBuilder) ErrorMessage(v string) *ModifyDeviceTrustNumResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *ModifyDeviceTrustNumResponseBuilder) ErrorCode(v string) *ModifyDeviceTrustNumResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *ModifyDeviceTrustNumResponseBuilder) State(v ModifyDeviceTrustNumResponseStateEnum) *ModifyDeviceTrustNumResponseBuilder {
	b.s.State = &v
	return b
}

func (b *ModifyDeviceTrustNumResponseBuilder) Body(v bool) *ModifyDeviceTrustNumResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *ModifyDeviceTrustNumResponseBuilder) Build() *ModifyDeviceTrustNumResponse {
	return b.s
}


