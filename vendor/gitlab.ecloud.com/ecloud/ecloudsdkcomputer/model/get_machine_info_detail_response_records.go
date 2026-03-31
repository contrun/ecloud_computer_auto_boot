// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetMachineInfoDetailResponseRecords struct {

    // vpc名称
	VpcName *string `json:"vpcName,omitempty"`
    // Qos名称
	QosName *string `json:"qosName,omitempty"`
    // ad用户与桌面操作的状态，0：已绑定、1:绑定中、2：绑定失败、3：换绑中、4：换绑失败、5：解绑中、6:解绑失败
	AdStatus *string `json:"adStatus,omitempty"`
    // 生效时间
	EffectiveTime *string `json:"effectiveTime,omitempty"`
    // 企业名称
	CompanyName *string `json:"companyName,omitempty"`
    // 可用区名称
	RegionName *string `json:"regionName,omitempty"`
    // 是否订购网络服务包（0：未订购，1：已订购）
	NetService *bool `json:"netService,omitempty"`
    // 是否分配，分配0，未分配1
	IsActive *string `json:"isActive,omitempty"`
    // 主机名
	MachineName *string `json:"machineName,omitempty"`
    // 企业门户组织ID
	OrgId *string `json:"orgId,omitempty"`
    // 资源实例ID
	InstanceId *string `json:"instanceId,omitempty"`
    // 操作系统
	OsVersion *string `json:"osVersion,omitempty"`
    // 操作系统
	OperateSystem *string `json:"operateSystem,omitempty"`
    // 渠道
	OrderSourceCode *string `json:"orderSourceCode,omitempty"`
    // 创建时间
	CreatedTime *string `json:"createdTime,omitempty"`
    // 可用区编码
	AzCode *string `json:"azCode,omitempty"`
    // 系统盘大小
	SysDiskSize *string `json:"sysDiskSize,omitempty"`
    // 资源管理状态（0：正常，1：锁定，2：禁用，3：锁定且禁用）
	ResourceManageState *string `json:"resourceManageState,omitempty"`
    // 资源池名称
	PoolName *string `json:"poolName,omitempty"`
    // 厂商
	CompanyCode *string `json:"companyCode,omitempty"`
    // 安全组名称
	SecurityGroupName *string `json:"securityGroupName,omitempty"`
    // 可用状态（available开机、exception异常、shutdown关机、onShutdown关机中、onAvailable开机中、onReload重装系统中、onRestart重启中、transferring 迁移中/漫游中）
	MachineStatus *string `json:"machineStatus,omitempty"`
    // 渠道名称
	OrderSource *string `json:"orderSource,omitempty"`
    // 企业门户组织名称
	OrgName *string `json:"orgName,omitempty"`
    // 快照ID
	SnapshotId *string `json:"snapshotId,omitempty"`
    // 电脑协议（V2:自研CMSSZTE，V1.1：新华三，V1.2：中兴，V 1.4：中兴大云融合）
	ComputerProtocol *string `json:"computerProtocol,omitempty"`
    // 分配状态（free空闲、occupied已分配、releasing待释放、freeze冻结、released已释放）
	DistributeStatus *string `json:"distributeStatus,omitempty"`
    // 电脑策略名称
	PolicyName *string `json:"policyName,omitempty"`
    // 登录用户名
	UserName *string `json:"userName,omitempty"`
    // 登录用户编号
	UserId *string `json:"userId,omitempty"`
    // 失效时间
	ExpireTime *string `json:"expireTime,omitempty"`
    // 数据盘大小
	ExtDataDiskSize *int64 `json:"extDataDiskSize,omitempty"`
    // 电脑ID
	MachineId *string `json:"machineId,omitempty"`
    // 电脑策略id
	PolicyId *string `json:"policyId,omitempty"`
    // 最大带宽
	BandwidthMax *int32 `json:"bandwidthMax,omitempty"`
    // 租户ID
	TenantId *string `json:"tenantId,omitempty"`
    // mop用户名
	MopUserName *string `json:"mopUserName,omitempty"`
    // 资源池编码
	PoolCode *string `json:"poolCode,omitempty"`
    // mop用户ID
	MopUserId *string `json:"mopUserId,omitempty"`
    // 可用区编码
	Region *string `json:"region,omitempty"`
    // 定制规格参数
	ResourceTypeName *string `json:"resourceTypeName,omitempty"`
    // 数据中心
	Dc *string `json:"dc,omitempty"`
    // 内网ip地址
	IntranetIP *string `json:"intranetIP,omitempty"`
}

