// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type MultiDiskChangeBody struct {
    position.Body
    // 新增数据盘大小
	DiskSize *int32 `json:"diskSize,omitempty"`
    // 资源实例id，示例值：CCA-XXX
	InstanceId *string `json:"instanceId,omitempty"`
    // 新增数据盘名称
	DiskDescription *string `json:"diskDescription,omitempty"`
}

func (s MultiDiskChangeBody) String() string {
	return utils.Beautify(s)
}

func (s MultiDiskChangeBody) GoString() string {
	return s.String()
}

func (s MultiDiskChangeBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *MultiDiskChangeBody) SetDiskSize(v int32) *MultiDiskChangeBody {
	s.DiskSize = &v
	return s
}

func (s *MultiDiskChangeBody) SetInstanceId(v string) *MultiDiskChangeBody {
	s.InstanceId = &v
	return s
}

func (s *MultiDiskChangeBody) SetDiskDescription(v string) *MultiDiskChangeBody {
	s.DiskDescription = &v
	return s
}


type MultiDiskChangeBodyBuilder struct {
	s *MultiDiskChangeBody
}

func NewMultiDiskChangeBodyBuilder() *MultiDiskChangeBodyBuilder {
	s := &MultiDiskChangeBody{}
	b := &MultiDiskChangeBodyBuilder{s: s}
	return b
}

func (b *MultiDiskChangeBodyBuilder) DiskSize(v int32) *MultiDiskChangeBodyBuilder {
	b.s.DiskSize = &v
	return b
}

func (b *MultiDiskChangeBodyBuilder) InstanceId(v string) *MultiDiskChangeBodyBuilder {
	b.s.InstanceId = &v
	return b
}

func (b *MultiDiskChangeBodyBuilder) DiskDescription(v string) *MultiDiskChangeBodyBuilder {
	b.s.DiskDescription = &v
	return b
}

func (b *MultiDiskChangeBodyBuilder) Build() *MultiDiskChangeBody {
	return b.s
}


