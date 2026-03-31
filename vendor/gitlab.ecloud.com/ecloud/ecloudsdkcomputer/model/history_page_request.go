// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type HistoryPageRequest struct {


	HistoryPageBody *HistoryPageBody `json:"historyPageBody,omitempty"`
}

func (s HistoryPageRequest) String() string {
	return utils.Beautify(s)
}

func (s HistoryPageRequest) GoString() string {
	return s.String()
}

func (s HistoryPageRequest) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *HistoryPageRequest) SetHistoryPageBody(v *HistoryPageBody) *HistoryPageRequest {
	s.HistoryPageBody = v
	return s
}


type HistoryPageRequestBuilder struct {
	s *HistoryPageRequest
}

func NewHistoryPageRequestBuilder() *HistoryPageRequestBuilder {
	s := &HistoryPageRequest{}
	b := &HistoryPageRequestBuilder{s: s}
	return b
}

func (b *HistoryPageRequestBuilder) HistoryPageBody(v *HistoryPageBody) *HistoryPageRequestBuilder {
	b.s.HistoryPageBody = v
	return b
}

func (b *HistoryPageRequestBuilder) Build() *HistoryPageRequest {
	return b.s
}


