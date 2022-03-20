package service

import (
	"fmt"
	"lib/src/go_code/customerManage/model"
)

// CustomerService 主要完成对customer的操作
type CustomerService struct {
	customers   []model.Customer
	customerNum int
}

func NewCustomerService() *CustomerService {
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(1, "张三", "男", 30, "13992984775", "388764@qq.com")
	customerService.customers = append(customerService.customers, customer)
	return customerService
}

func (this *CustomerService) List() []model.Customer {
	return this.customers
}

func (this *CustomerService) Add(cus model.Customer) {
	this.customerNum++
	cus.Id = this.customerNum
	this.customers = append(this.customers, cus)
}

func (this *CustomerService) Delete(id int) bool {
	index := this.FindById(id)
	if index == -1 {
		fmt.Println("当前ID对应的客户不存在，无法删除...")
		return false
	}
	this.customers = append(this.customers[:index], this.customers[index+1:]...)
	return true

}

func (this *CustomerService) FindById(id int) int {
	index := -1
	for i := 0; i < len(this.customers); i++ {
		if id == this.customers[i].Id {
			index = i
		}
	}
	return index
}
