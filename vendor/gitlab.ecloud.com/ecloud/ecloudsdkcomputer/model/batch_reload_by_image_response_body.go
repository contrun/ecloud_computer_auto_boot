// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BatchReloadByImageResponseBody struct {

    // 云电脑id
	MachineId *string `json:"machineId,omitempty"`
    // 请求ID
	RequestId *string `json:"requestId,omitempty"`
    // 镜像ID
	TemplateId *string `json:"templateId,omitempty"`
    // 镜像切换记录表主键ID
	ImageSnapshotRecordIds *string `json:"imageSnapshotRecordIds,omitempty"`
}

func (s BatchReloadByImageResponseBody) String() string {
	return utils.Beautify(s)
}

func (s BatchReloadByImageResponseBody) GoString() string {
	return s.String()
}

func (s BatchReloadByImageResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BatchReloadByImageResponseBody) SetMachineId(v string) *BatchReloadByImageResponseBody {
	s.MachineId = &v
	return s
}

func (s *BatchReloadByImageResponseBody) SetRequestId(v string) *BatchReloadByImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *BatchReloadByImageResponseBody) SetTemplateId(v string) *BatchReloadByImageResponseBody {
	s.TemplateId = &v
	return s
}

func (s *BatchReloadByImageResponseBody) SetImageSnapshotRecordIds(v string) *BatchReloadByImageResponseBody {
	s.ImageSnapshotRecordIds = &v
	return s
}


type BatchReloadByImageResponseBodyBuilder struct {
	s *BatchReloadByImageResponseBody
}

func NewBatchReloadByImageResponseBodyBuilder() *BatchReloadByImageResponseBodyBuilder {
	s := &BatchReloadByImageResponseBody{}
	b := &BatchReloadByImageResponseBodyBuilder{s: s}
	return b
}

func (b *BatchReloadByImageResponseBodyBuilder) MachineId(v string) *BatchReloadByImageResponseBodyBuilder {
	b.s.MachineId = &v
	return b
}

func (b *BatchReloadByImageResponseBodyBuilder) RequestId(v string) *BatchReloadByImageResponseBodyBuilder {
	b.s.RequestId = &v
	return b
}

func (b *BatchReloadByImageResponseBodyBuilder) TemplateId(v string) *BatchReloadByImageResponseBodyBuilder {
	b.s.TemplateId = &v
	return b
}

func (b *BatchReloadByImageResponseBodyBuilder) ImageSnapshotRecordIds(v string) *BatchReloadByImageResponseBodyBuilder {
	b.s.ImageSnapshotRecordIds = &v
	return b
}

func (b *BatchReloadByImageResponseBodyBuilder) Build() *BatchReloadByImageResponseBody {
	return b.s
}


