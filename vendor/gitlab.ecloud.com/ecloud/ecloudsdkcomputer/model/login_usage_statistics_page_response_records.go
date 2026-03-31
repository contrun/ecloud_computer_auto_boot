// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type LoginUsageStatisticsPageResponseRecords struct {

    // 厂商
	CompanyCode *string `json:"companyCode,omitempty"`
    // 组织名称
	OrgName *string `json:"orgName,omitempty"`
    // 最近一次登录时间
	LatestLoginTime *string `json:"latestLoginTime,omitempty"`
    // 厂商名称
	CompanyName *string `json:"companyName,omitempty"`
    // 登录账号
	UserName *string `json:"userName,omitempty"`
    // 用户Id
	UserId *string `json:"userId,omitempty"`
    // 登录次数
	LoginCount *int64 `json:"loginCount,omitempty"`
    // 最长登录时长
	LongestLoginTime *string `json:"longestLoginTime,omitempty"`
    // 组织id
	OrgId *string `json:"orgId,omitempty"`
    // 租户id
	TenantId *string `json:"tenantId,omitempty"`
    // 平均登录时长
	AvgLoginTime *string `json:"avgLoginTime,omitempty"`
    // 所属租户（展示租户名称字段）
	MopUserName *string `json:"mopUserName,omitempty"`
    // 租户Id
	MopUserId *string `json:"mopUserId,omitempty"`
    // 登录总时长
	TotalLoginTime *string `json:"totalLoginTime,omitempty"`
}

func (s LoginUsageStatisticsPageResponseRecords) String() string {
	return utils.Beautify(s)
}

func (s LoginUsageStatisticsPageResponseRecords) GoString() string {
	return s.String()
}

func (s LoginUsageStatisticsPageResponseRecords) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *LoginUsageStatisticsPageResponseRecords) SetCompanyCode(v string) *LoginUsageStatisticsPageResponseRecords {
	s.CompanyCode = &v
	return s
}

func (s *LoginUsageStatisticsPageResponseRecords) SetOrgName(v string) *LoginUsageStatisticsPageResponseRecords {
	s.OrgName = &v
	return s
}

func (s *LoginUsageStatisticsPageResponseRecords) SetLatestLoginTime(v string) *LoginUsageStatisticsPageResponseRecords {
	s.LatestLoginTime = &v
	return s
}

func (s *LoginUsageStatisticsPageResponseRecords) SetCompanyName(v string) *LoginUsageStatisticsPageResponseRecords {
	s.CompanyName = &v
	return s
}

func (s *LoginUsageStatisticsPageResponseRecords) SetUserName(v string) *LoginUsageStatisticsPageResponseRecords {
	s.UserName = &v
	return s
}

func (s *LoginUsageStatisticsPageResponseRecords) SetUserId(v string) *LoginUsageStatisticsPageResponseRecords {
	s.UserId = &v
	return s
}

func (s *LoginUsageStatisticsPageResponseRecords) SetLoginCount(v int64) *LoginUsageStatisticsPageResponseRecords {
	s.LoginCount = &v
	return s
}

func (s *LoginUsageStatisticsPageResponseRecords) SetLongestLoginTime(v string) *LoginUsageStatisticsPageResponseRecords {
	s.LongestLoginTime = &v
	return s
}

func (s *LoginUsageStatisticsPageResponseRecords) SetOrgId(v string) *LoginUsageStatisticsPageResponseRecords {
	s.OrgId = &v
	return s
}

func (s *LoginUsageStatisticsPageResponseRecords) SetTenantId(v string) *LoginUsageStatisticsPageResponseRecords {
	s.TenantId = &v
	return s
}

func (s *LoginUsageStatisticsPageResponseRecords) SetAvgLoginTime(v string) *LoginUsageStatisticsPageResponseRecords {
	s.AvgLoginTime = &v
	return s
}

func (s *LoginUsageStatisticsPageResponseRecords) SetMopUserName(v string) *LoginUsageStatisticsPageResponseRecords {
	s.MopUserName = &v
	return s
}

