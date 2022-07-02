package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"cs101/cmd"
	"cs101/pkg/utils"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	cp := os.Getenv("CONFIG_PATH")
	if cp == "" {
		log.Fatalf("config path is empty")
	}
	cfg := cmd.Config{}
	if err := utils.Load(cp, &cfg); err != nil {
		log.Fatalf("load config : %+v", err)
	}
	s := cmd.New(cfg)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGTERM, os.Interrupt)

	go s.Run()
	<-interrupt
	s.Shutdown()
}
