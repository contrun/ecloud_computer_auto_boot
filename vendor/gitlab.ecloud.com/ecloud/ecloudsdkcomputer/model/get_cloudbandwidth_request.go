// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetCloudbandwidthRequest struct {


	GetCloudbandwidthBody *GetCloudbandwidthBody `json:"getCloudbandwidthBody,omitempty"`
}

func (s GetCloudbandwidthRequest) String() string {
	return utils.Beautify(s)
}

func (s GetCloudbandwidthRequest) GoString() string {
	return s.String()
}

func (s GetCloudbandwidthRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetCloudbandwidthRequest) SetGetCloudbandwidthBody(v *GetCloudbandwidthBody) *GetCloudbandwidthRequest {
	s.GetCloudbandwidthBody = v
	return s
}


type GetCloudbandwidthRequestBuilder struct {
	s *GetCloudbandwidthRequest
}

func NewGetCloudbandwidthRequestBuilder() *GetCloudbandwidthRequestBuilder {
	s := &GetCloudbandwidthRequest{}
	b := &GetCloudbandwidthRequestBuilder{s: s}
	return b
}

func (b *GetCloudbandwidthRequestBuilder) GetCloudbandwidthBody(v *GetCloudbandwidthBody) *GetCloudbandwidthRequestBuilder {
	b.s.GetCloudbandwidthBody = v
	return b
}

func (b *GetCloudbandwidthRequestBuilder) Build() *GetCloudbandwidthRequest {
	return b.s
}


