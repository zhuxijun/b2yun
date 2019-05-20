package mssql

import root "b2yun/pkg"

// CatModel 门店模型
type CatModel struct {
	CatSN    string `db:"fitem_clsno"`
	CatName  string `db:"fitem_clsname"`
	ParentSN string `db:"fprt_no"`
	IsShow   string `db:"is_show"`
	TransID  string `db:"ftransid"`
}

func (s CatModel) toCat() root.Cat {

	var cat root.Cat

	cat.CatSN = s.CatSN
	cat.CatName = s.CatName
	cat.ParentSN = s.ParentSN
	cat.IsShow = s.IsShow
	cat.TransID = s.TransID

	return cat
}
