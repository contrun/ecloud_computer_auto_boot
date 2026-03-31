// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetSubnetBody struct {
    position.Body
    // NAT网关资源实例id
	NatInstanceId *string `json:"natInstanceId,omitempty"`
    // vpcUid
	VpcUid *string `json:"vpcUid,omitempty"`
}

func (s GetSubnetBody) String() string {
	return utils.Beautify(s)
}

func (s GetSubnetBody) GoString() string {
	return s.String()
}

func (s GetSubnetBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetSubnetBody) SetNatInstanceId(v string) *GetSubnetBody {
	s.NatInstanceId = &v
	return s
}

func (s *GetSubnetBody) SetVpcUid(v string) *GetSubnetBody {
	s.VpcUid = &v
	return s
}


type GetSubnetBodyBuilder struct {
	s *GetSubnetBody
}

func NewGetSubnetBodyBuilder() *GetSubnetBodyBuilder {
	s := &GetSubnetBody{}
	b := &GetSubnetBodyBuilder{s: s}
	return b
}

func (b *GetSubnetBodyBuilder) NatInstanceId(v string) *GetSubnetBodyBuilder {
	b.s.NatInstanceId = &v
	return b
}

func (b *GetSubnetBodyBuilder) VpcUid(v string) *GetSubnetBodyBuilder {
	b.s.VpcUid = &v
	return b
}

func (b *GetSubnetBodyBuilder) Build() *GetSubnetBody {
	return b.s
}


