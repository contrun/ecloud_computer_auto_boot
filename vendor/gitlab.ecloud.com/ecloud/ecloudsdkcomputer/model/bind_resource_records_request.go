// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BindResourceRecordsRequest struct {


	BindResourceRecordsBody *BindResourceRecordsBody `json:"bindResourceRecordsBody,omitempty"`
}

func (s BindResourceRecordsRequest) String() string {
	return utils.Beautify(s)
}

func (s BindResourceRecordsRequest) GoString() string {
	return s.String()
}

func (s BindResourceRecordsRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BindResourceRecordsRequest) SetBindResourceRecordsBody(v *BindResourceRecordsBody) *BindResourceRecordsRequest {
	s.BindResourceRecordsBody = v
	return s
}


type BindResourceRecordsRequestBuilder struct {
	s *BindResourceRecordsRequest
}

func NewBindResourceRecordsRequestBuilder() *BindResourceRecordsRequestBuilder {
	s := &BindResourceRecordsRequest{}
	b := &BindResourceRecordsRequestBuilder{s: s}
	return b
}

func (b *BindResourceRecordsRequestBuilder) BindResourceRecordsBody(v *BindResourceRecordsBody) *BindResourceRecordsRequestBuilder {
	b.s.BindResourceRecordsBody = v
	return b
}

func (b *BindResourceRecordsRequestBuilder) Build() *BindResourceRecordsRequest {
	return b.s
}


