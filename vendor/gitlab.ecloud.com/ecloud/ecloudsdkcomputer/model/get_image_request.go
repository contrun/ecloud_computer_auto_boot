// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetImageRequest struct {


	GetImageBody *GetImageBody `json:"getImageBody,omitempty"`
}

func (s GetImageRequest) String() string {
	return utils.Beautify(s)
}

func (s GetImageRequest) GoString() string {
	return s.String()
}

func (s GetImageRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetImageRequest) SetGetImageBody(v *GetImageBody) *GetImageRequest {
	s.GetImageBody = v
	return s
}


type GetImageRequestBuilder struct {
	s *GetImageRequest
}

func NewGetImageRequestBuilder() *GetImageRequestBuilder {
	s := &GetImageRequest{}
	b := &GetImageRequestBuilder{s: s}
	return b
}

func (b *GetImageRequestBuilder) GetImageBody(v *GetImageBody) *GetImageRequestBuilder {
	b.s.GetImageBody = v
	return b
}

func (b *GetImageRequestBuilder) Build() *GetImageRequest {
	return b.s
}


