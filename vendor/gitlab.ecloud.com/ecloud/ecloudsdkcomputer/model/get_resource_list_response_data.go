// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetResourceListResponseData struct {

    // VPC
	VpcName *string `json:"vpcName,omitempty"`
    // 生效时间
	EffectiveTime *string `json:"effectiveTime,omitempty"`
    // 云电脑名称
	MachineName *string `json:"machineName,omitempty"`
    // 订购资源ID
	InstanceId *string `json:"instanceId,omitempty"`
    // 订购渠道
	OrderSourceCode *string `json:"orderSourceCode,omitempty"`
    // 创建时间
	CreatedTime *string `json:"createdTime,omitempty"`
    // 系统盘
	SysDiskSize *string `json:"sysDiskSize,omitempty"`
    // 规格code
	ResourceTypeCode *string `json:"resourceTypeCode,omitempty"`
    // 资源池名称
	PoolName *string `json:"poolName,omitempty"`
    // 运行状态（available可用、shutdown关机、onShutdown关机中、onAvailable开机中、onReload重装系统中、onRestart重启中、exception异常、imageCreating镜像制作中）
	MachineStatus *string `json:"machineStatus,omitempty"`
    // 电脑协议（V2:自研CMSSZTE，V1.1：新华三，V1.2：中兴，V 1.4：中兴大云融合）
	ComputerProtocol *string `json:"computerProtocol,omitempty"`
    // 所属组织
	OrganizationName *string `json:"organizationName,omitempty"`
    // 用户名
	Uname *string `json:"uname,omitempty"`
    // 策略名称
	PolicyName *string `json:"policyName,omitempty"`
    // 登录账号
	UserName *string `json:"userName,omitempty"`
    // 数据盘大小
	ExtDataDiskSize *int32 `json:"extDataDiskSize,omitempty"`
    // 云电脑主机id
	MachineId *string `json:"machineId,omitempty"`
    // 策略ID
	PolicyId *string `json:"policyId,omitempty"`
    // 所在子网的带宽大小
	BandwidthMax *int32 `json:"bandwidthMax,omitempty"`
    // 资源池id
	PoolId *string `json:"poolId,omitempty"`
    // 租户表主键id
	TenantId *string `json:"tenantId,omitempty"`
    // mop用户名
	MopUserName *string `json:"mopUserName,omitempty"`
    // 资源池编码
	PoolCode *string `json:"poolCode,omitempty"`
    // mop用户id
	MopUserId *string `json:"mopUserId,omitempty"`
    // 可用区
	Region *string `json:"region,omitempty"`
    // 规格名称
	ResourceTypeName *string `json:"resourceTypeName,omitempty"`
    // 数据中心/网络工作区
	Dc *string `json:"dc,omitempty"`
    // 内网IP
	IntranetIP *string `json:"intranetIP,omitempty"`
}

func (s GetResourceListResponseData) String() string {
	return utils.Beautify(s)
}

func (s GetResourceListResponseData) GoString() string {
	return s.String()
}

func (s GetResourceListResponseData) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetResourceListResponseData) SetVpcName(v string) *GetResourceListResponseData {
	s.VpcName = &v
	return s
}

func (s *GetResourceListResponseData) SetEffectiveTime(v string) *GetResourceListResponseData {
	s.EffectiveTime = &v
	return s
}

func (s *GetResourceListResponseData) SetMachineName(v string) *GetResourceListResponseData {
	s.MachineName = &v
	return s
}

func (s *GetResourceListResponseData) SetInstanceId(v string) *GetResourceListResponseData {
	s.InstanceId = &v
	return s
}

func (s *GetResourceListResponseData) SetOrderSourceCode(v string) *GetResourceListResponseData {
	s.OrderSourceCode = &v
	return s
}

func (s *GetResourceListResponseData) SetCreatedTime(v string) *GetResourceListResponseData {
	s.CreatedTime = &v
	return s
}

func (s *GetResourceListResponseData) SetSysDiskSize(v string) *GetResourceListResponseData {
	s.SysDiskSize = &v
	return s
}

