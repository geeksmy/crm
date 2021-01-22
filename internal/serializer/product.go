package serializer

import (
	"crm/gen/product"
	"crm/internal/model"
)

func ModelProduct2Product(productModel *model.Product) *product.Product {
	founder := product.Founder{
		ID:   productModel.Founder.ID,
		Name: productModel.Founder.Name,
	}

	res := product.Product{
		ID:          productModel.ID,
		Name:        productModel.Name,
		Unit:        productModel.Unit,
		CostPrice:   productModel.CostPrice,
		MarketPrice: productModel.MarketPrice,
		Note:        productModel.Note,
		Image:       productModel.Image,
		Code:        productModel.Code,
		Size:        productModel.Size,
		Type:        productModel.Type,
		IsShelves:   productModel.IsShelves,
		Founder:     &founder,
	}

	return &res
}

func ModelProducts2Products(productsModel []*model.Product) []*product.Product {
	res := make([]*product.Product, 0, len(productsModel))
	for i := 0; i < len(productsModel); i++ {
		res = append(res, ModelProduct2Product(productsModel[i]))
	}

	return res
}
