package main

import (
	"fmt"
	"mempoolwatcher/process"
	"os"
	"time"
)

// 交易监听机器人
func main() {
	for {

		ShowMenu()
		time.Sleep(2 * time.Second)
	}

}

func ShowMenu() {
	fmt.Println("-------欢迎使用交易监听机器人-------")
	fmt.Println("\t\t\t 1.启用监听以太坊机器人")
	fmt.Println("\t\t\t 2.退出程序")
	fmt.Println("\t\t\t 3.请选择")

	var key int
	var url string
	fmt.Scanf("%d\n", &key)

	switch key {
	case 1:
		//userprocess 包处理登录
		fmt.Println("需要输入url")
		fmt.Println("请输入url:")
		fmt.Scanf("%s\n", &url)
		up := process.UserProcess{}
		go up.Watch(url)

	case 2:
		fmt.Println("退出程序")
		os.Exit(0)
	default:
		fmt.Println("请重新选择")
	}
}
