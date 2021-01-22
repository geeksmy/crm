package handler

import (
	"context"

	"crm/gen/log"
	"crm/gen/product"
	"crm/internal/dao"
	"crm/internal/serializer"
	"crm/internal/service"

	"go.uber.org/zap"
)

// Product service example implementation.
// The example methods log the requests and return zero values.
type productsrvc struct {
	authsrvc
	logger *log.Logger
}

// NewProduct returns the Product service implementation.
func NewProduct(logger *log.Logger) product.Service {
	return &productsrvc{authsrvc{logger: logger}, logger}
}

// 获取单个产品
func (s *productsrvc) Get(ctx context.Context, p *product.GetPayload) (res *product.Product, err error) {
	res = &product.Product{}
	logger := L(ctx, s.logger)
	logger.Info("product.Get")

	productService := service.NewProductService(ctx, dao.DB, logger)
	productModel, errMsg := productService.Get(p.ID)

	if errMsg != nil {
		logger.Error("获取单个产品失败", zap.Error(errMsg))
		return nil, service.ErrInternalError
	}

	res = serializer.ModelProduct2Product(productModel)
	return res, nil
}

// 获取产品列表
func (s *productsrvc) List(ctx context.Context, p *product.ListPayload) (res *product.ListResult, err error) {
	res = &product.ListResult{}
	logger := L(ctx, s.logger)
	logger.Info("product.List")

	productService := service.NewProductService(ctx, dao.DB, logger)
	productsModel, page, errMsg := productService.List(*p.Limit, *p.Cursor)

	if errMsg != nil {
		logger.Error("获取产品列表失败", zap.Error(errMsg))
		return nil, service.ErrInternalError
	}

	res.Items = serializer.ModelProducts2Products(productsModel)
	res.Total = page.TotalRecord
	res.NextCursor = page.NextCursor
	return res, nil
}

// 更新产品
func (s *productsrvc) Update(ctx context.Context, p *product.UpdatePayload) (res *product.Product, err error) {
	res = &product.Product{}
	logger := L(ctx, s.logger)
	logger.Info("product.Update")

	productService := service.NewProductService(ctx, dao.DB, logger)
	productModel, errMsg := productService.Update(p)

	if errMsg != nil {
		logger.Error("更新产品失败", zap.Error(errMsg))
		return nil, service.ErrInternalError
	}

	res = serializer.ModelProduct2Product(productModel)
	return res, nil
}

// 创建产品
func (s *productsrvc) Create(ctx context.Context, p *product.CreatePayload) (res *product.Product, err error) {
	res = &product.Product{}
	logger := L(ctx, s.logger)
	logger.Info("product.Create")

	productService := service.NewProductService(ctx, dao.DB, logger)
	productModel, errMsg := productService.Create(p)

	if errMsg != nil {
		logger.Error("创建产品失败", zap.Error(errMsg))
		return nil, service.ErrInternalError
	}

	res = serializer.ModelProduct2Product(productModel)
	return res, nil
}

// 删除产品
func (s *productsrvc) Delete(ctx context.Context, p *product.DeletePayload) (res *product.SuccessResult, err error) {
	res = &product.SuccessResult{}
	logger := L(ctx, s.logger)
	logger.Info("product.Delete")

	productService := service.NewProductService(ctx, dao.DB, logger)
	errMsg := productService.Delete(p.Ids)

	if errMsg != nil {
		logger.Error("更新产品失败", zap.Error(errMsg))
		return nil, service.ErrInternalError
	}

	res.OK = true
	return res, nil
}
