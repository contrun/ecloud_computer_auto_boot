// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type SubmitCcaOrderGovBody struct {
    position.Body
    // 订购时长，包月：1-60，包年：1-5
	PeriodNum *string `json:"periodNum,omitempty"`
    // 是否订购高级管理服务包标志，true：是，false：否，默认为false
	IsOrderAdvancedPackage *bool `json:"isOrderAdvancedPackage,omitempty"`
    // 订购数量（范围1-50的整数）
	GroupCount *string `json:"groupCount,omitempty"`
    // 操作系统编码，0：Windows 10，1：统信UOS_1050，2：麒麟KylinOS_V10，3：Windows 11，4：Windows Server 2016
	ImageCode *string `json:"imageCode,omitempty"`
    // 是否自动续订标识，true：是，false：否，默认为false
	AutoRenew *bool `json:"autoRenew,omitempty"`
    // 套餐标识，1：办公型-政企版（不带数据盘），2：办公型-政企版（带数据盘），3：信创型-政企版
	Combo *string `json:"combo,omitempty"`
    // 计费方式，month：包月，year：包年
	MeasureUnit *string `json:"measureUnit,omitempty"`
    // 分类类型，3：办公型-政企版；信创型-政企版，5：办公型-政企版（边缘节点）；信创型-政企版（边缘节点），默认为3
	CategoryCode *string `json:"categoryCode,omitempty"`
    // 额外数据盘数量，单位GB
	DataDiskCount *string `json:"dataDiskCount,omitempty"`
    // 网络工作区编码
	DcCode *string `json:"dcCode,omitempty"`
    // 规格类型，0：2vCPU/4GB内存/80GB系统盘/50M带宽，1：4vCPU/8GB内存/80GB系统盘/50M带宽，2：8vCPU/16GB内存/80GB系统盘/50M带宽，4：4vCPU/4GB内存/80GB系统盘/50M带宽，5：8vCPU/8GB内存/80GB系统盘/50M带宽，6：16vCPU/16GB内存/80GB系统盘/50M带宽，7：16vCPU/32GB内存/80GB系统盘/50M带宽
	SpecificationCode *string `json:"specificationCode,omitempty"`
    // 可用区编码。示例值：CIDC-CORE-00
	ZoneCode *string `json:"zoneCode,omitempty"`
}

func (s SubmitCcaOrderGovBody) String() string {
	return utils.Beautify(s)
}

func (s SubmitCcaOrderGovBody) GoString() string {
	return s.String()
}

func (s SubmitCcaOrderGovBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *SubmitCcaOrderGovBody) SetPeriodNum(v string) *SubmitCcaOrderGovBody {
	s.PeriodNum = &v
	return s
}

func (s *SubmitCcaOrderGovBody) SetIsOrderAdvancedPackage(v bool) *SubmitCcaOrderGovBody {
	s.IsOrderAdvancedPackage = &v
	return s
}

func (s *SubmitCcaOrderGovBody) SetGroupCount(v string) *SubmitCcaOrderGovBody {
	s.GroupCount = &v
	return s
}

func (s *SubmitCcaOrderGovBody) SetImageCode(v string) *SubmitCcaOrderGovBody {
	s.ImageCode = &v
	return s
}

func (s *SubmitCcaOrderGovBody) SetAutoRenew(v bool) *SubmitCcaOrderGovBody {
	s.AutoRenew = &v
	return s
}

func (s *SubmitCcaOrderGovBody) SetCombo(v string) *SubmitCcaOrderGovBody {
	s.Combo = &v
	return s
}

func (s *SubmitCcaOrderGovBody) SetMeasureUnit(v string) *SubmitCcaOrderGovBody {
	s.MeasureUnit = &v
	return s
}

func (s *SubmitCcaOrderGovBody) SetCategoryCode(v string) *SubmitCcaOrderGovBody {
	s.CategoryCode = &v
	return s
}

func (s *SubmitCcaOrderGovBody) SetDataDiskCount(v string) *SubmitCcaOrderGovBody {
	s.DataDiskCount = &v
	return s
}

func (s *SubmitCcaOrderGovBody) SetDcCode(v string) *SubmitCcaOrderGovBody {
	s.DcCode = &v
	return s
}

func (s *SubmitCcaOrderGovBody) SetSpecificationCode(v string) *SubmitCcaOrderGovBody {
	s.SpecificationCode = &v
	return s
}

func (s *SubmitCcaOrderGovBody) SetZoneCode(v string) *SubmitCcaOrderGovBody {
	s.ZoneCode = &v
	return s
}


type SubmitCcaOrderGovBodyBuilder struct {
	s *SubmitCcaOrderGovBody
}

func NewSubmitCcaOrderGovBodyBuilder() *SubmitCcaOrderGovBodyBuilder {
	s := &SubmitCcaOrderGovBody{}
	b := &SubmitCcaOrderGovBodyBuilder{s: s}
	return b
}

func (b *SubmitCcaOrderGovBodyBuilder) PeriodNum(v string) *SubmitCcaOrderGovBodyBuilder {
	b.s.PeriodNum = &v
	return b
}

func (b *SubmitCcaOrderGovBodyBuilder) IsOrderAdvancedPackage(v bool) *SubmitCcaOrderGovBodyBuilder {
	b.s.IsOrderAdvancedPackage = &v
	return b
}

func (b *SubmitCcaOrderGovBodyBuilder) GroupCount(v string) *SubmitCcaOrderGovBodyBuilder {
	b.s.GroupCount = &v
	return b
}

func (b *SubmitCcaOrderGovBodyBuilder) ImageCode(v string) *SubmitCcaOrderGovBodyBuilder {
	b.s.ImageCode = &v
	return b
}

func (b *SubmitCcaOrderGovBodyBuilder) AutoRenew(v bool) *SubmitCcaOrderGovBodyBuilder {
	b.s.AutoRenew = &v
	return b
}

func (b *SubmitCcaOrderGovBodyBuilder) Combo(v string) *SubmitCcaOrderGovBodyBuilder {
	b.s.Combo = &v
	return b
}

func (b *SubmitCcaOrderGovBodyBuilder) MeasureUnit(v string) *SubmitCcaOrderGovBodyBuilder {
	b.s.MeasureUnit = &v
	return b
}

func (b *SubmitCcaOrderGovBodyBuilder) CategoryCode(v string) *SubmitCcaOrderGovBodyBuilder {
	b.s.CategoryCode = &v
	return b
}

func (b *SubmitCcaOrderGovBodyBuilder) DataDiskCount(v string) *SubmitCcaOrderGovBodyBuilder {
	b.s.DataDiskCount = &v
	return b
}

func (b *SubmitCcaOrderGovBodyBuilder) DcCode(v string) *SubmitCcaOrderGovBodyBuilder {
	b.s.DcCode = &v
	return b
}

func (b *SubmitCcaOrderGovBodyBuilder) SpecificationCode(v string) *SubmitCcaOrderGovBodyBuilder {
	b.s.SpecificationCode = &v
	return b
}

func (b *SubmitCcaOrderGovBodyBuilder) ZoneCode(v string) *SubmitCcaOrderGovBodyBuilder {
	b.s.ZoneCode = &v
	return b
}

func (b *SubmitCcaOrderGovBodyBuilder) Build() *SubmitCcaOrderGovBody {
	return b.s
}


