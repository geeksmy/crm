package serializer

import (
	"crm/gen/customer"
	"crm/internal/model"
)

func ModelCustomer2Customer(customerModel *model.Customer) *customer.Customer {
	res := customer.Customer{
		ID:       customerModel.ID,
		Name:     customerModel.Name,
		Src:      customerModel.Src,
		Mobile:   customerModel.Mobile,
		URL:      customerModel.Url,
		Email:    customerModel.Email,
		Industry: customerModel.Industry,
		Level:    customerModel.Level,
		Note:     customerModel.Note,
		Address:  customerModel.Address,
	}

	return &res
}

func ModelCustomers2Customers(customersModel []*model.Customer) []*customer.Customer {
	res := make([]*customer.Customer, 0, len(customersModel))
	for i := 0; i < len(customersModel); i++ {
		res = append(res, ModelCustomer2Customer(customersModel[i]))
	}

	return res
}
