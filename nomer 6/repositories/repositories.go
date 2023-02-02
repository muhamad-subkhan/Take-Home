package repositories

import (
	"ShoppingCart/models"

	"gorm.io/gorm"
)

type Repositories interface {
	Create(entity models.Product) (models.Product, error)
	GetProduk(id int) (models.Product, error)
	UpdateProduk(product models.Product) (models.Product, error)
	DeleteProduk(product models.Product) (models.Product, error)
	FindProduk() ([]models.Product, error)
	FilterByName(name string) ([]models.Product, error)
	FilterByQty(qty int) ([]models.Product, error)
}

type Repository struct {
		db *gorm.DB
	
}

func Repo(db *gorm.DB) *Repository {
	return &Repository{db}
}


func (r *Repository) Create(entity models.Product) (models.Product, error) {
	err := r.db.Create(&entity).Error
	
	return entity, err
}

func (r *Repository) GetProduk(id int) (models.Product, error) {
	var product models.Product
	err := r.db.First(&product, id).Error


	return product, err
}


func (r *Repository) UpdateProduk(product models.Product) (models.Product, error) {
	err := r.db.Save(&product).Error

	return product, err
}

func (r *Repository) DeleteProduk(product models.Product) (models.Product, error) {
	err := r.db.Delete(&product).Error

	return product, err
}

func (r *Repository) FindProduk() ([]models.Product, error) {
	var Product []models.Product
	
	err := r.db.Find(Product).Error

	return Product, err

}

func (r *Repository) FilterByName(name string) ([]models.Product, error) {
	var products []models.Product
	err :=  r.db.Where("name_product = ?", name).Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil


}

func (r *Repository) FilterByQty(qty int) ([]models.Product, error) {
	var products []models.Product
	err :=  r.db.Where("qty_product = ?", qty).Find(&products).Error
	if err != nil {
		return nil, err
	}
	return products, nil


}
