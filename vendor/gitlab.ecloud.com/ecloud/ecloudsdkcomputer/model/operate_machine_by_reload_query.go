// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type OperateMachineByReloadQuery struct {
    position.Query
    // 云电脑id
	MachineId *string `json:"machineId,omitempty"`
}

func (s OperateMachineByReloadQuery) String() string {
	return utils.Beautify(s)
}

func (s OperateMachineByReloadQuery) GoString() string {
	return s.String()
}

func (s OperateMachineByReloadQuery) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *OperateMachineByReloadQuery) SetMachineId(v string) *OperateMachineByReloadQuery {
	s.MachineId = &v
	return s
}


type OperateMachineByReloadQueryBuilder struct {
	s *OperateMachineByReloadQuery
}

func NewOperateMachineByReloadQueryBuilder() *OperateMachineByReloadQueryBuilder {
	s := &OperateMachineByReloadQuery{}
	b := &OperateMachineByReloadQueryBuilder{s: s}
	return b
}

func (b *OperateMachineByReloadQueryBuilder) MachineId(v string) *OperateMachineByReloadQueryBuilder {
	b.s.MachineId = &v
	return b
}

func (b *OperateMachineByReloadQueryBuilder) Build() *OperateMachineByReloadQuery {
	return b.s
}


