// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type RegionsResponse struct {

    // 响应数据
	Data *[]RegionsResponseData `json:"data,omitempty"`
    // 请求ID
	RequestId *string `json:"requestId,omitempty"`
    // 响应信息
	RespMsg *string `json:"respMsg,omitempty"`
    // 响应码
	RespCode *string `json:"respCode,omitempty"`
}

func (s RegionsResponse) String() string {
	return utils.Beautify(s)
}

func (s RegionsResponse) GoString() string {
	return s.String()
}

func (s RegionsResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *RegionsResponse) SetData(v []RegionsResponseData) *RegionsResponse {
	s.Data = &v
	return s
}

func (s *RegionsResponse) SetRequestId(v string) *RegionsResponse {
	s.RequestId = &v
	return s
}

func (s *RegionsResponse) SetRespMsg(v string) *RegionsResponse {
	s.RespMsg = &v
	return s
}

func (s *RegionsResponse) SetRespCode(v string) *RegionsResponse {
	s.RespCode = &v
	return s
}


type RegionsResponseBuilder struct {
	s *RegionsResponse
}

func NewRegionsResponseBuilder() *RegionsResponseBuilder {
	s := &RegionsResponse{}
	b := &RegionsResponseBuilder{s: s}
	return b
}

func (b *RegionsResponseBuilder) Data(v []RegionsResponseData) *RegionsResponseBuilder {
	b.s.Data = &v
	return b
}

func (b *RegionsResponseBuilder) RequestId(v string) *RegionsResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *RegionsResponseBuilder) RespMsg(v string) *RegionsResponseBuilder {
	b.s.RespMsg = &v
	return b
}

func (b *RegionsResponseBuilder) RespCode(v string) *RegionsResponseBuilder {
	b.s.RespCode = &v
	return b
}

func (b *RegionsResponseBuilder) Build() *RegionsResponse {
	return b.s
}


