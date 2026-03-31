// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type SubmitCcaOrderBody struct {
    position.Body
    // 订购时长，包月：1-60，包年：1-5
	PeriodNum *string `json:"periodNum,omitempty"`
    // 订购数量（范围1-50的整数）
	GroupCount *string `json:"groupCount,omitempty"`
    // 操作系统编码，0：Windows 10，3：Windows 11，4：Windows Server 2016
	ImageCode *string `json:"imageCode,omitempty"`
    // 是否自动续订标识，true：是，false：否，默认为false
	AutoRenew *bool `json:"autoRenew,omitempty"`
    // 套餐标识，1：办公型-公众版（不带数据盘），2：办公型-公众版（带数据盘），4：行业型-娱乐版
	Combo *string `json:"combo,omitempty"`
    // 邮箱。示例值：zhangsan@qq.com
	OrderUserMail *string `json:"orderUserMail,omitempty"`
    // 计费方式，month：包月，year：包年
	MeasureUnit *string `json:"measureUnit,omitempty"`
    // 额外数据盘数量，单位GB
	DataDiskCount *string `json:"dataDiskCount,omitempty"`
    // 规格类型，0：2vCPU/4GB内存/80GB系统盘/50M带宽；娱乐型5小时月包，1：4vCPU/8GB内存/80GB系统盘/50M带宽；娱乐型15小时月包，2：8vCPU/16GB内存/80GB系统盘/50M带宽；娱乐型30小时月包，4：4vCPU/4GB内存/80GB系统盘/50M带宽，5：8vCPU/8GB内存/80GB系统盘/50M带宽，6：16vCPU/16GB内存/80GB系统盘/50M带宽，7：16vCPU/32GB内存/80GB系统盘/50M带宽
	SpecificationCode *string `json:"specificationCode,omitempty"`
}

func (s SubmitCcaOrderBody) String() string {
	return utils.Beautify(s)
}

func (s SubmitCcaOrderBody) GoString() string {
	return s.String()
}

func (s SubmitCcaOrderBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *SubmitCcaOrderBody) SetPeriodNum(v string) *SubmitCcaOrderBody {
	s.PeriodNum = &v
	return s
}

func (s *SubmitCcaOrderBody) SetGroupCount(v string) *SubmitCcaOrderBody {
	s.GroupCount = &v
	return s
}

func (s *SubmitCcaOrderBody) SetImageCode(v string) *SubmitCcaOrderBody {
	s.ImageCode = &v
	return s
}

func (s *SubmitCcaOrderBody) SetAutoRenew(v bool) *SubmitCcaOrderBody {
	s.AutoRenew = &v
	return s
}

func (s *SubmitCcaOrderBody) SetCombo(v string) *SubmitCcaOrderBody {
	s.Combo = &v
	return s
}

func (s *SubmitCcaOrderBody) SetOrderUserMail(v string) *SubmitCcaOrderBody {
	s.OrderUserMail = &v
	return s
}

func (s *SubmitCcaOrderBody) SetMeasureUnit(v string) *SubmitCcaOrderBody {
	s.MeasureUnit = &v
	return s
}

func (s *SubmitCcaOrderBody) SetDataDiskCount(v string) *SubmitCcaOrderBody {
	s.DataDiskCount = &v
	return s
}

func (s *SubmitCcaOrderBody) SetSpecificationCode(v string) *SubmitCcaOrderBody {
	s.SpecificationCode = &v
	return s
}


type SubmitCcaOrderBodyBuilder struct {
	s *SubmitCcaOrderBody
}

func NewSubmitCcaOrderBodyBuilder() *SubmitCcaOrderBodyBuilder {
	s := &SubmitCcaOrderBody{}
	b := &SubmitCcaOrderBodyBuilder{s: s}
	return b
}

func (b *SubmitCcaOrderBodyBuilder) PeriodNum(v string) *SubmitCcaOrderBodyBuilder {
	b.s.PeriodNum = &v
	return b
}

func (b *SubmitCcaOrderBodyBuilder) GroupCount(v string) *SubmitCcaOrderBodyBuilder {
	b.s.GroupCount = &v
	return b
}

func (b *SubmitCcaOrderBodyBuilder) ImageCode(v string) *SubmitCcaOrderBodyBuilder {
	b.s.ImageCode = &v
	return b
}

func (b *SubmitCcaOrderBodyBuilder) AutoRenew(v bool) *SubmitCcaOrderBodyBuilder {
	b.s.AutoRenew = &v
	return b
}

func (b *SubmitCcaOrderBodyBuilder) Combo(v string) *SubmitCcaOrderBodyBuilder {
	b.s.Combo = &v
	return b
}

func (b *SubmitCcaOrderBodyBuilder) OrderUserMail(v string) *SubmitCcaOrderBodyBuilder {
	b.s.OrderUserMail = &v
	return b
}

func (b *SubmitCcaOrderBodyBuilder) MeasureUnit(v string) *SubmitCcaOrderBodyBuilder {
	b.s.MeasureUnit = &v
	return b
}

func (b *SubmitCcaOrderBodyBuilder) DataDiskCount(v string) *SubmitCcaOrderBodyBuilder {
	b.s.DataDiskCount = &v
	return b
}

func (b *SubmitCcaOrderBodyBuilder) SpecificationCode(v string) *SubmitCcaOrderBodyBuilder {
	b.s.SpecificationCode = &v
	return b
}

func (b *SubmitCcaOrderBodyBuilder) Build() *SubmitCcaOrderBody {
	return b.s
}


