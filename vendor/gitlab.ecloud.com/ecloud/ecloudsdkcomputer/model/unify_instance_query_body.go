// @Title  Golang SDK Client
// @Description  This code is auto generated
// @Author  Ecloud SDK

package model

import (
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/position"
    "gitlab.ecloud.com/ecloud/ecloudsdkcore/utils"
)

type UnifyInstanceQueryBody struct {
    position.Body
    // 用户
	OrderUserId *string `json:"orderUserId,omitempty"`
    // 省管理員审核结果
	AdminAuditResult *bool `json:"adminAuditResult,omitempty"`
    // 页码
	PageNumber *int32 `json:"pageNumber,omitempty"`
    // 硬件订单状态列表列表 示例：102：已确认；103 已发货；104：已签收；105：已交付；106：订单取消；107：订单审核中；108：签收中
	StatusList []int32 `json:"statusList,omitempty"`
    // 订单号
	OrderId *string `json:"orderId,omitempty"`
    // sn号
	SnNumber *string `json:"snNumber,omitempty"`
    // 驳回原因
	DueRejectReason *string `json:"dueRejectReason,omitempty"`
    // 是否为大订单
	IsLarge *bool `json:"isLarge,omitempty"`
    // 邮箱
	OrderUserMail *string `json:"orderUserMail,omitempty"`
    // 页大小
	PageSize *int32 `json:"pageSize,omitempty"`
    // 审核结果
	AuditResult *bool `json:"auditResult,omitempty"`
    // 硬件订单状态示例：102：已确认；103 已发货；104：已签收；105：已交付；106：订单取消；107：订单审核中；108：签收中
	Status *int32 `json:"status,omitempty"`
}

func (s UnifyInstanceQueryBody) String() string {
	return utils.Beautify(s)
}

func (s UnifyInstanceQueryBody) GoString() string {
	return s.String()
}

func (s UnifyInstanceQueryBody) ToJsonString() string {
	return utils.ToJsonString(s)
}

func (s *UnifyInstanceQueryBody) SetOrderUserId(v string) *UnifyInstanceQueryBody {
	s.OrderUserId = &v
	return s
}

func (s *UnifyInstanceQueryBody) SetAdminAuditResult(v bool) *UnifyInstanceQueryBody {
	s.AdminAuditResult = &v
	return s
}

func (s *UnifyInstanceQueryBody) SetPageNumber(v int32) *UnifyInstanceQueryBody {
	s.PageNumber = &v
	return s
}

func (s *UnifyInstanceQueryBody) SetStatusList(v []int32) *UnifyInstanceQueryBody {
	s.StatusList = v
	return s
}

func (s *UnifyInstanceQueryBody) SetOrderId(v string) *UnifyInstanceQueryBody {
	s.OrderId = &v
	return s
}

func (s *UnifyInstanceQueryBody) SetSnNumber(v string) *UnifyInstanceQueryBody {
	s.SnNumber = &v
	return s
}

func (s *UnifyInstanceQueryBody) SetDueRejectReason(v string) *UnifyInstanceQueryBody {
	s.DueRejectReason = &v
	return s
}

func (s *UnifyInstanceQueryBody) SetIsLarge(v bool) *UnifyInstanceQueryBody {
	s.IsLarge = &v
	return s
}

func (s *UnifyInstanceQueryBody) SetOrderUserMail(v string) *UnifyInstanceQueryBody {
	s.OrderUserMail = &v
	return s
}

func (s *UnifyInstanceQueryBody) SetPageSize(v int32) *UnifyInstanceQueryBody {
	s.PageSize = &v
	return s
}

func (s *UnifyInstanceQueryBody) SetAuditResult(v bool) *UnifyInstanceQueryBody {
	s.AuditResult = &v
	return s
}

func (s *UnifyInstanceQueryBody) SetStatus(v int32) *UnifyInstanceQueryBody {
	s.Status = &v
	return s
}


type UnifyInstanceQueryBodyBuilder struct {
	s *UnifyInstanceQueryBody
}

func NewUnifyInstanceQueryBodyBuilder() *UnifyInstanceQueryBodyBuilder {
	s := &UnifyInstanceQueryBody{}
	b := &UnifyInstanceQueryBodyBuilder{s: s}
	return b
}

func (b *UnifyInstanceQueryBodyBuilder) OrderUserId(v string) *UnifyInstanceQueryBodyBuilder {
	b.s.OrderUserId = &v
	return b
}

func (b *UnifyInstanceQueryBodyBuilder) AdminAuditResult(v bool) *UnifyInstanceQueryBodyBuilder {
	b.s.AdminAuditResult = &v
	return b
}

func (b *UnifyInstanceQueryBodyBuilder) PageNumber(v int32) *UnifyInstanceQueryBodyBuilder {
	b.s.PageNumber = &v
	return b
}

func (b *UnifyInstanceQueryBodyBuilder) StatusList(v []int32) *UnifyInstanceQueryBodyBuilder {
	b.s.StatusList = v
	return b
}

func (b *UnifyInstanceQueryBodyBuilder) OrderId(v string) *UnifyInstanceQueryBodyBuilder {
	b.s.OrderId = &v
	return b
}

func (b *UnifyInstanceQueryBodyBuilder) SnNumber(v string) *UnifyInstanceQueryBodyBuilder {
	b.s.SnNumber = &v
	return b
}

func (b *UnifyInstanceQueryBodyBuilder) DueRejectReason(v string) *UnifyInstanceQueryBodyBuilder {
	b.s.DueRejectReason = &v
	return b
}

func (b *UnifyInstanceQueryBodyBuilder) IsLarge(v bool) *UnifyInstanceQueryBodyBuilder {
	b.s.IsLarge = &v
	return b
}

func (b *UnifyInstanceQueryBodyBuilder) OrderUserMail(v string) *UnifyInstanceQueryBodyBuilder {
	b.s.OrderUserMail = &v
	return b
}

func (b *UnifyInstanceQueryBodyBuilder) PageSize(v int32) *UnifyInstanceQueryBodyBuilder {
	b.s.PageSize = &v
	return b
}

func (b *UnifyInstanceQueryBodyBuilder) AuditResult(v bool) *UnifyInstanceQueryBodyBuilder {
	b.s.AuditResult = &v
	return b
}

func (b *UnifyInstanceQueryBodyBuilder) Status(v int32) *UnifyInstanceQueryBodyBuilder {
	b.s.Status = &v
	return b
}

func (b *UnifyInstanceQueryBodyBuilder) Build() *UnifyInstanceQueryBody {
	return b.s
}


