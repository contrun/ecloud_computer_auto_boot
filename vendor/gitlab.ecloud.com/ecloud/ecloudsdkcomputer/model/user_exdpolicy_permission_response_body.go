// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UserExdpolicyPermissionResponseBody struct {

    // 是否可用网络打印机, true:可用, false:不可用 开启可传参数:netPrinterEnable
	NetPrinterEnable *bool `json:"netPrinterEnable,omitempty"`
    // 是否可用安全组, true:可用, false:不可用 开启可传参数:securityGroupUid
	SecurityGroupEnable *bool `json:"securityGroupEnable,omitempty"`
    // 是否可用网络扫描仪, true:可用, false:不可用 开启可传参数:netScannerEnable
	NetScannerEnable *bool `json:"netScannerEnable,omitempty"`
    // 是否可用明水印, true:可用, false:不可用 开启可传参数:明水印全参
	ObviousWatermarkEnable *bool `json:"obviousWatermarkEnable,omitempty"`
    // 是否可用扫描仪, true:可用, false:不可用 开启可传参数:scannerEnable
	ScannerEnable *bool `json:"scannerEnable,omitempty"`
    // 是否可用音频播放模式, true:可用, false:不可用 开启可传参数:audioPlayBackMode
	AudioAndVideoSwitch *bool `json:"audioAndVideoSwitch,omitempty"`
    // 是否可用Qos, true:可用, false:不可用 开启可传参数:qosUid
	QosEnable *bool `json:"qosEnable,omitempty"`
    // 是否可用重定向本地磁盘, true:可用, false:不可用 开启可传参数:redirectLocalDisk
	RedirectLocalDiskEnable *bool `json:"redirectLocalDiskEnable,omitempty"`
    // 是否可用磁盘映射作用范围, true:可用, false:不可用 开启可传参数:fixedDrive,removableDrive,opticalDiscDrive,networkDrive,terminalSystemDisk
	ScopeDiskMappingEnable *bool `json:"scopeDiskMappingEnable,omitempty"`
    // 是否可用云侧网络重定向, true:可用, false:不可用 开启可传参数:exdCloudSideNetRedirectClass
	NetRedirectSwitch *bool `json:"netRedirectSwitch,omitempty"`
    // 画面设置开关, true:展示, false:不展示 开启可传参数:streamingVideo,desktopStreamRat,autoDesktopQuality
	ScreenSettingsSwitch *bool `json:"screenSettingsSwitch,omitempty"`
    // 是否可用盲水印, true:可用, false:不可用 开启可传参数:盲水印单参
	HideWatermarkEnable *bool `json:"hideWatermarkEnable,omitempty"`
    // 是否可用USB重定向存储RW,RO权限, true:可用, false:不可用 开启可传参数:redirectUsbStorage
	RedirectUsbStorageRWEnable *bool `json:"redirectUsbStorageRWEnable,omitempty"`
    // 是否可用自定义USB重定向, true:可用, false:不可用 开启可传参数:exdPolicyExts
	RedirectCustomEnable *bool `json:"redirectCustomEnable,omitempty"`
    // 是否可用防截屏, true:可用, false:不可用 开启可传参数:screenShotEnable
	ScreenShotEnable *bool `json:"screenShotEnable,omitempty"`
    // 是否可用自定义壁纸, true:可用, false:不可用
	CustomizableWallpapersEnable *bool `json:"customizableWallpapersEnable,omitempty"`
    // 是否可用录屏, true:可用, false:不可用 开启可传参数:exdPolicyScreenRecordeClass
	ScreenRecordingSwitch *bool `json:"screenRecordingSwitch,omitempty"`
    // 是否可用重定向剪贴板, true:可用, false:不可用 开启可传参数:redirectClipboard
	RedirectClipboardEnable *bool `json:"redirectClipboardEnable,omitempty"`
    // 是否可用带宽, true:可用, false:不可用 开启可传参数:usbBandwidth,diskBandwidth,clipboardBandwidth
	BandwidthLimitEnable *bool `json:"bandwidthLimitEnable,omitempty"`
    // 是否可用打印机, true:可用, false:不可用 开启可传参数:printerEnable
	PrinterEnable *bool `json:"printerEnable,omitempty"`
}

func (s UserExdpolicyPermissionResponseBody) String() string {
	return utils.Beautify(s)
}

func (s UserExdpolicyPermissionResponseBody) GoString() string {
	return s.String()
}

func (s UserExdpolicyPermissionResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UserExdpolicyPermissionResponseBody) SetNetPrinterEnable(v bool) *UserExdpolicyPermissionResponseBody {
	s.NetPrinterEnable = &v
	return s
}

