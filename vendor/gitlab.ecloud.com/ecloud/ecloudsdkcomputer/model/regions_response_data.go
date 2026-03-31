// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type RegionsResponseData struct {

    // 编码
	Code *string `json:"code,omitempty"`
    // 名称
	Name *string `json:"name,omitempty"`
}

func (s RegionsResponseData) String() string {
	return utils.Beautify(s)
}

func (s RegionsResponseData) GoString() string {
	return s.String()
}

func (s RegionsResponseData) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *RegionsResponseData) SetCode(v string) *RegionsResponseData {
	s.Code = &v
	return s
}

func (s *RegionsResponseData) SetName(v string) *RegionsResponseData {
	s.Name = &v
	return s
}


type RegionsResponseDataBuilder struct {
	s *RegionsResponseData
}

func NewRegionsResponseDataBuilder() *RegionsResponseDataBuilder {
	s := &RegionsResponseData{}
	b := &RegionsResponseDataBuilder{s: s}
	return b
}

func (b *RegionsResponseDataBuilder) Code(v string) *RegionsResponseDataBuilder {
	b.s.Code = &v
	return b
}

func (b *RegionsResponseDataBuilder) Name(v string) *RegionsResponseDataBuilder {
	b.s.Name = &v
	return b
}

func (b *RegionsResponseDataBuilder) Build() *RegionsResponseData {
	return b.s
}


