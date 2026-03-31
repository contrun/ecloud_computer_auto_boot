// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListServiceUnifyApiResponseServices struct {

    // 可选服务包ID
	ChaId *string `json:"chaId,omitempty"`
    // 可选服务包名称
	Name *string `json:"name,omitempty"`
}

func (s ListServiceUnifyApiResponseServices) String() string {
	return utils.Beautify(s)
}

func (s ListServiceUnifyApiResponseServices) GoString() string {
	return s.String()
}

func (s ListServiceUnifyApiResponseServices) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListServiceUnifyApiResponseServices) SetChaId(v string) *ListServiceUnifyApiResponseServices {
	s.ChaId = &v
	return s
}

func (s *ListServiceUnifyApiResponseServices) SetName(v string) *ListServiceUnifyApiResponseServices {
	s.Name = &v
	return s
}


type ListServiceUnifyApiResponseServicesBuilder struct {
	s *ListServiceUnifyApiResponseServices
}

func NewListServiceUnifyApiResponseServicesBuilder() *ListServiceUnifyApiResponseServicesBuilder {
	s := &ListServiceUnifyApiResponseServices{}
	b := &ListServiceUnifyApiResponseServicesBuilder{s: s}
	return b
}

func (b *ListServiceUnifyApiResponseServicesBuilder) ChaId(v string) *ListServiceUnifyApiResponseServicesBuilder {
	b.s.ChaId = &v
	return b
}

func (b *ListServiceUnifyApiResponseServicesBuilder) Name(v string) *ListServiceUnifyApiResponseServicesBuilder {
	b.s.Name = &v
	return b
}

func (b *ListServiceUnifyApiResponseServicesBuilder) Build() *ListServiceUnifyApiResponseServices {
	return b.s
}


