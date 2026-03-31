// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListServiceUnifyApiResponseData struct {

    // 服务包ID
	ChaId *string `json:"chaId,omitempty"`
    // 服务载体名称
	HardwareName *string `json:"hardwareName,omitempty"`
    // 服务包名称
	ServicePackName *string `json:"servicePackName,omitempty"`
    // 可选服务包信息
	Services *[]ListServiceUnifyApiResponseServices `json:"services,omitempty"`
    // 服务载体颜色
	Colours *[]ListServiceUnifyApiResponseColours `json:"colours,omitempty"`
}

func (s ListServiceUnifyApiResponseData) String() string {
	return utils.Beautify(s)
}

func (s ListServiceUnifyApiResponseData) GoString() string {
	return s.String()
}

func (s ListServiceUnifyApiResponseData) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListServiceUnifyApiResponseData) SetChaId(v string) *ListServiceUnifyApiResponseData {
	s.ChaId = &v
	return s
}

func (s *ListServiceUnifyApiResponseData) SetHardwareName(v string) *ListServiceUnifyApiResponseData {
	s.HardwareName = &v
	return s
}

func (s *ListServiceUnifyApiResponseData) SetServicePackName(v string) *ListServiceUnifyApiResponseData {
	s.ServicePackName = &v
	return s
}

func (s *ListServiceUnifyApiResponseData) SetServices(v []ListServiceUnifyApiResponseServices) *ListServiceUnifyApiResponseData {
	s.Services = &v
	return s
}

func (s *ListServiceUnifyApiResponseData) SetColours(v []ListServiceUnifyApiResponseColours) *ListServiceUnifyApiResponseData {
	s.Colours = &v
	return s
}


type ListServiceUnifyApiResponseDataBuilder struct {
	s *ListServiceUnifyApiResponseData
}

func NewListServiceUnifyApiResponseDataBuilder() *ListServiceUnifyApiResponseDataBuilder {
	s := &ListServiceUnifyApiResponseData{}
	b := &ListServiceUnifyApiResponseDataBuilder{s: s}
	return b
}

func (b *ListServiceUnifyApiResponseDataBuilder) ChaId(v string) *ListServiceUnifyApiResponseDataBuilder {
	b.s.ChaId = &v
	return b
}

func (b *ListServiceUnifyApiResponseDataBuilder) HardwareName(v string) *ListServiceUnifyApiResponseDataBuilder {
	b.s.HardwareName = &v
	return b
}

func (b *ListServiceUnifyApiResponseDataBuilder) ServicePackName(v string) *ListServiceUnifyApiResponseDataBuilder {
	b.s.ServicePackName = &v
	return b
}

func (b *ListServiceUnifyApiResponseDataBuilder) Services(v []ListServiceUnifyApiResponseServices) *ListServiceUnifyApiResponseDataBuilder {
	b.s.Services = &v
	return b
}

func (b *ListServiceUnifyApiResponseDataBuilder) Colours(v []ListServiceUnifyApiResponseColours) *ListServiceUnifyApiResponseDataBuilder {
	b.s.Colours = &v
	return b
}

func (b *ListServiceUnifyApiResponseDataBuilder) Build() *ListServiceUnifyApiResponseData {
	return b.s
}


