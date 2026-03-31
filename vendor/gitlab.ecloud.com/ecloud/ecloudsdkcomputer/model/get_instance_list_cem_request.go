// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetInstanceListCemRequest struct {


	GetInstanceListCemBody *GetInstanceListCemBody `json:"getInstanceListCemBody,omitempty"`
}

func (s GetInstanceListCemRequest) String() string {
	return utils.Beautify(s)
}

func (s GetInstanceListCemRequest) GoString() string {
	return s.String()
}

func (s GetInstanceListCemRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetInstanceListCemRequest) SetGetInstanceListCemBody(v *GetInstanceListCemBody) *GetInstanceListCemRequest {
	s.GetInstanceListCemBody = v
	return s
}


type GetInstanceListCemRequestBuilder struct {
	s *GetInstanceListCemRequest
}

func NewGetInstanceListCemRequestBuilder() *GetInstanceListCemRequestBuilder {
	s := &GetInstanceListCemRequest{}
	b := &GetInstanceListCemRequestBuilder{s: s}
	return b
}

func (b *GetInstanceListCemRequestBuilder) GetInstanceListCemBody(v *GetInstanceListCemBody) *GetInstanceListCemRequestBuilder {
	b.s.GetInstanceListCemBody = v
	return b
}

func (b *GetInstanceListCemRequestBuilder) Build() *GetInstanceListCemRequest {
	return b.s
}


