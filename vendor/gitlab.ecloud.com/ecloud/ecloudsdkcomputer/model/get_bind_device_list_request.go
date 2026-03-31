// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetBindDeviceListRequest struct {


	GetBindDeviceListBody *GetBindDeviceListBody `json:"getBindDeviceListBody,omitempty"`
}

func (s GetBindDeviceListRequest) String() string {
	return utils.Beautify(s)
}

func (s GetBindDeviceListRequest) GoString() string {
	return s.String()
}

func (s GetBindDeviceListRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetBindDeviceListRequest) SetGetBindDeviceListBody(v *GetBindDeviceListBody) *GetBindDeviceListRequest {
	s.GetBindDeviceListBody = v
	return s
}


type GetBindDeviceListRequestBuilder struct {
	s *GetBindDeviceListRequest
}

func NewGetBindDeviceListRequestBuilder() *GetBindDeviceListRequestBuilder {
	s := &GetBindDeviceListRequest{}
	b := &GetBindDeviceListRequestBuilder{s: s}
	return b
}

func (b *GetBindDeviceListRequestBuilder) GetBindDeviceListBody(v *GetBindDeviceListBody) *GetBindDeviceListRequestBuilder {
	b.s.GetBindDeviceListBody = v
	return b
}

func (b *GetBindDeviceListRequestBuilder) Build() *GetBindDeviceListRequest {
	return b.s
}


