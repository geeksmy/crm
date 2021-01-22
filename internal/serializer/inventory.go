package serializer

import (
	"crm/gen/inventory"
	"crm/internal/model"
)

func ModelInventory2Inventory(inventoryModel *model.Inventory) *inventory.Inventory {
	product := inventory.Product{
		ID:          inventoryModel.Product.ID,
		Name:        inventoryModel.Product.Name,
		Unit:        inventoryModel.Product.Unit,
		CostPrice:   inventoryModel.Product.CostPrice,
		MarketPrice: inventoryModel.Product.MarketPrice,
		Note:        inventoryModel.Product.Note,
		Image:       inventoryModel.Product.Image,
		Code:        inventoryModel.Product.Code,
		Size:        inventoryModel.Product.Size,
		Type:        inventoryModel.Product.Type,
		IsShelves:   inventoryModel.Product.IsShelves,
	}
	warehouse := inventory.Warehouse{
		ID:      inventoryModel.Warehouse.ID,
		Name:    inventoryModel.Warehouse.Name,
		Code:    inventoryModel.Warehouse.Code,
		Address: inventoryModel.Warehouse.Address,
		Type:    inventoryModel.Warehouse.Type,
	}

	head := inventory.Head{
		ID:   inventoryModel.Head.ID,
		Name: inventoryModel.Head.Name,
	}

	founder := inventory.Founder{
		ID:   inventoryModel.Founder.ID,
		Name: inventoryModel.Founder.Name,
	}

	res := inventory.Inventory{
		ID:            inventoryModel.ID,
		Product:       &product,
		Number:        inventoryModel.Number,
		Code:          inventoryModel.Code,
		Warehouse:     &warehouse,
		Type:          inventoryModel.Type,
		InventoryDate: inventoryModel.InventoryDate.Format("2006-01-02 15:04:05"),
		InAndOut:      inventoryModel.InAndOut,
		Note:          inventoryModel.Note,
		Head:          &head,
		Founder:       &founder,
	}

	return &res
}

func ModelInventorys2Inventorys(inventorysModel []*model.Inventory) []*inventory.Inventory {
	res := make([]*inventory.Inventory, 0, len(inventorysModel))

	for i := 0; i < len(inventorysModel); i++ {
		res = append(res, ModelInventory2Inventory(inventorysModel[i]))
	}

	return res
}
