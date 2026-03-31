// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UnsubscribeNatBody struct {
    position.Body
    // 实例id集合,示例值：CCA-XX
	InstanceIds []string `json:"instanceIds,omitempty"`
    // 租户id
	UserId *string `json:"userId,omitempty"`
}

func (s UnsubscribeNatBody) String() string {
	return utils.Beautify(s)
}

func (s UnsubscribeNatBody) GoString() string {
	return s.String()
}

func (s UnsubscribeNatBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UnsubscribeNatBody) SetInstanceIds(v []string) *UnsubscribeNatBody {
	s.InstanceIds = v
	return s
}

func (s *UnsubscribeNatBody) SetUserId(v string) *UnsubscribeNatBody {
	s.UserId = &v
	return s
}


type UnsubscribeNatBodyBuilder struct {
	s *UnsubscribeNatBody
}

func NewUnsubscribeNatBodyBuilder() *UnsubscribeNatBodyBuilder {
	s := &UnsubscribeNatBody{}
	b := &UnsubscribeNatBodyBuilder{s: s}
	return b
}

func (b *UnsubscribeNatBodyBuilder) InstanceIds(v []string) *UnsubscribeNatBodyBuilder {
	b.s.InstanceIds = v
	return b
}

func (b *UnsubscribeNatBodyBuilder) UserId(v string) *UnsubscribeNatBodyBuilder {
	b.s.UserId = &v
	return b
}

func (b *UnsubscribeNatBodyBuilder) Build() *UnsubscribeNatBody {
	return b.s
}


