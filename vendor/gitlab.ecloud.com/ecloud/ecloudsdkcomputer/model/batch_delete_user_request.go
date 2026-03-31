// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BatchDeleteUserRequest struct {


	BatchDeleteUserBody *BatchDeleteUserBody `json:"batchDeleteUserBody,omitempty"`
}

func (s BatchDeleteUserRequest) String() string {
	return utils.Beautify(s)
}

func (s BatchDeleteUserRequest) GoString() string {
	return s.String()
}

func (s BatchDeleteUserRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BatchDeleteUserRequest) SetBatchDeleteUserBody(v *BatchDeleteUserBody) *BatchDeleteUserRequest {
	s.BatchDeleteUserBody = v
	return s
}


type BatchDeleteUserRequestBuilder struct {
	s *BatchDeleteUserRequest
}

func NewBatchDeleteUserRequestBuilder() *BatchDeleteUserRequestBuilder {
	s := &BatchDeleteUserRequest{}
	b := &BatchDeleteUserRequestBuilder{s: s}
	return b
}

func (b *BatchDeleteUserRequestBuilder) BatchDeleteUserBody(v *BatchDeleteUserBody) *BatchDeleteUserRequestBuilder {
	b.s.BatchDeleteUserBody = v
	return b
}

func (b *BatchDeleteUserRequestBuilder) Build() *BatchDeleteUserRequest {
	return b.s
}


