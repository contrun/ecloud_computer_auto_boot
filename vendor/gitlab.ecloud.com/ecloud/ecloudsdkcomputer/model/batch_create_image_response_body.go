// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BatchCreateImageResponseBody struct {

    // 云电脑id
	MachineId *string `json:"machineId,omitempty"`
    // 镜像模板ID
	TemplateId *string `json:"templateId,omitempty"`
}

func (s BatchCreateImageResponseBody) String() string {
	return utils.Beautify(s)
}

func (s BatchCreateImageResponseBody) GoString() string {
	return s.String()
}

func (s BatchCreateImageResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BatchCreateImageResponseBody) SetMachineId(v string) *BatchCreateImageResponseBody {
	s.MachineId = &v
	return s
}

func (s *BatchCreateImageResponseBody) SetTemplateId(v string) *BatchCreateImageResponseBody {
	s.TemplateId = &v
	return s
}


type BatchCreateImageResponseBodyBuilder struct {
	s *BatchCreateImageResponseBody
}

func NewBatchCreateImageResponseBodyBuilder() *BatchCreateImageResponseBodyBuilder {
	s := &BatchCreateImageResponseBody{}
	b := &BatchCreateImageResponseBodyBuilder{s: s}
	return b
}

func (b *BatchCreateImageResponseBodyBuilder) MachineId(v string) *BatchCreateImageResponseBodyBuilder {
	b.s.MachineId = &v
	return b
}

func (b *BatchCreateImageResponseBodyBuilder) TemplateId(v string) *BatchCreateImageResponseBodyBuilder {
	b.s.TemplateId = &v
	return b
}

func (b *BatchCreateImageResponseBodyBuilder) Build() *BatchCreateImageResponseBody {
	return b.s
}


