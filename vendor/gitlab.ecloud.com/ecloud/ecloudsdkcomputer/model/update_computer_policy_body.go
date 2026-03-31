// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UpdateComputerPolicyBody struct {
    position.Body
    // 摄像头 关闭(disable)、可用(enable) 默认enable
	RedirectCamera *string `json:"redirectCamera,omitempty"`
    // 文件拷贝带宽控制-剪贴板范围, 500 - 2000 KB , 默认 500 KB, 增减步长 500, 单位(KB)
	ClipboardBandwidth *int32 `json:"clipboardBandwidth,omitempty"`
    // 录屏管控
	ExdPolicyScreenRecordeClass *UpdateComputerPolicyRequestExdPolicyScreenRecordeClass `json:"exdPolicyScreenRecordeClass,omitempty"`
    // 音频播放模式, 1:默认模式 2:极速模式/加速模式 3:流畅模式
	AudioPlayBackMode *string `json:"audioPlayBackMode,omitempty"`
    // 自动锁屏,单位分钟 默认锁屏时间 0 不锁屏
	AutoLock *string `json:"autoLock,omitempty"`
    // 网络控制
	NetworkPolicys *[]UpdateComputerPolicyRequestNetworkPolicys `json:"networkPolicys,omitempty"`
    // 是否开启网络扫描仪，关闭(disable)、开启（enable）, 默认关闭
	NetworkScannerEnable *string `json:"networkScannerEnable,omitempty"`
    // 并口 关闭(disable)、可用(enable) 默认 enable
	RedirectParallelport *string `json:"redirectParallelport,omitempty"`
    // 云侧网络重定向
	ExdCloudSideNetRedirectClass *UpdateComputerPolicyRequestExdCloudSideNetRedirectClass `json:"exdCloudSideNetRedirectClass,omitempty"`
    // 重定向剪贴板，禁用(disable)、云端到本地(c2l), 本地到云端(l2c), 双向(both)，默认both
	RedirectClipboard *string `json:"redirectClipboard,omitempty"`
    // 是否开启网络打印机，禁用(disable)、可用(enable)
	NetworkPrinterEnable *string `json:"networkPrinterEnable,omitempty"`
    // 本地磁盘映射作用范围
	RedirectLocalDiskScope *UpdateComputerPolicyRequestRedirectLocalDiskScope `json:"redirectLocalDiskScope,omitempty"`
    // 自定义USB重定向策略配置 关闭(disable)、可用(enable) 默认 disable
	RedirectCustom *string `json:"redirectCustom,omitempty"`
    // 水印详情
	WatermarkClass *UpdateComputerPolicyRequestWatermarkClass `json:"watermarkClass,omitempty"`
    // 防截屏 禁用(disable)、可用(enable) 默认 enable
	ScreenShotEnable *string `json:"screenShotEnable,omitempty"`
    // 磁盘映射范围, 500 - 2000 KB , 默认 500 KB, 增减步长 500, 单位(KB)
	DiskBandwidth *int32 `json:"diskBandwidth,omitempty"`
    // 图像上限帧率 默认30 最小值12 最大值100
	DesktopStreamRat *int32 `json:"desktopStreamRat,omitempty"`
    // 策略描述
	PolicyRemark *string `json:"policyRemark,omitempty"`
    // 自定义电脑壁纸
	ExdPolicyCustomWallpaperClass *UpdateComputerPolicyRequestExdPolicyCustomWallpaperClass `json:"exdPolicyCustomWallpaperClass,omitempty"`
    // usb带宽范围, 500 - 2000 KB , 默认 500 KB, 增减步长 500,  单位(KB)
	UsbBandwidth *int32 `json:"usbBandwidth,omitempty"`
    // 音频播放模式是否开启, 禁用(disable)、开启(enable) , 默认 关闭 disable
	AudioPlayBackModeEnable *string `json:"audioPlayBackModeEnable,omitempty"`
    // 是否开启打印机，禁用(disable)、可用(enable)
	PrinterEnable *string `json:"printerEnable,omitempty"`
    // Qos规则的UID
	QosUid *string `json:"qosUid,omitempty"`
    // 串口 关闭(disable)、可用(enable) 默认 enable
	RedirectSerialport *string `json:"redirectSerialport,omitempty"`
    // 电脑协议（V2.0:自研CMSSZTE，V1.1：新华三，V1.2：中兴）
	ComputerProtocol *string `json:"computerProtocol,omitempty"`
    // 策略名称
	PolicyName *string `json:"policyName,omitempty"`
    // 是否开启USB扫描仪，禁用(disable)、可用(enable)
	ScannerEnable *string `json:"scannerEnable,omitempty"`
    // 自定义USB重定向策略详情列表
	ExdPolicyExts *[]UpdateComputerPolicyRequestExdPolicyExts `json:"exdPolicyExts,omitempty"`
    // 重定向USB存储，禁用(disable),读写(enable),只读(ro),默认enable
	RedirectUsbStorage *string `json:"redirectUsbStorage,omitempty"`
    // 策略id
	PolicyId *string `json:"policyId,omitempty"`
    // 图像画质 默认 2 最小值0 最大值3
	AutoDesktopQuality *int32 `json:"autoDesktopQuality,omitempty"`
    // 所属资源池
	PoolCode *string `json:"poolCode,omitempty"`
    // 带宽是否启用 关闭(disable)、可用(enable) 默认 disable
	BandwidthLimitEnable *string `json:"bandwidthLimitEnable,omitempty"`
    // 所属可用区
	Region *string `json:"region,omitempty"`
    // 终端登录策略
	TerminalLoginPolicys *[]UpdateComputerPolicyRequestTerminalLoginPolicys `json:"terminalLoginPolicys,omitempty"`
    // 桌面工具栏策略
	DesktopToolbarPolicys *[]UpdateComputerPolicyRequestDesktopToolbarPolicys `json:"desktopToolbarPolicys,omitempty"`
    // 安全组UID
	SecurityGroupUid *string `json:"securityGroupUid,omitempty"`
    // 网络工作区，V2.0协议非必传
	Dc *string `json:"dc,omitempty"`
    // 图流设置 默认 2 最小值1 最大值2
	StreamingVideo *int32 `json:"streamingVideo,omitempty"`
}

