// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetRestartRecordListRequest struct {


	GetRestartRecordListBody *GetRestartRecordListBody `json:"getRestartRecordListBody,omitempty"`
}

func (s GetRestartRecordListRequest) String() string {
	return utils.Beautify(s)
}

func (s GetRestartRecordListRequest) GoString() string {
	return s.String()
}

func (s GetRestartRecordListRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetRestartRecordListRequest) SetGetRestartRecordListBody(v *GetRestartRecordListBody) *GetRestartRecordListRequest {
	s.GetRestartRecordListBody = v
	return s
}


type GetRestartRecordListRequestBuilder struct {
	s *GetRestartRecordListRequest
}

func NewGetRestartRecordListRequestBuilder() *GetRestartRecordListRequestBuilder {
	s := &GetRestartRecordListRequest{}
	b := &GetRestartRecordListRequestBuilder{s: s}
	return b
}

func (b *GetRestartRecordListRequestBuilder) GetRestartRecordListBody(v *GetRestartRecordListBody) *GetRestartRecordListRequestBuilder {
	b.s.GetRestartRecordListBody = v
	return b
}

func (b *GetRestartRecordListRequestBuilder) Build() *GetRestartRecordListRequest {
	return b.s
}


