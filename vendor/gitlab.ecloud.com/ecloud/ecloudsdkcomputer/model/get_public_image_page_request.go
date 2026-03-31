// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetPublicImagePageRequest struct {


	GetPublicImagePageBody *GetPublicImagePageBody `json:"getPublicImagePageBody,omitempty"`
}

func (s GetPublicImagePageRequest) String() string {
	return utils.Beautify(s)
}

func (s GetPublicImagePageRequest) GoString() string {
	return s.String()
}

func (s GetPublicImagePageRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetPublicImagePageRequest) SetGetPublicImagePageBody(v *GetPublicImagePageBody) *GetPublicImagePageRequest {
	s.GetPublicImagePageBody = v
	return s
}


type GetPublicImagePageRequestBuilder struct {
	s *GetPublicImagePageRequest
}

func NewGetPublicImagePageRequestBuilder() *GetPublicImagePageRequestBuilder {
	s := &GetPublicImagePageRequest{}
	b := &GetPublicImagePageRequestBuilder{s: s}
	return b
}

func (b *GetPublicImagePageRequestBuilder) GetPublicImagePageBody(v *GetPublicImagePageBody) *GetPublicImagePageRequestBuilder {
	b.s.GetPublicImagePageBody = v
	return b
}

func (b *GetPublicImagePageRequestBuilder) Build() *GetPublicImagePageRequest {
	return b.s
}