func (s GetMachineInfoDetailResponseRecords) String() string {
	return utils.Beautify(s)
}

func (s GetMachineInfoDetailResponseRecords) GoString() string {
	return s.String()
}

func (s GetMachineInfoDetailResponseRecords) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetMachineInfoDetailResponseRecords) SetVpcName(v string) *GetMachineInfoDetailResponseRecords {
	s.VpcName = &v
	return s
}

func (s *GetMachineInfoDetailResponseRecords) SetQosName(v string) *GetMachineInfoDetailResponseRecords {
	s.QosName = &v
	return s
}

func (s *GetMachineInfoDetailResponseRecords) SetAdStatus(v string) *GetMachineInfoDetailResponseRecords {
	s.AdStatus = &v
	return s
}

func (s *GetMachineInfoDetailResponseRecords) SetEffectiveTime(v string) *GetMachineInfoDetailResponseRecords {
	s.EffectiveTime = &v
	return s
}

func (s *GetMachineInfoDetailResponseRecords) SetCompanyName(v string) *GetMachineInfoDetailResponseRecords {
	s.CompanyName = &v
	return s
}

func (s *GetMachineInfoDetailResponseRecords) SetRegionName(v string) *GetMachineInfoDetailResponseRecords {
	s.RegionName = &v
	return s
}

func (s *GetMachineInfoDetailResponseRecords) SetNetService(v bool) *GetMachineInfoDetailResponseRecords {
	s.NetService = &v
	return s
}

func (s *GetMachineInfoDetailResponseRecords) SetIsActive(v string) *GetMachineInfoDetailResponseRecords {
	s.IsActive = &v
	return s
}

func (s *GetMachineInfoDetailResponseRecords) SetMachineName(v string) *GetMachineInfoDetailResponseRecords {
	s.MachineName = &v
	return s
}

func (s *GetMachineInfoDetailResponseRecords) SetOrgId(v string) *GetMachineInfoDetailResponseRecords {
	s.OrgId = &v
	return s
}

func (s *GetMachineInfoDetailResponseRecords) SetInstanceId(v string) *GetMachineInfoDetailResponseRecords {
	s.InstanceId = &v
	return s
}

func (s *GetMachineInfoDetailResponseRecords) SetOsVersion(v string) *GetMachineInfoDetailResponseRecords {
	s.OsVersion = &v
	return s
}

func (s *GetMachineInfoDetailResponseRecords) SetOperateSystem(v string) *GetMachineInfoDetailResponseRecords {
	s.OperateSystem = &v
	return s
}

func (s *GetMachineInfoDetailResponseRecords) SetOrderSourceCode(v string) *GetMachineInfoDetailResponseRecords {
	s.OrderSourceCode = &v
	return s
}

func (s *GetMachineInfoDetailResponseRecords) SetCreatedTime(v string) *GetMachineInfoDetailResponseRecords {
	s.CreatedTime = &v
	return s
}

func (s *GetMachineInfoDetailResponseRecords) SetAzCode(v string) *GetMachineInfoDetailResponseRecords {
	s.AzCode = &v
	return s
}

func (s *GetMachineInfoDetailResponseRecords) SetSysDiskSize(v string) *GetMachineInfoDetailResponseRecords {
	s.SysDiskSize = &v
	return s
}

func (s *GetMachineInfoDetailResponseRecords) SetResourceManageState(v string) *GetMachineInfoDetailResponseRecords {
	s.ResourceManageState = &v
	return s
}

func (s *GetMachineInfoDetailResponseRecords) SetPoolName(v string) *GetMachineInfoDetailResponseRecords {
	s.PoolName = &v
	return s
}

func (s *GetMachineInfoDetailResponseRecords) SetCompanyCode(v string) *GetMachineInfoDetailResponseRecords {
	s.CompanyCode = &v
	return s
}

func (s *GetMachineInfoDetailResponseRecords) SetSecurityGroupName(v string) *GetMachineInfoDetailResponseRecords {
	s.SecurityGroupName = &v
	return s
}

func (s *GetMachineInfoDetailResponseRecords) SetMachineStatus(v string) *GetMachineInfoDetailResponseRecords {
	s.MachineStatus = &v
	return s
}

