package mssql

import (
	root "b2yun/pkg"
)

// GoodsModel 门店模型
type GoodsModel struct {
	GoodsSN          string      `db:"goods_sn"`      // (是、是)商品货号
	GoodsName        string      `db:"goods_name"`    // (是、否)商品名称
	CatSN            string      `db:"cat_sn"`        // (是、否)类别ID
	BrandSN          string      `db:"brand_sn"`      // 品牌ID
	SuppliersID      string      `db:"suppliers_id"`  // 货商ID;0为自营商品,其他对应货商id则为货商商品
	GoodsBarcode     string      `db:"goods_barcode"` // 条码; 多个条码，请用英文逗号（“,”）隔开。
	GoodsSize        string      `db:"goods_size"`    // 规格
	GoodsUnit        string      `db:"goods_unit"`    // 单位
	GoodsArea        string      `db:"goods_area"`    // 产地
	GoodsNumber      string      `db:"goods_number"`  // 库存数量
	GoodsPrice       string      `db:"goods_price"`   // 参考进价
	ShopPrice        string      `db:"shop_price"`    // 默认商品售价;如存在等级价格时将按等级价格显示
	RankPrice        []RandPrice // 会员等级价格: user_rank会员等级ID;user_price会员等级价格;
	GoodsWeight      string      `db:"goods_weight"`       // 商品重量KG
	WarnNumber       string      `db:"warn_number"`        // 库存预警数量
	MinNumber        string      `db:"min_number"`         // 配送倍数（订货数量按倍数递增）
	LimiteNumber     string      `db:"limite_number"`      // 限购数量
	IsOnSale         string      `db:"is_on_sale"`         // 销售状态 1上架,0下架
	IsBest           string      `db:"is_best"`            // 是否为精品;1是,0不是
	IsHot            string      `db:"is_hot"`             // 是否为热销;1是,0不是
	IsNew            string      `db:"is_new"`             // 是否为新品;1是,0不是
	IsNotAccumulated string      `db:"is_not_accumulated"` // 不计算配送金额（如800元起送，那么这个商品金额不包含在内，需其他商品累计达到800元方达成起送金额）1代表不计算，0代表计算
	IsNotBonus       string      `db:"is_not_bonus"`       // 不可使用优惠券（红包） 1代表不可使用，0代表可使用
	IsNotRebonus     string      `db:"is_not_rebonus"`     // 不参与送优惠券（红包）活动 1代表不参与，0代表参与
	TransID          string      `db:"ftransid"`           //门店修改时间戳
}

//RandPrice 商品会员等级、及对应的价格
type RandPrice struct {
	UserRank  string `db:"user_rank"`
	UserPrice string `db:"user_price"`
}

func (s GoodsModel) toGoods() root.Goods {

	var goods root.Goods

	goods.GoodsSN = s.GoodsSN
	goods.GoodsName = s.GoodsName
	goods.CatSN = s.CatSN
	goods.BrandSN = s.BrandSN
	goods.SuppliersID = s.SuppliersID
	goods.GoodsBarcode = s.GoodsBarcode
	goods.GoodsSize = s.GoodsSize
	goods.GoodsUnit = s.GoodsUnit
	goods.GoodsArea = s.GoodsArea
	goods.GoodsNumber = s.GoodsNumber
	goods.GoodsPrice = s.GoodsPrice
	goods.ShopPrice = s.ShopPrice
	//goods.RankPrice = s.RankPrice
	goods.GoodsWeight = s.GoodsWeight
	goods.WarnNumber = s.WarnNumber
	goods.MinNumber = s.MinNumber
	goods.LimiteNumber = s.LimiteNumber
	goods.IsOnSale = s.IsOnSale
	goods.IsBest = s.IsBest
	goods.IsHot = s.IsHot
	goods.IsNew = s.IsNew
	goods.IsNotAccumulated = s.IsNotAccumulated
	goods.IsNotBonus = s.IsNotBonus
	goods.IsNotRebonus = s.IsNotRebonus
	goods.TransID = s.TransID

	return goods

}
