// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateVpcBody struct {
    position.Body
    // vpc名称
	VpcName *string `json:"vpcName,omitempty"`
    // vpcid
	VpcUid *string `json:"vpcUid,omitempty"`
    // vpc描述
	Description *string `json:"description,omitempty"`
}

func (s UpdateVpcBody) String() string {
	return utils.Beautify(s)
}

func (s UpdateVpcBody) GoString() string {
	return s.String()
}

func (s UpdateVpcBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateVpcBody) SetVpcName(v string) *UpdateVpcBody {
	s.VpcName = &v
	return s
}

func (s *UpdateVpcBody) SetVpcUid(v string) *UpdateVpcBody {
	s.VpcUid = &v
	return s
}

func (s *UpdateVpcBody) SetDescription(v string) *UpdateVpcBody {
	s.Description = &v
	return s
}


type UpdateVpcBodyBuilder struct {
	s *UpdateVpcBody
}

func NewUpdateVpcBodyBuilder() *UpdateVpcBodyBuilder {
	s := &UpdateVpcBody{}
	b := &UpdateVpcBodyBuilder{s: s}
	return b
}

func (b *UpdateVpcBodyBuilder) VpcName(v string) *UpdateVpcBodyBuilder {
	b.s.VpcName = &v
	return b
}

func (b *UpdateVpcBodyBuilder) VpcUid(v string) *UpdateVpcBodyBuilder {
	b.s.VpcUid = &v
	return b
}

func (b *UpdateVpcBodyBuilder) Description(v string) *UpdateVpcBodyBuilder {
	b.s.Description = &v
	return b
}

func (b *UpdateVpcBodyBuilder) Build() *UpdateVpcBody {
	return b.s
}


