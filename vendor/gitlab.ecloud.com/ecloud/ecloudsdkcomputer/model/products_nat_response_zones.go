// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ProductsNatResponseZones struct {

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
	Specifications *[]ProductsNatResponseSpecifications `json:"specifications,omitempty"`
    // 可用区编码
	ZoneCode *string `json:"zoneCode,omitempty"`
}

func (s ProductsNatResponseZones) String() string {
	return utils.Beautify(s)
}

func (s ProductsNatResponseZones) GoString() string {
	return s.String()
}

func (s ProductsNatResponseZones) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ProductsNatResponseZones) SetCompanyCode(v string) *ProductsNatResponseZones {
	s.CompanyCode = &v
	return s
}

func (s *ProductsNatResponseZones) SetZoneId(v string) *ProductsNatResponseZones {
	s.ZoneId = &v
	return s
}

func (s *ProductsNatResponseZones) SetZoneCodeMop(v string) *ProductsNatResponseZones {
	s.ZoneCodeMop = &v
	return s
}

func (s *ProductsNatResponseZones) SetZoneName(v string) *ProductsNatResponseZones {
	s.ZoneName = &v
	return s
}

func (s *ProductsNatResponseZones) SetZoneStatus(v string) *ProductsNatResponseZones {
	s.ZoneStatus = &v
	return s
}

func (s *ProductsNatResponseZones) SetSpecifications(v []ProductsNatResponseSpecifications) *ProductsNatResponseZones {
	s.Specifications = &v
	return s
}

func (s *ProductsNatResponseZones) SetZoneCode(v string) *ProductsNatResponseZones {
	s.ZoneCode = &v
	return s
}


type ProductsNatResponseZonesBuilder struct {
	s *ProductsNatResponseZones
}

func NewProductsNatResponseZonesBuilder() *ProductsNatResponseZonesBuilder {
	s := &ProductsNatResponseZones{}
	b := &ProductsNatResponseZonesBuilder{s: s}
	return b
}

func (b *ProductsNatResponseZonesBuilder) CompanyCode(v string) *ProductsNatResponseZonesBuilder {
	b.s.CompanyCode = &v
	return b
}

func (b *ProductsNatResponseZonesBuilder) ZoneId(v string) *ProductsNatResponseZonesBuilder {
	b.s.ZoneId = &v
	return b
}

func (b *ProductsNatResponseZonesBuilder) ZoneCodeMop(v string) *ProductsNatResponseZonesBuilder {
	b.s.ZoneCodeMop = &v
	return b
}

func (b *ProductsNatResponseZonesBuilder) ZoneName(v string) *ProductsNatResponseZonesBuilder {
	b.s.ZoneName = &v
	return b
}

func (b *ProductsNatResponseZonesBuilder) ZoneStatus(v string) *ProductsNatResponseZonesBuilder {
	b.s.ZoneStatus = &v
	return b
}

func (b *ProductsNatResponseZonesBuilder) Specifications(v []ProductsNatResponseSpecifications) *ProductsNatResponseZonesBuilder {
	b.s.Specifications = &v
	return b
}

func (b *ProductsNatResponseZonesBuilder) ZoneCode(v string) *ProductsNatResponseZonesBuilder {
	b.s.ZoneCode = &v
	return b
}

func (b *ProductsNatResponseZonesBuilder) Build() *ProductsNatResponseZones {
	return b.s
}


