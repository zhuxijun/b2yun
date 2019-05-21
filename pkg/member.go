package root

// ReqMemberLevel 获取会员等级
type ReqMemberLevel struct {
	RankID   string `json:"rank_id"`
	RankName string `json:"rank_name"`
}

// ReqMemberInfo 获取会员信息
type ReqMemberInfo struct {
	UserID       string `json:"user_id"`       //会员用户ID（唯一值不可修改）
	UserName     string `json:"user_name"`     //会员用户名称（唯一值可修改）
	AddressID    string `json:"address_id"`    //收货地址唯一ID
	AddressName  string `json:"address_name"`  //地址名称（店铺名称）必填项
	Consignee    string `json:"consignee"`     //收货人必填项
	Tel          string `json:"tel"`           //联系电话必填项
	Mobile       string `json:"mobile"`        //手机
	Email        string `json:"email"`         //邮件地址
	Country      string `json:"country"`       //国家
	Province     string `json:"province"`      //省份
	City         string `json:"city"`          //城市
	District     string `json:"district"`      //地区
	Address      string `json:"address"`       //详细地址必填项
	Zipcode      string `json:"zipcode"`       //邮政编码
	SignBuilding string `json:"sign_building"` //标志建筑
	BestTime     string `json:"best_time"`     //最佳送货时间
	OperTime     string `json:"oper_time"`     //记录最后操作时间
}

// MemberInfo 更新会员信息
type MemberInfo struct {
	UserID   string `json:"user_id"`
	UserName string `json:"user_name,omitempty"`
	UserRank string `json:"user_rank"`
	BranchNO string `json:"branch_no"`
	IsEnable string `json:"is_enable,omitempty"`
	TransID  string `json:"-"` //修改时间戳
}

// Memberer 获取会员信息接口
type Memberer interface {
	InsertMemberLevels([]ReqMemberLevel) error
	InsertMemberInfos([]ReqMemberInfo) error
	GetMemberInfos() ([]MemberInfo, error)
}
