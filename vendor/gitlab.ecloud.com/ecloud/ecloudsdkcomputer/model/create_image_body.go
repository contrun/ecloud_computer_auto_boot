// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreateImageBody struct {
    position.Body
    // 镜像名称
	ImageName *string `json:"imageName,omitempty"`
    // 桌面id
	MachineId *string `json:"machineId,omitempty"`
    // 镜像模板描述
	ImageRemark *string `json:"imageRemark,omitempty"`
    // 磁盘性质（当前仅支持系统盘system）
	DiskType *string `json:"diskType,omitempty"`
    // 镜像类型（当前仅支持自定义custom）
	ImageType *string `json:"imageType,omitempty"`
    // 镜像模板范围（当前仅支持系统盘system）
	ImageTarget *string `json:"imageTarget,omitempty"`
}

func (s CreateImageBody) String() string {
	return utils.Beautify(s)
}

func (s CreateImageBody) GoString() string {
	return s.String()
}

func (s CreateImageBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateImageBody) SetImageName(v string) *CreateImageBody {
	s.ImageName = &v
	return s
}

func (s *CreateImageBody) SetMachineId(v string) *CreateImageBody {
	s.MachineId = &v
	return s
}

func (s *CreateImageBody) SetImageRemark(v string) *CreateImageBody {
	s.ImageRemark = &v
	return s
}

func (s *CreateImageBody) SetDiskType(v string) *CreateImageBody {
	s.DiskType = &v
	return s
}

func (s *CreateImageBody) SetImageType(v string) *CreateImageBody {
	s.ImageType = &v
	return s
}

func (s *CreateImageBody) SetImageTarget(v string) *CreateImageBody {
	s.ImageTarget = &v
	return s
}


type CreateImageBodyBuilder struct {
	s *CreateImageBody
}

func NewCreateImageBodyBuilder() *CreateImageBodyBuilder {
	s := &CreateImageBody{}
	b := &CreateImageBodyBuilder{s: s}
	return b
}

func (b *CreateImageBodyBuilder) ImageName(v string) *CreateImageBodyBuilder {
	b.s.ImageName = &v
	return b
}

func (b *CreateImageBodyBuilder) MachineId(v string) *CreateImageBodyBuilder {
	b.s.MachineId = &v
	return b
}

func (b *CreateImageBodyBuilder) ImageRemark(v string) *CreateImageBodyBuilder {
	b.s.ImageRemark = &v
	return b
}

func (b *CreateImageBodyBuilder) DiskType(v string) *CreateImageBodyBuilder {
	b.s.DiskType = &v
	return b
}

func (b *CreateImageBodyBuilder) ImageType(v string) *CreateImageBodyBuilder {
	b.s.ImageType = &v
	return b
}

func (b *CreateImageBodyBuilder) ImageTarget(v string) *CreateImageBodyBuilder {
	b.s.ImageTarget = &v
	return b
}

func (b *CreateImageBodyBuilder) Build() *CreateImageBody {
	return b.s
}


