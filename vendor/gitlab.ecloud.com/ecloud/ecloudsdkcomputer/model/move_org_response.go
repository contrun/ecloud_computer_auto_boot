// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type MoveOrgResponseStateEnum string

// List of State
const (
    MoveOrgResponseStateEnumOk MoveOrgResponseStateEnum = "OK"
    MoveOrgResponseStateEnumError MoveOrgResponseStateEnum = "ERROR"
    MoveOrgResponseStateEnumException MoveOrgResponseStateEnum = "EXCEPTION"
    MoveOrgResponseStateEnumForbidden MoveOrgResponseStateEnum = "FORBIDDEN"
    MoveOrgResponseStateEnumRemaining MoveOrgResponseStateEnum = "REMAINING"
    MoveOrgResponseStateEnumTimeout MoveOrgResponseStateEnum = "TIMEOUT"
)

type MoveOrgResponse struct {

    // 每个请求的序列号
	RequestId *string `json:"requestId,omitempty"`
    // 错误提示信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 统一错误码
	ErrorCode *string `json:"errorCode,omitempty"`
    // 返回状态码;返回正常：OK,返回错误：ERROR,返回异常：EXCEPTION,需要告警实现的：ALARM,禁止访问：FORBIDDEN
	State *MoveOrgResponseStateEnum `json:"state,omitempty"`
    // 请求成功时返回的数据
	Body *string `json:"body,omitempty"`
}

func (s MoveOrgResponse) String() string {
	return utils.Beautify(s)
}

func (s MoveOrgResponse) GoString() string {
	return s.String()
}

func (s MoveOrgResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *MoveOrgResponse) SetRequestId(v string) *MoveOrgResponse {
	s.RequestId = &v
	return s
}

func (s *MoveOrgResponse) SetErrorMessage(v string) *MoveOrgResponse {
	s.ErrorMessage = &v
	return s
}

func (s *MoveOrgResponse) SetErrorCode(v string) *MoveOrgResponse {
	s.ErrorCode = &v
	return s
}

func (s *MoveOrgResponse) SetState(v MoveOrgResponseStateEnum) *MoveOrgResponse {
	s.State = &v
	return s
}

func (s *MoveOrgResponse) SetBody(v string) *MoveOrgResponse {
	s.Body = &v
	return s
}


type MoveOrgResponseBuilder struct {
	s *MoveOrgResponse
}

func NewMoveOrgResponseBuilder() *MoveOrgResponseBuilder {
	s := &MoveOrgResponse{}
	b := &MoveOrgResponseBuilder{s: s}
	return b
}

func (b *MoveOrgResponseBuilder) RequestId(v string) *MoveOrgResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *MoveOrgResponseBuilder) ErrorMessage(v string) *MoveOrgResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *MoveOrgResponseBuilder) ErrorCode(v string) *MoveOrgResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *MoveOrgResponseBuilder) State(v MoveOrgResponseStateEnum) *MoveOrgResponseBuilder {
	b.s.State = &v
	return b
}

func (b *MoveOrgResponseBuilder) Body(v string) *MoveOrgResponseBuilder {
	b.s.Body = &v
	return b
}

func (b *MoveOrgResponseBuilder) Build() *MoveOrgResponse {
	return b.s
}


