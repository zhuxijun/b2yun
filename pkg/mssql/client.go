package mssql

import (
	root "b2yun/pkg"
	"time"

	_ "github.com/denisenkom/go-mssqldb" // 注册驱动
	"github.com/jmoiron/sqlx"
)

//Client 数据库交互对象
type Client struct {

	// 数据库连接
	db *sqlx.DB

	// 数据库连接
	Configer root.Configer

	// now
	Now func() time.Time
}

// NewClient 创建新连接,参数为接口，实现获取配置参数功能
func NewClient(configer root.Configer) *Client {

	c := Client{Configer: configer}
	c.Now = time.Now
	return &c
}

// Connect 打开并连接数据库，返回Session
func (c *Client) Connect() *Session {

	s := NewSession(c.db)

	s.now = c.Now()

	return s
}

// Open 打开数据库连接
func (c *Client) Open() error {

	const op = "mssql.Client.Open"
	var customError root.Error
	customError.Op = op

	config, err := c.Configer.GetConfig()
	if err != nil {
		customError.Code = ECONFIGINVALID
		customError.Err = err
		return &customError
	}

	if config.Mssql == nil {
		customError.Code = ECONFIGMSSQLNOTFOUND
		return &customError
	}

	var dialects = config.Mssql.Dialects
	if config.Mssql.Dialects == "" {
		dialects = "mssql"
	}

	db, err := sqlx.Open(dialects, config.Mssql.Parm)
	if err != nil {
		customError.Code = EDBOPENERROR
		customError.Err = err
		return &customError
	}

	c.db = db
	return nil
}

// Close 关闭数据库连接
func (c *Client) Close() error {
	if c.db != nil {
		return c.db.Close()
	}
	return nil
}
