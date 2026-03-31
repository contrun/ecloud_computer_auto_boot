// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetNatListResponse struct {

    // 响应数据
	Data *GetNatListResponseData `json:"data,omitempty"`
    // 请求ID
	RequestId *string `json:"requestId,omitempty"`
    // 响应信息
	RespMsg *string `json:"respMsg,omitempty"`
    // 响应编码
	RespCode *string `json:"respCode,omitempty"`
}

func (s GetNatListResponse) String() string {
	return utils.Beautify(s)
}

func (s GetNatListResponse) GoString() string {
	return s.String()
}

func (s GetNatListResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetNatListResponse) SetData(v *GetNatListResponseData) *GetNatListResponse {
	s.Data = v
	return s
}

func (s *GetNatListResponse) SetRequestId(v string) *GetNatListResponse {
	s.RequestId = &v
	return s
}

func (s *GetNatListResponse) SetRespMsg(v string) *GetNatListResponse {
	s.RespMsg = &v
	return s
}

func (s *GetNatListResponse) SetRespCode(v string) *GetNatListResponse {
	s.RespCode = &v
	return s
}


type GetNatListResponseBuilder struct {
	s *GetNatListResponse
}

func NewGetNatListResponseBuilder() *GetNatListResponseBuilder {
	s := &GetNatListResponse{}
	b := &GetNatListResponseBuilder{s: s}
	return b
}

func (b *GetNatListResponseBuilder) Data(v *GetNatListResponseData) *GetNatListResponseBuilder {
	b.s.Data = v
	return b
}

func (b *GetNatListResponseBuilder) RequestId(v string) *GetNatListResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *GetNatListResponseBuilder) RespMsg(v string) *GetNatListResponseBuilder {
	b.s.RespMsg = &v
	return b
}

func (b *GetNatListResponseBuilder) RespCode(v string) *GetNatListResponseBuilder {
	b.s.RespCode = &v
	return b
}

func (b *GetNatListResponseBuilder) Build() *GetNatListResponse {
	return b.s
}