func (s *UserExdpolicyPermissionResponseBody) SetSecurityGroupEnable(v bool) *UserExdpolicyPermissionResponseBody {
	s.SecurityGroupEnable = &v
	return s
}

func (s *UserExdpolicyPermissionResponseBody) SetNetScannerEnable(v bool) *UserExdpolicyPermissionResponseBody {
	s.NetScannerEnable = &v
	return s
}

func (s *UserExdpolicyPermissionResponseBody) SetObviousWatermarkEnable(v bool) *UserExdpolicyPermissionResponseBody {
	s.ObviousWatermarkEnable = &v
	return s
}

func (s *UserExdpolicyPermissionResponseBody) SetScannerEnable(v bool) *UserExdpolicyPermissionResponseBody {
	s.ScannerEnable = &v
	return s
}

func (s *UserExdpolicyPermissionResponseBody) SetAudioAndVideoSwitch(v bool) *UserExdpolicyPermissionResponseBody {
	s.AudioAndVideoSwitch = &v
	return s
}

func (s *UserExdpolicyPermissionResponseBody) SetQosEnable(v bool) *UserExdpolicyPermissionResponseBody {
	s.QosEnable = &v
	return s
}

func (s *UserExdpolicyPermissionResponseBody) SetRedirectLocalDiskEnable(v bool) *UserExdpolicyPermissionResponseBody {
	s.RedirectLocalDiskEnable = &v
	return s
}

func (s *UserExdpolicyPermissionResponseBody) SetScopeDiskMappingEnable(v bool) *UserExdpolicyPermissionResponseBody {
	s.ScopeDiskMappingEnable = &v
	return s
}

func (s *UserExdpolicyPermissionResponseBody) SetNetRedirectSwitch(v bool) *UserExdpolicyPermissionResponseBody {
	s.NetRedirectSwitch = &v
	return s
}

func (s *UserExdpolicyPermissionResponseBody) SetScreenSettingsSwitch(v bool) *UserExdpolicyPermissionResponseBody {
	s.ScreenSettingsSwitch = &v
	return s
}

func (s *UserExdpolicyPermissionResponseBody) SetHideWatermarkEnable(v bool) *UserExdpolicyPermissionResponseBody {
	s.HideWatermarkEnable = &v
	return s
}

func (s *UserExdpolicyPermissionResponseBody) SetRedirectUsbStorageRWEnable(v bool) *UserExdpolicyPermissionResponseBody {
	s.RedirectUsbStorageRWEnable = &v
	return s
}

func (s *UserExdpolicyPermissionResponseBody) SetRedirectCustomEnable(v bool) *UserExdpolicyPermissionResponseBody {
	s.RedirectCustomEnable = &v
	return s
}

func (s *UserExdpolicyPermissionResponseBody) SetScreenShotEnable(v bool) *UserExdpolicyPermissionResponseBody {
	s.ScreenShotEnable = &v
	return s
}

func (s *UserExdpolicyPermissionResponseBody) SetCustomizableWallpapersEnable(v bool) *UserExdpolicyPermissionResponseBody {
	s.CustomizableWallpapersEnable = &v
	return s
}

func (s *UserExdpolicyPermissionResponseBody) SetScreenRecordingSwitch(v bool) *UserExdpolicyPermissionResponseBody {
	s.ScreenRecordingSwitch = &v
	return s
}

func (s *UserExdpolicyPermissionResponseBody) SetRedirectClipboardEnable(v bool) *UserExdpolicyPermissionResponseBody {
	s.RedirectClipboardEnable = &v
	return s
}

func (s *UserExdpolicyPermissionResponseBody) SetBandwidthLimitEnable(v bool) *UserExdpolicyPermissionResponseBody {
	s.BandwidthLimitEnable = &v
	return s
}

func (s *UserExdpolicyPermissionResponseBody) SetPrinterEnable(v bool) *UserExdpolicyPermissionResponseBody {
	s.PrinterEnable = &v
	return s
}


type UserExdpolicyPermissionResponseBodyBuilder struct {
	s *UserExdpolicyPermissionResponseBody
}

func NewUserExdpolicyPermissionResponseBodyBuilder() *UserExdpolicyPermissionResponseBodyBuilder {
	s := &UserExdpolicyPermissionResponseBody{}
	b := &UserExdpolicyPermissionResponseBodyBuilder{s: s}
	return b
}

func (b *UserExdpolicyPermissionResponseBodyBuilder) NetPrinterEnable(v bool) *UserExdpolicyPermissionResponseBodyBuilder {
	b.s.NetPrinterEnable = &v
	return b
}

