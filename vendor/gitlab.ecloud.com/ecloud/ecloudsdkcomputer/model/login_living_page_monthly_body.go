// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type LoginLivingPageMonthlyBody struct {
    position.Body
    // 组织id列表
	OrgIdList []string `json:"orgIdList,omitempty"`
    // 分页大小
	PageSize *int32 `json:"pageSize,omitempty"`
    // 开始时间
	StartTime *string `json:"startTime,omitempty"`
    // 结束时间
	EndTime *string `json:"endTime,omitempty"`
    // 当前页
	Page *int32 `json:"page,omitempty"`
}

func (s LoginLivingPageMonthlyBody) String() string {
	return utils.Beautify(s)
}

func (s LoginLivingPageMonthlyBody) GoString() string {
	return s.String()
}

func (s LoginLivingPageMonthlyBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *LoginLivingPageMonthlyBody) SetOrgIdList(v []string) *LoginLivingPageMonthlyBody {
	s.OrgIdList = v
	return s
}

func (s *LoginLivingPageMonthlyBody) SetPageSize(v int32) *LoginLivingPageMonthlyBody {
	s.PageSize = &v
	return s
}

func (s *LoginLivingPageMonthlyBody) SetStartTime(v string) *LoginLivingPageMonthlyBody {
	s.StartTime = &v
	return s
}

func (s *LoginLivingPageMonthlyBody) SetEndTime(v string) *LoginLivingPageMonthlyBody {
	s.EndTime = &v
	return s
}

func (s *LoginLivingPageMonthlyBody) SetPage(v int32) *LoginLivingPageMonthlyBody {
	s.Page = &v
	return s
}


type LoginLivingPageMonthlyBodyBuilder struct {
	s *LoginLivingPageMonthlyBody
}

func NewLoginLivingPageMonthlyBodyBuilder() *LoginLivingPageMonthlyBodyBuilder {
	s := &LoginLivingPageMonthlyBody{}
	b := &LoginLivingPageMonthlyBodyBuilder{s: s}
	return b
}

func (b *LoginLivingPageMonthlyBodyBuilder) OrgIdList(v []string) *LoginLivingPageMonthlyBodyBuilder {
	b.s.OrgIdList = v
	return b
}

func (b *LoginLivingPageMonthlyBodyBuilder) PageSize(v int32) *LoginLivingPageMonthlyBodyBuilder {
	b.s.PageSize = &v
	return b
}

func (b *LoginLivingPageMonthlyBodyBuilder) StartTime(v string) *LoginLivingPageMonthlyBodyBuilder {
	b.s.StartTime = &v
	return b
}

func (b *LoginLivingPageMonthlyBodyBuilder) EndTime(v string) *LoginLivingPageMonthlyBodyBuilder {
	b.s.EndTime = &v
	return b
}

func (b *LoginLivingPageMonthlyBodyBuilder) Page(v int32) *LoginLivingPageMonthlyBodyBuilder {
	b.s.Page = &v
	return b
}

func (b *LoginLivingPageMonthlyBodyBuilder) Build() *LoginLivingPageMonthlyBody {
	return b.s
}


