// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetPolicyInfoByIdResponseBody struct {

    // 剪贴板带宽
	ClipboardBandwidth *int32 `json:"clipboardBandwidth,omitempty"`
    // 摄像头 关闭(disable)、可用(enable) 
	RedirectCamera *string `json:"redirectCamera,omitempty"`
    // 修改时间
	ModifiedTime *string `json:"modifiedTime,omitempty"`
    // 网络重定向开关
	NetRedirectEnable *string `json:"netRedirectEnable,omitempty"`
    // 是否开启水印 禁用(disable)、开启(enable)
	WatermarkEnable *string `json:"watermarkEnable,omitempty"`
    // 音频播放模式
	AudioPlayBackMode *string `json:"audioPlayBackMode,omitempty"`
    // 桌面绑定外设策略状态
	RedirectCustomDetail *[]GetPolicyInfoByIdResponseRedirectCustomDetail `json:"redirectCustomDetail,omitempty"`
    // 厂商名称
	CompanyName *string `json:"companyName,omitempty"`
    // 水印透明度
	WatermarkTransparency *int32 `json:"watermarkTransparency,omitempty"`
    // 安全组相关信息列表
	SecurityGroupDetails *[]GetPolicyInfoByIdResponseSecurityGroupDetails `json:"securityGroupDetails,omitempty"`
    // 是否开启网络扫描仪，关闭(disable)、开启（enable）, 默认关闭 5.3.0_PT新增
	NetworkScannerEnable *string `json:"networkScannerEnable,omitempty"`
    // 并口 关闭(disable)、可用(enable) 默认 enable
	RedirectParallelport *string `json:"redirectParallelport,omitempty"`
    // 重定向剪贴板
	RedirectClipboard *string `json:"redirectClipboard,omitempty"`
    // 是否开启网络打印机，禁用(disable)、可用(enable)
	NetworkPrinterEnable *string `json:"networkPrinterEnable,omitempty"`
    // 防止截屏 禁用(disable)、可用(enable)
	ScreenShotEnable *string `json:"screenShotEnable,omitempty"`
    // 桌面流帧率
	DesktopStreamRat *int32 `json:"desktopStreamRat,omitempty"`
    // 创建时间
	CreatedTime *string `json:"createdTime,omitempty"`
    // 是否开启打印机，禁用(disable)、可用(enable)
	PrinterEnable *string `json:"printerEnable,omitempty"`
    // 资源池名称
	PoolName *string `json:"poolName,omitempty"`
    // 失败原因
	ResponseMsg *string `json:"responseMsg,omitempty"`
    // 串口 关闭(disable)、可用(enable) 默认 enable
	RedirectSerialport *string `json:"redirectSerialport,omitempty"`
    // 终端系统磁盘（即本地系统盘，通常指C盘） ：禁用(disable)、启用(enable)
	TerminalSystemDisk *string `json:"terminalSystemDisk,omitempty"`
    // 网络驱动器 禁用(disable)、启用(enable)
	NetworkDrive *string `json:"networkDrive,omitempty"`
    // 策略名称
	PolicyName *string `json:"policyName,omitempty"`
    // 是否开启USB扫描仪，禁用(disable)、可用(enable)
	ScannerEnable *string `json:"scannerEnable,omitempty"`
    // 字体, 幼圆:YouYuan, 宋体:SimSun, 楷体:KaiTi
	WatermarkFont *string `json:"watermarkFont,omitempty"`
    // 自定义USB重定向策略详情列表
	ExdPolicyExts *[]GetPolicyInfoByIdResponseExdPolicyExts `json:"exdPolicyExts,omitempty"`
    // 重定向USB存储，禁用(disable),读写(enable),只读(ro)
	RedirectUsbStorage *string `json:"redirectUsbStorage,omitempty"`
    // 是否显示云电脑ID 禁用(disable)、开启(enable)
	WatermarkMachineIdEnable *string `json:"watermarkMachineIdEnable,omitempty"`
    // 自定义壁纸绝对路径
	CustomWallpaperAbsolutePath *string `json:"customWallpaperAbsolutePath,omitempty"`
    // 重定向本地磁盘  禁用(disable)、只读(ro)、读写(rw)
	RedirectLocalDisk *string `json:"redirectLocalDisk,omitempty"`
    // 固定驱动器（即本地系统除系统盘外其他磁盘）禁用(disable)、启用(enable)
	FixedDrive *string `json:"fixedDrive,omitempty"`
    // 可用区
	Region *string `json:"region,omitempty"`
    // 重定向模式, 白名单模式:whitelistMode 全局模式:globalMode
	NetRedirectMode *string `json:"netRedirectMode,omitempty"`
    // 水印类型，明水印(0)、盲水印（1）
	WatermarkType *string `json:"watermarkType,omitempty"`
    // 关联电脑数
	BindMachineCount *int32 `json:"bindMachineCount,omitempty"`
    // 可移动驱动器 禁用(disable)、启用(enable)
	RemovableDrive *string `json:"removableDrive,omitempty"`
    // 可用区名称
	RegionName *string `json:"regionName,omitempty"`
    // 自动锁屏,单位分钟 默认锁屏时间 0 不锁屏
	AutoLock *string `json:"autoLock,omitempty"`
    // 网络控制
	NetworkPolicys *[]GetPolicyInfoByIdResponseNetworkPolicys `json:"networkPolicys,omitempty"`
    // 水印倾斜度
	WatermarkAngle *int32 `json:"watermarkAngle,omitempty"`
    // 策略状态
	PolicyStatus *string `json:"policyStatus,omitempty"`
    // 是否显示自定义字段 禁用(disable)、开启(enable)
	WatermarkCustomFieldEnable *string `json:"watermarkCustomFieldEnable,omitempty"`
    // 自定义usb重定向, 禁用(disable)、可用(enable)
	RedirectCustom *string `json:"redirectCustom,omitempty"`
    // 磁盘带宽
	DiskBandwidth *int32 `json:"diskBandwidth,omitempty"`

	TenantName *string `json:"tenantName,omitempty"`
    // 策略描述
	PolicyRemark *string `json:"policyRemark,omitempty"`
    // qos信息
	QosInfoDetail *GetPolicyInfoByIdResponseQosInfoDetail `json:"qosInfoDetail,omitempty"`
    // 录屏相关信息
	ScreenDetail *GetPolicyInfoByIdResponseScreenDetail `json:"screenDetail,omitempty"`
    // usb带宽
	UsbBandwidth *int32 `json:"usbBandwidth,omitempty"`
    // 是否显示时间 禁用(disable)、开启(enable)
	WatermarkTimeEnable *string `json:"watermarkTimeEnable,omitempty"`
    // 安全组相关信息
	SecurityGroupDetail *GetPolicyInfoByIdResponseSecurityGroupDetail `json:"securityGroupDetail,omitempty"`
    // 白名单规则
	NetRedirectWhitelistInfoDtos *[]GetPolicyInfoByIdResponseNetRedirectWhitelistInfoDtos `json:"netRedirectWhitelistInfoDtos,omitempty"`
    // 策略id
	PolicyId *string `json:"policyId,omitempty"`
    // 自定义壁纸是否启用
	CustomWallpaperEnable *string `json:"customWallpaperEnable,omitempty"`
    // 是否显示云电脑mac 禁用(disable)、开启(enable)
	WatermarkMachineMacEnable *string `json:"watermarkMachineMacEnable,omitempty"`
    // 画质
	AutoDesktopQuality *int32 `json:"autoDesktopQuality,omitempty"`
    // 光盘驱动器 禁用(disable)、启用(enable)
	OpticalDiscDrive *string `json:"opticalDiscDrive,omitempty"`
    // 是否显示云电脑IP 禁用(disable)、开启(enable)
	WatermarkMachineIpEnable *string `json:"watermarkMachineIpEnable,omitempty"`

	TenantId *string `json:"tenantId,omitempty"`
    // 自定义字段
	WatermarkCustomField *string `json:"watermarkCustomField,omitempty"`
    // 资源池编码
	PoolCode *string `json:"poolCode,omitempty"`
    // 带宽是否启用
	BandwidthLimitEnable *string `json:"bandwidthLimitEnable,omitempty"`
    // 终端登录策略
	TerminalLoginPolicys *[]GetPolicyInfoByIdResponseTerminalLoginPolicys `json:"terminalLoginPolicys,omitempty"`
    // 水印字体颜色
	WatermarkColor *string `json:"watermarkColor,omitempty"`
    // 桌面工具栏策略
	DesktopToolbarPolicys *[]GetPolicyInfoByIdResponseDesktopToolbarPolicys `json:"desktopToolbarPolicys,omitempty"`
    // 水印字体大小，单位px
	WatermarkSize *int32 `json:"watermarkSize,omitempty"`
    // 网络工作区
	Dc *string `json:"dc,omitempty"`
    // 图流设置
	StreamingVideo *int32 `json:"streamingVideo,omitempty"`
}

