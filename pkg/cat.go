package root

// Cat 分类信息
type Cat struct {
	CatSN    string `json:"cat_sn"`              // 添加分类时忽略此参数，带参数则为更新操作
	CatName  string `json:"cat_name"`            // 分类名称
	ParentSN string `json:"parent_sn,omitempty"` // 上级分类ID
	IsShow   string `json:"is_show"`             // 是否显示,1显示,0不显示;添加分类不带此参数时默认1显示
	TransID  string `json:"-"`                   //门店修改时间戳
}

// Cater 获取分类接口
type Cater interface {
	GetCats() ([]Cat, error)
}
