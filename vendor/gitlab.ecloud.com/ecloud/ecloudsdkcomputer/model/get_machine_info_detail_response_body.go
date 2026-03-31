// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetMachineInfoDetailResponseBody struct {

    // 成功集合
	Records *[]GetMachineInfoDetailResponseRecords `json:"records,omitempty"`
}

func (s GetMachineInfoDetailResponseBody) String() string {
	return utils.Beautify(s)
}

func (s GetMachineInfoDetailResponseBody) GoString() string {
	return s.String()
}

func (s GetMachineInfoDetailResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetMachineInfoDetailResponseBody) SetRecords(v []GetMachineInfoDetailResponseRecords) *GetMachineInfoDetailResponseBody {
	s.Records = &v
	return s
}


type GetMachineInfoDetailResponseBodyBuilder struct {
	s *GetMachineInfoDetailResponseBody
}

func NewGetMachineInfoDetailResponseBodyBuilder() *GetMachineInfoDetailResponseBodyBuilder {
	s := &GetMachineInfoDetailResponseBody{}
	b := &GetMachineInfoDetailResponseBodyBuilder{s: s}
	return b
}

func (b *GetMachineInfoDetailResponseBodyBuilder) Records(v []GetMachineInfoDetailResponseRecords) *GetMachineInfoDetailResponseBodyBuilder {
	b.s.Records = &v
	return b
}

func (b *GetMachineInfoDetailResponseBodyBuilder) Build() *GetMachineInfoDetailResponseBody {
	return b.s
}


