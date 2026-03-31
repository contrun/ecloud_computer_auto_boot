// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ProductsGovResponseImageTypeNums struct {

    // 网络工作区下的操作系统可售数量
	Num *int32 `json:"num,omitempty"`
    // 操作系统标识，UOS：统信UOS_1050，Kylin：麒麟KylinOS_V10，Windows：Windows 10，Windows11：Windows 11，WinS2016：Windows Server 2016
	ImageType *string `json:"imageType,omitempty"`
}

func (s ProductsGovResponseImageTypeNums) String() string {
	return utils.Beautify(s)
}

func (s ProductsGovResponseImageTypeNums) GoString() string {
	return s.String()
}

func (s ProductsGovResponseImageTypeNums) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ProductsGovResponseImageTypeNums) SetNum(v int32) *ProductsGovResponseImageTypeNums {
	s.Num = &v
	return s
}

func (s *ProductsGovResponseImageTypeNums) SetImageType(v string) *ProductsGovResponseImageTypeNums {
	s.ImageType = &v
	return s
}


type ProductsGovResponseImageTypeNumsBuilder struct {
	s *ProductsGovResponseImageTypeNums
}

func NewProductsGovResponseImageTypeNumsBuilder() *ProductsGovResponseImageTypeNumsBuilder {
	s := &ProductsGovResponseImageTypeNums{}
	b := &ProductsGovResponseImageTypeNumsBuilder{s: s}
	return b
}

func (b *ProductsGovResponseImageTypeNumsBuilder) Num(v int32) *ProductsGovResponseImageTypeNumsBuilder {
	b.s.Num = &v
	return b
}

func (b *ProductsGovResponseImageTypeNumsBuilder) ImageType(v string) *ProductsGovResponseImageTypeNumsBuilder {
	b.s.ImageType = &v
	return b
}

func (b *ProductsGovResponseImageTypeNumsBuilder) Build() *ProductsGovResponseImageTypeNums {
	return b.s
}


