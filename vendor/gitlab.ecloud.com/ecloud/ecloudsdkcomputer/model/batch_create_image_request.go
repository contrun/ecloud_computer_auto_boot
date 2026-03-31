// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BatchCreateImageRequest struct {


	BatchCreateImageBody *BatchCreateImageBody `json:"batchCreateImageBody,omitempty"`
}

func (s BatchCreateImageRequest) String() string {
	return utils.Beautify(s)
}

func (s BatchCreateImageRequest) GoString() string {
	return s.String()
}

func (s BatchCreateImageRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BatchCreateImageRequest) SetBatchCreateImageBody(v *BatchCreateImageBody) *BatchCreateImageRequest {
	s.BatchCreateImageBody = v
	return s
}


type BatchCreateImageRequestBuilder struct {
	s *BatchCreateImageRequest
}

func NewBatchCreateImageRequestBuilder() *BatchCreateImageRequestBuilder {
	s := &BatchCreateImageRequest{}
	b := &BatchCreateImageRequestBuilder{s: s}
	return b
}

func (b *BatchCreateImageRequestBuilder) BatchCreateImageBody(v *BatchCreateImageBody) *BatchCreateImageRequestBuilder {
	b.s.BatchCreateImageBody = v
	return b
}

func (b *BatchCreateImageRequestBuilder) Build() *BatchCreateImageRequest {
	return b.s
}


