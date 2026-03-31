// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetDeviceDetailQuery struct {
    position.Query
    // 设备UID
	DeviceUid *string `json:"deviceUid,omitempty"`
}

func (s GetDeviceDetailQuery) String() string {
	return utils.Beautify(s)
}

func (s GetDeviceDetailQuery) GoString() string {
	return s.String()
}

func (s GetDeviceDetailQuery) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetDeviceDetailQuery) SetDeviceUid(v string) *GetDeviceDetailQuery {
	s.DeviceUid = &v
	return s
}


type GetDeviceDetailQueryBuilder struct {
	s *GetDeviceDetailQuery
}

func NewGetDeviceDetailQueryBuilder() *GetDeviceDetailQueryBuilder {
	s := &GetDeviceDetailQuery{}
	b := &GetDeviceDetailQueryBuilder{s: s}
	return b
}

func (b *GetDeviceDetailQueryBuilder) DeviceUid(v string) *GetDeviceDetailQueryBuilder {
	b.s.DeviceUid = &v
	return b
}

func (b *GetDeviceDetailQueryBuilder) Build() *GetDeviceDetailQuery {
	return b.s
}


