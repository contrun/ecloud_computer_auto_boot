// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ReloadByPublicImageRequest struct {


	ReloadByPublicImageBody *ReloadByPublicImageBody `json:"reloadByPublicImageBody,omitempty"`
}

func (s ReloadByPublicImageRequest) String() string {
	return utils.Beautify(s)
}

func (s ReloadByPublicImageRequest) GoString() string {
	return s.String()
}

func (s ReloadByPublicImageRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ReloadByPublicImageRequest) SetReloadByPublicImageBody(v *ReloadByPublicImageBody) *ReloadByPublicImageRequest {
	s.ReloadByPublicImageBody = v
	return s
}


type ReloadByPublicImageRequestBuilder struct {
	s *ReloadByPublicImageRequest
}

func NewReloadByPublicImageRequestBuilder() *ReloadByPublicImageRequestBuilder {
	s := &ReloadByPublicImageRequest{}
	b := &ReloadByPublicImageRequestBuilder{s: s}
	return b
}

func (b *ReloadByPublicImageRequestBuilder) ReloadByPublicImageBody(v *ReloadByPublicImageBody) *ReloadByPublicImageRequestBuilder {
	b.s.ReloadByPublicImageBody = v
	return b
}

func (b *ReloadByPublicImageRequestBuilder) Build() *ReloadByPublicImageRequest {
	return b.s
}


