package mssql

import (
	root "b2yun/pkg"
)

// OrderService 订单服务
type OrderService struct {
	session *Session
}

// GetOrderCancels 获取取消订单列表
func (s *OrderService) GetOrderCancels() ([]root.OrderCancel, error) {
	var models []OrderCancelModel

	err := s.session.db.Select(&models, `
			select forder_sn 
			from b2yun_order_master
			where fstatus = '9'`)

	if err != nil {
		return nil, err
	}

	var ordercancels []root.OrderCancel

	for _, model := range models {
		ordercancels = append(ordercancels, model.toOrderCancel())
	}

	return ordercancels, nil
}

// GetOrderStatuss 获取需要更新物流状态的订单列表
func (s *OrderService) GetOrderStatuss() ([]root.OrderStatus, error) {
	var models []OrderStatusModel

	err := s.session.db.Select(&models, `
			select forder_sn
			from b2yun_order_master
			where fstatus = '9'`)

	if err != nil {
		return nil, err
	}

	var orderstatuss []root.OrderStatus

	for _, model := range models {
		orderstatuss = append(orderstatuss, model.toOrderStatus())
	}

	return orderstatuss, nil
}
