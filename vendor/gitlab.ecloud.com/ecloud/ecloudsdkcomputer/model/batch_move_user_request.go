// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BatchMoveUserRequest struct {


	BatchMoveUserBody *BatchMoveUserBody `json:"batchMoveUserBody,omitempty"`
}

func (s BatchMoveUserRequest) String() string {
	return utils.Beautify(s)
}

func (s BatchMoveUserRequest) GoString() string {
	return s.String()
}

func (s BatchMoveUserRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BatchMoveUserRequest) SetBatchMoveUserBody(v *BatchMoveUserBody) *BatchMoveUserRequest {
	s.BatchMoveUserBody = v
	return s
}


type BatchMoveUserRequestBuilder struct {
	s *BatchMoveUserRequest
}

func NewBatchMoveUserRequestBuilder() *BatchMoveUserRequestBuilder {
	s := &BatchMoveUserRequest{}
	b := &BatchMoveUserRequestBuilder{s: s}
	return b
}

func (b *BatchMoveUserRequestBuilder) BatchMoveUserBody(v *BatchMoveUserBody) *BatchMoveUserRequestBuilder {
	b.s.BatchMoveUserBody = v
	return b
}

func (b *BatchMoveUserRequestBuilder) Build() *BatchMoveUserRequest {
	return b.s
}


