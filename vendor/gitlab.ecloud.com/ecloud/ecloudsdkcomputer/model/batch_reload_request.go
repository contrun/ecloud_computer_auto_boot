// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BatchReloadRequest struct {


	BatchReloadBody *BatchReloadBody `json:"batchReloadBody,omitempty"`
}

func (s BatchReloadRequest) String() string {
	return utils.Beautify(s)
}

func (s BatchReloadRequest) GoString() string {
	return s.String()
}

func (s BatchReloadRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BatchReloadRequest) SetBatchReloadBody(v *BatchReloadBody) *BatchReloadRequest {
	s.BatchReloadBody = v
	return s
}


type BatchReloadRequestBuilder struct {
	s *BatchReloadRequest
}

func NewBatchReloadRequestBuilder() *BatchReloadRequestBuilder {
	s := &BatchReloadRequest{}
	b := &BatchReloadRequestBuilder{s: s}
	return b
}

func (b *BatchReloadRequestBuilder) BatchReloadBody(v *BatchReloadBody) *BatchReloadRequestBuilder {
	b.s.BatchReloadBody = v
	return b
}

func (b *BatchReloadRequestBuilder) Build() *BatchReloadRequest {
	return b.s
}


