package main

import (
	"flag"
	"fmt"

	"github.com/evalphobia/minfraud-api-go/config"
	"github.com/evalphobia/minfraud-api-go/minfraud"
)

func main() {
	var ipaddr string
	flag.StringVar(&ipaddr, "ipaddr", "", "set target ip address")
	flag.Parse()

	conf := config.Config{
		AccountID:  "",
		LicenseKey: "",
		Debug:      true,
	}

	svc, err := minfraud.New(conf)
	if err != nil {
		panic(err)
	}

	resp, err := svc.InsightsByIP(ipaddr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("[%+v]\n", resp)
}
