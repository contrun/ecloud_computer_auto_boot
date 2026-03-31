// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type EditCompanyTagRequest struct {


	EditCompanyTagBody *EditCompanyTagBody `json:"editCompanyTagBody,omitempty"`
}

func (s EditCompanyTagRequest) String() string {
	return utils.Beautify(s)
}

func (s EditCompanyTagRequest) GoString() string {
	return s.String()
}

func (s EditCompanyTagRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *EditCompanyTagRequest) SetEditCompanyTagBody(v *EditCompanyTagBody) *EditCompanyTagRequest {
	s.EditCompanyTagBody = v
	return s
}


type EditCompanyTagRequestBuilder struct {
	s *EditCompanyTagRequest
}

func NewEditCompanyTagRequestBuilder() *EditCompanyTagRequestBuilder {
	s := &EditCompanyTagRequest{}
	b := &EditCompanyTagRequestBuilder{s: s}
	return b
}

func (b *EditCompanyTagRequestBuilder) EditCompanyTagBody(v *EditCompanyTagBody) *EditCompanyTagRequestBuilder {
	b.s.EditCompanyTagBody = v
	return b
}

func (b *EditCompanyTagRequestBuilder) Build() *EditCompanyTagRequest {
	return b.s
}


