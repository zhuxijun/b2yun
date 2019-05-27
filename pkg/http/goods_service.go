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

	//上传商品基本信息
	error := s.uploadGoodss("GoodsBasicEntity")
	if error != nil {
		return error
	}

	//上传商品价格信息
	error1 := s.uploadGoodss("GoodsPriceEntity")
	if error1 != nil {
		return error1
	}

	//上传商品库存信息
	error2 := s.uploadGoodss("GoodsStockEntity")
	if error2 != nil {
		return error2
	}

	return nil
}

// uploadGoodss 对对对
func (s *GoodsService) uploadGoodss(entity string) error {
	// 获取商品信息
	goodss, error := s.goodsService.GetGoodss(entity)
	if error != nil {
		return error
	}

	if len(goodss) == 0 {
		return nil
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

	goods := goodss[len(goodss)-1]

	task := root.Task{
		Name: entity, //"GoodsBasicEntity",
		ID:   goods.TransID,
	}

	error = s.client.taskService.UpdateID(task)
	if error != nil {
		return error
	}

	return nil
}
