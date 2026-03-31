// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetIsBindUserInfoListRequest struct {


	GetIsBindUserInfoListBody *GetIsBindUserInfoListBody `json:"getIsBindUserInfoListBody,omitempty"`
}

func (s GetIsBindUserInfoListRequest) String() string {
	return utils.Beautify(s)
}

func (s GetIsBindUserInfoListRequest) GoString() string {
	return s.String()
}

func (s GetIsBindUserInfoListRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetIsBindUserInfoListRequest) SetGetIsBindUserInfoListBody(v *GetIsBindUserInfoListBody) *GetIsBindUserInfoListRequest {
	s.GetIsBindUserInfoListBody = v
	return s
}


type GetIsBindUserInfoListRequestBuilder struct {
	s *GetIsBindUserInfoListRequest
}

func NewGetIsBindUserInfoListRequestBuilder() *GetIsBindUserInfoListRequestBuilder {
	s := &GetIsBindUserInfoListRequest{}
	b := &GetIsBindUserInfoListRequestBuilder{s: s}
	return b
}

func (b *GetIsBindUserInfoListRequestBuilder) GetIsBindUserInfoListBody(v *GetIsBindUserInfoListBody) *GetIsBindUserInfoListRequestBuilder {
	b.s.GetIsBindUserInfoListBody = v
	return b
}

func (b *GetIsBindUserInfoListRequestBuilder) Build() *GetIsBindUserInfoListRequest {
	return b.s
}