func (s *GetResourceListResponseData) SetResourceTypeCode(v string) *GetResourceListResponseData {
	s.ResourceTypeCode = &v
	return s
}

func (s *GetResourceListResponseData) SetPoolName(v string) *GetResourceListResponseData {
	s.PoolName = &v
	return s
}

func (s *GetResourceListResponseData) SetMachineStatus(v string) *GetResourceListResponseData {
	s.MachineStatus = &v
	return s
}

func (s *GetResourceListResponseData) SetComputerProtocol(v string) *GetResourceListResponseData {
	s.ComputerProtocol = &v
	return s
}

func (s *GetResourceListResponseData) SetOrganizationName(v string) *GetResourceListResponseData {
	s.OrganizationName = &v
	return s
}

func (s *GetResourceListResponseData) SetUname(v string) *GetResourceListResponseData {
	s.Uname = &v
	return s
}

func (s *GetResourceListResponseData) SetPolicyName(v string) *GetResourceListResponseData {
	s.PolicyName = &v
	return s
}

func (s *GetResourceListResponseData) SetUserName(v string) *GetResourceListResponseData {
	s.UserName = &v
	return s
}

func (s *GetResourceListResponseData) SetExtDataDiskSize(v int32) *GetResourceListResponseData {
	s.ExtDataDiskSize = &v
	return s
}

func (s *GetResourceListResponseData) SetMachineId(v string) *GetResourceListResponseData {
	s.MachineId = &v
	return s
}

func (s *GetResourceListResponseData) SetPolicyId(v string) *GetResourceListResponseData {
	s.PolicyId = &v
	return s
}

func (s *GetResourceListResponseData) SetBandwidthMax(v int32) *GetResourceListResponseData {
	s.BandwidthMax = &v
	return s
}

func (s *GetResourceListResponseData) SetPoolId(v string) *GetResourceListResponseData {
	s.PoolId = &v
	return s
}

func (s *GetResourceListResponseData) SetTenantId(v string) *GetResourceListResponseData {
	s.TenantId = &v
	return s
}

func (s *GetResourceListResponseData) SetMopUserName(v string) *GetResourceListResponseData {
	s.MopUserName = &v
	return s
}

func (s *GetResourceListResponseData) SetPoolCode(v string) *GetResourceListResponseData {
	s.PoolCode = &v
	return s
}

func (s *GetResourceListResponseData) SetMopUserId(v string) *GetResourceListResponseData {
	s.MopUserId = &v
	return s
}

func (s *GetResourceListResponseData) SetRegion(v string) *GetResourceListResponseData {
	s.Region = &v
	return s
}

func (s *GetResourceListResponseData) SetResourceTypeName(v string) *GetResourceListResponseData {
	s.ResourceTypeName = &v
	return s
}

func (s *GetResourceListResponseData) SetDc(v string) *GetResourceListResponseData {
	s.Dc = &v
	return s
}

func (s *GetResourceListResponseData) SetIntranetIP(v string) *GetResourceListResponseData {
	s.IntranetIP = &v
	return s
}


type GetResourceListResponseDataBuilder struct {
	s *GetResourceListResponseData
}

func NewGetResourceListResponseDataBuilder() *GetResourceListResponseDataBuilder {
	s := &GetResourceListResponseData{}
	b := &GetResourceListResponseDataBuilder{s: s}
	return b
}

func (b *GetResourceListResponseDataBuilder) VpcName(v string) *GetResourceListResponseDataBuilder {
	b.s.VpcName = &v
	return b
}

func (b *GetResourceListResponseDataBuilder) EffectiveTime(v string) *GetResourceListResponseDataBuilder {
	b.s.EffectiveTime = &v
	return b
}

func (b *GetResourceListResponseDataBuilder) MachineName(v string) *GetResourceListResponseDataBuilder {
	b.s.MachineName = &v
	return b
}

func (b *GetResourceListResponseDataBuilder) InstanceId(v string) *GetResourceListResponseDataBuilder {
	b.s.InstanceId = &v
	return b
}

func (b *GetResourceListResponseDataBuilder) OrderSourceCode(v string) *GetResourceListResponseDataBuilder {
	b.s.OrderSourceCode = &v
	return b
}

