// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ImageRecordSharedWithOthersResponseData struct {

    // 接受者用户名
	Acceptor *string `json:"acceptor,omitempty"`
    // 共享时间
	ShareCreatedTime *string `json:"shareCreatedTime,omitempty"`
    // 共享状态
	ShareStatus *string `json:"shareStatus,omitempty"`
}

func (s ImageRecordSharedWithOthersResponseData) String() string {
	return utils.Beautify(s)
}

func (s ImageRecordSharedWithOthersResponseData) GoString() string {
	return s.String()
}

func (s ImageRecordSharedWithOthersResponseData) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ImageRecordSharedWithOthersResponseData) SetAcceptor(v string) *ImageRecordSharedWithOthersResponseData {
	s.Acceptor = &v
	return s
}

func (s *ImageRecordSharedWithOthersResponseData) SetShareCreatedTime(v string) *ImageRecordSharedWithOthersResponseData {
	s.ShareCreatedTime = &v
	return s
}

func (s *ImageRecordSharedWithOthersResponseData) SetShareStatus(v string) *ImageRecordSharedWithOthersResponseData {
	s.ShareStatus = &v
	return s
}


type ImageRecordSharedWithOthersResponseDataBuilder struct {
	s *ImageRecordSharedWithOthersResponseData
}

func NewImageRecordSharedWithOthersResponseDataBuilder() *ImageRecordSharedWithOthersResponseDataBuilder {
	s := &ImageRecordSharedWithOthersResponseData{}
	b := &ImageRecordSharedWithOthersResponseDataBuilder{s: s}
	return b
}

func (b *ImageRecordSharedWithOthersResponseDataBuilder) Acceptor(v string) *ImageRecordSharedWithOthersResponseDataBuilder {
	b.s.Acceptor = &v
	return b
}

func (b *ImageRecordSharedWithOthersResponseDataBuilder) ShareCreatedTime(v string) *ImageRecordSharedWithOthersResponseDataBuilder {
	b.s.ShareCreatedTime = &v
	return b
}

func (b *ImageRecordSharedWithOthersResponseDataBuilder) ShareStatus(v string) *ImageRecordSharedWithOthersResponseDataBuilder {
	b.s.ShareStatus = &v
	return b
}

func (b *ImageRecordSharedWithOthersResponseDataBuilder) Build() *ImageRecordSharedWithOthersResponseData {
	return b.s
}


