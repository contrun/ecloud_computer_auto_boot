// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ProductsNatResponseData struct {

    // 资源池ID
	PoolId *string `json:"poolId,omitempty"`
    // 资源池区域
	PoolArea *string `json:"poolArea,omitempty"`
    // 可用区集合
	Zones *[]ProductsNatResponseZones `json:"zones,omitempty"`
    // 资源池状态，1：上架，0：下架
	ChaStatus *string `json:"chaStatus,omitempty"`
    // 资源池名称
	PoolName *string `json:"poolName,omitempty"`
}

func (s ProductsNatResponseData) String() string {
	return utils.Beautify(s)
}

func (s ProductsNatResponseData) GoString() string {
	return s.String()
}

func (s ProductsNatResponseData) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ProductsNatResponseData) SetPoolId(v string) *ProductsNatResponseData {
	s.PoolId = &v
	return s
}

func (s *ProductsNatResponseData) SetPoolArea(v string) *ProductsNatResponseData {
	s.PoolArea = &v
	return s
}

func (s *ProductsNatResponseData) SetZones(v []ProductsNatResponseZones) *ProductsNatResponseData {
	s.Zones = &v
	return s
}

func (s *ProductsNatResponseData) SetChaStatus(v string) *ProductsNatResponseData {
	s.ChaStatus = &v
	return s
}

func (s *ProductsNatResponseData) SetPoolName(v string) *ProductsNatResponseData {
	s.PoolName = &v
	return s
}


type ProductsNatResponseDataBuilder struct {
	s *ProductsNatResponseData
}

func NewProductsNatResponseDataBuilder() *ProductsNatResponseDataBuilder {
	s := &ProductsNatResponseData{}
	b := &ProductsNatResponseDataBuilder{s: s}
	return b
}

func (b *ProductsNatResponseDataBuilder) PoolId(v string) *ProductsNatResponseDataBuilder {
	b.s.PoolId = &v
	return b
}

func (b *ProductsNatResponseDataBuilder) PoolArea(v string) *ProductsNatResponseDataBuilder {
	b.s.PoolArea = &v
	return b
}

func (b *ProductsNatResponseDataBuilder) Zones(v []ProductsNatResponseZones) *ProductsNatResponseDataBuilder {
	b.s.Zones = &v
	return b
}

func (b *ProductsNatResponseDataBuilder) ChaStatus(v string) *ProductsNatResponseDataBuilder {
	b.s.ChaStatus = &v
	return b
}

func (b *ProductsNatResponseDataBuilder) PoolName(v string) *ProductsNatResponseDataBuilder {
	b.s.PoolName = &v
	return b
}

func (b *ProductsNatResponseDataBuilder) Build() *ProductsNatResponseData {
	return b.s
}


