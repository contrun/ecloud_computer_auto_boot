// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ProductsBwResponseZones struct {

    // 厂商编码
	CompanyCode *string `json:"companyCode,omitempty"`
    // 可用区ID
	ZoneId *string `json:"zoneId,omitempty"`
    // Mop可用区编码
	ZoneCodeMop *string `json:"zoneCodeMop,omitempty"`
    // 可用区名称
	ZoneName *string `json:"zoneName,omitempty"`
    // 可用区状态，1：上架，0：下架
	ZoneStatus *string `json:"zoneStatus,omitempty"`
    // 产品规格集合
	Specifications *[]ProductsBwResponseSpecifications `json:"specifications,omitempty"`
    // 可用区编码
	ZoneCode *string `json:"zoneCode,omitempty"`
}

func (s ProductsBwResponseZones) String() string {
	return utils.Beautify(s)
}

func (s ProductsBwResponseZones) GoString() string {
	return s.String()
}

func (s ProductsBwResponseZones) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ProductsBwResponseZones) SetCompanyCode(v string) *ProductsBwResponseZones {
	s.CompanyCode = &v
	return s
}

func (s *ProductsBwResponseZones) SetZoneId(v string) *ProductsBwResponseZones {
	s.ZoneId = &v
	return s
}

func (s *ProductsBwResponseZones) SetZoneCodeMop(v string) *ProductsBwResponseZones {
	s.ZoneCodeMop = &v
	return s
}

func (s *ProductsBwResponseZones) SetZoneName(v string) *ProductsBwResponseZones {
	s.ZoneName = &v
	return s
}

func (s *ProductsBwResponseZones) SetZoneStatus(v string) *ProductsBwResponseZones {
	s.ZoneStatus = &v
	return s
}

func (s *ProductsBwResponseZones) SetSpecifications(v []ProductsBwResponseSpecifications) *ProductsBwResponseZones {
	s.Specifications = &v
	return s
}

func (s *ProductsBwResponseZones) SetZoneCode(v string) *ProductsBwResponseZones {
	s.ZoneCode = &v
	return s
}


type ProductsBwResponseZonesBuilder struct {
	s *ProductsBwResponseZones
}

func NewProductsBwResponseZonesBuilder() *ProductsBwResponseZonesBuilder {
	s := &ProductsBwResponseZones{}
	b := &ProductsBwResponseZonesBuilder{s: s}
	return b
}

func (b *ProductsBwResponseZonesBuilder) CompanyCode(v string) *ProductsBwResponseZonesBuilder {
	b.s.CompanyCode = &v
	return b
}

func (b *ProductsBwResponseZonesBuilder) ZoneId(v string) *ProductsBwResponseZonesBuilder {
	b.s.ZoneId = &v
	return b
}

func (b *ProductsBwResponseZonesBuilder) ZoneCodeMop(v string) *ProductsBwResponseZonesBuilder {
	b.s.ZoneCodeMop = &v
	return b
}

func (b *ProductsBwResponseZonesBuilder) ZoneName(v string) *ProductsBwResponseZonesBuilder {
	b.s.ZoneName = &v
	return b
}

func (b *ProductsBwResponseZonesBuilder) ZoneStatus(v string) *ProductsBwResponseZonesBuilder {
	b.s.ZoneStatus = &v
	return b
}

func (b *ProductsBwResponseZonesBuilder) Specifications(v []ProductsBwResponseSpecifications) *ProductsBwResponseZonesBuilder {
	b.s.Specifications = &v
	return b
}

func (b *ProductsBwResponseZonesBuilder) ZoneCode(v string) *ProductsBwResponseZonesBuilder {
	b.s.ZoneCode = &v
	return b
}

func (b *ProductsBwResponseZonesBuilder) Build() *ProductsBwResponseZones {
	return b.s
}


