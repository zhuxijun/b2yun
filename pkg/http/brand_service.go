package http

import (
	root "b2yun/pkg"
	"encoding/json"
)

// BrandService 品牌信息服务
type BrandService struct {
	client       *Client
	brandService root.Brander
}

// UploadBrands 上传品牌信息
func (s *BrandService) UploadBrands() error {

	// 获取品牌信息
	brands, error := s.brandService.GetBrands()
	if error != nil {
		return error
	}

	//----------------------------------------------------
	path := "/goods/index.php?action=update_brand_info"

	output, err := json.Marshal(brands)
	if err != nil {
		return err
	}

	reqStr := string(output)

	err1 := s.client.Post(path, reqStr)
	if err1 != nil {
		return err1
	}

	brand := brands[len(brands)]

	task := root.Task{
		Name: "BrandEntity",
		ID:   brand.TransID,
	}

	error1 := s.client.taskService.UpdateID(task)
	if error1 != nil {
		return error1
	}
	//----------------------------------------------------

	// // 上传品牌信息
	// for _, brand := range brands {

	// 	err := s.updateBrand(brand)
	// 	if err != nil {
	// 		s.client.log.Logger.Errorf("本次post品牌信息请求出现错误,错误信息[%s]", err)
	// 		//return err
	// 	}
	// }
	return nil
}

// // updateBrand 上传单个品牌信息
// func (s *BrandService) updateBrand(brand root.Brand) error {

// 	path := "/goods/index.php?action=update_brand_info"

// 	// reqStr, err := s.client.getReqStr(brand)
// 	// if err != nil {
// 	// 	return err
// 	// }

// 	output, err := json.Marshal(brand)
// 	if err != nil {
// 		return err
// 	}

// 	reqStr := string(output)

// 	if err := s.client.Post(path, reqStr); err != nil {
// 		return err
// 	}

// 	task := root.Task{
// 		Name: "BrandEntity",
// 		ID:   brand.TransID,
// 	}

// 	error := s.client.taskService.UpdateID(task)
// 	if error != nil {
// 		return error
// 	}

// 	return nil

// }
