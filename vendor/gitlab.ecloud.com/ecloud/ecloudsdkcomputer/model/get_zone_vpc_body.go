// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetZoneVpcBody struct {
    position.Body
    // 云管可用区编码,示例值 CIDC-CORE-00
	ZoneCode *string `json:"zoneCode,omitempty"`
}

func (s GetZoneVpcBody) String() string {
	return utils.Beautify(s)
}

func (s GetZoneVpcBody) GoString() string {
	return s.String()
}

func (s GetZoneVpcBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetZoneVpcBody) SetZoneCode(v string) *GetZoneVpcBody {
	s.ZoneCode = &v
	return s
}


type GetZoneVpcBodyBuilder struct {
	s *GetZoneVpcBody
}

func NewGetZoneVpcBodyBuilder() *GetZoneVpcBodyBuilder {
	s := &GetZoneVpcBody{}
	b := &GetZoneVpcBodyBuilder{s: s}
	return b
}

func (b *GetZoneVpcBodyBuilder) ZoneCode(v string) *GetZoneVpcBodyBuilder {
	b.s.ZoneCode = &v
	return b
}

func (b *GetZoneVpcBodyBuilder) Build() *GetZoneVpcBody {
	return b.s
}


