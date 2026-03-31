// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type MultiDiskChangeResponse struct {

    // 响应数据
	Data *MultiDiskChangeResponseData `json:"data,omitempty"`
    // 请求ID
	RequestId *string `json:"requestId,omitempty"`
    // 响应信息
	RespMsg *string `json:"respMsg,omitempty"`
    // 响应码
	RespCode *string `json:"respCode,omitempty"`
}

func (s MultiDiskChangeResponse) String() string {
	return utils.Beautify(s)
}

func (s MultiDiskChangeResponse) GoString() string {
	return s.String()
}

func (s MultiDiskChangeResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *MultiDiskChangeResponse) SetData(v *MultiDiskChangeResponseData) *MultiDiskChangeResponse {
	s.Data = v
	return s
}

func (s *MultiDiskChangeResponse) SetRequestId(v string) *MultiDiskChangeResponse {
	s.RequestId = &v
	return s
}

func (s *MultiDiskChangeResponse) SetRespMsg(v string) *MultiDiskChangeResponse {
	s.RespMsg = &v
	return s
}

func (s *MultiDiskChangeResponse) SetRespCode(v string) *MultiDiskChangeResponse {
	s.RespCode = &v
	return s
}


type MultiDiskChangeResponseBuilder struct {
	s *MultiDiskChangeResponse
}

func NewMultiDiskChangeResponseBuilder() *MultiDiskChangeResponseBuilder {
	s := &MultiDiskChangeResponse{}
	b := &MultiDiskChangeResponseBuilder{s: s}
	return b
}

func (b *MultiDiskChangeResponseBuilder) Data(v *MultiDiskChangeResponseData) *MultiDiskChangeResponseBuilder {
	b.s.Data = v
	return b
}

func (b *MultiDiskChangeResponseBuilder) RequestId(v string) *MultiDiskChangeResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *MultiDiskChangeResponseBuilder) RespMsg(v string) *MultiDiskChangeResponseBuilder {
	b.s.RespMsg = &v
	return b
}

func (b *MultiDiskChangeResponseBuilder) RespCode(v string) *MultiDiskChangeResponseBuilder {
	b.s.RespCode = &v
	return b
}

func (b *MultiDiskChangeResponseBuilder) Build() *MultiDiskChangeResponse {
	return b.s
}


