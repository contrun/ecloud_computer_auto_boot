// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetOrgListBody struct {
    position.Body
    // 分页大小
	PageSize *string `json:"pageSize,omitempty"`
    // 当前页
	PageNum *string `json:"pageNum,omitempty"`
}

func (s GetOrgListBody) String() string {
	return utils.Beautify(s)
}

func (s GetOrgListBody) GoString() string {
	return s.String()
}

func (s GetOrgListBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetOrgListBody) SetPageSize(v string) *GetOrgListBody {
	s.PageSize = &v
	return s
}

func (s *GetOrgListBody) SetPageNum(v string) *GetOrgListBody {
	s.PageNum = &v
	return s
}


type GetOrgListBodyBuilder struct {
	s *GetOrgListBody
}

func NewGetOrgListBodyBuilder() *GetOrgListBodyBuilder {
	s := &GetOrgListBody{}
	b := &GetOrgListBodyBuilder{s: s}
	return b
}

func (b *GetOrgListBodyBuilder) PageSize(v string) *GetOrgListBodyBuilder {
	b.s.PageSize = &v
	return b
}

func (b *GetOrgListBodyBuilder) PageNum(v string) *GetOrgListBodyBuilder {
	b.s.PageNum = &v
	return b
}

func (b *GetOrgListBodyBuilder) Build() *GetOrgListBody {
	return b.s
}


