// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetApplySnapshotRecordResponseData struct {

    // 厂商编码
	CompanyCode *string `json:"companyCode,omitempty"`
    // 快照还原记录修改时间
	ModifiedTime *string `json:"modifiedTime,omitempty"`
    // 资源表主键
	ResourceId *int64 `json:"resourceId,omitempty"`
    // 快照uid
	SnapshotId *string `json:"snapshotId,omitempty"`
    // 错误信息
	ErrorMessage *string `json:"errorMessage,omitempty"`
    // 镜像模板uid
	TemplateId *string `json:"templateId,omitempty"`
    // 云电脑名称
	MachineName *string `json:"machineName,omitempty"`
    // 云电脑ID
	MachineId *string `json:"machineId,omitempty"`
    // 快照还原记录删除标识，0表示未删除，1表示已删除
	IsDeleted *string `json:"isDeleted,omitempty"`
    // 与底层交互的请求id
	RequestId *string `json:"requestId,omitempty"`
    // 0:没有告警；1:已经告警
	Alerted *string `json:"alerted,omitempty"`
    // 快照还原记录创建时间
	CreatedTime *string `json:"createdTime,omitempty"`
    // 资源池编码
	PoolCode *string `json:"poolCode,omitempty"`
    // 快照还原记录主键id
	Id *int64 `json:"id,omitempty"`
    // 记录的业务类型：镜像：1，快照：2
	BusinessType *string `json:"businessType,omitempty"`
    // 可用区
	Region *string `json:"region,omitempty"`
    // 状态,0:快照恢复中（镜像重装中）,1:处理成功,2:处理失败
	Status *string `json:"status,omitempty"`
}

func (s GetApplySnapshotRecordResponseData) String() string {
	return utils.Beautify(s)
}

func (s GetApplySnapshotRecordResponseData) GoString() string {
	return s.String()
}

func (s GetApplySnapshotRecordResponseData) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetApplySnapshotRecordResponseData) SetCompanyCode(v string) *GetApplySnapshotRecordResponseData {
	s.CompanyCode = &v
	return s
}

func (s *GetApplySnapshotRecordResponseData) SetModifiedTime(v string) *GetApplySnapshotRecordResponseData {
	s.ModifiedTime = &v
	return s
}

func (s *GetApplySnapshotRecordResponseData) SetResourceId(v int64) *GetApplySnapshotRecordResponseData {
	s.ResourceId = &v
	return s
}

func (s *GetApplySnapshotRecordResponseData) SetSnapshotId(v string) *GetApplySnapshotRecordResponseData {
	s.SnapshotId = &v
	return s
}

func (s *GetApplySnapshotRecordResponseData) SetErrorMessage(v string) *GetApplySnapshotRecordResponseData {
	s.ErrorMessage = &v
	return s
}

func (s *GetApplySnapshotRecordResponseData) SetTemplateId(v string) *GetApplySnapshotRecordResponseData {
	s.TemplateId = &v
	return s
}

func (s *GetApplySnapshotRecordResponseData) SetMachineName(v string) *GetApplySnapshotRecordResponseData {
	s.MachineName = &v
	return s
}

func (s *GetApplySnapshotRecordResponseData) SetMachineId(v string) *GetApplySnapshotRecordResponseData {
	s.MachineId = &v
	return s
}

func (s *GetApplySnapshotRecordResponseData) SetIsDeleted(v string) *GetApplySnapshotRecordResponseData {
	s.IsDeleted = &v
	return s
}

func (s *GetApplySnapshotRecordResponseData) SetRequestId(v string) *GetApplySnapshotRecordResponseData {
	s.RequestId = &v
	return s
}

func (s *GetApplySnapshotRecordResponseData) SetAlerted(v string) *GetApplySnapshotRecordResponseData {
	s.Alerted = &v
	return s
}

func (s *GetApplySnapshotRecordResponseData) SetCreatedTime(v string) *GetApplySnapshotRecordResponseData {
	s.CreatedTime = &v
	return s
}

func (s *GetApplySnapshotRecordResponseData) SetPoolCode(v string) *GetApplySnapshotRecordResponseData {
	s.PoolCode = &v
	return s
}

func (s *GetApplySnapshotRecordResponseData) SetId(v int64) *GetApplySnapshotRecordResponseData {
	s.Id = &v
	return s
}

