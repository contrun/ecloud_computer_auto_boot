// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type RenameResourceBatchResponseStateEnum string

// List of State
const (
    RenameResourceBatchResponseStateEnumOk RenameResourceBatchResponseStateEnum = "OK"
    RenameResourceBatchResponseStateEnumError RenameResourceBatchResponseStateEnum = "ERROR"
    RenameResourceBatchResponseStateEnumException RenameResourceBatchResponseStateEnum = "EXCEPTION"
    RenameResourceBatchResponseStateEnumForbidden RenameResourceBatchResponseStateEnum = "FORBIDDEN"
    RenameResourceBatchResponseStateEnumRemaining RenameResourceBatchResponseStateEnum = "REMAINING"
    RenameResourceBatchResponseStateEnumTimeout RenameResourceBatchResponseStateEnum = "TIMEOUT"
)

type RenameResourceBatchResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误码提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *RenameResourceBatchResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *bool `json:"body,omitempty"`
}

func (s RenameResourceBatchResponse) String() string {
	return utils.Beautify(s)
}

func (s RenameResourceBatchResponse) GoString() string {
	return s.String()
}

func (s RenameResourceBatchResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *RenameResourceBatchResponse) SetRequestId(v string) *RenameResourceBatchResponse {
	s.RequestId = &v
	return s
}

func (s *RenameResourceBatchResponse) SetErrorMessage(v string) *RenameResourceBatchResponse {
	s.ErrorMessage = &v
	return s
}

func (s *RenameResourceBatchResponse) SetErrorCode(v string) *RenameResourceBatchResponse {
	s.ErrorCode = &v
	return s
}

func (s *RenameResourceBatchResponse) SetState(v RenameResourceBatchResponseStateEnum) *RenameResourceBatchResponse {
	s.State = &v
	return s
}

func (s *RenameResourceBatchResponse) SetBody(v bool) *RenameResourceBatchResponse {
	s.Body = &v
	return s
}


type RenameResourceBatchResponseBuilder struct {
	s *RenameResourceBatchResponse
}

func NewRenameResourceBatchResponseBuilder() *RenameResourceBatchResponseBuilder {
	s := &RenameResourceBatchResponse{}
	b := &RenameResourceBatchResponseBuilder{s: s}
	return b
}

func (b *RenameResourceBatchResponseBuilder) RequestId(v string) *RenameResourceBatchResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *RenameResourceBatchResponseBuilder) ErrorMessage(v string) *RenameResourceBatchResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *RenameResourceBatchResponseBuilder) ErrorCode(v string) *RenameResourceBatchResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *RenameResourceBatchResponseBuilder) State(v RenameResourceBatchResponseStateEnum) *RenameResourceBatchResponseBuilder {
	b.s.State = &v
	return b
}

func (b *RenameResourceBatchResponseBuilder) Body(v bool) *RenameResourceBatchResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *RenameResourceBatchResponseBuilder) Build() *RenameResourceBatchResponse {
	return b.s
}


