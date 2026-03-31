// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetZoneVpcResponseData struct {

    // 厂商编码
	CompanyCode *string `json:"companyCode,omitempty"`
    // vpc名称
	VpcName *string `json:"vpcName,omitempty"`
    // vpc ID编码
	VpcUid *string `json:"vpcUid,omitempty"`
    // 可用区编码
	Region *string `json:"region,omitempty"`
}

func (s GetZoneVpcResponseData) String() string {
	return utils.Beautify(s)
}

func (s GetZoneVpcResponseData) GoString() string {
	return s.String()
}

func (s GetZoneVpcResponseData) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetZoneVpcResponseData) SetCompanyCode(v string) *GetZoneVpcResponseData {
	s.CompanyCode = &v
	return s
}

func (s *GetZoneVpcResponseData) SetVpcName(v string) *GetZoneVpcResponseData {
	s.VpcName = &v
	return s
}

func (s *GetZoneVpcResponseData) SetVpcUid(v string) *GetZoneVpcResponseData {
	s.VpcUid = &v
	return s
}

func (s *GetZoneVpcResponseData) SetRegion(v string) *GetZoneVpcResponseData {
	s.Region = &v
	return s
}


type GetZoneVpcResponseDataBuilder struct {
	s *GetZoneVpcResponseData
}

func NewGetZoneVpcResponseDataBuilder() *GetZoneVpcResponseDataBuilder {
	s := &GetZoneVpcResponseData{}
	b := &GetZoneVpcResponseDataBuilder{s: s}
	return b
}

func (b *GetZoneVpcResponseDataBuilder) CompanyCode(v string) *GetZoneVpcResponseDataBuilder {
	b.s.CompanyCode = &v
	return b
}

func (b *GetZoneVpcResponseDataBuilder) VpcName(v string) *GetZoneVpcResponseDataBuilder {
	b.s.VpcName = &v
	return b
}

func (b *GetZoneVpcResponseDataBuilder) VpcUid(v string) *GetZoneVpcResponseDataBuilder {
	b.s.VpcUid = &v
	return b
}

func (b *GetZoneVpcResponseDataBuilder) Region(v string) *GetZoneVpcResponseDataBuilder {
	b.s.Region = &v
	return b
}

func (b *GetZoneVpcResponseDataBuilder) Build() *GetZoneVpcResponseData {
	return b.s
}


