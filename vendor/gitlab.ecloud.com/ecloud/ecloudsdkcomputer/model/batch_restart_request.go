// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BatchRestartRequest struct {


	BatchRestartBody *BatchRestartBody `json:"batchRestartBody,omitempty"`
}

func (s BatchRestartRequest) String() string {
	return utils.Beautify(s)
}

func (s BatchRestartRequest) GoString() string {
	return s.String()
}

func (s BatchRestartRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BatchRestartRequest) SetBatchRestartBody(v *BatchRestartBody) *BatchRestartRequest {
	s.BatchRestartBody = v
	return s
}


type BatchRestartRequestBuilder struct {
	s *BatchRestartRequest
}

func NewBatchRestartRequestBuilder() *BatchRestartRequestBuilder {
	s := &BatchRestartRequest{}
	b := &BatchRestartRequestBuilder{s: s}
	return b
}

func (b *BatchRestartRequestBuilder) BatchRestartBody(v *BatchRestartBody) *BatchRestartRequestBuilder {
	b.s.BatchRestartBody = v
	return b
}

func (b *BatchRestartRequestBuilder) Build() *BatchRestartRequest {
	return b.s
}


