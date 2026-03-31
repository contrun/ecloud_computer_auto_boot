// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetResourceDiskInfoBody struct {
    position.Body
    // 订购资源实例ID
	InstanceId *string `json:"instanceId,omitempty"`
}

func (s GetResourceDiskInfoBody) String() string {
	return utils.Beautify(s)
}

func (s GetResourceDiskInfoBody) GoString() string {
	return s.String()
}

func (s GetResourceDiskInfoBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetResourceDiskInfoBody) SetInstanceId(v string) *GetResourceDiskInfoBody {
	s.InstanceId = &v
	return s
}


type GetResourceDiskInfoBodyBuilder struct {
	s *GetResourceDiskInfoBody
}

func NewGetResourceDiskInfoBodyBuilder() *GetResourceDiskInfoBodyBuilder {
	s := &GetResourceDiskInfoBody{}
	b := &GetResourceDiskInfoBodyBuilder{s: s}
	return b
}

func (b *GetResourceDiskInfoBodyBuilder) InstanceId(v string) *GetResourceDiskInfoBodyBuilder {
	b.s.InstanceId = &v
	return b
}

func (b *GetResourceDiskInfoBodyBuilder) Build() *GetResourceDiskInfoBody {
	return b.s
}