func (s *GetMachineInfoDetailResponseRecords) SetOrderSource(v string) *GetMachineInfoDetailResponseRecords {
	s.OrderSource = &v
	return s
}

func (s *GetMachineInfoDetailResponseRecords) SetOrgName(v string) *GetMachineInfoDetailResponseRecords {
	s.OrgName = &v
	return s
}

func (s *GetMachineInfoDetailResponseRecords) SetSnapshotId(v string) *GetMachineInfoDetailResponseRecords {
	s.SnapshotId = &v
	return s
}

func (s *GetMachineInfoDetailResponseRecords) SetComputerProtocol(v string) *GetMachineInfoDetailResponseRecords {
	s.ComputerProtocol = &v
	return s
}

func (s *GetMachineInfoDetailResponseRecords) SetDistributeStatus(v string) *GetMachineInfoDetailResponseRecords {
	s.DistributeStatus = &v
	return s
}

func (s *GetMachineInfoDetailResponseRecords) SetPolicyName(v string) *GetMachineInfoDetailResponseRecords {
	s.PolicyName = &v
	return s
}

func (s *GetMachineInfoDetailResponseRecords) SetUserName(v string) *GetMachineInfoDetailResponseRecords {
	s.UserName = &v
	return s
}

func (s *GetMachineInfoDetailResponseRecords) SetUserId(v string) *GetMachineInfoDetailResponseRecords {
	s.UserId = &v
	return s
}

func (s *GetMachineInfoDetailResponseRecords) SetExpireTime(v string) *GetMachineInfoDetailResponseRecords {
	s.ExpireTime = &v
	return s
}

func (s *GetMachineInfoDetailResponseRecords) SetExtDataDiskSize(v int64) *GetMachineInfoDetailResponseRecords {
	s.ExtDataDiskSize = &v
	return s
}

func (s *GetMachineInfoDetailResponseRecords) SetMachineId(v string) *GetMachineInfoDetailResponseRecords {
	s.MachineId = &v
	return s
}

func (s *GetMachineInfoDetailResponseRecords) SetPolicyId(v string) *GetMachineInfoDetailResponseRecords {
	s.PolicyId = &v
	return s
}

func (s *GetMachineInfoDetailResponseRecords) SetBandwidthMax(v int32) *GetMachineInfoDetailResponseRecords {
	s.BandwidthMax = &v
	return s
}

func (s *GetMachineInfoDetailResponseRecords) SetTenantId(v string) *GetMachineInfoDetailResponseRecords {
	s.TenantId = &v
	return s
}

func (s *GetMachineInfoDetailResponseRecords) SetMopUserName(v string) *GetMachineInfoDetailResponseRecords {
	s.MopUserName = &v
	return s
}

func (s *GetMachineInfoDetailResponseRecords) SetPoolCode(v string) *GetMachineInfoDetailResponseRecords {
	s.PoolCode = &v
	return s
}

func (s *GetMachineInfoDetailResponseRecords) SetMopUserId(v string) *GetMachineInfoDetailResponseRecords {
	s.MopUserId = &v
	return s
}

func (s *GetMachineInfoDetailResponseRecords) SetRegion(v string) *GetMachineInfoDetailResponseRecords {
	s.Region = &v
	return s
}

func (s *GetMachineInfoDetailResponseRecords) SetResourceTypeName(v string) *GetMachineInfoDetailResponseRecords {
	s.ResourceTypeName = &v
	return s
}

func (s *GetMachineInfoDetailResponseRecords) SetDc(v string) *GetMachineInfoDetailResponseRecords {
	s.Dc = &v
	return s
}

func (s *GetMachineInfoDetailResponseRecords) SetIntranetIP(v string) *GetMachineInfoDetailResponseRecords {
	s.IntranetIP = &v
	return s
}


type GetMachineInfoDetailResponseRecordsBuilder struct {
	s *GetMachineInfoDetailResponseRecords
}

func NewGetMachineInfoDetailResponseRecordsBuilder() *GetMachineInfoDetailResponseRecordsBuilder {
	s := &GetMachineInfoDetailResponseRecords{}
	b := &GetMachineInfoDetailResponseRecordsBuilder{s: s}
	return b
}

func (b *GetMachineInfoDetailResponseRecordsBuilder) VpcName(v string) *GetMachineInfoDetailResponseRecordsBuilder {
	b.s.VpcName = &v
	return b
}

