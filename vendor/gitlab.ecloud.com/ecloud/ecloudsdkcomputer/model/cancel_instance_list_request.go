// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CancelInstanceListRequest struct {


	CancelInstanceListBody *CancelInstanceListBody `json:"cancelInstanceListBody,omitempty"`
}

func (s CancelInstanceListRequest) String() string {
	return utils.Beautify(s)
}

func (s CancelInstanceListRequest) GoString() string {
	return s.String()
}

func (s CancelInstanceListRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CancelInstanceListRequest) SetCancelInstanceListBody(v *CancelInstanceListBody) *CancelInstanceListRequest {
	s.CancelInstanceListBody = v
	return s
}


type CancelInstanceListRequestBuilder struct {
	s *CancelInstanceListRequest
}

func NewCancelInstanceListRequestBuilder() *CancelInstanceListRequestBuilder {
	s := &CancelInstanceListRequest{}
	b := &CancelInstanceListRequestBuilder{s: s}
	return b
}

func (b *CancelInstanceListRequestBuilder) CancelInstanceListBody(v *CancelInstanceListBody) *CancelInstanceListRequestBuilder {
	b.s.CancelInstanceListBody = v
	return b
}

func (b *CancelInstanceListRequestBuilder) Build() *CancelInstanceListRequest {
	return b.s
}


