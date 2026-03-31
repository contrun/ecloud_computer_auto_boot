// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ModifyDeliveryInfoBody struct {
    position.Body
    // 收货人手机号。示例值：139xxx
	DeliveryPhone *string `json:"deliveryPhone,omitempty"`
    // 订单号。示例值：1549224402274672641
	OrderNo *string `json:"orderNo,omitempty"`
    // 收货人区域名称。示例值：东城区
	DeliveryAreaName *string `json:"deliveryAreaName,omitempty"`
    // 收货人区域编码。示例值：110101
	DeliveryAreaCode *string `json:"deliveryAreaCode,omitempty"`
    // 收货人城市名称。示例值：市辖区
	DeliveryCityName *string `json:"deliveryCityName,omitempty"`
    // 批次号
	BatchId *string `json:"batchId,omitempty"`
    // 操作类型：1修改 2取消
	Type *int32 `json:"type,omitempty"`
    // 收货人城市编码。示例值：110100
	DeliveryCityCode *string `json:"deliveryCityCode,omitempty"`
    // 收货人姓名。示例值：张三
	DeliveryUserName *string `json:"deliveryUserName,omitempty"`
    // 收货人详细地址。示例值：xx小区xx栋
	DeliveryAddress *string `json:"deliveryAddress,omitempty"`
    // 收货人邮箱。示例值：11@qq.com
	OrderUserMail *string `json:"orderUserMail,omitempty"`
    // 收货人省份编码。示例值：110000
	DeliveryProvinceCode *string `json:"deliveryProvinceCode,omitempty"`
    // 收货人省份名称。示例值：北京市
	DeliveryProvinceName *string `json:"deliveryProvinceName,omitempty"`
}

func (s ModifyDeliveryInfoBody) String() string {
	return utils.Beautify(s)
}

func (s ModifyDeliveryInfoBody) GoString() string {
	return s.String()
}

func (s ModifyDeliveryInfoBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ModifyDeliveryInfoBody) SetDeliveryPhone(v string) *ModifyDeliveryInfoBody {
	s.DeliveryPhone = &v
	return s
}

func (s *ModifyDeliveryInfoBody) SetOrderNo(v string) *ModifyDeliveryInfoBody {
	s.OrderNo = &v
	return s
}

func (s *ModifyDeliveryInfoBody) SetDeliveryAreaName(v string) *ModifyDeliveryInfoBody {
	s.DeliveryAreaName = &v
	return s
}

func (s *ModifyDeliveryInfoBody) SetDeliveryAreaCode(v string) *ModifyDeliveryInfoBody {
	s.DeliveryAreaCode = &v
	return s
}

func (s *ModifyDeliveryInfoBody) SetDeliveryCityName(v string) *ModifyDeliveryInfoBody {
	s.DeliveryCityName = &v
	return s
}

func (s *ModifyDeliveryInfoBody) SetBatchId(v string) *ModifyDeliveryInfoBody {
	s.BatchId = &v
	return s
}

func (s *ModifyDeliveryInfoBody) SetType(v int32) *ModifyDeliveryInfoBody {
	s.Type = &v
	return s
}

func (s *ModifyDeliveryInfoBody) SetDeliveryCityCode(v string) *ModifyDeliveryInfoBody {
	s.DeliveryCityCode = &v
	return s
}

func (s *ModifyDeliveryInfoBody) SetDeliveryUserName(v string) *ModifyDeliveryInfoBody {
	s.DeliveryUserName = &v
	return s
}

func (s *ModifyDeliveryInfoBody) SetDeliveryAddress(v string) *ModifyDeliveryInfoBody {
	s.DeliveryAddress = &v
	return s
}

func (s *ModifyDeliveryInfoBody) SetOrderUserMail(v string) *ModifyDeliveryInfoBody {
	s.OrderUserMail = &v
	return s
}

func (s *ModifyDeliveryInfoBody) SetDeliveryProvinceCode(v string) *ModifyDeliveryInfoBody {
	s.DeliveryProvinceCode = &v
	return s
}

func (s *ModifyDeliveryInfoBody) SetDeliveryProvinceName(v string) *ModifyDeliveryInfoBody {
	s.DeliveryProvinceName = &v
	return s
}


type ModifyDeliveryInfoBodyBuilder struct {
	s *ModifyDeliveryInfoBody
}

func NewModifyDeliveryInfoBodyBuilder() *ModifyDeliveryInfoBodyBuilder {
	s := &ModifyDeliveryInfoBody{}
	b := &ModifyDeliveryInfoBodyBuilder{s: s}
	return b
}

func (b *ModifyDeliveryInfoBodyBuilder) DeliveryPhone(v string) *ModifyDeliveryInfoBodyBuilder {
	b.s.DeliveryPhone = &v
	return b
}

func (b *ModifyDeliveryInfoBodyBuilder) OrderNo(v string) *ModifyDeliveryInfoBodyBuilder {
	b.s.OrderNo = &v
	return b
}

func (b *ModifyDeliveryInfoBodyBuilder) DeliveryAreaName(v string) *ModifyDeliveryInfoBodyBuilder {
	b.s.DeliveryAreaName = &v
	return b
}

func (b *ModifyDeliveryInfoBodyBuilder) DeliveryAreaCode(v string) *ModifyDeliveryInfoBodyBuilder {
	b.s.DeliveryAreaCode = &v
	return b
}

func (b *ModifyDeliveryInfoBodyBuilder) DeliveryCityName(v string) *ModifyDeliveryInfoBodyBuilder {
	b.s.DeliveryCityName = &v
	return b
}

func (b *ModifyDeliveryInfoBodyBuilder) BatchId(v string) *ModifyDeliveryInfoBodyBuilder {
	b.s.BatchId = &v
	return b
}

func (b *ModifyDeliveryInfoBodyBuilder) Type(v int32) *ModifyDeliveryInfoBodyBuilder {
	b.s.Type = &v
	return b
}

func (b *ModifyDeliveryInfoBodyBuilder) DeliveryCityCode(v string) *ModifyDeliveryInfoBodyBuilder {
	b.s.DeliveryCityCode = &v
	return b
}

func (b *ModifyDeliveryInfoBodyBuilder) DeliveryUserName(v string) *ModifyDeliveryInfoBodyBuilder {
	b.s.DeliveryUserName = &v
	return b
}

func (b *ModifyDeliveryInfoBodyBuilder) DeliveryAddress(v string) *ModifyDeliveryInfoBodyBuilder {
	b.s.DeliveryAddress = &v
	return b
}

func (b *ModifyDeliveryInfoBodyBuilder) OrderUserMail(v string) *ModifyDeliveryInfoBodyBuilder {
	b.s.OrderUserMail = &v
	return b
}

func (b *ModifyDeliveryInfoBodyBuilder) DeliveryProvinceCode(v string) *ModifyDeliveryInfoBodyBuilder {
	b.s.DeliveryProvinceCode = &v
	return b
}

func (b *ModifyDeliveryInfoBodyBuilder) DeliveryProvinceName(v string) *ModifyDeliveryInfoBodyBuilder {
	b.s.DeliveryProvinceName = &v
	return b
}

func (b *ModifyDeliveryInfoBodyBuilder) Build() *ModifyDeliveryInfoBody {
	return b.s
}


