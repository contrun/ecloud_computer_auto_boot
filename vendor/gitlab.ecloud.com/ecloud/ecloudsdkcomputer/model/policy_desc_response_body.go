// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type PolicyDescResponseBody struct {

    // 修改时间
	ModifiedTime *string `json:"modifiedTime,omitempty"`
    // 创建人
	Creator *string `json:"creator,omitempty"`
    // 策略名称
	PolicyName *string `json:"policyName,omitempty"`
    // 修改人
	Modifier *string `json:"modifier,omitempty"`
    // 任务详情
	TaskDescInfoList *[]PolicyDescResponseTaskDescInfoList `json:"taskDescInfoList,omitempty"`
    // 定时任务策略uuid
	PolicyUid *string `json:"policyUid,omitempty"`
    // 策略描述
	PolicyDesc *string `json:"policyDesc,omitempty"`
    // 租户主键id
	TenantId *string `json:"tenantId,omitempty"`
    // 创建时间
	CreatedTime *string `json:"createdTime,omitempty"`
    // 租户名称
	MopUserName *string `json:"mopUserName,omitempty"`
    // 所属资源池编码
	PoolCode *string `json:"poolCode,omitempty"`
    // 租户id
	MopUserId *string `json:"mopUserId,omitempty"`
    // 所属资源池名称
	PoolName *string `json:"poolName,omitempty"`
    // create:创建成功,deleted:已删除
	Status *string `json:"status,omitempty"`
}

func (s PolicyDescResponseBody) String() string {
	return utils.Beautify(s)
}

func (s PolicyDescResponseBody) GoString() string {
	return s.String()
}

func (s PolicyDescResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *PolicyDescResponseBody) SetModifiedTime(v string) *PolicyDescResponseBody {
	s.ModifiedTime = &v
	return s
}

func (s *PolicyDescResponseBody) SetCreator(v string) *PolicyDescResponseBody {
	s.Creator = &v
	return s
}

func (s *PolicyDescResponseBody) SetPolicyName(v string) *PolicyDescResponseBody {
	s.PolicyName = &v
	return s
}

func (s *PolicyDescResponseBody) SetModifier(v string) *PolicyDescResponseBody {
	s.Modifier = &v
	return s
}

func (s *PolicyDescResponseBody) SetTaskDescInfoList(v []PolicyDescResponseTaskDescInfoList) *PolicyDescResponseBody {
	s.TaskDescInfoList = &v
	return s
}

func (s *PolicyDescResponseBody) SetPolicyUid(v string) *PolicyDescResponseBody {
	s.PolicyUid = &v
	return s
}

func (s *PolicyDescResponseBody) SetPolicyDesc(v string) *PolicyDescResponseBody {
	s.PolicyDesc = &v
	return s
}

func (s *PolicyDescResponseBody) SetTenantId(v string) *PolicyDescResponseBody {
	s.TenantId = &v
	return s
}

func (s *PolicyDescResponseBody) SetCreatedTime(v string) *PolicyDescResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *PolicyDescResponseBody) SetMopUserName(v string) *PolicyDescResponseBody {
	s.MopUserName = &v
	return s
}

func (s *PolicyDescResponseBody) SetPoolCode(v string) *PolicyDescResponseBody {
	s.PoolCode = &v
	return s
}

func (s *PolicyDescResponseBody) SetMopUserId(v string) *PolicyDescResponseBody {
	s.MopUserId = &v
	return s
}

func (s *PolicyDescResponseBody) SetPoolName(v string) *PolicyDescResponseBody {
	s.PoolName = &v
	return s
}

func (s *PolicyDescResponseBody) SetStatus(v string) *PolicyDescResponseBody {
	s.Status = &v
	return s
}


type PolicyDescResponseBodyBuilder struct {
	s *PolicyDescResponseBody
}

func NewPolicyDescResponseBodyBuilder() *PolicyDescResponseBodyBuilder {
	s := &PolicyDescResponseBody{}
	b := &PolicyDescResponseBodyBuilder{s: s}
	return b
}

func (b *PolicyDescResponseBodyBuilder) ModifiedTime(v string) *PolicyDescResponseBodyBuilder {
	b.s.ModifiedTime = &v
	return b
}

func (b *PolicyDescResponseBodyBuilder) Creator(v string) *PolicyDescResponseBodyBuilder {
	b.s.Creator = &v
	return b
}

func (b *PolicyDescResponseBodyBuilder) PolicyName(v string) *PolicyDescResponseBodyBuilder {
	b.s.PolicyName = &v
	return b
}

func (b *PolicyDescResponseBodyBuilder) Modifier(v string) *PolicyDescResponseBodyBuilder {
	b.s.Modifier = &v
	return b
}

func (b *PolicyDescResponseBodyBuilder) TaskDescInfoList(v []PolicyDescResponseTaskDescInfoList) *PolicyDescResponseBodyBuilder {
	b.s.TaskDescInfoList = &v
	return b
}

func (b *PolicyDescResponseBodyBuilder) PolicyUid(v string) *PolicyDescResponseBodyBuilder {
	b.s.PolicyUid = &v
	return b
}

func (b *PolicyDescResponseBodyBuilder) PolicyDesc(v string) *PolicyDescResponseBodyBuilder {
	b.s.PolicyDesc = &v
	return b
}

func (b *PolicyDescResponseBodyBuilder) TenantId(v string) *PolicyDescResponseBodyBuilder {
	b.s.TenantId = &v
	return b
}

func (b *PolicyDescResponseBodyBuilder) CreatedTime(v string) *PolicyDescResponseBodyBuilder {
	b.s.CreatedTime = &v
	return b
}

func (b *PolicyDescResponseBodyBuilder) MopUserName(v string) *PolicyDescResponseBodyBuilder {
	b.s.MopUserName = &v
	return b
}

func (b *PolicyDescResponseBodyBuilder) PoolCode(v string) *PolicyDescResponseBodyBuilder {
	b.s.PoolCode = &v
	return b
}

func (b *PolicyDescResponseBodyBuilder) MopUserId(v string) *PolicyDescResponseBodyBuilder {
	b.s.MopUserId = &v
	return b
}

func (b *PolicyDescResponseBodyBuilder) PoolName(v string) *PolicyDescResponseBodyBuilder {
	b.s.PoolName = &v
	return b
}

func (b *PolicyDescResponseBodyBuilder) Status(v string) *PolicyDescResponseBodyBuilder {
	b.s.Status = &v
	return b
}

func (b *PolicyDescResponseBodyBuilder) Build() *PolicyDescResponseBody {
	return b.s
}


