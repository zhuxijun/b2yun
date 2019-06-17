package root

// ReqOrder 获取订单详情
type ReqOrder struct {
	OrderSN      string   `json:"order_sn"`      //商城订单号
	UserID       string   `json:"user_id"`       //商城用户唯一ID
	UserName     string   `json:"user_name"`     //商城用户名称
	ParentID     string   `json:"parent_id"`     //业务推荐人ID
	BranchNO     string   `json:"branch_no"`     //第三方关联机构(分店)编码
	OrderStatus  string   `json:"order_status"`  //订单状态122 完成（结算）,	510 已发货（出车）,530 配货中（复核）,500 已分单（打印）,100 已确认（取单）,000 待确认
	Consignee    string   `json:"consignee"`     //收货人
	Country      string   `json:"country"`       //国家
	Province     string   `json:"province"`      //省份
	City         string   `json:"city"`          //城市
	District     string   `json:"district"`      //地区
	Address      string   `json:"address"`       //详细地址
	Zipcode      string   `json:"zipcode"`       //邮政编码
	Tel          string   `json:"tel"`           //订单联系电话
	Mobile       string   `json:"mobile"`        //订单联系手机
	PostScript   string   `json:"postscript"`    //订单留言,由用户提交订单前填写
	ShippingName string   `json:"shipping_name"` //配送方式
	PayName      string   `json:"pay_name"`      //付款方式
	GoodsAmount  string   `json:"goods_amount"`  //商品金额
	ShippingFee  string   `json:"shipping_fee"`  //配送费
	InsureFee    string   `json:"insure_fee"`    //保价费用
	PayFee       string   `json:"pay_fee"`       //支付费用（如支付手续费）
	PackFee      string   `json:"pack_fee"`      //包装费用
	CardFee      string   `json:"card_fee"`      //贺卡费用
	MoneyPaid    string   `json:"money_paid"`    //已付款金额
	Bonus        string   `json:"bonus"`         //使用红包金额
	OrderAmount  string   `json:"order_amount"`  //应付款金额（订单金额）
	AddTime      string   `json:"add_time"`      //下单时间
	ConfirmTime  string   `json:"confirm_time"`  //确认时间
	PayTime      string   `json:"pay_time"`      //付款时间
	ShippingTime string   `json:"shipping_time"` //配送时间
	BonusID      string   `json:"bonus_id"`      //红包ID
	InvoiceNO    string   `json:"invoice_no"`    //第三方关联单号
	SuppliersID  string   `json:"suppliers_id"`  //货商ID的商城订单
	Detail       []Detail `json:"detail"`        //当前订单商品明细
}

// Detail 获取订单详情商品明细
type Detail struct {
	GoodsID       string `json:"goods_id"`       //商城商品唯一ID
	GoodsName     string `json:"goods_name"`     //下订单时商品名称
	GoodsSN       string `json:"goods_sn"`       //商品货号(第三方商品唯一标识)
	GoodsNumber   string `json:"goods_number"`   //商品数量
	MarketPrice   string `json:"market_price"`   //市场价格（参考）
	GoodsPrice    string `json:"goods_price"`    //商品价格（实际下单价格）
	GoodsAttr     string `json:"goods_attr"`     //商品属性
	IsReal        string `json:"is_real"`        //1实物;0虚拟商品
	ExtensionCode string `json:"extension_code"` //拓展属性："package_buy"礼包,明细对应促销信息
	ParentID      string `json:"parent_id"`      //父商品ID（当前为配件商品）
	IsGift        string `json:"is_gift"`        //是否参加优惠活动0为正常销售,大于0则对应促销活动act_id
	IsNotRebonus  string `json:"is_not_rebonus"` //1不参与返红包（优惠券）
	IsNotExtract  string `json:"is_not_extract"` //1不参与提成商品
	GoodsAttrID   string `json:"goods_attr_id"`  //商品属性id，无属性默认为0
}

// OrderCancel 取消确认订单
type OrderCancel struct {
	OrderSN string //商城订单号
}

// OrderStatus 更新物流状态
/*字段order_status参数相关
状态码	说明
100	已确认（取单）
500	已分单（打印）
530	配货中（分拣/复核）
510	已发货（出车）
122	完成（结算）*/
type OrderStatus struct {
	OrderSN     string `json:"order_sn"`              //是否必须：是	所要查询状态的订单号,不提交订单号则返回全部订单号及状态码
	OrderStatus string `json:"order_status"`          //是否必须：是	需要更新的订单状态码
	InvoiceNO   string `json:"invoice_no,omitempty"`  //是否必须：是	510 已发货（出车）状态时必须，多个发货单号，请用英文逗号（“,”）隔开。
	ActionNote  string `json:"action_note,omitempty"` //是否必须：否	操作备注信息。510 已发货（出车）状态时为配送信息如：“张三:13800138000”。
}

// Orderer 订单相关接口
type Orderer interface {
	InsertOrders([]ReqOrder) error //从小程序获取订单详情后插入erp接口
	GetOrderCancels() ([]OrderCancel, error)
	GetOrderStatuss() ([]OrderStatus, error) //更新物流状态,将erp订单状态更新至小程序

}
