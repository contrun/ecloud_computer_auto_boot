// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetInstanceListCemResponseRecords struct {

    // 数据盘
	DataDiskExt *string `json:"dataDiskExt,omitempty"`
    // 操作系统类型
	ImageName *string `json:"imageName,omitempty"`
    // 高级服务包:为空表示未订购
	BwCompanyCode *string `json:"bwCompanyCode,omitempty"`
    // 套餐类型标识， 1：办公型-公众版；办公型-政企版（不带数据盘） 2：办公型-公众版；办公型-政企版（带数据盘） 3：信创型-政企版 4：行业型-娱乐版
	Combo *int32 `json:"combo,omitempty"`
    // 计费方式(month:包月计费；year：包年计费）
	MeasureUnit *string `json:"measureUnit,omitempty"`
    // 分类类型， 1：办公型-公众版；行业型-娱乐版 3：办公型-政企版；信创型-政企版 5：办公型-政企版（边缘节点）
	CategoryCode *string `json:"categoryCode,omitempty"`
    // ram账号
	RamUserName *string `json:"ramUserName,omitempty"`
    // 用户名
	LoginUserName *string `json:"loginUserName,omitempty"`
    // 订单批次号
	MopO *string `json:"mopO,omitempty"`
    // 订购资源ID
	InstanceId *string `json:"instanceId,omitempty"`
    // 订单编号
	MopT *string `json:"mopT,omitempty"`
    // 续订方式
	AutoRenew *bool `json:"autoRenew,omitempty"`
    // 激活邮箱
	OrderUserMail *string `json:"orderUserMail,omitempty"`
    // 开通时间
	EffTime *string `json:"effTime,omitempty"`
    // 云电脑机器ID
	DesktopId *string `json:"desktopId,omitempty"`
    // 到期时间
	EndTime *string `json:"endTime,omitempty"`
    // 规格编码  0：办公型-公众版2vCPU/4GB内存；办公型-公众版2vCPU/4GB内存/200GB数据盘；行业型-娱乐版5小时月包；办公型-政企版2vCPU/4GB内存；办公型-政企版2vCPU/4GB内存/200GB数据盘；办公型-政企版（边缘节点）2vCPU/4GB内存；办公型-政企版（边缘节点）2vCPU/4GB内存/200GB数据盘  1：办公型-公众版4vCPU/8GB内存；办公型-公众版4vCPU/8GB内存/500GB数据盘；行业型-娱乐版15小时月包；办公型-政企版4vCPU/8GB内存；办公型-政企版4vCPU/8GB内存/500GB数据盘；办公型-政企版（边缘节点）4vCPU/8GB内存；办公型-政企版（边缘节点）4vCPU/8GB内存/500GB数据盘；信创型-政企版4vCPU/8GB内存  2：办公型-公众版8vCPU/16GB内存；办公型-公众版8vCPU/16GB内存/1TB数据盘；行业型-娱乐版30小时月包；办公型-政企版8vCPU/16GB内存；办公型-政企版8vCPU/16GB内存/1TB数据盘；办公型-政企版（边缘节点）8vCPU/16GB内存；办公型-政企版（边缘节点）8vCPU/16GB内存/1TB数据盘；信创型-政企版8vCPU/16GB内存  3：办公型-政企版6vCPU/12GB内存；办公型-政企版6vCPU/12GB内存/200GB数据盘  4:办公型-公众版4vCPU/4GB内存；办公型-政企版4vCPU/4GB内存；办公型-政企版（边缘节点）4vCPU/4GB内存；信创型-政企版4vCPU/4GB内存。  5:办公型-公众版8vCPU/8GB内存；办公型-政企版8vCPU/8GB内存；办公型-政企版（边缘节点）8vCPU/8GB内存；信创型-政企版8vCPU/8GB内存。  6:办公型-公众版16vCPU/16GB内存；办公型-政企版16vCPU/16GB内存；办公型-政企版（边缘节点）16vCPU/16GB内存；信创型-政企版16vCPU/16GB内存。  7:办公型-公众版16vCPU/32GB内存；办公型-政企版16vCPU/32GB内存；办公型-政企版（边缘节点）16vCPU/32GB内存；信创型-政企版16vCPU/32GB内存。
	SpecificationCode *string `json:"specificationCode,omitempty"`
    // 状态枚举（2:已开通、1:开通中、10:开通失败、5:退订中、6/601:已退订、8:冻结中、7:已冻结、3:续订中、4:已到期（到期冻结期间）、16:续订冻结中、18:变更中、20:变更成功、1002:取消订单）
	Status *int32 `json:"status,omitempty"`
}

