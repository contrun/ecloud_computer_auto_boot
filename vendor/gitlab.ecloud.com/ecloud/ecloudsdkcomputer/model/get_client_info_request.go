// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetClientInfoRequest struct {


	GetClientInfoBody *GetClientInfoBody `json:"getClientInfoBody,omitempty"`
}

func (s GetClientInfoRequest) String() string {
	return utils.Beautify(s)
}

func (s GetClientInfoRequest) GoString() string {
	return s.String()
}

func (s GetClientInfoRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetClientInfoRequest) SetGetClientInfoBody(v *GetClientInfoBody) *GetClientInfoRequest {
	s.GetClientInfoBody = v
	return s
}


type GetClientInfoRequestBuilder struct {
	s *GetClientInfoRequest
}

func NewGetClientInfoRequestBuilder() *GetClientInfoRequestBuilder {
	s := &GetClientInfoRequest{}
	b := &GetClientInfoRequestBuilder{s: s}
	return b
}

func (b *GetClientInfoRequestBuilder) GetClientInfoBody(v *GetClientInfoBody) *GetClientInfoRequestBuilder {
	b.s.GetClientInfoBody = v
	return b
}

func (b *GetClientInfoRequestBuilder) Build() *GetClientInfoRequest {
	return b.s
}


