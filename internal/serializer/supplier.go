package serializer

import (
	"crm/gen/supplier"

	"crm/internal/model"
)

func ModelSupplier2Supplier(supplierModel *model.Supplier) *supplier.Supplier {
	head := supplier.Head{
		ID:   supplierModel.Head.ID,
		Name: supplierModel.Head.Name,
	}
	founder := supplier.Founder{
		ID:   supplierModel.Founder.ID,
		Name: supplierModel.Founder.Name,
	}
	res := supplier.Supplier{
		ID:             supplierModel.ID,
		Name:           supplierModel.Name,
		Level:          supplierModel.Level,
		ContactName:    supplierModel.ContactName,
		ContactPhone:   supplierModel.ContactPhone,
		ContactAddress: supplierModel.ContactAddress,
		Note:           supplierModel.Note,
		Head:           &head,
		Founder:        &founder,
	}
	return &res
}

func ModelSuppliers2Suppliers(suppliers []*model.Supplier) []*supplier.Supplier {
	res := make([]*supplier.Supplier, 0, len(suppliers))
	for i := 0; i < len(suppliers); i++ {
		res = append(res, ModelSupplier2Supplier(suppliers[i]))
	}

	return res
}