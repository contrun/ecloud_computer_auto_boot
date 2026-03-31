// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetNatListResponseRecords struct {

    // 订购时长，免费订单此字段为空
	PeriodNum *int32 `json:"periodNum,omitempty"`
    // vpc名称
	VpcName *string `json:"vpcName,omitempty"`
    // natId
	NatId *string `json:"natId,omitempty"`
    // 可用区名称
	RegionName *string `json:"regionName,omitempty"`
    // 订单购买类型，0:赠送，1:购买
	SaleType *string `json:"saleType,omitempty"`
    // 计费类型 month：包月计费，year：包年计费，免费订单此字段为空
	MeasureUnit *string `json:"measureUnit,omitempty"`
    // mop可用区编码
	RegionIdMop *string `json:"regionIdMop,omitempty"`
    // mop订单id
	MopO *string `json:"mopO,omitempty"`
    // nat实例id
	InstanceId *string `json:"instanceId,omitempty"`
    // nat对应的vpcUid
	VpcUid *string `json:"vpcUid,omitempty"`
    // 订单规格编码
	ChaId *string `json:"chaId,omitempty"`
    // mop订单子项id
	MopT *string `json:"mopT,omitempty"`
    // 云管可用区编码
	RegionId *string `json:"regionId,omitempty"`
    // nat名称
	NatName *string `json:"natName,omitempty"`
    // 资源池编码
	PoolId *string `json:"poolId,omitempty"`
    // 是否自动续订，true:是，false:否
	AutoRenew *bool `json:"autoRenew,omitempty"`
    // 开通时间
	EffTime *string `json:"effTime,omitempty"`
    // 续订周期
	RenewDuration *int32 `json:"renewDuration,omitempty"`
    // 到期时间，免费订单此字段为空
	EndTime *string `json:"endTime,omitempty"`
    // 规格编码
	SpecificationCode *string `json:"specificationCode,omitempty"`
    // 资源池名称
	PoolName *string `json:"poolName,omitempty"`
    // 订单订购来源
	SourceFrom *string `json:"sourceFrom,omitempty"`
    // 实例状态，1开通中、2已开通、3续订中、6/601已退订、8冻结中、7已冻结、10开通失败、1004审批中、18变更中
	Status *int32 `json:"status,omitempty"`
}

func (s GetNatListResponseRecords) String() string {
	return utils.Beautify(s)
}

func (s GetNatListResponseRecords) GoString() string {
	return s.String()
}

func (s GetNatListResponseRecords) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetNatListResponseRecords) SetPeriodNum(v int32) *GetNatListResponseRecords {
	s.PeriodNum = &v
	return s
}

func (s *GetNatListResponseRecords) SetVpcName(v string) *GetNatListResponseRecords {
	s.VpcName = &v
	return s
}

func (s *GetNatListResponseRecords) SetNatId(v string) *GetNatListResponseRecords {
	s.NatId = &v
	return s
}

func (s *GetNatListResponseRecords) SetRegionName(v string) *GetNatListResponseRecords {
	s.RegionName = &v
	return s
}

func (s *GetNatListResponseRecords) SetSaleType(v string) *GetNatListResponseRecords {
	s.SaleType = &v
	return s
}

func (s *GetNatListResponseRecords) SetMeasureUnit(v string) *GetNatListResponseRecords {
	s.MeasureUnit = &v
	return s
}

func (s *GetNatListResponseRecords) SetRegionIdMop(v string) *GetNatListResponseRecords {
	s.RegionIdMop = &v
	return s
}

func (s *GetNatListResponseRecords) SetMopO(v string) *GetNatListResponseRecords {
	s.MopO = &v
	return s
}

func (s *GetNatListResponseRecords) SetInstanceId(v string) *GetNatListResponseRecords {
	s.InstanceId = &v
	return s
}

func (s *GetNatListResponseRecords) SetVpcUid(v string) *GetNatListResponseRecords {
	s.VpcUid = &v
	return s
}

func (s *GetNatListResponseRecords) SetChaId(v string) *GetNatListResponseRecords {
	s.ChaId = &v
	return s
}

func (s *GetNatListResponseRecords) SetMopT(v string) *GetNatListResponseRecords {
	s.MopT = &v
	return s
}

func (s *GetNatListResponseRecords) SetRegionId(v string) *GetNatListResponseRecords {
	s.RegionId = &v
	return s
}

func (s *GetNatListResponseRecords) SetNatName(v string) *GetNatListResponseRecords {
	s.NatName = &v
	return s
}

func (s *GetNatListResponseRecords) SetPoolId(v string) *GetNatListResponseRecords {
	s.PoolId = &v
	return s
}

func (s *GetNatListResponseRecords) SetAutoRenew(v bool) *GetNatListResponseRecords {
	s.AutoRenew = &v
	return s
}

func (s *GetNatListResponseRecords) SetEffTime(v string) *GetNatListResponseRecords {
	s.EffTime = &v
	return s
}

func (s *GetNatListResponseRecords) SetRenewDuration(v int32) *GetNatListResponseRecords {
	s.RenewDuration = &v
	return s
}

