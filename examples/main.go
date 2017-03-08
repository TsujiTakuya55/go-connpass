package main

import (
	"fmt"
	"github.com/TsujiTakuya55/go-connpass/connpass"
)

func main() {

}

// イベントID=2 で検索
func eventId() {
	param := &connpass.Param{
		EventId:2,
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
	}
}

// 「golang」というキーワードで検索＆表示件数を5件にしたいとき
func keyword() {
	param := &connpass.Param{
		Keyword: "golang",
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

// 「Tsuji_Taku50」というユーザー名で検索したい＆表示件数を10件にしたいとき
func nickname() {
	param := &connpass.Param{
		Nickname: "Tsuji_Taku50",
		Options:connpass.Options{
			Start:1,
			Order:1,
			Count:10,
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

// 2017年3月に開催されるイベントを10件表示させたいとき
func ym() {
	param := &connpass.Param{
		Ym: 201703,
		Options:connpass.Options{
			Start:1,
			Order:1,
			Count:10,
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

// 2017年3月11日に開催されるイベントを7件表示させたいとき
func ymd() {
	param := &connpass.Param{
		Ymd: 20170311,
		Options:connpass.Options{
			Start:1,
			Order:1,
			Count:7,
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

// 2017年3月に開催されるイベント＆「golang」というキーワードで検索＆7件表示させたいとき
func keywordYm() {
	param := &connpass.Param{
		Keyword:"golang",
		Ym: 201703,
		Options:connpass.Options{
			Start:1,
			Order:1,
			Count:7,
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