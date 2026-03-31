// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetApplyImageRecordResponseData struct {

    // 厂商编码
	CompanyCode *string `json:"companyCode,omitempty"`
    // 镜像还原记录修改时间
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
    // 镜像还原记录删除标识，0标识未删除，1标识已删除
	IsDeleted *string `json:"isDeleted,omitempty"`
    // 与底层交互的请求id
	RequestId *string `json:"requestId,omitempty"`
    // 0:没有告警；1:已经告警
	Alerted *string `json:"alerted,omitempty"`
    // 租户主键id
	TenantId *int64 `json:"tenantId,omitempty"`
    // 镜像还原记录创建时间
	CreatedTime *string `json:"createdTime,omitempty"`
    // 资源池编码
	PoolCode *string `json:"poolCode,omitempty"`
    // 镜像还原记录主键id
	Id *int64 `json:"id,omitempty"`
    // 记录的业务类型：镜像：1，快照：2
	BusinessType *string `json:"businessType,omitempty"`
    // 可用区
	Region *string `json:"region,omitempty"`
    // 状态,0:快照恢复中（镜像重装中）,1:处理成功,2:处理失败
	Status *string `json:"status,omitempty"`
}

func (s GetApplyImageRecordResponseData) String() string {
	return utils.Beautify(s)
}

func (s GetApplyImageRecordResponseData) GoString() string {
	return s.String()
}

func (s GetApplyImageRecordResponseData) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetApplyImageRecordResponseData) SetCompanyCode(v string) *GetApplyImageRecordResponseData {
	s.CompanyCode = &v
	return s
}

func (s *GetApplyImageRecordResponseData) SetModifiedTime(v string) *GetApplyImageRecordResponseData {
	s.ModifiedTime = &v
	return s
}

func (s *GetApplyImageRecordResponseData) SetResourceId(v int64) *GetApplyImageRecordResponseData {
	s.ResourceId = &v
	return s
}

func (s *GetApplyImageRecordResponseData) SetSnapshotId(v string) *GetApplyImageRecordResponseData {
	s.SnapshotId = &v
	return s
}

func (s *GetApplyImageRecordResponseData) SetErrorMessage(v string) *GetApplyImageRecordResponseData {
	s.ErrorMessage = &v
	return s
}

func (s *GetApplyImageRecordResponseData) SetTemplateId(v string) *GetApplyImageRecordResponseData {
	s.TemplateId = &v
	return s
}

func (s *GetApplyImageRecordResponseData) SetMachineName(v string) *GetApplyImageRecordResponseData {
	s.MachineName = &v
	return s
}

func (s *GetApplyImageRecordResponseData) SetMachineId(v string) *GetApplyImageRecordResponseData {
	s.MachineId = &v
	return s
}

func (s *GetApplyImageRecordResponseData) SetIsDeleted(v string) *GetApplyImageRecordResponseData {
	s.IsDeleted = &v
	return s
}

func (s *GetApplyImageRecordResponseData) SetRequestId(v string) *GetApplyImageRecordResponseData {
	s.RequestId = &v
	return s
}

func (s *GetApplyImageRecordResponseData) SetAlerted(v string) *GetApplyImageRecordResponseData {
	s.Alerted = &v
	return s
}

func (s *GetApplyImageRecordResponseData) SetTenantId(v int64) *GetApplyImageRecordResponseData {
	s.TenantId = &v
	return s
}

func (s *GetApplyImageRecordResponseData) SetCreatedTime(v string) *GetApplyImageRecordResponseData {
	s.CreatedTime = &v
	return s
}

func (s *GetApplyImageRecordResponseData) SetPoolCode(v string) *GetApplyImageRecordResponseData {
	s.PoolCode = &v
	return s
}

func (s *GetApplyImageRecordResponseData) SetId(v int64) *GetApplyImageRecordResponseData {
	s.Id = &v
	return s
}

func (s *GetApplyImageRecordResponseData) SetBusinessType(v string) *GetApplyImageRecordResponseData {
	s.BusinessType = &v
	return s
}