func (s UpdateComputerPolicyBody) String() string {
	return utils.Beautify(s)
}

func (s UpdateComputerPolicyBody) GoString() string {
	return s.String()
}

func (s UpdateComputerPolicyBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UpdateComputerPolicyBody) SetRedirectCamera(v string) *UpdateComputerPolicyBody {
	s.RedirectCamera = &v
	return s
}

func (s *UpdateComputerPolicyBody) SetClipboardBandwidth(v int32) *UpdateComputerPolicyBody {
	s.ClipboardBandwidth = &v
	return s
}

func (s *UpdateComputerPolicyBody) SetExdPolicyScreenRecordeClass(v *UpdateComputerPolicyRequestExdPolicyScreenRecordeClass) *UpdateComputerPolicyBody {
	s.ExdPolicyScreenRecordeClass = v
	return s
}

func (s *UpdateComputerPolicyBody) SetAudioPlayBackMode(v string) *UpdateComputerPolicyBody {
	s.AudioPlayBackMode = &v
	return s
}

func (s *UpdateComputerPolicyBody) SetAutoLock(v string) *UpdateComputerPolicyBody {
	s.AutoLock = &v
	return s
}

func (s *UpdateComputerPolicyBody) SetNetworkPolicys(v []UpdateComputerPolicyRequestNetworkPolicys) *UpdateComputerPolicyBody {
	s.NetworkPolicys = &v
	return s
}

func (s *UpdateComputerPolicyBody) SetNetworkScannerEnable(v string) *UpdateComputerPolicyBody {
	s.NetworkScannerEnable = &v
	return s
}

func (s *UpdateComputerPolicyBody) SetRedirectParallelport(v string) *UpdateComputerPolicyBody {
	s.RedirectParallelport = &v
	return s
}

func (s *UpdateComputerPolicyBody) SetExdCloudSideNetRedirectClass(v *UpdateComputerPolicyRequestExdCloudSideNetRedirectClass) *UpdateComputerPolicyBody {
	s.ExdCloudSideNetRedirectClass = v
	return s
}

func (s *UpdateComputerPolicyBody) SetRedirectClipboard(v string) *UpdateComputerPolicyBody {
	s.RedirectClipboard = &v
	return s
}

func (s *UpdateComputerPolicyBody) SetNetworkPrinterEnable(v string) *UpdateComputerPolicyBody {
	s.NetworkPrinterEnable = &v
	return s
}

func (s *UpdateComputerPolicyBody) SetRedirectLocalDiskScope(v *UpdateComputerPolicyRequestRedirectLocalDiskScope) *UpdateComputerPolicyBody {
	s.RedirectLocalDiskScope = v
	return s
}

func (s *UpdateComputerPolicyBody) SetRedirectCustom(v string) *UpdateComputerPolicyBody {
	s.RedirectCustom = &v
	return s
}

