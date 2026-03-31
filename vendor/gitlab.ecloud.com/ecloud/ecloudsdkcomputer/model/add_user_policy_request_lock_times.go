// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type AddUserPolicyRequestLockTimes struct {

    // 开始时间
	BeginTime *string `json:"beginTime,omitempty"`
    // 结束时间
	EndTime *string `json:"endTime,omitempty"`
}

func (s AddUserPolicyRequestLockTimes) String() string {
	return utils.Beautify(s)
}

func (s AddUserPolicyRequestLockTimes) GoString() string {
	return s.String()
}

func (s AddUserPolicyRequestLockTimes) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *AddUserPolicyRequestLockTimes) SetBeginTime(v string) *AddUserPolicyRequestLockTimes {
	s.BeginTime = &v
	return s
}

func (s *AddUserPolicyRequestLockTimes) SetEndTime(v string) *AddUserPolicyRequestLockTimes {
	s.EndTime = &v
	return s
}


type AddUserPolicyRequestLockTimesBuilder struct {
	s *AddUserPolicyRequestLockTimes
}

func NewAddUserPolicyRequestLockTimesBuilder() *AddUserPolicyRequestLockTimesBuilder {
	s := &AddUserPolicyRequestLockTimes{}
	b := &AddUserPolicyRequestLockTimesBuilder{s: s}
	return b
}

func (b *AddUserPolicyRequestLockTimesBuilder) BeginTime(v string) *AddUserPolicyRequestLockTimesBuilder {
	b.s.BeginTime = &v
	return b
}

func (b *AddUserPolicyRequestLockTimesBuilder) EndTime(v string) *AddUserPolicyRequestLockTimesBuilder {
	b.s.EndTime = &v
	return b
}

func (b *AddUserPolicyRequestLockTimesBuilder) Build() *AddUserPolicyRequestLockTimes {
	return b.s
}


