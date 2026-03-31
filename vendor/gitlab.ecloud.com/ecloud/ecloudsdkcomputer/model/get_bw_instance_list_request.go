// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetBwInstanceListRequest struct {


	GetBwInstanceListBody *GetBwInstanceListBody `json:"getBwInstanceListBody,omitempty"`
}

func (s GetBwInstanceListRequest) String() string {
	return utils.Beautify(s)
}

func (s GetBwInstanceListRequest) GoString() string {
	return s.String()
}

func (s GetBwInstanceListRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetBwInstanceListRequest) SetGetBwInstanceListBody(v *GetBwInstanceListBody) *GetBwInstanceListRequest {
	s.GetBwInstanceListBody = v
	return s
}


type GetBwInstanceListRequestBuilder struct {
	s *GetBwInstanceListRequest
}

func NewGetBwInstanceListRequestBuilder() *GetBwInstanceListRequestBuilder {
	s := &GetBwInstanceListRequest{}
	b := &GetBwInstanceListRequestBuilder{s: s}
	return b
}

func (b *GetBwInstanceListRequestBuilder) GetBwInstanceListBody(v *GetBwInstanceListBody) *GetBwInstanceListRequestBuilder {
	b.s.GetBwInstanceListBody = v
	return b
}

func (b *GetBwInstanceListRequestBuilder) Build() *GetBwInstanceListRequest {
	return b.s
}


