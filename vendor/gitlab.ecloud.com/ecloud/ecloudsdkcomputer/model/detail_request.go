// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DetailRequest struct {


	DetailBody *DetailBody `json:"detailBody,omitempty"`
}

func (s DetailRequest) String() string {
	return utils.Beautify(s)
}

func (s DetailRequest) GoString() string {
	return s.String()
}

func (s DetailRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DetailRequest) SetDetailBody(v *DetailBody) *DetailRequest {
	s.DetailBody = v
	return s
}


type DetailRequestBuilder struct {
	s *DetailRequest
}

func NewDetailRequestBuilder() *DetailRequestBuilder {
	s := &DetailRequest{}
	b := &DetailRequestBuilder{s: s}
	return b
}

func (b *DetailRequestBuilder) DetailBody(v *DetailBody) *DetailRequestBuilder {
	b.s.DetailBody = v
	return b
}

func (b *DetailRequestBuilder) Build() *DetailRequest {
	return b.s
}


