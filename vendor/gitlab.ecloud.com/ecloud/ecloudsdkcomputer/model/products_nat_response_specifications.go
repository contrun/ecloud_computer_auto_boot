// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ProductsNatResponseSpecifications struct {

    // 规格ID
	ChaId *string `json:"chaId,omitempty"`
    // 规格名称
	Name *string `json:"name,omitempty"`
    // 是否存在免费领取
	IsExistsFree *bool `json:"isExistsFree,omitempty"`
}

func (s ProductsNatResponseSpecifications) String() string {
	return utils.Beautify(s)
}

func (s ProductsNatResponseSpecifications) GoString() string {
	return s.String()
}

func (s ProductsNatResponseSpecifications) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ProductsNatResponseSpecifications) SetChaId(v string) *ProductsNatResponseSpecifications {
	s.ChaId = &v
	return s
}

func (s *ProductsNatResponseSpecifications) SetName(v string) *ProductsNatResponseSpecifications {
	s.Name = &v
	return s
}

func (s *ProductsNatResponseSpecifications) SetIsExistsFree(v bool) *ProductsNatResponseSpecifications {
	s.IsExistsFree = &v
	return s
}


type ProductsNatResponseSpecificationsBuilder struct {
	s *ProductsNatResponseSpecifications
}

func NewProductsNatResponseSpecificationsBuilder() *ProductsNatResponseSpecificationsBuilder {
	s := &ProductsNatResponseSpecifications{}
	b := &ProductsNatResponseSpecificationsBuilder{s: s}
	return b
}

func (b *ProductsNatResponseSpecificationsBuilder) ChaId(v string) *ProductsNatResponseSpecificationsBuilder {
	b.s.ChaId = &v
	return b
}

func (b *ProductsNatResponseSpecificationsBuilder) Name(v string) *ProductsNatResponseSpecificationsBuilder {
	b.s.Name = &v
	return b
}

func (b *ProductsNatResponseSpecificationsBuilder) IsExistsFree(v bool) *ProductsNatResponseSpecificationsBuilder {
	b.s.IsExistsFree = &v
	return b
}

func (b *ProductsNatResponseSpecificationsBuilder) Build() *ProductsNatResponseSpecifications {
	return b.s
}


