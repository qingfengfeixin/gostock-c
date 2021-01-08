package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {

	gozhishu()
	time.Sleep(time.Second * 3)
	gostock()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)

	for s := range c {
		switch s {
		case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM:
			fmt.Println("Program Exit...", s)
		case syscall.SIGQUIT:
			fmt.Println("Program Quit", s)
		default:
			fmt.Println("other signal", s)
		}
	}

}