func (s GetPolicyInfoByIdResponseBody) String() string {
	return utils.Beautify(s)
}

func (s GetPolicyInfoByIdResponseBody) GoString() string {
	return s.String()
}

func (s GetPolicyInfoByIdResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetPolicyInfoByIdResponseBody) SetClipboardBandwidth(v int32) *GetPolicyInfoByIdResponseBody {
	s.ClipboardBandwidth = &v
	return s
}

func (s *GetPolicyInfoByIdResponseBody) SetRedirectCamera(v string) *GetPolicyInfoByIdResponseBody {
	s.RedirectCamera = &v
	return s
}

func (s *GetPolicyInfoByIdResponseBody) SetModifiedTime(v string) *GetPolicyInfoByIdResponseBody {
	s.ModifiedTime = &v
	return s
}

func (s *GetPolicyInfoByIdResponseBody) SetNetRedirectEnable(v string) *GetPolicyInfoByIdResponseBody {
	s.NetRedirectEnable = &v
	return s
}

func (s *GetPolicyInfoByIdResponseBody) SetWatermarkEnable(v string) *GetPolicyInfoByIdResponseBody {
	s.WatermarkEnable = &v
	return s
}

func (s *GetPolicyInfoByIdResponseBody) SetAudioPlayBackMode(v string) *GetPolicyInfoByIdResponseBody {
	s.AudioPlayBackMode = &v
	return s
}

