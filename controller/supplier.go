package handler

import (
	"context"

	"crm/gen/log"
	"crm/gen/supplier"
	"crm/internal/dao"
	"crm/internal/serializer"
	"crm/internal/service"

	"go.uber.org/zap"
)

// Supplier service example implementation.
// The example methods log the requests and return zero values.
type suppliersrvc struct {
	authsrvc
	logger *log.Logger
}

// NewSupplier returns the Supplier service implementation.
func NewSupplier(logger *log.Logger) supplier.Service {
	return &suppliersrvc{authsrvc{logger: logger}, logger}
}

// 获取单个供应商
func (s *suppliersrvc) Get(ctx context.Context, p *supplier.GetPayload) (res *supplier.Supplier, err error) {
	// res = &supplier.Supplier{}
	logger := L(ctx, s.logger)
	logger.Info("supplier.Get")

	supplierService := service.NewSupplierService(ctx, dao.DB, logger)
	supplierModel, errMsg := supplierService.Get(p.ID)

	if errMsg != nil {
		logger.Error("获取单个供应商失败", zap.Error(errMsg))
		return nil, service.ErrInternalError
	}

	res = serializer.ModelSupplier2Supplier(supplierModel)
	return res, nil
}

// 获取供应商列表
func (s *suppliersrvc) List(ctx context.Context, p *supplier.ListPayload) (res *supplier.ListResult, err error) {
	res = &supplier.ListResult{}
	logger := L(ctx, s.logger)
	logger.Info("supplier.List")

	supplierService := service.NewSupplierService(ctx, dao.DB, logger)
	suppliersModel, page, errMsg := supplierService.List(*p.Limit, *p.Cursor)

	if errMsg != nil {
		logger.Error("获取供应商列表失败", zap.Error(errMsg))
		return nil, service.ErrInternalError
	}

	res.Items = serializer.ModelSuppliers2Suppliers(suppliersModel)
	res.NextCursor = page.NextCursor
	res.Total = page.TotalRecord
	return res, nil
}

// 更新供应商
func (s *suppliersrvc) Update(ctx context.Context, p *supplier.UpdatePayload) (res *supplier.Supplier, err error) {
	// res = &supplier.Supplier{}
	logger := L(ctx, s.logger)
	logger.Info("supplier.Update")

	supplierService := service.NewSupplierService(ctx, dao.DB, logger)
	supplierModel, errMsg := supplierService.Update(p)

	if errMsg != nil {
		logger.Error("更新供应商失败", zap.Error(errMsg))
		return nil, service.ErrInternalError
	}

	res = serializer.ModelSupplier2Supplier(supplierModel)
	return res, nil
}

// 创建供应商
func (s *suppliersrvc) Create(ctx context.Context, p *supplier.CreatePayload) (res *supplier.Supplier, err error) {
	// res = &supplier.Supplier{}
	logger := L(ctx, s.logger)
	logger.Info("supplier.Create")

	supplierService := service.NewSupplierService(ctx, dao.DB, logger)
	supplierModel, errMsg := supplierService.Create(p)

	if errMsg != nil {
		logger.Error("创建供应商失败", zap.Error(errMsg))
		return nil, service.ErrInternalError
	}

	res = serializer.ModelSupplier2Supplier(supplierModel)
	return res, nil
}

// 删除供应商
func (s *suppliersrvc) Delete(ctx context.Context, p *supplier.DeletePayload) (res *supplier.SuccessResult, err error) {
	res = &supplier.SuccessResult{}
	logger := L(ctx, s.logger)
	logger.Info("supplier.Delete")

	supplierService := service.NewSupplierService(ctx, dao.DB, logger)
	errMsg := supplierService.Delete(p.Ids)

	if errMsg != nil {
		logger.Error("删除供应商失败", zap.Error(errMsg))
		return nil, service.ErrInternalError
	}

	res.OK = true
	return res, nil
}
