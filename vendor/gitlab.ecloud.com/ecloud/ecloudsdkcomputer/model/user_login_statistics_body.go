// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UserLoginStatisticsBody struct {
    position.Body
    // 开始时间
	StartTime *string `json:"startTime,omitempty"`
    // 结束时间
	EndTime *string `json:"endTime,omitempty"`
}

func (s UserLoginStatisticsBody) String() string {
	return utils.Beautify(s)
}

func (s UserLoginStatisticsBody) GoString() string {
	return s.String()
}

func (s UserLoginStatisticsBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UserLoginStatisticsBody) SetStartTime(v string) *UserLoginStatisticsBody {
	s.StartTime = &v
	return s
}

func (s *UserLoginStatisticsBody) SetEndTime(v string) *UserLoginStatisticsBody {
	s.EndTime = &v
	return s
}


type UserLoginStatisticsBodyBuilder struct {
	s *UserLoginStatisticsBody
}

func NewUserLoginStatisticsBodyBuilder() *UserLoginStatisticsBodyBuilder {
	s := &UserLoginStatisticsBody{}
	b := &UserLoginStatisticsBodyBuilder{s: s}
	return b
}

func (b *UserLoginStatisticsBodyBuilder) StartTime(v string) *UserLoginStatisticsBodyBuilder {
	b.s.StartTime = &v
	return b
}

func (b *UserLoginStatisticsBodyBuilder) EndTime(v string) *UserLoginStatisticsBodyBuilder {
	b.s.EndTime = &v
	return b
}

func (b *UserLoginStatisticsBodyBuilder) Build() *UserLoginStatisticsBody {
	return b.s
}