func (s *UpdateComputerPolicyBody) SetWatermarkClass(v *UpdateComputerPolicyRequestWatermarkClass) *UpdateComputerPolicyBody {
	s.WatermarkClass = v
	return s
}

func (s *UpdateComputerPolicyBody) SetScreenShotEnable(v string) *UpdateComputerPolicyBody {
	s.ScreenShotEnable = &v
	return s
}

func (s *UpdateComputerPolicyBody) SetDiskBandwidth(v int32) *UpdateComputerPolicyBody {
	s.DiskBandwidth = &v
	return s
}

func (s *UpdateComputerPolicyBody) SetDesktopStreamRat(v int32) *UpdateComputerPolicyBody {
	s.DesktopStreamRat = &v
	return s
}

func (s *UpdateComputerPolicyBody) SetPolicyRemark(v string) *UpdateComputerPolicyBody {
	s.PolicyRemark = &v
	return s
}

func (s *UpdateComputerPolicyBody) SetExdPolicyCustomWallpaperClass(v *UpdateComputerPolicyRequestExdPolicyCustomWallpaperClass) *UpdateComputerPolicyBody {
	s.ExdPolicyCustomWallpaperClass = v
	return s
}

func (s *UpdateComputerPolicyBody) SetUsbBandwidth(v int32) *UpdateComputerPolicyBody {
	s.UsbBandwidth = &v
	return s
}

func (s *UpdateComputerPolicyBody) SetAudioPlayBackModeEnable(v string) *UpdateComputerPolicyBody {
	s.AudioPlayBackModeEnable = &v
	return s
}

func (s *UpdateComputerPolicyBody) SetPrinterEnable(v string) *UpdateComputerPolicyBody {
	s.PrinterEnable = &v
	return s
}

func (s *UpdateComputerPolicyBody) SetQosUid(v string) *UpdateComputerPolicyBody {
	s.QosUid = &v
	return s
}

func (s *UpdateComputerPolicyBody) SetRedirectSerialport(v string) *UpdateComputerPolicyBody {
	s.RedirectSerialport = &v
	return s
}

func (s *UpdateComputerPolicyBody) SetComputerProtocol(v string) *UpdateComputerPolicyBody {
	s.ComputerProtocol = &v
	return s
}

func (s *UpdateComputerPolicyBody) SetPolicyName(v string) *UpdateComputerPolicyBody {
	s.PolicyName = &v
	return s
}

func (s *UpdateComputerPolicyBody) SetScannerEnable(v string) *UpdateComputerPolicyBody {
	s.ScannerEnable = &v
	return s
}

func (s *UpdateComputerPolicyBody) SetExdPolicyExts(v []UpdateComputerPolicyRequestExdPolicyExts) *UpdateComputerPolicyBody {
	s.ExdPolicyExts = &v
	return s
}

func (s *UpdateComputerPolicyBody) SetRedirectUsbStorage(v string) *UpdateComputerPolicyBody {
	s.RedirectUsbStorage = &v
	return s
}

func (s *UpdateComputerPolicyBody) SetPolicyId(v string) *UpdateComputerPolicyBody {
	s.PolicyId = &v
	return s
}

func (s *UpdateComputerPolicyBody) SetAutoDesktopQuality(v int32) *UpdateComputerPolicyBody {
	s.AutoDesktopQuality = &v
	return s
}

func (s *UpdateComputerPolicyBody) SetPoolCode(v string) *UpdateComputerPolicyBody {
	s.PoolCode = &v
	return s
}

func (s *UpdateComputerPolicyBody) SetBandwidthLimitEnable(v string) *UpdateComputerPolicyBody {
	s.BandwidthLimitEnable = &v
	return s
}

func (s *UpdateComputerPolicyBody) SetRegion(v string) *UpdateComputerPolicyBody {
	s.Region = &v
	return s
}

func (s *UpdateComputerPolicyBody) SetTerminalLoginPolicys(v []UpdateComputerPolicyRequestTerminalLoginPolicys) *UpdateComputerPolicyBody {
	s.TerminalLoginPolicys = &v
	return s
}

func (s *UpdateComputerPolicyBody) SetDesktopToolbarPolicys(v []UpdateComputerPolicyRequestDesktopToolbarPolicys) *UpdateComputerPolicyBody {
	s.DesktopToolbarPolicys = &v
	return s
}

func (s *UpdateComputerPolicyBody) SetSecurityGroupUid(v string) *UpdateComputerPolicyBody {
	s.SecurityGroupUid = &v
	return s
}

func (s *UpdateComputerPolicyBody) SetDc(v string) *UpdateComputerPolicyBody {
	s.Dc = &v
	return s
}