func (s *GetApplySnapshotRecordResponseData) SetBusinessType(v string) *GetApplySnapshotRecordResponseData {
	s.BusinessType = &v
	return s
}

func (s *GetApplySnapshotRecordResponseData) SetRegion(v string) *GetApplySnapshotRecordResponseData {
	s.Region = &v
	return s
}

func (s *GetApplySnapshotRecordResponseData) SetStatus(v string) *GetApplySnapshotRecordResponseData {
	s.Status = &v
	return s
}


type GetApplySnapshotRecordResponseDataBuilder struct {
	s *GetApplySnapshotRecordResponseData
}

func NewGetApplySnapshotRecordResponseDataBuilder() *GetApplySnapshotRecordResponseDataBuilder {
	s := &GetApplySnapshotRecordResponseData{}
	b := &GetApplySnapshotRecordResponseDataBuilder{s: s}
	return b
}

func (b *GetApplySnapshotRecordResponseDataBuilder) CompanyCode(v string) *GetApplySnapshotRecordResponseDataBuilder {
	b.s.CompanyCode = &v
	return b
}

func (b *GetApplySnapshotRecordResponseDataBuilder) ModifiedTime(v string) *GetApplySnapshotRecordResponseDataBuilder {
	b.s.ModifiedTime = &v
	return b
}

func (b *GetApplySnapshotRecordResponseDataBuilder) ResourceId(v int64) *GetApplySnapshotRecordResponseDataBuilder {
	b.s.ResourceId = &v
	return b
}

func (b *GetApplySnapshotRecordResponseDataBuilder) SnapshotId(v string) *GetApplySnapshotRecordResponseDataBuilder {
	b.s.SnapshotId = &v
	return b
}

func (b *GetApplySnapshotRecordResponseDataBuilder) ErrorMessage(v string) *GetApplySnapshotRecordResponseDataBuilder {
	b.s.ErrorMessage = &v
	return b
}

func (b *GetApplySnapshotRecordResponseDataBuilder) TemplateId(v string) *GetApplySnapshotRecordResponseDataBuilder {
	b.s.TemplateId = &v
	return b
}

func (b *GetApplySnapshotRecordResponseDataBuilder) MachineName(v string) *GetApplySnapshotRecordResponseDataBuilder {
	b.s.MachineName = &v
	return b
}

func (b *GetApplySnapshotRecordResponseDataBuilder) MachineId(v string) *GetApplySnapshotRecordResponseDataBuilder {
	b.s.MachineId = &v
	return b
}

func (b *GetApplySnapshotRecordResponseDataBuilder) IsDeleted(v string) *GetApplySnapshotRecordResponseDataBuilder {
	b.s.IsDeleted = &v
	return b
}

func (b *GetApplySnapshotRecordResponseDataBuilder) RequestId(v string) *GetApplySnapshotRecordResponseDataBuilder {
	b.s.RequestId = &v
	return b
}

func (b *GetApplySnapshotRecordResponseDataBuilder) Alerted(v string) *GetApplySnapshotRecordResponseDataBuilder {
	b.s.Alerted = &v
	return b
}

func (b *GetApplySnapshotRecordResponseDataBuilder) CreatedTime(v string) *GetApplySnapshotRecordResponseDataBuilder {
	b.s.CreatedTime = &v
	return b
}

func (b *GetApplySnapshotRecordResponseDataBuilder) PoolCode(v string) *GetApplySnapshotRecordResponseDataBuilder {
	b.s.PoolCode = &v
	return b
}

func (b *GetApplySnapshotRecordResponseDataBuilder) Id(v int64) *GetApplySnapshotRecordResponseDataBuilder {
	b.s.Id = &v
	return b
}

func (b *GetApplySnapshotRecordResponseDataBuilder) BusinessType(v string) *GetApplySnapshotRecordResponseDataBuilder {
	b.s.BusinessType = &v
	return b
}

func (b *GetApplySnapshotRecordResponseDataBuilder) Region(v string) *GetApplySnapshotRecordResponseDataBuilder {
	b.s.Region = &v
	return b
}

func (b *GetApplySnapshotRecordResponseDataBuilder) Status(v string) *GetApplySnapshotRecordResponseDataBuilder {
	b.s.Status = &v
	return b
}

func (b *GetApplySnapshotRecordResponseDataBuilder) Build() *GetApplySnapshotRecordResponseData {
	return b.s
}


