// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BatchCreateImageBody struct {
    position.Body
    // 起始编号
	StartNumber *int32 `json:"startNumber,omitempty"`
    // 镜像名称
	ImageName *string `json:"imageName,omitempty"`
    // 批量创建镜像信息列表
	ImageAddReqDtoList *[]BatchCreateImageRequestImageAddReqDtoList `json:"imageAddReqDtoList,omitempty"`
}

func (s BatchCreateImageBody) String() string {
	return utils.Beautify(s)
}

func (s BatchCreateImageBody) GoString() string {
	return s.String()
}

func (s BatchCreateImageBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BatchCreateImageBody) SetStartNumber(v int32) *BatchCreateImageBody {
	s.StartNumber = &v
	return s
}

func (s *BatchCreateImageBody) SetImageName(v string) *BatchCreateImageBody {
	s.ImageName = &v
	return s
}

func (s *BatchCreateImageBody) SetImageAddReqDtoList(v []BatchCreateImageRequestImageAddReqDtoList) *BatchCreateImageBody {
	s.ImageAddReqDtoList = &v
	return s
}


type BatchCreateImageBodyBuilder struct {
	s *BatchCreateImageBody
}

func NewBatchCreateImageBodyBuilder() *BatchCreateImageBodyBuilder {
	s := &BatchCreateImageBody{}
	b := &BatchCreateImageBodyBuilder{s: s}
	return b
}

func (b *BatchCreateImageBodyBuilder) StartNumber(v int32) *BatchCreateImageBodyBuilder {
	b.s.StartNumber = &v
	return b
}

func (b *BatchCreateImageBodyBuilder) ImageName(v string) *BatchCreateImageBodyBuilder {
	b.s.ImageName = &v
	return b
}

func (b *BatchCreateImageBodyBuilder) ImageAddReqDtoList(v []BatchCreateImageRequestImageAddReqDtoList) *BatchCreateImageBodyBuilder {
	b.s.ImageAddReqDtoList = &v
	return b
}

func (b *BatchCreateImageBodyBuilder) Build() *BatchCreateImageBody {
	return b.s
}


