// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetCcempDeviceListRequest struct {


	GetCcempDeviceListBody *GetCcempDeviceListBody `json:"getCcempDeviceListBody,omitempty"`
}

func (s GetCcempDeviceListRequest) String() string {
	return utils.Beautify(s)
}

func (s GetCcempDeviceListRequest) GoString() string {
	return s.String()
}

func (s GetCcempDeviceListRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetCcempDeviceListRequest) SetGetCcempDeviceListBody(v *GetCcempDeviceListBody) *GetCcempDeviceListRequest {
	s.GetCcempDeviceListBody = v
	return s
}


type GetCcempDeviceListRequestBuilder struct {
	s *GetCcempDeviceListRequest
}

func NewGetCcempDeviceListRequestBuilder() *GetCcempDeviceListRequestBuilder {
	s := &GetCcempDeviceListRequest{}
	b := &GetCcempDeviceListRequestBuilder{s: s}
	return b
}

func (b *GetCcempDeviceListRequestBuilder) GetCcempDeviceListBody(v *GetCcempDeviceListBody) *GetCcempDeviceListRequestBuilder {
	b.s.GetCcempDeviceListBody = v
	return b
}

func (b *GetCcempDeviceListRequestBuilder) Build() *GetCcempDeviceListRequest {
	return b.s
}


