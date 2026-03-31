// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type PolicyListResponseData struct {

    // 修改时间
	ModifiedTime *string `json:"modifiedTime,omitempty"`
    // 定时任务策略uuid
	PolicyUid *string `json:"policyUid,omitempty"`
    // 定时任务策略主键id
	PolicyId *string `json:"policyId,omitempty"`
    // 策略名称
	PolicyName *string `json:"policyName,omitempty"`
    // 关联云电脑数量
	BindNum *int32 `json:"bindNum,omitempty"`
    // 策略描述
	PolicyDesc *string `json:"policyDesc,omitempty"`
    // 租户名称
	MopUserName *string `json:"mopUserName,omitempty"`
    // 资源池编码
	PoolCode *string `json:"poolCode,omitempty"`
    // 租户id
	MopUserId *string `json:"mopUserId,omitempty"`
    // 状态：创建成功（create），已删除（deleted）
	Status *string `json:"status,omitempty"`
    // 资源池名称
	PoolName *string `json:"poolName,omitempty"`
}

func (s PolicyListResponseData) String() string {
	return utils.Beautify(s)
}

func (s PolicyListResponseData) GoString() string {
	return s.String()
}

func (s PolicyListResponseData) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *PolicyListResponseData) SetModifiedTime(v string) *PolicyListResponseData {
	s.ModifiedTime = &v
	return s
}

func (s *PolicyListResponseData) SetPolicyUid(v string) *PolicyListResponseData {
	s.PolicyUid = &v
	return s
}

func (s *PolicyListResponseData) SetPolicyId(v string) *PolicyListResponseData {
	s.PolicyId = &v
	return s
}

func (s *PolicyListResponseData) SetPolicyName(v string) *PolicyListResponseData {
	s.PolicyName = &v
	return s
}

func (s *PolicyListResponseData) SetBindNum(v int32) *PolicyListResponseData {
	s.BindNum = &v
	return s
}

func (s *PolicyListResponseData) SetPolicyDesc(v string) *PolicyListResponseData {
	s.PolicyDesc = &v
	return s
}

func (s *PolicyListResponseData) SetMopUserName(v string) *PolicyListResponseData {
	s.MopUserName = &v
	return s
}

func (s *PolicyListResponseData) SetPoolCode(v string) *PolicyListResponseData {
	s.PoolCode = &v
	return s
}

func (s *PolicyListResponseData) SetMopUserId(v string) *PolicyListResponseData {
	s.MopUserId = &v
	return s
}

func (s *PolicyListResponseData) SetStatus(v string) *PolicyListResponseData {
	s.Status = &v
	return s
}

func (s *PolicyListResponseData) SetPoolName(v string) *PolicyListResponseData {
	s.PoolName = &v
	return s
}


type PolicyListResponseDataBuilder struct {
	s *PolicyListResponseData
}

func NewPolicyListResponseDataBuilder() *PolicyListResponseDataBuilder {
	s := &PolicyListResponseData{}
	b := &PolicyListResponseDataBuilder{s: s}
	return b
}

func (b *PolicyListResponseDataBuilder) ModifiedTime(v string) *PolicyListResponseDataBuilder {
	b.s.ModifiedTime = &v
	return b
}

func (b *PolicyListResponseDataBuilder) PolicyUid(v string) *PolicyListResponseDataBuilder {
	b.s.PolicyUid = &v
	return b
}

func (b *PolicyListResponseDataBuilder) PolicyId(v string) *PolicyListResponseDataBuilder {
	b.s.PolicyId = &v
	return b
}

func (b *PolicyListResponseDataBuilder) PolicyName(v string) *PolicyListResponseDataBuilder {
	b.s.PolicyName = &v
	return b
}

func (b *PolicyListResponseDataBuilder) BindNum(v int32) *PolicyListResponseDataBuilder {
	b.s.BindNum = &v
	return b
}

func (b *PolicyListResponseDataBuilder) PolicyDesc(v string) *PolicyListResponseDataBuilder {
	b.s.PolicyDesc = &v
	return b
}

func (b *PolicyListResponseDataBuilder) MopUserName(v string) *PolicyListResponseDataBuilder {
	b.s.MopUserName = &v
	return b
}

func (b *PolicyListResponseDataBuilder) PoolCode(v string) *PolicyListResponseDataBuilder {
	b.s.PoolCode = &v
	return b
}

func (b *PolicyListResponseDataBuilder) MopUserId(v string) *PolicyListResponseDataBuilder {
	b.s.MopUserId = &v
	return b
}

func (b *PolicyListResponseDataBuilder) Status(v string) *PolicyListResponseDataBuilder {
	b.s.Status = &v
	return b
}

func (b *PolicyListResponseDataBuilder) PoolName(v string) *PolicyListResponseDataBuilder {
	b.s.PoolName = &v
	return b
}

func (b *PolicyListResponseDataBuilder) Build() *PolicyListResponseData {
	return b.s
}


