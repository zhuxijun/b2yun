package main

import (
	root "b2yun/pkg"
	"b2yun/pkg/config"
	"b2yun/pkg/http"
	"b2yun/pkg/mssql"
)

// App 应用程序对象
type App struct {
	config      *config.Config
	log         *root.Log
	mssqlclient *mssql.Client
	httpClient  *http.Client
}

// Initialize 初始化
func (a *App) Initialize() {
	config := config.NewConfig("app_config.json")
	a.config = config

	log, _ := root.NewLogrotate(config)
	a.log = log

	mssqlClient := mssql.NewClient(config)
	a.mssqlclient = mssqlClient

	httpClient := http.NewClient(config, log)
	a.httpClient = httpClient

}

// Run 运行
func (a *App) Run() error {

	a.log.Logger.Info("开始运行...")

	if err := a.mssqlclient.Open(); err != nil {
		return err
	}
	defer a.mssqlclient.Close()

	if err := a.httpClient.Open(); err != nil {
		return err
	}

	cater := a.mssqlclient.Connect().CatService()
	brander := a.mssqlclient.Connect().BrandService()
	goodser := a.mssqlclient.Connect().GoodsService()
	orderer := a.mssqlclient.Connect().OrderService()
	memberer := a.mssqlclient.Connect().MemberService()

	task := a.mssqlclient.Connect().TaskService()

	basicer := a.config

	option := http.ServiceOption{
		Cater:    cater,
		Brander:  brander,
		Goodser:  goodser,
		Orderer:  orderer,
		Memberer: memberer,

		Basicer: basicer,
		Tasker:  task,
	}

	a.httpClient.InitService(option)

	// //一、商品相关
	// //1)分类信息传输POST
	// catService := a.httpClient.CatService()
	// if err := catService.UploadCats(); err != nil {
	// 	return err
	// }

	// //2)品牌信息传输POST
	// brandService := a.httpClient.BrandService()
	// if err := brandService.UploadBrands(); err != nil {
	// 	return err
	// }

	// //3)商品信息传输POST
	// goodsService := a.httpClient.GoodsService()
	// if err := goodsService.UploadGoodss(); err != nil {
	// 	return err
	// }

	// //二、订单相关
	// orderService := a.httpClient.OrderService()
	// //1)获取订单详情GET
	// if err := orderService.DownOrderInfo(); err != nil {
	// 	return err
	// }
	// //2)取消确认订单GET
	// if err := orderService.CancelOrder(); err != nil {
	// 	return err
	// }

	// //3)更新物流状态POST
	// if err := orderService.UploadOrderInfo(); err != nil {
	// 	return err
	// }

	//三、会员相关
	memberService := a.httpClient.MemberService()
	//1)获取会员等级GET
	if err := memberService.DownloadMemberLevel(); err != nil {
		return err
	}
	//2)获取会员信息GET
	if err := memberService.DownloadMemberInfo(); err != nil {
		return err
	}
	//3)更新会员信息POST
	if err := memberService.UploadMemberInfo(); err != nil {
		return err
	}

	a.log.Logger.Info("结束运行...")

	return nil
}
