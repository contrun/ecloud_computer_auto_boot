// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListSoftwareUnifyApiResponseZones struct {

    // 可用区ID
	ZoneId *string `json:"zoneId,omitempty"`
    // 可用区名称
	ZoneName *string `json:"zoneName,omitempty"`
    // 可用区状态，1：上架，0：下架
	ZoneStatus *string `json:"zoneStatus,omitempty"`
    // 产品规格集合
	Specifications *[]ListSoftwareUnifyApiResponseSpecifications `json:"specifications,omitempty"`
    // 可用区编码
	ZoneCode *string `json:"zoneCode,omitempty"`
}

func (s ListSoftwareUnifyApiResponseZones) String() string {
	return utils.Beautify(s)
}

func (s ListSoftwareUnifyApiResponseZones) GoString() string {
	return s.String()
}

func (s ListSoftwareUnifyApiResponseZones) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListSoftwareUnifyApiResponseZones) SetZoneId(v string) *ListSoftwareUnifyApiResponseZones {
	s.ZoneId = &v
	return s
}

func (s *ListSoftwareUnifyApiResponseZones) SetZoneName(v string) *ListSoftwareUnifyApiResponseZones {
	s.ZoneName = &v
	return s
}

func (s *ListSoftwareUnifyApiResponseZones) SetZoneStatus(v string) *ListSoftwareUnifyApiResponseZones {
	s.ZoneStatus = &v
	return s
}

func (s *ListSoftwareUnifyApiResponseZones) SetSpecifications(v []ListSoftwareUnifyApiResponseSpecifications) *ListSoftwareUnifyApiResponseZones {
	s.Specifications = &v
	return s
}

func (s *ListSoftwareUnifyApiResponseZones) SetZoneCode(v string) *ListSoftwareUnifyApiResponseZones {
	s.ZoneCode = &v
	return s
}


type ListSoftwareUnifyApiResponseZonesBuilder struct {
	s *ListSoftwareUnifyApiResponseZones
}

func NewListSoftwareUnifyApiResponseZonesBuilder() *ListSoftwareUnifyApiResponseZonesBuilder {
	s := &ListSoftwareUnifyApiResponseZones{}
	b := &ListSoftwareUnifyApiResponseZonesBuilder{s: s}
	return b
}

func (b *ListSoftwareUnifyApiResponseZonesBuilder) ZoneId(v string) *ListSoftwareUnifyApiResponseZonesBuilder {
	b.s.ZoneId = &v
	return b
}

func (b *ListSoftwareUnifyApiResponseZonesBuilder) ZoneName(v string) *ListSoftwareUnifyApiResponseZonesBuilder {
	b.s.ZoneName = &v
	return b
}

func (b *ListSoftwareUnifyApiResponseZonesBuilder) ZoneStatus(v string) *ListSoftwareUnifyApiResponseZonesBuilder {
	b.s.ZoneStatus = &v
	return b
}

func (b *ListSoftwareUnifyApiResponseZonesBuilder) Specifications(v []ListSoftwareUnifyApiResponseSpecifications) *ListSoftwareUnifyApiResponseZonesBuilder {
	b.s.Specifications = &v
	return b
}

func (b *ListSoftwareUnifyApiResponseZonesBuilder) ZoneCode(v string) *ListSoftwareUnifyApiResponseZonesBuilder {
	b.s.ZoneCode = &v
	return b
}

func (b *ListSoftwareUnifyApiResponseZonesBuilder) Build() *ListSoftwareUnifyApiResponseZones {
	return b.s
}


