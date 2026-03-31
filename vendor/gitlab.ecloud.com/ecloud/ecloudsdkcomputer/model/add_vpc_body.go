// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type AddVpcBody struct {
    position.Body
    // vpc规格
	Specs *string `json:"specs,omitempty"`
    // vpc名称
	VpcName *string `json:"vpcName,omitempty"`
    // 子网信息
	BaseCloudNetworkReqDtoList *[]AddVpcRequestBaseCloudNetworkReqDtoList `json:"baseCloudNetworkReqDtoList,omitempty"`
    // vpc描述
	Description *string `json:"description,omitempty"`
    // CMDB可用区
	PoolCode *string `json:"poolCode,omitempty"`
}

func (s AddVpcBody) String() string {
	return utils.Beautify(s)
}

func (s AddVpcBody) GoString() string {
	return s.String()
}

func (s AddVpcBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *AddVpcBody) SetSpecs(v string) *AddVpcBody {
	s.Specs = &v
	return s
}

func (s *AddVpcBody) SetVpcName(v string) *AddVpcBody {
	s.VpcName = &v
	return s
}

func (s *AddVpcBody) SetBaseCloudNetworkReqDtoList(v []AddVpcRequestBaseCloudNetworkReqDtoList) *AddVpcBody {
	s.BaseCloudNetworkReqDtoList = &v
	return s
}

func (s *AddVpcBody) SetDescription(v string) *AddVpcBody {
	s.Description = &v
	return s
}

func (s *AddVpcBody) SetPoolCode(v string) *AddVpcBody {
	s.PoolCode = &v
	return s
}


type AddVpcBodyBuilder struct {
	s *AddVpcBody
}

func NewAddVpcBodyBuilder() *AddVpcBodyBuilder {
	s := &AddVpcBody{}
	b := &AddVpcBodyBuilder{s: s}
	return b
}

func (b *AddVpcBodyBuilder) Specs(v string) *AddVpcBodyBuilder {
	b.s.Specs = &v
	return b
}

func (b *AddVpcBodyBuilder) VpcName(v string) *AddVpcBodyBuilder {
	b.s.VpcName = &v
	return b
}

func (b *AddVpcBodyBuilder) BaseCloudNetworkReqDtoList(v []AddVpcRequestBaseCloudNetworkReqDtoList) *AddVpcBodyBuilder {
	b.s.BaseCloudNetworkReqDtoList = &v
	return b
}

func (b *AddVpcBodyBuilder) Description(v string) *AddVpcBodyBuilder {
	b.s.Description = &v
	return b
}

func (b *AddVpcBodyBuilder) PoolCode(v string) *AddVpcBodyBuilder {
	b.s.PoolCode = &v
	return b
}

func (b *AddVpcBodyBuilder) Build() *AddVpcBody {
	return b.s
}


