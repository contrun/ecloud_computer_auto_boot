// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BatchReloadByImageRequest struct {


	BatchReloadByImageBody *BatchReloadByImageBody `json:"batchReloadByImageBody,omitempty"`
}

func (s BatchReloadByImageRequest) String() string {
	return utils.Beautify(s)
}

func (s BatchReloadByImageRequest) GoString() string {
	return s.String()
}

func (s BatchReloadByImageRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BatchReloadByImageRequest) SetBatchReloadByImageBody(v *BatchReloadByImageBody) *BatchReloadByImageRequest {
	s.BatchReloadByImageBody = v
	return s
}


type BatchReloadByImageRequestBuilder struct {
	s *BatchReloadByImageRequest
}

func NewBatchReloadByImageRequestBuilder() *BatchReloadByImageRequestBuilder {
	s := &BatchReloadByImageRequest{}
	b := &BatchReloadByImageRequestBuilder{s: s}
	return b
}

func (b *BatchReloadByImageRequestBuilder) BatchReloadByImageBody(v *BatchReloadByImageBody) *BatchReloadByImageRequestBuilder {
	b.s.BatchReloadByImageBody = v
	return b
}

func (b *BatchReloadByImageRequestBuilder) Build() *BatchReloadByImageRequest {
	return b.s
}


