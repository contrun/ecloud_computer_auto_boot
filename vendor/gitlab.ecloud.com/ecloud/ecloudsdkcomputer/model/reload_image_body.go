// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ReloadImageBody struct {
    position.Body
    // 云电脑ID
	MachineId *string `json:"machineId,omitempty"`
    // 镜像模板ID
	TemplateId *string `json:"templateId,omitempty"`
}

func (s ReloadImageBody) String() string {
	return utils.Beautify(s)
}

func (s ReloadImageBody) GoString() string {
	return s.String()
}

func (s ReloadImageBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ReloadImageBody) SetMachineId(v string) *ReloadImageBody {
	s.MachineId = &v
	return s
}

func (s *ReloadImageBody) SetTemplateId(v string) *ReloadImageBody {
	s.TemplateId = &v
	return s
}


type ReloadImageBodyBuilder struct {
	s *ReloadImageBody
}

func NewReloadImageBodyBuilder() *ReloadImageBodyBuilder {
	s := &ReloadImageBody{}
	b := &ReloadImageBodyBuilder{s: s}
	return b
}

func (b *ReloadImageBodyBuilder) MachineId(v string) *ReloadImageBodyBuilder {
	b.s.MachineId = &v
	return b
}

func (b *ReloadImageBodyBuilder) TemplateId(v string) *ReloadImageBodyBuilder {
	b.s.TemplateId = &v
	return b
}

func (b *ReloadImageBodyBuilder) Build() *ReloadImageBody {
	return b.s
}


