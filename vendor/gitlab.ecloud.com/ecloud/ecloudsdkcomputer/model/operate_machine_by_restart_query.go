// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type OperateMachineByRestartQuery struct {
    position.Query
    // 云电脑id
	MachineId *string `json:"machineId,omitempty"`
}

func (s OperateMachineByRestartQuery) String() string {
	return utils.Beautify(s)
}

func (s OperateMachineByRestartQuery) GoString() string {
	return s.String()
}

func (s OperateMachineByRestartQuery) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *OperateMachineByRestartQuery) SetMachineId(v string) *OperateMachineByRestartQuery {
	s.MachineId = &v
	return s
}


type OperateMachineByRestartQueryBuilder struct {
	s *OperateMachineByRestartQuery
}

func NewOperateMachineByRestartQueryBuilder() *OperateMachineByRestartQueryBuilder {
	s := &OperateMachineByRestartQuery{}
	b := &OperateMachineByRestartQueryBuilder{s: s}
	return b
}

func (b *OperateMachineByRestartQueryBuilder) MachineId(v string) *OperateMachineByRestartQueryBuilder {
	b.s.MachineId = &v
	return b
}

func (b *OperateMachineByRestartQueryBuilder) Build() *OperateMachineByRestartQuery {
	return b.s
}


