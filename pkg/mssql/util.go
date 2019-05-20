package mssql

//SheetNo string类型，订单号
type SheetNo string

//GetGroupNo 获取订单中分组编号
func (s SheetNo) GetGroupNo() string {
	grpNo := string(s[6:8])
	if grpNo == "00" {
		return ""
	}
	return grpNo
}
