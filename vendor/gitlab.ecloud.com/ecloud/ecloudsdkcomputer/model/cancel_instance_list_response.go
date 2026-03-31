// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CancelInstanceListResponse struct {

    // 响应数据
	Data *CancelInstanceListResponseData `json:"data,omitempty"`
    // 请求ID
	RequestId *string `json:"requestId,omitempty"`
    // 响应信息
	RespMsg *string `json:"respMsg,omitempty"`
    // 响应编码
	RespCode *string `json:"respCode,omitempty"`
}

func (s CancelInstanceListResponse) String() string {
	return utils.Beautify(s)
}

func (s CancelInstanceListResponse) GoString() string {
	return s.String()
}

func (s CancelInstanceListResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CancelInstanceListResponse) SetData(v *CancelInstanceListResponseData) *CancelInstanceListResponse {
	s.Data = v
	return s
}

func (s *CancelInstanceListResponse) SetRequestId(v string) *CancelInstanceListResponse {
	s.RequestId = &v
	return s
}

func (s *CancelInstanceListResponse) SetRespMsg(v string) *CancelInstanceListResponse {
	s.RespMsg = &v
	return s
}

func (s *CancelInstanceListResponse) SetRespCode(v string) *CancelInstanceListResponse {
	s.RespCode = &v
	return s
}


type CancelInstanceListResponseBuilder struct {
	s *CancelInstanceListResponse
}

func NewCancelInstanceListResponseBuilder() *CancelInstanceListResponseBuilder {
	s := &CancelInstanceListResponse{}
	b := &CancelInstanceListResponseBuilder{s: s}
	return b
}

func (b *CancelInstanceListResponseBuilder) Data(v *CancelInstanceListResponseData) *CancelInstanceListResponseBuilder {
	b.s.Data = v
	return b
}

func (b *CancelInstanceListResponseBuilder) RequestId(v string) *CancelInstanceListResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *CancelInstanceListResponseBuilder) RespMsg(v string) *CancelInstanceListResponseBuilder {
	b.s.RespMsg = &v
	return b
}

func (b *CancelInstanceListResponseBuilder) RespCode(v string) *CancelInstanceListResponseBuilder {
	b.s.RespCode = &v
	return b
}

func (b *CancelInstanceListResponseBuilder) Build() *CancelInstanceListResponse {
	return b.s
}


