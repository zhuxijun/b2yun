package mssql

import (
	root "b2yun/pkg"
)

// OrderCancelModel 取消订单模型 GET类型，循环多次调GET,不用转json
type OrderCancelModel struct {
	OrderSN string `db:"order_sn"` //商城订单号
}

// OrderStatusModel 更新物流状态 POST类型
/*
字段order_status参数相关
状态码	说明
100	已确认（取单）
500	已分单（打印）
530	配货中（分拣/复核）
510	已发货（出车）
122	完成（结算）*/
type OrderStatusModel struct {
	OrderSN     string `db:"order_sn"`     //是否必须：是	所要查询状态的订单号,不提交订单号则返回全部订单号及状态码
	OrderStatus string `db:"order_status"` //是否必须：是	需要更新的订单状态码
	InvoiceNO   string `db:"invoice_no"`   //是否必须：是	510 已发货（出车）状态时必须，多个发货单号，请用英文逗号（“,”）隔开。
	ActionNote  string `db:"action_note"`  //是否必须：否	操作备注信息。510 已发货（出车）状态时为配送信息如：“张三:13800138000”。
}

func (s OrderCancelModel) toOrderCancel() root.OrderCancel {
	var ordercancel root.OrderCancel

	ordercancel.OrderSN = s.OrderSN

	return ordercancel
}

func (s OrderStatusModel) toOrderStatus() root.OrderStatus {

	var orderstatus root.OrderStatus

	orderstatus.OrderSN = s.OrderSN
	orderstatus.OrderStatus = s.OrderStatus
	orderstatus.InvoiceNO = s.InvoiceNO
	orderstatus.ActionNote = s.ActionNote

	return orderstatus
}
