// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)
type LoginUsageStatisticsPageResponseStateEnum string

// List of State
const (
    LoginUsageStatisticsPageResponseStateEnumOk LoginUsageStatisticsPageResponseStateEnum = "OK"
    LoginUsageStatisticsPageResponseStateEnumError LoginUsageStatisticsPageResponseStateEnum = "ERROR"
    LoginUsageStatisticsPageResponseStateEnumException LoginUsageStatisticsPageResponseStateEnum = "EXCEPTION"
    LoginUsageStatisticsPageResponseStateEnumForbidden LoginUsageStatisticsPageResponseStateEnum = "FORBIDDEN"
    LoginUsageStatisticsPageResponseStateEnumRemaining LoginUsageStatisticsPageResponseStateEnum = "REMAINING"
    LoginUsageStatisticsPageResponseStateEnumTimeout LoginUsageStatisticsPageResponseStateEnum = "TIMEOUT"
)

type LoginUsageStatisticsPageResponse struct {


	RequestId *string `json:"requestId,omitempty"`

	ErrorMessage *string `json:"errorMessage,omitempty"`

	ErrorCode *string `json:"errorCode,omitempty"`

	State *LoginUsageStatisticsPageResponseStateEnum `json:"state,omitempty"`

	Body *LoginUsageStatisticsPageResponseBody `json:"body,omitempty"`
}

func (s LoginUsageStatisticsPageResponse) String() string {
	return utils.Beautify(s)
}

func (s LoginUsageStatisticsPageResponse) GoString() string {
	return s.String()
}

func (s LoginUsageStatisticsPageResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *LoginUsageStatisticsPageResponse) SetRequestId(v string) *LoginUsageStatisticsPageResponse {
	s.RequestId = &v
	return s
}

func (s *LoginUsageStatisticsPageResponse) SetErrorMessage(v string) *LoginUsageStatisticsPageResponse {
	s.ErrorMessage = &v
	return s
}

func (s *LoginUsageStatisticsPageResponse) SetErrorCode(v string) *LoginUsageStatisticsPageResponse {
	s.ErrorCode = &v
	return s
}

func (s *LoginUsageStatisticsPageResponse) SetState(v LoginUsageStatisticsPageResponseStateEnum) *LoginUsageStatisticsPageResponse {
	s.State = &v
	return s
}

func (s *LoginUsageStatisticsPageResponse) SetBody(v *LoginUsageStatisticsPageResponseBody) *LoginUsageStatisticsPageResponse {
	s.Body = v
	return s
}


type LoginUsageStatisticsPageResponseBuilder struct {
	s *LoginUsageStatisticsPageResponse
}

func NewLoginUsageStatisticsPageResponseBuilder() *LoginUsageStatisticsPageResponseBuilder {
	s := &LoginUsageStatisticsPageResponse{}
	b := &LoginUsageStatisticsPageResponseBuilder{s: s}
	return b
}

func (b *LoginUsageStatisticsPageResponseBuilder) RequestId(v string) *LoginUsageStatisticsPageResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *LoginUsageStatisticsPageResponseBuilder) ErrorMessage(v string) *LoginUsageStatisticsPageResponseBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *LoginUsageStatisticsPageResponseBuilder) ErrorCode(v string) *LoginUsageStatisticsPageResponseBuilder {
	b.s.ErrorCode = &v
	return b
}

func (b *LoginUsageStatisticsPageResponseBuilder) State(v LoginUsageStatisticsPageResponseStateEnum) *LoginUsageStatisticsPageResponseBuilder {
	b.s.State = &v
	return b
}

func (b *LoginUsageStatisticsPageResponseBuilder) Body(v *LoginUsageStatisticsPageResponseBody) *LoginUsageStatisticsPageResponseBuilder {
	b.s.Body = v
	return b
}

func (b *LoginUsageStatisticsPageResponseBuilder) Build() *LoginUsageStatisticsPageResponse {
	return b.s
}


