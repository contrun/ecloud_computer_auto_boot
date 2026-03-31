// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetBwInstanceListBody struct {
    position.Body
    // 页码
	PageNumber *int32 `json:"pageNumber,omitempty"`
    // 带宽实例状态，1开通中、2已开通、3续订中、6/601已退订、8冻结中、7已冻结、10开通失败、1004审批中、18变更中
	StatusList []int32 `json:"statusList,omitempty"`
    // 实例列表
	InstanceIds []string `json:"instanceIds,omitempty"`
    // 页大小
	PageSize *int32 `json:"pageSize,omitempty"`
    // 计费类型 month：包月计费，year：包年计费
	MeasureUnit *string `json:"measureUnit,omitempty"`
    // 带宽名称
	BwName *string `json:"bwName,omitempty"`
}

func (s GetBwInstanceListBody) String() string {
	return utils.Beautify(s)
}

func (s GetBwInstanceListBody) GoString() string {
	return s.String()
}

func (s GetBwInstanceListBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetBwInstanceListBody) SetPageNumber(v int32) *GetBwInstanceListBody {
	s.PageNumber = &v
	return s
}

func (s *GetBwInstanceListBody) SetStatusList(v []int32) *GetBwInstanceListBody {
	s.StatusList = v
	return s
}

func (s *GetBwInstanceListBody) SetInstanceIds(v []string) *GetBwInstanceListBody {
	s.InstanceIds = v
	return s
}

func (s *GetBwInstanceListBody) SetPageSize(v int32) *GetBwInstanceListBody {
	s.PageSize = &v
	return s
}

func (s *GetBwInstanceListBody) SetMeasureUnit(v string) *GetBwInstanceListBody {
	s.MeasureUnit = &v
	return s
}

func (s *GetBwInstanceListBody) SetBwName(v string) *GetBwInstanceListBody {
	s.BwName = &v
	return s
}


type GetBwInstanceListBodyBuilder struct {
	s *GetBwInstanceListBody
}

func NewGetBwInstanceListBodyBuilder() *GetBwInstanceListBodyBuilder {
	s := &GetBwInstanceListBody{}
	b := &GetBwInstanceListBodyBuilder{s: s}
	return b
}

func (b *GetBwInstanceListBodyBuilder) PageNumber(v int32) *GetBwInstanceListBodyBuilder {
	b.s.PageNumber = &v
	return b
}

func (b *GetBwInstanceListBodyBuilder) StatusList(v []int32) *GetBwInstanceListBodyBuilder {
	b.s.StatusList = v
	return b
}

func (b *GetBwInstanceListBodyBuilder) InstanceIds(v []string) *GetBwInstanceListBodyBuilder {
	b.s.InstanceIds = v
	return b
}

func (b *GetBwInstanceListBodyBuilder) PageSize(v int32) *GetBwInstanceListBodyBuilder {
	b.s.PageSize = &v
	return b
}

func (b *GetBwInstanceListBodyBuilder) MeasureUnit(v string) *GetBwInstanceListBodyBuilder {
	b.s.MeasureUnit = &v
	return b
}

func (b *GetBwInstanceListBodyBuilder) BwName(v string) *GetBwInstanceListBodyBuilder {
	b.s.BwName = &v
	return b
}

func (b *GetBwInstanceListBodyBuilder) Build() *GetBwInstanceListBody {
	return b.s
}


