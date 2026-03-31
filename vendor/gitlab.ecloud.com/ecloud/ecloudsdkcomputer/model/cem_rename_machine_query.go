// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CemRenameMachineQuery struct {
    position.Query
    // 云电脑id
	MachineId *string `json:"machineId,omitempty"`
    // 云电脑名称
	MachineName *string `json:"machineName,omitempty"`
}

func (s CemRenameMachineQuery) String() string {
	return utils.Beautify(s)
}

func (s CemRenameMachineQuery) GoString() string {
	return s.String()
}

func (s CemRenameMachineQuery) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CemRenameMachineQuery) SetMachineId(v string) *CemRenameMachineQuery {
	s.MachineId = &v
	return s
}

func (s *CemRenameMachineQuery) SetMachineName(v string) *CemRenameMachineQuery {
	s.MachineName = &v
	return s
}


type CemRenameMachineQueryBuilder struct {
	s *CemRenameMachineQuery
}

func NewCemRenameMachineQueryBuilder() *CemRenameMachineQueryBuilder {
	s := &CemRenameMachineQuery{}
	b := &CemRenameMachineQueryBuilder{s: s}
	return b
}

func (b *CemRenameMachineQueryBuilder) MachineId(v string) *CemRenameMachineQueryBuilder {
	b.s.MachineId = &v
	return b
}

func (b *CemRenameMachineQueryBuilder) MachineName(v string) *CemRenameMachineQueryBuilder {
	b.s.MachineName = &v
	return b
}

func (b *CemRenameMachineQueryBuilder) Build() *CemRenameMachineQuery {
	return b.s
}


