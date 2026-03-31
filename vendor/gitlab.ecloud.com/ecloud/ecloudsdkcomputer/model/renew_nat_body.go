// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type RenewNatBody struct {
    position.Body
    // 续订时长，包月：1-60，包年：1-5
	Duration *int32 `json:"duration,omitempty"`
    // 资源id，示例值：CCA-XXX
	InstanceId *string `json:"instanceId,omitempty"`
    // 计费方式，包年：year，包月：month
	MeasureUnit *string `json:"measureUnit,omitempty"`
}

func (s RenewNatBody) String() string {
	return utils.Beautify(s)
}

func (s RenewNatBody) GoString() string {
	return s.String()
}

func (s RenewNatBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *RenewNatBody) SetDuration(v int32) *RenewNatBody {
	s.Duration = &v
	return s
}

func (s *RenewNatBody) SetInstanceId(v string) *RenewNatBody {
	s.InstanceId = &v
	return s
}

func (s *RenewNatBody) SetMeasureUnit(v string) *RenewNatBody {
	s.MeasureUnit = &v
	return s
}


type RenewNatBodyBuilder struct {
	s *RenewNatBody
}

func NewRenewNatBodyBuilder() *RenewNatBodyBuilder {
	s := &RenewNatBody{}
	b := &RenewNatBodyBuilder{s: s}
	return b
}

func (b *RenewNatBodyBuilder) Duration(v int32) *RenewNatBodyBuilder {
	b.s.Duration = &v
	return b
}

func (b *RenewNatBodyBuilder) InstanceId(v string) *RenewNatBodyBuilder {
	b.s.InstanceId = &v
	return b
}

func (b *RenewNatBodyBuilder) MeasureUnit(v string) *RenewNatBodyBuilder {
	b.s.MeasureUnit = &v
	return b
}

func (b *RenewNatBodyBuilder) Build() *RenewNatBody {
	return b.s
}


