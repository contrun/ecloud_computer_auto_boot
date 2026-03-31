// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeleteNetworkBody struct {
    position.Body
    // vpcUid
	VpcUid *string `json:"vpcUid,omitempty"`
    // networkUid
	NetworkUid *string `json:"networkUid,omitempty"`
}

func (s DeleteNetworkBody) String() string {
	return utils.Beautify(s)
}

func (s DeleteNetworkBody) GoString() string {
	return s.String()
}

func (s DeleteNetworkBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteNetworkBody) SetVpcUid(v string) *DeleteNetworkBody {
	s.VpcUid = &v
	return s
}

func (s *DeleteNetworkBody) SetNetworkUid(v string) *DeleteNetworkBody {
	s.NetworkUid = &v
	return s
}


type DeleteNetworkBodyBuilder struct {
	s *DeleteNetworkBody
}

func NewDeleteNetworkBodyBuilder() *DeleteNetworkBodyBuilder {
	s := &DeleteNetworkBody{}
	b := &DeleteNetworkBodyBuilder{s: s}
	return b
}

func (b *DeleteNetworkBodyBuilder) VpcUid(v string) *DeleteNetworkBodyBuilder {
	b.s.VpcUid = &v
	return b
}

func (b *DeleteNetworkBodyBuilder) NetworkUid(v string) *DeleteNetworkBodyBuilder {
	b.s.NetworkUid = &v
	return b
}

func (b *DeleteNetworkBodyBuilder) Build() *DeleteNetworkBody {
	return b.s
}


