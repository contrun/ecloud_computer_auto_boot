// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetOrgListResponseRecords struct {

    // 组织名称
	OrgName *string `json:"orgName,omitempty"`
    // 组织级别
	Level *int32 `json:"level,omitempty"`
    // 父组织id
	Pid *string `json:"pid,omitempty"`
    // 组织id
	OrgId *string `json:"orgId,omitempty"`
}

func (s GetOrgListResponseRecords) String() string {
	return utils.Beautify(s)
}

func (s GetOrgListResponseRecords) GoString() string {
	return s.String()
}

func (s GetOrgListResponseRecords) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetOrgListResponseRecords) SetOrgName(v string) *GetOrgListResponseRecords {
	s.OrgName = &v
	return s
}

func (s *GetOrgListResponseRecords) SetLevel(v int32) *GetOrgListResponseRecords {
	s.Level = &v
	return s
}

func (s *GetOrgListResponseRecords) SetPid(v string) *GetOrgListResponseRecords {
	s.Pid = &v
	return s
}

func (s *GetOrgListResponseRecords) SetOrgId(v string) *GetOrgListResponseRecords {
	s.OrgId = &v
	return s
}


type GetOrgListResponseRecordsBuilder struct {
	s *GetOrgListResponseRecords
}

func NewGetOrgListResponseRecordsBuilder() *GetOrgListResponseRecordsBuilder {
	s := &GetOrgListResponseRecords{}
	b := &GetOrgListResponseRecordsBuilder{s: s}
	return b
}

func (b *GetOrgListResponseRecordsBuilder) OrgName(v string) *GetOrgListResponseRecordsBuilder {
	b.s.OrgName = &v
	return b
}

func (b *GetOrgListResponseRecordsBuilder) Level(v int32) *GetOrgListResponseRecordsBuilder {
	b.s.Level = &v
	return b
}

func (b *GetOrgListResponseRecordsBuilder) Pid(v string) *GetOrgListResponseRecordsBuilder {
	b.s.Pid = &v
	return b
}

func (b *GetOrgListResponseRecordsBuilder) OrgId(v string) *GetOrgListResponseRecordsBuilder {
	b.s.OrgId = &v
	return b
}

func (b *GetOrgListResponseRecordsBuilder) Build() *GetOrgListResponseRecords {
	return b.s
}


