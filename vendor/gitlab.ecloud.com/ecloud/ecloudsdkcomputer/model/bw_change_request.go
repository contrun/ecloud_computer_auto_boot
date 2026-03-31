// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BwChangeRequest struct {


	BwChangeBody *BwChangeBody `json:"bwChangeBody,omitempty"`
}

func (s BwChangeRequest) String() string {
	return utils.Beautify(s)
}

func (s BwChangeRequest) GoString() string {
	return s.String()
}

func (s BwChangeRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BwChangeRequest) SetBwChangeBody(v *BwChangeBody) *BwChangeRequest {
	s.BwChangeBody = v
	return s
}


type BwChangeRequestBuilder struct {
	s *BwChangeRequest
}

func NewBwChangeRequestBuilder() *BwChangeRequestBuilder {
	s := &BwChangeRequest{}
	b := &BwChangeRequestBuilder{s: s}
	return b
}

func (b *BwChangeRequestBuilder) BwChangeBody(v *BwChangeBody) *BwChangeRequestBuilder {
	b.s.BwChangeBody = v
	return b
}

func (b *BwChangeRequestBuilder) Build() *BwChangeRequest {
	return b.s
}


