// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type NatChangeBody struct {
    position.Body
    // 实例id，示例值：CCA-XXX
	InstanceId *string `json:"instanceId,omitempty"`
    // 变更后的规格编码，0:小型，1:中型，2:大型，3:超大型
	SpecificationCode *string `json:"specificationCode,omitempty"`
}

func (s NatChangeBody) String() string {
	return utils.Beautify(s)
}

func (s NatChangeBody) GoString() string {
	return s.String()
}

func (s NatChangeBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *NatChangeBody) SetInstanceId(v string) *NatChangeBody {
	s.InstanceId = &v
	return s
}

func (s *NatChangeBody) SetSpecificationCode(v string) *NatChangeBody {
	s.SpecificationCode = &v
	return s
}


type NatChangeBodyBuilder struct {
	s *NatChangeBody
}

func NewNatChangeBodyBuilder() *NatChangeBodyBuilder {
	s := &NatChangeBody{}
	b := &NatChangeBodyBuilder{s: s}
	return b
}

func (b *NatChangeBodyBuilder) InstanceId(v string) *NatChangeBodyBuilder {
	b.s.InstanceId = &v
	return b
}

func (b *NatChangeBodyBuilder) SpecificationCode(v string) *NatChangeBodyBuilder {
	b.s.SpecificationCode = &v
	return b
}

func (b *NatChangeBodyBuilder) Build() *NatChangeBody {
	return b.s
}


