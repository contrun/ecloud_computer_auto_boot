// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type MultiDiskChangeRequest struct {


	MultiDiskChangeBody *MultiDiskChangeBody `json:"multiDiskChangeBody,omitempty"`
}

func (s MultiDiskChangeRequest) String() string {
	return utils.Beautify(s)
}

func (s MultiDiskChangeRequest) GoString() string {
	return s.String()
}

func (s MultiDiskChangeRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *MultiDiskChangeRequest) SetMultiDiskChangeBody(v *MultiDiskChangeBody) *MultiDiskChangeRequest {
	s.MultiDiskChangeBody = v
	return s
}


type MultiDiskChangeRequestBuilder struct {
	s *MultiDiskChangeRequest
}

func NewMultiDiskChangeRequestBuilder() *MultiDiskChangeRequestBuilder {
	s := &MultiDiskChangeRequest{}
	b := &MultiDiskChangeRequestBuilder{s: s}
	return b
}

func (b *MultiDiskChangeRequestBuilder) MultiDiskChangeBody(v *MultiDiskChangeBody) *MultiDiskChangeRequestBuilder {
	b.s.MultiDiskChangeBody = v
	return b
}

func (b *MultiDiskChangeRequestBuilder) Build() *MultiDiskChangeRequest {
	return b.s
}


