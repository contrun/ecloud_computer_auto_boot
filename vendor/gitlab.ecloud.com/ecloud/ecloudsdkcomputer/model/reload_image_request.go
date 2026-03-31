// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ReloadImageRequest struct {


	ReloadImageBody *ReloadImageBody `json:"reloadImageBody,omitempty"`
}

func (s ReloadImageRequest) String() string {
	return utils.Beautify(s)
}

func (s ReloadImageRequest) GoString() string {
	return s.String()
}

func (s ReloadImageRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ReloadImageRequest) SetReloadImageBody(v *ReloadImageBody) *ReloadImageRequest {
	s.ReloadImageBody = v
	return s
}


type ReloadImageRequestBuilder struct {
	s *ReloadImageRequest
}

func NewReloadImageRequestBuilder() *ReloadImageRequestBuilder {
	s := &ReloadImageRequest{}
	b := &ReloadImageRequestBuilder{s: s}
	return b
}

func (b *ReloadImageRequestBuilder) ReloadImageBody(v *ReloadImageBody) *ReloadImageRequestBuilder {
	b.s.ReloadImageBody = v
	return b
}

func (b *ReloadImageRequestBuilder) Build() *ReloadImageRequest {
	return b.s
}


