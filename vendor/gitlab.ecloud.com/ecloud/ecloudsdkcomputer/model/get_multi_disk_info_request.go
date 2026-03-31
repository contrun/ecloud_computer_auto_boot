// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetMultiDiskInfoRequest struct {


	GetMultiDiskInfoBody *GetMultiDiskInfoBody `json:"getMultiDiskInfoBody,omitempty"`
}

func (s GetMultiDiskInfoRequest) String() string {
	return utils.Beautify(s)
}

func (s GetMultiDiskInfoRequest) GoString() string {
	return s.String()
}

func (s GetMultiDiskInfoRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetMultiDiskInfoRequest) SetGetMultiDiskInfoBody(v *GetMultiDiskInfoBody) *GetMultiDiskInfoRequest {
	s.GetMultiDiskInfoBody = v
	return s
}


type GetMultiDiskInfoRequestBuilder struct {
	s *GetMultiDiskInfoRequest
}

func NewGetMultiDiskInfoRequestBuilder() *GetMultiDiskInfoRequestBuilder {
	s := &GetMultiDiskInfoRequest{}
	b := &GetMultiDiskInfoRequestBuilder{s: s}
	return b
}

func (b *GetMultiDiskInfoRequestBuilder) GetMultiDiskInfoBody(v *GetMultiDiskInfoBody) *GetMultiDiskInfoRequestBuilder {
	b.s.GetMultiDiskInfoBody = v
	return b
}

func (b *GetMultiDiskInfoRequestBuilder) Build() *GetMultiDiskInfoRequest {
	return b.s
}


