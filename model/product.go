package model

type Product struct {
	Id             int              `gorm:"unique;primaryKey;column:id" json:"id"`
	NameProduct    string           `gorm:"not null;type:varchar(255)" json:"name_product"`
	Price          int              `gorm:"not null" json:"price"`
	Review_product []Review_product `gorm:"foreignKey:ProductId" json:"product_review"`
}
