// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ProductsGovResponseImages struct {

    // 操作系统名称
	ImageName *string `json:"imageName,omitempty"`
    // 网络工作区数量
	Num *int32 `json:"num,omitempty"`
    // 操作系统编码，0：Windows 10，1：统信UOS_1050，2：麒麟KylinOS_V10，3：Windows 11，4：Windows Server 2016
	ImageCode *string `json:"imageCode,omitempty"`
    // 操作系统标识，UOS：统信UOS_1050，Kylin：麒麟KylinOS_V10，Windows：Windows 10，Windows11：Windows 11，WinS2016：Windows Server 2016
	ImageType *string `json:"imageType,omitempty"`
}

func (s ProductsGovResponseImages) String() string {
	return utils.Beautify(s)
}

func (s ProductsGovResponseImages) GoString() string {
	return s.String()
}

func (s ProductsGovResponseImages) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ProductsGovResponseImages) SetImageName(v string) *ProductsGovResponseImages {
	s.ImageName = &v
	return s
}

func (s *ProductsGovResponseImages) SetNum(v int32) *ProductsGovResponseImages {
	s.Num = &v
	return s
}

func (s *ProductsGovResponseImages) SetImageCode(v string) *ProductsGovResponseImages {
	s.ImageCode = &v
	return s
}

func (s *ProductsGovResponseImages) SetImageType(v string) *ProductsGovResponseImages {
	s.ImageType = &v
	return s
}


type ProductsGovResponseImagesBuilder struct {
	s *ProductsGovResponseImages
}

func NewProductsGovResponseImagesBuilder() *ProductsGovResponseImagesBuilder {
	s := &ProductsGovResponseImages{}
	b := &ProductsGovResponseImagesBuilder{s: s}
	return b
}

func (b *ProductsGovResponseImagesBuilder) ImageName(v string) *ProductsGovResponseImagesBuilder {
	b.s.ImageName = &v
	return b
}

func (b *ProductsGovResponseImagesBuilder) Num(v int32) *ProductsGovResponseImagesBuilder {
	b.s.Num = &v
	return b
}

func (b *ProductsGovResponseImagesBuilder) ImageCode(v string) *ProductsGovResponseImagesBuilder {
	b.s.ImageCode = &v
	return b
}

func (b *ProductsGovResponseImagesBuilder) ImageType(v string) *ProductsGovResponseImagesBuilder {
	b.s.ImageType = &v
	return b
}

func (b *ProductsGovResponseImagesBuilder) Build() *ProductsGovResponseImages {
	return b.s
}


