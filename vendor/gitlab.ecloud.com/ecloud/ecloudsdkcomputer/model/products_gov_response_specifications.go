// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ProductsGovResponseSpecifications struct {

    // 操作系统集合
	Images *[]ProductsGovResponseImages `json:"images,omitempty"`
    // 带宽大小（M）
	Bandwidth *int32 `json:"bandwidth,omitempty"`
    // 可售数量
	Count *int32 `json:"count,omitempty"`
    // 套餐标识，1：办公型-政企版（不带数据盘），2：办公型-政企版（带数据盘），3：信创型-政企版
	Combo *int32 `json:"combo,omitempty"`
    // CPU数量
	Cpu *int32 `json:"cpu,omitempty"`
    // 此规格是否存在高级管理服务包，true：是，false：否，默认为false
	IsExistsAdvancedPackage *bool `json:"isExistsAdvancedPackage,omitempty"`
    // 分类类型，3：办公型-政企版；信创型-政企版，5：办公型-政企版（边缘节点）
	CategoryCode *string `json:"categoryCode,omitempty"`
    // 规格ID
	ChaId *string `json:"chaId,omitempty"`
    // 网络工作区集合
	Dcs *[]ProductsGovResponseDcs `json:"dcs,omitempty"`
    // 系统盘大小（GB）
	SysDisk *int32 `json:"sysDisk,omitempty"`
    // 规格名称
	Name *string `json:"name,omitempty"`
    // 数据盘大小（GB）
	DataDisk *int32 `json:"dataDisk,omitempty"`
    // 规格类型，0：通用规格2vCPU/4GB内存，1：通用规格4vCPU/8GB内存，2：通用规格8vCPU/16GB内存，4：通用规格4vCPU/4GB内存，5：通用规格8vCPU/8GB内存，6：通用规格16vCPU/16GB内存，7：通用规格16vCPU/32GB内存
	SpecificationCode *string `json:"specificationCode,omitempty"`
    // 内存大小（GB）
	Ram *int32 `json:"ram,omitempty"`
}

func (s ProductsGovResponseSpecifications) String() string {
	return utils.Beautify(s)
}

func (s ProductsGovResponseSpecifications) GoString() string {
	return s.String()
}

func (s ProductsGovResponseSpecifications) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ProductsGovResponseSpecifications) SetImages(v []ProductsGovResponseImages) *ProductsGovResponseSpecifications {
	s.Images = &v
	return s
}

func (s *ProductsGovResponseSpecifications) SetBandwidth(v int32) *ProductsGovResponseSpecifications {
	s.Bandwidth = &v
	return s
}

func (s *ProductsGovResponseSpecifications) SetCount(v int32) *ProductsGovResponseSpecifications {
	s.Count = &v
	return s
}

func (s *ProductsGovResponseSpecifications) SetCombo(v int32) *ProductsGovResponseSpecifications {
	s.Combo = &v
	return s
}

func (s *ProductsGovResponseSpecifications) SetCpu(v int32) *ProductsGovResponseSpecifications {
	s.Cpu = &v
	return s
}

func (s *ProductsGovResponseSpecifications) SetIsExistsAdvancedPackage(v bool) *ProductsGovResponseSpecifications {
	s.IsExistsAdvancedPackage = &v
	return s
}

func (s *ProductsGovResponseSpecifications) SetCategoryCode(v string) *ProductsGovResponseSpecifications {
	s.CategoryCode = &v
	return s
}

func (s *ProductsGovResponseSpecifications) SetChaId(v string) *ProductsGovResponseSpecifications {
	s.ChaId = &v
	return s
}

func (s *ProductsGovResponseSpecifications) SetDcs(v []ProductsGovResponseDcs) *ProductsGovResponseSpecifications {
	s.Dcs = &v
	return s
}

func (s *ProductsGovResponseSpecifications) SetSysDisk(v int32) *ProductsGovResponseSpecifications {
	s.SysDisk = &v
	return s
}

