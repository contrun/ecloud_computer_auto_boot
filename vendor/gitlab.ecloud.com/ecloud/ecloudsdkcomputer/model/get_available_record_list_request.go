// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetAvailableRecordListRequest struct {


	GetAvailableRecordListBody *GetAvailableRecordListBody `json:"getAvailableRecordListBody,omitempty"`
}

func (s GetAvailableRecordListRequest) String() string {
	return utils.Beautify(s)
}

func (s GetAvailableRecordListRequest) GoString() string {
	return s.String()
}

func (s GetAvailableRecordListRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetAvailableRecordListRequest) SetGetAvailableRecordListBody(v *GetAvailableRecordListBody) *GetAvailableRecordListRequest {
	s.GetAvailableRecordListBody = v
	return s
}


type GetAvailableRecordListRequestBuilder struct {
	s *GetAvailableRecordListRequest
}

func NewGetAvailableRecordListRequestBuilder() *GetAvailableRecordListRequestBuilder {
	s := &GetAvailableRecordListRequest{}
	b := &GetAvailableRecordListRequestBuilder{s: s}
	return b
}

func (b *GetAvailableRecordListRequestBuilder) GetAvailableRecordListBody(v *GetAvailableRecordListBody) *GetAvailableRecordListRequestBuilder {
	b.s.GetAvailableRecordListBody = v
	return b
}

func (b *GetAvailableRecordListRequestBuilder) Build() *GetAvailableRecordListRequest {
	return b.s
}


