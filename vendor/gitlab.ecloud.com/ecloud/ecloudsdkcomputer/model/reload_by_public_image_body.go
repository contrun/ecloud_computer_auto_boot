// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ReloadByPublicImageBody struct {
    position.Body
    // 云电脑id
	MachineIds []string `json:"machineIds,omitempty"`
    // 镜像模板id
	TemplateId *string `json:"templateId,omitempty"`
}

func (s ReloadByPublicImageBody) String() string {
	return utils.Beautify(s)
}

func (s ReloadByPublicImageBody) GoString() string {
	return s.String()
}

func (s ReloadByPublicImageBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ReloadByPublicImageBody) SetMachineIds(v []string) *ReloadByPublicImageBody {
	s.MachineIds = v
	return s
}

func (s *ReloadByPublicImageBody) SetTemplateId(v string) *ReloadByPublicImageBody {
	s.TemplateId = &v
	return s
}


type ReloadByPublicImageBodyBuilder struct {
	s *ReloadByPublicImageBody
}

func NewReloadByPublicImageBodyBuilder() *ReloadByPublicImageBodyBuilder {
	s := &ReloadByPublicImageBody{}
	b := &ReloadByPublicImageBodyBuilder{s: s}
	return b
}

func (b *ReloadByPublicImageBodyBuilder) MachineIds(v []string) *ReloadByPublicImageBodyBuilder {
	b.s.MachineIds = v
	return b
}

func (b *ReloadByPublicImageBodyBuilder) TemplateId(v string) *ReloadByPublicImageBodyBuilder {
	b.s.TemplateId = &v
	return b
}

func (b *ReloadByPublicImageBodyBuilder) Build() *ReloadByPublicImageBody {
	return b.s
}


