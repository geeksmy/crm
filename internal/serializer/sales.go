package serializer

import (
	"crm/gen/sales"
	"crm/internal/model"
)

func ModelSales2Sales(salesModel *model.Sales) *sales.Sales {
	customer := sales.Customer{
		ID:       salesModel.Customer.ID,
		Name:     salesModel.Customer.Name,
		Src:      salesModel.Customer.Src,
		Mobile:   salesModel.Customer.Mobile,
		URL:      salesModel.Customer.Url,
		Email:    salesModel.Customer.Email,
		Industry: salesModel.Customer.Industry,
		Level:    salesModel.Customer.Level,
		Note:     salesModel.Customer.Note,
		Address:  salesModel.Customer.Address,
	}

	head := sales.Head{
		ID:   salesModel.Head.ID,
		Name: salesModel.Head.Name,
	}

	founder := sales.Founder{
		ID:   salesModel.Founder.ID,
		Name: salesModel.Founder.Name,
	}

	res := sales.Sales{
		ID:              salesModel.ID,
		Name:            salesModel.Name,
		Code:            salesModel.Code,
		Customer:        &customer,
		Money:           salesModel.Money,
		ConsignmentDate: salesModel.ConsignmentDate.Format("2006-01-02 15:04:05"),
		IsSalesReturn:   salesModel.IsSalesReturn,
		Note:            salesModel.Note,
		Head:            &head,
		Founder:         &founder,
	}

	return &res
}

func ModelSaless2Saless(salessModel []*model.Sales) []*sales.Sales {
	res := make([]*sales.Sales, 0, len(salessModel))
	for i := 0; i < len(salessModel); i++ {
		res = append(res, ModelSales2Sales(salessModel[i]))
	}

	return res
}
