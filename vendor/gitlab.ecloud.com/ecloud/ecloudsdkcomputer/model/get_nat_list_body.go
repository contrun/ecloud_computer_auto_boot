// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetNatListBody struct {
    position.Body
    // vpc名称
	VpcName *string `json:"vpcName,omitempty"`
    // 页码
	PageNumber *int32 `json:"pageNumber,omitempty"`
    // 实例状态，1开通中、2已开通、3续订中、6/601已退订、8冻结中、7已冻结、10开通失败、1004审批中、18变更中
	StatusList []int32 `json:"statusList,omitempty"`
    // 可用区编码
	RegionId *string `json:"regionId,omitempty"`
    // 实例集合
	InstanceIds []string `json:"instanceIds,omitempty"`
    // 网关id
	NatId *string `json:"natId,omitempty"`
    // 网关名称
	NatName *string `json:"natName,omitempty"`
    // 资源池ID
	PoolId *string `json:"poolId,omitempty"`
    // 页大小
	PageSize *int32 `json:"pageSize,omitempty"`
    // 计费类型
	MeasureUnit *string `json:"measureUnit,omitempty"`
}

func (s GetNatListBody) String() string {
	return utils.Beautify(s)
}

func (s GetNatListBody) GoString() string {
	return s.String()
}

func (s GetNatListBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetNatListBody) SetVpcName(v string) *GetNatListBody {
	s.VpcName = &v
	return s
}

func (s *GetNatListBody) SetPageNumber(v int32) *GetNatListBody {
	s.PageNumber = &v
	return s
}

func (s *GetNatListBody) SetStatusList(v []int32) *GetNatListBody {
	s.StatusList = v
	return s
}

func (s *GetNatListBody) SetRegionId(v string) *GetNatListBody {
	s.RegionId = &v
	return s
}

func (s *GetNatListBody) SetInstanceIds(v []string) *GetNatListBody {
	s.InstanceIds = v
	return s
}

func (s *GetNatListBody) SetNatId(v string) *GetNatListBody {
	s.NatId = &v
	return s
}

func (s *GetNatListBody) SetNatName(v string) *GetNatListBody {
	s.NatName = &v
	return s
}

func (s *GetNatListBody) SetPoolId(v string) *GetNatListBody {
	s.PoolId = &v
	return s
}

func (s *GetNatListBody) SetPageSize(v int32) *GetNatListBody {
	s.PageSize = &v
	return s
}

func (s *GetNatListBody) SetMeasureUnit(v string) *GetNatListBody {
	s.MeasureUnit = &v
	return s
}


type GetNatListBodyBuilder struct {
	s *GetNatListBody
}

func NewGetNatListBodyBuilder() *GetNatListBodyBuilder {
	s := &GetNatListBody{}
	b := &GetNatListBodyBuilder{s: s}
	return b
}

func (b *GetNatListBodyBuilder) VpcName(v string) *GetNatListBodyBuilder {
	b.s.VpcName = &v
	return b
}

func (b *GetNatListBodyBuilder) PageNumber(v int32) *GetNatListBodyBuilder {
	b.s.PageNumber = &v
	return b
}

func (b *GetNatListBodyBuilder) StatusList(v []int32) *GetNatListBodyBuilder {
	b.s.StatusList = v
	return b
}

func (b *GetNatListBodyBuilder) RegionId(v string) *GetNatListBodyBuilder {
	b.s.RegionId = &v
	return b
}

func (b *GetNatListBodyBuilder) InstanceIds(v []string) *GetNatListBodyBuilder {
	b.s.InstanceIds = v
	return b
}

func (b *GetNatListBodyBuilder) NatId(v string) *GetNatListBodyBuilder {
	b.s.NatId = &v
	return b
}

func (b *GetNatListBodyBuilder) NatName(v string) *GetNatListBodyBuilder {
	b.s.NatName = &v
	return b
}

func (b *GetNatListBodyBuilder) PoolId(v string) *GetNatListBodyBuilder {
	b.s.PoolId = &v
	return b
}

func (b *GetNatListBodyBuilder) PageSize(v int32) *GetNatListBodyBuilder {
	b.s.PageSize = &v
	return b
}

func (b *GetNatListBodyBuilder) MeasureUnit(v string) *GetNatListBodyBuilder {
	b.s.MeasureUnit = &v
	return b
}

func (b *GetNatListBodyBuilder) Build() *GetNatListBody {
	return b.s
}


