package model

type Member struct {
	Id             int              `gorm:"primaryKey" json:"id"`
	Username       string           `gorm:"not null;unique;type:varchar(255)" json:"username"`
	Gender         string           `gorm:"not null;type:varchar(255)" json:"gender"`
	SkinType       string           `gorm:"not null;type:varchar(255)" json:"skin_type"`
	SkinColor      string           `gorm:"not null;type:varchar(255)" json:"skin_color"`
	Review_product []Review_product `gorm:"foreignKey:MemberId" json:"member_review"`
	Like_review    []Like_review    `gorm:"foreignKey:MemberId" json:"product_review"`
}
