// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type AddOrgBody struct {
    position.Body
    // 组织名称
	Name *string `json:"name,omitempty"`
    // 描述
	Description *string `json:"description,omitempty"`
    // 父组织id
	Pid *string `json:"pid,omitempty"`
}

func (s AddOrgBody) String() string {
	return utils.Beautify(s)
}

func (s AddOrgBody) GoString() string {
	return s.String()
}

func (s AddOrgBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *AddOrgBody) SetName(v string) *AddOrgBody {
	s.Name = &v
	return s
}

func (s *AddOrgBody) SetDescription(v string) *AddOrgBody {
	s.Description = &v
	return s
}

func (s *AddOrgBody) SetPid(v string) *AddOrgBody {
	s.Pid = &v
	return s
}


type AddOrgBodyBuilder struct {
	s *AddOrgBody
}

func NewAddOrgBodyBuilder() *AddOrgBodyBuilder {
	s := &AddOrgBody{}
	b := &AddOrgBodyBuilder{s: s}
	return b
}

func (b *AddOrgBodyBuilder) Name(v string) *AddOrgBodyBuilder {
	b.s.Name = &v
	return b
}

func (b *AddOrgBodyBuilder) Description(v string) *AddOrgBodyBuilder {
	b.s.Description = &v
	return b
}

func (b *AddOrgBodyBuilder) Pid(v string) *AddOrgBodyBuilder {
	b.s.Pid = &v
	return b
}

func (b *AddOrgBodyBuilder) Build() *AddOrgBody {
	return b.s
}


