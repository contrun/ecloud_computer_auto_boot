// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetDeviceDetailRequest struct {


	GetDeviceDetailQuery *GetDeviceDetailQuery `json:"getDeviceDetailQuery,omitempty"`
}

func (s GetDeviceDetailRequest) String() string {
	return utils.Beautify(s)
}

func (s GetDeviceDetailRequest) GoString() string {
	return s.String()
}

func (s GetDeviceDetailRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetDeviceDetailRequest) SetGetDeviceDetailQuery(v *GetDeviceDetailQuery) *GetDeviceDetailRequest {
	s.GetDeviceDetailQuery = v
	return s
}


type GetDeviceDetailRequestBuilder struct {
	s *GetDeviceDetailRequest
}

func NewGetDeviceDetailRequestBuilder() *GetDeviceDetailRequestBuilder {
	s := &GetDeviceDetailRequest{}
	b := &GetDeviceDetailRequestBuilder{s: s}
	return b
}

func (b *GetDeviceDetailRequestBuilder) GetDeviceDetailQuery(v *GetDeviceDetailQuery) *GetDeviceDetailRequestBuilder {
	b.s.GetDeviceDetailQuery = v
	return b
}

func (b *GetDeviceDetailRequestBuilder) Build() *GetDeviceDetailRequest {
	return b.s
}