func (s *GetPolicyInfoByIdResponseBody) SetRedirectCustomDetail(v []GetPolicyInfoByIdResponseRedirectCustomDetail) *GetPolicyInfoByIdResponseBody {
	s.RedirectCustomDetail = &v
	return s
}

func (s *GetPolicyInfoByIdResponseBody) SetCompanyName(v string) *GetPolicyInfoByIdResponseBody {
	s.CompanyName = &v
	return s
}

func (s *GetPolicyInfoByIdResponseBody) SetWatermarkTransparency(v int32) *GetPolicyInfoByIdResponseBody {
	s.WatermarkTransparency = &v
	return s
}

func (s *GetPolicyInfoByIdResponseBody) SetSecurityGroupDetails(v []GetPolicyInfoByIdResponseSecurityGroupDetails) *GetPolicyInfoByIdResponseBody {
	s.SecurityGroupDetails = &v
	return s
}

func (s *GetPolicyInfoByIdResponseBody) SetNetworkScannerEnable(v string) *GetPolicyInfoByIdResponseBody {
	s.NetworkScannerEnable = &v
	return s
}

func (s *GetPolicyInfoByIdResponseBody) SetRedirectParallelport(v string) *GetPolicyInfoByIdResponseBody {
	s.RedirectParallelport = &v
	return s
}

func (s *GetPolicyInfoByIdResponseBody) SetRedirectClipboard(v string) *GetPolicyInfoByIdResponseBody {
	s.RedirectClipboard = &v
	return s
}

func (s *GetPolicyInfoByIdResponseBody) SetNetworkPrinterEnable(v string) *GetPolicyInfoByIdResponseBody {
	s.NetworkPrinterEnable = &v
	return s
}

func (s *GetPolicyInfoByIdResponseBody) SetScreenShotEnable(v string) *GetPolicyInfoByIdResponseBody {
	s.ScreenShotEnable = &v
	return s
}

func (s *GetPolicyInfoByIdResponseBody) SetDesktopStreamRat(v int32) *GetPolicyInfoByIdResponseBody {
	s.DesktopStreamRat = &v
	return s
}

func (s *GetPolicyInfoByIdResponseBody) SetCreatedTime(v string) *GetPolicyInfoByIdResponseBody {
	s.CreatedTime = &v
	return s
}

func (s *GetPolicyInfoByIdResponseBody) SetPrinterEnable(v string) *GetPolicyInfoByIdResponseBody {
	s.PrinterEnable = &v
	return s
}

func (s *GetPolicyInfoByIdResponseBody) SetPoolName(v string) *GetPolicyInfoByIdResponseBody {
	s.PoolName = &v
	return s
}

func (s *GetPolicyInfoByIdResponseBody) SetResponseMsg(v string) *GetPolicyInfoByIdResponseBody {
	s.ResponseMsg = &v
	return s
}

func (s *GetPolicyInfoByIdResponseBody) SetRedirectSerialport(v string) *GetPolicyInfoByIdResponseBody {
	s.RedirectSerialport = &v
	return s
}

func (s *GetPolicyInfoByIdResponseBody) SetTerminalSystemDisk(v string) *GetPolicyInfoByIdResponseBody {
	s.TerminalSystemDisk = &v
	return s
}

func (s *GetPolicyInfoByIdResponseBody) SetNetworkDrive(v string) *GetPolicyInfoByIdResponseBody {
	s.NetworkDrive = &v
	return s
}

func (s *GetPolicyInfoByIdResponseBody) SetPolicyName(v string) *GetPolicyInfoByIdResponseBody {
	s.PolicyName = &v
	return s
}

func (s *GetPolicyInfoByIdResponseBody) SetScannerEnable(v string) *GetPolicyInfoByIdResponseBody {
	s.ScannerEnable = &v
	return s
}

func (s *GetPolicyInfoByIdResponseBody) SetWatermarkFont(v string) *GetPolicyInfoByIdResponseBody {
	s.WatermarkFont = &v
	return s
}

func (s *GetPolicyInfoByIdResponseBody) SetExdPolicyExts(v []GetPolicyInfoByIdResponseExdPolicyExts) *GetPolicyInfoByIdResponseBody {
	s.ExdPolicyExts = &v
	return s
}

func (s *GetPolicyInfoByIdResponseBody) SetRedirectUsbStorage(v string) *GetPolicyInfoByIdResponseBody {
	s.RedirectUsbStorage = &v
	return s
}

