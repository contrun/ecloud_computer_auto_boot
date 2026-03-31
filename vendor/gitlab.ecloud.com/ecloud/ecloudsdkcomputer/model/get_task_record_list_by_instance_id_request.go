// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetTaskRecordListByInstanceIdRequest struct {


	GetTaskRecordListByInstanceIdBody *GetTaskRecordListByInstanceIdBody `json:"getTaskRecordListByInstanceIdBody,omitempty"`
}

func (s GetTaskRecordListByInstanceIdRequest) String() string {
	return utils.Beautify(s)
}

func (s GetTaskRecordListByInstanceIdRequest) GoString() string {
	return s.String()
}

func (s GetTaskRecordListByInstanceIdRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetTaskRecordListByInstanceIdRequest) SetGetTaskRecordListByInstanceIdBody(v *GetTaskRecordListByInstanceIdBody) *GetTaskRecordListByInstanceIdRequest {
	s.GetTaskRecordListByInstanceIdBody = v
	return s
}


type GetTaskRecordListByInstanceIdRequestBuilder struct {
	s *GetTaskRecordListByInstanceIdRequest
}

func NewGetTaskRecordListByInstanceIdRequestBuilder() *GetTaskRecordListByInstanceIdRequestBuilder {
	s := &GetTaskRecordListByInstanceIdRequest{}
	b := &GetTaskRecordListByInstanceIdRequestBuilder{s: s}
	return b
}

func (b *GetTaskRecordListByInstanceIdRequestBuilder) GetTaskRecordListByInstanceIdBody(v *GetTaskRecordListByInstanceIdBody) *GetTaskRecordListByInstanceIdRequestBuilder {
	b.s.GetTaskRecordListByInstanceIdBody = v
	return b
}

func (b *GetTaskRecordListByInstanceIdRequestBuilder) Build() *GetTaskRecordListByInstanceIdRequest {
	return b.s
}


