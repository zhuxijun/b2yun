package mssql

import (
	root "b2yun/pkg"
)

// CatService 门店服务
type CatService struct {
	session *Session
}

// GetCats 获取门店列表
func (s *CatService) GetCats() ([]root.Cat, error) {

	var models []CatModel

	err := s.session.db.Select(&models,
		`   SELECT fitem_clsno = t1.fitem_clsno,
				   fitem_clsname = t1.fitem_clsname,
				   fprt_no = (case when t1.fprt_no = '*' then '' else t1.fprt_no end),
				   is_show = '1',
				   ftransid = (select max(ftransid) from t_bc_master)
			FROM t_bc_master t1
			LEFT JOIN ts_t_transtype_info_mtq t5 WITH (NOLOCK) on (t5.fun_name='CatEntity')
			WHERE t1.ftransid > t5.ftransid
			ORDER BY t1.fitem_clsno asc`)

	if err != nil {
		return nil, err
	}

	var cats []root.Cat

	for _, model := range models {
		cats = append(cats, model.toCat())
	}

	return cats, nil

}