func (s *GetPolicyInfoByIdResponseBody) SetWatermarkMachineIdEnable(v string) *GetPolicyInfoByIdResponseBody {
	s.WatermarkMachineIdEnable = &v
	return s
}

func (s *GetPolicyInfoByIdResponseBody) SetCustomWallpaperAbsolutePath(v string) *GetPolicyInfoByIdResponseBody {
	s.CustomWallpaperAbsolutePath = &v
	return s
}

func (s *GetPolicyInfoByIdResponseBody) SetRedirectLocalDisk(v string) *GetPolicyInfoByIdResponseBody {
	s.RedirectLocalDisk = &v
	return s
}

func (s *GetPolicyInfoByIdResponseBody) SetFixedDrive(v string) *GetPolicyInfoByIdResponseBody {
	s.FixedDrive = &v
	return s
}

func (s *GetPolicyInfoByIdResponseBody) SetRegion(v string) *GetPolicyInfoByIdResponseBody {
	s.Region = &v
	return s
}

func (s *GetPolicyInfoByIdResponseBody) SetNetRedirectMode(v string) *GetPolicyInfoByIdResponseBody {
	s.NetRedirectMode = &v
	return s
}

func (s *GetPolicyInfoByIdResponseBody) SetWatermarkType(v string) *GetPolicyInfoByIdResponseBody {
	s.WatermarkType = &v
	return s
}

func (s *GetPolicyInfoByIdResponseBody) SetBindMachineCount(v int32) *GetPolicyInfoByIdResponseBody {
	s.BindMachineCount = &v
	return s
}

func (s *GetPolicyInfoByIdResponseBody) SetRemovableDrive(v string) *GetPolicyInfoByIdResponseBody {
	s.RemovableDrive = &v
	return s
}

func (s *GetPolicyInfoByIdResponseBody) SetRegionName(v string) *GetPolicyInfoByIdResponseBody {
	s.RegionName = &v
	return s
}

func (s *GetPolicyInfoByIdResponseBody) SetAutoLock(v string) *GetPolicyInfoByIdResponseBody {
	s.AutoLock = &v
	return s
}

func (s *GetPolicyInfoByIdResponseBody) SetNetworkPolicys(v []GetPolicyInfoByIdResponseNetworkPolicys) *GetPolicyInfoByIdResponseBody {
	s.NetworkPolicys = &v
	return s
}

func (s *GetPolicyInfoByIdResponseBody) SetWatermarkAngle(v int32) *GetPolicyInfoByIdResponseBody {
	s.WatermarkAngle = &v
	return s
}

func (s *GetPolicyInfoByIdResponseBody) SetPolicyStatus(v string) *GetPolicyInfoByIdResponseBody {
	s.PolicyStatus = &v
	return s
}

func (s *GetPolicyInfoByIdResponseBody) SetWatermarkCustomFieldEnable(v string) *GetPolicyInfoByIdResponseBody {
	s.WatermarkCustomFieldEnable = &v
	return s
}

func (s *GetPolicyInfoByIdResponseBody) SetRedirectCustom(v string) *GetPolicyInfoByIdResponseBody {
	s.RedirectCustom = &v
	return s
}

func (s *GetPolicyInfoByIdResponseBody) SetDiskBandwidth(v int32) *GetPolicyInfoByIdResponseBody {
	s.DiskBandwidth = &v
	return s
}

func (s *GetPolicyInfoByIdResponseBody) SetTenantName(v string) *GetPolicyInfoByIdResponseBody {
	s.TenantName = &v
	return s
}

func (s *GetPolicyInfoByIdResponseBody) SetPolicyRemark(v string) *GetPolicyInfoByIdResponseBody {
	s.PolicyRemark = &v
	return s
}

func (s *GetPolicyInfoByIdResponseBody) SetQosInfoDetail(v *GetPolicyInfoByIdResponseQosInfoDetail) *GetPolicyInfoByIdResponseBody {
	s.QosInfoDetail = v
	return s
}

func (s *GetPolicyInfoByIdResponseBody) SetScreenDetail(v *GetPolicyInfoByIdResponseScreenDetail) *GetPolicyInfoByIdResponseBody {
	s.ScreenDetail = v
	return s
}

func (s *GetPolicyInfoByIdResponseBody) SetUsbBandwidth(v int32) *GetPolicyInfoByIdResponseBody {
	s.UsbBandwidth = &v
	return s
}

func (s *GetPolicyInfoByIdResponseBody) SetWatermarkTimeEnable(v string) *GetPolicyInfoByIdResponseBody {
	s.WatermarkTimeEnable = &v
	return s
}

func (s *GetPolicyInfoByIdResponseBody) SetSecurityGroupDetail(v *GetPolicyInfoByIdResponseSecurityGroupDetail) *GetPolicyInfoByIdResponseBody {
	s.SecurityGroupDetail = v
	return s
}

