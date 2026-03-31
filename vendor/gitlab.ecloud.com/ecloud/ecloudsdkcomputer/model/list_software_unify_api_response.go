// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListSoftwareUnifyApiResponse struct {

    // 响应数据
	Data *[]ListSoftwareUnifyApiResponseData `json:"data,omitempty"`
    // 请求ID
	RequestId *string `json:"requestId,omitempty"`
    // 响应信息
	RespMsg *string `json:"respMsg,omitempty"`
    // 响应码
	RespCode *string `json:"respCode,omitempty"`
}

func (s ListSoftwareUnifyApiResponse) String() string {
	return utils.Beautify(s)
}

func (s ListSoftwareUnifyApiResponse) GoString() string {
	return s.String()
}

func (s ListSoftwareUnifyApiResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListSoftwareUnifyApiResponse) SetData(v []ListSoftwareUnifyApiResponseData) *ListSoftwareUnifyApiResponse {
	s.Data = &v
	return s
}

func (s *ListSoftwareUnifyApiResponse) SetRequestId(v string) *ListSoftwareUnifyApiResponse {
	s.RequestId = &v
	return s
}

func (s *ListSoftwareUnifyApiResponse) SetRespMsg(v string) *ListSoftwareUnifyApiResponse {
	s.RespMsg = &v
	return s
}

func (s *ListSoftwareUnifyApiResponse) SetRespCode(v string) *ListSoftwareUnifyApiResponse {
	s.RespCode = &v
	return s
}


type ListSoftwareUnifyApiResponseBuilder struct {
	s *ListSoftwareUnifyApiResponse
}

func NewListSoftwareUnifyApiResponseBuilder() *ListSoftwareUnifyApiResponseBuilder {
	s := &ListSoftwareUnifyApiResponse{}
	b := &ListSoftwareUnifyApiResponseBuilder{s: s}
	return b
}

func (b *ListSoftwareUnifyApiResponseBuilder) Data(v []ListSoftwareUnifyApiResponseData) *ListSoftwareUnifyApiResponseBuilder {
	b.s.Data = &v
	return b
}

func (b *ListSoftwareUnifyApiResponseBuilder) RequestId(v string) *ListSoftwareUnifyApiResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *ListSoftwareUnifyApiResponseBuilder) RespMsg(v string) *ListSoftwareUnifyApiResponseBuilder {
	b.s.RespMsg = &v
	return b
}

func (b *ListSoftwareUnifyApiResponseBuilder) RespCode(v string) *ListSoftwareUnifyApiResponseBuilder {
	b.s.RespCode = &v
	return b
}

func (b *ListSoftwareUnifyApiResponseBuilder) Build() *ListSoftwareUnifyApiResponse {
	return b.s
}


