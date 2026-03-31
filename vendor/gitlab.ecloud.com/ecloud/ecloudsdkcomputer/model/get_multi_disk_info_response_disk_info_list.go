// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetMultiDiskInfoResponseDiskInfoList struct {

    // 盘符
	DiskName *string `json:"diskName,omitempty"`
    // 数据盘id
	DiskUid *string `json:"diskUid,omitempty"`
    // 数据盘大小
	DiskSize *int32 `json:"diskSize,omitempty"`
    // 新增数据盘名称
	DiskDescription *string `json:"diskDescription,omitempty"`
}

func (s GetMultiDiskInfoResponseDiskInfoList) String() string {
	return utils.Beautify(s)
}

func (s GetMultiDiskInfoResponseDiskInfoList) GoString() string {
	return s.String()
}

func (s GetMultiDiskInfoResponseDiskInfoList) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetMultiDiskInfoResponseDiskInfoList) SetDiskName(v string) *GetMultiDiskInfoResponseDiskInfoList {
	s.DiskName = &v
	return s
}

func (s *GetMultiDiskInfoResponseDiskInfoList) SetDiskUid(v string) *GetMultiDiskInfoResponseDiskInfoList {
	s.DiskUid = &v
	return s
}

func (s *GetMultiDiskInfoResponseDiskInfoList) SetDiskSize(v int32) *GetMultiDiskInfoResponseDiskInfoList {
	s.DiskSize = &v
	return s
}

func (s *GetMultiDiskInfoResponseDiskInfoList) SetDiskDescription(v string) *GetMultiDiskInfoResponseDiskInfoList {
	s.DiskDescription = &v
	return s
}


type GetMultiDiskInfoResponseDiskInfoListBuilder struct {
	s *GetMultiDiskInfoResponseDiskInfoList
}

func NewGetMultiDiskInfoResponseDiskInfoListBuilder() *GetMultiDiskInfoResponseDiskInfoListBuilder {
	s := &GetMultiDiskInfoResponseDiskInfoList{}
	b := &GetMultiDiskInfoResponseDiskInfoListBuilder{s: s}
	return b
}

func (b *GetMultiDiskInfoResponseDiskInfoListBuilder) DiskName(v string) *GetMultiDiskInfoResponseDiskInfoListBuilder {
	b.s.DiskName = &v
	return b
}

func (b *GetMultiDiskInfoResponseDiskInfoListBuilder) DiskUid(v string) *GetMultiDiskInfoResponseDiskInfoListBuilder {
	b.s.DiskUid = &v
	return b
}

func (b *GetMultiDiskInfoResponseDiskInfoListBuilder) DiskSize(v int32) *GetMultiDiskInfoResponseDiskInfoListBuilder {
	b.s.DiskSize = &v
	return b
}

func (b *GetMultiDiskInfoResponseDiskInfoListBuilder) DiskDescription(v string) *GetMultiDiskInfoResponseDiskInfoListBuilder {
	b.s.DiskDescription = &v
	return b
}

func (b *GetMultiDiskInfoResponseDiskInfoListBuilder) Build() *GetMultiDiskInfoResponseDiskInfoList {
	return b.s
}


