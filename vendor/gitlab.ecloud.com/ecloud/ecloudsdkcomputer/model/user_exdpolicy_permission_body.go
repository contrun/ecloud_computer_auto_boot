// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UserExdpolicyPermissionBody struct {
    position.Body
    // 电脑协议（V2.0:自研CMSSZTE(dc可不传)，V1.1：新华三，V1.2：中兴）
	ComputerProtocol *string `json:"computerProtocol,omitempty"`
    // 资源池编码
	PoolCode *string `json:"poolCode,omitempty"`
    // 可用区
	Region *string `json:"region,omitempty"`
    // 网络工作区
	Dc *string `json:"dc,omitempty"`
}

func (s UserExdpolicyPermissionBody) String() string {
	return utils.Beautify(s)
}

func (s UserExdpolicyPermissionBody) GoString() string {
	return s.String()
}

func (s UserExdpolicyPermissionBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UserExdpolicyPermissionBody) SetComputerProtocol(v string) *UserExdpolicyPermissionBody {
	s.ComputerProtocol = &v
	return s
}

func (s *UserExdpolicyPermissionBody) SetPoolCode(v string) *UserExdpolicyPermissionBody {
	s.PoolCode = &v
	return s
}

func (s *UserExdpolicyPermissionBody) SetRegion(v string) *UserExdpolicyPermissionBody {
	s.Region = &v
	return s
}

func (s *UserExdpolicyPermissionBody) SetDc(v string) *UserExdpolicyPermissionBody {
	s.Dc = &v
	return s
}


type UserExdpolicyPermissionBodyBuilder struct {
	s *UserExdpolicyPermissionBody
}

func NewUserExdpolicyPermissionBodyBuilder() *UserExdpolicyPermissionBodyBuilder {
	s := &UserExdpolicyPermissionBody{}
	b := &UserExdpolicyPermissionBodyBuilder{s: s}
	return b
}

func (b *UserExdpolicyPermissionBodyBuilder) ComputerProtocol(v string) *UserExdpolicyPermissionBodyBuilder {
	b.s.ComputerProtocol = &v
	return b
}

func (b *UserExdpolicyPermissionBodyBuilder) PoolCode(v string) *UserExdpolicyPermissionBodyBuilder {
	b.s.PoolCode = &v
	return b
}

func (b *UserExdpolicyPermissionBodyBuilder) Region(v string) *UserExdpolicyPermissionBodyBuilder {
	b.s.Region = &v
	return b
}

func (b *UserExdpolicyPermissionBodyBuilder) Dc(v string) *UserExdpolicyPermissionBodyBuilder {
	b.s.Dc = &v
	return b
}

func (b *UserExdpolicyPermissionBodyBuilder) Build() *UserExdpolicyPermissionBody {
	return b.s
}


