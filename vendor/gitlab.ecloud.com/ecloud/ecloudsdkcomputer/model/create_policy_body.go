// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreatePolicyBody struct {
    position.Body
    // 摄像头 关闭(disable)、可用(enable) 默认enable
	RedirectCamera *string `json:"redirectCamera,omitempty"`
    // 文件拷贝带宽控制-剪贴板范围, 500 - 2000 KB , 默认 500 KB, 增减步长 500, 单位(KB)
	ClipboardBandwidth *int32 `json:"clipboardBandwidth,omitempty"`
    // 录屏管控
	ExdPolicyScreenRecordeClass *CreatePolicyRequestExdPolicyScreenRecordeClass `json:"exdPolicyScreenRecordeClass,omitempty"`
    // 音频播放模式, 1:默认模式 2:极速模式/加速模式 3:流畅模式
	AudioPlayBackMode *string `json:"audioPlayBackMode,omitempty"`
    // 自动锁屏,单位分钟 默认锁屏时间 0 不锁屏
	AutoLock *string `json:"autoLock,omitempty"`
    // 网络控制
	NetworkPolicys *[]CreatePolicyRequestNetworkPolicys `json:"networkPolicys,omitempty"`
    // 是否开启网络扫描仪，关闭(disable)、开启（enable）, 默认关闭
	NetworkScannerEnable *string `json:"networkScannerEnable,omitempty"`
    // 并口 关闭(disable)、可用(enable) 默认 enable
	RedirectParallelport *string `json:"redirectParallelport,omitempty"`
    // 云侧网络重定向
	ExdCloudSideNetRedirectClass *CreatePolicyRequestExdCloudSideNetRedirectClass `json:"exdCloudSideNetRedirectClass,omitempty"`
    // 重定向剪贴板，禁用(disable)、云端到本地(c2l), 本地到云端(l2c), 双向(both)，默认both
	RedirectClipboard *string `json:"redirectClipboard,omitempty"`
    // 是否开启网络打印机，禁用(disable)、可用(enable)
	NetworkPrinterEnable *string `json:"networkPrinterEnable,omitempty"`
    // 本地磁盘映射作用范围
	RedirectLocalDiskScope *CreatePolicyRequestRedirectLocalDiskScope `json:"redirectLocalDiskScope,omitempty"`
    // 自定义USB重定向策略配置 关闭(disable)、可用(enable) 默认 disable
	RedirectCustom *string `json:"redirectCustom,omitempty"`
    // 水印详情
	WatermarkClass *CreatePolicyRequestWatermarkClass `json:"watermarkClass,omitempty"`
    // 防截屏 禁用(disable)、可用(enable) 默认 enable
	ScreenShotEnable *string `json:"screenShotEnable,omitempty"`
    // 磁盘映射范围, 500 - 2000 KB , 默认 500 KB, 增减步长 500, 单位(KB)
	DiskBandwidth *int32 `json:"diskBandwidth,omitempty"`
    // 图像上限帧率 默认30 最小值12 最大值100
	DesktopStreamRat *int32 `json:"desktopStreamRat,omitempty"`
    // 策略描述
	PolicyRemark *string `json:"policyRemark,omitempty"`
    // 自定义电脑壁纸
	ExdPolicyCustomWallpaperClass *CreatePolicyRequestExdPolicyCustomWallpaperClass `json:"exdPolicyCustomWallpaperClass,omitempty"`
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
	ExdPolicyExts *[]CreatePolicyRequestExdPolicyExts `json:"exdPolicyExts,omitempty"`
    // 重定向USB存储，禁用(disable),读写(enable),只读(ro),默认enable
	RedirectUsbStorage *string `json:"redirectUsbStorage,omitempty"`
    // 图像画质 默认 2 最小值0 最大值3
	AutoDesktopQuality *int32 `json:"autoDesktopQuality,omitempty"`
    // 所属资源池
	PoolCode *string `json:"poolCode,omitempty"`
    // 带宽是否启用 关闭(disable)、可用(enable) 默认 disable
	BandwidthLimitEnable *string `json:"bandwidthLimitEnable,omitempty"`
    // 所属可用区
	Region *string `json:"region,omitempty"`
    // 终端登录策略
	TerminalLoginPolicys *[]CreatePolicyRequestTerminalLoginPolicys `json:"terminalLoginPolicys,omitempty"`
    // 桌面工具栏策略
	DesktopToolbarPolicys *[]CreatePolicyRequestDesktopToolbarPolicys `json:"desktopToolbarPolicys,omitempty"`
    // 安全组UID
	SecurityGroupUid *string `json:"securityGroupUid,omitempty"`
    // 网络工作区，V2.0协议非必传
	Dc *string `json:"dc,omitempty"`
    // 图流设置 默认 2 最小值1 最大值2
	StreamingVideo *int32 `json:"streamingVideo,omitempty"`
}

