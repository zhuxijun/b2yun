package mssql

import (
	root "b2yun/pkg"
)

// GoodsService 门店服务
type GoodsService struct {
	session *Session
}

// GetGoodss 获取门店列表
func (s *GoodsService) GetGoodss(entity string) ([]root.Goods, error) {

	var models []GoodsModel

	err := s.session.db.Select(&models, getStr(entity))

	if err != nil {
		return nil, err
	}

	var goodss []root.Goods

	goodsSNTmp := ""
	j := -1

	var rankPrice root.RankPrice

	for _, model := range models {
		if goodsSNTmp != model.GoodsSN { //商品信息重复，去重
			j++
			goodss = append(goodss, model.toGoods())
		}
		if entity == "GoodsPriceEntity" {
			//赋值嵌套子集
			rankPrice.UserRank = model.UserRank
			rankPrice.UserPrice = model.UserPrice

			goodss[j].RankPrice = append(goodss[j].RankPrice, rankPrice)

			goodsSNTmp = model.GoodsSN
		}
	}

	return goodss, nil

}

// getStr 将商品信息拆分成三部分获取（basic基本信息，price价格信息，stock库存信息）
func getStr(entity string) string {

	var sql string

	if entity == "GoodsBasicEntity" {
		sql = `   
		select 
			goods_sn = CAST(t1.fitem_id as varchar(10)),   --	商品货号(第三方商品唯一标识)
			goods_name = t1.fitem_name,    --	商品名称
			cat_sn = t1.fitem_clsno,   --	自定义类别ID
			brand_sn = t1.fitem_brdno,    --	自定义品牌ID
			suppliers_id = '0',    --货商ID;0为自营商品,其他对应货商id则为货商商品
			goods_barcode = t1.fitem_subno,   --条码; 多个条码，请用英文逗号（“,”）隔开。
			goods_size = t1.fitem_size,  --规格
			goods_unit = t1.funit_no,  --单位
			goods_area = t1.fplace,  --产地
			--goods_number = CAST(isnull(isnull(t6.fqty,t7.fqty),0) as varchar(19)),    --库存数量
			goods_price = CAST(t8.fin_price as varchar(18)), --参考进价
			--shop_price = CAST(t2.fps_price as varchar(18)),  --默认商品售价;如存在等级价格时将按等级价格显示
			--rank_price = '1.3',  --会员等级价格: user_rank会员等级ID;user_price会员等级价格;
			goods_weight = '',    --商品重量KG
			warn_number = '', --库存预警数量
			min_number = cast(t1.fdc_spec_num as varchar(10)),  --	配送倍数（订货数量按倍数递增）
			limite_number = '',   --限购数量
			is_on_sale = (case when t1.fstatus in ('5','6','7') then '1' else '0' end),  --销售状态 1上架,0下架
			is_best = ISNULL(t4.frec_flag,'0'), --是否为精品;1是,0不是
			is_hot = '0',  --	是否为热销;1是,0不是
			is_new = ISNULL(t3.fnew_flag,'0'),  --是否为新品;1是,0不是
			is_not_accumulated = '0',  --不计算配送金额（如800元起送，那么这个商品金额不包含在内，需其他商品累计达到800元方达成起送金额）1代表不计算，0代表计算
			is_not_bonus = '0',    --不可使用优惠券（红包） 1代表不可使用，0代表可使用
			is_not_rebonus = '0',  --不参与送优惠券（红包）活动 1代表不参与，0代表参与
			ftransid = ((case when isnull(t4.ftransid,0) > (case when t1.ftransid > isnull(t3.ftransid,0) then t1.ftransid else isnull(t3.ftransid,0) end) 
						then isnull(t4.ftransid,0) else (case when t1.ftransid > isnull(t3.ftransid,0) then t1.ftransid else isnull(t3.ftransid,0) end) end))
			from t_bi_master t1 
			--inner join t_bn_bi t2 on (t1.fitem_id = t2.fitem_id and t2.fbn_no = '0000')
			left join app_t_bi_new t3 on (t1.fitem_id = t3.fitem_id)
			left join app_t_bi_recommand t4 on (t1.fitem_id = t4.fitem_id)
			left join ts_t_transtype_info_mtq t5  WITH (NOLOCK) on (t5.fun_name='GoodsBasicEntity')
			--left join (select fitem_id,SUM(fqty) as fqty from t_sk_master_02 group by fitem_id) t6 on (t1.fitem_id = t6.fitem_id)
			--left join (select fitem_id,SUM(fqty) as fqty from t_sk_master_03 group by fitem_id) t7 on (t1.fitem_id = t7.fitem_id)
			left join t_bi_price t8 on (t1.fitem_id = t8.fitem_id)
		where t1.fstatus >= '5' and t1.freward_type = '0' and t1.fbom_type = '0'
		and ((case when isnull(t4.ftransid,0) > (case when t1.ftransid > isnull(t3.ftransid,0) then t1.ftransid else isnull(t3.ftransid,0) end) 
			then isnull(t4.ftransid,0) else (case when t1.ftransid > isnull(t3.ftransid,0) then t1.ftransid else isnull(t3.ftransid,0) end) end)) > t5.ftransid
		order by ((case when isnull(t4.ftransid,0) > (case when t1.ftransid > isnull(t3.ftransid,0) then t1.ftransid else isnull(t3.ftransid,0) end) 
				then isnull(t4.ftransid,0) else (case when t1.ftransid > isnull(t3.ftransid,0) then t1.ftransid else isnull(t3.ftransid,0) end) end))
		`
	} else if entity == "GoodsPriceEntity" {
		sql = `
		select 
			goods_sn = CAST(t1.fitem_id as varchar(10)),   --	商品货号(第三方商品唯一标识)
			goods_name = t1.fitem_name,    --	商品名称
			shop_price = CAST(t222.fps_price as varchar(18)), --默认商品售价;如存在等级价格时将按等级价格显示
			user_rank = t22.flevel_no,   --会员等级
			user_price = CAST(t2.fps_price as varchar(18)),  --默认商品售价;如存在等级价格时将按等级价格显示
			ftransid = (select max(ftransid) from t_bn_bi)
			from t_bi_master t1 
			inner join t_bn_bi t2 on (t1.fitem_id = t2.fitem_id)
			inner join t_bn_master t22 on (t2.fbn_no = t22.fbn_no)
			left join t_bn_bi t222 on (t1.fitem_id = t222.fitem_id and t222.fbn_no = '0000')
			left join ts_t_transtype_info_mtq t5  WITH (NOLOCK) on (t5.fun_name='GoodsPriceEntity')
		where t1.fstatus >= '5' and t1.freward_type = '0' and t1.fbom_type = '0' and isnull(t22.flevel_no,'') <> ''
		and t2.ftransid > t5.ftransid
		order by t1.fitem_id,cast(t22.flevel_no as int)
		`
	} else {
		sql = `   
		select 
			goods_sn = cast(t1.fitem_id as varchar(10)),   --	商品货号(第三方商品唯一标识)
			goods_name = t1.fitem_name,    --	商品名称
			goods_number = cast(isnull(isnull(t6.fqty,t7.fqty),0) as varchar(19)),    --库存数量
			ftransid = 		 (case when t6.ftransid > isnull(t7.ftransid,0) then t6.ftransid else isnull(t7.ftransid,0) end)
			from t_bi_master t1
			left join ts_t_transtype_info_mtq t5  WITH (NOLOCK) on (t5.fun_name='GoodsStockEntity')
			left join (select fitem_id,SUM(fqty) as fqty,MAX(ftransid) as ftransid from t_sk_master_02
			            where fitem_id in (select fitem_id from t_sk_master_02 where ftransid > (select ftransid from ts_t_transtype_info_mtq where fun_name='GoodsStockEntity')) 
			            group by fitem_id) t6 on (t1.fitem_id = t6.fitem_id)
			left join (select fitem_id,SUM(fqty) as fqty,MAX(ftransid) as ftransid from t_sk_master_03
			            where fitem_id in (select fitem_id from t_sk_master_03 where ftransid > (select ftransid from ts_t_transtype_info_mtq where fun_name='GoodsStockEntity')) 
			            group by fitem_id) t7 on (t1.fitem_id = t7.fitem_id)
		where t1.fstatus >= '5' and t1.freward_type = '0' and t1.fbom_type = '0'
		and (ISNULL(t6.fqty,0) <> 0 or ISNULL(t7.fqty,0) <> 0)
		and (case when t6.ftransid > isnull(t7.ftransid,0) then t6.ftransid else isnull(t7.ftransid,0) end) > t5.ftransid
		order by (case when t6.ftransid > isnull(t7.ftransid,0) then t6.ftransid else isnull(t7.ftransid,0) end)
		`
	}
	return sql
}
