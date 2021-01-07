package main

import (
	"fmt"
	"strconv"
)

type stock struct {
	stockcode    string
	stockname    string
	jinrikaipan  float64
	zuorishoupan float64
	dangqianjia  float64
	jinrizuigao  float64
	jinrizuidi   float64
	riqi         string
	shijian      string
	zhangdie     string
	zhangdielv   float64
}

func Newstock(stockcode string) (s *stock){
	s =&stock{
		stockcode: stockcode,
	}
	return s
}

func (si *stock) getReal() {

	sl := getinfo(si.stockcode)

	si.stockname = sl[0]
	si.jinrikaipan, _ = strconv.ParseFloat(sl[1], 64)
	si.zuorishoupan, _ = strconv.ParseFloat(sl[2], 64)

	if newPrice, _ := strconv.ParseFloat(sl[3], 64); newPrice > si.dangqianjia {
		si.zhangdie = "上涨"
	} else if newPrice < si.dangqianjia {
		si.zhangdie = "下跌"
	}

	si.dangqianjia, _ = strconv.ParseFloat(sl[3], 64)
	si.jinrizuigao, _ = strconv.ParseFloat(sl[4], 64)
	si.jinrizuidi, _ = strconv.ParseFloat(sl[5], 64)
	si.riqi = sl[30]
	si.shijian = sl[31]
	si.zhangdielv = (si.dangqianjia - si.zuorishoupan) / si.zuorishoupan * 100

	fmt.Printf("%s %s %s 当前涨跌方向:%s 涨跌:%v%% 当前价格:%v 开盘价格:%v 今日最高:%v 今日最低:%v \n",
		si.riqi, si.shijian, si.stockname, si.zhangdie, strconv.FormatFloat(si.zhangdielv, 'f', 2, 64),
		si.dangqianjia, si.jinrikaipan, si.jinrizuigao, si.jinrizuidi)

}
