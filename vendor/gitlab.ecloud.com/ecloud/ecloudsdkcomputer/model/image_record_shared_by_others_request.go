// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ImageRecordSharedByOthersRequest struct {


	ImageRecordSharedByOthersBody *ImageRecordSharedByOthersBody `json:"ImageRecordSharedByOthersBody,omitempty"`
}

func (s ImageRecordSharedByOthersRequest) String() string {
	return utils.Beautify(s)
}

func (s ImageRecordSharedByOthersRequest) GoString() string {
	return s.String()
}

func (s ImageRecordSharedByOthersRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ImageRecordSharedByOthersRequest) SetImageRecordSharedByOthersBody(v *ImageRecordSharedByOthersBody) *ImageRecordSharedByOthersRequest {
	s.ImageRecordSharedByOthersBody = v
	return s
}


type ImageRecordSharedByOthersRequestBuilder struct {
	s *ImageRecordSharedByOthersRequest
}

func NewImageRecordSharedByOthersRequestBuilder() *ImageRecordSharedByOthersRequestBuilder {
	s := &ImageRecordSharedByOthersRequest{}
	b := &ImageRecordSharedByOthersRequestBuilder{s: s}
	return b
}

func (b *ImageRecordSharedByOthersRequestBuilder) ImageRecordSharedByOthersBody(v *ImageRecordSharedByOthersBody) *ImageRecordSharedByOthersRequestBuilder {
	b.s.ImageRecordSharedByOthersBody = v
	return b
}

func (b *ImageRecordSharedByOthersRequestBuilder) Build() *ImageRecordSharedByOthersRequest {
	return b.s
}


