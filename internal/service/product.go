package service

import (
	"context"

	"crm/gen/product"
	"crm/internal/dao"
	"crm/internal/model"

	libsgorm "github.com/geeksmy/go-libs/gorm"
	"github.com/jinzhu/gorm"
	"go.uber.org/zap"
)

func NewProductService(ctx context.Context, db *gorm.DB, logger *zap.Logger) *ProductService {
	return &ProductService{
		ctx:    ctx,
		db:     db,
		logger: logger,
	}
}

type ProductService struct {
	ctx    context.Context
	db     *gorm.DB
	logger *zap.Logger
}

func (srv ProductService) Get(pk string) (*model.Product, error) {
	productDao := dao.NewProductDao(srv.ctx, srv.db)
	scopes := dao.NewProductScope()
	scopes = scopes.WithPk(pk)

	res, err := productDao.Get(scopes)

	if err != nil {
		srv.logger.Error("获取单个产品失败", zap.Error(err))
		return nil, err
	}

	product, err := srv.getUserByFounder(res)

	if err != nil {
		srv.logger.Error("获取创建人失败", zap.Error(err))
		return nil, err
	}

	return product, nil
}

func (srv ProductService) List(limit, cursor int) ([]*model.Product, *libsgorm.Page, error) {
	var products []*model.Product
	productDao := dao.NewProductDao(srv.ctx, srv.db)
	scopes := dao.NewProductScope()

	res, page, err := productDao.List(limit, cursor, scopes)

	if err != nil {
		srv.logger.Error("获取单个产品失败", zap.Error(err))
		return nil, nil, err
	}

	for i := 0; i < len(res); i++ {
		product, err := srv.getUserByFounder(res[i])

		if err != nil {
			srv.logger.Error("获取创建人失败", zap.Error(err))
			return nil, nil, err
		}

		products = append(products, product)
	}

	return products, page, nil
}

func (srv ProductService) Create(p *product.CreatePayload) (*model.Product, error) {
	productDao := dao.NewProductDao(srv.ctx, srv.db)

	productModel := model.Product{
		Name:          p.Name,
		Unit:          p.Unit,
		CostPrice:     p.CostPrice,
		MarketPrice:   p.MarketPrice,
		Note:          p.Note,
		Image:         p.Image,
		Code:          p.Code,
		Size:          p.Size,
		Type:          p.Type,
		IsShelves:     p.IsShelves,
		FounderID:     p.FounderID,
	}

	res, err := productDao.Create(&productModel)

	if err != nil {
		srv.logger.Error("创建产品失败", zap.Error(err))
		return nil, err
	}

	resModel, err := srv.getUserByFounder(res)

	if err != nil {
		srv.logger.Error("获取创建人失败", zap.Error(err))
		return nil, err
	}

	return resModel, nil
}

func (srv ProductService) Update(p *product.UpdatePayload) (*model.Product, error) {
	productDao := dao.NewProductDao(srv.ctx, srv.db)
	scopes := dao.NewProductScope()
	scopes = scopes.WithPk(p.ID)

	fields := make(map[string]interface{})
	if p.Name != nil && *p.Name != "" {
		fields["name"] = *p.Name
	}
	if p.Unit != nil && *p.Unit != 0 {
		fields["unit"] = *p.Unit
	}
	if p.CostPrice != nil {
		fields["cost_price"] = *p.CostPrice
	}
	if p.MarketPrice != nil {
		fields["market_price"] = *p.MarketPrice
	}
	if p.Note != nil {
		fields["note"] = *p.Note
	}
	if p.Image != nil {
		fields["image"] = *p.Image
	}
	if p.Size != nil && *p.Size != "" {
		fields["size"] = *p.Size
	}
	if p.Type != nil && *p.Type != 0 {
		fields["type"] = *p.Type
	}
	if p.IsShelves != nil {
		fields["is_shelves"] = *p.IsShelves
	}

	res, err := productDao.Update(scopes, fields)

	if err != nil {
		srv.logger.Error("更新产品失败", zap.Error(err))
		return nil, err
	}

	resModel, err := srv.getUserByFounder(res)

	if err != nil {
		srv.logger.Error("获取创建人失败", zap.Error(err))
		return nil, err
	}

	return resModel, nil
}

func (srv ProductService) Delete(pks []string) error {
	productDao := dao.NewProductDao(srv.ctx, srv.db)

	err := productDao.Delete(pks)

	return err
}

func (srv ProductService) getUserByFounder(product *model.Product) (*model.Product, error) {
	userDao := dao.NewUserDao(srv.ctx, srv.db)

	userModel, err := userDao.GetByID(product.FounderID)

	if err != nil {
		srv.logger.Error("获取创建用户失败", zap.Error(err))
		return nil, err
	}

	product.Founder.ID = userModel.ID
	product.Founder.Name = userModel.Name

	return product, nil
}
