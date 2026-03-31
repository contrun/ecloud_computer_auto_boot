// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type EditUserPolicyRequestLockTimes struct {

    // 开始时间
	BeginTime *string `json:"beginTime,omitempty"`
    // 结束时间
	EndTime *string `json:"endTime,omitempty"`
}

func (s EditUserPolicyRequestLockTimes) String() string {
	return utils.Beautify(s)
}

func (s EditUserPolicyRequestLockTimes) GoString() string {
	return s.String()
}

func (s EditUserPolicyRequestLockTimes) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *EditUserPolicyRequestLockTimes) SetBeginTime(v string) *EditUserPolicyRequestLockTimes {
	s.BeginTime = &v
	return s
}

func (s *EditUserPolicyRequestLockTimes) SetEndTime(v string) *EditUserPolicyRequestLockTimes {
	s.EndTime = &v
	return s
}


type EditUserPolicyRequestLockTimesBuilder struct {
	s *EditUserPolicyRequestLockTimes
}

func NewEditUserPolicyRequestLockTimesBuilder() *EditUserPolicyRequestLockTimesBuilder {
	s := &EditUserPolicyRequestLockTimes{}
	b := &EditUserPolicyRequestLockTimesBuilder{s: s}
	return b
}

func (b *EditUserPolicyRequestLockTimesBuilder) BeginTime(v string) *EditUserPolicyRequestLockTimesBuilder {
	b.s.BeginTime = &v
	return b
}

func (b *EditUserPolicyRequestLockTimesBuilder) EndTime(v string) *EditUserPolicyRequestLockTimesBuilder {
	b.s.EndTime = &v
	return b
}

func (b *EditUserPolicyRequestLockTimesBuilder) Build() *EditUserPolicyRequestLockTimes {
	return b.s
}


