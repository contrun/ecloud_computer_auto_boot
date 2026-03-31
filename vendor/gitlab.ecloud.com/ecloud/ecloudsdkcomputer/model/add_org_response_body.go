// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type AddOrgResponseBody struct {

    // 组织名称
	OrgName *string `json:"orgName,omitempty"`
    // 组织级别
	Level *int32 `json:"level,omitempty"`
    // 父组织id
	Pid *string `json:"pid,omitempty"`
    // 组织id
	OrgId *string `json:"orgId,omitempty"`
}

func (s AddOrgResponseBody) String() string {
	return utils.Beautify(s)
}

func (s AddOrgResponseBody) GoString() string {
	return s.String()
}

func (s AddOrgResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *AddOrgResponseBody) SetOrgName(v string) *AddOrgResponseBody {
	s.OrgName = &v
	return s
}

func (s *AddOrgResponseBody) SetLevel(v int32) *AddOrgResponseBody {
	s.Level = &v
	return s
}

func (s *AddOrgResponseBody) SetPid(v string) *AddOrgResponseBody {
	s.Pid = &v
	return s
}

func (s *AddOrgResponseBody) SetOrgId(v string) *AddOrgResponseBody {
	s.OrgId = &v
	return s
}


type AddOrgResponseBodyBuilder struct {
	s *AddOrgResponseBody
}

func NewAddOrgResponseBodyBuilder() *AddOrgResponseBodyBuilder {
	s := &AddOrgResponseBody{}
	b := &AddOrgResponseBodyBuilder{s: s}
	return b
}

func (b *AddOrgResponseBodyBuilder) OrgName(v string) *AddOrgResponseBodyBuilder {
	b.s.OrgName = &v
	return b
}

func (b *AddOrgResponseBodyBuilder) Level(v int32) *AddOrgResponseBodyBuilder {
	b.s.Level = &v
	return b
}

func (b *AddOrgResponseBodyBuilder) Pid(v string) *AddOrgResponseBodyBuilder {
	b.s.Pid = &v
	return b
}

func (b *AddOrgResponseBodyBuilder) OrgId(v string) *AddOrgResponseBodyBuilder {
	b.s.OrgId = &v
	return b
}

func (b *AddOrgResponseBodyBuilder) Build() *AddOrgResponseBody {
	return b.s
}


