// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type AddNetworkBody struct {
    position.Body
    // ipv6子网掩码
	CidrV6 *string `json:"cidrV6,omitempty"`
    // vpcUid
	VpcUid *string `json:"vpcUid,omitempty"`
    // 子网名称
	NetworkName *string `json:"networkName,omitempty"`
    // ipv4子网地址
	Cidr *string `json:"cidr,omitempty"`
}

func (s AddNetworkBody) String() string {
	return utils.Beautify(s)
}

func (s AddNetworkBody) GoString() string {
	return s.String()
}

func (s AddNetworkBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *AddNetworkBody) SetCidrV6(v string) *AddNetworkBody {
	s.CidrV6 = &v
	return s
}

func (s *AddNetworkBody) SetVpcUid(v string) *AddNetworkBody {
	s.VpcUid = &v
	return s
}

func (s *AddNetworkBody) SetNetworkName(v string) *AddNetworkBody {
	s.NetworkName = &v
	return s
}

func (s *AddNetworkBody) SetCidr(v string) *AddNetworkBody {
	s.Cidr = &v
	return s
}


type AddNetworkBodyBuilder struct {
	s *AddNetworkBody
}

func NewAddNetworkBodyBuilder() *AddNetworkBodyBuilder {
	s := &AddNetworkBody{}
	b := &AddNetworkBodyBuilder{s: s}
	return b
}

func (b *AddNetworkBodyBuilder) CidrV6(v string) *AddNetworkBodyBuilder {
	b.s.CidrV6 = &v
	return b
}

func (b *AddNetworkBodyBuilder) VpcUid(v string) *AddNetworkBodyBuilder {
	b.s.VpcUid = &v
	return b
}

func (b *AddNetworkBodyBuilder) NetworkName(v string) *AddNetworkBodyBuilder {
	b.s.NetworkName = &v
	return b
}

func (b *AddNetworkBodyBuilder) Cidr(v string) *AddNetworkBodyBuilder {
	b.s.Cidr = &v
	return b
}

func (b *AddNetworkBodyBuilder) Build() *AddNetworkBody {
	return b.s
}


