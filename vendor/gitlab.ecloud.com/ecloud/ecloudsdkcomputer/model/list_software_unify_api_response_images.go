// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListSoftwareUnifyApiResponseImages struct {

    // 操作系统名称
	ImageName *string `json:"imageName,omitempty"`
    // 操作系统总可售数量
	Num *int32 `json:"num,omitempty"`
    // 操作系统编码，0：Windows 10，1：统信UOS_1050，2：麒麟KylinOS_V10，3：Windows 11，4：Windows Server 2016
	ImageCode *string `json:"imageCode,omitempty"`
    // 操作系统标识，UOS：统信UOS_1050，Kylin：麒麟KylinOS_V10，Windows：Windows 10，Windows11：Windows 11，WinS2016：Windows Server 2016
	ImageType *string `json:"imageType,omitempty"`
}

func (s ListSoftwareUnifyApiResponseImages) String() string {
	return utils.Beautify(s)
}

func (s ListSoftwareUnifyApiResponseImages) GoString() string {
	return s.String()
}

func (s ListSoftwareUnifyApiResponseImages) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListSoftwareUnifyApiResponseImages) SetImageName(v string) *ListSoftwareUnifyApiResponseImages {
	s.ImageName = &v
	return s
}

func (s *ListSoftwareUnifyApiResponseImages) SetNum(v int32) *ListSoftwareUnifyApiResponseImages {
	s.Num = &v
	return s
}

func (s *ListSoftwareUnifyApiResponseImages) SetImageCode(v string) *ListSoftwareUnifyApiResponseImages {
	s.ImageCode = &v
	return s
}

func (s *ListSoftwareUnifyApiResponseImages) SetImageType(v string) *ListSoftwareUnifyApiResponseImages {
	s.ImageType = &v
	return s
}


type ListSoftwareUnifyApiResponseImagesBuilder struct {
	s *ListSoftwareUnifyApiResponseImages
}

func NewListSoftwareUnifyApiResponseImagesBuilder() *ListSoftwareUnifyApiResponseImagesBuilder {
	s := &ListSoftwareUnifyApiResponseImages{}
	b := &ListSoftwareUnifyApiResponseImagesBuilder{s: s}
	return b
}

func (b *ListSoftwareUnifyApiResponseImagesBuilder) ImageName(v string) *ListSoftwareUnifyApiResponseImagesBuilder {
	b.s.ImageName = &v
	return b
}

func (b *ListSoftwareUnifyApiResponseImagesBuilder) Num(v int32) *ListSoftwareUnifyApiResponseImagesBuilder {
	b.s.Num = &v
	return b
}

func (b *ListSoftwareUnifyApiResponseImagesBuilder) ImageCode(v string) *ListSoftwareUnifyApiResponseImagesBuilder {
	b.s.ImageCode = &v
	return b
}

func (b *ListSoftwareUnifyApiResponseImagesBuilder) ImageType(v string) *ListSoftwareUnifyApiResponseImagesBuilder {
	b.s.ImageType = &v
	return b
}

func (b *ListSoftwareUnifyApiResponseImagesBuilder) Build() *ListSoftwareUnifyApiResponseImages {
	return b.s
}