func (b *GetMachineInfoDetailResponseRecordsBuilder) QosName(v string) *GetMachineInfoDetailResponseRecordsBuilder {
	b.s.QosName = &v
	return b
}

func (b *GetMachineInfoDetailResponseRecordsBuilder) AdStatus(v string) *GetMachineInfoDetailResponseRecordsBuilder {
	b.s.AdStatus = &v
	return b
}

func (b *GetMachineInfoDetailResponseRecordsBuilder) EffectiveTime(v string) *GetMachineInfoDetailResponseRecordsBuilder {
	b.s.EffectiveTime = &v
	return b
}

func (b *GetMachineInfoDetailResponseRecordsBuilder) CompanyName(v string) *GetMachineInfoDetailResponseRecordsBuilder {
	b.s.CompanyName = &v
	return b
}

func (b *GetMachineInfoDetailResponseRecordsBuilder) RegionName(v string) *GetMachineInfoDetailResponseRecordsBuilder {
	b.s.RegionName = &v
	return b
}

func (b *GetMachineInfoDetailResponseRecordsBuilder) NetService(v bool) *GetMachineInfoDetailResponseRecordsBuilder {
	b.s.NetService = &v
	return b
}

func (b *GetMachineInfoDetailResponseRecordsBuilder) IsActive(v string) *GetMachineInfoDetailResponseRecordsBuilder {
	b.s.IsActive = &v
	return b
}

func (b *GetMachineInfoDetailResponseRecordsBuilder) MachineName(v string) *GetMachineInfoDetailResponseRecordsBuilder {
	b.s.MachineName = &v
	return b
}

func (b *GetMachineInfoDetailResponseRecordsBuilder) OrgId(v string) *GetMachineInfoDetailResponseRecordsBuilder {
	b.s.OrgId = &v
	return b
}

func (b *GetMachineInfoDetailResponseRecordsBuilder) InstanceId(v string) *GetMachineInfoDetailResponseRecordsBuilder {
	b.s.InstanceId = &v
	return b
}

func (b *GetMachineInfoDetailResponseRecordsBuilder) OsVersion(v string) *GetMachineInfoDetailResponseRecordsBuilder {
	b.s.OsVersion = &v
	return b
}

func (b *GetMachineInfoDetailResponseRecordsBuilder) OperateSystem(v string) *GetMachineInfoDetailResponseRecordsBuilder {
	b.s.OperateSystem = &v
	return b
}

func (b *GetMachineInfoDetailResponseRecordsBuilder) OrderSourceCode(v string) *GetMachineInfoDetailResponseRecordsBuilder {
	b.s.OrderSourceCode = &v
	return b
}

func (b *GetMachineInfoDetailResponseRecordsBuilder) CreatedTime(v string) *GetMachineInfoDetailResponseRecordsBuilder {
	b.s.CreatedTime = &v
	return b
}

func (b *GetMachineInfoDetailResponseRecordsBuilder) AzCode(v string) *GetMachineInfoDetailResponseRecordsBuilder {
	b.s.AzCode = &v
	return b
}

func (b *GetMachineInfoDetailResponseRecordsBuilder) SysDiskSize(v string) *GetMachineInfoDetailResponseRecordsBuilder {
	b.s.SysDiskSize = &v
	return b
}

func (b *GetMachineInfoDetailResponseRecordsBuilder) ResourceManageState(v string) *GetMachineInfoDetailResponseRecordsBuilder {
	b.s.ResourceManageState = &v
	return b
}

func (b *GetMachineInfoDetailResponseRecordsBuilder) PoolName(v string) *GetMachineInfoDetailResponseRecordsBuilder {
	b.s.PoolName = &v
	return b
}

func (b *GetMachineInfoDetailResponseRecordsBuilder) CompanyCode(v string) *GetMachineInfoDetailResponseRecordsBuilder {
	b.s.CompanyCode = &v
	return b
}

func (b *GetMachineInfoDetailResponseRecordsBuilder) SecurityGroupName(v string) *GetMachineInfoDetailResponseRecordsBuilder {
	b.s.SecurityGroupName = &v
	return b
}

func (b *GetMachineInfoDetailResponseRecordsBuilder) MachineStatus(v string) *GetMachineInfoDetailResponseRecordsBuilder {
	b.s.MachineStatus = &v
	return b
}

