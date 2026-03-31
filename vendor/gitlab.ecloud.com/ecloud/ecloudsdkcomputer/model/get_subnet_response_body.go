// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetSubnetResponseBody struct {


	ModifiedTime *string `json:"modifiedTime,omitempty"`
    // 修改DNS服务器地址接口可访问性，1代表修改完成，0代表修改中，2代表修改失败
	Access *string `json:"access,omitempty"`

	IsDeleted *string `json:"isDeleted,omitempty"`
    // subnetUid
	SubnetUid *string `json:"subnetUid,omitempty"`

	CreatedTime *string `json:"createdTime,omitempty"`
    // cidr地址块，可存V4和V6网段
	Cidr *string `json:"cidr,omitempty"`

	Id *int64 `json:"id,omitempty"`
    // 所属networkUid
	NetworkUid *string `json:"networkUid,omitempty"`
    // dns服务器地址
	DnsAddr *string `json:"dnsAddr,omitempty"`
    // subnet名称
	SubnetName *string `json:"subnetName,omitempty"`
    // 子网类型，ipV4或ipV6
	SubnetType *string `json:"subnetType,omitempty"`
}

func (s GetSubnetResponseBody) String() string {
	return utils.Beautify(s)
}

func (s GetSubnetResponseBody) GoString() string {
	return s.String()
}

func (s GetSubnetResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetSubnetResponseBody) SetModifiedTime(v string) *GetSubnetResponseBody {
	s.ModifiedTime = &v
	return s
}

func (s *GetSubnetResponseBody) SetAccess(v string) *GetSubnetResponseBody {
	s.Access = &v
	return s
}

func (s *GetSubnetResponseBody) SetIsDeleted(v string) *GetSubnetResponseBody {
	s.IsDeleted = &v
	return s
}

func (s *GetSubnetResponseBody) SetSubnetUid(v string) *GetSubnetResponseBody {
	s.SubnetUid = &v
	return s
}

func (s *GetSubnetResponseBody) SetCreatedTime(v string) *GetSubnetResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *GetSubnetResponseBody) SetCidr(v string) *GetSubnetResponseBody {
	s.Cidr = &v
	return s
}

func (s *GetSubnetResponseBody) SetId(v int64) *GetSubnetResponseBody {
	s.Id = &v
	return s
}

func (s *GetSubnetResponseBody) SetNetworkUid(v string) *GetSubnetResponseBody {
	s.NetworkUid = &v
	return s
}

func (s *GetSubnetResponseBody) SetDnsAddr(v string) *GetSubnetResponseBody {
	s.DnsAddr = &v
	return s
}

func (s *GetSubnetResponseBody) SetSubnetName(v string) *GetSubnetResponseBody {
	s.SubnetName = &v
	return s
}

func (s *GetSubnetResponseBody) SetSubnetType(v string) *GetSubnetResponseBody {
	s.SubnetType = &v
	return s
}


type GetSubnetResponseBodyBuilder struct {
	s *GetSubnetResponseBody
}

func NewGetSubnetResponseBodyBuilder() *GetSubnetResponseBodyBuilder {
	s := &GetSubnetResponseBody{}
	b := &GetSubnetResponseBodyBuilder{s: s}
	return b
}

func (b *GetSubnetResponseBodyBuilder) ModifiedTime(v string) *GetSubnetResponseBodyBuilder {
	b.s.ModifiedTime = &v
	return b
}

func (b *GetSubnetResponseBodyBuilder) Access(v string) *GetSubnetResponseBodyBuilder {
	b.s.Access = &v
	return b
}

func (b *GetSubnetResponseBodyBuilder) IsDeleted(v string) *GetSubnetResponseBodyBuilder {
	b.s.IsDeleted = &v
	return b
}

func (b *GetSubnetResponseBodyBuilder) SubnetUid(v string) *GetSubnetResponseBodyBuilder {
	b.s.SubnetUid = &v
	return b
}

func (b *GetSubnetResponseBodyBuilder) CreatedTime(v string) *GetSubnetResponseBodyBuilder {
	b.s.CreatedTime = &v
	return b
}

func (b *GetSubnetResponseBodyBuilder) Cidr(v string) *GetSubnetResponseBodyBuilder {
	b.s.Cidr = &v
	return b
}

func (b *GetSubnetResponseBodyBuilder) Id(v int64) *GetSubnetResponseBodyBuilder {
	b.s.Id = &v
	return b
}

func (b *GetSubnetResponseBodyBuilder) NetworkUid(v string) *GetSubnetResponseBodyBuilder {
	b.s.NetworkUid = &v
	return b
}

func (b *GetSubnetResponseBodyBuilder) DnsAddr(v string) *GetSubnetResponseBodyBuilder {
	b.s.DnsAddr = &v
	return b
}

func (b *GetSubnetResponseBodyBuilder) SubnetName(v string) *GetSubnetResponseBodyBuilder {
	b.s.SubnetName = &v
	return b
}

func (b *GetSubnetResponseBodyBuilder) SubnetType(v string) *GetSubnetResponseBodyBuilder {
	b.s.SubnetType = &v
	return b
}

func (b *GetSubnetResponseBodyBuilder) Build() *GetSubnetResponseBody {
	return b.s
}


