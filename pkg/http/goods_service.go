package http

import (
	root "b2yun/pkg"
	"encoding/json"
)

// GoodsService 商品信息服务
type GoodsService struct {
	client       *Client
	goodsService root.Goodser
}

// UploadGoodss 上传商品信息
func (s *GoodsService) UploadGoodss() error {

	// 获取商品信息
	goodss, error := s.goodsService.GetGoodss()
	if error != nil {
		return error
	}

	//----------------------------------------------------
	path := "/goods/index.php?action=update_goods_info"

	output, err := json.Marshal(goodss)
	if err != nil {
		return err
	}

	reqStr := string(output)

	if err := s.client.Post(path, reqStr); err != nil {
		return err
	}

	goods := goodss[0]

	task := root.Task{
		Name: "GoodsEntity",
		ID:   goods.TransID,
	}

	error = s.client.taskService.UpdateID(task)
	if error != nil {
		return error
	}
	//----------------------------------------------------

	// // 上传商品信息
	// for _, goods := range goodss {

	// 	err := s.updateGoods(goods)
	// 	if err != nil {
	// 		s.client.log.Logger.Errorf("本次post商品信息请求出现错误,错误信息[%s]", err)
	// 		//return err
	// 	}
	// }
	return nil
}

// // updateGoods 上传单个商品信息
// func (s *GoodsService) updateGoods(goods root.Goods) error {

// 	path := "/goods/index.php?action=update_goods_info"

// 	// reqStr, err := s.client.getReqStr(goods)
// 	// if err != nil {
// 	// 	return err
// 	// }

// 	output, err := json.Marshal(goods)
// 	if err != nil {
// 		return err
// 	}

// 	reqStr := string(output)

// 	if err := s.client.Post(path, reqStr); err != nil {
// 		return err
// 	}

// 	task := root.Task{
// 		Name: "GoodsEntity",
// 		ID:   goods.TransID,
// 	}

// 	error := s.client.taskService.UpdateID(task)
// 	if error != nil {
// 		return error
// 	}

// 	return nil

// }
