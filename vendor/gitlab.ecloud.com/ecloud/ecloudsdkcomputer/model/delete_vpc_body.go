// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type DeleteVpcBody struct {
    position.Body
    // vpcid
	VpcUid *string `json:"vpcUid,omitempty"`
}

func (s DeleteVpcBody) String() string {
	return utils.Beautify(s)
}

func (s DeleteVpcBody) GoString() string {
	return s.String()
}

func (s DeleteVpcBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *DeleteVpcBody) SetVpcUid(v string) *DeleteVpcBody {
	s.VpcUid = &v
	return s
}


type DeleteVpcBodyBuilder struct {
	s *DeleteVpcBody
}

func NewDeleteVpcBodyBuilder() *DeleteVpcBodyBuilder {
	s := &DeleteVpcBody{}
	b := &DeleteVpcBodyBuilder{s: s}
	return b
}

func (b *DeleteVpcBodyBuilder) VpcUid(v string) *DeleteVpcBodyBuilder {
	b.s.VpcUid = &v
	return b
}

func (b *DeleteVpcBodyBuilder) Build() *DeleteVpcBody {
	return b.s
}


