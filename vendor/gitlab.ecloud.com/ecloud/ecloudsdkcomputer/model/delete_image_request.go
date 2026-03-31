// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeleteImageRequest struct {


	DeleteImageBody *DeleteImageBody `json:"deleteImageBody,omitempty"`
}

func (s DeleteImageRequest) String() string {
	return utils.Beautify(s)
}

func (s DeleteImageRequest) GoString() string {
	return s.String()
}

func (s DeleteImageRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteImageRequest) SetDeleteImageBody(v *DeleteImageBody) *DeleteImageRequest {
	s.DeleteImageBody = v
	return s
}


type DeleteImageRequestBuilder struct {
	s *DeleteImageRequest
}

func NewDeleteImageRequestBuilder() *DeleteImageRequestBuilder {
	s := &DeleteImageRequest{}
	b := &DeleteImageRequestBuilder{s: s}
	return b
}

func (b *DeleteImageRequestBuilder) DeleteImageBody(v *DeleteImageBody) *DeleteImageRequestBuilder {
	b.s.DeleteImageBody = v
	return b
}

func (b *DeleteImageRequestBuilder) Build() *DeleteImageRequest {
	return b.s
}


