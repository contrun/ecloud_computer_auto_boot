// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ProductsGovResponseZones struct {

    // 可用区ID
	ZoneId *string `json:"zoneId,omitempty"`
    // 可用区名称
	ZoneName *string `json:"zoneName,omitempty"`
    // 可用区状态，1：上架，0：下架
	ZoneStatus *string `json:"zoneStatus,omitempty"`
    // 产品规格集合
	Specifications *[]ProductsGovResponseSpecifications `json:"specifications,omitempty"`
    // 可用区编码
	ZoneCode *string `json:"zoneCode,omitempty"`
}

func (s ProductsGovResponseZones) String() string {
	return utils.Beautify(s)
}

func (s ProductsGovResponseZones) GoString() string {
	return s.String()
}

func (s ProductsGovResponseZones) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ProductsGovResponseZones) SetZoneId(v string) *ProductsGovResponseZones {
	s.ZoneId = &v
	return s
}

func (s *ProductsGovResponseZones) SetZoneName(v string) *ProductsGovResponseZones {
	s.ZoneName = &v
	return s
}

func (s *ProductsGovResponseZones) SetZoneStatus(v string) *ProductsGovResponseZones {
	s.ZoneStatus = &v
	return s
}

func (s *ProductsGovResponseZones) SetSpecifications(v []ProductsGovResponseSpecifications) *ProductsGovResponseZones {
	s.Specifications = &v
	return s
}

func (s *ProductsGovResponseZones) SetZoneCode(v string) *ProductsGovResponseZones {
	s.ZoneCode = &v
	return s
}


type ProductsGovResponseZonesBuilder struct {
	s *ProductsGovResponseZones
}

func NewProductsGovResponseZonesBuilder() *ProductsGovResponseZonesBuilder {
	s := &ProductsGovResponseZones{}
	b := &ProductsGovResponseZonesBuilder{s: s}
	return b
}

func (b *ProductsGovResponseZonesBuilder) ZoneId(v string) *ProductsGovResponseZonesBuilder {
	b.s.ZoneId = &v
	return b
}

func (b *ProductsGovResponseZonesBuilder) ZoneName(v string) *ProductsGovResponseZonesBuilder {
	b.s.ZoneName = &v
	return b
}

func (b *ProductsGovResponseZonesBuilder) ZoneStatus(v string) *ProductsGovResponseZonesBuilder {
	b.s.ZoneStatus = &v
	return b
}

func (b *ProductsGovResponseZonesBuilder) Specifications(v []ProductsGovResponseSpecifications) *ProductsGovResponseZonesBuilder {
	b.s.Specifications = &v
	return b
}

func (b *ProductsGovResponseZonesBuilder) ZoneCode(v string) *ProductsGovResponseZonesBuilder {
	b.s.ZoneCode = &v
	return b
}

func (b *ProductsGovResponseZonesBuilder) Build() *ProductsGovResponseZones {
	return b.s
}


