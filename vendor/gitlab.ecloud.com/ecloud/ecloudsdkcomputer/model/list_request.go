// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ListRequest struct {


	ListBody *ListBody `json:"listBody,omitempty"`
}

func (s ListRequest) String() string {
	return utils.Beautify(s)
}

func (s ListRequest) GoString() string {
	return s.String()
}

func (s ListRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ListRequest) SetListBody(v *ListBody) *ListRequest {
	s.ListBody = v
	return s
}


type ListRequestBuilder struct {
	s *ListRequest
}

func NewListRequestBuilder() *ListRequestBuilder {
	s := &ListRequest{}
	b := &ListRequestBuilder{s: s}
	return b
}

func (b *ListRequestBuilder) ListBody(v *ListBody) *ListRequestBuilder {
	b.s.ListBody = v
	return b
}

func (b *ListRequestBuilder) Build() *ListRequest {
	return b.s
}


