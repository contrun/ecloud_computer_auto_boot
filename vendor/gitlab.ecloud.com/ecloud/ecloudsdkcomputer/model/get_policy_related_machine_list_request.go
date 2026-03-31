// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetPolicyRelatedMachineListRequest struct {


	GetPolicyRelatedMachineListBody *GetPolicyRelatedMachineListBody `json:"getPolicyRelatedMachineListBody,omitempty"`
}

func (s GetPolicyRelatedMachineListRequest) String() string {
	return utils.Beautify(s)
}

func (s GetPolicyRelatedMachineListRequest) GoString() string {
	return s.String()
}

func (s GetPolicyRelatedMachineListRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetPolicyRelatedMachineListRequest) SetGetPolicyRelatedMachineListBody(v *GetPolicyRelatedMachineListBody) *GetPolicyRelatedMachineListRequest {
	s.GetPolicyRelatedMachineListBody = v
	return s
}


type GetPolicyRelatedMachineListRequestBuilder struct {
	s *GetPolicyRelatedMachineListRequest
}

func NewGetPolicyRelatedMachineListRequestBuilder() *GetPolicyRelatedMachineListRequestBuilder {
	s := &GetPolicyRelatedMachineListRequest{}
	b := &GetPolicyRelatedMachineListRequestBuilder{s: s}
	return b
}

func (b *GetPolicyRelatedMachineListRequestBuilder) GetPolicyRelatedMachineListBody(v *GetPolicyRelatedMachineListBody) *GetPolicyRelatedMachineListRequestBuilder {
	b.s.GetPolicyRelatedMachineListBody = v
	return b
}

func (b *GetPolicyRelatedMachineListRequestBuilder) Build() *GetPolicyRelatedMachineListRequest {
	return b.s
}


