// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type LoginLivingPageDailyBody struct {
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

func (s LoginLivingPageDailyBody) String() string {
	return utils.Beautify(s)
}

func (s LoginLivingPageDailyBody) GoString() string {
	return s.String()
}

func (s LoginLivingPageDailyBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *LoginLivingPageDailyBody) SetOrgIdList(v []string) *LoginLivingPageDailyBody {
	s.OrgIdList = v
	return s
}

func (s *LoginLivingPageDailyBody) SetPageSize(v int32) *LoginLivingPageDailyBody {
	s.PageSize = &v
	return s
}

func (s *LoginLivingPageDailyBody) SetStartTime(v string) *LoginLivingPageDailyBody {
	s.StartTime = &v
	return s
}

func (s *LoginLivingPageDailyBody) SetEndTime(v string) *LoginLivingPageDailyBody {
	s.EndTime = &v
	return s
}

func (s *LoginLivingPageDailyBody) SetPage(v int32) *LoginLivingPageDailyBody {
	s.Page = &v
	return s
}


type LoginLivingPageDailyBodyBuilder struct {
	s *LoginLivingPageDailyBody
}

func NewLoginLivingPageDailyBodyBuilder() *LoginLivingPageDailyBodyBuilder {
	s := &LoginLivingPageDailyBody{}
	b := &LoginLivingPageDailyBodyBuilder{s: s}
	return b
}

func (b *LoginLivingPageDailyBodyBuilder) OrgIdList(v []string) *LoginLivingPageDailyBodyBuilder {
	b.s.OrgIdList = v
	return b
}

func (b *LoginLivingPageDailyBodyBuilder) PageSize(v int32) *LoginLivingPageDailyBodyBuilder {
	b.s.PageSize = &v
	return b
}

func (b *LoginLivingPageDailyBodyBuilder) StartTime(v string) *LoginLivingPageDailyBodyBuilder {
	b.s.StartTime = &v
	return b
}

func (b *LoginLivingPageDailyBodyBuilder) EndTime(v string) *LoginLivingPageDailyBodyBuilder {
	b.s.EndTime = &v
	return b
}

func (b *LoginLivingPageDailyBodyBuilder) Page(v int32) *LoginLivingPageDailyBodyBuilder {
	b.s.Page = &v
	return b
}

func (b *LoginLivingPageDailyBodyBuilder) Build() *LoginLivingPageDailyBody {
	return b.s
}


