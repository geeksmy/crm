package serializer

import (
	"crm/gen/procurement"
	"crm/internal/model"
)

func ModelProcurement2Procurement(procurementModel *model.Procurement) *procurement.Procurement {
	supplier := procurement.Supplier{
		ID:             procurementModel.Supplier.ID,
		Name:           procurementModel.Supplier.Name,
		Level:          procurementModel.Supplier.Level,
		ContactName:    procurementModel.Supplier.ContactName,
		ContactPhone:   procurementModel.Supplier.ContactPhone,
		ContactAddress: procurementModel.Supplier.ContactAddress,
		Note:           procurementModel.Supplier.Note,
	}

	head := procurement.Head{
		ID:   procurementModel.Head.ID,
		Name: procurementModel.Head.Name,
	}

	founder := procurement.Founder{
		ID:   procurementModel.Founder.ID,
		Name: procurementModel.Founder.Name,
	}

	res := procurement.Procurement{
		ID:            procurementModel.ID,
		Supplier:      &supplier,
		Code:          procurementModel.Code,
		Money:         procurementModel.Money,
		IsSalesReturn: procurementModel.IsSalesReturn,
		Note:          procurementModel.Note,
		Head:          &head,
		Founder:       &founder,
	}

	return &res
}

func ModelProcurements2Procurements(procurementsModel []*model.Procurement) []*procurement.Procurement {
	res := make([]*procurement.Procurement, 0, len(procurementsModel))

	for i := 0; i < len(procurementsModel); i++ {
		res = append(res, ModelProcurement2Procurement(procurementsModel[i]))
	}

	return res
}
