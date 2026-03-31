// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ImageSharedRecord2OwnRequest struct {


	ImageSharedRecord2OwnBody *ImageSharedRecord2OwnBody `json:"imageSharedRecord2OwnBody,omitempty"`
}

func (s ImageSharedRecord2OwnRequest) String() string {
	return utils.Beautify(s)
}

func (s ImageSharedRecord2OwnRequest) GoString() string {
	return s.String()
}

func (s ImageSharedRecord2OwnRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ImageSharedRecord2OwnRequest) SetImageSharedRecord2OwnBody(v *ImageSharedRecord2OwnBody) *ImageSharedRecord2OwnRequest {
	s.ImageSharedRecord2OwnBody = v
	return s
}


type ImageSharedRecord2OwnRequestBuilder struct {
	s *ImageSharedRecord2OwnRequest
}

func NewImageSharedRecord2OwnRequestBuilder() *ImageSharedRecord2OwnRequestBuilder {
	s := &ImageSharedRecord2OwnRequest{}
	b := &ImageSharedRecord2OwnRequestBuilder{s: s}
	return b
}

func (b *ImageSharedRecord2OwnRequestBuilder) ImageSharedRecord2OwnBody(v *ImageSharedRecord2OwnBody) *ImageSharedRecord2OwnRequestBuilder {
	b.s.ImageSharedRecord2OwnBody = v
	return b
}

func (b *ImageSharedRecord2OwnRequestBuilder) Build() *ImageSharedRecord2OwnRequest {
	return b.s
}


