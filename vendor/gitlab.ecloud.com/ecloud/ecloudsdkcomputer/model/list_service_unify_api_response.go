// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListServiceUnifyApiResponse struct {

    // 响应数据
	Data *[]ListServiceUnifyApiResponseData `json:"data,omitempty"`
    // 请求ID
	RequestId *string `json:"requestId,omitempty"`
    // 响应信息
	RespMsg *string `json:"respMsg,omitempty"`
    // 响应码
	RespCode *string `json:"respCode,omitempty"`
}

func (s ListServiceUnifyApiResponse) String() string {
	return utils.Beautify(s)
}

func (s ListServiceUnifyApiResponse) GoString() string {
	return s.String()
}

func (s ListServiceUnifyApiResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListServiceUnifyApiResponse) SetData(v []ListServiceUnifyApiResponseData) *ListServiceUnifyApiResponse {
	s.Data = &v
	return s
}

func (s *ListServiceUnifyApiResponse) SetRequestId(v string) *ListServiceUnifyApiResponse {
	s.RequestId = &v
	return s
}

func (s *ListServiceUnifyApiResponse) SetRespMsg(v string) *ListServiceUnifyApiResponse {
	s.RespMsg = &v
	return s
}

func (s *ListServiceUnifyApiResponse) SetRespCode(v string) *ListServiceUnifyApiResponse {
	s.RespCode = &v
	return s
}


type ListServiceUnifyApiResponseBuilder struct {
	s *ListServiceUnifyApiResponse
}

func NewListServiceUnifyApiResponseBuilder() *ListServiceUnifyApiResponseBuilder {
	s := &ListServiceUnifyApiResponse{}
	b := &ListServiceUnifyApiResponseBuilder{s: s}
	return b
}

func (b *ListServiceUnifyApiResponseBuilder) Data(v []ListServiceUnifyApiResponseData) *ListServiceUnifyApiResponseBuilder {
	b.s.Data = &v
	return b
}

func (b *ListServiceUnifyApiResponseBuilder) RequestId(v string) *ListServiceUnifyApiResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *ListServiceUnifyApiResponseBuilder) RespMsg(v string) *ListServiceUnifyApiResponseBuilder {
	b.s.RespMsg = &v
	return b
}

func (b *ListServiceUnifyApiResponseBuilder) RespCode(v string) *ListServiceUnifyApiResponseBuilder {
	b.s.RespCode = &v
	return b
}

func (b *ListServiceUnifyApiResponseBuilder) Build() *ListServiceUnifyApiResponse {
	return b.s
}