func (s *GetPolicyInfoByIdResponseBody) SetNetRedirectWhitelistInfoDtos(v []GetPolicyInfoByIdResponseNetRedirectWhitelistInfoDtos) *GetPolicyInfoByIdResponseBody {
	s.NetRedirectWhitelistInfoDtos = &v
	return s
}

func (s *GetPolicyInfoByIdResponseBody) SetPolicyId(v string) *GetPolicyInfoByIdResponseBody {
	s.PolicyId = &v
	return s
}

func (s *GetPolicyInfoByIdResponseBody) SetCustomWallpaperEnable(v string) *GetPolicyInfoByIdResponseBody {
	s.CustomWallpaperEnable = &v
	return s
}

func (s *GetPolicyInfoByIdResponseBody) SetWatermarkMachineMacEnable(v string) *GetPolicyInfoByIdResponseBody {
	s.WatermarkMachineMacEnable = &v
	return s
}

func (s *GetPolicyInfoByIdResponseBody) SetAutoDesktopQuality(v int32) *GetPolicyInfoByIdResponseBody {
	s.AutoDesktopQuality = &v
	return s
}

func (s *GetPolicyInfoByIdResponseBody) SetOpticalDiscDrive(v string) *GetPolicyInfoByIdResponseBody {
	s.OpticalDiscDrive = &v
	return s
}

func (s *GetPolicyInfoByIdResponseBody) SetWatermarkMachineIpEnable(v string) *GetPolicyInfoByIdResponseBody {
	s.WatermarkMachineIpEnable = &v
	return s
}

func (s *GetPolicyInfoByIdResponseBody) SetTenantId(v string) *GetPolicyInfoByIdResponseBody {
	s.TenantId = &v
	return s
}

func (s *GetPolicyInfoByIdResponseBody) SetWatermarkCustomField(v string) *GetPolicyInfoByIdResponseBody {
	s.WatermarkCustomField = &v
	return s
}

func (s *GetPolicyInfoByIdResponseBody) SetPoolCode(v string) *GetPolicyInfoByIdResponseBody {
	s.PoolCode = &v
	return s
}

func (s *GetPolicyInfoByIdResponseBody) SetBandwidthLimitEnable(v string) *GetPolicyInfoByIdResponseBody {
	s.BandwidthLimitEnable = &v
	return s
}

func (s *GetPolicyInfoByIdResponseBody) SetTerminalLoginPolicys(v []GetPolicyInfoByIdResponseTerminalLoginPolicys) *GetPolicyInfoByIdResponseBody {
	s.TerminalLoginPolicys = &v
	return s
}

func (s *GetPolicyInfoByIdResponseBody) SetWatermarkColor(v string) *GetPolicyInfoByIdResponseBody {
	s.WatermarkColor = &v
	return s
}

func (s *GetPolicyInfoByIdResponseBody) SetDesktopToolbarPolicys(v []GetPolicyInfoByIdResponseDesktopToolbarPolicys) *GetPolicyInfoByIdResponseBody {
	s.DesktopToolbarPolicys = &v
	return s
}

func (s *GetPolicyInfoByIdResponseBody) SetWatermarkSize(v int32) *GetPolicyInfoByIdResponseBody {
	s.WatermarkSize = &v
	return s
}

func (s *GetPolicyInfoByIdResponseBody) SetDc(v string) *GetPolicyInfoByIdResponseBody {
	s.Dc = &v
	return s
}

func (s *GetPolicyInfoByIdResponseBody) SetStreamingVideo(v int32) *GetPolicyInfoByIdResponseBody {
	s.StreamingVideo = &v
	return s
}


type GetPolicyInfoByIdResponseBodyBuilder struct {
	s *GetPolicyInfoByIdResponseBody
}

func NewGetPolicyInfoByIdResponseBodyBuilder() *GetPolicyInfoByIdResponseBodyBuilder {
	s := &GetPolicyInfoByIdResponseBody{}
	b := &GetPolicyInfoByIdResponseBodyBuilder{s: s}
	return b
}

func (b *GetPolicyInfoByIdResponseBodyBuilder) ClipboardBandwidth(v int32) *GetPolicyInfoByIdResponseBodyBuilder {
	b.s.ClipboardBandwidth = &v
	return b
}

func (b *GetPolicyInfoByIdResponseBodyBuilder) RedirectCamera(v string) *GetPolicyInfoByIdResponseBodyBuilder {
	b.s.RedirectCamera = &v
	return b
}

func (b *GetPolicyInfoByIdResponseBodyBuilder) ModifiedTime(v string) *GetPolicyInfoByIdResponseBodyBuilder {
	b.s.ModifiedTime = &v
	return b
}

func (b *GetPolicyInfoByIdResponseBodyBuilder) NetRedirectEnable(v string) *GetPolicyInfoByIdResponseBodyBuilder {
	b.s.NetRedirectEnable = &v
	return b
}

func (b *GetPolicyInfoByIdResponseBodyBuilder) WatermarkEnable(v string) *GetPolicyInfoByIdResponseBodyBuilder {
	b.s.WatermarkEnable = &v
	return b
}

