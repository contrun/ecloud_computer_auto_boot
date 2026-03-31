// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetNetworkSelectBody struct {
    position.Body
    // vpcUid
	VpcUid *string `json:"vpcUid,omitempty"`
}

func (s GetNetworkSelectBody) String() string {
	return utils.Beautify(s)
}

func (s GetNetworkSelectBody) GoString() string {
	return s.String()
}

func (s GetNetworkSelectBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetNetworkSelectBody) SetVpcUid(v string) *GetNetworkSelectBody {
	s.VpcUid = &v
	return s
}


type GetNetworkSelectBodyBuilder struct {
	s *GetNetworkSelectBody
}

func NewGetNetworkSelectBodyBuilder() *GetNetworkSelectBodyBuilder {
	s := &GetNetworkSelectBody{}
	b := &GetNetworkSelectBodyBuilder{s: s}
	return b
}

func (b *GetNetworkSelectBodyBuilder) VpcUid(v string) *GetNetworkSelectBodyBuilder {
	b.s.VpcUid = &v
	return b
}

func (b *GetNetworkSelectBodyBuilder) Build() *GetNetworkSelectBody {
	return b.s
}


