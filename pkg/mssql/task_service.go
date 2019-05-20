package mssql

import (
	root "b2yun/pkg"
	"errors"
	"fmt"
)

// TaskService 任务服务
type TaskService struct {
	session *Session
}

// UpdateID 实现更新ID接口
func (s TaskService) UpdateID(task root.Task) error {

	result, err := s.session.db.Exec(`
		UPDATE ts_t_transtype_info_mtq
		SET  ftransid = ?
		WHERE fun_name = ?
	`, task.ID, task.Name)

	if err != nil {
		return err
	}

	rows, err := result.RowsAffected()
	if rows != 1 {
		return errors.New("影响行数错误")
	}

	return nil

}

// UpdateFlag 更新传输标志
func (s TaskService) UpdateFlag(record root.Record) error {

	var sqlClause string
	for key, value := range record.KeyMaps {
		sqlClause = sqlClause + fmt.Sprintf(`%s='%s'`, key, value)
	}

	sqlMain := fmt.Sprintf(`
		UPDATE %s 
		SET %s = '%s'
		WHERE %s
	`, record.Table, record.Flags[0], record.Flags[1], sqlClause)

	_, err := s.session.db.Exec(sqlMain)
	if err != nil {
		return err
	}

	return nil

}
