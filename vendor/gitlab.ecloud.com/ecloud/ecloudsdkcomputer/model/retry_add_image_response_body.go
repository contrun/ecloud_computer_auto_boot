// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type RetryAddImageResponseBody struct {

    // 厂商编码
	CompanyCode *string `json:"companyCode,omitempty"`
    // 修改时间
	ModifiedTime *string `json:"modifiedTime,omitempty"`
    // 资源表主键
	ResourceId *int64 `json:"resourceId,omitempty"`
    // 快照uid
	SnapshotId *string `json:"snapshotId,omitempty"`
    // 错误信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 镜像模板uid
	TemplateId *string `json:"templateId,omitempty"`
    // 主机名/桌面名称/虚机名称
	MachineName *string `json:"machineName,omitempty"`
    // 主机ID/云电脑ID/虚机ID
	MachineId *string `json:"machineId,omitempty"`
    // 删除标识
	IsDeleted *string `json:"isDeleted,omitempty"`
    // 与底层交互的请求id
	RequestId *string `json:"requestId,omitempty"`
    // 0:没有告警；1:已经告警
	Alerted *string `json:"alerted,omitempty"`
    // 创建时间
	CreatedTime *string `json:"createdTime,omitempty"`
    // 资源池编码
	PoolCode *string `json:"poolCode,omitempty"`
    // 主键id
	Id *int64 `json:"id,omitempty"`
    // 记录的业务类型：镜像：1，快照：2
	BusinessType *string `json:"businessType,omitempty"`
    // 可用区
	Region *string `json:"region,omitempty"`
    // 状态,0:快照恢复中（镜像重装中）,1:处理成功,2:处理失败
	Status *string `json:"status,omitempty"`
}

func (s RetryAddImageResponseBody) String() string {
	return utils.Beautify(s)
}

func (s RetryAddImageResponseBody) GoString() string {
	return s.String()
}

func (s RetryAddImageResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *RetryAddImageResponseBody) SetCompanyCode(v string) *RetryAddImageResponseBody {
	s.CompanyCode = &v
	return s
}

func (s *RetryAddImageResponseBody) SetModifiedTime(v string) *RetryAddImageResponseBody {
	s.ModifiedTime = &v
	return s
}

func (s *RetryAddImageResponseBody) SetResourceId(v int64) *RetryAddImageResponseBody {
	s.ResourceId = &v
	return s
}

func (s *RetryAddImageResponseBody) SetSnapshotId(v string) *RetryAddImageResponseBody {
	s.SnapshotId = &v
	return s
}

func (s *RetryAddImageResponseBody) SetErrorMessage(v string) *RetryAddImageResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *RetryAddImageResponseBody) SetTemplateId(v string) *RetryAddImageResponseBody {
	s.TemplateId = &v
	return s
}

func (s *RetryAddImageResponseBody) SetMachineName(v string) *RetryAddImageResponseBody {
	s.MachineName = &v
	return s
}

func (s *RetryAddImageResponseBody) SetMachineId(v string) *RetryAddImageResponseBody {
	s.MachineId = &v
	return s
}

func (s *RetryAddImageResponseBody) SetIsDeleted(v string) *RetryAddImageResponseBody {
	s.IsDeleted = &v
	return s
}

func (s *RetryAddImageResponseBody) SetRequestId(v string) *RetryAddImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *RetryAddImageResponseBody) SetAlerted(v string) *RetryAddImageResponseBody {
	s.Alerted = &v
	return s
}

func (s *RetryAddImageResponseBody) SetCreatedTime(v string) *RetryAddImageResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *RetryAddImageResponseBody) SetPoolCode(v string) *RetryAddImageResponseBody {
	s.PoolCode = &v
	return s
}

func (s *RetryAddImageResponseBody) SetId(v int64) *RetryAddImageResponseBody {
	s.Id = &v
	return s
}

func (s *RetryAddImageResponseBody) SetBusinessType(v string) *RetryAddImageResponseBody {
	s.BusinessType = &v
	return s
}

func (s *RetryAddImageResponseBody) SetRegion(v string) *RetryAddImageResponseBody {
	s.Region = &v
	return s
}

func (s *RetryAddImageResponseBody) SetStatus(v string) *RetryAddImageResponseBody {
	s.Status = &v
	return s
}


type RetryAddImageResponseBodyBuilder struct {
	s *RetryAddImageResponseBody
}

func NewRetryAddImageResponseBodyBuilder() *RetryAddImageResponseBodyBuilder {
	s := &RetryAddImageResponseBody{}
	b := &RetryAddImageResponseBodyBuilder{s: s}
	return b
}

func (b *RetryAddImageResponseBodyBuilder) CompanyCode(v string) *RetryAddImageResponseBodyBuilder {
	b.s.CompanyCode = &v
	return b
}

func (b *RetryAddImageResponseBodyBuilder) ModifiedTime(v string) *RetryAddImageResponseBodyBuilder {
	b.s.ModifiedTime = &v
	return b
}

func (b *RetryAddImageResponseBodyBuilder) ResourceId(v int64) *RetryAddImageResponseBodyBuilder {
	b.s.ResourceId = &v
	return b
}

func (b *RetryAddImageResponseBodyBuilder) SnapshotId(v string) *RetryAddImageResponseBodyBuilder {
	b.s.SnapshotId = &v
	return b
}

func (b *RetryAddImageResponseBodyBuilder) ErrorMessage(v string) *RetryAddImageResponseBodyBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *RetryAddImageResponseBodyBuilder) TemplateId(v string) *RetryAddImageResponseBodyBuilder {
	b.s.TemplateId = &v
	return b
}

func (b *RetryAddImageResponseBodyBuilder) MachineName(v string) *RetryAddImageResponseBodyBuilder {
	b.s.MachineName = &v
	return b
}

func (b *RetryAddImageResponseBodyBuilder) MachineId(v string) *RetryAddImageResponseBodyBuilder {
	b.s.MachineId = &v
	return b
}

func (b *RetryAddImageResponseBodyBuilder) IsDeleted(v string) *RetryAddImageResponseBodyBuilder {
	b.s.IsDeleted = &v
	return b
}

func (b *RetryAddImageResponseBodyBuilder) RequestId(v string) *RetryAddImageResponseBodyBuilder {
	b.s.RequestId = &v
	return b
}

func (b *RetryAddImageResponseBodyBuilder) Alerted(v string) *RetryAddImageResponseBodyBuilder {
	b.s.Alerted = &v
	return b
}

func (b *RetryAddImageResponseBodyBuilder) CreatedTime(v string) *RetryAddImageResponseBodyBuilder {
	b.s.CreatedTime = &v
	return b
}

func (b *RetryAddImageResponseBodyBuilder) PoolCode(v string) *RetryAddImageResponseBodyBuilder {
	b.s.PoolCode = &v
	return b
}

func (b *RetryAddImageResponseBodyBuilder) Id(v int64) *RetryAddImageResponseBodyBuilder {
	b.s.Id = &v
	return b
}

func (b *RetryAddImageResponseBodyBuilder) BusinessType(v string) *RetryAddImageResponseBodyBuilder {
	b.s.BusinessType = &v
	return b
}

func (b *RetryAddImageResponseBodyBuilder) Region(v string) *RetryAddImageResponseBodyBuilder {
	b.s.Region = &v
	return b
}

func (b *RetryAddImageResponseBodyBuilder) Status(v string) *RetryAddImageResponseBodyBuilder {
	b.s.Status = &v
	return b
}

func (b *RetryAddImageResponseBodyBuilder) Build() *RetryAddImageResponseBody {
	return b.s
}


