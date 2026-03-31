// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type SecurityGroupManagementListBody struct {
    position.Body
    // 安全组名
	SecurityGroupName *string `json:"securityGroupName,omitempty"`
    // 当前页起始下标
	Start *int32 `json:"start,omitempty"`
    // 每页显示条数
	PageSize *int32 `json:"pageSize,omitempty"`
    // 资源池编码
	PoolCode *string `json:"poolCode,omitempty"`
    // 安全组状态
	SecurityGroupStatus *string `json:"securityGroupStatus,omitempty"`
    // 当前页
	CurrentPage *int32 `json:"currentPage,omitempty"`
    // 安全组实例ID
	SecurityGroupUid *string `json:"securityGroupUid,omitempty"`
    // 所属资源池名称
	PoolName *string `json:"poolName,omitempty"`
}

func (s SecurityGroupManagementListBody) String() string {
	return utils.Beautify(s)
}

func (s SecurityGroupManagementListBody) GoString() string {
	return s.String()
}

func (s SecurityGroupManagementListBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *SecurityGroupManagementListBody) SetSecurityGroupName(v string) *SecurityGroupManagementListBody {
	s.SecurityGroupName = &v
	return s
}

func (s *SecurityGroupManagementListBody) SetStart(v int32) *SecurityGroupManagementListBody {
	s.Start = &v
	return s
}

func (s *SecurityGroupManagementListBody) SetPageSize(v int32) *SecurityGroupManagementListBody {
	s.PageSize = &v
	return s
}

func (s *SecurityGroupManagementListBody) SetPoolCode(v string) *SecurityGroupManagementListBody {
	s.PoolCode = &v
	return s
}

func (s *SecurityGroupManagementListBody) SetSecurityGroupStatus(v string) *SecurityGroupManagementListBody {
	s.SecurityGroupStatus = &v
	return s
}

func (s *SecurityGroupManagementListBody) SetCurrentPage(v int32) *SecurityGroupManagementListBody {
	s.CurrentPage = &v
	return s
}

func (s *SecurityGroupManagementListBody) SetSecurityGroupUid(v string) *SecurityGroupManagementListBody {
	s.SecurityGroupUid = &v
	return s
}

func (s *SecurityGroupManagementListBody) SetPoolName(v string) *SecurityGroupManagementListBody {
	s.PoolName = &v
	return s
}


type SecurityGroupManagementListBodyBuilder struct {
	s *SecurityGroupManagementListBody
}

func NewSecurityGroupManagementListBodyBuilder() *SecurityGroupManagementListBodyBuilder {
	s := &SecurityGroupManagementListBody{}
	b := &SecurityGroupManagementListBodyBuilder{s: s}
	return b
}

func (b *SecurityGroupManagementListBodyBuilder) SecurityGroupName(v string) *SecurityGroupManagementListBodyBuilder {
	b.s.SecurityGroupName = &v
	return b
}

func (b *SecurityGroupManagementListBodyBuilder) Start(v int32) *SecurityGroupManagementListBodyBuilder {
	b.s.Start = &v
	return b
}

func (b *SecurityGroupManagementListBodyBuilder) PageSize(v int32) *SecurityGroupManagementListBodyBuilder {
	b.s.PageSize = &v
	return b
}

func (b *SecurityGroupManagementListBodyBuilder) PoolCode(v string) *SecurityGroupManagementListBodyBuilder {
	b.s.PoolCode = &v
	return b
}

func (b *SecurityGroupManagementListBodyBuilder) SecurityGroupStatus(v string) *SecurityGroupManagementListBodyBuilder {
	b.s.SecurityGroupStatus = &v
	return b
}

func (b *SecurityGroupManagementListBodyBuilder) CurrentPage(v int32) *SecurityGroupManagementListBodyBuilder {
	b.s.CurrentPage = &v
	return b
}

func (b *SecurityGroupManagementListBodyBuilder) SecurityGroupUid(v string) *SecurityGroupManagementListBodyBuilder {
	b.s.SecurityGroupUid = &v
	return b
}

func (b *SecurityGroupManagementListBodyBuilder) PoolName(v string) *SecurityGroupManagementListBodyBuilder {
	b.s.PoolName = &v
	return b
}

func (b *SecurityGroupManagementListBodyBuilder) Build() *SecurityGroupManagementListBody {
	return b.s
}