func (b *GetPolicyInfoByIdResponseBodyBuilder) AudioPlayBackMode(v string) *GetPolicyInfoByIdResponseBodyBuilder {
	b.s.AudioPlayBackMode = &v
	return b
}

func (b *GetPolicyInfoByIdResponseBodyBuilder) RedirectCustomDetail(v []GetPolicyInfoByIdResponseRedirectCustomDetail) *GetPolicyInfoByIdResponseBodyBuilder {
	b.s.RedirectCustomDetail = &v
	return b
}

func (b *GetPolicyInfoByIdResponseBodyBuilder) CompanyName(v string) *GetPolicyInfoByIdResponseBodyBuilder {
	b.s.CompanyName = &v
	return b
}

func (b *GetPolicyInfoByIdResponseBodyBuilder) WatermarkTransparency(v int32) *GetPolicyInfoByIdResponseBodyBuilder {
	b.s.WatermarkTransparency = &v
	return b
}

func (b *GetPolicyInfoByIdResponseBodyBuilder) SecurityGroupDetails(v []GetPolicyInfoByIdResponseSecurityGroupDetails) *GetPolicyInfoByIdResponseBodyBuilder {
	b.s.SecurityGroupDetails = &v
	return b
}

func (b *GetPolicyInfoByIdResponseBodyBuilder) NetworkScannerEnable(v string) *GetPolicyInfoByIdResponseBodyBuilder {
	b.s.NetworkScannerEnable = &v
	return b
}

func (b *GetPolicyInfoByIdResponseBodyBuilder) RedirectParallelport(v string) *GetPolicyInfoByIdResponseBodyBuilder {
	b.s.RedirectParallelport = &v
	return b
}

func (b *GetPolicyInfoByIdResponseBodyBuilder) RedirectClipboard(v string) *GetPolicyInfoByIdResponseBodyBuilder {
	b.s.RedirectClipboard = &v
	return b
}

func (b *GetPolicyInfoByIdResponseBodyBuilder) NetworkPrinterEnable(v string) *GetPolicyInfoByIdResponseBodyBuilder {
	b.s.NetworkPrinterEnable = &v
	return b
}

func (b *GetPolicyInfoByIdResponseBodyBuilder) ScreenShotEnable(v string) *GetPolicyInfoByIdResponseBodyBuilder {
	b.s.ScreenShotEnable = &v
	return b
}

func (b *GetPolicyInfoByIdResponseBodyBuilder) DesktopStreamRat(v int32) *GetPolicyInfoByIdResponseBodyBuilder {
	b.s.DesktopStreamRat = &v
	return b
}

func (b *GetPolicyInfoByIdResponseBodyBuilder) CreatedTime(v string) *GetPolicyInfoByIdResponseBodyBuilder {
	b.s.CreatedTime = &v
	return b
}

func (b *GetPolicyInfoByIdResponseBodyBuilder) PrinterEnable(v string) *GetPolicyInfoByIdResponseBodyBuilder {
	b.s.PrinterEnable = &v
	return b
}

func (b *GetPolicyInfoByIdResponseBodyBuilder) PoolName(v string) *GetPolicyInfoByIdResponseBodyBuilder {
	b.s.PoolName = &v
	return b
}

func (b *GetPolicyInfoByIdResponseBodyBuilder) ResponseMsg(v string) *GetPolicyInfoByIdResponseBodyBuilder {
	b.s.ResponseMsg = &v
	return b
}

func (b *GetPolicyInfoByIdResponseBodyBuilder) RedirectSerialport(v string) *GetPolicyInfoByIdResponseBodyBuilder {
	b.s.RedirectSerialport = &v
	return b
}

func (b *GetPolicyInfoByIdResponseBodyBuilder) TerminalSystemDisk(v string) *GetPolicyInfoByIdResponseBodyBuilder {
	b.s.TerminalSystemDisk = &v
	return b
}

func (b *GetPolicyInfoByIdResponseBodyBuilder) NetworkDrive(v string) *GetPolicyInfoByIdResponseBodyBuilder {
	b.s.NetworkDrive = &v
	return b
}

func (b *GetPolicyInfoByIdResponseBodyBuilder) PolicyName(v string) *GetPolicyInfoByIdResponseBodyBuilder {
	b.s.PolicyName = &v
	return b
}

func (b *GetPolicyInfoByIdResponseBodyBuilder) ScannerEnable(v string) *GetPolicyInfoByIdResponseBodyBuilder {
	b.s.ScannerEnable = &v
	return b
}

func (b *GetPolicyInfoByIdResponseBodyBuilder) WatermarkFont(v string) *GetPolicyInfoByIdResponseBodyBuilder {
	b.s.WatermarkFont = &v
	return b
}

func (b *GetPolicyInfoByIdResponseBodyBuilder) ExdPolicyExts(v []GetPolicyInfoByIdResponseExdPolicyExts) *GetPolicyInfoByIdResponseBodyBuilder {
	b.s.ExdPolicyExts = &v
	return b
}

