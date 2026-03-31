// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetMultiDiskInfoResponseData struct {

    // 资源实例id
	InstanceId *string `json:"instanceId,omitempty"`
    // 多数据盘信息
	DiskInfoList *[]GetMultiDiskInfoResponseDiskInfoList `json:"diskInfoList,omitempty"`
}

func (s GetMultiDiskInfoResponseData) String() string {
	return utils.Beautify(s)
}

func (s GetMultiDiskInfoResponseData) GoString() string {
	return s.String()
}

func (s GetMultiDiskInfoResponseData) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetMultiDiskInfoResponseData) SetInstanceId(v string) *GetMultiDiskInfoResponseData {
	s.InstanceId = &v
	return s
}

func (s *GetMultiDiskInfoResponseData) SetDiskInfoList(v []GetMultiDiskInfoResponseDiskInfoList) *GetMultiDiskInfoResponseData {
	s.DiskInfoList = &v
	return s
}


type GetMultiDiskInfoResponseDataBuilder struct {
	s *GetMultiDiskInfoResponseData
}

func NewGetMultiDiskInfoResponseDataBuilder() *GetMultiDiskInfoResponseDataBuilder {
	s := &GetMultiDiskInfoResponseData{}
	b := &GetMultiDiskInfoResponseDataBuilder{s: s}
	return b
}

func (b *GetMultiDiskInfoResponseDataBuilder) InstanceId(v string) *GetMultiDiskInfoResponseDataBuilder {
	b.s.InstanceId = &v
	return b
}

func (b *GetMultiDiskInfoResponseDataBuilder) DiskInfoList(v []GetMultiDiskInfoResponseDiskInfoList) *GetMultiDiskInfoResponseDataBuilder {
	b.s.DiskInfoList = &v
	return b
}

func (b *GetMultiDiskInfoResponseDataBuilder) Build() *GetMultiDiskInfoResponseData {
	return b.s
}


