// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type SubmitBwBody struct {
    position.Body
    // 付费订购时长，包月：1-60，包年：1-5，免费时此字段不生效
	PeriodNum *int32 `json:"periodNum,omitempty"`
    // 是否免费
	IsFree *bool `json:"isFree,omitempty"`
    // 是否自动续订标识
	AutoRenew *bool `json:"autoRenew,omitempty"`
    // 带宽大小（Mbit/s）,范围1-2048
	BwSize *int32 `json:"bwSize,omitempty"`
    // 付费计费方式，month：包月，year：包年，免费时此字段不生效
	MeasureUnit *string `json:"measureUnit,omitempty"`
    // 带宽名称，字母、中划线和下划线的组合（必须以字母开头）。示例值：bandwidth-xxx
	BwName *string `json:"bwName,omitempty"`
    // 可用区编码。示例值：CIDC-CORE-00
	ZoneCode *string `json:"zoneCode,omitempty"`
}

func (s SubmitBwBody) String() string {
	return utils.Beautify(s)
}

func (s SubmitBwBody) GoString() string {
	return s.String()
}

func (s SubmitBwBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *SubmitBwBody) SetPeriodNum(v int32) *SubmitBwBody {
	s.PeriodNum = &v
	return s
}

func (s *SubmitBwBody) SetIsFree(v bool) *SubmitBwBody {
	s.IsFree = &v
	return s
}

func (s *SubmitBwBody) SetAutoRenew(v bool) *SubmitBwBody {
	s.AutoRenew = &v
	return s
}

func (s *SubmitBwBody) SetBwSize(v int32) *SubmitBwBody {
	s.BwSize = &v
	return s
}

func (s *SubmitBwBody) SetMeasureUnit(v string) *SubmitBwBody {
	s.MeasureUnit = &v
	return s
}

func (s *SubmitBwBody) SetBwName(v string) *SubmitBwBody {
	s.BwName = &v
	return s
}

func (s *SubmitBwBody) SetZoneCode(v string) *SubmitBwBody {
	s.ZoneCode = &v
	return s
}


type SubmitBwBodyBuilder struct {
	s *SubmitBwBody
}

func NewSubmitBwBodyBuilder() *SubmitBwBodyBuilder {
	s := &SubmitBwBody{}
	b := &SubmitBwBodyBuilder{s: s}
	return b
}

func (b *SubmitBwBodyBuilder) PeriodNum(v int32) *SubmitBwBodyBuilder {
	b.s.PeriodNum = &v
	return b
}

func (b *SubmitBwBodyBuilder) IsFree(v bool) *SubmitBwBodyBuilder {
	b.s.IsFree = &v
	return b
}

func (b *SubmitBwBodyBuilder) AutoRenew(v bool) *SubmitBwBodyBuilder {
	b.s.AutoRenew = &v
	return b
}

func (b *SubmitBwBodyBuilder) BwSize(v int32) *SubmitBwBodyBuilder {
	b.s.BwSize = &v
	return b
}

func (b *SubmitBwBodyBuilder) MeasureUnit(v string) *SubmitBwBodyBuilder {
	b.s.MeasureUnit = &v
	return b
}

func (b *SubmitBwBodyBuilder) BwName(v string) *SubmitBwBodyBuilder {
	b.s.BwName = &v
	return b
}

func (b *SubmitBwBodyBuilder) ZoneCode(v string) *SubmitBwBodyBuilder {
	b.s.ZoneCode = &v
	return b
}

func (b *SubmitBwBodyBuilder) Build() *SubmitBwBody {
	return b.s
}