func (s *UpdateComputerPolicyBody) SetStreamingVideo(v int32) *UpdateComputerPolicyBody {
	s.StreamingVideo = &v
	return s
}


type UpdateComputerPolicyBodyBuilder struct {
	s *UpdateComputerPolicyBody
}

func NewUpdateComputerPolicyBodyBuilder() *UpdateComputerPolicyBodyBuilder {
	s := &UpdateComputerPolicyBody{}
	b := &UpdateComputerPolicyBodyBuilder{s: s}
	return b
}

func (b *UpdateComputerPolicyBodyBuilder) RedirectCamera(v string) *UpdateComputerPolicyBodyBuilder {
	b.s.RedirectCamera = &v
	return b
}

func (b *UpdateComputerPolicyBodyBuilder) ClipboardBandwidth(v int32) *UpdateComputerPolicyBodyBuilder {
	b.s.ClipboardBandwidth = &v
	return b
}

func (b *UpdateComputerPolicyBodyBuilder) ExdPolicyScreenRecordeClass(v *UpdateComputerPolicyRequestExdPolicyScreenRecordeClass) *UpdateComputerPolicyBodyBuilder {
	b.s.ExdPolicyScreenRecordeClass = v
	return b
}

func (b *UpdateComputerPolicyBodyBuilder) AudioPlayBackMode(v string) *UpdateComputerPolicyBodyBuilder {
	b.s.AudioPlayBackMode = &v
	return b
}

func (b *UpdateComputerPolicyBodyBuilder) AutoLock(v string) *UpdateComputerPolicyBodyBuilder {
	b.s.AutoLock = &v
	return b
}

func (b *UpdateComputerPolicyBodyBuilder) NetworkPolicys(v []UpdateComputerPolicyRequestNetworkPolicys) *UpdateComputerPolicyBodyBuilder {
	b.s.NetworkPolicys = &v
	return b
}

func (b *UpdateComputerPolicyBodyBuilder) NetworkScannerEnable(v string) *UpdateComputerPolicyBodyBuilder {
	b.s.NetworkScannerEnable = &v
	return b
}

func (b *UpdateComputerPolicyBodyBuilder) RedirectParallelport(v string) *UpdateComputerPolicyBodyBuilder {
	b.s.RedirectParallelport = &v
	return b
}

func (b *UpdateComputerPolicyBodyBuilder) ExdCloudSideNetRedirectClass(v *UpdateComputerPolicyRequestExdCloudSideNetRedirectClass) *UpdateComputerPolicyBodyBuilder {
	b.s.ExdCloudSideNetRedirectClass = v
	return b
}

func (b *UpdateComputerPolicyBodyBuilder) RedirectClipboard(v string) *UpdateComputerPolicyBodyBuilder {
	b.s.RedirectClipboard = &v
	return b
}

func (b *UpdateComputerPolicyBodyBuilder) NetworkPrinterEnable(v string) *UpdateComputerPolicyBodyBuilder {
	b.s.NetworkPrinterEnable = &v
	return b
}

func (b *UpdateComputerPolicyBodyBuilder) RedirectLocalDiskScope(v *UpdateComputerPolicyRequestRedirectLocalDiskScope) *UpdateComputerPolicyBodyBuilder {
	b.s.RedirectLocalDiskScope = v
	return b
}

func (b *UpdateComputerPolicyBodyBuilder) RedirectCustom(v string) *UpdateComputerPolicyBodyBuilder {
	b.s.RedirectCustom = &v
	return b
}

func (b *UpdateComputerPolicyBodyBuilder) WatermarkClass(v *UpdateComputerPolicyRequestWatermarkClass) *UpdateComputerPolicyBodyBuilder {
	b.s.WatermarkClass = v
	return b
}

func (b *UpdateComputerPolicyBodyBuilder) ScreenShotEnable(v string) *UpdateComputerPolicyBodyBuilder {
	b.s.ScreenShotEnable = &v
	return b
}

func (b *UpdateComputerPolicyBodyBuilder) DiskBandwidth(v int32) *UpdateComputerPolicyBodyBuilder {
	b.s.DiskBandwidth = &v
	return b
}

func (b *UpdateComputerPolicyBodyBuilder) DesktopStreamRat(v int32) *UpdateComputerPolicyBodyBuilder {
	b.s.DesktopStreamRat = &v
	return b
}

func (b *UpdateComputerPolicyBodyBuilder) PolicyRemark(v string) *UpdateComputerPolicyBodyBuilder {
	b.s.PolicyRemark = &v
	return b
}

