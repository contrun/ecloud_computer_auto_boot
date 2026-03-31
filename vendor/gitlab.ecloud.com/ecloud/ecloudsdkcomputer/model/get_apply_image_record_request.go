// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetApplyImageRecordRequest struct {


	GetApplyImageRecordBody *GetApplyImageRecordBody `json:"getApplyImageRecordBody,omitempty"`
}

func (s GetApplyImageRecordRequest) String() string {
	return utils.Beautify(s)
}

func (s GetApplyImageRecordRequest) GoString() string {
	return s.String()
}

func (s GetApplyImageRecordRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetApplyImageRecordRequest) SetGetApplyImageRecordBody(v *GetApplyImageRecordBody) *GetApplyImageRecordRequest {
	s.GetApplyImageRecordBody = v
	return s
}


type GetApplyImageRecordRequestBuilder struct {
	s *GetApplyImageRecordRequest
}

func NewGetApplyImageRecordRequestBuilder() *GetApplyImageRecordRequestBuilder {
	s := &GetApplyImageRecordRequest{}
	b := &GetApplyImageRecordRequestBuilder{s: s}
	return b
}

func (b *GetApplyImageRecordRequestBuilder) GetApplyImageRecordBody(v *GetApplyImageRecordBody) *GetApplyImageRecordRequestBuilder {
	b.s.GetApplyImageRecordBody = v
	return b
}

func (b *GetApplyImageRecordRequestBuilder) Build() *GetApplyImageRecordRequest {
	return b.s
}


