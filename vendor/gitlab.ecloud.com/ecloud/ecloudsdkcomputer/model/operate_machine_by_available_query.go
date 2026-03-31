// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type OperateMachineByAvailableQuery struct {
    position.Query
    // 云电脑id
	MachineId *string `json:"machineId,omitempty"`
}

func (s OperateMachineByAvailableQuery) String() string {
	return utils.Beautify(s)
}

func (s OperateMachineByAvailableQuery) GoString() string {
	return s.String()
}

func (s OperateMachineByAvailableQuery) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *OperateMachineByAvailableQuery) SetMachineId(v string) *OperateMachineByAvailableQuery {
	s.MachineId = &v
	return s
}


type OperateMachineByAvailableQueryBuilder struct {
	s *OperateMachineByAvailableQuery
}

func NewOperateMachineByAvailableQueryBuilder() *OperateMachineByAvailableQueryBuilder {
	s := &OperateMachineByAvailableQuery{}
	b := &OperateMachineByAvailableQueryBuilder{s: s}
	return b
}

func (b *OperateMachineByAvailableQueryBuilder) MachineId(v string) *OperateMachineByAvailableQueryBuilder {
	b.s.MachineId = &v
	return b
}

func (b *OperateMachineByAvailableQueryBuilder) Build() *OperateMachineByAvailableQuery {
	return b.s
}


