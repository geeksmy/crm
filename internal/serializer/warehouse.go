package serializer

import (
	"crm/gen/warehouse"
	"crm/internal/model"
)

func ModelWarehouse2Warehouse(warehouseModel *model.Warehouse) *warehouse.Warehouse {
	founder := warehouse.Founder{
		ID:   warehouseModel.Founder.ID,
		Name: warehouseModel.Founder.Name,
	}

	res := warehouse.Warehouse{
		ID:      warehouseModel.ID,
		Name:    warehouseModel.Name,
		Code:    warehouseModel.Code,
		Address: warehouseModel.Address,
		Type:    warehouseModel.Type,
		Founder: &founder,
	}

	return &res
}

func ModelWarehouses2Warehouses(warehousesModel []*model.Warehouse) []*warehouse.Warehouse {
	res := make([]*warehouse.Warehouse, 0, len(warehousesModel))

	for i := 0; i < len(warehousesModel); i++ {
		res = append(res, ModelWarehouse2Warehouse(warehousesModel[i]))
	}

	return res
}
