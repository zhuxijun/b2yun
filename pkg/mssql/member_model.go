package mssql

import (
	root "b2yun/pkg"
)

// MemberInfoModel 更新会员信息模型
type MemberInfoModel struct {
	UserID   string `db:"user_id"`
	UserName string `db:"user_name"`
	UserRank string `db:"user_rank"`
	BranchNO string `db:"branch_no"`
	IsEnable string `db:"is_enable"`
	TransID  string `db:"ftransid"`
}

func (s MemberInfoModel) toMemberInfo() root.MemberInfo {
	var memberInfo root.MemberInfo

	memberInfo.UserID = s.UserID
	memberInfo.UserName = s.UserName
	memberInfo.UserRank = s.UserRank
	memberInfo.BranchNO = s.BranchNO
	memberInfo.IsEnable = s.IsEnable
	memberInfo.TransID = s.TransID

	return memberInfo
}
