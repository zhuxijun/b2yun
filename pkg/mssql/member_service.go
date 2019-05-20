package mssql

import (
	root "b2yun/pkg"
)

// MemberService 会员服务
type MemberService struct {
	session *Session
}

// GetMemberInfos 获取需要更新的会员信息列表
func (s *MemberService) GetMemberInfos() ([]root.MemberInfo, error) {
	var models []MemberInfoModel

	err := s.session.db.Select(&models, `
			select * from t_br_user`)
	if err != nil {
		return nil, err
	}

	var memberInfos []root.MemberInfo

	for _, model := range models {
		memberInfos = append(memberInfos, model.toMemberInfo())
	}

	return memberInfos, nil
}
