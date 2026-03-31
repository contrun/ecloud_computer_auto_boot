// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type BatchResetPwdBody struct {
    position.Body
    // 批量重置密码列表
	ResetPwdList *[]BatchResetPwdRequestResetPwdList `json:"resetPwdList,omitempty"`
}

func (s BatchResetPwdBody) String() string {
	return utils.Beautify(s)
}

func (s BatchResetPwdBody) GoString() string {
	return s.String()
}

func (s BatchResetPwdBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *BatchResetPwdBody) SetResetPwdList(v []BatchResetPwdRequestResetPwdList) *BatchResetPwdBody {
	s.ResetPwdList = &v
	return s
}


type BatchResetPwdBodyBuilder struct {
	s *BatchResetPwdBody
}

func NewBatchResetPwdBodyBuilder() *BatchResetPwdBodyBuilder {
	s := &BatchResetPwdBody{}
	b := &BatchResetPwdBodyBuilder{s: s}
	return b
}

func (b *BatchResetPwdBodyBuilder) ResetPwdList(v []BatchResetPwdRequestResetPwdList) *BatchResetPwdBodyBuilder {
	b.s.ResetPwdList = &v
	return b
}

func (b *BatchResetPwdBodyBuilder) Build() *BatchResetPwdBody {
	return b.s
}