func (s *LoginUsageStatisticsPageResponseRecords) SetMopUserId(v string) *LoginUsageStatisticsPageResponseRecords {
	s.MopUserId = &v
	return s
}

func (s *LoginUsageStatisticsPageResponseRecords) SetTotalLoginTime(v string) *LoginUsageStatisticsPageResponseRecords {
	s.TotalLoginTime = &v
	return s
}


type LoginUsageStatisticsPageResponseRecordsBuilder struct {
	s *LoginUsageStatisticsPageResponseRecords
}

func NewLoginUsageStatisticsPageResponseRecordsBuilder() *LoginUsageStatisticsPageResponseRecordsBuilder {
	s := &LoginUsageStatisticsPageResponseRecords{}
	b := &LoginUsageStatisticsPageResponseRecordsBuilder{s: s}
	return b
}

func (b *LoginUsageStatisticsPageResponseRecordsBuilder) CompanyCode(v string) *LoginUsageStatisticsPageResponseRecordsBuilder {
	b.s.CompanyCode = &v
	return b
}

func (b *LoginUsageStatisticsPageResponseRecordsBuilder) OrgName(v string) *LoginUsageStatisticsPageResponseRecordsBuilder {
	b.s.OrgName = &v
	return b
}

func (b *LoginUsageStatisticsPageResponseRecordsBuilder) LatestLoginTime(v string) *LoginUsageStatisticsPageResponseRecordsBuilder {
	b.s.LatestLoginTime = &v
	return b
}

func (b *LoginUsageStatisticsPageResponseRecordsBuilder) CompanyName(v string) *LoginUsageStatisticsPageResponseRecordsBuilder {
	b.s.CompanyName = &v
	return b
}

func (b *LoginUsageStatisticsPageResponseRecordsBuilder) UserName(v string) *LoginUsageStatisticsPageResponseRecordsBuilder {
	b.s.UserName = &v
	return b
}

func (b *LoginUsageStatisticsPageResponseRecordsBuilder) UserId(v string) *LoginUsageStatisticsPageResponseRecordsBuilder {
	b.s.UserId = &v
	return b
}

func (b *LoginUsageStatisticsPageResponseRecordsBuilder) LoginCount(v int64) *LoginUsageStatisticsPageResponseRecordsBuilder {
	b.s.LoginCount = &v
	return b
}

func (b *LoginUsageStatisticsPageResponseRecordsBuilder) LongestLoginTime(v string) *LoginUsageStatisticsPageResponseRecordsBuilder {
	b.s.LongestLoginTime = &v
	return b
}

func (b *LoginUsageStatisticsPageResponseRecordsBuilder) OrgId(v string) *LoginUsageStatisticsPageResponseRecordsBuilder {
	b.s.OrgId = &v
	return b
}

func (b *LoginUsageStatisticsPageResponseRecordsBuilder) TenantId(v string) *LoginUsageStatisticsPageResponseRecordsBuilder {
	b.s.TenantId = &v
	return b
}

func (b *LoginUsageStatisticsPageResponseRecordsBuilder) AvgLoginTime(v string) *LoginUsageStatisticsPageResponseRecordsBuilder {
	b.s.AvgLoginTime = &v
	return b
}

func (b *LoginUsageStatisticsPageResponseRecordsBuilder) MopUserName(v string) *LoginUsageStatisticsPageResponseRecordsBuilder {
	b.s.MopUserName = &v
	return b
}

func (b *LoginUsageStatisticsPageResponseRecordsBuilder) MopUserId(v string) *LoginUsageStatisticsPageResponseRecordsBuilder {
	b.s.MopUserId = &v
	return b
}

func (b *LoginUsageStatisticsPageResponseRecordsBuilder) TotalLoginTime(v string) *LoginUsageStatisticsPageResponseRecordsBuilder {
	b.s.TotalLoginTime = &v
	return b
}

func (b *LoginUsageStatisticsPageResponseRecordsBuilder) Build() *LoginUsageStatisticsPageResponseRecords {
	return b.s
}


