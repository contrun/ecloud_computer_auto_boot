// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetVmListResponseBody struct {

    // 桌面信息列表
	MachineList *[]GetVmListResponseMachineList `json:"machineList,omitempty"`
}

func (s GetVmListResponseBody) String() string {
	return utils.Beautify(s)
}

func (s GetVmListResponseBody) GoString() string {
	return s.String()
}

func (s GetVmListResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetVmListResponseBody) SetMachineList(v []GetVmListResponseMachineList) *GetVmListResponseBody {
	s.MachineList = &v
	return s
}


type GetVmListResponseBodyBuilder struct {
	s *GetVmListResponseBody
}

func NewGetVmListResponseBodyBuilder() *GetVmListResponseBodyBuilder {
	s := &GetVmListResponseBody{}
	b := &GetVmListResponseBodyBuilder{s: s}
	return b
}

func (b *GetVmListResponseBodyBuilder) MachineList(v []GetVmListResponseMachineList) *GetVmListResponseBodyBuilder {
	b.s.MachineList = &v
	return b
}

func (b *GetVmListResponseBodyBuilder) Build() *GetVmListResponseBody {
	return b.s
}


