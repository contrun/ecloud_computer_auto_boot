// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetNatListRequest struct {


	GetNatListBody *GetNatListBody `json:"getNatListBody,omitempty"`
}

func (s GetNatListRequest) String() string {
	return utils.Beautify(s)
}

func (s GetNatListRequest) GoString() string {
	return s.String()
}

func (s GetNatListRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetNatListRequest) SetGetNatListBody(v *GetNatListBody) *GetNatListRequest {
	s.GetNatListBody = v
	return s
}


type GetNatListRequestBuilder struct {
	s *GetNatListRequest
}

func NewGetNatListRequestBuilder() *GetNatListRequestBuilder {
	s := &GetNatListRequest{}
	b := &GetNatListRequestBuilder{s: s}
	return b
}

func (b *GetNatListRequestBuilder) GetNatListBody(v *GetNatListBody) *GetNatListRequestBuilder {
	b.s.GetNatListBody = v
	return b
}

func (b *GetNatListRequestBuilder) Build() *GetNatListRequest {
	return b.s
}


