// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ProductsGovResponseData struct {

    // 资源池ID
	PoolId *string `json:"poolId,omitempty"`
    // 资源池区域
	PoolArea *string `json:"poolArea,omitempty"`
    // 可用区集合
	Zones *[]ProductsGovResponseZones `json:"zones,omitempty"`
    // 资源池状态，1：上架，0：下架
	ChaStatus *string `json:"chaStatus,omitempty"`
    // 资源池名称
	PoolName *string `json:"poolName,omitempty"`
}

func (s ProductsGovResponseData) String() string {
	return utils.Beautify(s)
}

func (s ProductsGovResponseData) GoString() string {
	return s.String()
}

func (s ProductsGovResponseData) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ProductsGovResponseData) SetPoolId(v string) *ProductsGovResponseData {
	s.PoolId = &v
	return s
}

func (s *ProductsGovResponseData) SetPoolArea(v string) *ProductsGovResponseData {
	s.PoolArea = &v
	return s
}

func (s *ProductsGovResponseData) SetZones(v []ProductsGovResponseZones) *ProductsGovResponseData {
	s.Zones = &v
	return s
}

func (s *ProductsGovResponseData) SetChaStatus(v string) *ProductsGovResponseData {
	s.ChaStatus = &v
	return s
}

func (s *ProductsGovResponseData) SetPoolName(v string) *ProductsGovResponseData {
	s.PoolName = &v
	return s
}


type ProductsGovResponseDataBuilder struct {
	s *ProductsGovResponseData
}

func NewProductsGovResponseDataBuilder() *ProductsGovResponseDataBuilder {
	s := &ProductsGovResponseData{}
	b := &ProductsGovResponseDataBuilder{s: s}
	return b
}

func (b *ProductsGovResponseDataBuilder) PoolId(v string) *ProductsGovResponseDataBuilder {
	b.s.PoolId = &v
	return b
}

func (b *ProductsGovResponseDataBuilder) PoolArea(v string) *ProductsGovResponseDataBuilder {
	b.s.PoolArea = &v
	return b
}

func (b *ProductsGovResponseDataBuilder) Zones(v []ProductsGovResponseZones) *ProductsGovResponseDataBuilder {
	b.s.Zones = &v
	return b
}

func (b *ProductsGovResponseDataBuilder) ChaStatus(v string) *ProductsGovResponseDataBuilder {
	b.s.ChaStatus = &v
	return b
}

func (b *ProductsGovResponseDataBuilder) PoolName(v string) *ProductsGovResponseDataBuilder {
	b.s.PoolName = &v
	return b
}

func (b *ProductsGovResponseDataBuilder) Build() *ProductsGovResponseData {
	return b.s
}


