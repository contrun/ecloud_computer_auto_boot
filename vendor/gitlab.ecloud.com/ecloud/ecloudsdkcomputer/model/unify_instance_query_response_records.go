// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UnifyInstanceQueryResponseRecords struct {

    // 收货人手机
	DeliveryPhone *string `json:"deliveryPhone,omitempty"`
    // 协议时长
	PeriodNum *string `json:"periodNum,omitempty"`
    // 收货人区域名称
	DeliveryAreaName *string `json:"deliveryAreaName,omitempty"`
    // 收货人区域编码
	DeliveryAreaCode *string `json:"deliveryAreaCode,omitempty"`
    // 收货人城市名称
	DeliveryCityName *string `json:"deliveryCityName,omitempty"`
    // 订单配置
	ProductConfig *string `json:"productConfig,omitempty"`
    // 收货人城市编码
	DeliveryCityCode *string `json:"deliveryCityCode,omitempty"`
    // 订购来源
	MarketName *string `json:"marketName,omitempty"`
    // 收货人姓名
	DeliveryUserName *string `json:"deliveryUserName,omitempty"`
    // 订购时间
	CreateTime *string `json:"createTime,omitempty"`
    // 收货地址
	DeliveryAddress *string `json:"deliveryAddress,omitempty"`
    // 是否是大单
	IsLarge *bool `json:"isLarge,omitempty"`
    // 邮箱
	OrderUserMail *string `json:"orderUserMail,omitempty"`
    // 交付进度
	Progress *string `json:"progress,omitempty"`
    // 订单号
	Id *string `json:"id,omitempty"`
    // 收货人省份编码
	DeliveryProvinceCode *string `json:"deliveryProvinceCode,omitempty"`
    // 收货人省份名称
	DeliveryProvinceName *string `json:"deliveryProvinceName,omitempty"`
    // 物流单号
	TrackingNumber *string `json:"trackingNumber,omitempty"`
    // 是否自研
	IsSelf *bool `json:"isSelf,omitempty"`
    // 软件配置/可选服务包说明
	ProductConfigDesc *string `json:"productConfigDesc,omitempty"`
    // 状态
	Status *int32 `json:"status,omitempty"`
}

func (s UnifyInstanceQueryResponseRecords) String() string {
	return utils.Beautify(s)
}

func (s UnifyInstanceQueryResponseRecords) GoString() string {
	return s.String()
}

func (s UnifyInstanceQueryResponseRecords) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UnifyInstanceQueryResponseRecords) SetDeliveryPhone(v string) *UnifyInstanceQueryResponseRecords {
	s.DeliveryPhone = &v
	return s
}

func (s *UnifyInstanceQueryResponseRecords) SetPeriodNum(v string) *UnifyInstanceQueryResponseRecords {
	s.PeriodNum = &v
	return s
}

func (s *UnifyInstanceQueryResponseRecords) SetDeliveryAreaName(v string) *UnifyInstanceQueryResponseRecords {
	s.DeliveryAreaName = &v
	return s
}

func (s *UnifyInstanceQueryResponseRecords) SetDeliveryAreaCode(v string) *UnifyInstanceQueryResponseRecords {
	s.DeliveryAreaCode = &v
	return s
}

func (s *UnifyInstanceQueryResponseRecords) SetDeliveryCityName(v string) *UnifyInstanceQueryResponseRecords {
	s.DeliveryCityName = &v
	return s
}

func (s *UnifyInstanceQueryResponseRecords) SetProductConfig(v string) *UnifyInstanceQueryResponseRecords {
	s.ProductConfig = &v
	return s
}

func (s *UnifyInstanceQueryResponseRecords) SetDeliveryCityCode(v string) *UnifyInstanceQueryResponseRecords {
	s.DeliveryCityCode = &v
	return s
}

func (s *UnifyInstanceQueryResponseRecords) SetMarketName(v string) *UnifyInstanceQueryResponseRecords {
	s.MarketName = &v
	return s
}

func (s *UnifyInstanceQueryResponseRecords) SetDeliveryUserName(v string) *UnifyInstanceQueryResponseRecords {
	s.DeliveryUserName = &v
	return s
}

func (s *UnifyInstanceQueryResponseRecords) SetCreateTime(v string) *UnifyInstanceQueryResponseRecords {
	s.CreateTime = &v
	return s
}

func (s *UnifyInstanceQueryResponseRecords) SetDeliveryAddress(v string) *UnifyInstanceQueryResponseRecords {
	s.DeliveryAddress = &v
	return s
}

func (s *UnifyInstanceQueryResponseRecords) SetIsLarge(v bool) *UnifyInstanceQueryResponseRecords {
	s.IsLarge = &v
	return s
}

func (s *UnifyInstanceQueryResponseRecords) SetOrderUserMail(v string) *UnifyInstanceQueryResponseRecords {
	s.OrderUserMail = &v
	return s
}

func (s *UnifyInstanceQueryResponseRecords) SetProgress(v string) *UnifyInstanceQueryResponseRecords {
	s.Progress = &v
	return s
}

func (s *UnifyInstanceQueryResponseRecords) SetId(v string) *UnifyInstanceQueryResponseRecords {
	s.Id = &v
	return s
}

func (s *UnifyInstanceQueryResponseRecords) SetDeliveryProvinceCode(v string) *UnifyInstanceQueryResponseRecords {
	s.DeliveryProvinceCode = &v
	return s
}

func (s *UnifyInstanceQueryResponseRecords) SetDeliveryProvinceName(v string) *UnifyInstanceQueryResponseRecords {
	s.DeliveryProvinceName = &v
	return s
}

