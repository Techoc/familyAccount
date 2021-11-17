package main

import "fmt"

func main() {
	//声明一个变量，保存接受用户输入的选项
	key := ""
	//声明一个变量，控制循环
	loop := true

	//定义账户的余额
	balance := 10000.0
	//每次收支的金额
	money := 0.0
	//每次收支的说明
	notes := ""
	//定义一个变量，记录是否有收支的记录
	flag := false
	//收支的一个详情使用字符串来记录
	//当有收支时，只需要对details 进行拼接处理即可
	details := "收支\t账户金额\t收支金额\t说	明"
	//显示主菜单
	for {
		fmt.Println("--------------------家庭收支记账软件--------------------")
		fmt.Println("                     1 收支明细")
		fmt.Println("                     2 登记收入")
		fmt.Println("                     3 登记支出")
		fmt.Println("                     4 退出软件")
		fmt.Print("请选择(1-4):")

		_, err := fmt.Scanln(&key)
		if err != nil {
			return
		}
		switch key {
		case "1":
			fmt.Println("--------------------收支明细--------------------")
			if flag {
				fmt.Println(details)
			} else {
				fmt.Println("当前没有收支明细，请先进行收支")
			}
		case "2":
			fmt.Print("本次收入金额:")
			_, err := fmt.Scanln(&money)
			if err != nil {
				return
			}
			balance += money //修改账户余额
			fmt.Print("本次收入说明:")
			_, err = fmt.Scanln(&notes)
			if err != nil {
				return
			}
			//将这个收入情况，拼接到details变量
			//收入	11000	100		有人发红包
			details += fmt.Sprintf("\n收入\t%v\t%v\t%v", balance, money, notes)
			flag = true
		case "3":
			fmt.Print("本次支出金额:")
			_, err := fmt.Scanln(&money)
			if err != nil {
				return
			}
			//这里需要做一个必要的判断
			if money > balance {
				fmt.Println("余额的金额不足")
				break
			}
			balance -= money //修改账户余额
			fmt.Print("本次支出说明:")
			_, err = fmt.Scanln(&notes)
			if err != nil {
				return
			}
			//将这个支出情况，拼接到details变量
			//支出	11000	-100		买了烟
			details += fmt.Sprintf("\n支出\t%v\t%v\t%v", balance, money, notes)
			flag = true
		case "4":
			fmt.Println("你确定要退出吗？ y/n")
			choices := ""
			for {
				_, err := fmt.Scanln(&choices)
				if err != nil {
					return
				}
				if choices == "y" || choices == "n" {
					break
				}
				fmt.Println("你的输入有误，请重新输入")
			}
			if choices == "y" {
				loop = false
			}
		default:
			fmt.Println("请输入正确的选项！")
		}

		if !loop {
			break
		}
	}
	fmt.Println("已退出家庭记账软件的使用！")
}
