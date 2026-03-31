// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetPolicyInfoByInstanceIdResponseBody struct {

    // 修改时间
	ModifiedTime *string `json:"modifiedTime,omitempty"`
    // 创建人
	Creator *string `json:"creator,omitempty"`
    // 策略名称
	PolicyName *string `json:"policyName,omitempty"`
    // 修改人
	Modifier *string `json:"modifier,omitempty"`
    // 任务详情
	TaskDescInfoList *[]GetPolicyInfoByInstanceIdResponseTaskDescInfoList `json:"taskDescInfoList,omitempty"`
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

func (s GetPolicyInfoByInstanceIdResponseBody) String() string {
	return utils.Beautify(s)
}

func (s GetPolicyInfoByInstanceIdResponseBody) GoString() string {
	return s.String()
}

func (s GetPolicyInfoByInstanceIdResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetPolicyInfoByInstanceIdResponseBody) SetModifiedTime(v string) *GetPolicyInfoByInstanceIdResponseBody {
	s.ModifiedTime = &v
	return s
}

func (s *GetPolicyInfoByInstanceIdResponseBody) SetCreator(v string) *GetPolicyInfoByInstanceIdResponseBody {
	s.Creator = &v
	return s
}

func (s *GetPolicyInfoByInstanceIdResponseBody) SetPolicyName(v string) *GetPolicyInfoByInstanceIdResponseBody {
	s.PolicyName = &v
	return s
}

func (s *GetPolicyInfoByInstanceIdResponseBody) SetModifier(v string) *GetPolicyInfoByInstanceIdResponseBody {
	s.Modifier = &v
	return s
}

func (s *GetPolicyInfoByInstanceIdResponseBody) SetTaskDescInfoList(v []GetPolicyInfoByInstanceIdResponseTaskDescInfoList) *GetPolicyInfoByInstanceIdResponseBody {
	s.TaskDescInfoList = &v
	return s
}

func (s *GetPolicyInfoByInstanceIdResponseBody) SetPolicyUid(v string) *GetPolicyInfoByInstanceIdResponseBody {
	s.PolicyUid = &v
	return s
}

func (s *GetPolicyInfoByInstanceIdResponseBody) SetPolicyDesc(v string) *GetPolicyInfoByInstanceIdResponseBody {
	s.PolicyDesc = &v
	return s
}

func (s *GetPolicyInfoByInstanceIdResponseBody) SetTenantId(v string) *GetPolicyInfoByInstanceIdResponseBody {
	s.TenantId = &v
	return s
}

func (s *GetPolicyInfoByInstanceIdResponseBody) SetCreatedTime(v string) *GetPolicyInfoByInstanceIdResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *GetPolicyInfoByInstanceIdResponseBody) SetMopUserName(v string) *GetPolicyInfoByInstanceIdResponseBody {
	s.MopUserName = &v
	return s
}

func (s *GetPolicyInfoByInstanceIdResponseBody) SetPoolCode(v string) *GetPolicyInfoByInstanceIdResponseBody {
	s.PoolCode = &v
	return s
}

func (s *GetPolicyInfoByInstanceIdResponseBody) SetMopUserId(v string) *GetPolicyInfoByInstanceIdResponseBody {
	s.MopUserId = &v
	return s
}

func (s *GetPolicyInfoByInstanceIdResponseBody) SetPoolName(v string) *GetPolicyInfoByInstanceIdResponseBody {
	s.PoolName = &v
	return s
}

func (s *GetPolicyInfoByInstanceIdResponseBody) SetStatus(v string) *GetPolicyInfoByInstanceIdResponseBody {
	s.Status = &v
	return s
}


type GetPolicyInfoByInstanceIdResponseBodyBuilder struct {
	s *GetPolicyInfoByInstanceIdResponseBody
}

func NewGetPolicyInfoByInstanceIdResponseBodyBuilder() *GetPolicyInfoByInstanceIdResponseBodyBuilder {
	s := &GetPolicyInfoByInstanceIdResponseBody{}
	b := &GetPolicyInfoByInstanceIdResponseBodyBuilder{s: s}
	return b
}

func (b *GetPolicyInfoByInstanceIdResponseBodyBuilder) ModifiedTime(v string) *GetPolicyInfoByInstanceIdResponseBodyBuilder {
	b.s.ModifiedTime = &v
	return b
}

func (b *GetPolicyInfoByInstanceIdResponseBodyBuilder) Creator(v string) *GetPolicyInfoByInstanceIdResponseBodyBuilder {
	b.s.Creator = &v
	return b
}

func (b *GetPolicyInfoByInstanceIdResponseBodyBuilder) PolicyName(v string) *GetPolicyInfoByInstanceIdResponseBodyBuilder {
	b.s.PolicyName = &v
	return b
}

func (b *GetPolicyInfoByInstanceIdResponseBodyBuilder) Modifier(v string) *GetPolicyInfoByInstanceIdResponseBodyBuilder {
	b.s.Modifier = &v
	return b
}

func (b *GetPolicyInfoByInstanceIdResponseBodyBuilder) TaskDescInfoList(v []GetPolicyInfoByInstanceIdResponseTaskDescInfoList) *GetPolicyInfoByInstanceIdResponseBodyBuilder {
	b.s.TaskDescInfoList = &v
	return b
}

func (b *GetPolicyInfoByInstanceIdResponseBodyBuilder) PolicyUid(v string) *GetPolicyInfoByInstanceIdResponseBodyBuilder {
	b.s.PolicyUid = &v
	return b
}

func (b *GetPolicyInfoByInstanceIdResponseBodyBuilder) PolicyDesc(v string) *GetPolicyInfoByInstanceIdResponseBodyBuilder {
	b.s.PolicyDesc = &v
	return b
}

func (b *GetPolicyInfoByInstanceIdResponseBodyBuilder) TenantId(v string) *GetPolicyInfoByInstanceIdResponseBodyBuilder {
	b.s.TenantId = &v
	return b
}

func (b *GetPolicyInfoByInstanceIdResponseBodyBuilder) CreatedTime(v string) *GetPolicyInfoByInstanceIdResponseBodyBuilder {
	b.s.CreatedTime = &v
	return b
}

func (b *GetPolicyInfoByInstanceIdResponseBodyBuilder) MopUserName(v string) *GetPolicyInfoByInstanceIdResponseBodyBuilder {
	b.s.MopUserName = &v
	return b
}

func (b *GetPolicyInfoByInstanceIdResponseBodyBuilder) PoolCode(v string) *GetPolicyInfoByInstanceIdResponseBodyBuilder {
	b.s.PoolCode = &v
	return b
}

func (b *GetPolicyInfoByInstanceIdResponseBodyBuilder) MopUserId(v string) *GetPolicyInfoByInstanceIdResponseBodyBuilder {
	b.s.MopUserId = &v
	return b
}

func (b *GetPolicyInfoByInstanceIdResponseBodyBuilder) PoolName(v string) *GetPolicyInfoByInstanceIdResponseBodyBuilder {
	b.s.PoolName = &v
	return b
}

func (b *GetPolicyInfoByInstanceIdResponseBodyBuilder) Status(v string) *GetPolicyInfoByInstanceIdResponseBodyBuilder {
	b.s.Status = &v
	return b
}

func (b *GetPolicyInfoByInstanceIdResponseBodyBuilder) Build() *GetPolicyInfoByInstanceIdResponseBody {
	return b.s
}