func (b *GetPolicyInfoByIdResponseBodyBuilder) RedirectUsbStorage(v string) *GetPolicyInfoByIdResponseBodyBuilder {
	b.s.RedirectUsbStorage = &v
	return b
}

func (b *GetPolicyInfoByIdResponseBodyBuilder) WatermarkMachineIdEnable(v string) *GetPolicyInfoByIdResponseBodyBuilder {
	b.s.WatermarkMachineIdEnable = &v
	return b
}

func (b *GetPolicyInfoByIdResponseBodyBuilder) CustomWallpaperAbsolutePath(v string) *GetPolicyInfoByIdResponseBodyBuilder {
	b.s.CustomWallpaperAbsolutePath = &v
	return b
}

func (b *GetPolicyInfoByIdResponseBodyBuilder) RedirectLocalDisk(v string) *GetPolicyInfoByIdResponseBodyBuilder {
	b.s.RedirectLocalDisk = &v
	return b
}

func (b *GetPolicyInfoByIdResponseBodyBuilder) FixedDrive(v string) *GetPolicyInfoByIdResponseBodyBuilder {
	b.s.FixedDrive = &v
	return b
}

func (b *GetPolicyInfoByIdResponseBodyBuilder) Region(v string) *GetPolicyInfoByIdResponseBodyBuilder {
	b.s.Region = &v
	return b
}

func (b *GetPolicyInfoByIdResponseBodyBuilder) NetRedirectMode(v string) *GetPolicyInfoByIdResponseBodyBuilder {
	b.s.NetRedirectMode = &v
	return b
}

func (b *GetPolicyInfoByIdResponseBodyBuilder) WatermarkType(v string) *GetPolicyInfoByIdResponseBodyBuilder {
	b.s.WatermarkType = &v
	return b
}

func (b *GetPolicyInfoByIdResponseBodyBuilder) BindMachineCount(v int32) *GetPolicyInfoByIdResponseBodyBuilder {
	b.s.BindMachineCount = &v
	return b
}

func (b *GetPolicyInfoByIdResponseBodyBuilder) RemovableDrive(v string) *GetPolicyInfoByIdResponseBodyBuilder {
	b.s.RemovableDrive = &v
	return b
}

func (b *GetPolicyInfoByIdResponseBodyBuilder) RegionName(v string) *GetPolicyInfoByIdResponseBodyBuilder {
	b.s.RegionName = &v
	return b
}

func (b *GetPolicyInfoByIdResponseBodyBuilder) AutoLock(v string) *GetPolicyInfoByIdResponseBodyBuilder {
	b.s.AutoLock = &v
	return b
}

func (b *GetPolicyInfoByIdResponseBodyBuilder) NetworkPolicys(v []GetPolicyInfoByIdResponseNetworkPolicys) *GetPolicyInfoByIdResponseBodyBuilder {
	b.s.NetworkPolicys = &v
	return b
}

func (b *GetPolicyInfoByIdResponseBodyBuilder) WatermarkAngle(v int32) *GetPolicyInfoByIdResponseBodyBuilder {
	b.s.WatermarkAngle = &v
	return b
}

func (b *GetPolicyInfoByIdResponseBodyBuilder) PolicyStatus(v string) *GetPolicyInfoByIdResponseBodyBuilder {
	b.s.PolicyStatus = &v
	return b
}

func (b *GetPolicyInfoByIdResponseBodyBuilder) WatermarkCustomFieldEnable(v string) *GetPolicyInfoByIdResponseBodyBuilder {
	b.s.WatermarkCustomFieldEnable = &v
	return b
}

func (b *GetPolicyInfoByIdResponseBodyBuilder) RedirectCustom(v string) *GetPolicyInfoByIdResponseBodyBuilder {
	b.s.RedirectCustom = &v
	return b
}

func (b *GetPolicyInfoByIdResponseBodyBuilder) DiskBandwidth(v int32) *GetPolicyInfoByIdResponseBodyBuilder {
	b.s.DiskBandwidth = &v
	return b
}

func (b *GetPolicyInfoByIdResponseBodyBuilder) TenantName(v string) *GetPolicyInfoByIdResponseBodyBuilder {
	b.s.TenantName = &v
	return b
}

func (b *GetPolicyInfoByIdResponseBodyBuilder) PolicyRemark(v string) *GetPolicyInfoByIdResponseBodyBuilder {
	b.s.PolicyRemark = &v
	return b
}

func (b *GetPolicyInfoByIdResponseBodyBuilder) QosInfoDetail(v *GetPolicyInfoByIdResponseQosInfoDetail) *GetPolicyInfoByIdResponseBodyBuilder {
	b.s.QosInfoDetail = v
	return b
}

func (b *GetPolicyInfoByIdResponseBodyBuilder) ScreenDetail(v *GetPolicyInfoByIdResponseScreenDetail) *GetPolicyInfoByIdResponseBodyBuilder {
	b.s.ScreenDetail = v
	return b
}