func (s *GetNatListResponseRecords) SetEndTime(v string) *GetNatListResponseRecords {
	s.EndTime = &v
	return s
}

func (s *GetNatListResponseRecords) SetSpecificationCode(v string) *GetNatListResponseRecords {
	s.SpecificationCode = &v
	return s
}

func (s *GetNatListResponseRecords) SetPoolName(v string) *GetNatListResponseRecords {
	s.PoolName = &v
	return s
}

func (s *GetNatListResponseRecords) SetSourceFrom(v string) *GetNatListResponseRecords {
	s.SourceFrom = &v
	return s
}

func (s *GetNatListResponseRecords) SetStatus(v int32) *GetNatListResponseRecords {
	s.Status = &v
	return s
}


type GetNatListResponseRecordsBuilder struct {
	s *GetNatListResponseRecords
}

func NewGetNatListResponseRecordsBuilder() *GetNatListResponseRecordsBuilder {
	s := &GetNatListResponseRecords{}
	b := &GetNatListResponseRecordsBuilder{s: s}
	return b
}

func (b *GetNatListResponseRecordsBuilder) PeriodNum(v int32) *GetNatListResponseRecordsBuilder {
	b.s.PeriodNum = &v
	return b
}

func (b *GetNatListResponseRecordsBuilder) VpcName(v string) *GetNatListResponseRecordsBuilder {
	b.s.VpcName = &v
	return b
}

func (b *GetNatListResponseRecordsBuilder) NatId(v string) *GetNatListResponseRecordsBuilder {
	b.s.NatId = &v
	return b
}

func (b *GetNatListResponseRecordsBuilder) RegionName(v string) *GetNatListResponseRecordsBuilder {
	b.s.RegionName = &v
	return b
}

func (b *GetNatListResponseRecordsBuilder) SaleType(v string) *GetNatListResponseRecordsBuilder {
	b.s.SaleType = &v
	return b
}

func (b *GetNatListResponseRecordsBuilder) MeasureUnit(v string) *GetNatListResponseRecordsBuilder {
	b.s.MeasureUnit = &v
	return b
}

func (b *GetNatListResponseRecordsBuilder) RegionIdMop(v string) *GetNatListResponseRecordsBuilder {
	b.s.RegionIdMop = &v
	return b
}

func (b *GetNatListResponseRecordsBuilder) MopO(v string) *GetNatListResponseRecordsBuilder {
	b.s.MopO = &v
	return b
}

func (b *GetNatListResponseRecordsBuilder) InstanceId(v string) *GetNatListResponseRecordsBuilder {
	b.s.InstanceId = &v
	return b
}

func (b *GetNatListResponseRecordsBuilder) VpcUid(v string) *GetNatListResponseRecordsBuilder {
	b.s.VpcUid = &v
	return b
}

func (b *GetNatListResponseRecordsBuilder) ChaId(v string) *GetNatListResponseRecordsBuilder {
	b.s.ChaId = &v
	return b
}

func (b *GetNatListResponseRecordsBuilder) MopT(v string) *GetNatListResponseRecordsBuilder {
	b.s.MopT = &v
	return b
}

func (b *GetNatListResponseRecordsBuilder) RegionId(v string) *GetNatListResponseRecordsBuilder {
	b.s.RegionId = &v
	return b
}

func (b *GetNatListResponseRecordsBuilder) NatName(v string) *GetNatListResponseRecordsBuilder {
	b.s.NatName = &v
	return b
}

func (b *GetNatListResponseRecordsBuilder) PoolId(v string) *GetNatListResponseRecordsBuilder {
	b.s.PoolId = &v
	return b
}

func (b *GetNatListResponseRecordsBuilder) AutoRenew(v bool) *GetNatListResponseRecordsBuilder {
	b.s.AutoRenew = &v
	return b
}

func (b *GetNatListResponseRecordsBuilder) EffTime(v string) *GetNatListResponseRecordsBuilder {
	b.s.EffTime = &v
	return b
}

func (b *GetNatListResponseRecordsBuilder) RenewDuration(v int32) *GetNatListResponseRecordsBuilder {
	b.s.RenewDuration = &v
	return b
}

func (b *GetNatListResponseRecordsBuilder) EndTime(v string) *GetNatListResponseRecordsBuilder {
	b.s.EndTime = &v
	return b
}

func (b *GetNatListResponseRecordsBuilder) SpecificationCode(v string) *GetNatListResponseRecordsBuilder {
	b.s.SpecificationCode = &v
	return b
}

func (b *GetNatListResponseRecordsBuilder) PoolName(v string) *GetNatListResponseRecordsBuilder {
	b.s.PoolName = &v
	return b
}

func (b *GetNatListResponseRecordsBuilder) SourceFrom(v string) *GetNatListResponseRecordsBuilder {
	b.s.SourceFrom = &v
	return b
}

func (b *GetNatListResponseRecordsBuilder) Status(v int32) *GetNatListResponseRecordsBuilder {
	b.s.Status = &v
	return b
}

func (b *GetNatListResponseRecordsBuilder) Build() *GetNatListResponseRecords {
	return b.s
}


