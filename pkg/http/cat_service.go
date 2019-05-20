package http

import (
	root "b2yun/pkg"
	"encoding/json"
)

// CatService 分类信息服务
type CatService struct {
	client     *Client
	catService root.Cater
}

// UploadCats 上传分类信息
func (s *CatService) UploadCats() error {

	// 获取分类信息
	cats, error := s.catService.GetCats()
	if error != nil {
		return error
	}

	//----------------------------------------------------
	path := "/goods/index.php?action=update_cat_info"

	output, err := json.Marshal(cats)
	if err != nil {
		return err
	}

	reqStr := string(output)

	if err := s.client.Post(path, reqStr); err != nil {
		return err
	}

	cat := cats[0]

	task := root.Task{
		Name: "CatEntity",
		ID:   cat.TransID,
	}

	error = s.client.taskService.UpdateID(task)
	if error != nil {
		return error
	}
	//----------------------------------------------------

	// // 上传分类信息
	// for _, cat := range cats {

	// 	err := s.updateCat(cat)
	// 	if err != nil {
	// 		s.client.log.Logger.Errorf("本次post分类信息请求出现错误,错误信息[%s]", err)
	// 		//return err
	// 	}
	// }
	return nil
}

// // updateCat 上传单个分类信息
// func (s *CatService) updateCat(cat root.Cat) error {

// 	path := "/goods/index.php?action=update_cat_info"

// 	output, err := json.Marshal(cat)
// 	if err != nil {
// 		return err
// 	}

// 	reqStr := string(output)

// 	if err := s.client.Post(path, reqStr); err != nil {
// 		return err
// 	}

// 	task := root.Task{
// 		Name: "CatEntity",
// 		ID:   cat.TransID,
// 	}

// 	error := s.client.taskService.UpdateID(task)
// 	if error != nil {
// 		return error
// 	}

// 	return nil

// }
