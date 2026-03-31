// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetBwInstanceListResponseRecords struct {

    // 订购周期，免费订单时此字段为空
	PeriodNum *int32 `json:"periodNum,omitempty"`
    // nat资源id
	NatInstanceId *string `json:"natInstanceId,omitempty"`
    // nat网关id
	NatGatWayId *string `json:"natGatWayId,omitempty"`
    // 公网ip地址
	PublicIpAddress *string `json:"publicIpAddress,omitempty"`
    // 云管可用区名称
	RegionName *string `json:"regionName,omitempty"`
    // 订单购买类型，0:赠送，1:购买'
	SaleType *string `json:"saleType,omitempty"`
    // month：包月计费，year：包年计费，免费订单此字段为空
	MeasureUnit *string `json:"measureUnit,omitempty"`
    // 带宽名称
	BwName *string `json:"bwName,omitempty"`
    // 带宽VPC的ID
	BwVpcId *string `json:"bwVpcId,omitempty"`
    // 带宽VPC名称
	BwVpcName *string `json:"bwVpcName,omitempty"`
    // mopO
	MopO *string `json:"mopO,omitempty"`
    // 实例ID
	InstanceId *string `json:"instanceId,omitempty"`
    // 业务子网名称，多个用英文逗号隔开
	BwSubnetName *string `json:"bwSubnetName,omitempty"`
    // mopT
	MopT *string `json:"mopT,omitempty"`
    // 是否自动续订，true:是，false:否
	AutoRenew *bool `json:"autoRenew,omitempty"`
    // 带宽大小
	BwSize *int32 `json:"bwSize,omitempty"`
    // 开通时间
	EffTime *string `json:"effTime,omitempty"`
    // 续订周期
	RenewDuration *int32 `json:"renewDuration,omitempty"`
    // 到期时间，免费订单此字段为空
	EndTime *string `json:"endTime,omitempty"`
    // 资源池名称
	PoolName *string `json:"poolName,omitempty"`
    // 带宽实例状态，1开通中、2已开通、3续订中、6/601已退订、8冻结中、7已冻结、10开通失败、1004审批中、18变更中
	Status *int32 `json:"status,omitempty"`
}

func (s GetBwInstanceListResponseRecords) String() string {
	return utils.Beautify(s)
}

func (s GetBwInstanceListResponseRecords) GoString() string {
	return s.String()
}

func (s GetBwInstanceListResponseRecords) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetBwInstanceListResponseRecords) SetPeriodNum(v int32) *GetBwInstanceListResponseRecords {
	s.PeriodNum = &v
	return s
}

func (s *GetBwInstanceListResponseRecords) SetNatInstanceId(v string) *GetBwInstanceListResponseRecords {
	s.NatInstanceId = &v
	return s
}

func (s *GetBwInstanceListResponseRecords) SetNatGatWayId(v string) *GetBwInstanceListResponseRecords {
	s.NatGatWayId = &v
	return s
}

func (s *GetBwInstanceListResponseRecords) SetPublicIpAddress(v string) *GetBwInstanceListResponseRecords {
	s.PublicIpAddress = &v
	return s
}

func (s *GetBwInstanceListResponseRecords) SetRegionName(v string) *GetBwInstanceListResponseRecords {
	s.RegionName = &v
	return s
}

func (s *GetBwInstanceListResponseRecords) SetSaleType(v string) *GetBwInstanceListResponseRecords {
	s.SaleType = &v
	return s
}

func (s *GetBwInstanceListResponseRecords) SetMeasureUnit(v string) *GetBwInstanceListResponseRecords {
	s.MeasureUnit = &v
	return s
}

func (s *GetBwInstanceListResponseRecords) SetBwName(v string) *GetBwInstanceListResponseRecords {
	s.BwName = &v
	return s
}

func (s *GetBwInstanceListResponseRecords) SetBwVpcId(v string) *GetBwInstanceListResponseRecords {
	s.BwVpcId = &v
	return s
}

func (s *GetBwInstanceListResponseRecords) SetBwVpcName(v string) *GetBwInstanceListResponseRecords {
	s.BwVpcName = &v
	return s
}

func (s *GetBwInstanceListResponseRecords) SetMopO(v string) *GetBwInstanceListResponseRecords {
	s.MopO = &v
	return s
}

func (s *GetBwInstanceListResponseRecords) SetInstanceId(v string) *GetBwInstanceListResponseRecords {
	s.InstanceId = &v
	return s
}

func (s *GetBwInstanceListResponseRecords) SetBwSubnetName(v string) *GetBwInstanceListResponseRecords {
	s.BwSubnetName = &v
	return s
}

func (s *GetBwInstanceListResponseRecords) SetMopT(v string) *GetBwInstanceListResponseRecords {
	s.MopT = &v
	return s
}

func (s *GetBwInstanceListResponseRecords) SetAutoRenew(v bool) *GetBwInstanceListResponseRecords {
	s.AutoRenew = &v
	return s
}

func (s *GetBwInstanceListResponseRecords) SetBwSize(v int32) *GetBwInstanceListResponseRecords {
	s.BwSize = &v
	return s
}

func (s *GetBwInstanceListResponseRecords) SetEffTime(v string) *GetBwInstanceListResponseRecords {
	s.EffTime = &v
	return s
}