func (s CreatePolicyBody) String() string {
	return utils.Beautify(s)
}

func (s CreatePolicyBody) GoString() string {
	return s.String()
}

func (s CreatePolicyBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreatePolicyBody) SetRedirectCamera(v string) *CreatePolicyBody {
	s.RedirectCamera = &v
	return s
}

func (s *CreatePolicyBody) SetClipboardBandwidth(v int32) *CreatePolicyBody {
	s.ClipboardBandwidth = &v
	return s
}

func (s *CreatePolicyBody) SetExdPolicyScreenRecordeClass(v *CreatePolicyRequestExdPolicyScreenRecordeClass) *CreatePolicyBody {
	s.ExdPolicyScreenRecordeClass = v
	return s
}

func (s *CreatePolicyBody) SetAudioPlayBackMode(v string) *CreatePolicyBody {
	s.AudioPlayBackMode = &v
	return s
}

func (s *CreatePolicyBody) SetAutoLock(v string) *CreatePolicyBody {
	s.AutoLock = &v
	return s
}

func (s *CreatePolicyBody) SetNetworkPolicys(v []CreatePolicyRequestNetworkPolicys) *CreatePolicyBody {
	s.NetworkPolicys = &v
	return s
}

func (s *CreatePolicyBody) SetNetworkScannerEnable(v string) *CreatePolicyBody {
	s.NetworkScannerEnable = &v
	return s
}

func (s *CreatePolicyBody) SetRedirectParallelport(v string) *CreatePolicyBody {
	s.RedirectParallelport = &v
	return s
}

func (s *CreatePolicyBody) SetExdCloudSideNetRedirectClass(v *CreatePolicyRequestExdCloudSideNetRedirectClass) *CreatePolicyBody {
	s.ExdCloudSideNetRedirectClass = v
	return s
}

func (s *CreatePolicyBody) SetRedirectClipboard(v string) *CreatePolicyBody {
	s.RedirectClipboard = &v
	return s
}

func (s *CreatePolicyBody) SetNetworkPrinterEnable(v string) *CreatePolicyBody {
	s.NetworkPrinterEnable = &v
	return s
}

func (s *CreatePolicyBody) SetRedirectLocalDiskScope(v *CreatePolicyRequestRedirectLocalDiskScope) *CreatePolicyBody {
	s.RedirectLocalDiskScope = v
	return s
}

func (s *CreatePolicyBody) SetRedirectCustom(v string) *CreatePolicyBody {
	s.RedirectCustom = &v
	return s
}

func (s *CreatePolicyBody) SetWatermarkClass(v *CreatePolicyRequestWatermarkClass) *CreatePolicyBody {
	s.WatermarkClass = v
	return s
}

func (s *CreatePolicyBody) SetScreenShotEnable(v string) *CreatePolicyBody {
	s.ScreenShotEnable = &v
	return s
}

func (s *CreatePolicyBody) SetDiskBandwidth(v int32) *CreatePolicyBody {
	s.DiskBandwidth = &v
	return s
}

func (s *CreatePolicyBody) SetDesktopStreamRat(v int32) *CreatePolicyBody {
	s.DesktopStreamRat = &v
	return s
}

func (s *CreatePolicyBody) SetPolicyRemark(v string) *CreatePolicyBody {
	s.PolicyRemark = &v
	return s
}

func (s *CreatePolicyBody) SetExdPolicyCustomWallpaperClass(v *CreatePolicyRequestExdPolicyCustomWallpaperClass) *CreatePolicyBody {
	s.ExdPolicyCustomWallpaperClass = v
	return s
}

func (s *CreatePolicyBody) SetUsbBandwidth(v int32) *CreatePolicyBody {
	s.UsbBandwidth = &v
	return s
}

func (s *CreatePolicyBody) SetAudioPlayBackModeEnable(v string) *CreatePolicyBody {
	s.AudioPlayBackModeEnable = &v
	return s
}

func (s *CreatePolicyBody) SetPrinterEnable(v string) *CreatePolicyBody {
	s.PrinterEnable = &v
	return s
}

func (s *CreatePolicyBody) SetQosUid(v string) *CreatePolicyBody {
	s.QosUid = &v
	return s
}

func (s *CreatePolicyBody) SetRedirectSerialport(v string) *CreatePolicyBody {
	s.RedirectSerialport = &v
	return s
}

func (s *CreatePolicyBody) SetComputerProtocol(v string) *CreatePolicyBody {
	s.ComputerProtocol = &v
	return s
}

func (s *CreatePolicyBody) SetPolicyName(v string) *CreatePolicyBody {
	s.PolicyName = &v
	return s
}

func (s *CreatePolicyBody) SetScannerEnable(v string) *CreatePolicyBody {
	s.ScannerEnable = &v
	return s
}

