// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetResourceDiskInfoRequest struct {


	GetResourceDiskInfoBody *GetResourceDiskInfoBody `json:"getResourceDiskInfoBody,omitempty"`
}

func (s GetResourceDiskInfoRequest) String() string {
	return utils.Beautify(s)
}

func (s GetResourceDiskInfoRequest) GoString() string {
	return s.String()
}

func (s GetResourceDiskInfoRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetResourceDiskInfoRequest) SetGetResourceDiskInfoBody(v *GetResourceDiskInfoBody) *GetResourceDiskInfoRequest {
	s.GetResourceDiskInfoBody = v
	return s
}


type GetResourceDiskInfoRequestBuilder struct {
	s *GetResourceDiskInfoRequest
}

func NewGetResourceDiskInfoRequestBuilder() *GetResourceDiskInfoRequestBuilder {
	s := &GetResourceDiskInfoRequest{}
	b := &GetResourceDiskInfoRequestBuilder{s: s}
	return b
}

func (b *GetResourceDiskInfoRequestBuilder) GetResourceDiskInfoBody(v *GetResourceDiskInfoBody) *GetResourceDiskInfoRequestBuilder {
	b.s.GetResourceDiskInfoBody = v
	return b
}

func (b *GetResourceDiskInfoRequestBuilder) Build() *GetResourceDiskInfoRequest {
	return b.s
}


