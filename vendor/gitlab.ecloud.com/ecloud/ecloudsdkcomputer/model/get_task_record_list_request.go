// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetTaskRecordListRequest struct {


	GetTaskRecordListBody *GetTaskRecordListBody `json:"getTaskRecordListBody,omitempty"`
}

func (s GetTaskRecordListRequest) String() string {
	return utils.Beautify(s)
}

func (s GetTaskRecordListRequest) GoString() string {
	return s.String()
}

func (s GetTaskRecordListRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetTaskRecordListRequest) SetGetTaskRecordListBody(v *GetTaskRecordListBody) *GetTaskRecordListRequest {
	s.GetTaskRecordListBody = v
	return s
}


type GetTaskRecordListRequestBuilder struct {
	s *GetTaskRecordListRequest
}

func NewGetTaskRecordListRequestBuilder() *GetTaskRecordListRequestBuilder {
	s := &GetTaskRecordListRequest{}
	b := &GetTaskRecordListRequestBuilder{s: s}
	return b
}

func (b *GetTaskRecordListRequestBuilder) GetTaskRecordListBody(v *GetTaskRecordListBody) *GetTaskRecordListRequestBuilder {
	b.s.GetTaskRecordListBody = v
	return b
}

func (b *GetTaskRecordListRequestBuilder) Build() *GetTaskRecordListRequest {
	return b.s
}


