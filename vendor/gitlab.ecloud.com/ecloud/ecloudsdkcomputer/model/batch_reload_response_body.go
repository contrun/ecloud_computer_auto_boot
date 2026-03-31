// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BatchReloadResponseBody struct {

    // 操作成功的云电脑id
	SuccessList []string `json:"successList,omitempty"`
    // 操作失败的云云电脑id
	FailList []string `json:"failList,omitempty"`
}

func (s BatchReloadResponseBody) String() string {
	return utils.Beautify(s)
}

func (s BatchReloadResponseBody) GoString() string {
	return s.String()
}

func (s BatchReloadResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BatchReloadResponseBody) SetSuccessList(v []string) *BatchReloadResponseBody {
	s.SuccessList = v
	return s
}

func (s *BatchReloadResponseBody) SetFailList(v []string) *BatchReloadResponseBody {
	s.FailList = v
	return s
}


type BatchReloadResponseBodyBuilder struct {
	s *BatchReloadResponseBody
}

func NewBatchReloadResponseBodyBuilder() *BatchReloadResponseBodyBuilder {
	s := &BatchReloadResponseBody{}
	b := &BatchReloadResponseBodyBuilder{s: s}
	return b
}

func (b *BatchReloadResponseBodyBuilder) SuccessList(v []string) *BatchReloadResponseBodyBuilder {
	b.s.SuccessList = v
	return b
}

func (b *BatchReloadResponseBodyBuilder) FailList(v []string) *BatchReloadResponseBodyBuilder {
	b.s.FailList = v
	return b
}

func (b *BatchReloadResponseBodyBuilder) Build() *BatchReloadResponseBody {
	return b.s
}


