// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CancelOrderBody struct {
    position.Body
    // 退订来源:1：官网  2：下沉 3：bc
	SourceCode *string `json:"sourceCode,omitempty"`
    // 订单号。示例值：1549224402274672641
	OrderNo *string `json:"orderNo,omitempty"`
    // 退订申请类型：1：未发货退订（类似之前的取消订单） 2:未拆箱 3：质量类问题 
	UnsubscribeType *int32 `json:"unsubscribeType,omitempty"`
    // 物流公司
	TrackingCompany *string `json:"trackingCompany,omitempty"`
    // 申请备注
	Remark *string `json:"remark,omitempty"`
    // 操作类型：1修改 2取消
	Type *int32 `json:"type,omitempty"`
    // 物流单号
	TrackingNumber *string `json:"trackingNumber,omitempty"`
}

func (s CancelOrderBody) String() string {
	return utils.Beautify(s)
}

func (s CancelOrderBody) GoString() string {
	return s.String()
}

func (s CancelOrderBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CancelOrderBody) SetSourceCode(v string) *CancelOrderBody {
	s.SourceCode = &v
	return s
}

func (s *CancelOrderBody) SetOrderNo(v string) *CancelOrderBody {
	s.OrderNo = &v
	return s
}

func (s *CancelOrderBody) SetUnsubscribeType(v int32) *CancelOrderBody {
	s.UnsubscribeType = &v
	return s
}

func (s *CancelOrderBody) SetTrackingCompany(v string) *CancelOrderBody {
	s.TrackingCompany = &v
	return s
}

func (s *CancelOrderBody) SetRemark(v string) *CancelOrderBody {
	s.Remark = &v
	return s
}

func (s *CancelOrderBody) SetType(v int32) *CancelOrderBody {
	s.Type = &v
	return s
}

func (s *CancelOrderBody) SetTrackingNumber(v string) *CancelOrderBody {
	s.TrackingNumber = &v
	return s
}


type CancelOrderBodyBuilder struct {
	s *CancelOrderBody
}

func NewCancelOrderBodyBuilder() *CancelOrderBodyBuilder {
	s := &CancelOrderBody{}
	b := &CancelOrderBodyBuilder{s: s}
	return b
}

func (b *CancelOrderBodyBuilder) SourceCode(v string) *CancelOrderBodyBuilder {
	b.s.SourceCode = &v
	return b
}

func (b *CancelOrderBodyBuilder) OrderNo(v string) *CancelOrderBodyBuilder {
	b.s.OrderNo = &v
	return b
}

func (b *CancelOrderBodyBuilder) UnsubscribeType(v int32) *CancelOrderBodyBuilder {
	b.s.UnsubscribeType = &v
	return b
}

func (b *CancelOrderBodyBuilder) TrackingCompany(v string) *CancelOrderBodyBuilder {
	b.s.TrackingCompany = &v
	return b
}

func (b *CancelOrderBodyBuilder) Remark(v string) *CancelOrderBodyBuilder {
	b.s.Remark = &v
	return b
}

func (b *CancelOrderBodyBuilder) Type(v int32) *CancelOrderBodyBuilder {
	b.s.Type = &v
	return b
}

func (b *CancelOrderBodyBuilder) TrackingNumber(v string) *CancelOrderBodyBuilder {
	b.s.TrackingNumber = &v
	return b
}

func (b *CancelOrderBodyBuilder) Build() *CancelOrderBody {
	return b.s
}