func (s GetInstanceListCemResponseRecords) String() string {
	return utils.Beautify(s)
}

func (s GetInstanceListCemResponseRecords) GoString() string {
	return s.String()
}

func (s GetInstanceListCemResponseRecords) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetInstanceListCemResponseRecords) SetDataDiskExt(v string) *GetInstanceListCemResponseRecords {
	s.DataDiskExt = &v
	return s
}

func (s *GetInstanceListCemResponseRecords) SetImageName(v string) *GetInstanceListCemResponseRecords {
	s.ImageName = &v
	return s
}

func (s *GetInstanceListCemResponseRecords) SetBwCompanyCode(v string) *GetInstanceListCemResponseRecords {
	s.BwCompanyCode = &v
	return s
}

func (s *GetInstanceListCemResponseRecords) SetCombo(v int32) *GetInstanceListCemResponseRecords {
	s.Combo = &v
	return s
}

func (s *GetInstanceListCemResponseRecords) SetMeasureUnit(v string) *GetInstanceListCemResponseRecords {
	s.MeasureUnit = &v
	return s
}

func (s *GetInstanceListCemResponseRecords) SetCategoryCode(v string) *GetInstanceListCemResponseRecords {
	s.CategoryCode = &v
	return s
}

func (s *GetInstanceListCemResponseRecords) SetRamUserName(v string) *GetInstanceListCemResponseRecords {
	s.RamUserName = &v
	return s
}

func (s *GetInstanceListCemResponseRecords) SetLoginUserName(v string) *GetInstanceListCemResponseRecords {
	s.LoginUserName = &v
	return s
}

func (s *GetInstanceListCemResponseRecords) SetMopO(v string) *GetInstanceListCemResponseRecords {
	s.MopO = &v
	return s
}

func (s *GetInstanceListCemResponseRecords) SetInstanceId(v string) *GetInstanceListCemResponseRecords {
	s.InstanceId = &v
	return s
}

func (s *GetInstanceListCemResponseRecords) SetMopT(v string) *GetInstanceListCemResponseRecords {
	s.MopT = &v
	return s
}

func (s *GetInstanceListCemResponseRecords) SetAutoRenew(v bool) *GetInstanceListCemResponseRecords {
	s.AutoRenew = &v
	return s
}

func (s *GetInstanceListCemResponseRecords) SetOrderUserMail(v string) *GetInstanceListCemResponseRecords {
	s.OrderUserMail = &v
	return s
}

func (s *GetInstanceListCemResponseRecords) SetEffTime(v string) *GetInstanceListCemResponseRecords {
	s.EffTime = &v
	return s
}

func (s *GetInstanceListCemResponseRecords) SetDesktopId(v string) *GetInstanceListCemResponseRecords {
	s.DesktopId = &v
	return s
}

func (s *GetInstanceListCemResponseRecords) SetEndTime(v string) *GetInstanceListCemResponseRecords {
	s.EndTime = &v
	return s
}

func (s *GetInstanceListCemResponseRecords) SetSpecificationCode(v string) *GetInstanceListCemResponseRecords {
	s.SpecificationCode = &v
	return s
}

func (s *GetInstanceListCemResponseRecords) SetStatus(v int32) *GetInstanceListCemResponseRecords {
	s.Status = &v
	return s
}


