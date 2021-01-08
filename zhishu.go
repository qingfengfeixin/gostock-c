package main

import (
	"fmt"
	"strconv"
	"time"
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

func NewZS(stockcode string) (s *zhishu) {
	s = &zhishu{
		stockcode: stockcode,
	}
	return s
}

func (si *zhishu) getReal() {
	sl := getinfo(si.stockcode)

	si.stockname = sl[0]
	si.dangqiandianshu, _ = strconv.ParseFloat(sl[1], 64)
	si.zhangdielv, _ = strconv.ParseFloat(sl[3], 64)
	si.chengjiaoliang, _ = strconv.ParseFloat(sl[4], 64)
	si.chengjiaoe, _ = strconv.ParseFloat(sl[5], 64)

	fmt.Printf("%s 涨跌:%v%% 当前价格:%v 成交量:%v 成交额:%v \n", si.stockname, si.zhangdielv,
		si.dangqiandianshu, si.chengjiaoliang/10000, si.chengjiaoe/10000)
}

func gozhishu() {

	z := []zhishu{*NewZS("s_sh000001"), *NewZS("s_sz399001"), *NewZS("s_sz399006")}

	go func(z []zhishu){
		for {
			fmt.Println("----------------------------------------------------------------------------------------")
			z[0].getReal() //上证
			z[1].getReal() //深证
			z[2].getReal() //创业板
			fmt.Printf("证券成交额 =%v \n", (z[0].chengjiaoe+z[1].chengjiaoe+z[2].chengjiaoe)/10000)

			select {
			case <-time.NewTimer(20 * 1000 * time.Millisecond).C:
			}
		}
	}(z)

}
