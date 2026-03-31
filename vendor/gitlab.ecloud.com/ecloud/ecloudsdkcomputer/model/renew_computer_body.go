// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type RenewComputerBody struct {
    position.Body
    // 续订时长，包月：1-60，包年：1-5
	Duration *string `json:"duration,omitempty"`
    // 资源id列表，示例值：[CCA-XXX1,CCA-XXX2]
	InstanceIdList []string `json:"instanceIdList,omitempty"`
    // 计费方式，包年year，包月month。目前支持普通续订（续订前后计费方式一致）及包年转包月续订。根据此值判断包年转包月续订还是普通续订，不传则默认普通续订。不支持包月转包年续订，不支持批量包年转包月续订
	MeasureUnit *string `json:"measureUnit,omitempty"`
}

func (s RenewComputerBody) String() string {
	return utils.Beautify(s)
}

func (s RenewComputerBody) GoString() string {
	return s.String()
}

func (s RenewComputerBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *RenewComputerBody) SetDuration(v string) *RenewComputerBody {
	s.Duration = &v
	return s
}

func (s *RenewComputerBody) SetInstanceIdList(v []string) *RenewComputerBody {
	s.InstanceIdList = v
	return s
}

func (s *RenewComputerBody) SetMeasureUnit(v string) *RenewComputerBody {
	s.MeasureUnit = &v
	return s
}


type RenewComputerBodyBuilder struct {
	s *RenewComputerBody
}

func NewRenewComputerBodyBuilder() *RenewComputerBodyBuilder {
	s := &RenewComputerBody{}
	b := &RenewComputerBodyBuilder{s: s}
	return b
}

func (b *RenewComputerBodyBuilder) Duration(v string) *RenewComputerBodyBuilder {
	b.s.Duration = &v
	return b
}

func (b *RenewComputerBodyBuilder) InstanceIdList(v []string) *RenewComputerBodyBuilder {
	b.s.InstanceIdList = v
	return b
}

func (b *RenewComputerBodyBuilder) MeasureUnit(v string) *RenewComputerBodyBuilder {
	b.s.MeasureUnit = &v
	return b
}

func (b *RenewComputerBodyBuilder) Build() *RenewComputerBody {
	return b.s
}


