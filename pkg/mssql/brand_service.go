package mssql

import (
	root "b2yun/pkg"
)

// BrandService 门店服务
type BrandService struct {
	session *Session
}

// GetBrands 获取门店列表
func (s *BrandService) GetBrands() ([]root.Brand, error) {

	var models []BrandModel

	err := s.session.db.Select(&models,
		`   SELECT top 1 fitem_brdno = t1.fitem_brdno,
				   fitem_brdname = t1.fitem_brdname,
				   is_show = '1',
				   ftransid = (select max(ftransid) from t_bb_master)
			FROM t_bb_master t1
			LEFT JOIN ts_t_transtype_info_mtq t5 WITH (NOLOCK) on (t5.fun_name='BrandEntity')
			WHERE t1.ftransid > t5.ftransid
			ORDER BY t1.ftransid asc`)

	if err != nil {
		return nil, err
	}

	var brands []root.Brand

	for _, model := range models {
		brands = append(brands, model.toBrand())
	}

	return brands, nil

}
