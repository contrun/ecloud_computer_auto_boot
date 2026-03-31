// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type CreateImageResponseBody struct {

    // 云电脑ID
	MachineId *string `json:"machineId,omitempty"`
    // 镜像模板ID，uuid
	TemplateId *string `json:"templateId,omitempty"`
}

func (s CreateImageResponseBody) String() string {
	return utils.Beautify(s)
}

func (s CreateImageResponseBody) GoString() string {
	return s.String()
}

func (s CreateImageResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *CreateImageResponseBody) SetMachineId(v string) *CreateImageResponseBody {
	s.MachineId = &v
	return s
}

func (s *CreateImageResponseBody) SetTemplateId(v string) *CreateImageResponseBody {
	s.TemplateId = &v
	return s
}


type CreateImageResponseBodyBuilder struct {
	s *CreateImageResponseBody
}

func NewCreateImageResponseBodyBuilder() *CreateImageResponseBodyBuilder {
	s := &CreateImageResponseBody{}
	b := &CreateImageResponseBodyBuilder{s: s}
	return b
}

func (b *CreateImageResponseBodyBuilder) MachineId(v string) *CreateImageResponseBodyBuilder {
	b.s.MachineId = &v
	return b
}

func (b *CreateImageResponseBodyBuilder) TemplateId(v string) *CreateImageResponseBodyBuilder {
	b.s.TemplateId = &v
	return b
}

func (b *CreateImageResponseBodyBuilder) Build() *CreateImageResponseBody {
	return b.s
}


