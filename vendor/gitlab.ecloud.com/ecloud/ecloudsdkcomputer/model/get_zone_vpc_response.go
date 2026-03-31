// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetZoneVpcResponse struct {

    // 响应数据
	Data *[]GetZoneVpcResponseData `json:"data,omitempty"`
    // 请求ID
	RequestId *string `json:"requestId,omitempty"`
    // 响应信息
	RespMsg *string `json:"respMsg,omitempty"`
    // 响应编码
	RespCode *string `json:"respCode,omitempty"`
}

func (s GetZoneVpcResponse) String() string {
	return utils.Beautify(s)
}

func (s GetZoneVpcResponse) GoString() string {
	return s.String()
}

func (s GetZoneVpcResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetZoneVpcResponse) SetData(v []GetZoneVpcResponseData) *GetZoneVpcResponse {
	s.Data = &v
	return s
}

func (s *GetZoneVpcResponse) SetRequestId(v string) *GetZoneVpcResponse {
	s.RequestId = &v
	return s
}

func (s *GetZoneVpcResponse) SetRespMsg(v string) *GetZoneVpcResponse {
	s.RespMsg = &v
	return s
}

func (s *GetZoneVpcResponse) SetRespCode(v string) *GetZoneVpcResponse {
	s.RespCode = &v
	return s
}


type GetZoneVpcResponseBuilder struct {
	s *GetZoneVpcResponse
}

func NewGetZoneVpcResponseBuilder() *GetZoneVpcResponseBuilder {
	s := &GetZoneVpcResponse{}
	b := &GetZoneVpcResponseBuilder{s: s}
	return b
}

func (b *GetZoneVpcResponseBuilder) Data(v []GetZoneVpcResponseData) *GetZoneVpcResponseBuilder {
	b.s.Data = &v
	return b
}

func (b *GetZoneVpcResponseBuilder) RequestId(v string) *GetZoneVpcResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *GetZoneVpcResponseBuilder) RespMsg(v string) *GetZoneVpcResponseBuilder {
	b.s.RespMsg = &v
	return b
}

func (b *GetZoneVpcResponseBuilder) RespCode(v string) *GetZoneVpcResponseBuilder {
	b.s.RespCode = &v
	return b
}

func (b *GetZoneVpcResponseBuilder) Build() *GetZoneVpcResponse {
	return b.s
}


