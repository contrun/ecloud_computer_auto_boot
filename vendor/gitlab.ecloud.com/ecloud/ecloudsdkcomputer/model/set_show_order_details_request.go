// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type SetShowOrderDetailsRequest struct {


	SetShowOrderDetailsBody *SetShowOrderDetailsBody `json:"setShowOrderDetailsBody,omitempty"`
}

func (s SetShowOrderDetailsRequest) String() string {
	return utils.Beautify(s)
}

func (s SetShowOrderDetailsRequest) GoString() string {
	return s.String()
}

func (s SetShowOrderDetailsRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *SetShowOrderDetailsRequest) SetSetShowOrderDetailsBody(v *SetShowOrderDetailsBody) *SetShowOrderDetailsRequest {
	s.SetShowOrderDetailsBody = v
	return s
}


type SetShowOrderDetailsRequestBuilder struct {
	s *SetShowOrderDetailsRequest
}

func NewSetShowOrderDetailsRequestBuilder() *SetShowOrderDetailsRequestBuilder {
	s := &SetShowOrderDetailsRequest{}
	b := &SetShowOrderDetailsRequestBuilder{s: s}
	return b
}

func (b *SetShowOrderDetailsRequestBuilder) SetShowOrderDetailsBody(v *SetShowOrderDetailsBody) *SetShowOrderDetailsRequestBuilder {
	b.s.SetShowOrderDetailsBody = v
	return b
}

func (b *SetShowOrderDetailsRequestBuilder) Build() *SetShowOrderDetailsRequest {
	return b.s
}


