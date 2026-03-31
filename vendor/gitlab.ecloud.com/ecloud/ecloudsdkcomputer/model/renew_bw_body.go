// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type RenewBwBody struct {
    position.Body
    // 续订时长，包月：1-60，包年：1-5
	Duration *int32 `json:"duration,omitempty"`
    // 资源id，示例值：CCA-XXX
	InstanceId *string `json:"instanceId,omitempty"`
    // 计费方式，包年：year，包月：month
	MeasureUnit *string `json:"measureUnit,omitempty"`
}

func (s RenewBwBody) String() string {
	return utils.Beautify(s)
}

func (s RenewBwBody) GoString() string {
	return s.String()
}

func (s RenewBwBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *RenewBwBody) SetDuration(v int32) *RenewBwBody {
	s.Duration = &v
	return s
}

func (s *RenewBwBody) SetInstanceId(v string) *RenewBwBody {
	s.InstanceId = &v
	return s
}

func (s *RenewBwBody) SetMeasureUnit(v string) *RenewBwBody {
	s.MeasureUnit = &v
	return s
}


type RenewBwBodyBuilder struct {
	s *RenewBwBody
}

func NewRenewBwBodyBuilder() *RenewBwBodyBuilder {
	s := &RenewBwBody{}
	b := &RenewBwBodyBuilder{s: s}
	return b
}

func (b *RenewBwBodyBuilder) Duration(v int32) *RenewBwBodyBuilder {
	b.s.Duration = &v
	return b
}

func (b *RenewBwBodyBuilder) InstanceId(v string) *RenewBwBodyBuilder {
	b.s.InstanceId = &v
	return b
}

func (b *RenewBwBodyBuilder) MeasureUnit(v string) *RenewBwBodyBuilder {
	b.s.MeasureUnit = &v
	return b
}

func (b *RenewBwBodyBuilder) Build() *RenewBwBody {
	return b.s
}


