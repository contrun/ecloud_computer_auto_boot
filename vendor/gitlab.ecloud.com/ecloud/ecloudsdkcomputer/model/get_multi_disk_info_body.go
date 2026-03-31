// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetMultiDiskInfoBody struct {
    position.Body
    // 资源实例id，示例值：CCA-XXX
	InstanceId *string `json:"instanceId,omitempty"`
}

func (s GetMultiDiskInfoBody) String() string {
	return utils.Beautify(s)
}

func (s GetMultiDiskInfoBody) GoString() string {
	return s.String()
}

func (s GetMultiDiskInfoBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetMultiDiskInfoBody) SetInstanceId(v string) *GetMultiDiskInfoBody {
	s.InstanceId = &v
	return s
}


type GetMultiDiskInfoBodyBuilder struct {
	s *GetMultiDiskInfoBody
}

func NewGetMultiDiskInfoBodyBuilder() *GetMultiDiskInfoBodyBuilder {
	s := &GetMultiDiskInfoBody{}
	b := &GetMultiDiskInfoBodyBuilder{s: s}
	return b
}

func (b *GetMultiDiskInfoBodyBuilder) InstanceId(v string) *GetMultiDiskInfoBodyBuilder {
	b.s.InstanceId = &v
	return b
}

func (b *GetMultiDiskInfoBodyBuilder) Build() *GetMultiDiskInfoBody {
	return b.s
}


