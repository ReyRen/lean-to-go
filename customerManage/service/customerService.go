package service

import (
	"lean-to-go/customerManage/model"
)

// customer add/delete/modify/find
type CustomerService struct {
	customers   []model.Customer
	customerNum int
}

func NewCustomerService() *CustomerService {
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.NewCustomer(1, "zhangsan", "male", 20, "112", "zs@122.com")
	customerService.customers = append(customerService.customers, customer)
	return customerService // default is zhangsan
}

func (this *CustomerService) List() []model.Customer {
	return this.customers
}

// add customer to []customers
func (this *CustomerService) Add(customer model.Customer) bool { // 这里一定要用*CustomerService, 因为其中有切片，必须保证每次添加的都是针对同一个切片，不然之前的customers就会都丢掉
	this.customerNum++
	customer.Id = this.customerNum
	this.customers = append(this.customers, customer)
	return true
}

func (this *CustomerService) FindById(id int) int {
	index := -1
	for i := 0; i < len(this.customers); i++ {
		if this.customers[i].Id == id {
			index = i
		}
	}
	return index
}

func (this *CustomerService) DeleteById(id int) bool {
	index := this.FindById(id)
	if index == -1 {
		return false
	}

	this.customers = append(this.customers[:index], this.customers[index+1:]...) // ...表示把切片打散为一个个，这样就能append进去整个切片了
	return true
}
