package main

import (
	"fmt"
	"strconv"
)

// 指数名称，当前点数，当前价格，涨跌率，成交量（手），成交额（万元）
type zhishu struct {
	stockcode       string
	stockname       string
	dangqiandianshu float64
	zhangdielv      float64
	chengjiaoliang  float64
	chengjiaoe      float64
}

func NewZS(stockcode string) (s *zhishu){
	s =&zhishu{
		stockcode: stockcode,
	}
	return s
}

func (si *zhishu) getReal() {
	sl := getinfo(si.stockcode)

	si.stockname = sl[0]
	si.dangqiandianshu,_ = strconv.ParseFloat(sl[1], 64)
	si.zhangdielv,_ = strconv.ParseFloat(sl[3], 64)
	si.chengjiaoliang,_ = strconv.ParseFloat(sl[4], 64)
	si.chengjiaoe,_ = strconv.ParseFloat(sl[5], 64)

	fmt.Printf("%s 涨跌:%v%% 当前价格:%v 成交量:%v 成交额:%v \n", si.stockname, si.zhangdielv, si.dangqiandianshu, si.chengjiaoliang/10000, si.chengjiaoe/10000)
}
