// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetResourceListRequest struct {


	GetResourceListBody *GetResourceListBody `json:"getResourceListBody,omitempty"`
}

func (s GetResourceListRequest) String() string {
	return utils.Beautify(s)
}

func (s GetResourceListRequest) GoString() string {
	return s.String()
}

func (s GetResourceListRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetResourceListRequest) SetGetResourceListBody(v *GetResourceListBody) *GetResourceListRequest {
	s.GetResourceListBody = v
	return s
}


type GetResourceListRequestBuilder struct {
	s *GetResourceListRequest
}

func NewGetResourceListRequestBuilder() *GetResourceListRequestBuilder {
	s := &GetResourceListRequest{}
	b := &GetResourceListRequestBuilder{s: s}
	return b
}

func (b *GetResourceListRequestBuilder) GetResourceListBody(v *GetResourceListBody) *GetResourceListRequestBuilder {
	b.s.GetResourceListBody = v
	return b
}

func (b *GetResourceListRequestBuilder) Build() *GetResourceListRequest {
	return b.s
}


