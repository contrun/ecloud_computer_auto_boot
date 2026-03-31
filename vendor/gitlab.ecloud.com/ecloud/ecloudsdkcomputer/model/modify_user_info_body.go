// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type ModifyUserInfoBody struct {
    position.Body
    // 批量修改用户信息列表
	BatchReqList *[]ModifyUserInfoRequestBatchReqList `json:"batchReqList,omitempty"`
}

func (s ModifyUserInfoBody) String() string {
	return utils.Beautify(s)
}

func (s ModifyUserInfoBody) GoString() string {
	return s.String()
}

func (s ModifyUserInfoBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *ModifyUserInfoBody) SetBatchReqList(v []ModifyUserInfoRequestBatchReqList) *ModifyUserInfoBody {
	s.BatchReqList = &v
	return s
}


type ModifyUserInfoBodyBuilder struct {
	s *ModifyUserInfoBody
}

func NewModifyUserInfoBodyBuilder() *ModifyUserInfoBodyBuilder {
	s := &ModifyUserInfoBody{}
	b := &ModifyUserInfoBodyBuilder{s: s}
	return b
}

func (b *ModifyUserInfoBodyBuilder) BatchReqList(v []ModifyUserInfoRequestBatchReqList) *ModifyUserInfoBodyBuilder {
	b.s.BatchReqList = &v
	return b
}

func (b *ModifyUserInfoBodyBuilder) Build() *ModifyUserInfoBody {
	return b.s
}


