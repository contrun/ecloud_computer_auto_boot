// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type OperateMachineByShutdownQuery struct {
    position.Query
    // 云电脑id
	MachineId *string `json:"machineId,omitempty"`
}

func (s OperateMachineByShutdownQuery) String() string {
	return utils.Beautify(s)
}

func (s OperateMachineByShutdownQuery) GoString() string {
	return s.String()
}

func (s OperateMachineByShutdownQuery) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *OperateMachineByShutdownQuery) SetMachineId(v string) *OperateMachineByShutdownQuery {
	s.MachineId = &v
	return s
}


type OperateMachineByShutdownQueryBuilder struct {
	s *OperateMachineByShutdownQuery
}

func NewOperateMachineByShutdownQueryBuilder() *OperateMachineByShutdownQueryBuilder {
	s := &OperateMachineByShutdownQuery{}
	b := &OperateMachineByShutdownQueryBuilder{s: s}
	return b
}

func (b *OperateMachineByShutdownQueryBuilder) MachineId(v string) *OperateMachineByShutdownQueryBuilder {
	b.s.MachineId = &v
	return b
}

func (b *OperateMachineByShutdownQueryBuilder) Build() *OperateMachineByShutdownQuery {
	return b.s
}


