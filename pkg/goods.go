package root

// Goods 商品信息(添加时是否必须、编辑时是否必须；默认为否、否)
type Goods struct {
	GoodsSN          string      `json:"goods_sn"`                     // (是、是)商品货号
	GoodsName        string      `json:"goods_name"`                   // (是、否)商品名称
	CatSN            string      `json:"cat_sn,omitempty"`             // (是、否)类别ID
	BrandSN          string      `json:"brand_sn,omitempty"`           // 品牌ID
	SuppliersID      string      `json:"suppliers_id,omitempty"`       // 货商ID;0为自营商品,其他对应货商id则为货商商品
	GoodsBarcode     string      `json:"goods_barcode,omitempty"`      // 条码; 多个条码，请用英文逗号（“,”）隔开。
	GoodsSize        string      `json:"goods_size,omitempty"`         // 规格
	GoodsUnit        string      `json:"goods_unit,omitempty"`         // 单位
	GoodsArea        string      `json:"goods_area,omitempty"`         // 产地
	GoodsNumber      string      `json:"goods_number,omitempty"`       // 库存数量
	GoodsPrice       string      `json:"goods_price,omitempty"`        // 参考进价
	ShopPrice        string      `json:"shop_price,omitempty"`         // 默认商品售价;如存在等级价格时将按等级价格显示
	RankPrice        []RankPrice `json:"rank_price,omitempty"`         // 会员等级价格: user_rank会员等级ID;user_price会员等级价格;
	GoodsWeight      string      `json:"goods_weight,omitempty"`       // 商品重量KG
	WarnNumber       string      `json:"warn_number,omitempty"`        // 库存预警数量
	MinNumber        string      `json:"min_number,omitempty"`         // 配送倍数（订货数量按倍数递增）
	LimiteNumber     string      `json:"limite_number,omitempty"`      // 限购数量
	IsOnSale         string      `json:"is_on_sale,omitempty"`         // 销售状态 1上架,0下架
	IsBest           string      `json:"is_best,omitempty"`            // 是否为精品;1是,0不是
	IsHot            string      `json:"is_hot,omitempty"`             // 是否为热销;1是,0不是
	IsNew            string      `json:"is_new,omitempty"`             // 是否为新品;1是,0不是
	IsNotAccumulated string      `json:"is_not_accumulated,omitempty"` // 不计算配送金额（如800元起送，那么这个商品金额不包含在内，需其他商品累计达到800元方达成起送金额）1代表不计算，0代表计算
	IsNotBonus       string      `json:"is_not_bonus,omitempty"`       // 不可使用优惠券（红包） 1代表不可使用，0代表可使用
	IsNotRebonus     string      `json:"is_not_rebonus,omitempty"`     // 不参与送优惠券（红包）活动 1代表不参与，0代表参与
	TransID          string      `json:"-"`                            // 门店修改时间戳
}

// RankPrice 会员等级对应的价格
type RankPrice struct {
	UserRank  string `json:"user_rank"`
	UserPrice string `json:"user_price"`
}

// Goodser 获取分类接口
type Goodser interface {
	GetGoodss(string) ([]Goods, error)
}
