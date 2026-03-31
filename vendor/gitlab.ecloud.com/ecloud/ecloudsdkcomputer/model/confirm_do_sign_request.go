// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ConfirmDoSignRequest struct {


	ConfirmDoSignBody *ConfirmDoSignBody `json:"confirmDoSignBody,omitempty"`
}

func (s ConfirmDoSignRequest) String() string {
	return utils.Beautify(s)
}

func (s ConfirmDoSignRequest) GoString() string {
	return s.String()
}

func (s ConfirmDoSignRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ConfirmDoSignRequest) SetConfirmDoSignBody(v *ConfirmDoSignBody) *ConfirmDoSignRequest {
	s.ConfirmDoSignBody = v
	return s
}


type ConfirmDoSignRequestBuilder struct {
	s *ConfirmDoSignRequest
}

func NewConfirmDoSignRequestBuilder() *ConfirmDoSignRequestBuilder {
	s := &ConfirmDoSignRequest{}
	b := &ConfirmDoSignRequestBuilder{s: s}
	return b
}

func (b *ConfirmDoSignRequestBuilder) ConfirmDoSignBody(v *ConfirmDoSignBody) *ConfirmDoSignRequestBuilder {
	b.s.ConfirmDoSignBody = v
	return b
}

func (b *ConfirmDoSignRequestBuilder) Build() *ConfirmDoSignRequest {
	return b.s
}


