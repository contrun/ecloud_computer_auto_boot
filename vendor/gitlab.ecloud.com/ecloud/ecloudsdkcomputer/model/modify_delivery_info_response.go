// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ModifyDeliveryInfoResponse struct {

    // 响应数据
	Data *string `json:"data,omitempty"`
    // 请求ID
	RequestId *string `json:"requestId,omitempty"`
    // 响应信息
	RespMsg *string `json:"respMsg,omitempty"`
    // 响应码
	RespCode *string `json:"respCode,omitempty"`
}

func (s ModifyDeliveryInfoResponse) String() string {
	return utils.Beautify(s)
}

func (s ModifyDeliveryInfoResponse) GoString() string {
	return s.String()
}

func (s ModifyDeliveryInfoResponse) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ModifyDeliveryInfoResponse) SetData(v string) *ModifyDeliveryInfoResponse {
	s.Data = &v
	return s
}

func (s *ModifyDeliveryInfoResponse) SetRequestId(v string) *ModifyDeliveryInfoResponse {
	s.RequestId = &v
	return s
}

func (s *ModifyDeliveryInfoResponse) SetRespMsg(v string) *ModifyDeliveryInfoResponse {
	s.RespMsg = &v
	return s
}

func (s *ModifyDeliveryInfoResponse) SetRespCode(v string) *ModifyDeliveryInfoResponse {
	s.RespCode = &v
	return s
}


type ModifyDeliveryInfoResponseBuilder struct {
	s *ModifyDeliveryInfoResponse
}

func NewModifyDeliveryInfoResponseBuilder() *ModifyDeliveryInfoResponseBuilder {
	s := &ModifyDeliveryInfoResponse{}
	b := &ModifyDeliveryInfoResponseBuilder{s: s}
	return b
}

func (b *ModifyDeliveryInfoResponseBuilder) Data(v string) *ModifyDeliveryInfoResponseBuilder {
	b.s.Data = &v
	return b
}

func (b *ModifyDeliveryInfoResponseBuilder) RequestId(v string) *ModifyDeliveryInfoResponseBuilder {
	b.s.RequestId = &v
	return b
}

func (b *ModifyDeliveryInfoResponseBuilder) RespMsg(v string) *ModifyDeliveryInfoResponseBuilder {
	b.s.RespMsg = &v
	return b
}

func (b *ModifyDeliveryInfoResponseBuilder) RespCode(v string) *ModifyDeliveryInfoResponseBuilder {
	b.s.RespCode = &v
	return b
}

func (b *ModifyDeliveryInfoResponseBuilder) Build() *ModifyDeliveryInfoResponse {
	return b.s
}


