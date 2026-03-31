// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListSoftwareUnifyApiQuery struct {
    position.Query
    // 服务包ID
	ChaId *string `json:"chaId,omitempty"`
}

func (s ListSoftwareUnifyApiQuery) String() string {
	return utils.Beautify(s)
}

func (s ListSoftwareUnifyApiQuery) GoString() string {
	return s.String()
}

func (s ListSoftwareUnifyApiQuery) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListSoftwareUnifyApiQuery) SetChaId(v string) *ListSoftwareUnifyApiQuery {
	s.ChaId = &v
	return s
}


type ListSoftwareUnifyApiQueryBuilder struct {
	s *ListSoftwareUnifyApiQuery
}

func NewListSoftwareUnifyApiQueryBuilder() *ListSoftwareUnifyApiQueryBuilder {
	s := &ListSoftwareUnifyApiQuery{}
	b := &ListSoftwareUnifyApiQueryBuilder{s: s}
	return b
}

func (b *ListSoftwareUnifyApiQueryBuilder) ChaId(v string) *ListSoftwareUnifyApiQueryBuilder {
	b.s.ChaId = &v
	return b
}

func (b *ListSoftwareUnifyApiQueryBuilder) Build() *ListSoftwareUnifyApiQuery {
	return b.s
}