func (b *GetResourceListResponseDataBuilder) CreatedTime(v string) *GetResourceListResponseDataBuilder {
	b.s.CreatedTime = &v
	return b
}

func (b *GetResourceListResponseDataBuilder) SysDiskSize(v string) *GetResourceListResponseDataBuilder {
	b.s.SysDiskSize = &v
	return b
}

func (b *GetResourceListResponseDataBuilder) ResourceTypeCode(v string) *GetResourceListResponseDataBuilder {
	b.s.ResourceTypeCode = &v
	return b
}

func (b *GetResourceListResponseDataBuilder) PoolName(v string) *GetResourceListResponseDataBuilder {
	b.s.PoolName = &v
	return b
}

func (b *GetResourceListResponseDataBuilder) MachineStatus(v string) *GetResourceListResponseDataBuilder {
	b.s.MachineStatus = &v
	return b
}

func (b *GetResourceListResponseDataBuilder) ComputerProtocol(v string) *GetResourceListResponseDataBuilder {
	b.s.ComputerProtocol = &v
	return b
}

func (b *GetResourceListResponseDataBuilder) OrganizationName(v string) *GetResourceListResponseDataBuilder {
	b.s.OrganizationName = &v
	return b
}

func (b *GetResourceListResponseDataBuilder) Uname(v string) *GetResourceListResponseDataBuilder {
	b.s.Uname = &v
	return b
}

func (b *GetResourceListResponseDataBuilder) PolicyName(v string) *GetResourceListResponseDataBuilder {
	b.s.PolicyName = &v
	return b
}

func (b *GetResourceListResponseDataBuilder) UserName(v string) *GetResourceListResponseDataBuilder {
	b.s.UserName = &v
	return b
}

func (b *GetResourceListResponseDataBuilder) ExtDataDiskSize(v int32) *GetResourceListResponseDataBuilder {
	b.s.ExtDataDiskSize = &v
	return b
}

func (b *GetResourceListResponseDataBuilder) MachineId(v string) *GetResourceListResponseDataBuilder {
	b.s.MachineId = &v
	return b
}

func (b *GetResourceListResponseDataBuilder) PolicyId(v string) *GetResourceListResponseDataBuilder {
	b.s.PolicyId = &v
	return b
}

func (b *GetResourceListResponseDataBuilder) BandwidthMax(v int32) *GetResourceListResponseDataBuilder {
	b.s.BandwidthMax = &v
	return b
}

func (b *GetResourceListResponseDataBuilder) PoolId(v string) *GetResourceListResponseDataBuilder {
	b.s.PoolId = &v
	return b
}

func (b *GetResourceListResponseDataBuilder) TenantId(v string) *GetResourceListResponseDataBuilder {
	b.s.TenantId = &v
	return b
}

func (b *GetResourceListResponseDataBuilder) MopUserName(v string) *GetResourceListResponseDataBuilder {
	b.s.MopUserName = &v
	return b
}

func (b *GetResourceListResponseDataBuilder) PoolCode(v string) *GetResourceListResponseDataBuilder {
	b.s.PoolCode = &v
	return b
}

func (b *GetResourceListResponseDataBuilder) MopUserId(v string) *GetResourceListResponseDataBuilder {
	b.s.MopUserId = &v
	return b
}

func (b *GetResourceListResponseDataBuilder) Region(v string) *GetResourceListResponseDataBuilder {
	b.s.Region = &v
	return b
}

func (b *GetResourceListResponseDataBuilder) ResourceTypeName(v string) *GetResourceListResponseDataBuilder {
	b.s.ResourceTypeName = &v
	return b
}

func (b *GetResourceListResponseDataBuilder) Dc(v string) *GetResourceListResponseDataBuilder {
	b.s.Dc = &v
	return b
}

func (b *GetResourceListResponseDataBuilder) IntranetIP(v string) *GetResourceListResponseDataBuilder {
	b.s.IntranetIP = &v
	return b
}

func (b *GetResourceListResponseDataBuilder) Build() *GetResourceListResponseData {
	return b.s
}


