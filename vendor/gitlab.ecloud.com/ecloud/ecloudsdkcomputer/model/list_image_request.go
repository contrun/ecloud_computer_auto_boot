// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListImageRequest struct {


	ListImageBody *ListImageBody `json:"listImageBody,omitempty"`
}

func (s ListImageRequest) String() string {
	return utils.Beautify(s)
}

func (s ListImageRequest) GoString() string {
	return s.String()
}

func (s ListImageRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListImageRequest) SetListImageBody(v *ListImageBody) *ListImageRequest {
	s.ListImageBody = v
	return s
}


type ListImageRequestBuilder struct {
	s *ListImageRequest
}

func NewListImageRequestBuilder() *ListImageRequestBuilder {
	s := &ListImageRequest{}
	b := &ListImageRequestBuilder{s: s}
	return b
}

func (b *ListImageRequestBuilder) ListImageBody(v *ListImageBody) *ListImageRequestBuilder {
	b.s.ListImageBody = v
	return b
}

func (b *ListImageRequestBuilder) Build() *ListImageRequest {
	return b.s
}


