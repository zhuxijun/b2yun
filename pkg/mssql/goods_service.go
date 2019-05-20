package mssql

import (
	root "b2yun/pkg"
)

// GoodsService 门店服务
type GoodsService struct {
	session *Session
}

// GetGoodss 获取门店列表
func (s *GoodsService) GetGoodss() ([]root.Goods, error) {

	var models []GoodsModel

	err := s.session.db.Select(&models,
		`   select 
		goods_sn = cast(t1.fitem_id as varchar(10)),   --	商品货号(第三方商品唯一标识)
		goods_name = t1.fitem_name,    --	商品名称
		cat_sn = t1.fitem_clsno,   --	自定义类别ID
		brand_sn = t1.fitem_brdno,    --	自定义品牌ID
		suppliers_id = '0',    --货商ID;0为自营商品,其他对应货商id则为货商商品
		goods_barcode = t1.fitem_subno,   --条码; 多个条码，请用英文逗号（“,”）隔开。
		goods_size = t1.fitem_size,  --规格
		goods_unit = t1.funit_no,  --单位
		goods_area = t1.fplace,  --产地
		goods_price = '', --参考进价
		goods_weight = '',    --商品重量KG
		warn_number = '', --库存预警数量
		min_number = cast(t1.fdc_spec_num as varchar(10)),  --	配送倍数（订货数量按倍数递增）
		limite_number = '',   --限购数量
		is_on_sale = (case when t1.fstatus in ('5','6','7') then '1' else '0' end),  --销售状态 1上架,0下架
		is_hot = '0',  --	是否为热销;1是,0不是
		is_not_accumulated = '0',  --不计算配送金额（如800元起送，那么这个商品金额不包含在内，需其他商品累计达到800元方达成起送金额）1代表不计算，0代表计算
		is_not_bonus = '0',    --不可使用优惠券（红包） 1代表不可使用，0代表可使用
		is_not_rebonus = '0'  --不参与送优惠券（红包）活动 1代表不参与，0代表参与
		from t_bi_master t1 
		left join ts_t_transtype_info_mtq t5  WITH (NOLOCK) on (t5.fun_name='GoodsEntity')
		where t1.fstatus >= '5'
		  and t1.ftransid > t5.ftransid`)

	if err != nil {
		return nil, err
	}

	var goodss []root.Goods

	for _, model := range models {
		goodss = append(goodss, model.toGoods())
	}

	return goodss, nil

}
