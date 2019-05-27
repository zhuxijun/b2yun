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
	path := "/users/index.php?action=get_user_rank"

	reqStr, err := s.client.Get(path)
	if err != nil {
		return err
	}

	//插入信息
	var commonResponse CommonResponseMemberLevel
	err1 := json.Unmarshal([]byte(reqStr), &commonResponse)
	if err1 != nil {
		return err1
	}

	error := s.memberService.InsertMemberLevels(commonResponse.Data)
	if error != nil {
		return error
	}

	return nil
}

// DownloadMemberInfo 下载会员信息
func (s *MemberService) DownloadMemberInfo() error {
	path := "/users/index.php?action=get_users_info&last_day=10"

	reqStr, err := s.client.Get(path)
	if err != nil {
		return err
	}

	//插入信息
	var commonResponse CommonResponseMemberInfo
	err1 := json.Unmarshal([]byte(reqStr), &commonResponse)
	if err1 != nil {
		return err1
	}

	error := s.memberService.InsertMemberInfos(commonResponse.Data)
	if error != nil {
		return error
	}

	return nil
}

// UploadMemberInfo 上传需要更新的会员信息
func (s *MemberService) UploadMemberInfo() error {
	//获取信息
	memberInfos, error := s.memberService.GetMemberInfos()

	if error != nil {
		return error
	}

	if len(memberInfos) == 0 {
		return nil
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

	memberInfo := memberInfos[len(memberInfos)-1]

	task := root.Task{
		Name: "MemberInfoEntity",
		ID:   memberInfo.TransID,
	}

	error1 := s.client.taskService.UpdateID(task)
	if error1 != nil {
		return error1
	}

	return nil
}
