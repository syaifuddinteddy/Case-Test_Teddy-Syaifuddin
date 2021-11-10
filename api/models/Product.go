package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Product struct {
	ID          uint64    `gorm:"primary_key;auto_increment" json:"id"`
	SKU         string    `gorm:"size:255;not null;unique" json:"sku"`
	Name        string    `gorm:"size:255;not null;unique" json:"name"`
	ImgString   string    `gorm:"size:255;not null;" json:"imgstring"`
	Description string    `gorm:"size:255;not null;" json:"description"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (p *Product) Prepare() {
	p.ID = 0
	p.SKU = html.EscapeString(strings.TrimSpace(p.SKU))
	p.Name = html.EscapeString(strings.TrimSpace(p.Name))
	p.ImgString = html.EscapeString(strings.TrimSpace(p.ImgString))
	p.Description = html.EscapeString(strings.TrimSpace(p.Description))
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
}

func (p *Product) Validate() error {

	if p.SKU == "" {
		return errors.New("Required SKU")
	}
	if p.Name == "" {
		return errors.New("Required Name")
	}
	if p.Description == "" {
		return errors.New("Required Description")
	}
	return nil
}

func (p *Product) SaveProduct(db *gorm.DB) (*Product, error) {
	var err error
	err = db.Debug().Model(&Product{}).Create(&p).Error
	if err != nil {
		return &Product{}, err
	}
	return p, nil
}

func (p *Product) FindAllProducts(db *gorm.DB) (*[]Product, error) {
	var err error
	products := []Product{}
	err = db.Debug().Model(&Product{}).Limit(100).Find(&products).Error
	if err != nil {
		return &[]Product{}, err
	}
	return &products, nil
}

func (p *Product) FindProductByID(db *gorm.DB, pid uint64) (*Product, error) {
	var err error
	err = db.Debug().Model(&Product{}).Where("id = ?", pid).Take(&p).Error
	if err != nil {
		return &Product{}, err
	}
	return p, nil
}

func (p *Product) UpdateAProduct(db *gorm.DB) (*Product, error) {

	var err error

	err = db.Debug().Model(&Product{}).Where("id = ?", p.ID).Updates(Product{Name: p.Name, Description: p.Description, UpdatedAt: time.Now()}).Error
	if err != nil {
		return &Product{}, err
	}
	return p, nil
}

func (p *Product) DeleteAProduct(db *gorm.DB, pid uint64) (int64, error) {

	db = db.Debug().Model(&Product{}).Where("id = ?", pid).Take(&Product{}).Delete(&Product{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("Product not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
