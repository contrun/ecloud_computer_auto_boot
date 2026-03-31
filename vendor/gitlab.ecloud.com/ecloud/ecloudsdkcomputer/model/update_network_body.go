// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateNetworkBody struct {
    position.Body
    // vpcUid
	VpcUid *string `json:"vpcUid,omitempty"`
    // 子网名称
	NetworkName *string `json:"networkName,omitempty"`
    // networkUid
	NetworkUid *string `json:"networkUid,omitempty"`
}

func (s UpdateNetworkBody) String() string {
	return utils.Beautify(s)
}

func (s UpdateNetworkBody) GoString() string {
	return s.String()
}

func (s UpdateNetworkBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateNetworkBody) SetVpcUid(v string) *UpdateNetworkBody {
	s.VpcUid = &v
	return s
}

func (s *UpdateNetworkBody) SetNetworkName(v string) *UpdateNetworkBody {
	s.NetworkName = &v
	return s
}

func (s *UpdateNetworkBody) SetNetworkUid(v string) *UpdateNetworkBody {
	s.NetworkUid = &v
	return s
}


type UpdateNetworkBodyBuilder struct {
	s *UpdateNetworkBody
}

func NewUpdateNetworkBodyBuilder() *UpdateNetworkBodyBuilder {
	s := &UpdateNetworkBody{}
	b := &UpdateNetworkBodyBuilder{s: s}
	return b
}

func (b *UpdateNetworkBodyBuilder) VpcUid(v string) *UpdateNetworkBodyBuilder {
	b.s.VpcUid = &v
	return b
}

func (b *UpdateNetworkBodyBuilder) NetworkName(v string) *UpdateNetworkBodyBuilder {
	b.s.NetworkName = &v
	return b
}

func (b *UpdateNetworkBodyBuilder) NetworkUid(v string) *UpdateNetworkBodyBuilder {
	b.s.NetworkUid = &v
	return b
}

func (b *UpdateNetworkBodyBuilder) Build() *UpdateNetworkBody {
	return b.s
}


