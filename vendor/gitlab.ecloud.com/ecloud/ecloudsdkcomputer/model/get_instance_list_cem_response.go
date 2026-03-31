// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetInstanceListCemResponse struct {

    // 响应数据
	Data *GetInstanceListCemResponseData `json:"data,omitempty"`
    // 请求ID
	RequestId *string `json:"requestId,omitempty"`
    // 响应信息
	RespMsg *string `json:"respMsg,omitempty"`
    // 响应码
	RespCode *string `json:"respCode,omitempty"`
}

func (s GetInstanceListCemResponse) String() string {
	return utils.Beautify(s)
}

func (s GetInstanceListCemResponse) GoString() string {
	return s.String()
}

func (s GetInstanceListCemResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetInstanceListCemResponse) SetData(v *GetInstanceListCemResponseData) *GetInstanceListCemResponse {
	s.Data = v
	return s
}

func (s *GetInstanceListCemResponse) SetRequestId(v string) *GetInstanceListCemResponse {
	s.RequestId = &v
	return s
}

func (s *GetInstanceListCemResponse) SetRespMsg(v string) *GetInstanceListCemResponse {
	s.RespMsg = &v
	return s
}

func (s *GetInstanceListCemResponse) SetRespCode(v string) *GetInstanceListCemResponse {
	s.RespCode = &v
	return s
}


type GetInstanceListCemResponseBuilder struct {
	s *GetInstanceListCemResponse
}

func NewGetInstanceListCemResponseBuilder() *GetInstanceListCemResponseBuilder {
	s := &GetInstanceListCemResponse{}
	b := &GetInstanceListCemResponseBuilder{s: s}
	return b
}

func (b *GetInstanceListCemResponseBuilder) Data(v *GetInstanceListCemResponseData) *GetInstanceListCemResponseBuilder {
	b.s.Data = v
	return b
}

func (b *GetInstanceListCemResponseBuilder) RequestId(v string) *GetInstanceListCemResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *GetInstanceListCemResponseBuilder) RespMsg(v string) *GetInstanceListCemResponseBuilder {
	b.s.RespMsg = &v
	return b
}

func (b *GetInstanceListCemResponseBuilder) RespCode(v string) *GetInstanceListCemResponseBuilder {
	b.s.RespCode = &v
	return b
}

func (b *GetInstanceListCemResponseBuilder) Build() *GetInstanceListCemResponse {
	return b.s
}


