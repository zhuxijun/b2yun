package http

import (
	root "b2yun/pkg"
	"encoding/json"
)

// MemberService 会员服务
type MemberService struct {
	client        *Client
	memberService root.Memberer
}

// DownloadMemberLevel 下载会员等级信息
func (s *MemberService) DownloadMemberLevel() error {

	return nil
}

// DownloadMemberInfo 下载会员信息
func (s *MemberService) DownloadMemberInfo() error {

	return nil
}

// UploadMemberInfo 上传需要更新的会员信息
func (s *MemberService) UploadMemberInfo() error {
	//获取信息
	memberInfos, error := s.memberService.GetMemberInfos()

	if error != nil {
		return error
	}

	path := "/users/index.php?action=update_users_info"

	output, err := json.Marshal(memberInfos)

	if err != nil {
		return err
	}

	reqStr := string(output)

	err1 := s.client.Post(path, reqStr)
	if err1 != nil {
		return err1
	}

	memberInfo := memberInfos[0]

	task := root.Task{
		Name: "MemberEntity",
		ID:   memberInfo.TransID,
	}

	error1 := s.client.taskService.UpdateID(task)
	if error1 != nil {
		return error1
	}

	return nil
}
