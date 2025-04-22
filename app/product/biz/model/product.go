package model

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

// Product 商品模型
type Product struct {
	ID          int64          `gorm:"primarykey"`
	Name        string         `gorm:"type:varchar(255);not null"`
	Description string         `gorm:"type:text"`
	Price       int64          `gorm:"not null"` // 单位：分
	Stock       int32          `gorm:"not null"`
	ImageURL    string         `gorm:"type:varchar(255)"`
	Gallery     string         `gorm:"type:text"` // JSON 格式存储图片集
	Category    string         `gorm:"type:varchar(100);index"`
	IsOnSale    bool           `gorm:"default:true"`
	Attributes  string         `gorm:"type:text"` // JSON 格式存储属性
	Rating      float32        `gorm:"default:5.0"`
	SalesCount  int32          `gorm:"default:0"`
	CreatedAt   time.Time      `gorm:"not null"`
	UpdatedAt   time.Time      `gorm:"not null"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

// TableName 设置表名
func (Product) TableName() string {
	return "products"
}

// GetGallery 获取图片集
func (p *Product) GetGallery() []string {
	var gallery []string
	if p.Gallery != "" {
		_ = json.Unmarshal([]byte(p.Gallery), &gallery)
	}
	return gallery
}

// SetGallery 设置图片集
func (p *Product) SetGallery(gallery []string) error {
	data, err := json.Marshal(gallery)
	if err != nil {
		return err
	}
	p.Gallery = string(data)
	return nil
}

// GetAttributes 获取属性
func (p *Product) GetAttributes() map[string]string {
	var attributes map[string]string
	if p.Attributes != "" {
		_ = json.Unmarshal([]byte(p.Attributes), &attributes)
	}
	if attributes == nil {
		attributes = make(map[string]string)
	}
	return attributes
}

// SetAttributes 设置属性
func (p *Product) SetAttributes(attributes map[string]string) error {
	data, err := json.Marshal(attributes)
	if err != nil {
		return err
	}
	p.Attributes = string(data)
	return nil
}
