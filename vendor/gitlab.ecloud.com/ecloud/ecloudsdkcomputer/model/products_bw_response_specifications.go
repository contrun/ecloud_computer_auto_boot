// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ProductsBwResponseSpecifications struct {

    // 规格ID
	ChaId *string `json:"chaId,omitempty"`
    // 规格名称
	Name *string `json:"name,omitempty"`
    // 是否存在免费领取
	IsExistsFree *bool `json:"isExistsFree,omitempty"`
}

func (s ProductsBwResponseSpecifications) String() string {
	return utils.Beautify(s)
}

func (s ProductsBwResponseSpecifications) GoString() string {
	return s.String()
}

func (s ProductsBwResponseSpecifications) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ProductsBwResponseSpecifications) SetChaId(v string) *ProductsBwResponseSpecifications {
	s.ChaId = &v
	return s
}

func (s *ProductsBwResponseSpecifications) SetName(v string) *ProductsBwResponseSpecifications {
	s.Name = &v
	return s
}

func (s *ProductsBwResponseSpecifications) SetIsExistsFree(v bool) *ProductsBwResponseSpecifications {
	s.IsExistsFree = &v
	return s
}


type ProductsBwResponseSpecificationsBuilder struct {
	s *ProductsBwResponseSpecifications
}

func NewProductsBwResponseSpecificationsBuilder() *ProductsBwResponseSpecificationsBuilder {
	s := &ProductsBwResponseSpecifications{}
	b := &ProductsBwResponseSpecificationsBuilder{s: s}
	return b
}

func (b *ProductsBwResponseSpecificationsBuilder) ChaId(v string) *ProductsBwResponseSpecificationsBuilder {
	b.s.ChaId = &v
	return b
}

func (b *ProductsBwResponseSpecificationsBuilder) Name(v string) *ProductsBwResponseSpecificationsBuilder {
	b.s.Name = &v
	return b
}

func (b *ProductsBwResponseSpecificationsBuilder) IsExistsFree(v bool) *ProductsBwResponseSpecificationsBuilder {
	b.s.IsExistsFree = &v
	return b
}

func (b *ProductsBwResponseSpecificationsBuilder) Build() *ProductsBwResponseSpecifications {
	return b.s
}


