// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetNetworkResponseData struct {

    // 子网名称
	NetworkName *string `json:"networkName,omitempty"`
    // ipv4子网地址
	CidrIpv4 *string `json:"cidrIpv4,omitempty"`

	AccessV6 *int32 `json:"accessV6,omitempty"`
    // ipv6子网地址
	CidrIpv6 *string `json:"cidrIpv6,omitempty"`
    // 子网来源
	RequestSource *string `json:"requestSource,omitempty"`

	AccessV4 *int32 `json:"accessV4,omitempty"`
    // vpcUid
	VpcUid *string `json:"vpcUid,omitempty"`
    // 创建时间
	CreatedTime *string `json:"createdTime,omitempty"`

	PoolCode *string `json:"poolCode,omitempty"`
    // ipv4dns服务器信息
	DnsAddrV4 *string `json:"dnsAddrV4,omitempty"`
    // networkUid
	NetworkUid *string `json:"networkUid,omitempty"`

	Region *string `json:"region,omitempty"`
    // ipv6dns服务器信息
	DnsAddrV6 *string `json:"dnsAddrV6,omitempty"`
    // 子网状态
	Status *string `json:"status,omitempty"`
}

func (s GetNetworkResponseData) String() string {
	return utils.Beautify(s)
}

func (s GetNetworkResponseData) GoString() string {
	return s.String()
}

func (s GetNetworkResponseData) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetNetworkResponseData) SetNetworkName(v string) *GetNetworkResponseData {
	s.NetworkName = &v
	return s
}

func (s *GetNetworkResponseData) SetCidrIpv4(v string) *GetNetworkResponseData {
	s.CidrIpv4 = &v
	return s
}

func (s *GetNetworkResponseData) SetAccessV6(v int32) *GetNetworkResponseData {
	s.AccessV6 = &v
	return s
}

func (s *GetNetworkResponseData) SetCidrIpv6(v string) *GetNetworkResponseData {
	s.CidrIpv6 = &v
	return s
}

func (s *GetNetworkResponseData) SetRequestSource(v string) *GetNetworkResponseData {
	s.RequestSource = &v
	return s
}

func (s *GetNetworkResponseData) SetAccessV4(v int32) *GetNetworkResponseData {
	s.AccessV4 = &v
	return s
}

func (s *GetNetworkResponseData) SetVpcUid(v string) *GetNetworkResponseData {
	s.VpcUid = &v
	return s
}

func (s *GetNetworkResponseData) SetCreatedTime(v string) *GetNetworkResponseData {
	s.CreatedTime = &v
	return s
}

func (s *GetNetworkResponseData) SetPoolCode(v string) *GetNetworkResponseData {
	s.PoolCode = &v
	return s
}

func (s *GetNetworkResponseData) SetDnsAddrV4(v string) *GetNetworkResponseData {
	s.DnsAddrV4 = &v
	return s
}

func (s *GetNetworkResponseData) SetNetworkUid(v string) *GetNetworkResponseData {
	s.NetworkUid = &v
	return s
}

func (s *GetNetworkResponseData) SetRegion(v string) *GetNetworkResponseData {
	s.Region = &v
	return s
}

func (s *GetNetworkResponseData) SetDnsAddrV6(v string) *GetNetworkResponseData {
	s.DnsAddrV6 = &v
	return s
}

func (s *GetNetworkResponseData) SetStatus(v string) *GetNetworkResponseData {
	s.Status = &v
	return s
}


type GetNetworkResponseDataBuilder struct {
	s *GetNetworkResponseData
}

func NewGetNetworkResponseDataBuilder() *GetNetworkResponseDataBuilder {
	s := &GetNetworkResponseData{}
	b := &GetNetworkResponseDataBuilder{s: s}
	return b
}

func (b *GetNetworkResponseDataBuilder) NetworkName(v string) *GetNetworkResponseDataBuilder {
	b.s.NetworkName = &v
	return b
}

func (b *GetNetworkResponseDataBuilder) CidrIpv4(v string) *GetNetworkResponseDataBuilder {
	b.s.CidrIpv4 = &v
	return b
}

func (b *GetNetworkResponseDataBuilder) AccessV6(v int32) *GetNetworkResponseDataBuilder {
	b.s.AccessV6 = &v
	return b
}

func (b *GetNetworkResponseDataBuilder) CidrIpv6(v string) *GetNetworkResponseDataBuilder {
	b.s.CidrIpv6 = &v
	return b
}

func (b *GetNetworkResponseDataBuilder) RequestSource(v string) *GetNetworkResponseDataBuilder {
	b.s.RequestSource = &v
	return b
}

func (b *GetNetworkResponseDataBuilder) AccessV4(v int32) *GetNetworkResponseDataBuilder {
	b.s.AccessV4 = &v
	return b
}

func (b *GetNetworkResponseDataBuilder) VpcUid(v string) *GetNetworkResponseDataBuilder {
	b.s.VpcUid = &v
	return b
}

func (b *GetNetworkResponseDataBuilder) CreatedTime(v string) *GetNetworkResponseDataBuilder {
	b.s.CreatedTime = &v
	return b
}

func (b *GetNetworkResponseDataBuilder) PoolCode(v string) *GetNetworkResponseDataBuilder {
	b.s.PoolCode = &v
	return b
}

func (b *GetNetworkResponseDataBuilder) DnsAddrV4(v string) *GetNetworkResponseDataBuilder {
	b.s.DnsAddrV4 = &v
	return b
}

func (b *GetNetworkResponseDataBuilder) NetworkUid(v string) *GetNetworkResponseDataBuilder {
	b.s.NetworkUid = &v
	return b
}

func (b *GetNetworkResponseDataBuilder) Region(v string) *GetNetworkResponseDataBuilder {
	b.s.Region = &v
	return b
}

func (b *GetNetworkResponseDataBuilder) DnsAddrV6(v string) *GetNetworkResponseDataBuilder {
	b.s.DnsAddrV6 = &v
	return b
}

func (b *GetNetworkResponseDataBuilder) Status(v string) *GetNetworkResponseDataBuilder {
	b.s.Status = &v
	return b
}

func (b *GetNetworkResponseDataBuilder) Build() *GetNetworkResponseData {
	return b.s
}


