// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetPolicyRelatedMachineListResponseBody struct {

    // 数据
	Data *[]GetPolicyRelatedMachineListResponseData `json:"data,omitempty"`
    // 总行数
	TotalCount *int32 `json:"totalCount,omitempty"`
}

func (s GetPolicyRelatedMachineListResponseBody) String() string {
	return utils.Beautify(s)
}

func (s GetPolicyRelatedMachineListResponseBody) GoString() string {
	return s.String()
}

func (s GetPolicyRelatedMachineListResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetPolicyRelatedMachineListResponseBody) SetData(v []GetPolicyRelatedMachineListResponseData) *GetPolicyRelatedMachineListResponseBody {
	s.Data = &v
	return s
}

func (s *GetPolicyRelatedMachineListResponseBody) SetTotalCount(v int32) *GetPolicyRelatedMachineListResponseBody {
	s.TotalCount = &v
	return s
}


type GetPolicyRelatedMachineListResponseBodyBuilder struct {
	s *GetPolicyRelatedMachineListResponseBody
}

func NewGetPolicyRelatedMachineListResponseBodyBuilder() *GetPolicyRelatedMachineListResponseBodyBuilder {
	s := &GetPolicyRelatedMachineListResponseBody{}
	b := &GetPolicyRelatedMachineListResponseBodyBuilder{s: s}
	return b
}

func (b *GetPolicyRelatedMachineListResponseBodyBuilder) Data(v []GetPolicyRelatedMachineListResponseData) *GetPolicyRelatedMachineListResponseBodyBuilder {
	b.s.Data = &v
	return b
}

func (b *GetPolicyRelatedMachineListResponseBodyBuilder) TotalCount(v int32) *GetPolicyRelatedMachineListResponseBodyBuilder {
	b.s.TotalCount = &v
	return b
}

func (b *GetPolicyRelatedMachineListResponseBodyBuilder) Build() *GetPolicyRelatedMachineListResponseBody {
	return b.s
}


