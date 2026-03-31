// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetShutdownRecordListRequest struct {


	GetShutdownRecordListBody *GetShutdownRecordListBody `json:"getShutdownRecordListBody,omitempty"`
}

func (s GetShutdownRecordListRequest) String() string {
	return utils.Beautify(s)
}

func (s GetShutdownRecordListRequest) GoString() string {
	return s.String()
}

func (s GetShutdownRecordListRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetShutdownRecordListRequest) SetGetShutdownRecordListBody(v *GetShutdownRecordListBody) *GetShutdownRecordListRequest {
	s.GetShutdownRecordListBody = v
	return s
}


type GetShutdownRecordListRequestBuilder struct {
	s *GetShutdownRecordListRequest
}

func NewGetShutdownRecordListRequestBuilder() *GetShutdownRecordListRequestBuilder {
	s := &GetShutdownRecordListRequest{}
	b := &GetShutdownRecordListRequestBuilder{s: s}
	return b
}

func (b *GetShutdownRecordListRequestBuilder) GetShutdownRecordListBody(v *GetShutdownRecordListBody) *GetShutdownRecordListRequestBuilder {
	b.s.GetShutdownRecordListBody = v
	return b
}

func (b *GetShutdownRecordListRequestBuilder) Build() *GetShutdownRecordListRequest {
	return b.s
}


