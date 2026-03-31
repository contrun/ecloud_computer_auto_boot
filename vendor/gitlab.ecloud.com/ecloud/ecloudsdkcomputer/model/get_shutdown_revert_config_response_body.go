// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetShutdownRevertConfigResponseBody struct {

    // 关机还原配置修改时间
	ModifiedTime *string `json:"modifiedTime,omitempty"`
    // 镜像、快照id，默认的镜像传空
	Uid *string `json:"uid,omitempty"`
    // 桌面id
	MachineId *string `json:"machineId,omitempty"`
    // 删除标识，0表示未删除，1表示已删除
	IsDeleted *string `json:"isDeleted,omitempty"`
    // 镜像、快照名称
	Name *string `json:"name,omitempty"`
    // 关机还原配置创建时间
	CreatedTime *string `json:"createdTime,omitempty"`
    // 镜像、快照磁盘类型
	DiskType *string `json:"diskType,omitempty"`
    // 关机还原配置表主键id
	Id *int64 `json:"id,omitempty"`
    // 镜像、快照系统盘大小
	SysDiskSize *int32 `json:"sysDiskSize,omitempty"`
    // 还原类型，（指定快照、指定镜像、基础镜像）
	Type *string `json:"type,omitempty"`
}

func (s GetShutdownRevertConfigResponseBody) String() string {
	return utils.Beautify(s)
}

func (s GetShutdownRevertConfigResponseBody) GoString() string {
	return s.String()
}

func (s GetShutdownRevertConfigResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetShutdownRevertConfigResponseBody) SetModifiedTime(v string) *GetShutdownRevertConfigResponseBody {
	s.ModifiedTime = &v
	return s
}

func (s *GetShutdownRevertConfigResponseBody) SetUid(v string) *GetShutdownRevertConfigResponseBody {
	s.Uid = &v
	return s
}

func (s *GetShutdownRevertConfigResponseBody) SetMachineId(v string) *GetShutdownRevertConfigResponseBody {
	s.MachineId = &v
	return s
}

func (s *GetShutdownRevertConfigResponseBody) SetIsDeleted(v string) *GetShutdownRevertConfigResponseBody {
	s.IsDeleted = &v
	return s
}

func (s *GetShutdownRevertConfigResponseBody) SetName(v string) *GetShutdownRevertConfigResponseBody {
	s.Name = &v
	return s
}

func (s *GetShutdownRevertConfigResponseBody) SetCreatedTime(v string) *GetShutdownRevertConfigResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *GetShutdownRevertConfigResponseBody) SetDiskType(v string) *GetShutdownRevertConfigResponseBody {
	s.DiskType = &v
	return s
}

func (s *GetShutdownRevertConfigResponseBody) SetId(v int64) *GetShutdownRevertConfigResponseBody {
	s.Id = &v
	return s
}

func (s *GetShutdownRevertConfigResponseBody) SetSysDiskSize(v int32) *GetShutdownRevertConfigResponseBody {
	s.SysDiskSize = &v
	return s
}

func (s *GetShutdownRevertConfigResponseBody) SetType(v string) *GetShutdownRevertConfigResponseBody {
	s.Type = &v
	return s
}


type GetShutdownRevertConfigResponseBodyBuilder struct {
	s *GetShutdownRevertConfigResponseBody
}

func NewGetShutdownRevertConfigResponseBodyBuilder() *GetShutdownRevertConfigResponseBodyBuilder {
	s := &GetShutdownRevertConfigResponseBody{}
	b := &GetShutdownRevertConfigResponseBodyBuilder{s: s}
	return b
}

func (b *GetShutdownRevertConfigResponseBodyBuilder) ModifiedTime(v string) *GetShutdownRevertConfigResponseBodyBuilder {
	b.s.ModifiedTime = &v
	return b
}

func (b *GetShutdownRevertConfigResponseBodyBuilder) Uid(v string) *GetShutdownRevertConfigResponseBodyBuilder {
	b.s.Uid = &v
	return b
}

func (b *GetShutdownRevertConfigResponseBodyBuilder) MachineId(v string) *GetShutdownRevertConfigResponseBodyBuilder {
	b.s.MachineId = &v
	return b
}

func (b *GetShutdownRevertConfigResponseBodyBuilder) IsDeleted(v string) *GetShutdownRevertConfigResponseBodyBuilder {
	b.s.IsDeleted = &v
	return b
}

func (b *GetShutdownRevertConfigResponseBodyBuilder) Name(v string) *GetShutdownRevertConfigResponseBodyBuilder {
	b.s.Name = &v
	return b
}

func (b *GetShutdownRevertConfigResponseBodyBuilder) CreatedTime(v string) *GetShutdownRevertConfigResponseBodyBuilder {
	b.s.CreatedTime = &v
	return b
}

func (b *GetShutdownRevertConfigResponseBodyBuilder) DiskType(v string) *GetShutdownRevertConfigResponseBodyBuilder {
	b.s.DiskType = &v
	return b
}

func (b *GetShutdownRevertConfigResponseBodyBuilder) Id(v int64) *GetShutdownRevertConfigResponseBodyBuilder {
	b.s.Id = &v
	return b
}

func (b *GetShutdownRevertConfigResponseBodyBuilder) SysDiskSize(v int32) *GetShutdownRevertConfigResponseBodyBuilder {
	b.s.SysDiskSize = &v
	return b
}

func (b *GetShutdownRevertConfigResponseBodyBuilder) Type(v string) *GetShutdownRevertConfigResponseBodyBuilder {
	b.s.Type = &v
	return b
}

func (b *GetShutdownRevertConfigResponseBodyBuilder) Build() *GetShutdownRevertConfigResponseBody {
	return b.s
}


