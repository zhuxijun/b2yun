package mssql

import root "b2yun/pkg"

// BrandModel 门店模型
type BrandModel struct {
	BrandSN   string `db:"fitem_brdno"`
	BrandName string `db:"fitem_brdname"`
	IsShow    string `db:"is_show"`
	TransID   string `db:"ftransid"`
}

func (s BrandModel) toBrand() root.Brand {

	var brand root.Brand

	brand.BrandSN = s.BrandSN
	brand.BrandName = s.BrandName
	brand.IsShow = s.IsShow
	brand.TransID = s.TransID

	return brand
}
