// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetVpcSelectBody struct {
    position.Body
    // 资源池
	PoolCode *string `json:"poolCode,omitempty"`
    // 可用区
	Region *string `json:"region,omitempty"`
    // machineIdList
	MachineIdList []string `json:"machineIdList,omitempty"`
}

func (s GetVpcSelectBody) String() string {
	return utils.Beautify(s)
}

func (s GetVpcSelectBody) GoString() string {
	return s.String()
}

func (s GetVpcSelectBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetVpcSelectBody) SetPoolCode(v string) *GetVpcSelectBody {
	s.PoolCode = &v
	return s
}

func (s *GetVpcSelectBody) SetRegion(v string) *GetVpcSelectBody {
	s.Region = &v
	return s
}

func (s *GetVpcSelectBody) SetMachineIdList(v []string) *GetVpcSelectBody {
	s.MachineIdList = v
	return s
}


type GetVpcSelectBodyBuilder struct {
	s *GetVpcSelectBody
}

func NewGetVpcSelectBodyBuilder() *GetVpcSelectBodyBuilder {
	s := &GetVpcSelectBody{}
	b := &GetVpcSelectBodyBuilder{s: s}
	return b
}

func (b *GetVpcSelectBodyBuilder) PoolCode(v string) *GetVpcSelectBodyBuilder {
	b.s.PoolCode = &v
	return b
}

func (b *GetVpcSelectBodyBuilder) Region(v string) *GetVpcSelectBodyBuilder {
	b.s.Region = &v
	return b
}

func (b *GetVpcSelectBodyBuilder) MachineIdList(v []string) *GetVpcSelectBodyBuilder {
	b.s.MachineIdList = v
	return b
}

func (b *GetVpcSelectBodyBuilder) Build() *GetVpcSelectBody {
	return b.s
}


