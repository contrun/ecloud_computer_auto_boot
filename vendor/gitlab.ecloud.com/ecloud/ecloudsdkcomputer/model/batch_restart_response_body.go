// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BatchRestartResponseBody struct {

    // 操作成功的云电脑id
	SuccessList []string `json:"successList,omitempty"`
    // 操作失败的云电脑id
	FailList []string `json:"failList,omitempty"`
}

func (s BatchRestartResponseBody) String() string {
	return utils.Beautify(s)
}

func (s BatchRestartResponseBody) GoString() string {
	return s.String()
}

func (s BatchRestartResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BatchRestartResponseBody) SetSuccessList(v []string) *BatchRestartResponseBody {
	s.SuccessList = v
	return s
}

func (s *BatchRestartResponseBody) SetFailList(v []string) *BatchRestartResponseBody {
	s.FailList = v
	return s
}


type BatchRestartResponseBodyBuilder struct {
	s *BatchRestartResponseBody
}

func NewBatchRestartResponseBodyBuilder() *BatchRestartResponseBodyBuilder {
	s := &BatchRestartResponseBody{}
	b := &BatchRestartResponseBodyBuilder{s: s}
	return b
}

func (b *BatchRestartResponseBodyBuilder) SuccessList(v []string) *BatchRestartResponseBodyBuilder {
	b.s.SuccessList = v
	return b
}

func (b *BatchRestartResponseBodyBuilder) FailList(v []string) *BatchRestartResponseBodyBuilder {
	b.s.FailList = v
	return b
}

func (b *BatchRestartResponseBodyBuilder) Build() *BatchRestartResponseBody {
	return b.s
}


