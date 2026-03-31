// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type OrderChangeBody struct {
    position.Body
    // 计费模式 1:按周期，2:按量
	ChargingMode *int32 `json:"chargingMode,omitempty"`
    // 额外数据盘大小
	DataDiskExt *string `json:"dataDiskExt,omitempty"`
    // 数据盘扩容类型 1:普通，2:多数据盘
	DiskChangeType *int32 `json:"diskChangeType,omitempty"`
    // 资源实例id，示例值：CCA-XXX
	InstanceId *string `json:"instanceId,omitempty"`
    // 多数据盘信息
	MultiDiskInfo *OrderChangeRequestMultiDiskInfo `json:"multiDiskInfo,omitempty"`
    // 变更类型 1：规格，2：数据盘
	ChangeType *int32 `json:"changeType,omitempty"`
    // 变更后的规格编码，1：办公型-公众版4vCPU/8GB内存；办公型-政企版4vCPU/8GB内存；办公型-政企版（边缘节点）4vCPU/8GB内存；信创型-政企版4vCPU/8GB内存；信创型-政企版（边缘节点）4vCPU/8GB内存。2:办公型-公众版8vCPU/16GB内存；办公型-政企版8vCPU/16GB内存；办公型-政企版（边缘节点）8vCPU/16GB内存；信创型-政企版8vCPU/16GB内存；信创型-政企版（边缘节点）8vCPU/16GB内存。4:办公型-公众版4vCPU/4GB内存；办公型-政企版4vCPU/4GB内存；办公型-政企版（边缘节点）4vCPU/4GB内存；信创型-政企版4vCPU/4GB内存；信创型-政企版（边缘节点）4vCPU/4GB内存。5:办公型-公众版8vCPU/8GB内存；办公型-政企版8vCPU/8GB内存；办公型-政企版（边缘节点）8vCPU/8GB内存；信创型-政企版8vCPU/8GB内存；信创型-政企版（边缘节点）8vCPU/8GB内存。6:办公型-公众版16vCPU/16GB内存；办公型-政企版16vCPU/16GB内存；办公型-政企版（边缘节点）16vCPU/16GB内存；信创型-政企版16vCPU/16GB内存；信创型-政企版（边缘节点）16vCPU/16GB内存。7:办公型-公众版16vCPU/32GB内存；办公型-政企版16vCPU/32GB内存；办公型-政企版（边缘节点）16vCPU/32GB内存；信创型-政企版16vCPU/32GB内存；信创型-政企版（边缘节点）16vCPU/32GB内存。
	SpecificationCode *string `json:"specificationCode,omitempty"`
}

func (s OrderChangeBody) String() string {
	return utils.Beautify(s)
}

func (s OrderChangeBody) GoString() string {
	return s.String()
}

func (s OrderChangeBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *OrderChangeBody) SetChargingMode(v int32) *OrderChangeBody {
	s.ChargingMode = &v
	return s
}

func (s *OrderChangeBody) SetDataDiskExt(v string) *OrderChangeBody {
	s.DataDiskExt = &v
	return s
}

func (s *OrderChangeBody) SetDiskChangeType(v int32) *OrderChangeBody {
	s.DiskChangeType = &v
	return s
}

func (s *OrderChangeBody) SetInstanceId(v string) *OrderChangeBody {
	s.InstanceId = &v
	return s
}

func (s *OrderChangeBody) SetMultiDiskInfo(v *OrderChangeRequestMultiDiskInfo) *OrderChangeBody {
	s.MultiDiskInfo = v
	return s
}

func (s *OrderChangeBody) SetChangeType(v int32) *OrderChangeBody {
	s.ChangeType = &v
	return s
}

func (s *OrderChangeBody) SetSpecificationCode(v string) *OrderChangeBody {
	s.SpecificationCode = &v
	return s
}


type OrderChangeBodyBuilder struct {
	s *OrderChangeBody
}

func NewOrderChangeBodyBuilder() *OrderChangeBodyBuilder {
	s := &OrderChangeBody{}
	b := &OrderChangeBodyBuilder{s: s}
	return b
}

func (b *OrderChangeBodyBuilder) ChargingMode(v int32) *OrderChangeBodyBuilder {
	b.s.ChargingMode = &v
	return b
}

func (b *OrderChangeBodyBuilder) DataDiskExt(v string) *OrderChangeBodyBuilder {
	b.s.DataDiskExt = &v
	return b
}

func (b *OrderChangeBodyBuilder) DiskChangeType(v int32) *OrderChangeBodyBuilder {
	b.s.DiskChangeType = &v
	return b
}

func (b *OrderChangeBodyBuilder) InstanceId(v string) *OrderChangeBodyBuilder {
	b.s.InstanceId = &v
	return b
}

func (b *OrderChangeBodyBuilder) MultiDiskInfo(v *OrderChangeRequestMultiDiskInfo) *OrderChangeBodyBuilder {
	b.s.MultiDiskInfo = v
	return b
}

func (b *OrderChangeBodyBuilder) ChangeType(v int32) *OrderChangeBodyBuilder {
	b.s.ChangeType = &v
	return b
}

func (b *OrderChangeBodyBuilder) SpecificationCode(v string) *OrderChangeBodyBuilder {
	b.s.SpecificationCode = &v
	return b
}

func (b *OrderChangeBodyBuilder) Build() *OrderChangeBody {
	return b.s
}


