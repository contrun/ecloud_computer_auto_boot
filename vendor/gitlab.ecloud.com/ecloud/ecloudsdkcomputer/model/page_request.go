// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type PageRequest struct {


	PageBody *PageBody `json:"pageBody,omitempty"`
}

func (s PageRequest) String() string {
	return utils.Beautify(s)
}

func (s PageRequest) GoString() string {
	return s.String()
}

func (s PageRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *PageRequest) SetPageBody(v *PageBody) *PageRequest {
	s.PageBody = v
	return s
}


type PageRequestBuilder struct {
	s *PageRequest
}

func NewPageRequestBuilder() *PageRequestBuilder {
	s := &PageRequest{}
	b := &PageRequestBuilder{s: s}
	return b
}

func (b *PageRequestBuilder) PageBody(v *PageBody) *PageRequestBuilder {
	b.s.PageBody = v
	return b
}

func (b *PageRequestBuilder) Build() *PageRequest {
	return b.s
}


