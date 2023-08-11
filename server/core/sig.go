package core

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func SigInit() {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM,
		syscall.SIGQUIT)
	go func() {
		for s := range c {
			switch s {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				fmt.Println("Program Exit...", s)
				os.Exit(0)
			default:
				fmt.Println("other signal", s)
			}
		}
	}()
}
