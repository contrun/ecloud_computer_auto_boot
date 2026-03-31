// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetVmListRequest struct {


	GetVmListBody *GetVmListBody `json:"getVmListBody,omitempty"`
}

func (s GetVmListRequest) String() string {
	return utils.Beautify(s)
}

func (s GetVmListRequest) GoString() string {
	return s.String()
}

func (s GetVmListRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetVmListRequest) SetGetVmListBody(v *GetVmListBody) *GetVmListRequest {
	s.GetVmListBody = v
	return s
}


type GetVmListRequestBuilder struct {
	s *GetVmListRequest
}

func NewGetVmListRequestBuilder() *GetVmListRequestBuilder {
	s := &GetVmListRequest{}
	b := &GetVmListRequestBuilder{s: s}
	return b
}

func (b *GetVmListRequestBuilder) GetVmListBody(v *GetVmListBody) *GetVmListRequestBuilder {
	b.s.GetVmListBody = v
	return b
}

func (b *GetVmListRequestBuilder) Build() *GetVmListRequest {
	return b.s
}


