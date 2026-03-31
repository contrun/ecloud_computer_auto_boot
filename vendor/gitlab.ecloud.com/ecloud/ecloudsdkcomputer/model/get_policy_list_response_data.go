// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetPolicyListResponseData struct {

    // 修改时间
	ModifiedTime *string `json:"modifiedTime,omitempty"`
    // 关联电脑数
	BindMachineCount *int32 `json:"bindMachineCount,omitempty"`
    // 策略名称
	PolicyName *string `json:"policyName,omitempty"`
    // 厂商名称
	CompanyName *string `json:"companyName,omitempty"`
    // 可用区名称
	RegionName *string `json:"regionName,omitempty"`
    // 策略状态
	PolicyStatus *string `json:"policyStatus,omitempty"`
    // 策略id
	PolicyId *string `json:"policyId,omitempty"`
    // 策略描述
	PolicyRemark *string `json:"policyRemark,omitempty"`
    // 创建时间
	CreatedTime *string `json:"createdTime,omitempty"`
    // 资源池编码
	PoolCode *string `json:"poolCode,omitempty"`
    // 可用区
	Region *string `json:"region,omitempty"`
    // 网络工作区
	Dc *string `json:"dc,omitempty"`
    // 资源池名称
	PoolName *string `json:"poolName,omitempty"`
}

func (s GetPolicyListResponseData) String() string {
	return utils.Beautify(s)
}

func (s GetPolicyListResponseData) GoString() string {
	return s.String()
}

func (s GetPolicyListResponseData) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetPolicyListResponseData) SetModifiedTime(v string) *GetPolicyListResponseData {
	s.ModifiedTime = &v
	return s
}

func (s *GetPolicyListResponseData) SetBindMachineCount(v int32) *GetPolicyListResponseData {
	s.BindMachineCount = &v
	return s
}

func (s *GetPolicyListResponseData) SetPolicyName(v string) *GetPolicyListResponseData {
	s.PolicyName = &v
	return s
}

func (s *GetPolicyListResponseData) SetCompanyName(v string) *GetPolicyListResponseData {
	s.CompanyName = &v
	return s
}

func (s *GetPolicyListResponseData) SetRegionName(v string) *GetPolicyListResponseData {
	s.RegionName = &v
	return s
}

func (s *GetPolicyListResponseData) SetPolicyStatus(v string) *GetPolicyListResponseData {
	s.PolicyStatus = &v
	return s
}

func (s *GetPolicyListResponseData) SetPolicyId(v string) *GetPolicyListResponseData {
	s.PolicyId = &v
	return s
}

func (s *GetPolicyListResponseData) SetPolicyRemark(v string) *GetPolicyListResponseData {
	s.PolicyRemark = &v
	return s
}

func (s *GetPolicyListResponseData) SetCreatedTime(v string) *GetPolicyListResponseData {
	s.CreatedTime = &v
	return s
}

func (s *GetPolicyListResponseData) SetPoolCode(v string) *GetPolicyListResponseData {
	s.PoolCode = &v
	return s
}

func (s *GetPolicyListResponseData) SetRegion(v string) *GetPolicyListResponseData {
	s.Region = &v
	return s
}

func (s *GetPolicyListResponseData) SetDc(v string) *GetPolicyListResponseData {
	s.Dc = &v
	return s
}

func (s *GetPolicyListResponseData) SetPoolName(v string) *GetPolicyListResponseData {
	s.PoolName = &v
	return s
}


type GetPolicyListResponseDataBuilder struct {
	s *GetPolicyListResponseData
}

func NewGetPolicyListResponseDataBuilder() *GetPolicyListResponseDataBuilder {
	s := &GetPolicyListResponseData{}
	b := &GetPolicyListResponseDataBuilder{s: s}
	return b
}

func (b *GetPolicyListResponseDataBuilder) ModifiedTime(v string) *GetPolicyListResponseDataBuilder {
	b.s.ModifiedTime = &v
	return b
}

func (b *GetPolicyListResponseDataBuilder) BindMachineCount(v int32) *GetPolicyListResponseDataBuilder {
	b.s.BindMachineCount = &v
	return b
}

func (b *GetPolicyListResponseDataBuilder) PolicyName(v string) *GetPolicyListResponseDataBuilder {
	b.s.PolicyName = &v
	return b
}

func (b *GetPolicyListResponseDataBuilder) CompanyName(v string) *GetPolicyListResponseDataBuilder {
	b.s.CompanyName = &v
	return b
}

func (b *GetPolicyListResponseDataBuilder) RegionName(v string) *GetPolicyListResponseDataBuilder {
	b.s.RegionName = &v
	return b
}

func (b *GetPolicyListResponseDataBuilder) PolicyStatus(v string) *GetPolicyListResponseDataBuilder {
	b.s.PolicyStatus = &v
	return b
}

func (b *GetPolicyListResponseDataBuilder) PolicyId(v string) *GetPolicyListResponseDataBuilder {
	b.s.PolicyId = &v
	return b
}

func (b *GetPolicyListResponseDataBuilder) PolicyRemark(v string) *GetPolicyListResponseDataBuilder {
	b.s.PolicyRemark = &v
	return b
}

func (b *GetPolicyListResponseDataBuilder) CreatedTime(v string) *GetPolicyListResponseDataBuilder {
	b.s.CreatedTime = &v
	return b
}

func (b *GetPolicyListResponseDataBuilder) PoolCode(v string) *GetPolicyListResponseDataBuilder {
	b.s.PoolCode = &v
	return b
}

func (b *GetPolicyListResponseDataBuilder) Region(v string) *GetPolicyListResponseDataBuilder {
	b.s.Region = &v
	return b
}

func (b *GetPolicyListResponseDataBuilder) Dc(v string) *GetPolicyListResponseDataBuilder {
	b.s.Dc = &v
	return b
}

func (b *GetPolicyListResponseDataBuilder) PoolName(v string) *GetPolicyListResponseDataBuilder {
	b.s.PoolName = &v
	return b
}

func (b *GetPolicyListResponseDataBuilder) Build() *GetPolicyListResponseData {
	return b.s
}


