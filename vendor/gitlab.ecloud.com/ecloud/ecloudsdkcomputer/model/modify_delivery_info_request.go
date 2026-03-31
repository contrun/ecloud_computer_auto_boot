// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ModifyDeliveryInfoRequest struct {


	ModifyDeliveryInfoBody *ModifyDeliveryInfoBody `json:"modifyDeliveryInfoBody,omitempty"`
}

func (s ModifyDeliveryInfoRequest) String() string {
	return utils.Beautify(s)
}

func (s ModifyDeliveryInfoRequest) GoString() string {
	return s.String()
}

func (s ModifyDeliveryInfoRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ModifyDeliveryInfoRequest) SetModifyDeliveryInfoBody(v *ModifyDeliveryInfoBody) *ModifyDeliveryInfoRequest {
	s.ModifyDeliveryInfoBody = v
	return s
}


type ModifyDeliveryInfoRequestBuilder struct {
	s *ModifyDeliveryInfoRequest
}

func NewModifyDeliveryInfoRequestBuilder() *ModifyDeliveryInfoRequestBuilder {
	s := &ModifyDeliveryInfoRequest{}
	b := &ModifyDeliveryInfoRequestBuilder{s: s}
	return b
}

func (b *ModifyDeliveryInfoRequestBuilder) ModifyDeliveryInfoBody(v *ModifyDeliveryInfoBody) *ModifyDeliveryInfoRequestBuilder {
	b.s.ModifyDeliveryInfoBody = v
	return b
}

func (b *ModifyDeliveryInfoRequestBuilder) Build() *ModifyDeliveryInfoRequest {
	return b.s
}


