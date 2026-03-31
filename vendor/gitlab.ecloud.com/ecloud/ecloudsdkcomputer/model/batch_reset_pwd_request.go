// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BatchResetPwdRequest struct {


	BatchResetPwdBody *BatchResetPwdBody `json:"batchResetPwdBody,omitempty"`
}

func (s BatchResetPwdRequest) String() string {
	return utils.Beautify(s)
}

func (s BatchResetPwdRequest) GoString() string {
	return s.String()
}

func (s BatchResetPwdRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BatchResetPwdRequest) SetBatchResetPwdBody(v *BatchResetPwdBody) *BatchResetPwdRequest {
	s.BatchResetPwdBody = v
	return s
}


type BatchResetPwdRequestBuilder struct {
	s *BatchResetPwdRequest
}

func NewBatchResetPwdRequestBuilder() *BatchResetPwdRequestBuilder {
	s := &BatchResetPwdRequest{}
	b := &BatchResetPwdRequestBuilder{s: s}
	return b
}

func (b *BatchResetPwdRequestBuilder) BatchResetPwdBody(v *BatchResetPwdBody) *BatchResetPwdRequestBuilder {
	b.s.BatchResetPwdBody = v
	return b
}

func (b *BatchResetPwdRequestBuilder) Build() *BatchResetPwdRequest {
	return b.s
}


