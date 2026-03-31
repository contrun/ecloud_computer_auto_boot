// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ImageRecordSharedWithOthersRequest struct {


	ImageRecordSharedWithOthersBody *ImageRecordSharedWithOthersBody `json:"ImageRecordSharedWithOthersBody,omitempty"`
}

func (s ImageRecordSharedWithOthersRequest) String() string {
	return utils.Beautify(s)
}

func (s ImageRecordSharedWithOthersRequest) GoString() string {
	return s.String()
}

func (s ImageRecordSharedWithOthersRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ImageRecordSharedWithOthersRequest) SetImageRecordSharedWithOthersBody(v *ImageRecordSharedWithOthersBody) *ImageRecordSharedWithOthersRequest {
	s.ImageRecordSharedWithOthersBody = v
	return s
}


type ImageRecordSharedWithOthersRequestBuilder struct {
	s *ImageRecordSharedWithOthersRequest
}

func NewImageRecordSharedWithOthersRequestBuilder() *ImageRecordSharedWithOthersRequestBuilder {
	s := &ImageRecordSharedWithOthersRequest{}
	b := &ImageRecordSharedWithOthersRequestBuilder{s: s}
	return b
}

func (b *ImageRecordSharedWithOthersRequestBuilder) ImageRecordSharedWithOthersBody(v *ImageRecordSharedWithOthersBody) *ImageRecordSharedWithOthersRequestBuilder {
	b.s.ImageRecordSharedWithOthersBody = v
	return b
}

func (b *ImageRecordSharedWithOthersRequestBuilder) Build() *ImageRecordSharedWithOthersRequest {
	return b.s
}


