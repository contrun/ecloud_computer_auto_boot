// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type ConnectUsageStatisticsPageResponseStateEnum string

// List of State
const (
    ConnectUsageStatisticsPageResponseStateEnumOk ConnectUsageStatisticsPageResponseStateEnum = "OK"
    ConnectUsageStatisticsPageResponseStateEnumError ConnectUsageStatisticsPageResponseStateEnum = "ERROR"
    ConnectUsageStatisticsPageResponseStateEnumException ConnectUsageStatisticsPageResponseStateEnum = "EXCEPTION"
    ConnectUsageStatisticsPageResponseStateEnumForbidden ConnectUsageStatisticsPageResponseStateEnum = "FORBIDDEN"
    ConnectUsageStatisticsPageResponseStateEnumRemaining ConnectUsageStatisticsPageResponseStateEnum = "REMAINING"
    ConnectUsageStatisticsPageResponseStateEnumTimeout ConnectUsageStatisticsPageResponseStateEnum = "TIMEOUT"
)

type ConnectUsageStatisticsPageResponse struct {


	RequestId *string `json:"requestId,omitempty"`

	ErrorMessage *string `json:"errorMessage,omitempty"`

	ErrorCode *string `json:"errorCode,omitempty"`

	State *ConnectUsageStatisticsPageResponseStateEnum `json:"state,omitempty"`

	Body *ConnectUsageStatisticsPageResponseBody `json:"body,omitempty"`
}

func (s ConnectUsageStatisticsPageResponse) String() string {
	return utils.Beautify(s)
}

func (s ConnectUsageStatisticsPageResponse) GoString() string {
	return s.String()
}

func (s ConnectUsageStatisticsPageResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ConnectUsageStatisticsPageResponse) SetRequestId(v string) *ConnectUsageStatisticsPageResponse {
	s.RequestId = &v
	return s
}

func (s *ConnectUsageStatisticsPageResponse) SetErrorMessage(v string) *ConnectUsageStatisticsPageResponse {
	s.ErrorMessage = &v
	return s
}

func (s *ConnectUsageStatisticsPageResponse) SetErrorCode(v string) *ConnectUsageStatisticsPageResponse {
	s.ErrorCode = &v
	return s
}

func (s *ConnectUsageStatisticsPageResponse) SetState(v ConnectUsageStatisticsPageResponseStateEnum) *ConnectUsageStatisticsPageResponse {
	s.State = &v
	return s
}

func (s *ConnectUsageStatisticsPageResponse) SetBody(v *ConnectUsageStatisticsPageResponseBody) *ConnectUsageStatisticsPageResponse {
	s.Body = v
	return s
}


type ConnectUsageStatisticsPageResponseBuilder struct {
	s *ConnectUsageStatisticsPageResponse
}

func NewConnectUsageStatisticsPageResponseBuilder() *ConnectUsageStatisticsPageResponseBuilder {
	s := &ConnectUsageStatisticsPageResponse{}
	b := &ConnectUsageStatisticsPageResponseBuilder{s: s}
	return b
}

func (b *ConnectUsageStatisticsPageResponseBuilder) RequestId(v string) *ConnectUsageStatisticsPageResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *ConnectUsageStatisticsPageResponseBuilder) ErrorMessage(v string) *ConnectUsageStatisticsPageResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *ConnectUsageStatisticsPageResponseBuilder) ErrorCode(v string) *ConnectUsageStatisticsPageResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *ConnectUsageStatisticsPageResponseBuilder) State(v ConnectUsageStatisticsPageResponseStateEnum) *ConnectUsageStatisticsPageResponseBuilder {
	b.s.State = &v
	return b
}

func (b *ConnectUsageStatisticsPageResponseBuilder) Body(v *ConnectUsageStatisticsPageResponseBody) *ConnectUsageStatisticsPageResponseBuilder {
	b.s.Body = v
	return b
}

func (b *ConnectUsageStatisticsPageResponseBuilder) Build() *ConnectUsageStatisticsPageResponse {
	return b.s
}


