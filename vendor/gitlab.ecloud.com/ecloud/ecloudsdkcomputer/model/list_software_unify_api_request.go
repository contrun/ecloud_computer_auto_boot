// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListSoftwareUnifyApiRequest struct {


	ListSoftwareUnifyApiQuery *ListSoftwareUnifyApiQuery `json:"listSoftwareUnifyApiQuery,omitempty"`
}

func (s ListSoftwareUnifyApiRequest) String() string {
	return utils.Beautify(s)
}

func (s ListSoftwareUnifyApiRequest) GoString() string {
	return s.String()
}

func (s ListSoftwareUnifyApiRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListSoftwareUnifyApiRequest) SetListSoftwareUnifyApiQuery(v *ListSoftwareUnifyApiQuery) *ListSoftwareUnifyApiRequest {
	s.ListSoftwareUnifyApiQuery = v
	return s
}


type ListSoftwareUnifyApiRequestBuilder struct {
	s *ListSoftwareUnifyApiRequest
}

func NewListSoftwareUnifyApiRequestBuilder() *ListSoftwareUnifyApiRequestBuilder {
	s := &ListSoftwareUnifyApiRequest{}
	b := &ListSoftwareUnifyApiRequestBuilder{s: s}
	return b
}

func (b *ListSoftwareUnifyApiRequestBuilder) ListSoftwareUnifyApiQuery(v *ListSoftwareUnifyApiQuery) *ListSoftwareUnifyApiRequestBuilder {
	b.s.ListSoftwareUnifyApiQuery = v
	return b
}

func (b *ListSoftwareUnifyApiRequestBuilder) Build() *ListSoftwareUnifyApiRequest {
	return b.s
}


