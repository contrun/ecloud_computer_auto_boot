// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetMultiDiskInfoResponse struct {

    // 响应数据
	Data *GetMultiDiskInfoResponseData `json:"data,omitempty"`
    // 请求ID
	RequestId *string `json:"requestId,omitempty"`
    // 响应信息
	RespMsg *string `json:"respMsg,omitempty"`
    // 响应码
	RespCode *string `json:"respCode,omitempty"`
}

func (s GetMultiDiskInfoResponse) String() string {
	return utils.Beautify(s)
}

func (s GetMultiDiskInfoResponse) GoString() string {
	return s.String()
}

func (s GetMultiDiskInfoResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetMultiDiskInfoResponse) SetData(v *GetMultiDiskInfoResponseData) *GetMultiDiskInfoResponse {
	s.Data = v
	return s
}

func (s *GetMultiDiskInfoResponse) SetRequestId(v string) *GetMultiDiskInfoResponse {
	s.RequestId = &v
	return s
}

func (s *GetMultiDiskInfoResponse) SetRespMsg(v string) *GetMultiDiskInfoResponse {
	s.RespMsg = &v
	return s
}

func (s *GetMultiDiskInfoResponse) SetRespCode(v string) *GetMultiDiskInfoResponse {
	s.RespCode = &v
	return s
}


type GetMultiDiskInfoResponseBuilder struct {
	s *GetMultiDiskInfoResponse
}

func NewGetMultiDiskInfoResponseBuilder() *GetMultiDiskInfoResponseBuilder {
	s := &GetMultiDiskInfoResponse{}
	b := &GetMultiDiskInfoResponseBuilder{s: s}
	return b
}

func (b *GetMultiDiskInfoResponseBuilder) Data(v *GetMultiDiskInfoResponseData) *GetMultiDiskInfoResponseBuilder {
	b.s.Data = v
	return b
}

func (b *GetMultiDiskInfoResponseBuilder) RequestId(v string) *GetMultiDiskInfoResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *GetMultiDiskInfoResponseBuilder) RespMsg(v string) *GetMultiDiskInfoResponseBuilder {
	b.s.RespMsg = &v
	return b
}

func (b *GetMultiDiskInfoResponseBuilder) RespCode(v string) *GetMultiDiskInfoResponseBuilder {
	b.s.RespCode = &v
	return b
}

func (b *GetMultiDiskInfoResponseBuilder) Build() *GetMultiDiskInfoResponse {
	return b.s
}


