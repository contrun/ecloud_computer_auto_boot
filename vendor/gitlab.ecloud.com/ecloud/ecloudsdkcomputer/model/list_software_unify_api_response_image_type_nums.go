// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListSoftwareUnifyApiResponseImageTypeNums struct {

    // 网络工作区下的操作系统可售数量
	Num *int32 `json:"num,omitempty"`
    // 操作系统标识，UOS：统信UOS_1050，Kylin：麒麟KylinOS_V10，Windows：Windows 10，Windows11：Windows 11，WinS2016：Windows Server 2016
	ImageType *string `json:"imageType,omitempty"`
}

func (s ListSoftwareUnifyApiResponseImageTypeNums) String() string {
	return utils.Beautify(s)
}

func (s ListSoftwareUnifyApiResponseImageTypeNums) GoString() string {
	return s.String()
}

func (s ListSoftwareUnifyApiResponseImageTypeNums) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListSoftwareUnifyApiResponseImageTypeNums) SetNum(v int32) *ListSoftwareUnifyApiResponseImageTypeNums {
	s.Num = &v
	return s
}

func (s *ListSoftwareUnifyApiResponseImageTypeNums) SetImageType(v string) *ListSoftwareUnifyApiResponseImageTypeNums {
	s.ImageType = &v
	return s
}


type ListSoftwareUnifyApiResponseImageTypeNumsBuilder struct {
	s *ListSoftwareUnifyApiResponseImageTypeNums
}

func NewListSoftwareUnifyApiResponseImageTypeNumsBuilder() *ListSoftwareUnifyApiResponseImageTypeNumsBuilder {
	s := &ListSoftwareUnifyApiResponseImageTypeNums{}
	b := &ListSoftwareUnifyApiResponseImageTypeNumsBuilder{s: s}
	return b
}

func (b *ListSoftwareUnifyApiResponseImageTypeNumsBuilder) Num(v int32) *ListSoftwareUnifyApiResponseImageTypeNumsBuilder {
	b.s.Num = &v
	return b
}

func (b *ListSoftwareUnifyApiResponseImageTypeNumsBuilder) ImageType(v string) *ListSoftwareUnifyApiResponseImageTypeNumsBuilder {
	b.s.ImageType = &v
	return b
}

func (b *ListSoftwareUnifyApiResponseImageTypeNumsBuilder) Build() *ListSoftwareUnifyApiResponseImageTypeNums {
	return b.s
}


