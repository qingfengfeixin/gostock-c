package main

import (
	"fmt"
	"time"
)

func gostock() {
	z1 := NewZS("s_sh000001")
	z2 := NewZS("s_sz399001")
	z3 := NewZS("s_sz399006")

	s1 := Newstock("sz000002")
	s2 := Newstock("sh601318")

	for {
		z1.getReal()
		z2.getReal()
		z3.getReal()
		fmt.Printf("大盘成交额 =%v \n", (z1.chengjiaoe+z2.chengjiaoe+z3.chengjiaoe)/10000)

		s1.getReal()
		s2.getReal()

		select {
		case <-time.NewTimer(30 * 1000 * time.Millisecond).C:
		}
	}
}

