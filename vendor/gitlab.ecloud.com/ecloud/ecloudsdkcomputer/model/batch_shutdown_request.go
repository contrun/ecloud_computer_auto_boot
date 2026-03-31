// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BatchShutdownRequest struct {


	BatchShutdownBody *BatchShutdownBody `json:"batchShutdownBody,omitempty"`
}

func (s BatchShutdownRequest) String() string {
	return utils.Beautify(s)
}

func (s BatchShutdownRequest) GoString() string {
	return s.String()
}

func (s BatchShutdownRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BatchShutdownRequest) SetBatchShutdownBody(v *BatchShutdownBody) *BatchShutdownRequest {
	s.BatchShutdownBody = v
	return s
}


type BatchShutdownRequestBuilder struct {
	s *BatchShutdownRequest
}

func NewBatchShutdownRequestBuilder() *BatchShutdownRequestBuilder {
	s := &BatchShutdownRequest{}
	b := &BatchShutdownRequestBuilder{s: s}
	return b
}

func (b *BatchShutdownRequestBuilder) BatchShutdownBody(v *BatchShutdownBody) *BatchShutdownRequestBuilder {
	b.s.BatchShutdownBody = v
	return b
}

func (b *BatchShutdownRequestBuilder) Build() *BatchShutdownRequest {
	return b.s
}


