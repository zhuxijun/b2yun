package root

// ReqMemberLevel 获取会员等级
type ReqMemberLevel struct {
	RandID   string `json:"rand_id"`
	RandName string `json:"rand_name"`
}

// ReqMemberInfo 获取会员信息
type ReqMemberInfo struct {
	UserID   string `json:"user_id"`
	UserName string `json:"user_name"`
}

// MemberInfo 更新会员信息
type MemberInfo struct {
	UserID   string `json:"user_id"`
	UserName string `json:"user_name"`
	UserRank string `json:"user_rank"`
	BranchNO string `json:"branch_no"`
	IsEnable string `json:"is_enable"`
	TransID  string `json:"-"` //修改时间戳
}

// Memberer 获取会员信息接口
type Memberer interface {
	//InsertMemberLevels() ([]ReqMemberLevel, error)
	//InsertMemberInfos() ([]ReqMemberInfo, error)
	GetMemberInfos() ([]MemberInfo, error)
}
