// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetCloudbandwidthBody struct {
    position.Body
    // cmdb可用区
	Region *string `json:"region,omitempty"`
}

func (s GetCloudbandwidthBody) String() string {
	return utils.Beautify(s)
}

func (s GetCloudbandwidthBody) GoString() string {
	return s.String()
}

func (s GetCloudbandwidthBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetCloudbandwidthBody) SetRegion(v string) *GetCloudbandwidthBody {
	s.Region = &v
	return s
}


type GetCloudbandwidthBodyBuilder struct {
	s *GetCloudbandwidthBody
}

func NewGetCloudbandwidthBodyBuilder() *GetCloudbandwidthBodyBuilder {
	s := &GetCloudbandwidthBody{}
	b := &GetCloudbandwidthBodyBuilder{s: s}
	return b
}

func (b *GetCloudbandwidthBodyBuilder) Region(v string) *GetCloudbandwidthBodyBuilder {
	b.s.Region = &v
	return b
}

func (b *GetCloudbandwidthBodyBuilder) Build() *GetCloudbandwidthBody {
	return b.s
}


