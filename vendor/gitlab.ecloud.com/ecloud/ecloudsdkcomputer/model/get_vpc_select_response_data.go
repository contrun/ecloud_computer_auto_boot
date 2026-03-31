// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetVpcSelectResponseData struct {


	VpcName *string `json:"vpcName,omitempty"`

	VpcUid *string `json:"vpcUid,omitempty"`

	TenantName *string `json:"tenantName,omitempty"`

	IsEcloud *bool `json:"isEcloud,omitempty"`

	TenantId *string `json:"tenantId,omitempty"`

	CreatedTime *string `json:"createdTime,omitempty"`

	Id *string `json:"id,omitempty"`

	MopUserId *string `json:"mopUserId,omitempty"`
}

func (s GetVpcSelectResponseData) String() string {
	return utils.Beautify(s)
}

func (s GetVpcSelectResponseData) GoString() string {
	return s.String()
}

func (s GetVpcSelectResponseData) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetVpcSelectResponseData) SetVpcName(v string) *GetVpcSelectResponseData {
	s.VpcName = &v
	return s
}

func (s *GetVpcSelectResponseData) SetVpcUid(v string) *GetVpcSelectResponseData {
	s.VpcUid = &v
	return s
}

func (s *GetVpcSelectResponseData) SetTenantName(v string) *GetVpcSelectResponseData {
	s.TenantName = &v
	return s
}

func (s *GetVpcSelectResponseData) SetIsEcloud(v bool) *GetVpcSelectResponseData {
	s.IsEcloud = &v
	return s
}

func (s *GetVpcSelectResponseData) SetTenantId(v string) *GetVpcSelectResponseData {
	s.TenantId = &v
	return s
}

func (s *GetVpcSelectResponseData) SetCreatedTime(v string) *GetVpcSelectResponseData {
	s.CreatedTime = &v
	return s
}

func (s *GetVpcSelectResponseData) SetId(v string) *GetVpcSelectResponseData {
	s.Id = &v
	return s
}

func (s *GetVpcSelectResponseData) SetMopUserId(v string) *GetVpcSelectResponseData {
	s.MopUserId = &v
	return s
}


type GetVpcSelectResponseDataBuilder struct {
	s *GetVpcSelectResponseData
}

func NewGetVpcSelectResponseDataBuilder() *GetVpcSelectResponseDataBuilder {
	s := &GetVpcSelectResponseData{}
	b := &GetVpcSelectResponseDataBuilder{s: s}
	return b
}

func (b *GetVpcSelectResponseDataBuilder) VpcName(v string) *GetVpcSelectResponseDataBuilder {
	b.s.VpcName = &v
	return b
}

func (b *GetVpcSelectResponseDataBuilder) VpcUid(v string) *GetVpcSelectResponseDataBuilder {
	b.s.VpcUid = &v
	return b
}

func (b *GetVpcSelectResponseDataBuilder) TenantName(v string) *GetVpcSelectResponseDataBuilder {
	b.s.TenantName = &v
	return b
}

func (b *GetVpcSelectResponseDataBuilder) IsEcloud(v bool) *GetVpcSelectResponseDataBuilder {
	b.s.IsEcloud = &v
	return b
}

func (b *GetVpcSelectResponseDataBuilder) TenantId(v string) *GetVpcSelectResponseDataBuilder {
	b.s.TenantId = &v
	return b
}

func (b *GetVpcSelectResponseDataBuilder) CreatedTime(v string) *GetVpcSelectResponseDataBuilder {
	b.s.CreatedTime = &v
	return b
}

func (b *GetVpcSelectResponseDataBuilder) Id(v string) *GetVpcSelectResponseDataBuilder {
	b.s.Id = &v
	return b
}

func (b *GetVpcSelectResponseDataBuilder) MopUserId(v string) *GetVpcSelectResponseDataBuilder {
	b.s.MopUserId = &v
	return b
}

func (b *GetVpcSelectResponseDataBuilder) Build() *GetVpcSelectResponseData {
	return b.s
}


