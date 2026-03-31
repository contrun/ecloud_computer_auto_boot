// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetMachineVNCQuery struct {
    position.Query
    // 桌面ID
	MachineId *string `json:"machineId,omitempty"`
}

func (s GetMachineVNCQuery) String() string {
	return utils.Beautify(s)
}

func (s GetMachineVNCQuery) GoString() string {
	return s.String()
}

func (s GetMachineVNCQuery) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetMachineVNCQuery) SetMachineId(v string) *GetMachineVNCQuery {
	s.MachineId = &v
	return s
}


type GetMachineVNCQueryBuilder struct {
	s *GetMachineVNCQuery
}

func NewGetMachineVNCQueryBuilder() *GetMachineVNCQueryBuilder {
	s := &GetMachineVNCQuery{}
	b := &GetMachineVNCQueryBuilder{s: s}
	return b
}

func (b *GetMachineVNCQueryBuilder) MachineId(v string) *GetMachineVNCQueryBuilder {
	b.s.MachineId = &v
	return b
}

func (b *GetMachineVNCQueryBuilder) Build() *GetMachineVNCQuery {
	return b.s
}


