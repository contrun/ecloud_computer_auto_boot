// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListSoftwareUnifyApiResponseData struct {

    // 资源池ID
	PoolId *string `json:"poolId,omitempty"`
    // 资源池区域
	PoolArea *string `json:"poolArea,omitempty"`
    // 可用区集合
	Zones *[]ListSoftwareUnifyApiResponseZones `json:"zones,omitempty"`
    // 资源池状态，1：上架，0：下架
	ChaStatus *string `json:"chaStatus,omitempty"`
    // 资源池名称
	PoolName *string `json:"poolName,omitempty"`
}

func (s ListSoftwareUnifyApiResponseData) String() string {
	return utils.Beautify(s)
}

func (s ListSoftwareUnifyApiResponseData) GoString() string {
	return s.String()
}

func (s ListSoftwareUnifyApiResponseData) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListSoftwareUnifyApiResponseData) SetPoolId(v string) *ListSoftwareUnifyApiResponseData {
	s.PoolId = &v
	return s
}

func (s *ListSoftwareUnifyApiResponseData) SetPoolArea(v string) *ListSoftwareUnifyApiResponseData {
	s.PoolArea = &v
	return s
}

func (s *ListSoftwareUnifyApiResponseData) SetZones(v []ListSoftwareUnifyApiResponseZones) *ListSoftwareUnifyApiResponseData {
	s.Zones = &v
	return s
}

func (s *ListSoftwareUnifyApiResponseData) SetChaStatus(v string) *ListSoftwareUnifyApiResponseData {
	s.ChaStatus = &v
	return s
}

func (s *ListSoftwareUnifyApiResponseData) SetPoolName(v string) *ListSoftwareUnifyApiResponseData {
	s.PoolName = &v
	return s
}


type ListSoftwareUnifyApiResponseDataBuilder struct {
	s *ListSoftwareUnifyApiResponseData
}

func NewListSoftwareUnifyApiResponseDataBuilder() *ListSoftwareUnifyApiResponseDataBuilder {
	s := &ListSoftwareUnifyApiResponseData{}
	b := &ListSoftwareUnifyApiResponseDataBuilder{s: s}
	return b
}

func (b *ListSoftwareUnifyApiResponseDataBuilder) PoolId(v string) *ListSoftwareUnifyApiResponseDataBuilder {
	b.s.PoolId = &v
	return b
}

func (b *ListSoftwareUnifyApiResponseDataBuilder) PoolArea(v string) *ListSoftwareUnifyApiResponseDataBuilder {
	b.s.PoolArea = &v
	return b
}

func (b *ListSoftwareUnifyApiResponseDataBuilder) Zones(v []ListSoftwareUnifyApiResponseZones) *ListSoftwareUnifyApiResponseDataBuilder {
	b.s.Zones = &v
	return b
}

func (b *ListSoftwareUnifyApiResponseDataBuilder) ChaStatus(v string) *ListSoftwareUnifyApiResponseDataBuilder {
	b.s.ChaStatus = &v
	return b
}

func (b *ListSoftwareUnifyApiResponseDataBuilder) PoolName(v string) *ListSoftwareUnifyApiResponseDataBuilder {
	b.s.PoolName = &v
	return b
}

func (b *ListSoftwareUnifyApiResponseDataBuilder) Build() *ListSoftwareUnifyApiResponseData {
	return b.s
}


