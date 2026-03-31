// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetResourceListBody struct {
    position.Body
    // 运行状态（available可用、shutdown关机、onShutdown关机中、onAvailable开机中、onReload重装系统中、onRestart重启中、exception异常）
	MachineStatus *string `json:"machineStatus,omitempty"`
    // 订购资源ID
	InstanceId *string `json:"instanceId,omitempty"`
    // 云电脑主机id
	MachineId *string `json:"machineId,omitempty"`
    // 策略ID
	PolicyId *string `json:"policyId,omitempty"`
    // 策略名称
	PolicyName *string `json:"policyName,omitempty"`
    // 资源池id
	PoolId *string `json:"poolId,omitempty"`
    // 每页显示条数
	PageSize *int32 `json:"pageSize,omitempty"`
    // 当前页
	Page *int32 `json:"page,omitempty"`
    // 登录账号
	UserName *string `json:"userName,omitempty"`
    // 云电脑名称
	MachineName *string `json:"machineName,omitempty"`
    // 规格名称
	ResourceTypeName *string `json:"resourceTypeName,omitempty"`
}

func (s GetResourceListBody) String() string {
	return utils.Beautify(s)
}

func (s GetResourceListBody) GoString() string {
	return s.String()
}

func (s GetResourceListBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetResourceListBody) SetMachineStatus(v string) *GetResourceListBody {
	s.MachineStatus = &v
	return s
}

func (s *GetResourceListBody) SetInstanceId(v string) *GetResourceListBody {
	s.InstanceId = &v
	return s
}

func (s *GetResourceListBody) SetMachineId(v string) *GetResourceListBody {
	s.MachineId = &v
	return s
}

func (s *GetResourceListBody) SetPolicyId(v string) *GetResourceListBody {
	s.PolicyId = &v
	return s
}

func (s *GetResourceListBody) SetPolicyName(v string) *GetResourceListBody {
	s.PolicyName = &v
	return s
}

func (s *GetResourceListBody) SetPoolId(v string) *GetResourceListBody {
	s.PoolId = &v
	return s
}

func (s *GetResourceListBody) SetPageSize(v int32) *GetResourceListBody {
	s.PageSize = &v
	return s
}

func (s *GetResourceListBody) SetPage(v int32) *GetResourceListBody {
	s.Page = &v
	return s
}

func (s *GetResourceListBody) SetUserName(v string) *GetResourceListBody {
	s.UserName = &v
	return s
}

func (s *GetResourceListBody) SetMachineName(v string) *GetResourceListBody {
	s.MachineName = &v
	return s
}

func (s *GetResourceListBody) SetResourceTypeName(v string) *GetResourceListBody {
	s.ResourceTypeName = &v
	return s
}


type GetResourceListBodyBuilder struct {
	s *GetResourceListBody
}

func NewGetResourceListBodyBuilder() *GetResourceListBodyBuilder {
	s := &GetResourceListBody{}
	b := &GetResourceListBodyBuilder{s: s}
	return b
}

func (b *GetResourceListBodyBuilder) MachineStatus(v string) *GetResourceListBodyBuilder {
	b.s.MachineStatus = &v
	return b
}

func (b *GetResourceListBodyBuilder) InstanceId(v string) *GetResourceListBodyBuilder {
	b.s.InstanceId = &v
	return b
}

func (b *GetResourceListBodyBuilder) MachineId(v string) *GetResourceListBodyBuilder {
	b.s.MachineId = &v
	return b
}

func (b *GetResourceListBodyBuilder) PolicyId(v string) *GetResourceListBodyBuilder {
	b.s.PolicyId = &v
	return b
}

func (b *GetResourceListBodyBuilder) PolicyName(v string) *GetResourceListBodyBuilder {
	b.s.PolicyName = &v
	return b
}

func (b *GetResourceListBodyBuilder) PoolId(v string) *GetResourceListBodyBuilder {
	b.s.PoolId = &v
	return b
}

func (b *GetResourceListBodyBuilder) PageSize(v int32) *GetResourceListBodyBuilder {
	b.s.PageSize = &v
	return b
}

func (b *GetResourceListBodyBuilder) Page(v int32) *GetResourceListBodyBuilder {
	b.s.Page = &v
	return b
}

func (b *GetResourceListBodyBuilder) UserName(v string) *GetResourceListBodyBuilder {
	b.s.UserName = &v
	return b
}

func (b *GetResourceListBodyBuilder) MachineName(v string) *GetResourceListBodyBuilder {
	b.s.MachineName = &v
	return b
}

func (b *GetResourceListBodyBuilder) ResourceTypeName(v string) *GetResourceListBodyBuilder {
	b.s.ResourceTypeName = &v
	return b
}

func (b *GetResourceListBodyBuilder) Build() *GetResourceListBody {
	return b.s
}