func (b *GetMachineInfoDetailResponseRecordsBuilder) OrderSource(v string) *GetMachineInfoDetailResponseRecordsBuilder {
	b.s.OrderSource = &v
	return b
}

func (b *GetMachineInfoDetailResponseRecordsBuilder) OrgName(v string) *GetMachineInfoDetailResponseRecordsBuilder {
	b.s.OrgName = &v
	return b
}

func (b *GetMachineInfoDetailResponseRecordsBuilder) SnapshotId(v string) *GetMachineInfoDetailResponseRecordsBuilder {
	b.s.SnapshotId = &v
	return b
}

func (b *GetMachineInfoDetailResponseRecordsBuilder) ComputerProtocol(v string) *GetMachineInfoDetailResponseRecordsBuilder {
	b.s.ComputerProtocol = &v
	return b
}

func (b *GetMachineInfoDetailResponseRecordsBuilder) DistributeStatus(v string) *GetMachineInfoDetailResponseRecordsBuilder {
	b.s.DistributeStatus = &v
	return b
}

func (b *GetMachineInfoDetailResponseRecordsBuilder) PolicyName(v string) *GetMachineInfoDetailResponseRecordsBuilder {
	b.s.PolicyName = &v
	return b
}

func (b *GetMachineInfoDetailResponseRecordsBuilder) UserName(v string) *GetMachineInfoDetailResponseRecordsBuilder {
	b.s.UserName = &v
	return b
}

func (b *GetMachineInfoDetailResponseRecordsBuilder) UserId(v string) *GetMachineInfoDetailResponseRecordsBuilder {
	b.s.UserId = &v
	return b
}

func (b *GetMachineInfoDetailResponseRecordsBuilder) ExpireTime(v string) *GetMachineInfoDetailResponseRecordsBuilder {
	b.s.ExpireTime = &v
	return b
}

func (b *GetMachineInfoDetailResponseRecordsBuilder) ExtDataDiskSize(v int64) *GetMachineInfoDetailResponseRecordsBuilder {
	b.s.ExtDataDiskSize = &v
	return b
}

func (b *GetMachineInfoDetailResponseRecordsBuilder) MachineId(v string) *GetMachineInfoDetailResponseRecordsBuilder {
	b.s.MachineId = &v
	return b
}

func (b *GetMachineInfoDetailResponseRecordsBuilder) PolicyId(v string) *GetMachineInfoDetailResponseRecordsBuilder {
	b.s.PolicyId = &v
	return b
}

func (b *GetMachineInfoDetailResponseRecordsBuilder) BandwidthMax(v int32) *GetMachineInfoDetailResponseRecordsBuilder {
	b.s.BandwidthMax = &v
	return b
}

func (b *GetMachineInfoDetailResponseRecordsBuilder) TenantId(v string) *GetMachineInfoDetailResponseRecordsBuilder {
	b.s.TenantId = &v
	return b
}

func (b *GetMachineInfoDetailResponseRecordsBuilder) MopUserName(v string) *GetMachineInfoDetailResponseRecordsBuilder {
	b.s.MopUserName = &v
	return b
}

func (b *GetMachineInfoDetailResponseRecordsBuilder) PoolCode(v string) *GetMachineInfoDetailResponseRecordsBuilder {
	b.s.PoolCode = &v
	return b
}

func (b *GetMachineInfoDetailResponseRecordsBuilder) MopUserId(v string) *GetMachineInfoDetailResponseRecordsBuilder {
	b.s.MopUserId = &v
	return b
}

func (b *GetMachineInfoDetailResponseRecordsBuilder) Region(v string) *GetMachineInfoDetailResponseRecordsBuilder {
	b.s.Region = &v
	return b
}

func (b *GetMachineInfoDetailResponseRecordsBuilder) ResourceTypeName(v string) *GetMachineInfoDetailResponseRecordsBuilder {
	b.s.ResourceTypeName = &v
	return b
}

func (b *GetMachineInfoDetailResponseRecordsBuilder) Dc(v string) *GetMachineInfoDetailResponseRecordsBuilder {
	b.s.Dc = &v
	return b
}

func (b *GetMachineInfoDetailResponseRecordsBuilder) IntranetIP(v string) *GetMachineInfoDetailResponseRecordsBuilder {
	b.s.IntranetIP = &v
	return b
}

func (b *GetMachineInfoDetailResponseRecordsBuilder) Build() *GetMachineInfoDetailResponseRecords {
	return b.s
}