func (b *GetPolicyInfoByIdResponseBodyBuilder) UsbBandwidth(v int32) *GetPolicyInfoByIdResponseBodyBuilder {
	b.s.UsbBandwidth = &v
	return b
}

func (b *GetPolicyInfoByIdResponseBodyBuilder) WatermarkTimeEnable(v string) *GetPolicyInfoByIdResponseBodyBuilder {
	b.s.WatermarkTimeEnable = &v
	return b
}

func (b *GetPolicyInfoByIdResponseBodyBuilder) SecurityGroupDetail(v *GetPolicyInfoByIdResponseSecurityGroupDetail) *GetPolicyInfoByIdResponseBodyBuilder {
	b.s.SecurityGroupDetail = v
	return b
}

func (b *GetPolicyInfoByIdResponseBodyBuilder) NetRedirectWhitelistInfoDtos(v []GetPolicyInfoByIdResponseNetRedirectWhitelistInfoDtos) *GetPolicyInfoByIdResponseBodyBuilder {
	b.s.NetRedirectWhitelistInfoDtos = &v
	return b
}

func (b *GetPolicyInfoByIdResponseBodyBuilder) PolicyId(v string) *GetPolicyInfoByIdResponseBodyBuilder {
	b.s.PolicyId = &v
	return b
}

func (b *GetPolicyInfoByIdResponseBodyBuilder) CustomWallpaperEnable(v string) *GetPolicyInfoByIdResponseBodyBuilder {
	b.s.CustomWallpaperEnable = &v
	return b
}

func (b *GetPolicyInfoByIdResponseBodyBuilder) WatermarkMachineMacEnable(v string) *GetPolicyInfoByIdResponseBodyBuilder {
	b.s.WatermarkMachineMacEnable = &v
	return b
}

func (b *GetPolicyInfoByIdResponseBodyBuilder) AutoDesktopQuality(v int32) *GetPolicyInfoByIdResponseBodyBuilder {
	b.s.AutoDesktopQuality = &v
	return b
}

func (b *GetPolicyInfoByIdResponseBodyBuilder) OpticalDiscDrive(v string) *GetPolicyInfoByIdResponseBodyBuilder {
	b.s.OpticalDiscDrive = &v
	return b
}

func (b *GetPolicyInfoByIdResponseBodyBuilder) WatermarkMachineIpEnable(v string) *GetPolicyInfoByIdResponseBodyBuilder {
	b.s.WatermarkMachineIpEnable = &v
	return b
}

func (b *GetPolicyInfoByIdResponseBodyBuilder) TenantId(v string) *GetPolicyInfoByIdResponseBodyBuilder {
	b.s.TenantId = &v
	return b
}

func (b *GetPolicyInfoByIdResponseBodyBuilder) WatermarkCustomField(v string) *GetPolicyInfoByIdResponseBodyBuilder {
	b.s.WatermarkCustomField = &v
	return b
}

func (b *GetPolicyInfoByIdResponseBodyBuilder) PoolCode(v string) *GetPolicyInfoByIdResponseBodyBuilder {
	b.s.PoolCode = &v
	return b
}

func (b *GetPolicyInfoByIdResponseBodyBuilder) BandwidthLimitEnable(v string) *GetPolicyInfoByIdResponseBodyBuilder {
	b.s.BandwidthLimitEnable = &v
	return b
}

func (b *GetPolicyInfoByIdResponseBodyBuilder) TerminalLoginPolicys(v []GetPolicyInfoByIdResponseTerminalLoginPolicys) *GetPolicyInfoByIdResponseBodyBuilder {
	b.s.TerminalLoginPolicys = &v
	return b
}

func (b *GetPolicyInfoByIdResponseBodyBuilder) WatermarkColor(v string) *GetPolicyInfoByIdResponseBodyBuilder {
	b.s.WatermarkColor = &v
	return b
}

func (b *GetPolicyInfoByIdResponseBodyBuilder) DesktopToolbarPolicys(v []GetPolicyInfoByIdResponseDesktopToolbarPolicys) *GetPolicyInfoByIdResponseBodyBuilder {
	b.s.DesktopToolbarPolicys = &v
	return b
}

func (b *GetPolicyInfoByIdResponseBodyBuilder) WatermarkSize(v int32) *GetPolicyInfoByIdResponseBodyBuilder {
	b.s.WatermarkSize = &v
	return b
}

func (b *GetPolicyInfoByIdResponseBodyBuilder) Dc(v string) *GetPolicyInfoByIdResponseBodyBuilder {
	b.s.Dc = &v
	return b
}

func (b *GetPolicyInfoByIdResponseBodyBuilder) StreamingVideo(v int32) *GetPolicyInfoByIdResponseBodyBuilder {
	b.s.StreamingVideo = &v
	return b
}

func (b *GetPolicyInfoByIdResponseBodyBuilder) Build() *GetPolicyInfoByIdResponseBody {
	return b.s
}


