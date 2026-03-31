// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CancelInstanceListBody struct {
    position.Body
    // 实例id集合，示例值：CCA-XX
	InstanceIds []string `json:"instanceIds,omitempty"`
    // 租户id
	UserId *string `json:"userId,omitempty"`
}

func (s CancelInstanceListBody) String() string {
	return utils.Beautify(s)
}

func (s CancelInstanceListBody) GoString() string {
	return s.String()
}

func (s CancelInstanceListBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CancelInstanceListBody) SetInstanceIds(v []string) *CancelInstanceListBody {
	s.InstanceIds = v
	return s
}

func (s *CancelInstanceListBody) SetUserId(v string) *CancelInstanceListBody {
	s.UserId = &v
	return s
}


type CancelInstanceListBodyBuilder struct {
	s *CancelInstanceListBody
}

func NewCancelInstanceListBodyBuilder() *CancelInstanceListBodyBuilder {
	s := &CancelInstanceListBody{}
	b := &CancelInstanceListBodyBuilder{s: s}
	return b
}

func (b *CancelInstanceListBodyBuilder) InstanceIds(v []string) *CancelInstanceListBodyBuilder {
	b.s.InstanceIds = v
	return b
}

func (b *CancelInstanceListBodyBuilder) UserId(v string) *CancelInstanceListBodyBuilder {
	b.s.UserId = &v
	return b
}

func (b *CancelInstanceListBodyBuilder) Build() *CancelInstanceListBody {
	return b.s
}