func (s *GetApplyImageRecordResponseData) SetRegion(v string) *GetApplyImageRecordResponseData {
	s.Region = &v
	return s
}

func (s *GetApplyImageRecordResponseData) SetStatus(v string) *GetApplyImageRecordResponseData {
	s.Status = &v
	return s
}


type GetApplyImageRecordResponseDataBuilder struct {
	s *GetApplyImageRecordResponseData
}

func NewGetApplyImageRecordResponseDataBuilder() *GetApplyImageRecordResponseDataBuilder {
	s := &GetApplyImageRecordResponseData{}
	b := &GetApplyImageRecordResponseDataBuilder{s: s}
	return b
}

func (b *GetApplyImageRecordResponseDataBuilder) CompanyCode(v string) *GetApplyImageRecordResponseDataBuilder {
	b.s.CompanyCode = &v
	return b
}

func (b *GetApplyImageRecordResponseDataBuilder) ModifiedTime(v string) *GetApplyImageRecordResponseDataBuilder {
	b.s.ModifiedTime = &v
	return b
}

func (b *GetApplyImageRecordResponseDataBuilder) ResourceId(v int64) *GetApplyImageRecordResponseDataBuilder {
	b.s.ResourceId = &v
	return b
}

func (b *GetApplyImageRecordResponseDataBuilder) SnapshotId(v string) *GetApplyImageRecordResponseDataBuilder {
	b.s.SnapshotId = &v
	return b
}

func (b *GetApplyImageRecordResponseDataBuilder) ErrorMessage(v string) *GetApplyImageRecordResponseDataBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *GetApplyImageRecordResponseDataBuilder) TemplateId(v string) *GetApplyImageRecordResponseDataBuilder {
	b.s.TemplateId = &v
	return b
}

func (b *GetApplyImageRecordResponseDataBuilder) MachineName(v string) *GetApplyImageRecordResponseDataBuilder {
	b.s.MachineName = &v
	return b
}

func (b *GetApplyImageRecordResponseDataBuilder) MachineId(v string) *GetApplyImageRecordResponseDataBuilder {
	b.s.MachineId = &v
	return b
}

func (b *GetApplyImageRecordResponseDataBuilder) IsDeleted(v string) *GetApplyImageRecordResponseDataBuilder {
	b.s.IsDeleted = &v
	return b
}

func (b *GetApplyImageRecordResponseDataBuilder) RequestId(v string) *GetApplyImageRecordResponseDataBuilder {
	b.s.RequestId = &v
	return b
}

func (b *GetApplyImageRecordResponseDataBuilder) Alerted(v string) *GetApplyImageRecordResponseDataBuilder {
	b.s.Alerted = &v
	return b
}

func (b *GetApplyImageRecordResponseDataBuilder) TenantId(v int64) *GetApplyImageRecordResponseDataBuilder {
	b.s.TenantId = &v
	return b
}

func (b *GetApplyImageRecordResponseDataBuilder) CreatedTime(v string) *GetApplyImageRecordResponseDataBuilder {
	b.s.CreatedTime = &v
	return b
}

func (b *GetApplyImageRecordResponseDataBuilder) PoolCode(v string) *GetApplyImageRecordResponseDataBuilder {
	b.s.PoolCode = &v
	return b
}

func (b *GetApplyImageRecordResponseDataBuilder) Id(v int64) *GetApplyImageRecordResponseDataBuilder {
	b.s.Id = &v
	return b
}

func (b *GetApplyImageRecordResponseDataBuilder) BusinessType(v string) *GetApplyImageRecordResponseDataBuilder {
	b.s.BusinessType = &v
	return b
}

func (b *GetApplyImageRecordResponseDataBuilder) Region(v string) *GetApplyImageRecordResponseDataBuilder {
	b.s.Region = &v
	return b
}

func (b *GetApplyImageRecordResponseDataBuilder) Status(v string) *GetApplyImageRecordResponseDataBuilder {
	b.s.Status = &v
	return b
}

func (b *GetApplyImageRecordResponseDataBuilder) Build() *GetApplyImageRecordResponseData {
	return b.s
}


