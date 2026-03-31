// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetUserInfoListRequest struct {


	GetUserInfoListBody *GetUserInfoListBody `json:"getUserInfoListBody,omitempty"`
}

func (s GetUserInfoListRequest) String() string {
	return utils.Beautify(s)
}

func (s GetUserInfoListRequest) GoString() string {
	return s.String()
}

func (s GetUserInfoListRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetUserInfoListRequest) SetGetUserInfoListBody(v *GetUserInfoListBody) *GetUserInfoListRequest {
	s.GetUserInfoListBody = v
	return s
}


type GetUserInfoListRequestBuilder struct {
	s *GetUserInfoListRequest
}

func NewGetUserInfoListRequestBuilder() *GetUserInfoListRequestBuilder {
	s := &GetUserInfoListRequest{}
	b := &GetUserInfoListRequestBuilder{s: s}
	return b
}

func (b *GetUserInfoListRequestBuilder) GetUserInfoListBody(v *GetUserInfoListBody) *GetUserInfoListRequestBuilder {
	b.s.GetUserInfoListBody = v
	return b
}

func (b *GetUserInfoListRequestBuilder) Build() *GetUserInfoListRequest {
	return b.s
}


