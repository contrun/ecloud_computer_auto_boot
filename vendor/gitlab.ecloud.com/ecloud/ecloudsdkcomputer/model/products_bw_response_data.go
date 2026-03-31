// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ProductsBwResponseData struct {

    // 资源池ID
	PoolId *string `json:"poolId,omitempty"`
    // 资源池区域
	PoolArea *string `json:"poolArea,omitempty"`
    // 可用区集合
	Zones *[]ProductsBwResponseZones `json:"zones,omitempty"`
    // 资源池状态，1：上架，0：下架
	ChaStatus *string `json:"chaStatus,omitempty"`
    // 资源池名称
	PoolName *string `json:"poolName,omitempty"`
}

func (s ProductsBwResponseData) String() string {
	return utils.Beautify(s)
}

func (s ProductsBwResponseData) GoString() string {
	return s.String()
}

func (s ProductsBwResponseData) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ProductsBwResponseData) SetPoolId(v string) *ProductsBwResponseData {
	s.PoolId = &v
	return s
}

func (s *ProductsBwResponseData) SetPoolArea(v string) *ProductsBwResponseData {
	s.PoolArea = &v
	return s
}

func (s *ProductsBwResponseData) SetZones(v []ProductsBwResponseZones) *ProductsBwResponseData {
	s.Zones = &v
	return s
}

func (s *ProductsBwResponseData) SetChaStatus(v string) *ProductsBwResponseData {
	s.ChaStatus = &v
	return s
}

func (s *ProductsBwResponseData) SetPoolName(v string) *ProductsBwResponseData {
	s.PoolName = &v
	return s
}


type ProductsBwResponseDataBuilder struct {
	s *ProductsBwResponseData
}

func NewProductsBwResponseDataBuilder() *ProductsBwResponseDataBuilder {
	s := &ProductsBwResponseData{}
	b := &ProductsBwResponseDataBuilder{s: s}
	return b
}

func (b *ProductsBwResponseDataBuilder) PoolId(v string) *ProductsBwResponseDataBuilder {
	b.s.PoolId = &v
	return b
}

func (b *ProductsBwResponseDataBuilder) PoolArea(v string) *ProductsBwResponseDataBuilder {
	b.s.PoolArea = &v
	return b
}

func (b *ProductsBwResponseDataBuilder) Zones(v []ProductsBwResponseZones) *ProductsBwResponseDataBuilder {
	b.s.Zones = &v
	return b
}

func (b *ProductsBwResponseDataBuilder) ChaStatus(v string) *ProductsBwResponseDataBuilder {
	b.s.ChaStatus = &v
	return b
}

func (b *ProductsBwResponseDataBuilder) PoolName(v string) *ProductsBwResponseDataBuilder {
	b.s.PoolName = &v
	return b
}

func (b *ProductsBwResponseDataBuilder) Build() *ProductsBwResponseData {
	return b.s
}