func (b *UpdateComputerPolicyBodyBuilder) ExdPolicyCustomWallpaperClass(v *UpdateComputerPolicyRequestExdPolicyCustomWallpaperClass) *UpdateComputerPolicyBodyBuilder {
	b.s.ExdPolicyCustomWallpaperClass = v
	return b
}

func (b *UpdateComputerPolicyBodyBuilder) UsbBandwidth(v int32) *UpdateComputerPolicyBodyBuilder {
	b.s.UsbBandwidth = &v
	return b
}

func (b *UpdateComputerPolicyBodyBuilder) AudioPlayBackModeEnable(v string) *UpdateComputerPolicyBodyBuilder {
	b.s.AudioPlayBackModeEnable = &v
	return b
}

func (b *UpdateComputerPolicyBodyBuilder) PrinterEnable(v string) *UpdateComputerPolicyBodyBuilder {
	b.s.PrinterEnable = &v
	return b
}

func (b *UpdateComputerPolicyBodyBuilder) QosUid(v string) *UpdateComputerPolicyBodyBuilder {
	b.s.QosUid = &v
	return b
}

func (b *UpdateComputerPolicyBodyBuilder) RedirectSerialport(v string) *UpdateComputerPolicyBodyBuilder {
	b.s.RedirectSerialport = &v
	return b
}

func (b *UpdateComputerPolicyBodyBuilder) ComputerProtocol(v string) *UpdateComputerPolicyBodyBuilder {
	b.s.ComputerProtocol = &v
	return b
}

func (b *UpdateComputerPolicyBodyBuilder) PolicyName(v string) *UpdateComputerPolicyBodyBuilder {
	b.s.PolicyName = &v
	return b
}

func (b *UpdateComputerPolicyBodyBuilder) ScannerEnable(v string) *UpdateComputerPolicyBodyBuilder {
	b.s.ScannerEnable = &v
	return b
}

func (b *UpdateComputerPolicyBodyBuilder) ExdPolicyExts(v []UpdateComputerPolicyRequestExdPolicyExts) *UpdateComputerPolicyBodyBuilder {
	b.s.ExdPolicyExts = &v
	return b
}

func (b *UpdateComputerPolicyBodyBuilder) RedirectUsbStorage(v string) *UpdateComputerPolicyBodyBuilder {
	b.s.RedirectUsbStorage = &v
	return b
}

func (b *UpdateComputerPolicyBodyBuilder) PolicyId(v string) *UpdateComputerPolicyBodyBuilder {
	b.s.PolicyId = &v
	return b
}

func (b *UpdateComputerPolicyBodyBuilder) AutoDesktopQuality(v int32) *UpdateComputerPolicyBodyBuilder {
	b.s.AutoDesktopQuality = &v
	return b
}

func (b *UpdateComputerPolicyBodyBuilder) PoolCode(v string) *UpdateComputerPolicyBodyBuilder {
	b.s.PoolCode = &v
	return b
}

func (b *UpdateComputerPolicyBodyBuilder) BandwidthLimitEnable(v string) *UpdateComputerPolicyBodyBuilder {
	b.s.BandwidthLimitEnable = &v
	return b
}

func (b *UpdateComputerPolicyBodyBuilder) Region(v string) *UpdateComputerPolicyBodyBuilder {
	b.s.Region = &v
	return b
}

func (b *UpdateComputerPolicyBodyBuilder) TerminalLoginPolicys(v []UpdateComputerPolicyRequestTerminalLoginPolicys) *UpdateComputerPolicyBodyBuilder {
	b.s.TerminalLoginPolicys = &v
	return b
}

func (b *UpdateComputerPolicyBodyBuilder) DesktopToolbarPolicys(v []UpdateComputerPolicyRequestDesktopToolbarPolicys) *UpdateComputerPolicyBodyBuilder {
	b.s.DesktopToolbarPolicys = &v
	return b
}

func (b *UpdateComputerPolicyBodyBuilder) SecurityGroupUid(v string) *UpdateComputerPolicyBodyBuilder {
	b.s.SecurityGroupUid = &v
	return b
}

func (b *UpdateComputerPolicyBodyBuilder) Dc(v string) *UpdateComputerPolicyBodyBuilder {
	b.s.Dc = &v
	return b
}

func (b *UpdateComputerPolicyBodyBuilder) StreamingVideo(v int32) *UpdateComputerPolicyBodyBuilder {
	b.s.StreamingVideo = &v
	return b
}

func (b *UpdateComputerPolicyBodyBuilder) Build() *UpdateComputerPolicyBody {
	return b.s
}


