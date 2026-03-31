// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BatchCreateImageRequestImageAddReqDtoList struct {

    // 云电脑id
	MachineId *string `json:"machineId,omitempty"`
    // 镜像模板描述
	ImageRemark *string `json:"imageRemark,omitempty"`
}

func (s BatchCreateImageRequestImageAddReqDtoList) String() string {
	return utils.Beautify(s)
}

func (s BatchCreateImageRequestImageAddReqDtoList) GoString() string {
	return s.String()
}

func (s BatchCreateImageRequestImageAddReqDtoList) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BatchCreateImageRequestImageAddReqDtoList) SetMachineId(v string) *BatchCreateImageRequestImageAddReqDtoList {
	s.MachineId = &v
	return s
}

func (s *BatchCreateImageRequestImageAddReqDtoList) SetImageRemark(v string) *BatchCreateImageRequestImageAddReqDtoList {
	s.ImageRemark = &v
	return s
}


type BatchCreateImageRequestImageAddReqDtoListBuilder struct {
	s *BatchCreateImageRequestImageAddReqDtoList
}

func NewBatchCreateImageRequestImageAddReqDtoListBuilder() *BatchCreateImageRequestImageAddReqDtoListBuilder {
	s := &BatchCreateImageRequestImageAddReqDtoList{}
	b := &BatchCreateImageRequestImageAddReqDtoListBuilder{s: s}
	return b
}

func (b *BatchCreateImageRequestImageAddReqDtoListBuilder) MachineId(v string) *BatchCreateImageRequestImageAddReqDtoListBuilder {
	b.s.MachineId = &v
	return b
}

func (b *BatchCreateImageRequestImageAddReqDtoListBuilder) ImageRemark(v string) *BatchCreateImageRequestImageAddReqDtoListBuilder {
	b.s.ImageRemark = &v
	return b
}

func (b *BatchCreateImageRequestImageAddReqDtoListBuilder) Build() *BatchCreateImageRequestImageAddReqDtoList {
	return b.s
}


