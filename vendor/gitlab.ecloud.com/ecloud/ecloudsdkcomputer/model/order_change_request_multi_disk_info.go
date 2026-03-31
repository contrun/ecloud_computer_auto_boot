// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type OrderChangeRequestMultiDiskInfo struct {

    // 盘符
	DiskName *string `json:"diskName,omitempty"`
    // 数据盘id
	DiskUid *string `json:"diskUid,omitempty"`
    // 数据盘大小
	DiskSize *int32 `json:"diskSize,omitempty"`
    // 新增数据盘名称
	DiskDescription *string `json:"diskDescription,omitempty"`
}

func (s OrderChangeRequestMultiDiskInfo) String() string {
	return utils.Beautify(s)
}

func (s OrderChangeRequestMultiDiskInfo) GoString() string {
	return s.String()
}

func (s OrderChangeRequestMultiDiskInfo) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *OrderChangeRequestMultiDiskInfo) SetDiskName(v string) *OrderChangeRequestMultiDiskInfo {
	s.DiskName = &v
	return s
}

func (s *OrderChangeRequestMultiDiskInfo) SetDiskUid(v string) *OrderChangeRequestMultiDiskInfo {
	s.DiskUid = &v
	return s
}

func (s *OrderChangeRequestMultiDiskInfo) SetDiskSize(v int32) *OrderChangeRequestMultiDiskInfo {
	s.DiskSize = &v
	return s
}

func (s *OrderChangeRequestMultiDiskInfo) SetDiskDescription(v string) *OrderChangeRequestMultiDiskInfo {
	s.DiskDescription = &v
	return s
}


type OrderChangeRequestMultiDiskInfoBuilder struct {
	s *OrderChangeRequestMultiDiskInfo
}

func NewOrderChangeRequestMultiDiskInfoBuilder() *OrderChangeRequestMultiDiskInfoBuilder {
	s := &OrderChangeRequestMultiDiskInfo{}
	b := &OrderChangeRequestMultiDiskInfoBuilder{s: s}
	return b
}

func (b *OrderChangeRequestMultiDiskInfoBuilder) DiskName(v string) *OrderChangeRequestMultiDiskInfoBuilder {
	b.s.DiskName = &v
	return b
}

func (b *OrderChangeRequestMultiDiskInfoBuilder) DiskUid(v string) *OrderChangeRequestMultiDiskInfoBuilder {
	b.s.DiskUid = &v
	return b
}

func (b *OrderChangeRequestMultiDiskInfoBuilder) DiskSize(v int32) *OrderChangeRequestMultiDiskInfoBuilder {
	b.s.DiskSize = &v
	return b
}

func (b *OrderChangeRequestMultiDiskInfoBuilder) DiskDescription(v string) *OrderChangeRequestMultiDiskInfoBuilder {
	b.s.DiskDescription = &v
	return b
}

func (b *OrderChangeRequestMultiDiskInfoBuilder) Build() *OrderChangeRequestMultiDiskInfo {
	return b.s
}


