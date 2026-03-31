// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type RegionsRequest struct {


	RegionsQuery *RegionsQuery `json:"regionsQuery,omitempty"`
}

func (s RegionsRequest) String() string {
	return utils.Beautify(s)
}

func (s RegionsRequest) GoString() string {
	return s.String()
}

func (s RegionsRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *RegionsRequest) SetRegionsQuery(v *RegionsQuery) *RegionsRequest {
	s.RegionsQuery = v
	return s
}


type RegionsRequestBuilder struct {
	s *RegionsRequest
}

func NewRegionsRequestBuilder() *RegionsRequestBuilder {
	s := &RegionsRequest{}
	b := &RegionsRequestBuilder{s: s}
	return b
}

func (b *RegionsRequestBuilder) RegionsQuery(v *RegionsQuery) *RegionsRequestBuilder {
	b.s.RegionsQuery = v
	return b
}

func (b *RegionsRequestBuilder) Build() *RegionsRequest {
	return b.s
}


