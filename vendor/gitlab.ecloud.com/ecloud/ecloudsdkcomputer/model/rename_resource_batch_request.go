// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type RenameResourceBatchRequest struct {


	RenameResourceBatchBody *RenameResourceBatchBody `json:"renameResourceBatchBody,omitempty"`
}

func (s RenameResourceBatchRequest) String() string {
	return utils.Beautify(s)
}

func (s RenameResourceBatchRequest) GoString() string {
	return s.String()
}

func (s RenameResourceBatchRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *RenameResourceBatchRequest) SetRenameResourceBatchBody(v *RenameResourceBatchBody) *RenameResourceBatchRequest {
	s.RenameResourceBatchBody = v
	return s
}


type RenameResourceBatchRequestBuilder struct {
	s *RenameResourceBatchRequest
}

func NewRenameResourceBatchRequestBuilder() *RenameResourceBatchRequestBuilder {
	s := &RenameResourceBatchRequest{}
	b := &RenameResourceBatchRequestBuilder{s: s}
	return b
}

func (b *RenameResourceBatchRequestBuilder) RenameResourceBatchBody(v *RenameResourceBatchBody) *RenameResourceBatchRequestBuilder {
	b.s.RenameResourceBatchBody = v
	return b
}

func (b *RenameResourceBatchRequestBuilder) Build() *RenameResourceBatchRequest {
	return b.s
}


