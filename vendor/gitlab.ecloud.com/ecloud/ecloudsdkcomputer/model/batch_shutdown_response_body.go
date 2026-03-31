// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BatchShutdownResponseBody struct {

    // 操作成功的云电脑id
	SuccessList []string `json:"successList,omitempty"`
    // 操作失败的云电脑id
	FailList []string `json:"failList,omitempty"`
}

func (s BatchShutdownResponseBody) String() string {
	return utils.Beautify(s)
}

func (s BatchShutdownResponseBody) GoString() string {
	return s.String()
}

func (s BatchShutdownResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BatchShutdownResponseBody) SetSuccessList(v []string) *BatchShutdownResponseBody {
	s.SuccessList = v
	return s
}

func (s *BatchShutdownResponseBody) SetFailList(v []string) *BatchShutdownResponseBody {
	s.FailList = v
	return s
}


type BatchShutdownResponseBodyBuilder struct {
	s *BatchShutdownResponseBody
}

func NewBatchShutdownResponseBodyBuilder() *BatchShutdownResponseBodyBuilder {
	s := &BatchShutdownResponseBody{}
	b := &BatchShutdownResponseBodyBuilder{s: s}
	return b
}

func (b *BatchShutdownResponseBodyBuilder) SuccessList(v []string) *BatchShutdownResponseBodyBuilder {
	b.s.SuccessList = v
	return b
}

func (b *BatchShutdownResponseBodyBuilder) FailList(v []string) *BatchShutdownResponseBodyBuilder {
	b.s.FailList = v
	return b
}

func (b *BatchShutdownResponseBodyBuilder) Build() *BatchShutdownResponseBody {
	return b.s
}


