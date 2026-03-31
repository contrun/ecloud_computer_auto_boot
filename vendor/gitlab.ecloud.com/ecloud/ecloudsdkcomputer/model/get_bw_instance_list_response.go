// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetBwInstanceListResponse struct {

    // 响应数据
	Data *GetBwInstanceListResponseData `json:"data,omitempty"`
    // 请求ID
	RequestId *string `json:"requestId,omitempty"`
    // 响应信息
	RespMsg *string `json:"respMsg,omitempty"`
    // 响应编码
	RespCode *string `json:"respCode,omitempty"`
}

func (s GetBwInstanceListResponse) String() string {
	return utils.Beautify(s)
}

func (s GetBwInstanceListResponse) GoString() string {
	return s.String()
}

func (s GetBwInstanceListResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetBwInstanceListResponse) SetData(v *GetBwInstanceListResponseData) *GetBwInstanceListResponse {
	s.Data = v
	return s
}

func (s *GetBwInstanceListResponse) SetRequestId(v string) *GetBwInstanceListResponse {
	s.RequestId = &v
	return s
}

func (s *GetBwInstanceListResponse) SetRespMsg(v string) *GetBwInstanceListResponse {
	s.RespMsg = &v
	return s
}

func (s *GetBwInstanceListResponse) SetRespCode(v string) *GetBwInstanceListResponse {
	s.RespCode = &v
	return s
}


type GetBwInstanceListResponseBuilder struct {
	s *GetBwInstanceListResponse
}

func NewGetBwInstanceListResponseBuilder() *GetBwInstanceListResponseBuilder {
	s := &GetBwInstanceListResponse{}
	b := &GetBwInstanceListResponseBuilder{s: s}
	return b
}

func (b *GetBwInstanceListResponseBuilder) Data(v *GetBwInstanceListResponseData) *GetBwInstanceListResponseBuilder {
	b.s.Data = v
	return b
}

func (b *GetBwInstanceListResponseBuilder) RequestId(v string) *GetBwInstanceListResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *GetBwInstanceListResponseBuilder) RespMsg(v string) *GetBwInstanceListResponseBuilder {
	b.s.RespMsg = &v
	return b
}

func (b *GetBwInstanceListResponseBuilder) RespCode(v string) *GetBwInstanceListResponseBuilder {
	b.s.RespCode = &v
	return b
}

func (b *GetBwInstanceListResponseBuilder) Build() *GetBwInstanceListResponse {
	return b.s
}


