// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListSoftwareUnifyApiResponseSpecifications struct {

    // 操作系统集合
	Images *[]ListSoftwareUnifyApiResponseImages `json:"images,omitempty"`
    // 带宽大小（M）
	Bandwidth *int32 `json:"bandwidth,omitempty"`
    // 可售数量
	Count *int32 `json:"count,omitempty"`
    // 套餐标识，1：办公型-公众版（不带数据盘）、办公型-政企版（不带数据盘），2：办公型-公众版（带数据盘）、办公型-政企版（带数据盘），3：信创型-政企版，5：办公型-政企版-尊享规格
	Combo *int32 `json:"combo,omitempty"`
    // CPU数量
	Cpu *int32 `json:"cpu,omitempty"`
    // 分类类型，1：办公型-公众版，3：办公型-政企版；信创型-政企版，5：办公型-政企版（边缘节点）
	CategoryCode *string `json:"categoryCode,omitempty"`
    // 规格ID
	ChaId *string `json:"chaId,omitempty"`
    // 网络工作区集合
	Dcs *[]ListSoftwareUnifyApiResponseDcs `json:"dcs,omitempty"`
    // 系统盘大小（GB）
	SysDisk *int32 `json:"sysDisk,omitempty"`
    // 规格名称
	Name *string `json:"name,omitempty"`
    // 数据盘大小（GB）
	DataDisk *int32 `json:"dataDisk,omitempty"`
    // 规格类型，0：通用规格2vCPU/4GB内存，1：通用规格4vCPU/8GB内存，2：通用规格8vCPU/16GB内存，3：通用规格6vCPU/12GB内存，4：通用规格4vCPU/4GB内存，5：通用规格8vCPU/8GB内存，6：通用规格16vCPU/16GB内存，7：通用规格16vCPU/32GB内存套餐标识，1：办公型-公众版（不带数据盘）、办公型-政企版（不带数据盘），2：办公型-公众版（带数据盘）、办公型-政企版（带数据盘），3：信创型-政企版，5：办公型-政企版-尊享规格
	SpecificationCode *string `json:"specificationCode,omitempty"`
    // 内存大小（GB）
	Ram *int32 `json:"ram,omitempty"`
}

func (s ListSoftwareUnifyApiResponseSpecifications) String() string {
	return utils.Beautify(s)
}

func (s ListSoftwareUnifyApiResponseSpecifications) GoString() string {
	return s.String()
}

func (s ListSoftwareUnifyApiResponseSpecifications) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListSoftwareUnifyApiResponseSpecifications) SetImages(v []ListSoftwareUnifyApiResponseImages) *ListSoftwareUnifyApiResponseSpecifications {
	s.Images = &v
	return s
}

func (s *ListSoftwareUnifyApiResponseSpecifications) SetBandwidth(v int32) *ListSoftwareUnifyApiResponseSpecifications {
	s.Bandwidth = &v
	return s
}

func (s *ListSoftwareUnifyApiResponseSpecifications) SetCount(v int32) *ListSoftwareUnifyApiResponseSpecifications {
	s.Count = &v
	return s
}

func (s *ListSoftwareUnifyApiResponseSpecifications) SetCombo(v int32) *ListSoftwareUnifyApiResponseSpecifications {
	s.Combo = &v
	return s
}

func (s *ListSoftwareUnifyApiResponseSpecifications) SetCpu(v int32) *ListSoftwareUnifyApiResponseSpecifications {
	s.Cpu = &v
	return s
}

func (s *ListSoftwareUnifyApiResponseSpecifications) SetCategoryCode(v string) *ListSoftwareUnifyApiResponseSpecifications {
	s.CategoryCode = &v
	return s
}

func (s *ListSoftwareUnifyApiResponseSpecifications) SetChaId(v string) *ListSoftwareUnifyApiResponseSpecifications {
	s.ChaId = &v
	return s
}

func (s *ListSoftwareUnifyApiResponseSpecifications) SetDcs(v []ListSoftwareUnifyApiResponseDcs) *ListSoftwareUnifyApiResponseSpecifications {
	s.Dcs = &v
	return s
}

func (s *ListSoftwareUnifyApiResponseSpecifications) SetSysDisk(v int32) *ListSoftwareUnifyApiResponseSpecifications {
	s.SysDisk = &v
	return s
}

