// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListServiceUnifyApiResponseColours struct {

    // 颜色编码
	Code *string `json:"code,omitempty"`
    // 颜色名称
	Name *string `json:"name,omitempty"`
}

func (s ListServiceUnifyApiResponseColours) String() string {
	return utils.Beautify(s)
}

func (s ListServiceUnifyApiResponseColours) GoString() string {
	return s.String()
}

func (s ListServiceUnifyApiResponseColours) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListServiceUnifyApiResponseColours) SetCode(v string) *ListServiceUnifyApiResponseColours {
	s.Code = &v
	return s
}

func (s *ListServiceUnifyApiResponseColours) SetName(v string) *ListServiceUnifyApiResponseColours {
	s.Name = &v
	return s
}


type ListServiceUnifyApiResponseColoursBuilder struct {
	s *ListServiceUnifyApiResponseColours
}

func NewListServiceUnifyApiResponseColoursBuilder() *ListServiceUnifyApiResponseColoursBuilder {
	s := &ListServiceUnifyApiResponseColours{}
	b := &ListServiceUnifyApiResponseColoursBuilder{s: s}
	return b
}

func (b *ListServiceUnifyApiResponseColoursBuilder) Code(v string) *ListServiceUnifyApiResponseColoursBuilder {
	b.s.Code = &v
	return b
}

func (b *ListServiceUnifyApiResponseColoursBuilder) Name(v string) *ListServiceUnifyApiResponseColoursBuilder {
	b.s.Name = &v
	return b
}

func (b *ListServiceUnifyApiResponseColoursBuilder) Build() *ListServiceUnifyApiResponseColours {
	return b.s
}


