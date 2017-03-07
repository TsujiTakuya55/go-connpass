package main

import (
	"fmt"
	"github.com/TsujiTakuya55/go-connpass/connpass"
)

func main() {

	param := &connpass.Param{
		//EventId:2,
		//Keyword: "golang",
		//Ymd:20170311,
		Nickname:"Tsuji_Taku50",
		Options:connpass.Options{
			Start:1,
			Order:1,
			Count:5,
		},
	}

	client := connpass.NewClient()
	connpass, _, err := client.Get(param)

	if err != nil {
		fmt.Println(fmt.Errorf("err! \n", err))
		return
	}

	for _, k := range *connpass.Events {
		fmt.Println("eventId :", *k.EventId)
		fmt.Println("title :", *k.Title)
		fmt.Println("address :", *k.Address)
		fmt.Println("startedAt :", *k.StartedAt)
		fmt.Println("==================================")
	}
}