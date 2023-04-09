package model

type Review_product struct {
	Id          int           `gorm:"primaryKey" json:"id"`
	ProductId   int           `json:"id_product"`
	MemberId    int           `json:"id_member"`
	DescReview  string        `gorm:"type:varchar(255)" json:"desc_review"`
	Member      Member        `gorm:"foreignKey:MemberId" json:"member"`
	Like_review []Like_review `gorm:"foreignKey:ReviewId" json:"jumlah_like_review"`
}
