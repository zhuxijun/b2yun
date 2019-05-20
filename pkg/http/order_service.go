package http

import (
	root "b2yun/pkg"
	"encoding/json"
)

// OrderService 订单服务
type OrderService struct {
	client       *Client
	orderService root.Orderer
}

//DownOrderInfo 下载订单详情
func (s *OrderService) DownOrderInfo() error {

	return nil
}

//CancelOrder 取消订单
func (s *OrderService) CancelOrder() error {
	//获取需要取消订单的列表
	cancelOrders, error := s.orderService.GetOrderCancels()

	if error != nil {
		return error
	}

	for _, cancelOrder := range cancelOrders {
		path := "/order/index.php?action=cancel_order&order_sn=" + cancelOrder.OrderSN
		_, err := s.client.Get(path)

		if err != nil {
			return err
		}

		//回写已传输标志
		task := root.Record{}
		task.KeyMaps = map[string]string{"forder_sn": cancelOrder.OrderSN}
		task.Table = ""
		task.Flags = []string{"flag", "1"}

		error1 := s.client.taskService.UpdateFlag(task)
		if error1 != nil {
			return error1
		}
	}

	return nil
}

// UploadOrderInfo 更新订单物流信息
func (s *OrderService) UploadOrderInfo() error {
	//获取需要更新订单物流信息的列表
	orderStatuss, error := s.orderService.GetOrderStatuss()

	if error != nil {
		return error
	}

	path := "/order/index.php?action=update_order_status"

	output, err := json.Marshal(orderStatuss)

	if err != nil {
		return err
	}

	reqStr := string(output)

	err1 := s.client.Post(path, reqStr)
	if err1 != nil {
		return err1
	}

	//回写已传输标志
	for _, orderStatus := range orderStatuss {
		task := root.Record{}
		task.KeyMaps = map[string]string{"forder_sn": orderStatus.OrderSN}
		task.Table = ""
		task.Flags = []string{"flag", "1"}

		error1 := s.client.taskService.UpdateFlag(task)
		if error1 != nil {
			return error1
		}
	}

	return nil
}
