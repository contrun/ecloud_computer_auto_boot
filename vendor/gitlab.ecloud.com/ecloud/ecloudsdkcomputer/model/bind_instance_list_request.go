// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BindInstanceListRequest struct {


	BindInstanceListBody *BindInstanceListBody `json:"bindInstanceListBody,omitempty"`
}

func (s BindInstanceListRequest) String() string {
	return utils.Beautify(s)
}

func (s BindInstanceListRequest) GoString() string {
	return s.String()
}

func (s BindInstanceListRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BindInstanceListRequest) SetBindInstanceListBody(v *BindInstanceListBody) *BindInstanceListRequest {
	s.BindInstanceListBody = v
	return s
}


type BindInstanceListRequestBuilder struct {
	s *BindInstanceListRequest
}

func NewBindInstanceListRequestBuilder() *BindInstanceListRequestBuilder {
	s := &BindInstanceListRequest{}
	b := &BindInstanceListRequestBuilder{s: s}
	return b
}

func (b *BindInstanceListRequestBuilder) BindInstanceListBody(v *BindInstanceListBody) *BindInstanceListRequestBuilder {
	b.s.BindInstanceListBody = v
	return b
}

func (b *BindInstanceListRequestBuilder) Build() *BindInstanceListRequest {
	return b.s
}


