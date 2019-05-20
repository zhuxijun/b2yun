package mssql

import (
	"time"

	root "b2yun/pkg"

	"github.com/jmoiron/sqlx"
)

// Session 数据库连接
type Session struct {
	db  *sqlx.DB
	now time.Time // 提供同一个会话中的数据库操作时间戳

	catService    CatService
	brandService  BrandService
	goodsService  GoodsService
	orderSercice  OrderService
	memberService MemberService

	taskService TaskService
}

// NewSession 创建链接
func NewSession(db *sqlx.DB) *Session {

	s := &Session{db: db}

	s.catService.session = s
	s.brandService.session = s
	s.goodsService.session = s
	s.orderSercice.session = s
	s.memberService.session = s

	s.taskService.session = s

	return s
}

// CatService 分类信息服务
func (s *Session) CatService() root.Cater {
	return &s.catService
}

// BrandService 品牌信息服务
func (s *Session) BrandService() root.Brander {
	return &s.brandService
}

// GoodsService 商品信息服务
func (s *Session) GoodsService() root.Goodser {
	return &s.goodsService
}

// OrderService 商品信息服务
func (s *Session) OrderService() root.Orderer {
	return &s.orderSercice
}

// MemberService 商品信息服务
func (s *Session) MemberService() root.Memberer {
	return &s.memberService
}

// TaskService 任务服务
func (s *Session) TaskService() root.Tasker {
	return &s.taskService
}
