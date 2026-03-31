// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ChangeVpcBody struct {
    position.Body
    // 子网ID，/getNetworkSelect接口返参id处获取
	SubnetId *string `json:"subnetId,omitempty"`
    // 子网所属vpcUid
	VpcUid *string `json:"vpcUid,omitempty"`
    // 子网uID
	NetworkUid *string `json:"networkUid,omitempty"`

	OldSubnetId *string `json:"oldSubnetId,omitempty"`

	OldVpcUid *string `json:"oldVpcUid,omitempty"`
    // machineIdList
	MachineIdList []string `json:"machineIdList,omitempty"`
}

func (s ChangeVpcBody) String() string {
	return utils.Beautify(s)
}

func (s ChangeVpcBody) GoString() string {
	return s.String()
}

func (s ChangeVpcBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ChangeVpcBody) SetSubnetId(v string) *ChangeVpcBody {
	s.SubnetId = &v
	return s
}

func (s *ChangeVpcBody) SetVpcUid(v string) *ChangeVpcBody {
	s.VpcUid = &v
	return s
}

func (s *ChangeVpcBody) SetNetworkUid(v string) *ChangeVpcBody {
	s.NetworkUid = &v
	return s
}

func (s *ChangeVpcBody) SetOldSubnetId(v string) *ChangeVpcBody {
	s.OldSubnetId = &v
	return s
}

func (s *ChangeVpcBody) SetOldVpcUid(v string) *ChangeVpcBody {
	s.OldVpcUid = &v
	return s
}

func (s *ChangeVpcBody) SetMachineIdList(v []string) *ChangeVpcBody {
	s.MachineIdList = v
	return s
}


type ChangeVpcBodyBuilder struct {
	s *ChangeVpcBody
}

func NewChangeVpcBodyBuilder() *ChangeVpcBodyBuilder {
	s := &ChangeVpcBody{}
	b := &ChangeVpcBodyBuilder{s: s}
	return b
}

func (b *ChangeVpcBodyBuilder) SubnetId(v string) *ChangeVpcBodyBuilder {
	b.s.SubnetId = &v
	return b
}

func (b *ChangeVpcBodyBuilder) VpcUid(v string) *ChangeVpcBodyBuilder {
	b.s.VpcUid = &v
	return b
}

func (b *ChangeVpcBodyBuilder) NetworkUid(v string) *ChangeVpcBodyBuilder {
	b.s.NetworkUid = &v
	return b
}

func (b *ChangeVpcBodyBuilder) OldSubnetId(v string) *ChangeVpcBodyBuilder {
	b.s.OldSubnetId = &v
	return b
}

func (b *ChangeVpcBodyBuilder) OldVpcUid(v string) *ChangeVpcBodyBuilder {
	b.s.OldVpcUid = &v
	return b
}

func (b *ChangeVpcBodyBuilder) MachineIdList(v []string) *ChangeVpcBodyBuilder {
	b.s.MachineIdList = v
	return b
}

func (b *ChangeVpcBodyBuilder) Build() *ChangeVpcBody {
	return b.s
}


