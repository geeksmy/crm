package service

import (
	"context"
	"time"

	"crm/gen/sales"
	"crm/internal/dao"
	"crm/internal/model"

	libsgorm "github.com/geeksmy/go-libs/gorm"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

func NewSalesService(ctx context.Context, db *gorm.DB, logger *zap.Logger) *SalesService {
	return &SalesService{
		ctx:    ctx,
		db:     db,
		logger: logger,
	}
}

type SalesService struct {
	ctx    context.Context
	db     *gorm.DB
	logger *zap.Logger
}

func (srv SalesService) Get(pk string) (*model.Sales, error) {
	salesDao := dao.NewSalesDao(srv.ctx, srv.db)
	scopes := dao.NewSalesScope()
	scopes = scopes.WithPk(pk)

	res, err := salesDao.Get(scopes)

	if err != nil {
		srv.logger.Error("获取销售失败", zap.Error(err))
		return nil, err
	}

	res, err = srv.getByCustomerAndFounderAndHead(res)

	if err != nil {
		srv.logger.Error("获取用于返回值失败", zap.Error(err))
		return nil, err
	}

	return res, nil
}

func (srv SalesService) List(limit, cursor int) ([]*model.Sales, *libsgorm.Page, error) {
	var salessModel []*model.Sales
	salesDao := dao.NewSalesDao(srv.ctx, srv.db)
	scopes := dao.NewSalesScope()

	res, page, err := salesDao.List(limit, cursor, scopes)

	if err != nil {
		srv.logger.Error("获取销售列表失败", zap.Error(err))
		return nil, nil, err
	}

	for i := 0; i < len(res); i++ {
		salesModel, err := srv.getByCustomerAndFounderAndHead(res[i])

		if err != nil {
			srv.logger.Error("获取用于返回值失败", zap.Error(err))
			return nil, nil, err
		}

		salessModel = append(salessModel, salesModel)
	}

	return salessModel, page, nil
}

func (srv SalesService) Create(p *sales.CreatePayload) (*model.Sales, error) {
	salesDao := dao.NewSalesDao(srv.ctx, srv.db)

	consignmentDate, err := time.Parse("2006-01-02 15:04:05", p.ConsignmentDate)

	if err != nil {
		srv.logger.Error("string转time失败", zap.Error(err))
		return nil, err
	}

	res := &model.Sales{
		Name:            p.Name,
		Code:            p.Code,
		CustomerID:      p.CustomerID,
		Money:           p.Money,
		ConsignmentDate: consignmentDate,
		IsSalesReturn:   p.IsSalesReturn,
		Note:            p.Note,
		HeadID:          p.HeadID,
		FounderID:       p.FounderID,
	}

	res, err = salesDao.Create(*res)

	if err != nil {
		srv.logger.Error("创建销售失败", zap.Error(err))
		return nil, err
	}

	res, err = srv.getByCustomerAndFounderAndHead(res)

	if err != nil {
		srv.logger.Error("获取用于返回值失败", zap.Error(err))
		return nil, err
	}

	return res, nil
}

func (srv SalesService) Update(p *sales.UpdatePayload) (*model.Sales, error) {
	salesDao := dao.NewSalesDao(srv.ctx, srv.db)
	scopes := dao.NewSalesScope()
	scopes = scopes.WithPk(p.ID)

	fields := make(map[string]interface{})

	res, err := salesDao.Update(scopes, fields)

	if err != nil {
		srv.logger.Error("创建销售失败", zap.Error(err))
		return nil, err
	}

	res, err = srv.getByCustomerAndFounderAndHead(res)

	if err != nil {
		srv.logger.Error("获取用于返回值失败", zap.Error(err))
		return nil, err
	}

	return res, nil
}

func (srv SalesService) Delete(pks []string) error {
	salesDao := dao.NewSalesDao(srv.ctx, srv.db)

	err := salesDao.Delete(pks)

	return err
}

func (srv SalesService) getByCustomerAndFounderAndHead(salesModel *model.Sales) (*model.Sales, error) {
	userDao := dao.NewUserDao(srv.ctx, srv.db)
	customerDao := dao.NewCustomerDao(srv.ctx, srv.db)

	head, err := userDao.GetByID(salesModel.HeadID)

	if err != nil {
		srv.logger.Error("获取负责人失败", zap.Error(err))
		return nil, err
	}

	founder, err := userDao.GetByID(salesModel.FounderID)

	if err != nil {
		srv.logger.Error("获取创建人失败", zap.Error(err))
		return nil, err
	}

	customer, err := customerDao.GetByID(salesModel.CustomerID)

	if err != nil {
		srv.logger.Error("获取客户失败", zap.Error(err))
		return nil, err
	}

	salesModel.Founder.ID = founder.ID
	salesModel.Founder.Name = founder.Name
	salesModel.Head.ID = head.ID
	salesModel.Head.Name = head.Name
	salesModel.Customer = *customer

	return salesModel, nil
}