func (s *UnifyInstanceQueryResponseRecords) SetTrackingNumber(v string) *UnifyInstanceQueryResponseRecords {
	s.TrackingNumber = &v
	return s
}

func (s *UnifyInstanceQueryResponseRecords) SetIsSelf(v bool) *UnifyInstanceQueryResponseRecords {
	s.IsSelf = &v
	return s
}

func (s *UnifyInstanceQueryResponseRecords) SetProductConfigDesc(v string) *UnifyInstanceQueryResponseRecords {
	s.ProductConfigDesc = &v
	return s
}

func (s *UnifyInstanceQueryResponseRecords) SetStatus(v int32) *UnifyInstanceQueryResponseRecords {
	s.Status = &v
	return s
}


type UnifyInstanceQueryResponseRecordsBuilder struct {
	s *UnifyInstanceQueryResponseRecords
}

func NewUnifyInstanceQueryResponseRecordsBuilder() *UnifyInstanceQueryResponseRecordsBuilder {
	s := &UnifyInstanceQueryResponseRecords{}
	b := &UnifyInstanceQueryResponseRecordsBuilder{s: s}
	return b
}

func (b *UnifyInstanceQueryResponseRecordsBuilder) DeliveryPhone(v string) *UnifyInstanceQueryResponseRecordsBuilder {
	b.s.DeliveryPhone = &v
	return b
}

func (b *UnifyInstanceQueryResponseRecordsBuilder) PeriodNum(v string) *UnifyInstanceQueryResponseRecordsBuilder {
	b.s.PeriodNum = &v
	return b
}

func (b *UnifyInstanceQueryResponseRecordsBuilder) DeliveryAreaName(v string) *UnifyInstanceQueryResponseRecordsBuilder {
	b.s.DeliveryAreaName = &v
	return b
}

func (b *UnifyInstanceQueryResponseRecordsBuilder) DeliveryAreaCode(v string) *UnifyInstanceQueryResponseRecordsBuilder {
	b.s.DeliveryAreaCode = &v
	return b
}

func (b *UnifyInstanceQueryResponseRecordsBuilder) DeliveryCityName(v string) *UnifyInstanceQueryResponseRecordsBuilder {
	b.s.DeliveryCityName = &v
	return b
}

func (b *UnifyInstanceQueryResponseRecordsBuilder) ProductConfig(v string) *UnifyInstanceQueryResponseRecordsBuilder {
	b.s.ProductConfig = &v
	return b
}

func (b *UnifyInstanceQueryResponseRecordsBuilder) DeliveryCityCode(v string) *UnifyInstanceQueryResponseRecordsBuilder {
	b.s.DeliveryCityCode = &v
	return b
}

func (b *UnifyInstanceQueryResponseRecordsBuilder) MarketName(v string) *UnifyInstanceQueryResponseRecordsBuilder {
	b.s.MarketName = &v
	return b
}

func (b *UnifyInstanceQueryResponseRecordsBuilder) DeliveryUserName(v string) *UnifyInstanceQueryResponseRecordsBuilder {
	b.s.DeliveryUserName = &v
	return b
}

func (b *UnifyInstanceQueryResponseRecordsBuilder) CreateTime(v string) *UnifyInstanceQueryResponseRecordsBuilder {
	b.s.CreateTime = &v
	return b
}

func (b *UnifyInstanceQueryResponseRecordsBuilder) DeliveryAddress(v string) *UnifyInstanceQueryResponseRecordsBuilder {
	b.s.DeliveryAddress = &v
	return b
}

func (b *UnifyInstanceQueryResponseRecordsBuilder) IsLarge(v bool) *UnifyInstanceQueryResponseRecordsBuilder {
	b.s.IsLarge = &v
	return b
}

func (b *UnifyInstanceQueryResponseRecordsBuilder) OrderUserMail(v string) *UnifyInstanceQueryResponseRecordsBuilder {
	b.s.OrderUserMail = &v
	return b
}

func (b *UnifyInstanceQueryResponseRecordsBuilder) Progress(v string) *UnifyInstanceQueryResponseRecordsBuilder {
	b.s.Progress = &v
	return b
}

func (b *UnifyInstanceQueryResponseRecordsBuilder) Id(v string) *UnifyInstanceQueryResponseRecordsBuilder {
	b.s.Id = &v
	return b
}

func (b *UnifyInstanceQueryResponseRecordsBuilder) DeliveryProvinceCode(v string) *UnifyInstanceQueryResponseRecordsBuilder {
	b.s.DeliveryProvinceCode = &v
	return b
}

func (b *UnifyInstanceQueryResponseRecordsBuilder) DeliveryProvinceName(v string) *UnifyInstanceQueryResponseRecordsBuilder {
	b.s.DeliveryProvinceName = &v
	return b
}

func (b *UnifyInstanceQueryResponseRecordsBuilder) TrackingNumber(v string) *UnifyInstanceQueryResponseRecordsBuilder {
	b.s.TrackingNumber = &v
	return b
}

func (b *UnifyInstanceQueryResponseRecordsBuilder) IsSelf(v bool) *UnifyInstanceQueryResponseRecordsBuilder {
	b.s.IsSelf = &v
	return b
}

func (b *UnifyInstanceQueryResponseRecordsBuilder) ProductConfigDesc(v string) *UnifyInstanceQueryResponseRecordsBuilder {
	b.s.ProductConfigDesc = &v
	return b
}

func (b *UnifyInstanceQueryResponseRecordsBuilder) Status(v int32) *UnifyInstanceQueryResponseRecordsBuilder {
	b.s.Status = &v
	return b
}

func (b *UnifyInstanceQueryResponseRecordsBuilder) Build() *UnifyInstanceQueryResponseRecords {
	return b.s
}


