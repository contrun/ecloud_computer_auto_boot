// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BatchAvailableResponseBody struct {

    // 操作成功的云电脑id
	SuccessList []string `json:"successList,omitempty"`
    // 操作失败的云电脑id
	FailList []string `json:"failList,omitempty"`
}

func (s BatchAvailableResponseBody) String() string {
	return utils.Beautify(s)
}

func (s BatchAvailableResponseBody) GoString() string {
	return s.String()
}

func (s BatchAvailableResponseBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BatchAvailableResponseBody) SetSuccessList(v []string) *BatchAvailableResponseBody {
	s.SuccessList = v
	return s
}

func (s *BatchAvailableResponseBody) SetFailList(v []string) *BatchAvailableResponseBody {
	s.FailList = v
	return s
}


type BatchAvailableResponseBodyBuilder struct {
	s *BatchAvailableResponseBody
}

func NewBatchAvailableResponseBodyBuilder() *BatchAvailableResponseBodyBuilder {
	s := &BatchAvailableResponseBody{}
	b := &BatchAvailableResponseBodyBuilder{s: s}
	return b
}

func (b *BatchAvailableResponseBodyBuilder) SuccessList(v []string) *BatchAvailableResponseBodyBuilder {
	b.s.SuccessList = v
	return b
}

func (b *BatchAvailableResponseBodyBuilder) FailList(v []string) *BatchAvailableResponseBodyBuilder {
	b.s.FailList = v
	return b
}

func (b *BatchAvailableResponseBodyBuilder) Build() *BatchAvailableResponseBody {
	return b.s
}