type GetInstanceListCemResponseRecordsBuilder struct {
	s *GetInstanceListCemResponseRecords
}

func NewGetInstanceListCemResponseRecordsBuilder() *GetInstanceListCemResponseRecordsBuilder {
	s := &GetInstanceListCemResponseRecords{}
	b := &GetInstanceListCemResponseRecordsBuilder{s: s}
	return b
}

func (b *GetInstanceListCemResponseRecordsBuilder) DataDiskExt(v string) *GetInstanceListCemResponseRecordsBuilder {
	b.s.DataDiskExt = &v
	return b
}

func (b *GetInstanceListCemResponseRecordsBuilder) ImageName(v string) *GetInstanceListCemResponseRecordsBuilder {
	b.s.ImageName = &v
	return b
}

func (b *GetInstanceListCemResponseRecordsBuilder) BwCompanyCode(v string) *GetInstanceListCemResponseRecordsBuilder {
	b.s.BwCompanyCode = &v
	return b
}

func (b *GetInstanceListCemResponseRecordsBuilder) Combo(v int32) *GetInstanceListCemResponseRecordsBuilder {
	b.s.Combo = &v
	return b
}

func (b *GetInstanceListCemResponseRecordsBuilder) MeasureUnit(v string) *GetInstanceListCemResponseRecordsBuilder {
	b.s.MeasureUnit = &v
	return b
}

func (b *GetInstanceListCemResponseRecordsBuilder) CategoryCode(v string) *GetInstanceListCemResponseRecordsBuilder {
	b.s.CategoryCode = &v
	return b
}

func (b *GetInstanceListCemResponseRecordsBuilder) RamUserName(v string) *GetInstanceListCemResponseRecordsBuilder {
	b.s.RamUserName = &v
	return b
}

func (b *GetInstanceListCemResponseRecordsBuilder) LoginUserName(v string) *GetInstanceListCemResponseRecordsBuilder {
	b.s.LoginUserName = &v
	return b
}

func (b *GetInstanceListCemResponseRecordsBuilder) MopO(v string) *GetInstanceListCemResponseRecordsBuilder {
	b.s.MopO = &v
	return b
}

func (b *GetInstanceListCemResponseRecordsBuilder) InstanceId(v string) *GetInstanceListCemResponseRecordsBuilder {
	b.s.InstanceId = &v
	return b
}

func (b *GetInstanceListCemResponseRecordsBuilder) MopT(v string) *GetInstanceListCemResponseRecordsBuilder {
	b.s.MopT = &v
	return b
}

func (b *GetInstanceListCemResponseRecordsBuilder) AutoRenew(v bool) *GetInstanceListCemResponseRecordsBuilder {
	b.s.AutoRenew = &v
	return b
}

func (b *GetInstanceListCemResponseRecordsBuilder) OrderUserMail(v string) *GetInstanceListCemResponseRecordsBuilder {
	b.s.OrderUserMail = &v
	return b
}

func (b *GetInstanceListCemResponseRecordsBuilder) EffTime(v string) *GetInstanceListCemResponseRecordsBuilder {
	b.s.EffTime = &v
	return b
}

func (b *GetInstanceListCemResponseRecordsBuilder) DesktopId(v string) *GetInstanceListCemResponseRecordsBuilder {
	b.s.DesktopId = &v
	return b
}

func (b *GetInstanceListCemResponseRecordsBuilder) EndTime(v string) *GetInstanceListCemResponseRecordsBuilder {
	b.s.EndTime = &v
	return b
}

func (b *GetInstanceListCemResponseRecordsBuilder) SpecificationCode(v string) *GetInstanceListCemResponseRecordsBuilder {
	b.s.SpecificationCode = &v
	return b
}

func (b *GetInstanceListCemResponseRecordsBuilder) Status(v int32) *GetInstanceListCemResponseRecordsBuilder {
	b.s.Status = &v
	return b
}

func (b *GetInstanceListCemResponseRecordsBuilder) Build() *GetInstanceListCemResponseRecords {
	return b.s
}