func (s *CreatePolicyBody) SetExdPolicyExts(v []CreatePolicyRequestExdPolicyExts) *CreatePolicyBody {
	s.ExdPolicyExts = &v
	return s
}

func (s *CreatePolicyBody) SetRedirectUsbStorage(v string) *CreatePolicyBody {
	s.RedirectUsbStorage = &v
	return s
}

func (s *CreatePolicyBody) SetAutoDesktopQuality(v int32) *CreatePolicyBody {
	s.AutoDesktopQuality = &v
	return s
}

func (s *CreatePolicyBody) SetPoolCode(v string) *CreatePolicyBody {
	s.PoolCode = &v
	return s
}

func (s *CreatePolicyBody) SetBandwidthLimitEnable(v string) *CreatePolicyBody {
	s.BandwidthLimitEnable = &v
	return s
}

func (s *CreatePolicyBody) SetRegion(v string) *CreatePolicyBody {
	s.Region = &v
	return s
}

func (s *CreatePolicyBody) SetTerminalLoginPolicys(v []CreatePolicyRequestTerminalLoginPolicys) *CreatePolicyBody {
	s.TerminalLoginPolicys = &v
	return s
}

func (s *CreatePolicyBody) SetDesktopToolbarPolicys(v []CreatePolicyRequestDesktopToolbarPolicys) *CreatePolicyBody {
	s.DesktopToolbarPolicys = &v
	return s
}

func (s *CreatePolicyBody) SetSecurityGroupUid(v string) *CreatePolicyBody {
	s.SecurityGroupUid = &v
	return s
}

func (s *CreatePolicyBody) SetDc(v string) *CreatePolicyBody {
	s.Dc = &v
	return s
}

func (s *CreatePolicyBody) SetStreamingVideo(v int32) *CreatePolicyBody {
	s.StreamingVideo = &v
	return s
}


type CreatePolicyBodyBuilder struct {
	s *CreatePolicyBody
}

func NewCreatePolicyBodyBuilder() *CreatePolicyBodyBuilder {
	s := &CreatePolicyBody{}
	b := &CreatePolicyBodyBuilder{s: s}
	return b
}

func (b *CreatePolicyBodyBuilder) RedirectCamera(v string) *CreatePolicyBodyBuilder {
	b.s.RedirectCamera = &v
	return b
}

func (b *CreatePolicyBodyBuilder) ClipboardBandwidth(v int32) *CreatePolicyBodyBuilder {
	b.s.ClipboardBandwidth = &v
	return b
}

func (b *CreatePolicyBodyBuilder) ExdPolicyScreenRecordeClass(v *CreatePolicyRequestExdPolicyScreenRecordeClass) *CreatePolicyBodyBuilder {
	b.s.ExdPolicyScreenRecordeClass = v
	return b
}

func (b *CreatePolicyBodyBuilder) AudioPlayBackMode(v string) *CreatePolicyBodyBuilder {
	b.s.AudioPlayBackMode = &v
	return b
}

func (b *CreatePolicyBodyBuilder) AutoLock(v string) *CreatePolicyBodyBuilder {
	b.s.AutoLock = &v
	return b
}

func (b *CreatePolicyBodyBuilder) NetworkPolicys(v []CreatePolicyRequestNetworkPolicys) *CreatePolicyBodyBuilder {
	b.s.NetworkPolicys = &v
	return b
}

func (b *CreatePolicyBodyBuilder) NetworkScannerEnable(v string) *CreatePolicyBodyBuilder {
	b.s.NetworkScannerEnable = &v
	return b
}

func (b *CreatePolicyBodyBuilder) RedirectParallelport(v string) *CreatePolicyBodyBuilder {
	b.s.RedirectParallelport = &v
	return b
}

func (b *CreatePolicyBodyBuilder) ExdCloudSideNetRedirectClass(v *CreatePolicyRequestExdCloudSideNetRedirectClass) *CreatePolicyBodyBuilder {
	b.s.ExdCloudSideNetRedirectClass = v
	return b
}

func (b *CreatePolicyBodyBuilder) RedirectClipboard(v string) *CreatePolicyBodyBuilder {
	b.s.RedirectClipboard = &v
	return b
}

func (b *CreatePolicyBodyBuilder) NetworkPrinterEnable(v string) *CreatePolicyBodyBuilder {
	b.s.NetworkPrinterEnable = &v
	return b
}

func (b *CreatePolicyBodyBuilder) RedirectLocalDiskScope(v *CreatePolicyRequestRedirectLocalDiskScope) *CreatePolicyBodyBuilder {
	b.s.RedirectLocalDiskScope = v
	return b
}

func (b *CreatePolicyBodyBuilder) RedirectCustom(v string) *CreatePolicyBodyBuilder {
	b.s.RedirectCustom = &v
	return b
}

