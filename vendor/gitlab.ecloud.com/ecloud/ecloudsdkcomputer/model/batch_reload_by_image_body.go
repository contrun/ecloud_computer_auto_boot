// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BatchReloadByImageBody struct {
    position.Body
    // 主机ID/云电脑ID
	MachineIds []string `json:"machineIds,omitempty"`
    // 镜像模板uuid
	TemplateId *string `json:"templateId,omitempty"`
}

func (s BatchReloadByImageBody) String() string {
	return utils.Beautify(s)
}

func (s BatchReloadByImageBody) GoString() string {
	return s.String()
}

func (s BatchReloadByImageBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BatchReloadByImageBody) SetMachineIds(v []string) *BatchReloadByImageBody {
	s.MachineIds = v
	return s
}

func (s *BatchReloadByImageBody) SetTemplateId(v string) *BatchReloadByImageBody {
	s.TemplateId = &v
	return s
}


type BatchReloadByImageBodyBuilder struct {
	s *BatchReloadByImageBody
}

func NewBatchReloadByImageBodyBuilder() *BatchReloadByImageBodyBuilder {
	s := &BatchReloadByImageBody{}
	b := &BatchReloadByImageBodyBuilder{s: s}
	return b
}

func (b *BatchReloadByImageBodyBuilder) MachineIds(v []string) *BatchReloadByImageBodyBuilder {
	b.s.MachineIds = v
	return b
}

func (b *BatchReloadByImageBodyBuilder) TemplateId(v string) *BatchReloadByImageBodyBuilder {
	b.s.TemplateId = &v
	return b
}

func (b *BatchReloadByImageBodyBuilder) Build() *BatchReloadByImageBody {
	return b.s
}