func (b *UserExdpolicyPermissionResponseBodyBuilder) SecurityGroupEnable(v bool) *UserExdpolicyPermissionResponseBodyBuilder {
	b.s.SecurityGroupEnable = &v
	return b
}

func (b *UserExdpolicyPermissionResponseBodyBuilder) NetScannerEnable(v bool) *UserExdpolicyPermissionResponseBodyBuilder {
	b.s.NetScannerEnable = &v
	return b
}

func (b *UserExdpolicyPermissionResponseBodyBuilder) ObviousWatermarkEnable(v bool) *UserExdpolicyPermissionResponseBodyBuilder {
	b.s.ObviousWatermarkEnable = &v
	return b
}

func (b *UserExdpolicyPermissionResponseBodyBuilder) ScannerEnable(v bool) *UserExdpolicyPermissionResponseBodyBuilder {
	b.s.ScannerEnable = &v
	return b
}

func (b *UserExdpolicyPermissionResponseBodyBuilder) AudioAndVideoSwitch(v bool) *UserExdpolicyPermissionResponseBodyBuilder {
	b.s.AudioAndVideoSwitch = &v
	return b
}

func (b *UserExdpolicyPermissionResponseBodyBuilder) QosEnable(v bool) *UserExdpolicyPermissionResponseBodyBuilder {
	b.s.QosEnable = &v
	return b
}

func (b *UserExdpolicyPermissionResponseBodyBuilder) RedirectLocalDiskEnable(v bool) *UserExdpolicyPermissionResponseBodyBuilder {
	b.s.RedirectLocalDiskEnable = &v
	return b
}

func (b *UserExdpolicyPermissionResponseBodyBuilder) ScopeDiskMappingEnable(v bool) *UserExdpolicyPermissionResponseBodyBuilder {
	b.s.ScopeDiskMappingEnable = &v
	return b
}

func (b *UserExdpolicyPermissionResponseBodyBuilder) NetRedirectSwitch(v bool) *UserExdpolicyPermissionResponseBodyBuilder {
	b.s.NetRedirectSwitch = &v
	return b
}

func (b *UserExdpolicyPermissionResponseBodyBuilder) ScreenSettingsSwitch(v bool) *UserExdpolicyPermissionResponseBodyBuilder {
	b.s.ScreenSettingsSwitch = &v
	return b
}

func (b *UserExdpolicyPermissionResponseBodyBuilder) HideWatermarkEnable(v bool) *UserExdpolicyPermissionResponseBodyBuilder {
	b.s.HideWatermarkEnable = &v
	return b
}

func (b *UserExdpolicyPermissionResponseBodyBuilder) RedirectUsbStorageRWEnable(v bool) *UserExdpolicyPermissionResponseBodyBuilder {
	b.s.RedirectUsbStorageRWEnable = &v
	return b
}

func (b *UserExdpolicyPermissionResponseBodyBuilder) RedirectCustomEnable(v bool) *UserExdpolicyPermissionResponseBodyBuilder {
	b.s.RedirectCustomEnable = &v
	return b
}

func (b *UserExdpolicyPermissionResponseBodyBuilder) ScreenShotEnable(v bool) *UserExdpolicyPermissionResponseBodyBuilder {
	b.s.ScreenShotEnable = &v
	return b
}

func (b *UserExdpolicyPermissionResponseBodyBuilder) CustomizableWallpapersEnable(v bool) *UserExdpolicyPermissionResponseBodyBuilder {
	b.s.CustomizableWallpapersEnable = &v
	return b
}

func (b *UserExdpolicyPermissionResponseBodyBuilder) ScreenRecordingSwitch(v bool) *UserExdpolicyPermissionResponseBodyBuilder {
	b.s.ScreenRecordingSwitch = &v
	return b
}

func (b *UserExdpolicyPermissionResponseBodyBuilder) RedirectClipboardEnable(v bool) *UserExdpolicyPermissionResponseBodyBuilder {
	b.s.RedirectClipboardEnable = &v
	return b
}

func (b *UserExdpolicyPermissionResponseBodyBuilder) BandwidthLimitEnable(v bool) *UserExdpolicyPermissionResponseBodyBuilder {
	b.s.BandwidthLimitEnable = &v
	return b
}

func (b *UserExdpolicyPermissionResponseBodyBuilder) PrinterEnable(v bool) *UserExdpolicyPermissionResponseBodyBuilder {
	b.s.PrinterEnable = &v
	return b
}

func (b *UserExdpolicyPermissionResponseBodyBuilder) Build() *UserExdpolicyPermissionResponseBody {
	return b.s
}


