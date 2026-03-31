// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreateImageRequest struct {


	CreateImageBody *CreateImageBody `json:"createImageBody,omitempty"`
}

func (s CreateImageRequest) String() string {
	return utils.Beautify(s)
}

func (s CreateImageRequest) GoString() string {
	return s.String()
}

func (s CreateImageRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateImageRequest) SetCreateImageBody(v *CreateImageBody) *CreateImageRequest {
	s.CreateImageBody = v
	return s
}


type CreateImageRequestBuilder struct {
	s *CreateImageRequest
}

func NewCreateImageRequestBuilder() *CreateImageRequestBuilder {
	s := &CreateImageRequest{}
	b := &CreateImageRequestBuilder{s: s}
	return b
}

func (b *CreateImageRequestBuilder) CreateImageBody(v *CreateImageBody) *CreateImageRequestBuilder {
	b.s.CreateImageBody = v
	return b
}

func (b *CreateImageRequestBuilder) Build() *CreateImageRequest {
	return b.s
}


