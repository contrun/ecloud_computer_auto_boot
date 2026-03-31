// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type GetResourceDiskInfoResponseBody struct {

    // 存储uuid，格式：disk-xxxxxxxxxxxxxx
	DiskUid *string `json:"diskUid,omitempty"`
    // 存储总大小
	DiskSize *int32 `json:"diskSize,omitempty"`
    // 存储名称
	DiskBusinessName *string `json:"diskBusinessName,omitempty"`
    // 盘符
	LetterName *string `json:"letterName,omitempty"`
}

func (s GetResourceDiskInfoResponseBody) String() string {
	return utils.Beautify(s)
}

func (s GetResourceDiskInfoResponseBody) GoString() string {
	return s.String()
}

func (s GetResourceDiskInfoResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *GetResourceDiskInfoResponseBody) SetDiskUid(v string) *GetResourceDiskInfoResponseBody {
	s.DiskUid = &v
	return s
}

func (s *GetResourceDiskInfoResponseBody) SetDiskSize(v int32) *GetResourceDiskInfoResponseBody {
	s.DiskSize = &v
	return s
}

func (s *GetResourceDiskInfoResponseBody) SetDiskBusinessName(v string) *GetResourceDiskInfoResponseBody {
	s.DiskBusinessName = &v
	return s
}

func (s *GetResourceDiskInfoResponseBody) SetLetterName(v string) *GetResourceDiskInfoResponseBody {
	s.LetterName = &v
	return s
}


type GetResourceDiskInfoResponseBodyBuilder struct {
	s *GetResourceDiskInfoResponseBody
}

func NewGetResourceDiskInfoResponseBodyBuilder() *GetResourceDiskInfoResponseBodyBuilder {
	s := &GetResourceDiskInfoResponseBody{}
	b := &GetResourceDiskInfoResponseBodyBuilder{s: s}
	return b
}

func (b *GetResourceDiskInfoResponseBodyBuilder) DiskUid(v string) *GetResourceDiskInfoResponseBodyBuilder {
	b.s.DiskUid = &v
	return b
}

func (b *GetResourceDiskInfoResponseBodyBuilder) DiskSize(v int32) *GetResourceDiskInfoResponseBodyBuilder {
	b.s.DiskSize = &v
	return b
}

func (b *GetResourceDiskInfoResponseBodyBuilder) DiskBusinessName(v string) *GetResourceDiskInfoResponseBodyBuilder {
	b.s.DiskBusinessName = &v
	return b
}

func (b *GetResourceDiskInfoResponseBodyBuilder) LetterName(v string) *GetResourceDiskInfoResponseBodyBuilder {
	b.s.LetterName = &v
	return b
}

func (b *GetResourceDiskInfoResponseBodyBuilder) Build() *GetResourceDiskInfoResponseBody {
	return b.s
}


