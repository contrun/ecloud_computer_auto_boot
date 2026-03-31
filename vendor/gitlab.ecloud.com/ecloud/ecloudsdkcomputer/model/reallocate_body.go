// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ReallocateBody struct {
    position.Body
    // 订购资源实例ID列表
	Instances *[]ReallocateRequestInstances `json:"instances,omitempty"`
    // 云电脑用户id
	UserId *string `json:"userId,omitempty"`
}

func (s ReallocateBody) String() string {
	return utils.Beautify(s)
}

func (s ReallocateBody) GoString() string {
	return s.String()
}

func (s ReallocateBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ReallocateBody) SetInstances(v []ReallocateRequestInstances) *ReallocateBody {
	s.Instances = &v
	return s
}

func (s *ReallocateBody) SetUserId(v string) *ReallocateBody {
	s.UserId = &v
	return s
}


type ReallocateBodyBuilder struct {
	s *ReallocateBody
}

func NewReallocateBodyBuilder() *ReallocateBodyBuilder {
	s := &ReallocateBody{}
	b := &ReallocateBodyBuilder{s: s}
	return b
}

func (b *ReallocateBodyBuilder) Instances(v []ReallocateRequestInstances) *ReallocateBodyBuilder {
	b.s.Instances = &v
	return b
}

func (b *ReallocateBodyBuilder) UserId(v string) *ReallocateBodyBuilder {
	b.s.UserId = &v
	return b
}

func (b *ReallocateBodyBuilder) Build() *ReallocateBody {
	return b.s
}


