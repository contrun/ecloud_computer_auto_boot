// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type EditOrgBody struct {
    position.Body
    // 组织名称
	Name *string `json:"name,omitempty"`
    // 描述
	Description *string `json:"description,omitempty"`
    // 组织id
	Id *string `json:"id,omitempty"`
}

func (s EditOrgBody) String() string {
	return utils.Beautify(s)
}

func (s EditOrgBody) GoString() string {
	return s.String()
}

func (s EditOrgBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *EditOrgBody) SetName(v string) *EditOrgBody {
	s.Name = &v
	return s
}

func (s *EditOrgBody) SetDescription(v string) *EditOrgBody {
	s.Description = &v
	return s
}

func (s *EditOrgBody) SetId(v string) *EditOrgBody {
	s.Id = &v
	return s
}


type EditOrgBodyBuilder struct {
	s *EditOrgBody
}

func NewEditOrgBodyBuilder() *EditOrgBodyBuilder {
	s := &EditOrgBody{}
	b := &EditOrgBodyBuilder{s: s}
	return b
}

func (b *EditOrgBodyBuilder) Name(v string) *EditOrgBodyBuilder {
	b.s.Name = &v
	return b
}

func (b *EditOrgBodyBuilder) Description(v string) *EditOrgBodyBuilder {
	b.s.Description = &v
	return b
}

func (b *EditOrgBodyBuilder) Id(v string) *EditOrgBodyBuilder {
	b.s.Id = &v
	return b
}

func (b *EditOrgBodyBuilder) Build() *EditOrgBody {
	return b.s
}


