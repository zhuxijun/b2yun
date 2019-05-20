package root

// Brand 分类信息
type Brand struct {
	BrandSN   string `json:"brand_sn"`   // 添加品牌时忽略此参数，带参数则为更新操作
	BrandName string `json:"brand_name"` // 品牌名称
	IsShow    string `json:"is_show"`    // 是否显示,1显示,0不显示;添加品牌不带此参数时默认1显示
	TransID   string `json:"-"`          //门店修改时间戳
}

// Brander 获取分类接口
type Brander interface {
	GetBrands() ([]Brand, error)
}
