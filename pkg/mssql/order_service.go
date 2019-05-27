package mssql

import (
	root "b2yun/pkg"
	"database/sql"
)

// OrderService 订单服务
type OrderService struct {
	session *Session
}

// InsertOrders 将获取到的会员等级插入erp数据库
func (s *OrderService) InsertOrders(datas []root.ReqOrder) error {

	var fordersn string
	for _, data := range datas {

		err := s.session.db.QueryRow("select forder_sn as fordersn from b2yun_order_master where forder_sn = ?", data.OrderSN).Scan(&fordersn)

		if err == sql.ErrNoRows {
			_, err1 := s.session.db.Exec(`insert into b2yun_order_master(
					forder_sn,fuser_id,fuser_name,fparent_id,fbranch_no,forder_status,fconsignee,fcountry,fprovince,fcity,fdistrict,faddress,fzipcode,ftel,fmobile,fpostscript,fshipping_name,fpay_name,
					fgoods_amount,fshipping_fee,finsure_fee,fpay_fee,fpack_fee,fcard_fee,fmoney_paid,fbonus,forder_amount,fadd_time,fconfirm_time,fpay_time,fshipping_time,fbonus_id,finvoice_no,fsuppliers_id
					) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`,
				data.OrderSN, data.UserID, data.UserName, data.ParentID, data.BranchNO, data.OrderStatus, data.Consignee, data.Country, data.Province, data.City,
				data.District, data.Address, data.Zipcode, data.Tel, data.Mobile, data.PostScript, data.ShippingName, data.PayName, data.GoodsAmount, data.ShippingFee, data.InsureFee,
				data.PayFee, data.PackFee, data.CardFee, data.MoneyPaid, data.Bonus, data.OrderAmount, data.AddTime, data.ConfirmTime, data.PayTime, data.ShippingTime, data.BonusID,
				data.InvoiceNO, data.SuppliersID)

			if err1 != nil {
				return err1
			}

			for _, detail := range data.Detail {
				_, err1 := s.session.db.Exec(`insert into b2yun_order_detail(forder_sn,fgoods_id,fgoods_name,fgoods_sn,fgoods_number,fmarket_price,fgoods_price,fgoods_attr,
						fis_real,fextension_code,fparent_id,fis_gift,fis_not_rebonus,fis_not_extract,fgoods_attr_id
						) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`,
					data.OrderSN, detail.GoodsID, detail.GoodsName, detail.GoodsSN, detail.GoodsNumber, detail.MarketPrice, detail.GoodsPrice, detail.GoodsAttr, detail.IsReal, detail.ExtensionCode, detail.ParentID, detail.IsGift, detail.IsNotRebonus, detail.IsNotExtract, detail.GoodsAttrID)

				if err1 != nil {
					return err1
				}
			}
		}

	}

	return nil
}

// GetOrderCancels 获取取消订单列表
func (s *OrderService) GetOrderCancels() ([]root.OrderCancel, error) {
	var models []OrderCancelModel

	err := s.session.db.Select(&models, `
			select order_sn = forder_sn 
			from b2yun_order_master
			where fstatus = '9'
			  and ftrans_flag = '0'`)

	if err != nil {
		return nil, err
	}

	var ordercancels []root.OrderCancel

	for _, model := range models {
		ordercancels = append(ordercancels, model.toOrderCancel())
	}

	return ordercancels, nil
}

// GetOrderStatuss 获取需要更新物流状态的订单列表
func (s *OrderService) GetOrderStatuss() ([]root.OrderStatus, error) {
	var models []OrderStatusModel

	err := s.session.db.Select(&models, `
			select order_sn = forder_sn
			from b2yun_order_master
			where fstatus in ('1','2','3','4','5')
			  and ftrans_flag = '0'`)

	if err != nil {
		return nil, err
	}

	var orderstatuss []root.OrderStatus

	for _, model := range models {
		orderstatuss = append(orderstatuss, model.toOrderStatus())
	}

	return orderstatuss, nil
}