func (s *GetBwInstanceListResponseRecords) SetRenewDuration(v int32) *GetBwInstanceListResponseRecords {
	s.RenewDuration = &v
	return s
}

func (s *GetBwInstanceListResponseRecords) SetEndTime(v string) *GetBwInstanceListResponseRecords {
	s.EndTime = &v
	return s
}

func (s *GetBwInstanceListResponseRecords) SetPoolName(v string) *GetBwInstanceListResponseRecords {
	s.PoolName = &v
	return s
}

func (s *GetBwInstanceListResponseRecords) SetStatus(v int32) *GetBwInstanceListResponseRecords {
	s.Status = &v
	return s
}


type GetBwInstanceListResponseRecordsBuilder struct {
	s *GetBwInstanceListResponseRecords
}

func NewGetBwInstanceListResponseRecordsBuilder() *GetBwInstanceListResponseRecordsBuilder {
	s := &GetBwInstanceListResponseRecords{}
	b := &GetBwInstanceListResponseRecordsBuilder{s: s}
	return b
}

func (b *GetBwInstanceListResponseRecordsBuilder) PeriodNum(v int32) *GetBwInstanceListResponseRecordsBuilder {
	b.s.PeriodNum = &v
	return b
}

func (b *GetBwInstanceListResponseRecordsBuilder) NatInstanceId(v string) *GetBwInstanceListResponseRecordsBuilder {
	b.s.NatInstanceId = &v
	return b
}

func (b *GetBwInstanceListResponseRecordsBuilder) NatGatWayId(v string) *GetBwInstanceListResponseRecordsBuilder {
	b.s.NatGatWayId = &v
	return b
}

func (b *GetBwInstanceListResponseRecordsBuilder) PublicIpAddress(v string) *GetBwInstanceListResponseRecordsBuilder {
	b.s.PublicIpAddress = &v
	return b
}

func (b *GetBwInstanceListResponseRecordsBuilder) RegionName(v string) *GetBwInstanceListResponseRecordsBuilder {
	b.s.RegionName = &v
	return b
}

func (b *GetBwInstanceListResponseRecordsBuilder) SaleType(v string) *GetBwInstanceListResponseRecordsBuilder {
	b.s.SaleType = &v
	return b
}

func (b *GetBwInstanceListResponseRecordsBuilder) MeasureUnit(v string) *GetBwInstanceListResponseRecordsBuilder {
	b.s.MeasureUnit = &v
	return b
}

func (b *GetBwInstanceListResponseRecordsBuilder) BwName(v string) *GetBwInstanceListResponseRecordsBuilder {
	b.s.BwName = &v
	return b
}

func (b *GetBwInstanceListResponseRecordsBuilder) BwVpcId(v string) *GetBwInstanceListResponseRecordsBuilder {
	b.s.BwVpcId = &v
	return b
}

func (b *GetBwInstanceListResponseRecordsBuilder) BwVpcName(v string) *GetBwInstanceListResponseRecordsBuilder {
	b.s.BwVpcName = &v
	return b
}

func (b *GetBwInstanceListResponseRecordsBuilder) MopO(v string) *GetBwInstanceListResponseRecordsBuilder {
	b.s.MopO = &v
	return b
}

func (b *GetBwInstanceListResponseRecordsBuilder) InstanceId(v string) *GetBwInstanceListResponseRecordsBuilder {
	b.s.InstanceId = &v
	return b
}

func (b *GetBwInstanceListResponseRecordsBuilder) BwSubnetName(v string) *GetBwInstanceListResponseRecordsBuilder {
	b.s.BwSubnetName = &v
	return b
}

func (b *GetBwInstanceListResponseRecordsBuilder) MopT(v string) *GetBwInstanceListResponseRecordsBuilder {
	b.s.MopT = &v
	return b
}

func (b *GetBwInstanceListResponseRecordsBuilder) AutoRenew(v bool) *GetBwInstanceListResponseRecordsBuilder {
	b.s.AutoRenew = &v
	return b
}

func (b *GetBwInstanceListResponseRecordsBuilder) BwSize(v int32) *GetBwInstanceListResponseRecordsBuilder {
	b.s.BwSize = &v
	return b
}

func (b *GetBwInstanceListResponseRecordsBuilder) EffTime(v string) *GetBwInstanceListResponseRecordsBuilder {
	b.s.EffTime = &v
	return b
}

func (b *GetBwInstanceListResponseRecordsBuilder) RenewDuration(v int32) *GetBwInstanceListResponseRecordsBuilder {
	b.s.RenewDuration = &v
	return b
}

func (b *GetBwInstanceListResponseRecordsBuilder) EndTime(v string) *GetBwInstanceListResponseRecordsBuilder {
	b.s.EndTime = &v
	return b
}

func (b *GetBwInstanceListResponseRecordsBuilder) PoolName(v string) *GetBwInstanceListResponseRecordsBuilder {
	b.s.PoolName = &v
	return b
}

func (b *GetBwInstanceListResponseRecordsBuilder) Status(v int32) *GetBwInstanceListResponseRecordsBuilder {
	b.s.Status = &v
	return b
}

func (b *GetBwInstanceListResponseRecordsBuilder) Build() *GetBwInstanceListResponseRecords {
	return b.s
}


