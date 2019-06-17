package mssql

import (
	root "b2yun/pkg"
	"database/sql"
)

// MemberService 会员服务
type MemberService struct {
	session *Session
}

// InsertMemberLevels 将获取到的会员等级插入erp数据库
func (s *MemberService) InsertMemberLevels(datas []root.ReqMemberLevel) error {

	var flevelno string
	for _, data := range datas {

		err := s.session.db.QueryRow("select flevel_no as flevelno from t_bn_level where flevel_no = ?", data.RankID).Scan(&flevelno)

		if err == sql.ErrNoRows {
			_, err1 := s.session.db.Exec("insert into t_bn_level(flevel_no,flevel_name) values(?,?)", data.RankID, data.RankName)

			if err1 != nil {
				return err1
			}
		}

	}

	return nil
}

// InsertMemberInfos 将获取到的会员信息插入erp数据库
func (s *MemberService) InsertMemberInfos(datas []root.ReqMemberInfo) error {

	var fuserno string
	for _, data := range datas {

		err := s.session.db.QueryRow("select fuser_no as fuserno from t_br_user where fuser_no = ?", data.UserID).Scan(&fuserno)

		if err == sql.ErrNoRows {
			_, err1 := s.session.db.Exec(`insert into t_br_user(fuser_no,fuser_name,faddress_id,faddress_name,fconsignee,ftel,fmobile,femail,fcountry,fprovince,fcity,fdistrict,faddress,fzipcode,fsign_building,fbest_time,foper_time
											) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`, data.UserID, data.UserName, data.AddressID, data.AddressName, data.Consignee, data.Tel, data.Mobile, data.Email, data.Country, data.Province, data.City, data.District, data.Address, data.Zipcode, data.SignBuilding, data.BestTime, data.OperTime)

			if err1 != nil {
				return err1
			}
		}
		// else {
		// 	_, err1 := s.session.db.Exec(`update t_br_user set fuser_name = ?,faddress_id = ?,faddress_name = ?,fconsignee = ?,ftel = ?,fmobile = ?,femail = ?,fcountry = ?,fprovince = ?,fcity = ?,fdistrict = ?,faddress = ?,fzipcode = ?,fsign_building = ?,fbest_time = ?,foper_time = ? where fuser_no = ?
		// 		`, data.UserName, data.AddressID, data.AddressName, data.Consignee, data.Tel, data.Mobile, data.Email, data.Country, data.Province, data.City, data.District, data.Address, data.Zipcode, data.SignBuilding, data.BestTime, data.OperTime, data.UserID)

		// 	if err1 != nil {
		// 		return err1
		// 	}
		// }

	}

	return nil
}

// GetMemberInfos 获取需要更新的会员信息列表
func (s *MemberService) GetMemberInfos() ([]root.MemberInfo, error) {
	var models []MemberInfoModel

	err := s.session.db.Select(&models, `
			select user_id = t1.fuser_no,
				user_name = t1.fuser_name,
				user_rank = isnull(t3.flevel_no,''),
				branch_no = isnull(t2.fbrh_no,''),
				is_enable = '1',
				ftransid = (case when isnull(t3.ftransid,0) > (case when t1.ftransid > isnull(t2.ftransid,0) then t1.ftransid else isnull(t2.ftransid,0) end) 
						then isnull(t3.ftransid,0) else (case when t1.ftransid > isnull(t2.ftransid,0) then t1.ftransid else isnull(t2.ftransid,0) end) end)
			from t_br_user t1
			inner join t_br_master t2 on t1.fbrh_no = t2.fbrh_no
			inner join t_bn_master t3 on t2.fbn_no = t3.fbn_no
			left join ts_t_transtype_info_mtq t5 WITH (NOLOCK) on (t5.fun_name='MemberInfoEntity')
			WHERE t1.fstatus = '1'
				and (case when isnull(t3.ftransid,0) > (case when t1.ftransid > isnull(t2.ftransid,0) then t1.ftransid else isnull(t2.ftransid,0) end) 
					then isnull(t3.ftransid,0) else (case when t1.ftransid > isnull(t2.ftransid,0) then t1.ftransid else isnull(t2.ftransid,0) end) end) > t5.ftransid
			ORDER BY (case when isnull(t3.ftransid,0) > (case when t1.ftransid > isnull(t2.ftransid,0) then t1.ftransid else isnull(t2.ftransid,0) end) 
					then isnull(t3.ftransid,0) else (case when t1.ftransid > isnull(t2.ftransid,0) then t1.ftransid else isnull(t2.ftransid,0) end) end)
			`)
	if err != nil {
		return nil, err
	}

	var memberInfos []root.MemberInfo

	for _, model := range models {
		memberInfos = append(memberInfos, model.toMemberInfo())
	}

	return memberInfos, nil
}
