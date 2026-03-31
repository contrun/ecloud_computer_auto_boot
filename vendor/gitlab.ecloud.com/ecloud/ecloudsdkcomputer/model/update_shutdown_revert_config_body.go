// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateShutdownRevertConfigBody struct {
    position.Body
    // 快照id,镜像id,还原至初始镜像传空
	Uid *string `json:"uid,omitempty"`
    // 云电脑id
	MachineId *string `json:"machineId,omitempty"`
    // 还原配置项 快照:snapshot 镜像:image 初始镜像: initialImage 不还原: reset
	RevertType *string `json:"revertType,omitempty"`
}

func (s UpdateShutdownRevertConfigBody) String() string {
	return utils.Beautify(s)
}

func (s UpdateShutdownRevertConfigBody) GoString() string {
	return s.String()
}

func (s UpdateShutdownRevertConfigBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateShutdownRevertConfigBody) SetUid(v string) *UpdateShutdownRevertConfigBody {
	s.Uid = &v
	return s
}

func (s *UpdateShutdownRevertConfigBody) SetMachineId(v string) *UpdateShutdownRevertConfigBody {
	s.MachineId = &v
	return s
}

func (s *UpdateShutdownRevertConfigBody) SetRevertType(v string) *UpdateShutdownRevertConfigBody {
	s.RevertType = &v
	return s
}


type UpdateShutdownRevertConfigBodyBuilder struct {
	s *UpdateShutdownRevertConfigBody
}

func NewUpdateShutdownRevertConfigBodyBuilder() *UpdateShutdownRevertConfigBodyBuilder {
	s := &UpdateShutdownRevertConfigBody{}
	b := &UpdateShutdownRevertConfigBodyBuilder{s: s}
	return b
}

func (b *UpdateShutdownRevertConfigBodyBuilder) Uid(v string) *UpdateShutdownRevertConfigBodyBuilder {
	b.s.Uid = &v
	return b
}

func (b *UpdateShutdownRevertConfigBodyBuilder) MachineId(v string) *UpdateShutdownRevertConfigBodyBuilder {
	b.s.MachineId = &v
	return b
}

func (b *UpdateShutdownRevertConfigBodyBuilder) RevertType(v string) *UpdateShutdownRevertConfigBodyBuilder {
	b.s.RevertType = &v
	return b
}

func (b *UpdateShutdownRevertConfigBodyBuilder) Build() *UpdateShutdownRevertConfigBody {
	return b.s
}


