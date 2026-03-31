// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ProductsGovResponseDcs struct {

    // 厂商编码
	CompanyCode *string `json:"companyCode,omitempty"`
    // 网络工作区数量
	Num *int32 `json:"num,omitempty"`
    // 此dc是否支持高级管理服务包，true：是，false：否，默认为false
	IsSupportAdvancedPackage *bool `json:"isSupportAdvancedPackage,omitempty"`
    // 网络工作区编码
	Dc *string `json:"dc,omitempty"`
    // 网络工作区下操作系统信息
	ImageTypeNums *[]ProductsGovResponseImageTypeNums `json:"imageTypeNums,omitempty"`
}

func (s ProductsGovResponseDcs) String() string {
	return utils.Beautify(s)
}

func (s ProductsGovResponseDcs) GoString() string {
	return s.String()
}

func (s ProductsGovResponseDcs) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ProductsGovResponseDcs) SetCompanyCode(v string) *ProductsGovResponseDcs {
	s.CompanyCode = &v
	return s
}

func (s *ProductsGovResponseDcs) SetNum(v int32) *ProductsGovResponseDcs {
	s.Num = &v
	return s
}

func (s *ProductsGovResponseDcs) SetIsSupportAdvancedPackage(v bool) *ProductsGovResponseDcs {
	s.IsSupportAdvancedPackage = &v
	return s
}

func (s *ProductsGovResponseDcs) SetDc(v string) *ProductsGovResponseDcs {
	s.Dc = &v
	return s
}

func (s *ProductsGovResponseDcs) SetImageTypeNums(v []ProductsGovResponseImageTypeNums) *ProductsGovResponseDcs {
	s.ImageTypeNums = &v
	return s
}


type ProductsGovResponseDcsBuilder struct {
	s *ProductsGovResponseDcs
}

func NewProductsGovResponseDcsBuilder() *ProductsGovResponseDcsBuilder {
	s := &ProductsGovResponseDcs{}
	b := &ProductsGovResponseDcsBuilder{s: s}
	return b
}

func (b *ProductsGovResponseDcsBuilder) CompanyCode(v string) *ProductsGovResponseDcsBuilder {
	b.s.CompanyCode = &v
	return b
}

func (b *ProductsGovResponseDcsBuilder) Num(v int32) *ProductsGovResponseDcsBuilder {
	b.s.Num = &v
	return b
}

func (b *ProductsGovResponseDcsBuilder) IsSupportAdvancedPackage(v bool) *ProductsGovResponseDcsBuilder {
	b.s.IsSupportAdvancedPackage = &v
	return b
}

func (b *ProductsGovResponseDcsBuilder) Dc(v string) *ProductsGovResponseDcsBuilder {
	b.s.Dc = &v
	return b
}

func (b *ProductsGovResponseDcsBuilder) ImageTypeNums(v []ProductsGovResponseImageTypeNums) *ProductsGovResponseDcsBuilder {
	b.s.ImageTypeNums = &v
	return b
}

func (b *ProductsGovResponseDcsBuilder) Build() *ProductsGovResponseDcs {
	return b.s
}