func (s *ListSoftwareUnifyApiResponseSpecifications) SetName(v string) *ListSoftwareUnifyApiResponseSpecifications {
	s.Name = &v
	return s
}

func (s *ListSoftwareUnifyApiResponseSpecifications) SetDataDisk(v int32) *ListSoftwareUnifyApiResponseSpecifications {
	s.DataDisk = &v
	return s
}

func (s *ListSoftwareUnifyApiResponseSpecifications) SetSpecificationCode(v string) *ListSoftwareUnifyApiResponseSpecifications {
	s.SpecificationCode = &v
	return s
}

func (s *ListSoftwareUnifyApiResponseSpecifications) SetRam(v int32) *ListSoftwareUnifyApiResponseSpecifications {
	s.Ram = &v
	return s
}


type ListSoftwareUnifyApiResponseSpecificationsBuilder struct {
	s *ListSoftwareUnifyApiResponseSpecifications
}

func NewListSoftwareUnifyApiResponseSpecificationsBuilder() *ListSoftwareUnifyApiResponseSpecificationsBuilder {
	s := &ListSoftwareUnifyApiResponseSpecifications{}
	b := &ListSoftwareUnifyApiResponseSpecificationsBuilder{s: s}
	return b
}

func (b *ListSoftwareUnifyApiResponseSpecificationsBuilder) Images(v []ListSoftwareUnifyApiResponseImages) *ListSoftwareUnifyApiResponseSpecificationsBuilder {
	b.s.Images = &v
	return b
}

func (b *ListSoftwareUnifyApiResponseSpecificationsBuilder) Bandwidth(v int32) *ListSoftwareUnifyApiResponseSpecificationsBuilder {
	b.s.Bandwidth = &v
	return b
}

func (b *ListSoftwareUnifyApiResponseSpecificationsBuilder) Count(v int32) *ListSoftwareUnifyApiResponseSpecificationsBuilder {
	b.s.Count = &v
	return b
}

func (b *ListSoftwareUnifyApiResponseSpecificationsBuilder) Combo(v int32) *ListSoftwareUnifyApiResponseSpecificationsBuilder {
	b.s.Combo = &v
	return b
}

func (b *ListSoftwareUnifyApiResponseSpecificationsBuilder) Cpu(v int32) *ListSoftwareUnifyApiResponseSpecificationsBuilder {
	b.s.Cpu = &v
	return b
}

func (b *ListSoftwareUnifyApiResponseSpecificationsBuilder) CategoryCode(v string) *ListSoftwareUnifyApiResponseSpecificationsBuilder {
	b.s.CategoryCode = &v
	return b
}

func (b *ListSoftwareUnifyApiResponseSpecificationsBuilder) ChaId(v string) *ListSoftwareUnifyApiResponseSpecificationsBuilder {
	b.s.ChaId = &v
	return b
}

func (b *ListSoftwareUnifyApiResponseSpecificationsBuilder) Dcs(v []ListSoftwareUnifyApiResponseDcs) *ListSoftwareUnifyApiResponseSpecificationsBuilder {
	b.s.Dcs = &v
	return b
}

func (b *ListSoftwareUnifyApiResponseSpecificationsBuilder) SysDisk(v int32) *ListSoftwareUnifyApiResponseSpecificationsBuilder {
	b.s.SysDisk = &v
	return b
}

func (b *ListSoftwareUnifyApiResponseSpecificationsBuilder) Name(v string) *ListSoftwareUnifyApiResponseSpecificationsBuilder {
	b.s.Name = &v
	return b
}

func (b *ListSoftwareUnifyApiResponseSpecificationsBuilder) DataDisk(v int32) *ListSoftwareUnifyApiResponseSpecificationsBuilder {
	b.s.DataDisk = &v
	return b
}

func (b *ListSoftwareUnifyApiResponseSpecificationsBuilder) SpecificationCode(v string) *ListSoftwareUnifyApiResponseSpecificationsBuilder {
	b.s.SpecificationCode = &v
	return b
}

func (b *ListSoftwareUnifyApiResponseSpecificationsBuilder) Ram(v int32) *ListSoftwareUnifyApiResponseSpecificationsBuilder {
	b.s.Ram = &v
	return b
}

func (b *ListSoftwareUnifyApiResponseSpecificationsBuilder) Build() *ListSoftwareUnifyApiResponseSpecifications {
	return b.s
}


