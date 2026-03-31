// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetNetworkSelectResponseData struct {

    // isEcloud
	IsEcloud *bool `json:"isEcloud,omitempty"`
    // 子网名称
	NetworkName *string `json:"networkName,omitempty"`
    // id
	Id *string `json:"id,omitempty"`
    // networkUid
	NetworkUid *string `json:"networkUid,omitempty"`
}

func (s GetNetworkSelectResponseData) String() string {
	return utils.Beautify(s)
}

func (s GetNetworkSelectResponseData) GoString() string {
	return s.String()
}

func (s GetNetworkSelectResponseData) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetNetworkSelectResponseData) SetIsEcloud(v bool) *GetNetworkSelectResponseData {
	s.IsEcloud = &v
	return s
}

func (s *GetNetworkSelectResponseData) SetNetworkName(v string) *GetNetworkSelectResponseData {
	s.NetworkName = &v
	return s
}

func (s *GetNetworkSelectResponseData) SetId(v string) *GetNetworkSelectResponseData {
	s.Id = &v
	return s
}

func (s *GetNetworkSelectResponseData) SetNetworkUid(v string) *GetNetworkSelectResponseData {
	s.NetworkUid = &v
	return s
}


type GetNetworkSelectResponseDataBuilder struct {
	s *GetNetworkSelectResponseData
}

func NewGetNetworkSelectResponseDataBuilder() *GetNetworkSelectResponseDataBuilder {
	s := &GetNetworkSelectResponseData{}
	b := &GetNetworkSelectResponseDataBuilder{s: s}
	return b
}

func (b *GetNetworkSelectResponseDataBuilder) IsEcloud(v bool) *GetNetworkSelectResponseDataBuilder {
	b.s.IsEcloud = &v
	return b
}

func (b *GetNetworkSelectResponseDataBuilder) NetworkName(v string) *GetNetworkSelectResponseDataBuilder {
	b.s.NetworkName = &v
	return b
}

func (b *GetNetworkSelectResponseDataBuilder) Id(v string) *GetNetworkSelectResponseDataBuilder {
	b.s.Id = &v
	return b
}

func (b *GetNetworkSelectResponseDataBuilder) NetworkUid(v string) *GetNetworkSelectResponseDataBuilder {
	b.s.NetworkUid = &v
	return b
}

func (b *GetNetworkSelectResponseDataBuilder) Build() *GetNetworkSelectResponseData {
	return b.s
}


