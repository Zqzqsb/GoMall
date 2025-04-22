package mysql

import (
	"errors"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"zqzqsb/gomall/app/product/biz/model"
	product "zqzqsb/gomall/app/product/kitex_gen/product"
)

// CreateProduct 创建商品
func CreateProduct(db *gorm.DB, p *model.Product) (int64, error) {
	now := time.Now()
	p.CreatedAt = now
	p.UpdatedAt = now

	result := db.Create(p)
	if result.Error != nil {
		return 0, result.Error
	}
	return p.ID, nil
}

// GetProductByID 根据ID获取商品
func GetProductByID(db *gorm.DB, id int64) (*model.Product, error) {
	var product model.Product
	result := db.First(&product, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("product not found")
		}
		return nil, result.Error
	}
	return &product, nil
}

// UpdateProduct 更新商品信息
func UpdateProduct(db *gorm.DB, p *model.Product) error {
	p.UpdatedAt = time.Now()
	result := db.Save(p)
	return result.Error
}

// DeleteProduct 删除商品
func DeleteProduct(db *gorm.DB, id int64) error {
	result := db.Delete(&model.Product{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("product not found")
	}
	return nil
}

// ListProducts 获取商品列表
func ListProducts(db *gorm.DB, req *product.ListProductsReq) ([]*model.Product, int64, error) {
	var products []*model.Product
	var count int64

	query := db

	// 应用筛选条件
	if req.Category != "" {
		query = query.Where("category = ?", req.Category)
	}

	if req.Keyword != "" {
		query = query.Where("name LIKE ? OR description LIKE ?", "%"+req.Keyword+"%", "%"+req.Keyword+"%")
	}

	if req.OnSaleOnly {
		query = query.Where("is_on_sale = ?", true)
	}

	if req.MinPrice > 0 {
		query = query.Where("price >= ?", req.MinPrice)
	}

	if req.MaxPrice > 0 {
		query = query.Where("price <= ?", req.MaxPrice)
	}

	// 获取总数
	if err := query.Model(&model.Product{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}

	// 应用排序
	switch req.SortBy {
	case "price":
		if req.Ascending {
			query = query.Order("price ASC")
		} else {
			query = query.Order("price DESC")
		}
	case "create_time":
		if req.Ascending {
			query = query.Order("created_at ASC")
		} else {
			query = query.Order("created_at DESC")
		}
	case "sales":
		if req.Ascending {
			query = query.Order("sales_count ASC")
		} else {
			query = query.Order("sales_count DESC")
		}
	case "rating":
		if req.Ascending {
			query = query.Order("rating ASC")
		} else {
			query = query.Order("rating DESC")
		}
	default:
		query = query.Order("id DESC") // 默认按ID降序排序
	}

	// 分页
	pageSize := int(req.PageSize)
	if pageSize <= 0 {
		pageSize = 10 // 默认每页10条
	}

	page := int(req.Page)
	if page <= 0 {
		page = 1 // 默认第1页
	}

	offset := (page - 1) * pageSize
	query = query.Offset(offset).Limit(pageSize)

	// 执行查询
	if err := query.Find(&products).Error; err != nil {
		return nil, 0, err
	}

	return products, count, nil
}

// GetCategories 获取所有商品分类
func GetCategories(db *gorm.DB) ([]string, error) {
	var categories []string
	result := db.Model(&model.Product{}).Distinct().Pluck("category", &categories)
	return categories, result.Error
}

// UpdateStock 更新商品库存
func UpdateStock(db *gorm.DB, productID int64, quantity int32) (int32, error) {
	var product model.Product
	
	// 使用事务确保库存更新的原子性
	err := db.Transaction(func(tx *gorm.DB) error {
		// 查询商品并锁定行
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).First(&product, productID).Error; err != nil {
			return err
		}

		// 检查库存是否足够（如果是减少库存的操作）
		if quantity < 0 && product.Stock < -quantity {
			return errors.New("insufficient stock")
		}

		// 更新库存
		product.Stock += quantity
		product.UpdatedAt = time.Now()
		
		// 如果是减少库存，增加销量
		if quantity < 0 {
			product.SalesCount += -quantity
		}

		return tx.Save(&product).Error
	})

	if err != nil {
		return 0, err
	}

	return product.Stock, nil
}
