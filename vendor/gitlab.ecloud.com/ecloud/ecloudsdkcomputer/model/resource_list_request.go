// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ResourceListRequest struct {


	ResourceListBody *ResourceListBody `json:"resourceListBody,omitempty"`
}

func (s ResourceListRequest) String() string {
	return utils.Beautify(s)
}

func (s ResourceListRequest) GoString() string {
	return s.String()
}

func (s ResourceListRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ResourceListRequest) SetResourceListBody(v *ResourceListBody) *ResourceListRequest {
	s.ResourceListBody = v
	return s
}


type ResourceListRequestBuilder struct {
	s *ResourceListRequest
}

func NewResourceListRequestBuilder() *ResourceListRequestBuilder {
	s := &ResourceListRequest{}
	b := &ResourceListRequestBuilder{s: s}
	return b
}

func (b *ResourceListRequestBuilder) ResourceListBody(v *ResourceListBody) *ResourceListRequestBuilder {
	b.s.ResourceListBody = v
	return b
}

func (b *ResourceListRequestBuilder) Build() *ResourceListRequest {
	return b.s
}


