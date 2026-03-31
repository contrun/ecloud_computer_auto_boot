// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BatchAvailableRequest struct {


	BatchAvailableBody *BatchAvailableBody `json:"batchAvailableBody,omitempty"`
}

func (s BatchAvailableRequest) String() string {
	return utils.Beautify(s)
}

func (s BatchAvailableRequest) GoString() string {
	return s.String()
}

func (s BatchAvailableRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BatchAvailableRequest) SetBatchAvailableBody(v *BatchAvailableBody) *BatchAvailableRequest {
	s.BatchAvailableBody = v
	return s
}


type BatchAvailableRequestBuilder struct {
	s *BatchAvailableRequest
}

func NewBatchAvailableRequestBuilder() *BatchAvailableRequestBuilder {
	s := &BatchAvailableRequest{}
	b := &BatchAvailableRequestBuilder{s: s}
	return b
}

func (b *BatchAvailableRequestBuilder) BatchAvailableBody(v *BatchAvailableBody) *BatchAvailableRequestBuilder {
	b.s.BatchAvailableBody = v
	return b
}

func (b *BatchAvailableRequestBuilder) Build() *BatchAvailableRequest {
	return b.s
}


