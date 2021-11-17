package utils

import "fmt"

type FamilyAccount struct {
	//声明必须的字段.

	//声明一个字段，保存接受用户输入的选项
	key string
	//声明一个字段，控制循环
	loop bool

	//定义账户的余额
	balance float64
	//每次收支的金额
	money float64
	//每次收支的说明
	notes string
	//定义一个字段，记录是否有收支的记录
	flag bool
	//收支的一个详情使用字符串来记录
	//当有收支时，只需要对details 进行拼接处理即可
	details string
}

//编写工厂模式的构造方法，返回一个*FamilyAccount实例
func NewFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		key:     "",
		loop:    true,
		balance: 10000.0,
		money:   0.0,
		notes:   "",
		flag:    false,
		details: "收支\t账户金额\t收支金额\t说   明",
	}

}

//将显示明细写成一个方法
func (this *FamilyAccount) showDetails() {
	fmt.Println("--------------------收支明细--------------------")
	if this.flag {
		fmt.Println(this.details)
	} else {
		fmt.Println("当前没有收支明细，请先进行收支")
	}
}

//将登记收入写成一个方法，和*FamilyAccount绑定
func (this *FamilyAccount) income() {
	fmt.Print("本次收入金额:")
	_, err := fmt.Scanln(&this.money)
	if err != nil {
		return
	}
	this.balance += this.money //修改账户余额
	fmt.Print("本次收入说明:")
	_, err = fmt.Scanln(&this.notes)
	if err != nil {
		return
	}
	//将这个收入情况，拼接到details变量
	//收入	11000	100		有人发红包
	this.details += fmt.Sprintf("\n收入\t%v\t%v\t%v", this.balance, this.money, this.notes)
	this.flag = true
}

//将登记支出写成一个方法，和*FamilyAccount绑定
func (this *FamilyAccount) pay() {
	fmt.Print("本次支出金额:")
	_, err := fmt.Scanln(&this.money)
	if err != nil {
		return
	}
	//这里需要做一个必要的判断
	if this.money > this.balance {
		fmt.Println("余额的金额不足")
		return
	}
	this.balance -= this.money //修改账户余额
	fmt.Print("本次支出说明:")
	_, err = fmt.Scanln(&this.notes)
	if err != nil {
		return
	}
	//将这个支出情况，拼接到details变量
	//支出	11000	-100		买了烟
	this.details += fmt.Sprintf("\n支出\t%v\t%v\t%v", this.balance, this.money, this.notes)
}

//将退出系统写成一个方法，和*FamilyAccount绑定
func (this *FamilyAccount) exit() {
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
		this.loop = false
	}
}

//给该结构体绑定相应的方法
//显示主菜单
func (this *FamilyAccount) MainMenu() {
	for {
		fmt.Println("--------------------家庭收支记账软件--------------------")
		fmt.Println("                     1 收支明细")
		fmt.Println("                     2 登记收入")
		fmt.Println("                     3 登记支出")
		fmt.Println("                     4 退出软件")
		fmt.Print("请选择(1-4):")

		_, err := fmt.Scanln(&this.key)
		if err != nil {
			return
		}
		switch this.key {
		case "1":
			this.showDetails()
		case "2":
			this.income()
		case "3":
			this.pay()
		case "4":
			this.exit()
		default:
			fmt.Println("请输入正确的选项！")
		}

		if !this.loop {
			break
		}
	}
	fmt.Println("已退出家庭记账软件的使用！")
}
