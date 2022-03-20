package main

import (
	"fmt"
	"lib/src/go_code/customerManage/model"
	"lib/src/go_code/customerManage/service"
)

type customerView struct {
	key             string
	loop            bool
	customerService *service.CustomerService
}

//显示所有客户信息
func (cv *customerView) list() {
	customerList := cv.customerService.List()
	if len(customerList) > 0 {
		fmt.Println("---------------------客户列表Begin-------------------")
		fmt.Println("编号\t姓名\t性别\t年龄\t电话\t\t邮箱")
		for i := 0; i < len(customerList); i++ {
			fmt.Println(customerList[i].GetInfo())
		}
		fmt.Println("---------------------客户列表End--------------------")
		fmt.Println()

	} else {
		fmt.Println("平台暂时还没有客户信息")
	}
}

//新增客户

func (cv *customerView) addCustomer() {
	var name string
	fmt.Print("请输入客户姓名:")
	fmt.Scanln(&name)
	var gender string
	fmt.Print("请输入客户性别:")
	fmt.Scanln(&gender)
	var age int
	fmt.Print("请输入客户年龄:")
	fmt.Scanln(&age)
	var phone string
	fmt.Print("请输入客户电话:")
	fmt.Scanln(&phone)
	var email string
	fmt.Print("请输入客户邮箱:")
	fmt.Scanln(&email)
	cus := model.Customer{
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
	cv.customerService.Add(cus)
	fmt.Println("客户添加成功....")
}

func (cv *customerView) delete() {
	fmt.Println("---------------------删除客户-------------------")
	var id int
	fmt.Print("请输入需要删除的客户ID[-1退出]：")
	fmt.Scanln(&id)
	if id == -1 {
		return
	} else {
		var key string
		fmt.Print("您是否确定要删除当前客户？y/n:")
		fmt.Scanln(&key)
		if key == "y" || key == "Y" {
			if cv.customerService.Delete(id) {
				fmt.Println("---------------------删除客户完成-------------------")
			} else {
				fmt.Println("---------------------删除客户失败-------------------")
			}
		}
	}
}
func (cv *customerView) update() {
	fmt.Println("该功能还未上线，敬请期待....")
}
func (cv *customerView) mainMenu() {
	for {
		fmt.Println("<<<<<<<<<<<<<<<<<<<客户信息管理系统>>>>>>>>>>>>>>>>>")
		fmt.Println("		1.添加客户")
		fmt.Println("		2.修改客户")
		fmt.Println("		3.删除客户")
		fmt.Println("		4.客户列表")
		fmt.Println("		5.退   出")
		fmt.Println()
		fmt.Print("请选择(1-5):")

		fmt.Scanln(&cv.key)
		switch cv.key {
		case "1":
			cv.addCustomer()
		case "2":
			cv.update()
		case "3":
			cv.delete()
		case "4":
			cv.list()
		case "5":
			cv.loop = false
		default:
			fmt.Printf("你的输入有误，请重新输入...")
		}
		if !cv.loop {
			break
		}
	}
	fmt.Println("您退出了用户管理系统...")
}
func main() {
	var customerView = customerView{
		key:  "",
		loop: true,
	}
	customerView.customerService = service.NewCustomerService()
	customerView.mainMenu()
}
