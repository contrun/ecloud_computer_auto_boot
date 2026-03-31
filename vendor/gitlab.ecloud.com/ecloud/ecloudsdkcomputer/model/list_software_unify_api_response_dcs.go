// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListSoftwareUnifyApiResponseDcs struct {

    // 网络工作区数量
	Num *int32 `json:"num,omitempty"`
    // 网络工作区编码
	Dc *string `json:"dc,omitempty"`
    // 网络工作区下操作系统信息
	ImageTypeNums *[]ListSoftwareUnifyApiResponseImageTypeNums `json:"imageTypeNums,omitempty"`
}

func (s ListSoftwareUnifyApiResponseDcs) String() string {
	return utils.Beautify(s)
}

func (s ListSoftwareUnifyApiResponseDcs) GoString() string {
	return s.String()
}

func (s ListSoftwareUnifyApiResponseDcs) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListSoftwareUnifyApiResponseDcs) SetNum(v int32) *ListSoftwareUnifyApiResponseDcs {
	s.Num = &v
	return s
}

func (s *ListSoftwareUnifyApiResponseDcs) SetDc(v string) *ListSoftwareUnifyApiResponseDcs {
	s.Dc = &v
	return s
}

func (s *ListSoftwareUnifyApiResponseDcs) SetImageTypeNums(v []ListSoftwareUnifyApiResponseImageTypeNums) *ListSoftwareUnifyApiResponseDcs {
	s.ImageTypeNums = &v
	return s
}


type ListSoftwareUnifyApiResponseDcsBuilder struct {
	s *ListSoftwareUnifyApiResponseDcs
}

func NewListSoftwareUnifyApiResponseDcsBuilder() *ListSoftwareUnifyApiResponseDcsBuilder {
	s := &ListSoftwareUnifyApiResponseDcs{}
	b := &ListSoftwareUnifyApiResponseDcsBuilder{s: s}
	return b
}

func (b *ListSoftwareUnifyApiResponseDcsBuilder) Num(v int32) *ListSoftwareUnifyApiResponseDcsBuilder {
	b.s.Num = &v
	return b
}

func (b *ListSoftwareUnifyApiResponseDcsBuilder) Dc(v string) *ListSoftwareUnifyApiResponseDcsBuilder {
	b.s.Dc = &v
	return b
}

func (b *ListSoftwareUnifyApiResponseDcsBuilder) ImageTypeNums(v []ListSoftwareUnifyApiResponseImageTypeNums) *ListSoftwareUnifyApiResponseDcsBuilder {
	b.s.ImageTypeNums = &v
	return b
}

func (b *ListSoftwareUnifyApiResponseDcsBuilder) Build() *ListSoftwareUnifyApiResponseDcs {
	return b.s
}


