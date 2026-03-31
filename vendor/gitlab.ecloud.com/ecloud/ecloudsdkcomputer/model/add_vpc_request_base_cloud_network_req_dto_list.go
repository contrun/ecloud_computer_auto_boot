// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type AddVpcRequestBaseCloudNetworkReqDtoList struct {

    // ipv6子网掩码
	CidrV6 *string `json:"cidrV6,omitempty"`
    // 子网名称
	NetworkName *string `json:"networkName,omitempty"`
    // ipv4子网地址
	Cidr *string `json:"cidr,omitempty"`
}

func (s AddVpcRequestBaseCloudNetworkReqDtoList) String() string {
	return utils.Beautify(s)
}

func (s AddVpcRequestBaseCloudNetworkReqDtoList) GoString() string {
	return s.String()
}

func (s AddVpcRequestBaseCloudNetworkReqDtoList) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *AddVpcRequestBaseCloudNetworkReqDtoList) SetCidrV6(v string) *AddVpcRequestBaseCloudNetworkReqDtoList {
	s.CidrV6 = &v
	return s
}

func (s *AddVpcRequestBaseCloudNetworkReqDtoList) SetNetworkName(v string) *AddVpcRequestBaseCloudNetworkReqDtoList {
	s.NetworkName = &v
	return s
}

func (s *AddVpcRequestBaseCloudNetworkReqDtoList) SetCidr(v string) *AddVpcRequestBaseCloudNetworkReqDtoList {
	s.Cidr = &v
	return s
}


type AddVpcRequestBaseCloudNetworkReqDtoListBuilder struct {
	s *AddVpcRequestBaseCloudNetworkReqDtoList
}

func NewAddVpcRequestBaseCloudNetworkReqDtoListBuilder() *AddVpcRequestBaseCloudNetworkReqDtoListBuilder {
	s := &AddVpcRequestBaseCloudNetworkReqDtoList{}
	b := &AddVpcRequestBaseCloudNetworkReqDtoListBuilder{s: s}
	return b
}

func (b *AddVpcRequestBaseCloudNetworkReqDtoListBuilder) CidrV6(v string) *AddVpcRequestBaseCloudNetworkReqDtoListBuilder {
	b.s.CidrV6 = &v
	return b
}

func (b *AddVpcRequestBaseCloudNetworkReqDtoListBuilder) NetworkName(v string) *AddVpcRequestBaseCloudNetworkReqDtoListBuilder {
	b.s.NetworkName = &v
	return b
}

func (b *AddVpcRequestBaseCloudNetworkReqDtoListBuilder) Cidr(v string) *AddVpcRequestBaseCloudNetworkReqDtoListBuilder {
	b.s.Cidr = &v
	return b
}

func (b *AddVpcRequestBaseCloudNetworkReqDtoListBuilder) Build() *AddVpcRequestBaseCloudNetworkReqDtoList {
	return b.s
}


