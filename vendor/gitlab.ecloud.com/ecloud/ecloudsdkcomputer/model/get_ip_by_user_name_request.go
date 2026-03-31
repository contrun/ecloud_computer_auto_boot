// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetIPByUserNameRequest struct {


	GetIPByUserNameQuery *GetIPByUserNameQuery `json:"getIPByUserNameQuery,omitempty"`
}

func (s GetIPByUserNameRequest) String() string {
	return utils.Beautify(s)
}

func (s GetIPByUserNameRequest) GoString() string {
	return s.String()
}

func (s GetIPByUserNameRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetIPByUserNameRequest) SetGetIPByUserNameQuery(v *GetIPByUserNameQuery) *GetIPByUserNameRequest {
	s.GetIPByUserNameQuery = v
	return s
}


type GetIPByUserNameRequestBuilder struct {
	s *GetIPByUserNameRequest
}

func NewGetIPByUserNameRequestBuilder() *GetIPByUserNameRequestBuilder {
	s := &GetIPByUserNameRequest{}
	b := &GetIPByUserNameRequestBuilder{s: s}
	return b
}

func (b *GetIPByUserNameRequestBuilder) GetIPByUserNameQuery(v *GetIPByUserNameQuery) *GetIPByUserNameRequestBuilder {
	b.s.GetIPByUserNameQuery = v
	return b
}

func (b *GetIPByUserNameRequestBuilder) Build() *GetIPByUserNameRequest {
	return b.s
}