func (s *ProductsGovResponseSpecifications) SetName(v string) *ProductsGovResponseSpecifications {
	s.Name = &v
	return s
}

func (s *ProductsGovResponseSpecifications) SetDataDisk(v int32) *ProductsGovResponseSpecifications {
	s.DataDisk = &v
	return s
}

func (s *ProductsGovResponseSpecifications) SetSpecificationCode(v string) *ProductsGovResponseSpecifications {
	s.SpecificationCode = &v
	return s
}

func (s *ProductsGovResponseSpecifications) SetRam(v int32) *ProductsGovResponseSpecifications {
	s.Ram = &v
	return s
}


type ProductsGovResponseSpecificationsBuilder struct {
	s *ProductsGovResponseSpecifications
}

func NewProductsGovResponseSpecificationsBuilder() *ProductsGovResponseSpecificationsBuilder {
	s := &ProductsGovResponseSpecifications{}
	b := &ProductsGovResponseSpecificationsBuilder{s: s}
	return b
}

func (b *ProductsGovResponseSpecificationsBuilder) Images(v []ProductsGovResponseImages) *ProductsGovResponseSpecificationsBuilder {
	b.s.Images = &v
	return b
}

func (b *ProductsGovResponseSpecificationsBuilder) Bandwidth(v int32) *ProductsGovResponseSpecificationsBuilder {
	b.s.Bandwidth = &v
	return b
}

func (b *ProductsGovResponseSpecificationsBuilder) Count(v int32) *ProductsGovResponseSpecificationsBuilder {
	b.s.Count = &v
	return b
}

func (b *ProductsGovResponseSpecificationsBuilder) Combo(v int32) *ProductsGovResponseSpecificationsBuilder {
	b.s.Combo = &v
	return b
}

func (b *ProductsGovResponseSpecificationsBuilder) Cpu(v int32) *ProductsGovResponseSpecificationsBuilder {
	b.s.Cpu = &v
	return b
}

func (b *ProductsGovResponseSpecificationsBuilder) IsExistsAdvancedPackage(v bool) *ProductsGovResponseSpecificationsBuilder {
	b.s.IsExistsAdvancedPackage = &v
	return b
}

func (b *ProductsGovResponseSpecificationsBuilder) CategoryCode(v string) *ProductsGovResponseSpecificationsBuilder {
	b.s.CategoryCode = &v
	return b
}

func (b *ProductsGovResponseSpecificationsBuilder) ChaId(v string) *ProductsGovResponseSpecificationsBuilder {
	b.s.ChaId = &v
	return b
}

func (b *ProductsGovResponseSpecificationsBuilder) Dcs(v []ProductsGovResponseDcs) *ProductsGovResponseSpecificationsBuilder {
	b.s.Dcs = &v
	return b
}

func (b *ProductsGovResponseSpecificationsBuilder) SysDisk(v int32) *ProductsGovResponseSpecificationsBuilder {
	b.s.SysDisk = &v
	return b
}

func (b *ProductsGovResponseSpecificationsBuilder) Name(v string) *ProductsGovResponseSpecificationsBuilder {
	b.s.Name = &v
	return b
}

func (b *ProductsGovResponseSpecificationsBuilder) DataDisk(v int32) *ProductsGovResponseSpecificationsBuilder {
	b.s.DataDisk = &v
	return b
}

func (b *ProductsGovResponseSpecificationsBuilder) SpecificationCode(v string) *ProductsGovResponseSpecificationsBuilder {
	b.s.SpecificationCode = &v
	return b
}

func (b *ProductsGovResponseSpecificationsBuilder) Ram(v int32) *ProductsGovResponseSpecificationsBuilder {
	b.s.Ram = &v
	return b
}

func (b *ProductsGovResponseSpecificationsBuilder) Build() *ProductsGovResponseSpecifications {
	return b.s
}