func (b *CreatePolicyBodyBuilder) WatermarkClass(v *CreatePolicyRequestWatermarkClass) *CreatePolicyBodyBuilder {
	b.s.WatermarkClass = v
	return b
}

func (b *CreatePolicyBodyBuilder) ScreenShotEnable(v string) *CreatePolicyBodyBuilder {
	b.s.ScreenShotEnable = &v
	return b
}

func (b *CreatePolicyBodyBuilder) DiskBandwidth(v int32) *CreatePolicyBodyBuilder {
	b.s.DiskBandwidth = &v
	return b
}

func (b *CreatePolicyBodyBuilder) DesktopStreamRat(v int32) *CreatePolicyBodyBuilder {
	b.s.DesktopStreamRat = &v
	return b
}

func (b *CreatePolicyBodyBuilder) PolicyRemark(v string) *CreatePolicyBodyBuilder {
	b.s.PolicyRemark = &v
	return b
}

func (b *CreatePolicyBodyBuilder) ExdPolicyCustomWallpaperClass(v *CreatePolicyRequestExdPolicyCustomWallpaperClass) *CreatePolicyBodyBuilder {
	b.s.ExdPolicyCustomWallpaperClass = v
	return b
}

func (b *CreatePolicyBodyBuilder) UsbBandwidth(v int32) *CreatePolicyBodyBuilder {
	b.s.UsbBandwidth = &v
	return b
}

func (b *CreatePolicyBodyBuilder) AudioPlayBackModeEnable(v string) *CreatePolicyBodyBuilder {
	b.s.AudioPlayBackModeEnable = &v
	return b
}

func (b *CreatePolicyBodyBuilder) PrinterEnable(v string) *CreatePolicyBodyBuilder {
	b.s.PrinterEnable = &v
	return b
}

func (b *CreatePolicyBodyBuilder) QosUid(v string) *CreatePolicyBodyBuilder {
	b.s.QosUid = &v
	return b
}

func (b *CreatePolicyBodyBuilder) RedirectSerialport(v string) *CreatePolicyBodyBuilder {
	b.s.RedirectSerialport = &v
	return b
}

func (b *CreatePolicyBodyBuilder) ComputerProtocol(v string) *CreatePolicyBodyBuilder {
	b.s.ComputerProtocol = &v
	return b
}

func (b *CreatePolicyBodyBuilder) PolicyName(v string) *CreatePolicyBodyBuilder {
	b.s.PolicyName = &v
	return b
}

func (b *CreatePolicyBodyBuilder) ScannerEnable(v string) *CreatePolicyBodyBuilder {
	b.s.ScannerEnable = &v
	return b
}

func (b *CreatePolicyBodyBuilder) ExdPolicyExts(v []CreatePolicyRequestExdPolicyExts) *CreatePolicyBodyBuilder {
	b.s.ExdPolicyExts = &v
	return b
}

func (b *CreatePolicyBodyBuilder) RedirectUsbStorage(v string) *CreatePolicyBodyBuilder {
	b.s.RedirectUsbStorage = &v
	return b
}

func (b *CreatePolicyBodyBuilder) AutoDesktopQuality(v int32) *CreatePolicyBodyBuilder {
	b.s.AutoDesktopQuality = &v
	return b
}

func (b *CreatePolicyBodyBuilder) PoolCode(v string) *CreatePolicyBodyBuilder {
	b.s.PoolCode = &v
	return b
}

func (b *CreatePolicyBodyBuilder) BandwidthLimitEnable(v string) *CreatePolicyBodyBuilder {
	b.s.BandwidthLimitEnable = &v
	return b
}

func (b *CreatePolicyBodyBuilder) Region(v string) *CreatePolicyBodyBuilder {
	b.s.Region = &v
	return b
}

func (b *CreatePolicyBodyBuilder) TerminalLoginPolicys(v []CreatePolicyRequestTerminalLoginPolicys) *CreatePolicyBodyBuilder {
	b.s.TerminalLoginPolicys = &v
	return b
}

func (b *CreatePolicyBodyBuilder) DesktopToolbarPolicys(v []CreatePolicyRequestDesktopToolbarPolicys) *CreatePolicyBodyBuilder {
	b.s.DesktopToolbarPolicys = &v
	return b
}

func (b *CreatePolicyBodyBuilder) SecurityGroupUid(v string) *CreatePolicyBodyBuilder {
	b.s.SecurityGroupUid = &v
	return b
}

func (b *CreatePolicyBodyBuilder) Dc(v string) *CreatePolicyBodyBuilder {
	b.s.Dc = &v
	return b
}

func (b *CreatePolicyBodyBuilder) StreamingVideo(v int32) *CreatePolicyBodyBuilder {
	b.s.StreamingVideo = &v
	return b
}

func (b *CreatePolicyBodyBuilder) Build() *CreatePolicyBody {
	return b.s
}


