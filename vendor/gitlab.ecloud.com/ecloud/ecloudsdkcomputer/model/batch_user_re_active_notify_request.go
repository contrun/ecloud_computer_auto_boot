// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BatchUserReActiveNotifyRequest struct {


	BatchUserReActiveNotifyBody *BatchUserReActiveNotifyBody `json:"batchUserReActiveNotifyBody,omitempty"`
}

func (s BatchUserReActiveNotifyRequest) String() string {
	return utils.Beautify(s)
}

func (s BatchUserReActiveNotifyRequest) GoString() string {
	return s.String()
}

func (s BatchUserReActiveNotifyRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BatchUserReActiveNotifyRequest) SetBatchUserReActiveNotifyBody(v *BatchUserReActiveNotifyBody) *BatchUserReActiveNotifyRequest {
	s.BatchUserReActiveNotifyBody = v
	return s
}


type BatchUserReActiveNotifyRequestBuilder struct {
	s *BatchUserReActiveNotifyRequest
}

func NewBatchUserReActiveNotifyRequestBuilder() *BatchUserReActiveNotifyRequestBuilder {
	s := &BatchUserReActiveNotifyRequest{}
	b := &BatchUserReActiveNotifyRequestBuilder{s: s}
	return b
}

func (b *BatchUserReActiveNotifyRequestBuilder) BatchUserReActiveNotifyBody(v *BatchUserReActiveNotifyBody) *BatchUserReActiveNotifyRequestBuilder {
	b.s.BatchUserReActiveNotifyBody = v
	return b
}

func (b *BatchUserReActiveNotifyRequestBuilder) Build() *BatchUserReActiveNotifyRequest {
	return b.s
}


