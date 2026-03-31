// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetVmListResponseMachineList struct {

    // 桌面状态 available 运行中 shutdown 已关机 onShutdown 关机中 onAvailable 开机中 onReload 重装中 onRestart 重启中 exception 异常 imageCreating 镜像制作中 snapshotCreating 快照制作中 transferring 迁移中 freeze 已分配-冻结 alter 已分配-变更中
	ResourceStatus *string `json:"resourceStatus,omitempty"`
    // 桌面状态描述 运行中 已关机 关机中 开机中 重装中 重启中 异常 镜像制作中 快照制作中 迁移中 已分配-冻结 已分配-变更中
	ResourceStatusCn *string `json:"resourceStatusCn,omitempty"`
    // 桌面ID
	MachineId *string `json:"machineId,omitempty"`
    // 连接信息
	ConnectInfo *string `json:"connectInfo,omitempty"`
    // 桌面名称
	MachineName *string `json:"machineName,omitempty"`
}

func (s GetVmListResponseMachineList) String() string {
	return utils.Beautify(s)
}

func (s GetVmListResponseMachineList) GoString() string {
	return s.String()
}

func (s GetVmListResponseMachineList) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetVmListResponseMachineList) SetResourceStatus(v string) *GetVmListResponseMachineList {
	s.ResourceStatus = &v
	return s
}

func (s *GetVmListResponseMachineList) SetResourceStatusCn(v string) *GetVmListResponseMachineList {
	s.ResourceStatusCn = &v
	return s
}

func (s *GetVmListResponseMachineList) SetMachineId(v string) *GetVmListResponseMachineList {
	s.MachineId = &v
	return s
}

func (s *GetVmListResponseMachineList) SetConnectInfo(v string) *GetVmListResponseMachineList {
	s.ConnectInfo = &v
	return s
}

func (s *GetVmListResponseMachineList) SetMachineName(v string) *GetVmListResponseMachineList {
	s.MachineName = &v
	return s
}


type GetVmListResponseMachineListBuilder struct {
	s *GetVmListResponseMachineList
}

func NewGetVmListResponseMachineListBuilder() *GetVmListResponseMachineListBuilder {
	s := &GetVmListResponseMachineList{}
	b := &GetVmListResponseMachineListBuilder{s: s}
	return b
}

func (b *GetVmListResponseMachineListBuilder) ResourceStatus(v string) *GetVmListResponseMachineListBuilder {
	b.s.ResourceStatus = &v
	return b
}

func (b *GetVmListResponseMachineListBuilder) ResourceStatusCn(v string) *GetVmListResponseMachineListBuilder {
	b.s.ResourceStatusCn = &v
	return b
}

func (b *GetVmListResponseMachineListBuilder) MachineId(v string) *GetVmListResponseMachineListBuilder {
	b.s.MachineId = &v
	return b
}

func (b *GetVmListResponseMachineListBuilder) ConnectInfo(v string) *GetVmListResponseMachineListBuilder {
	b.s.ConnectInfo = &v
	return b
}

func (b *GetVmListResponseMachineListBuilder) MachineName(v string) *GetVmListResponseMachineListBuilder {
	b.s.MachineName = &v
	return b
}

func (b *GetVmListResponseMachineListBuilder) Build() *GetVmListResponseMachineList {
	return b.s
}


