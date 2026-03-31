// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetMachineInfoDetailBody struct {
    position.Body
    // 云电脑ID列表
	MachineIdList []string `json:"machineIdList,omitempty"`
}

func (s GetMachineInfoDetailBody) String() string {
	return utils.Beautify(s)
}

func (s GetMachineInfoDetailBody) GoString() string {
	return s.String()
}

func (s GetMachineInfoDetailBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetMachineInfoDetailBody) SetMachineIdList(v []string) *GetMachineInfoDetailBody {
	s.MachineIdList = v
	return s
}


type GetMachineInfoDetailBodyBuilder struct {
	s *GetMachineInfoDetailBody
}

func NewGetMachineInfoDetailBodyBuilder() *GetMachineInfoDetailBodyBuilder {
	s := &GetMachineInfoDetailBody{}
	b := &GetMachineInfoDetailBodyBuilder{s: s}
	return b
}

func (b *GetMachineInfoDetailBodyBuilder) MachineIdList(v []string) *GetMachineInfoDetailBodyBuilder {
	b.s.MachineIdList = v
	return b
}

func (b *GetMachineInfoDetailBodyBuilder) Build() *GetMachineInfoDetailBody {
	return b.s
}


